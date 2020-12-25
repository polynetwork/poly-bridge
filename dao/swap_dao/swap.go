package swap_dao

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"gorm.io/gorm"
	"poly-swap/conf"
	"poly-swap/models"
	"gorm.io/driver/mysql"
	"runtime/debug"
	"time"
)

type SwapDao struct {
	dbCfg  *conf.DBConfig
	db     *gorm.DB
}

func NewSwapDao(dbCfg *conf.DBConfig) *SwapDao {
	swapDao := &SwapDao{
		dbCfg: dbCfg,
	}
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	swapDao.db = db
	return swapDao
}

func (dao *SwapDao) UpdateEvents(chain *models.Chain, wrapperTransactions []*models.WrapperTransaction, srcTransactions []*models.SrcTransaction, polyTransactions []*models.PolyTransaction, dstTransactions []*models.DstTransaction) error {
	if wrapperTransactions != nil && len(wrapperTransactions) > 0 {
		res := dao.db.Save(wrapperTransactions)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return fmt.Errorf("update wrapper Transactions failed!")
		}
	}
	if srcTransactions != nil && len(srcTransactions) > 0 {
		res := dao.db.Save(srcTransactions)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return fmt.Errorf("update src Transactions failed!")
		}
	}
	if polyTransactions != nil && len(polyTransactions) > 0 {
		res := dao.db.Save(polyTransactions)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return fmt.Errorf("update poly Transactions failed!")
		}
	}
	if dstTransactions != nil && len(dstTransactions) > 0 {
		res := dao.db.Save(dstTransactions)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return fmt.Errorf("update dst Transactions failed!")
		}
	}
	if chain != nil {
		res := dao.db.Save(chain)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return fmt.Errorf("update chain failed!")
		}
	}
	return nil
}

func (dao *SwapDao) GetChain(chainId uint64) (*models.Chain, error) {
	chain := new(models.Chain)
	res := dao.db.Where("chain_id = ?", chainId).First(chain)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("no record!")
	}
	return chain, nil
}

func (dao *SwapDao) UpdateChain(chain *models.Chain) error {
	if chain == nil {
		return fmt.Errorf("no value!")
	}
	res := dao.db.Save(chain)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("no update!")
	}
	return nil
}

func (dao *SwapDao) Start() {
	go dao.Check()
}

func (dao *SwapDao) Check() {
	for {
		dao.check()
	}
}

func (dao *SwapDao) check() {
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
			err := dao.CheckHash()
			if err != nil {
				logs.Error("check - err: %s", err)
			}
		}
	}
}

func (dao *SwapDao) CheckHash() error {
	unUpdatePolyTransactions := make([]*models.PolyTransaction, 0)
	dao.db.Where("left(src_hash, 8) = ?", "00000000").Preload("SrcTransaction0").Find(&unUpdatePolyTransactions)
	updatePolyTransactions := make([]*models.PolyTransaction, 0)
	for _, unUpdatePolyTransaction := range unUpdatePolyTransactions {
		if unUpdatePolyTransaction.SrcTransaction0 != nil {
			unUpdatePolyTransaction.SrcHash = unUpdatePolyTransaction.SrcTransaction0.Hash
			unUpdatePolyTransaction.SrcTransaction0 = nil
			updatePolyTransactions = append(updatePolyTransactions, unUpdatePolyTransaction)
		}
	}
	if len(updatePolyTransactions) > 0 {
		dao.db.Save(updatePolyTransactions)
	}
	return nil
}
