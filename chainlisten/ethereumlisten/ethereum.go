package ethereumlisten

import (
	"context"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"poly-swap/chainlisten/ethereumlisten/lock_proxy_abi"
	"poly-swap/conf"
	"poly-swap/models"
	"runtime/debug"
	"time"
)

type EthereumChainListen struct {
	ethCfg *conf.EthereumChainListenConfig
	dbCfg  *conf.DBConfig
	db     *gorm.DB
	ethSdk *EthereumSdk
	chain  *models.Chain
}

func NewEthereumChainListen(cfg *conf.EthereumChainListenConfig, dbCfg *conf.DBConfig) *EthereumChainListen {
	ethListen := &EthereumChainListen{}
	ethListen.ethCfg = cfg
	ethListen.dbCfg = dbCfg
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	ethListen.db = db
	//
	sdk, err := NewEthereumSdk(cfg.RestURL)
	if err != nil {
		panic(err)
	}
	ethListen.ethSdk = sdk
	//
	chain := new(models.Chain)
	res := db.Where("chain_id = ?", conf.ETHEREUM_CROSSCHAIN_ID).First(chain)
	if res.RowsAffected == 0 {
		panic("there is no ethereum!")
	}
	height, err := sdk.GetCurrentBlockHeight()
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

func (this *EthereumChainListen) Start() {
	go this.ListenChain()
}

func (this *EthereumChainListen) ListenChain() {
	for {
		this.listenChain()
	}
}

func (this *EthereumChainListen) listenChain() {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("service start, recover info: %s", string(debug.Stack()))
		}
	}()

	logs.Debug("listen ethereum chain......")
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			height, err := this.ethSdk.GetCurrentBlockHeight()
			if err != nil || height == 0 {
				logs.Error("ListenChain - cannot get height, err: %s", err)
				continue
			}
			if this.chain.Height >= height-1 {
				continue
			}
			logs.Info("ListenChain - ethereum latest height is %d, parser height: %d", height, this.chain.Height)
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

func (this *EthereumChainListen) HandleNewBlock(height uint64) error {
	blockHeader, err := this.ethSdk.GetHeaderByNumber(height)
	if err != nil {
		return err
	}
	tt := blockHeader.Time
	ethCrossChainTxs, err := this.getProxyEventByBlockNumber(this.ethCfg.Contract, height)
	if err != nil {
		return err
	}
	for _, item := range ethCrossChainTxs {
		item.Time = tt
		item.BlockHeight = height
	}
	if len(ethCrossChainTxs) > 0 {
		this.db.Save(ethCrossChainTxs)
	}
	this.chain.Height = height
	this.db.Save(this.chain)
	return nil
}

func (this *EthereumChainListen) getProxyEventByBlockNumber(contractAddr string, height uint64) ([]*models.Transaction, error) {
	proxyAddress := common.HexToAddress(contractAddr)
	lockContract, err := lock_proxy_abi.NewLockProxy(proxyAddress, this.ethSdk.rawClient)
	if err != nil {
		return nil, fmt.Errorf("GetSmartContractEventByBlock, error: %s", err.Error())
	}
	opt := &bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: context.Background(),
	}
	// get ethereum lock events from given block
	ethCrossChainTxs := make([]*models.Transaction, 0)
	lockEvents, err := lockContract.FilterLockEvent(opt)
	if err != nil {
		return nil, fmt.Errorf("GetSmartContractEventByBlock, filter lock events :%s", err.Error())
	}
	for lockEvents.Next() {
		evt := lockEvents.Event
		ethCrossChainTxs = append(ethCrossChainTxs, &models.Transaction{
			Hash:         evt.Raw.TxHash.String()[2:],
			User:         evt.FromAddress.String(),
			SrcChainId:   uint64(evt.ToChainId),
			DstChainId:   uint64(evt.ToChainId),
			FeeTokenHash: evt.Raw.TxHash.String()[2:],
			FeeAmount:    evt.Amount.Uint64(),
		})
	}
	return ethCrossChainTxs, nil
}
