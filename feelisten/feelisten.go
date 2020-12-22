package feelisten

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/prometheus/common/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"poly-swap/conf"
	"poly-swap/models"
	"runtime/debug"
	"time"
)

type FeeListen struct {
	ethCfg   *conf.EthereumFeeListenConfig
	neoCfg   *conf.NeoFeeListenConfig
	bscCfg   *conf.BscFeeListenConfig
	dbCfg    *conf.DBConfig
	feeUpdateSlot int64
	db       *gorm.DB
}

func StartFeeListen(cfg *conf.FeeListenConfig, dbCfg *conf.DBConfig) {
	feeListen := NewFeeListen(cfg.EthereumFeeListenConfig, cfg.NeoFeeListenConfig, cfg.BscFeeListenConfig, cfg.FeeUpdateSlot, dbCfg)
	feeListen.Start()
}

func NewFeeListen(ethCfg *conf.EthereumFeeListenConfig, neoCfg *conf.NeoFeeListenConfig, bscCfg *conf.BscFeeListenConfig, feeUpdateSlot int64, dbCfg *conf.DBConfig) *FeeListen {
	feeListen := &FeeListen{}
	feeListen.ethCfg = ethCfg
	feeListen.neoCfg = neoCfg
	feeListen.bscCfg = bscCfg
	feeListen.dbCfg = dbCfg
	feeListen.feeUpdateSlot = feeUpdateSlot
	db, err := gorm.Open(mysql.Open(dbCfg.User + ":" + dbCfg.Password + "@tcp(" + dbCfg.URL + ")/" +
		dbCfg.Scheme + "?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	feeListen.db = db
	//
	chainFees := make([]*models.ChainFee, 0)
	res := db.Find(&chainFees)
	if res.RowsAffected == 0 {
		panic("there is no token basic!")
	}
	err = feeListen.getChainFee(chainFees)
	if err != nil {
		panic(err)
	}
	db.Save(chainFees)
	return feeListen
}


func (this *FeeListen) Start() {
	go this.ListenFee()
}

func (this *FeeListen) ListenFee() {
	for {
		this.listenFee()
	}
}

func (this *FeeListen) listenFee() {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("service start, recover info: %s", string(debug.Stack()))
		}
	}()

	logs.Debug("listen fee......")
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			now := time.Now().Unix() / 60
			if now % this.feeUpdateSlot != 0 {
				continue
			}
			log.Infof("do fee update at time: %s", time.Now().Format("2006-01-02 15:04:05"))
			chainFees := make([]*models.ChainFee, 0)
			res := this.db.Find(&chainFees)
			if res.RowsAffected == 0 {
				continue
			}
			err := this.getChainFee(chainFees)
			if err != nil {
				continue
			}
			this.db.Save(chainFees)
		}
	}
}

func (this *FeeListen) getChainFee(chainFees []*models.ChainFee) error {
	chainName2Item := make(map[uint64]*models.ChainFee, 0)
	for _, item := range chainFees {
		chainName2Item[item.ChainId] = item
	}
	//
	maxFee, minFee, err1 := this.getEthFee()
	chainFee, ok := chainName2Item[conf.ETHEREUM_CROSSCHAIN_ID]
	if err1 == nil && ok {
		chainFee.MaxFee = maxFee
		chainFee.MinFee = minFee
		chainFee.ProxyFee = minFee * this.ethCfg.ProxyFee / 100
	} else {
		log.Errorf("get eth fee err: %v", err1)
	}
	//
	maxFee, minFee, err2 := this.getNeoFee()
	chainFee, ok = chainName2Item[conf.NEO_CROSSCHAIN_ID]
	if err2 == nil && ok {
		chainFee.MaxFee = maxFee
		chainFee.MinFee = minFee
		chainFee.ProxyFee = minFee * this.neoCfg.ProxyFee / 100
	} else {
		log.Errorf("get neo fee err: %v", err2)
	}
	//
	maxFee, minFee, err3 := this.getBscFee()
	chainFee, ok = chainName2Item[conf.BSC_CROSSCHAIN_ID]
	if err2 == nil && ok {
		chainFee.MaxFee = maxFee
		chainFee.MinFee = minFee
		chainFee.ProxyFee = minFee * this.bscCfg.ProxyFee / 100
	} else {
		log.Errorf("get bsc fee err: %v", err2)
	}
	if err1 != nil || err2 != nil || err3 != nil {
		return fmt.Errorf("can not get the fee information of all chains!")
	}
	return nil
}

func (this *FeeListen) getEthFee() (uint64, uint64, error) {
	return 1000000000, 1000000000, nil
}

func (this *FeeListen) getNeoFee() (uint64, uint64, error) {
	return 1000000000, 1000000000, nil
}

func (this *FeeListen) getBscFee() (uint64, uint64, error) {
	return 1000000000, 1000000000, nil
}


