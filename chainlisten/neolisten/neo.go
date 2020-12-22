package neolisten

import (
	"encoding/hex"
	"github.com/astaxie/beego/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"poly-swap/conf"
	"poly-swap/models"
	"runtime/debug"
	"strconv"
	"time"
)

const (
	_neo_crosschainlock = "CrossChainLockEvent"
)

type NeoChainListen struct {
	neoCfg *conf.NeoChainListenConfig
	dbCfg  *conf.DBConfig
	db     *gorm.DB
	neoSdk *NeoSdk
	chain  *models.Chain
}

func NewNeoChainListen(cfg *conf.NeoChainListenConfig, dbCfg *conf.DBConfig) *NeoChainListen {
	ethListen := &NeoChainListen{}
	ethListen.neoCfg = cfg
	ethListen.dbCfg = dbCfg
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	ethListen.db = db
	//
	sdk := NewNeoSdk(cfg.RestURL)
	ethListen.neoSdk = sdk
	//
	chain := new(models.Chain)
	res := db.Where("chain_id = ?", conf.NEO_CROSSCHAIN_ID).First(chain)
	if res.RowsAffected == 0 {
		panic("there is no neo!")
	}
	height, err := sdk.GetBlockCount()
	if err != nil || height == 0 {
		panic(err)
	}
	if chain.Height == 0 {
		chain.Height = height
	}
	db.Save(chain)
	ethListen.chain = chain
	return ethListen
}

func (this *NeoChainListen) Start() {
	go this.ListenChain()
}

func (this *NeoChainListen) ListenChain() {
	for {
		this.listenChain()
	}
}

func (this *NeoChainListen) listenChain() {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("service start, recover info: %s", string(debug.Stack()))
		}
	}()

	logs.Debug("listen neo chain......")
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			height, err := this.neoSdk.GetBlockCount()
			if err != nil || height == 0 {
				logs.Error("ListenChain - cannot get height, err: %s", err)
				continue
			}
			if this.chain.Height >= height-1 {
				continue
			}
			logs.Info("ListenChain - neo latest height is %d, parser height: %d", height, this.chain.Height)
			for this.chain.Height < height-1 {
				err := this.HandleNewBlock(this.chain.Height + 1)
				if err != nil {
					logs.Error("HandleNewBlock err: %v", err)
					break
				}
			}
		}
	}
}

func (this *NeoChainListen) parseNeoMethod(v string) string {
	xx, _ := hex.DecodeString(v)
	return string(xx)
}

func (this *NeoChainListen) HandleNewBlock(height uint64) error {
	block, err := this.neoSdk.GetBlockByIndex(height)
	if err != nil {
		return err
	}
	tt := block.Time
	neoCrossChainTxs := make([]*models.Transaction, 0)
	for _, tx := range block.Tx {
		if tx.Type != "InvocationTransaction" {
			continue
		}
		appLog, err := this.neoSdk.GetApplicationLog(tx.Txid)
		if err != nil {
			continue
		}
		for _, exeitem := range appLog.Executions {
			for _, notify := range exeitem.Notifications {
					if notify.Contract[2:] != this.neoCfg.Contract {
						continue
					}
					if len(notify.State.Value) <= 0 {
						continue
					}
					contractMethod := this.parseNeoMethod(notify.State.Value[0].Value)
					switch contractMethod {
					case _neo_crosschainlock:
						logs.Info("from chain: %s, txhash: %s\n", this.chain.Name, tx.Txid[2:])
						if len(notify.State.Value) < 6 {
							continue
						}
						xx, _ := strconv.ParseUint(notify.State.Value[3].Value, 10, 64)
						neoCrossChainTxs = append(neoCrossChainTxs, &models.Transaction{
							Hash:         tx.Txid[2:],
							User:         notify.State.Value[4].Value,
							SrcChainId:   xx,
							DstChainId:   xx,
							FeeTokenHash: notify.State.Value[4].Value,
							FeeAmount:    xx,
						})
				}
			}
		}
	}
	for _, item := range neoCrossChainTxs {
		item.Time = uint64(tt)
		item.BlockHeight = height
	}
	if len(neoCrossChainTxs) > 0 {
		this.db.Save(neoCrossChainTxs)
	}
	this.chain.Height = height
	this.db.Save(this.chain)
	return nil
}

