package crosschainlisten

import (
	"context"
	"github.com/beego/beego/v2/core/logs"
	"github.com/devfans/cogroup"
	"poly-bridge/conf"
	"poly-bridge/crosschaindao/bridgedao"
	"time"
)

var handlerMap = make(map[uint64]ChainHandle, 0)

func StartCrossChainListenPatch(config *conf.Config) {
	dao := bridgedao.NewBridgeDao(config.DBConfig, config.Backup)
	if dao == nil {
		panic("NewBridgeDao err")
	}
	for _, cfg := range config.ChainListenConfig {
		handlerMap[cfg.ChainId] = NewChainHandle(cfg)
	}
	go startPatchWrapperMissingTx(dao)
}

func startPatchWrapperMissingTx(dao *bridgedao.BridgeDao) {
	ticker := time.NewTicker(time.Minute * 10)
	for {
		select {
		case <-ticker.C:
			go patchWrapperMissingTx(dao)
		}
	}
}

func patchWrapperMissingTx(dao *bridgedao.BridgeDao) {
	txs, err := dao.FilterMissingWrapperTransactions()
	if err != nil {
		logs.Error("filterMissingWrapperTransactions err:" + err.Error())
	}
	logs.Info("find %d TXs missing wrapper event", len(txs))
	for _, tx := range txs {
		logs.Info("srcTransactions hash: %s missing wrapper_transactions", tx.Hash)
		handler := handlerMap[tx.ChainId]
		if handler == nil {
			logs.Error("handle of chain id: %d is nil", tx.ChainId)
			continue
		}
		g := cogroup.Start(context.Background(), 4, 8, false)

		g.Insert(retry(func() error {
			wrapperTransactions, srcTransactions, polyTransactions, dstTransactions, wrapperDetails, polyDetails, locks, unlocks, err := handler.HandleNewBlock(tx.Height)
			if err != nil {
				logs.Error("HandleNewBlock chain：%s, height: %d err: %v", handler.GetChainName(), tx.Height, err)
				return err
			}
			logs.Info("Fetch block events success chain name: %s, height: %d, wrapper: %d, src: %d, poly: %d, dst: %d, locks: %d, unlocks: %d",
				handler.GetChainName(), tx.Height, len(wrapperTransactions), len(srcTransactions), len(polyTransactions), len(dstTransactions), locks, unlocks)
			logs.Info("wrapper : %+v", wrapperTransactions)
			logs.Info("src : %+v", srcTransactions)
			logs.Info("poly : %+v", polyTransactions)
			logs.Info("dst : %+v", dstTransactions)
			err = dao.UpdateEvents(wrapperTransactions, srcTransactions, polyTransactions, dstTransactions, wrapperDetails, polyDetails)
			if err != nil {
				logs.Error("UpdateEvents chain：%s, height: %d err: %v", handler.GetChainName(), tx.Height, err)
				return err
			}
			return nil
		}, 2, 5*time.Second))
		g.Wait()
	}
}

func retry(f func() error, count int, duration time.Duration) func(context.Context) error {
	return func(context.Context) error {
		i := 0
		for {
			i++
			if i > count && count != 0 {
				return nil
			}
			err := f()
			if err == nil {
				return nil
			}
			time.Sleep(duration)
		}
	}
}
