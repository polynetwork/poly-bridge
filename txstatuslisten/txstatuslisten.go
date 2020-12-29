package txstatuslisten

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"poly-swap/conf"
	"poly-swap/models"
	"runtime/debug"
	"time"
)

type TxStatusListen struct {
	txStatusCfg *conf.TxStatusListenConfig
	dbCfg       *conf.DBConfig
	db          *gorm.DB
	sdk         *PolyExplorerSdk
}

func StartTxStatusListen(txStatusCfg *conf.TxStatusListenConfig, dbCfg *conf.DBConfig) {
	txStatusListen := NewTxStatusListen(txStatusCfg, dbCfg)
	txStatusListen.Start()
}

func NewTxStatusListen(txStatusCfg *conf.TxStatusListenConfig, dbCfg *conf.DBConfig) *TxStatusListen {
	txStatusListen := &TxStatusListen{}
	txStatusListen.txStatusCfg = txStatusCfg
	txStatusListen.dbCfg = dbCfg
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	txStatusListen.db = db
	//
	sdk := NewPolyExplorerSdk(txStatusCfg.RestURL)
	txStatusListen.sdk = sdk
	return txStatusListen
}

func (this *TxStatusListen) Start() {
	go this.ListenTxStatus()
}

func (this *TxStatusListen) ListenTxStatus() {
	for {
		this.listenTxStatus()
	}
}

func (this *TxStatusListen) listenTxStatus() {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("service start, recover info: %s", string(debug.Stack()))
		}
	}()

	logs.Debug("listen tx status......")
	ticker := time.NewTicker(time.Second * time.Duration(this.txStatusCfg.UpdateSlot))
	for {
		select {
		case <-ticker.C:
			unfinishedTxs := make([]*models.WrapperTransaction, 0)
			res := this.db.Where("status != ?", conf.FINISHED).Find(&unfinishedTxs)
			if res.RowsAffected == 0 {
				continue
			}
			err := this.updateTxStatus(unfinishedTxs)
			if err != nil {
				logs.Error("update tx status err: %v", err)
				continue
			}
			this.db.Save(unfinishedTxs)
		}
	}
}

func (this *TxStatusListen) updateTxStatus(unfinishedTxs []*models.WrapperTransaction) error {
	for _, tx := range unfinishedTxs {
		status, err := this.getTxStatus(tx.Hash)
		if err != nil {
			logs.Error("get tx %s status err: %v", tx, err)
			continue
		}
		tx.Status = status
	}
	return nil
}

func (this *TxStatusListen) getTxStatus(txHash string) (uint64, error) {
	rsp, err := this.sdk.TxStatus(txHash)
	if err != nil {
		return 0, err
	}
	if rsp == nil {
		return 0, fmt.Errorf("can not get tx %s status!", txHash)
	}
	tx := rsp.Body.Result
	if tx.Fchaintx_valid == false || tx.Fchaintx == nil {
		return conf.PENDDING, nil
	}
	if tx.Mchaintx_valid == false || tx.Mchaintx == nil {
		if uint64(tx.Fchaintx.ChainId) == conf.ETHEREUM_CROSSCHAIN_ID {
			if uint64(tx.FchainHeight-tx.Fchaintx.Height) < this.txStatusCfg.EthereumConfirmed {
				return conf.SOURCE_DONE, nil
			} else {
				return conf.SOURCE_CONFIRMED, nil
			}
		} else if uint64(tx.Fchaintx.ChainId) == conf.NEO_CROSSCHAIN_ID {
			if uint64(tx.FchainHeight-tx.Fchaintx.Height) < this.txStatusCfg.NeoConfirmed {
				return conf.SOURCE_DONE, nil
			} else {
				return conf.SOURCE_CONFIRMED, nil
			}
		} else if uint64(tx.Fchaintx.ChainId) == conf.BSC_CROSSCHAIN_ID {
			if uint64(tx.FchainHeight-tx.Fchaintx.Height) < this.txStatusCfg.BscConfirmed {
				return conf.SOURCE_DONE, nil
			} else {
				return conf.SOURCE_CONFIRMED, nil
			}
		}
	}
	if tx.Tchaintx_valid == false || tx.Tchaintx == nil {
		return conf.POLY_CONFIRMED, nil
	}
	return conf.FINISHED, nil
}
