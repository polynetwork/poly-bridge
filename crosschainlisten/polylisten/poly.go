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

package polylisten

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"math/big"
	"poly-swap/chainsdk"
	"poly-swap/conf"
	"poly-swap/models"
	"poly-swap/utils"
)

type PolyChainListen struct {
	polyCfg *conf.ChainListenConfig
	polySdk *chainsdk.PolySDKPro
}

func NewPolyChainListen(cfg *conf.ChainListenConfig) *PolyChainListen {
	polyListen := &PolyChainListen{}
	polyListen.polyCfg = cfg
	urls := cfg.GetNodesUrl()
	sdk := chainsdk.NewPolySDKPro(urls, cfg.ListenSlot, cfg.ChainId)
	polyListen.polySdk = sdk
	return polyListen
}

func (this *PolyChainListen) GetLatestHeight() (uint64, error) {
	return this.polySdk.GetCurrentBlockHeight()
}

func (this *PolyChainListen) GetBackwardBlockNumber() uint64 {
	return this.polyCfg.BackwardBlockNumber
}

func (this *PolyChainListen) GetChainListenSlot() uint64 {
	return this.polyCfg.ListenSlot
}

func (this *PolyChainListen) GetChainId() uint64 {
	return this.polyCfg.ChainId
}

func (this *PolyChainListen) GetChainName() string {
	return this.polyCfg.ChainName
}

func (this *PolyChainListen) HandleNewBlock(height uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, error) {
	block, err := this.polySdk.GetBlockByHeight(height)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	if block == nil {
		return nil, nil, nil, nil, fmt.Errorf("there is no poly block!")
	}
	tt := block.Header.Timestamp
	events, err := this.polySdk.GetSmartContractEventByBlock(height)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	polyTransactions := make([]*models.PolyTransaction, 0)
	for _, event := range events {
		for _, notify := range event.Notify {
			if notify.ContractAddress == this.polyCfg.CCMContract {
				states := notify.States.([]interface{})
				contractMethod, _ := states[0].(string)
				logs.Info("chain: %s, tx hash: %s, method: %s, state: %d, gas: %d\n", this.GetChainName(), event.TxHash, contractMethod, event.State, 0)
				if contractMethod != "makeProof" && contractMethod != "btcTxToRelay" {
					continue
				}
				if len(states) < 4 {
					continue
				}
				fchainid := uint32(states[1].(float64))
				tchainid := uint32(states[2].(float64))
				mctx := &models.PolyTransaction{}
				mctx.ChainId = this.GetChainId()
				mctx.Hash = event.TxHash
				mctx.State = uint64(event.State)
				mctx.Fee = &models.BigInt{*big.NewInt(0)}
				mctx.Time = uint64(tt)
				mctx.Height = height
				mctx.SrcChainId = uint64(fchainid)
				mctx.DstChainId = uint64(tchainid)
				if uint64(fchainid) == conf.ETHEREUM_CROSSCHAIN_ID {
					mctx.SrcHash = states[3].(string)
				} else {
					mctx.SrcHash = utils.HexStringReverse(states[3].(string))
				}
				polyTransactions = append(polyTransactions, mctx)
			}
		}
	}
	return nil, nil, polyTransactions, nil, nil
}

func (this *PolyChainListen) GetExtendLatestHeight() (uint64, error) {
	if len(this.polyCfg.ExtendNodes) == 0 {
		return this.GetLatestHeight()
	}
	return this.GetLatestHeight()
}
