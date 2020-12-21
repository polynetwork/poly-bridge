package gaslisten

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

type GasListen struct {
	ethCfg   *conf.EthereumGasListenConfig
	neoCfg   *conf.NeoGasListenConfig
	bscCfg   *conf.BscGasListenConfig
	dbCfg    *conf.DBConfig
	gasUpdateSlot int64
	db       *gorm.DB
}

func StartGasListen(cfg *conf.GasListenConfig, dbCfg *conf.DBConfig) {
	gasListen := NewGasListen(cfg.EthereumGasListenConfig, cfg.NeoGasListenConfig, cfg.BscGasListenConfig, cfg.GasUpdateSlot, dbCfg)
	gasListen.Start()
}

func NewGasListen(ethCfg *conf.EthereumGasListenConfig, neoCfg *conf.NeoGasListenConfig, bscCfg *conf.BscGasListenConfig, gasUpdateSlot int64, dbCfg *conf.DBConfig) *GasListen {
	gasListen := &GasListen{}
	gasListen.ethCfg = ethCfg
	gasListen.neoCfg = neoCfg
	gasListen.bscCfg = bscCfg
	gasListen.dbCfg = dbCfg
	gasListen.gasUpdateSlot = gasUpdateSlot
	db, err := gorm.Open(mysql.Open(dbCfg.User + ":" + dbCfg.Password + "@tcp(" + dbCfg.URL + ")/" +
		dbCfg.Scheme + "?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	gasListen.db = db
	//
	chainGases := make([]*models.ChainGas, 0)
	res := db.Find(&chainGases)
	if res.RowsAffected == 0 {
		panic("there is no token basic!")
	}
	err = gasListen.getChainGas(chainGases)
	if err != nil {
		panic(err)
	}
	db.Save(chainGases)
	return gasListen
}


func (this *GasListen) Start() {
	go this.ListenGas()
}

func (this *GasListen) ListenGas() {
	for {
		this.listenGas()
	}
}

func (this *GasListen) listenGas() {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("service start, recover info: %s", string(debug.Stack()))
		}
	}()

	logs.Debug("listen gas......")
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			now := time.Now().Unix() / 60
			if now % this.gasUpdateSlot != 0 {
				continue
			}
			log.Infof("do gas update at time: %s", time.Now().Format("2006-01-02 15:04:05"))
			chainGases := make([]*models.ChainGas, 0)
			res := this.db.Find(&chainGases)
			if res.RowsAffected == 0 {
				continue
			}
			err := this.getChainGas(chainGases)
			if err != nil {
				continue
			}
			this.db.Save(chainGases)
		}
	}
}

func (this *GasListen) getChainGas(chainGases []*models.ChainGas) error {
	chainName2Item := make(map[uint64]*models.ChainGas, 0)
	for _, item := range chainGases {
		chainName2Item[item.ChainId] = item
	}
	//
	maxGas, minGas, err1 := this.getEthGas()
	chainGas, ok := chainName2Item[conf.ETHEREUM_CROSSCHAIN_ID]
	if err1 == nil && ok {
		chainGas.MaxFee = maxGas
		chainGas.MinFee = minGas
		chainGas.ProxyFee = minGas * this.ethCfg.ProxyFee / 100
	} else {
		log.Errorf("get eth gas err: %v", err1)
	}
	//
	maxGas, minGas, err2 := this.getNeoGas()
	chainGas, ok = chainName2Item[conf.NEO_CROSSCHAIN_ID]
	if err2 == nil && ok {
		chainGas.MaxFee = maxGas
		chainGas.MinFee = minGas
		chainGas.ProxyFee = minGas * this.neoCfg.ProxyFee / 100
	} else {
		log.Errorf("get neo gas err: %v", err2)
	}
	//
	maxGas, minGas, err3 := this.getBscGas()
	chainGas, ok = chainName2Item[conf.BSC_CROSSCHAIN_ID]
	if err2 == nil && ok {
		chainGas.MaxFee = maxGas
		chainGas.MinFee = minGas
		chainGas.ProxyFee = minGas * this.bscCfg.ProxyFee / 100
	} else {
		log.Errorf("get bsc gas err: %v", err2)
	}
	if err1 != nil || err2 != nil || err3 != nil {
		return fmt.Errorf("can not get the gas information of all chains!")
	}
	return nil
}

func (this *GasListen) getEthGas() (uint64, uint64, error) {
	return 1000000000, 1000000000, nil
}

func (this *GasListen) getNeoGas() (uint64, uint64, error) {
	return 1000000000, 1000000000, nil
}

func (this *GasListen) getBscGas() (uint64, uint64, error) {
	return 1000000000, 1000000000, nil
}


