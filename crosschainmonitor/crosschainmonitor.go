package crosschainmonitor

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"poly-swap/conf"
	"poly-swap/crosschaindao/explorer_dao"
	"poly-swap/models"
	"runtime/debug"
	"time"
)

func StartCrossChainMonitor(monitorCfg *conf.CrossChainMonitor, dbCfg *conf.DBConfig) {
	monitor := NewCrossChainMonitor(monitorCfg, dbCfg)
	if monitor == nil {
		panic("monitor is invalid")
	}
	monitor.Start()
}

type CrossChainMonitor struct {
	monitorCfg *conf.CrossChainMonitor
	dbCfg *conf.DBConfig
	db    *gorm.DB
}

func NewCrossChainMonitor(monitorCfg *conf.CrossChainMonitor, dbCfg *conf.DBConfig) *CrossChainMonitor {
	crossChainMonitor := &CrossChainMonitor{
		dbCfg: dbCfg,
		monitorCfg: monitorCfg,
	}
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	crossChainMonitor.db = db
	return crossChainMonitor
}

func (dao *CrossChainMonitor) Start() {
	go dao.Check()
}

func (dao *CrossChainMonitor) Check() {
	for {
		dao.check()
	}
}

func (dao *CrossChainMonitor) check() {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("service start, recover info: %s", string(debug.Stack()))
		}
	}()
	logs.Debug("check events %s......")
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			err := dao.checkHash()
			if err != nil {
				logs.Error("check hash- err: %s", err)
			}
			err = dao.checkStatus()
			if err != nil {
				logs.Error("check status- err: %s", err)
			}
		}
	}
}

func (dao *CrossChainMonitor) checkHash() error {
	if dao.monitorCfg.Server == conf.SERVER_POLY_SWAP {
		polySrcRelations := make([]*models.PolySrcRelation, 0)
		dao.db.Debug().Table("poly_transactions").Where("left(poly_transactions.src_hash, 8) = ?", "00000000").Select("poly_transactions.hash as poly_hash, src_transactions.hash as src_hash").Joins("left join src_transactions on poly_transactions.src_hash = src_transactions.key").Preload("SrcTransaction").Preload("PolyTransaction").Find(&polySrcRelations)
		updatePolyTransactions := make([]*models.PolyTransaction, 0)
		for _, polySrcRelation := range polySrcRelations {
			if polySrcRelation.SrcTransaction != nil {
				polySrcRelation.PolyTransaction.SrcHash = polySrcRelation.SrcHash
				updatePolyTransactions = append(updatePolyTransactions, polySrcRelation.PolyTransaction)
			}
		}
		if len(updatePolyTransactions) > 0 {
			dao.db.Save(updatePolyTransactions)
		}
	} else if dao.monitorCfg.Server == conf.SERVER_EXPLORER {
		polySrcRelations := make([]*explorer_dao.PolySrcRelation, 0)
		dao.db.Debug().Table("poly_transactions").Where("left(poly_transactions.src_hash, 8) = ? and chain_id != ?", "00000000", conf.ETHEREUM_CROSSCHAIN_ID).Select("poly_transactions.hash as poly_hash, src_transactions.hash as src_hash").Joins("left join src_transactions on poly_transactions.src_hash = src_transactions.key").Preload("SrcTransaction").Preload("PolyTransaction").Find(&polySrcRelations)
		updatePolyTransactions := make([]*explorer_dao.PolyTransaction, 0)
		for _, polySrcRelation := range polySrcRelations {
			if polySrcRelation.SrcTransaction != nil {
				polySrcRelation.PolyTransaction.SrcHash = polySrcRelation.SrcHash
				updatePolyTransactions = append(updatePolyTransactions, polySrcRelation.PolyTransaction)
			}
		}
		if len(updatePolyTransactions) > 0 {
			dao.db.Save(updatePolyTransactions)
		}
	}
	return nil
}

func (dao *CrossChainMonitor) checkStatus() error {
	if dao.monitorCfg.Server != conf.SERVER_POLY_SWAP {
		return nil
	}
	wrapperTransactions := make([]*models.WrapperTransaction, 0)
	now := time.Now().Unix() - dao.monitorCfg.HowOld
	dao.db.Model(models.WrapperTransaction{}).Where("status != ? and time < ?", conf.FINISHED, now).Find(&wrapperTransactions)
	if len(wrapperTransactions) > 0 {
		wrapperTransactionsJson, _ := json.Marshal(wrapperTransactions)
		logs.Error("There is unfinished transactions %s", string(wrapperTransactionsJson))
	}
	return nil
}
