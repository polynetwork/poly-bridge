/*
 * Copyright (C) 2020 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package crosschainlisten

import (
	"fmt"
	"math"
	"poly-bridge/cacheRedis"
	"poly-bridge/common"
	"poly-bridge/crosschainlisten/starcoinlisten"
	"poly-bridge/crosschainlisten/zilliqalisten"
	"poly-bridge/utils/decimal"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/polynetwork/bridge-common/metrics"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/crosschaindao"
	"poly-bridge/crosschainlisten/ethereumlisten"
	"poly-bridge/crosschainlisten/neo3listen"
	"poly-bridge/crosschainlisten/neolisten"
	"poly-bridge/crosschainlisten/o3listen"
	"poly-bridge/crosschainlisten/ontologylisten"
	"poly-bridge/crosschainlisten/polylisten"
	"poly-bridge/crosschainlisten/switcheolisten"
	"poly-bridge/models"

	"github.com/beego/beego/v2/core/logs"
)

var chainListens = make([]*CrossChainListen, 0)

func StartCrossChainListen(config *conf.Config) {
	dao := crosschaindao.NewCrossChainDao(config.Server, config.Backup, config.DBConfig)
	if dao == nil {
		panic("server is not valid")
	}
	for _, cfg := range config.ChainListenConfig {
		chainHandle := NewChainHandle(cfg)
		if chainHandle == nil {
			logs.Error("chain %d handler is invalid", cfg.ChainId)
			continue
		}
		chainListen := NewCrossChainListen(chainHandle, dao, config)
		chainListen.Start()
		chainListens = append(chainListens, chainListen)
	}
}

func StopCrossChainListen() {
	for _, chainListen := range chainListens {
		if chainListen != nil {
			chainListen.Stop()
		}
	}
}

type ChainHandle interface {
	GetExtendLatestHeight() (uint64, error)
	GetLatestHeight() (uint64, error)
	HandleNewBlock(height uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, int, int, error)
	GetChainListenSlot() uint64
	GetChainId() uint64
	GetChainName() string
	GetDefer() uint64
	GetBatchSize() uint64
}

func NewChainHandle(chainListenConfig *conf.ChainListenConfig) ChainHandle {
	switch chainListenConfig.ChainId {
	case basedef.POLY_CROSSCHAIN_ID:
		return polylisten.NewPolyChainListen(chainListenConfig)
	case basedef.ETHEREUM_CROSSCHAIN_ID, basedef.BSC_CROSSCHAIN_ID, basedef.PLT_CROSSCHAIN_ID, basedef.OK_CROSSCHAIN_ID,
		basedef.HECO_CROSSCHAIN_ID, basedef.MATIC_CROSSCHAIN_ID, basedef.ARBITRUM_CROSSCHAIN_ID, basedef.XDAI_CROSSCHAIN_ID,
		basedef.FANTOM_CROSSCHAIN_ID, basedef.AVAX_CROSSCHAIN_ID, basedef.OPTIMISTIC_CROSSCHAIN_ID, basedef.METIS_CROSSCHAIN_ID,
		basedef.PIXIE_CROSSCHAIN_ID, basedef.RINKEBY_CROSSCHAIN_ID, basedef.BOBA_CROSSCHAIN_ID, basedef.OASIS_CROSSCHAIN_ID,
		basedef.HARMONY_CROSSCHAIN_ID, basedef.HSC_CROSSCHAIN_ID, basedef.BCSPALETTE_CROSSCHAIN_ID, basedef.BYTOM_CROSSCHAIN_ID,
		basedef.KCC_CROSSCHAIN_ID:
		return ethereumlisten.NewEthereumChainListen(chainListenConfig)
	case basedef.NEO_CROSSCHAIN_ID:
		return neolisten.NewNeoChainListen(chainListenConfig)
	case basedef.ONT_CROSSCHAIN_ID, basedef.ONTEVM_CROSSCHAIN_ID:
		return ontologylisten.NewOntologyChainListen(chainListenConfig)
	case basedef.O3_CROSSCHAIN_ID:
		return o3listen.NewO3ChainListen(chainListenConfig)
	case basedef.SWITCHEO_CROSSCHAIN_ID:
		return switcheolisten.NewSwitcheoChainListen(chainListenConfig)
	case basedef.NEO3_CROSSCHAIN_ID:
		return neo3listen.NewNeo3ChainListen(chainListenConfig)
	case basedef.ZILLIQA_CROSSCHAIN_ID:
		return zilliqalisten.NewZilliqaChainListen(chainListenConfig)
	case basedef.STARCOIN_CROSSCHAIN_ID:
		return starcoinlisten.NewStarcoinChainListen(chainListenConfig)

	default:
		return nil
	}
}

type CrossChainListen struct {
	handle  ChainHandle
	db      crosschaindao.CrossChainDao
	exit    chan bool
	height  uint64
	config  *conf.Config
	dingMux sync.Mutex
}

func NewCrossChainListen(handle ChainHandle, db crosschaindao.CrossChainDao, config *conf.Config) *CrossChainListen {
	crossChainListen := &CrossChainListen{
		handle: handle,
		db:     db,
		exit:   make(chan bool, 0),
		config: config,
	}
	return crossChainListen
}

func (ccl *CrossChainListen) SetHeight(height uint64) {
	ccl.height = height
}

func (ccl *CrossChainListen) Start() {
	if ccl.config.Backup && ccl.handle.GetChainId() == basedef.POLY_CROSSCHAIN_ID {
		return
	}
	logs.Info("start cross chain listen: %s", ccl.handle.GetChainName())
	go ccl.ListenChain()
}

func (ccl *CrossChainListen) Stop() {
	ccl.exit <- true
	logs.Info("stop cross chain listen: %s", ccl.handle.GetChainName())
}

func (ccl *CrossChainListen) ListenChain() {
	for {
		exit := ccl.listenChain()
		if exit {
			close(ccl.exit)
			break
		}
		time.Sleep(time.Second * 5)
	}
}

func (ccl *CrossChainListen) HandleNewBlock(height uint64) (w []*models.WrapperTransaction, s []*models.SrcTransaction, p []*models.PolyTransaction, d []*models.DstTransaction, err error) {
	// chain := ccl.handle.GetChainId()
	// var locks, unlocks int
	w, s, p, d, _, _, err = ccl.handle.HandleNewBlock(height)
	if err != nil {
		return
	}
	// logs.Error("Possible inconsistent chain %d height %d wrapper %d/%d src %d/%d dst %d/%d", chain, height, len(w), locks, len(s), locks, len(d), unlocks)
	return
}
func (ccl *CrossChainListen) listenChain() (exit bool) {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("%s listenChain service start, recover info: %s", ccl.handle.GetChainName(), string(debug.Stack()))
			exit = false
		}
	}()
	chain, err := ccl.db.GetChain(ccl.handle.GetChainId())
	if err != nil {
		panic(err)
	}
	height, err := ccl.handle.GetLatestHeight()
	if err != nil || height == 0 {
		panic(err)
	}
	if chain.Height == 0 {
		chain.Height = height
	}
	ccl.db.UpdateChain(chain)
	if ccl.height != 0 {
		chain.Height = ccl.height
	}
	if ccl.config.Backup {
		chain.Height -= ccl.handle.GetDefer()
	}
	logs.Info("cross chain listen, chain: %s, dao: %s......", ccl.handle.GetChainName(), ccl.db.Name())
	ticker := time.NewTicker(time.Second * time.Duration(ccl.handle.GetChainListenSlot()))
	for {
		select {
		case <-ticker.C:
			if ccl.config.Backup {
				dbchain, err := ccl.db.GetChain(chain.ChainId)
				if err != nil {
					continue
				}
				height = dbchain.Height
				if chain.Height >= height-ccl.handle.GetDefer() {
					continue
				}
				logs.Info("backup ListenChain - chain %s db height is %d, listen height: %d", ccl.handle.GetChainName(), height, chain.Height)
			} else {
				height, err = ccl.handle.GetLatestHeight()
				if err != nil || height == 0 || height == math.MaxUint64 {
					logs.Error("listenChain - cannot get chain %s height, err: %s", ccl.handle.GetChainName(), err)
					continue
				}
				extendHeight, err := ccl.handle.GetExtendLatestHeight()
				if err != nil || extendHeight == 0 {
					logs.Error("ListenChain - cannot get chain %s extend height, err: %s", ccl.handle.GetChainName(), err)
				} else if extendHeight >= height+21 {
					logs.Error("ListenChain - chain %s node is too slow, node height: %d, really height: %d", ccl.handle.GetChainName(), height, extendHeight)
				}
				metrics.Record(height, "%v.lastest_height", chain.ChainId)
				metrics.Record(extendHeight, "%v.watch_height", chain.ChainId)
				metrics.Record(chain.Height, "%v.height", chain.ChainId)
				if chain.Height >= height-ccl.handle.GetDefer() {
					continue
				}
				logs.Info("ListenChain - chain %s latest height is %d, listen height: %d", ccl.handle.GetChainName(), height, chain.Height)
			}
			for chain.Height < height-ccl.handle.GetDefer() {
				batchSize := ccl.handle.GetBatchSize()
				if batchSize == 0 {
					batchSize = 1
				}
				if batchSize > height-chain.Height-ccl.handle.GetDefer() {
					batchSize = height - chain.Height - ccl.handle.GetDefer()
				}

				ch := make(chan bool, batchSize)
				for i := uint64(1); i <= batchSize; i++ {
					go func(height uint64) {
						wrapperTransactions, srcTransactions, polyTransactions, dstTransactions, err := ccl.HandleNewBlock(height)
						if err != nil {
							logs.Error("HandleNewBlock chain：%s, height: %d err: %v", ccl.handle.GetChainName(), height, err)
							ch <- false
							return
						}
						logs.Info("HandleNewBlock [chainName: %s, height: %d]. "+
							"len(wrapperTransactions)=%d, len(srcTransactions)=%d, len(polyTransactions)=%d, len(dstTransactions)=%d",
							chain.Name, height, len(wrapperTransactions), len(srcTransactions), len(polyTransactions), len(dstTransactions))
						err = ccl.db.UpdateEvents(wrapperTransactions, srcTransactions, polyTransactions, dstTransactions)
						if err != nil {
							logs.Error("UpdateEvents on block %d err: %v", height, err)
							ch <- false
						} else {
							if !ccl.config.Backup {
								go ccl.checkLargeTransaction(srcTransactions)
							}
							ch <- true
						}

					}(chain.Height + i)
				}
				allTaskSuccess := true
				for j := 0; j < int(batchSize); j++ {
					ok := <-ch
					if !ok {
						allTaskSuccess = false
					}
				}
				close(ch)
				if !allTaskSuccess {
					break
				}

				chain.Height += batchSize
				if err := ccl.db.UpdateChain(chain); err != nil {
					logs.Error("UpdateChain [chainId:%d, height:%d] err %v", chain.ChainId, chain.Height, err)
					chain.Height -= batchSize
				}
			}
		case <-ccl.exit:
			logs.Info("cross chain listen exit, chain: %s, dao: %s......", ccl.handle.GetChainName(), ccl.db.Name())
			return true
		}
	}
}

func (ccl *CrossChainListen) checkLargeTransaction(srcTransactions []*models.SrcTransaction) {
	ccl.dingMux.Lock()
	defer ccl.dingMux.Unlock()
	if srcTransactions != nil && len(srcTransactions) > 0 {
		for _, v := range srcTransactions {
			if existed, err := cacheRedis.Redis.Exists(cacheRedis.LargeTxAlarmPrefix + strings.ToLower(v.Hash)); err == nil && existed {
				logs.Info("large TX hash: %s alarm has been sent.", v.Hash)
				return
			}

			if ccl.isO3SwapTx(v) {
				logs.Info("hash: %s is O3Swap, skip large TX check.", v.Hash)
				return
			}

			if v.SrcTransfer != nil {
				token, err := ccl.db.GetTokenBasicByHash(v.SrcTransfer.ChainId, v.SrcTransfer.Asset)
				if err == nil {
					amount := decimal.NewFromBigInt(&v.SrcTransfer.Amount.Int, 0).
						Div(decimal.NewFromInt(basedef.Int64FromFigure(int(token.Precision)))).
						Mul(decimal.NewFromInt(token.TokenBasic.Price)).
						Div(decimal.NewFromInt(100000000))

					if amount.Cmp(decimal.NewFromInt(ccl.config.LargeTxAmount)) >= 0 {
						//cacheRedis.Redis.Unlink(cacheRedis.LargeTxList)
						if err := cacheRedis.Redis.RPush(cacheRedis.LargeTxList, v.Hash); err != nil {
							logs.Error("Save LargeTx[hash: %s] err: %s", v.Hash, err)
						}
						if err := ccl.sendLargeTransactionDingAlarm(v, token, ccl.config.LargeTxAmount, amount); err != nil {
							logs.Error("send LargeTxAmount alarm err:", err)
						} else {
							if _, err := cacheRedis.Redis.Set(cacheRedis.LargeTxAlarmPrefix+strings.ToLower(v.Hash), "done", time.Hour); err != nil {
								logs.Error("mark large TX hash: %s alarm done err: %s", v.Hash, err)
							}
						}
					}
				}
			}
		}
	}
}

func (ccl *CrossChainListen) isO3SwapTx(src *models.SrcTransaction) bool {
	if src.ChainId != basedef.O3_CROSSCHAIN_ID {
		return false
	}
	if dst, err := ccl.db.GetDstTransactionByHash(src.Hash); err == nil && dst != nil {
		return true
	}
	return false
}

func (ccl *CrossChainListen) sendLargeTransactionDingAlarm(srcTransaction *models.SrcTransaction, token *models.Token, largeTxAmount int64, amount decimal.Decimal) error {
	exceedingAmount := strconv.FormatInt(largeTxAmount, 10)
	if amount.Cmp(decimal.NewFromInt(10000000)) >= 0 {
		exceedingAmount = "1000w"
	} else if amount.Cmp(decimal.NewFromInt(5000000)) >= 0 {
		exceedingAmount = "500w"
	} else if amount.Cmp(decimal.NewFromInt(1000000)) >= 0 {
		exceedingAmount = "100w"
	}
	//ss := "A large transaction exceeding " + exceedingAmount + " USD was detected.\n"
	srcChainName := strconv.FormatUint(srcTransaction.ChainId, 10)
	srcChain, err := ccl.db.GetChain(srcTransaction.ChainId)
	if err == nil {
		srcChainName = srcChain.Name
	}

	dstChainName := strconv.FormatUint(srcTransaction.DstChainId, 10)
	dstChain, err := ccl.db.GetChain(srcTransaction.DstChainId)
	if err == nil {
		dstChainName = dstChain.Name
	}
	if srcTransaction.SrcSwap != nil && srcTransaction.SrcSwap.DstChainId != 0 {
		dstChainName = strconv.FormatUint(srcTransaction.SrcSwap.DstChainId, 10)
		dstChain, err := ccl.db.GetChain(srcTransaction.SrcSwap.DstChainId)
		if err == nil {
			dstChainName = dstChain.Name
		}
	}

	title := fmt.Sprintf("Large transaction exceeding %s USD (%s->%s)", exceedingAmount, srcChainName, dstChainName)
	//ss += "Asset " + token.Name + "(" + srcChainName + "->" + dstChainName + ")\n"
	txType := "SWAP"
	if srcTransaction.SrcSwap != nil {
		switch srcTransaction.SrcSwap.Type {
		case basedef.SWAP_SWAP:
			txType = "SWAP"
		case basedef.SWAP_ROLLBACK:
			txType = "ROLLBACK"
		case basedef.SWAP_ADDLIQUIDITY:
			txType = "ADDLIQUIDITY"
		case basedef.SWAP_REMOVELIQUIDITY:
			txType = "REMOVELIQUIDITY"
		}
	}
	largeTx := basedef.LargeTx{
		Asset:     token.Name,
		From:      srcChainName,
		To:        dstChainName,
		Type:      txType,
		Amount:    decimal.NewFromBigInt(&srcTransaction.SrcTransfer.Amount.Int, 0).Div(decimal.NewFromInt(basedef.Int64FromFigure(int(token.Precision)))).String(),
		USDAmount: amount.String(),
		Hash:      srcTransaction.Hash,
		User:      srcTransaction.User,
		Time:      time.Unix(int64(srcTransaction.Time), 0).Format("2006-01-02 15:04:05"),
	}
	body := fmt.Sprintf("## %s\n- Asset: %s\n- Type: %s\n- Amount: %s %s (%s USD)\n- Hash: %s\n- User: %s\n- Time: %s\n",
		title,
		largeTx.Asset,
		largeTx.Type,
		largeTx.Amount, largeTx.Asset, largeTx.USDAmount,
		largeTx.Hash,
		largeTx.User,
		largeTx.Time,
	)
	logs.Info(body)
	//cacheRedis.Redis.Unlink(cacheRedis.LargeTxList)
	//if largeTxJson, err := json.Marshal(largeTx); err == nil {
	//	value := string(largeTxJson)
	//	if err := cacheRedis.Redis.RPush(cacheRedis.LargeTxList, value); err != nil {
	//		logs.Error("Save LargeTx[hash: %s] err: %s", srcTransaction.Hash, err)
	//	}
	//}

	btns := []map[string]string{
		{
			"title":     "List All",
			"actionURL": fmt.Sprintf("%stoken=%s", conf.GlobalConfig.BotConfig.BaseUrl+conf.GlobalConfig.BotConfig.ListLargeTxUrl, conf.GlobalConfig.BotConfig.ApiToken),
		},
	}
	return common.PostDingCard(title, body, btns, conf.GlobalConfig.BotConfig.LargeTxDingUrl)
}
