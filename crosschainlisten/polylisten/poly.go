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
	"github.com/beego/beego/v2/core/logs"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"poly-bridge/models"
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

func (this *PolyChainListen) GetChainListenSlot() uint64 {
	return this.polyCfg.ListenSlot
}

func (this *PolyChainListen) GetChainId() uint64 {
	return this.polyCfg.ChainId
}

func (this *PolyChainListen) GetChainName() string {
	return this.polyCfg.ChainName
}

func (this *PolyChainListen) GetDefer() uint64 {
	return this.polyCfg.Defer
}

func (this *PolyChainListen) GetBatchSize() uint64 {
	return this.polyCfg.BatchSize
}

func (this *PolyChainListen) HandleNewBlock(height uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, []*models.WrapperDetail, []*models.PolyDetail, int, int, error) {
	block, err := this.polySdk.GetBlockByHeight(height)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, 0, 0, err
	}
	if block == nil {
		return nil, nil, nil, nil, nil, nil, 0, 0, fmt.Errorf("there is no poly block!")
	}
	tt := block.Header.Timestamp
	events, err := this.polySdk.GetSmartContractEventByBlock(height)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, 0, 0, err
	}
	polyTransactions := make([]*models.PolyTransaction, 0)
	polyDetails := make([]*models.PolyDetail, 0)
	for _, event := range events {
		for _, notify := range event.Notify {
			if notify.ContractAddress == this.polyCfg.CCMContract {
				states := notify.States.([]interface{})
				contractMethod, _ := states[0].(string)
				logs.Info("chain: %s, height: %d, tx hash: %s", this.GetChainName(), height, event.TxHash)
				if contractMethod != "makeProof" && contractMethod != "btcTxToRelay" && contractMethod != "multisignedTxJson" && contractMethod != "rippleTxJson" {
					continue
				}
				if len(states) < 4 {
					continue
				}
				fchainid := uint32(states[1].(float64))
				tchainid := uint32(states[2].(float64))

				switch contractMethod {
				case "rippleTxJson":
					polyDetail := &models.PolyDetail{}
					polyDetail.ChainId = this.GetChainId()
					polyDetail.Hash = event.TxHash
					polyDetail.State = uint64(event.State)
					polyDetail.Fee = &models.BigInt{*big.NewInt(0)}
					polyDetail.Time = uint64(tt)
					polyDetail.Height = height
					polyDetail.SrcChainId = uint64(fchainid)
					polyDetail.DstChainId = uint64(tchainid)
					switch uint64(fchainid) {
					case basedef.NEO_CROSSCHAIN_ID, basedef.NEO3_CROSSCHAIN_ID, basedef.ONT_CROSSCHAIN_ID:
						polyDetail.SrcHash = basedef.HexStringReverse(states[3].(string))
					default:
						polyDetail.SrcHash = states[3].(string)
					}
					switch uint64(tchainid) {
					case basedef.RIPPLE_CROSSCHAIN_ID:
						sequence := states[5].(float64)
						polyDetail.DstSequence = uint64(sequence)
					}

					polyDetails = append(polyDetails, polyDetail)
				default:
					mctx := &models.PolyTransaction{}
					mctx.ChainId = this.GetChainId()
					mctx.Hash = event.TxHash
					mctx.State = uint64(event.State)
					mctx.Fee = &models.BigInt{*big.NewInt(0)}
					mctx.Time = uint64(tt)
					mctx.Height = height
					mctx.SrcChainId = uint64(fchainid)
					mctx.DstChainId = uint64(tchainid)
					switch uint64(fchainid) {
					case basedef.NEO_CROSSCHAIN_ID, basedef.NEO3_CROSSCHAIN_ID, basedef.ONT_CROSSCHAIN_ID:
						mctx.SrcHash = basedef.HexStringReverse(states[3].(string))
					default:
						mctx.SrcHash = states[3].(string)
					}
					switch uint64(tchainid) {
					case basedef.RIPPLE_CROSSCHAIN_ID:
						sequence := states[5].(float64)
						mctx.DstSequence = uint64(sequence)
					}

					polyTransactions = append(polyTransactions, mctx)
				}
			}
		}
	}
	return nil, nil, polyTransactions, nil, nil, polyDetails, 0, 0, nil
}

func (this *PolyChainListen) GetExtendLatestHeight() (uint64, error) {
	if len(this.polyCfg.ExtendNodes) == 0 {
		return this.GetLatestHeight()
	}
	return this.GetLatestHeight()
}
