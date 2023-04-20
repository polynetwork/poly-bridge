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

package ontologylisten

import (
	"encoding/hex"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/ontio/ontology-go-sdk/utils"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"poly-bridge/models"
	"strconv"
)

const (
	_ont_crosschainlock   = "cross_chain"
	_ont_crosschainunlock = "verifyToOntProof"
	_ont_lock             = "lock"
	_ont_unlock           = "unlock"
	ont_wrapper_lock      = "PolyWrapperLock"
)

type OntologyChainListen struct {
	ontCfg *conf.ChainListenConfig
	ontSdk *chainsdk.OntologySdkPro
}

func NewOntologyChainListen(cfg *conf.ChainListenConfig) *OntologyChainListen {
	ontListen := &OntologyChainListen{}
	ontListen.ontCfg = cfg
	sdk := chainsdk.NewOntologySdkPro(cfg.Nodes, cfg.ListenSlot, cfg.ChainId)
	ontListen.ontSdk = sdk
	return ontListen
}

func (this *OntologyChainListen) GetLatestHeight() (uint64, error) {
	return this.ontSdk.GetCurrentBlockHeight()
}

func (this *OntologyChainListen) GetChainListenSlot() uint64 {
	return this.ontCfg.ListenSlot
}

func (this *OntologyChainListen) GetChainId() uint64 {
	return this.ontCfg.ChainId
}

func (this *OntologyChainListen) GetChainName() string {
	return this.ontCfg.ChainName
}

func (this *OntologyChainListen) parseOntolofyMethod(v string) string {
	xx, _ := hex.DecodeString(v)
	return string(xx)
}

func (this *OntologyChainListen) GetDefer() uint64 {
	return this.ontCfg.Defer
}

func (this *OntologyChainListen) GetBatchSize() uint64 {
	return this.ontCfg.BatchSize
}

func (this *OntologyChainListen) GetBatchLength() (uint64, uint64) {
	return this.ontCfg.MinBatchLength, this.ontCfg.MaxBatchLength
}

func (this *OntologyChainListen) isListeningContract(contract string, contracts []string) bool {
	for _, item := range contracts {
		if contract == item {
			return true
		}
	}
	return false
}

func (this *OntologyChainListen) HandleNewBatchBlock(start, end uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, []*models.WrapperDetail, []*models.PolyDetail, int, int, error) {
	return nil, nil, nil, nil, nil, nil, 0, 0, nil
}

func (this *OntologyChainListen) HandleNewBlock(height uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, []*models.WrapperDetail, []*models.PolyDetail, int, int, error) {
	block, err := this.ontSdk.GetBlockByHeight(uint32(height))
	if err != nil {
		err = fmt.Errorf("ontSdk.GetBlockByHeight of height:%d err:%s", height, err)
		return nil, nil, nil, nil, nil, nil, 0, 0, err
	}
	tt := uint64(block.Header.Timestamp)
	events, err := this.ontSdk.GetSmartContractEventByBlock(uint32(height))
	if err != nil {
		err = fmt.Errorf("ontSdk.GetSmartContractEventByBlock of height:%d err:%s", height, err)
		return nil, nil, nil, nil, nil, nil, 0, 0, err
	}
	wrapperTransactions := make([]*models.WrapperTransaction, 0)
	srcTransactions := make([]*models.SrcTransaction, 0)
	dstTransactions := make([]*models.DstTransaction, 0)
	for _, event := range events {
		for _, notify := range event.Notify {
			if this.isListeningContract(notify.ContractAddress, this.ontCfg.WrapperContract) {
				states := notify.States.([]interface{})
				contractMethod, ok := states[0].(string)
				if !ok {
					continue
				}
				contractMethod = this.parseOntolofyMethod(contractMethod)
				switch contractMethod {
				case ont_wrapper_lock:
					logs.Info("(wrapper) from chain: %s, height: %d, txhash: %s", this.GetChainName(), height, event.TxHash)
					if len(states) < 8 {
						continue
					}
					amount, _ := new(big.Int).SetString(basedef.HexStringReverse(states[6].(string)), 16)
					toChain, _ := new(big.Int).SetString(basedef.HexStringReverse(states[3].(string)), 16)
					serverId, _ := new(big.Int).SetString(basedef.HexStringReverse(states[7].(string)), 16)
					srcUser := states[2].(string)
					dstUser := states[4].(string)
					if len(srcUser) > basedef.ADDRESS_LENGTH || len(dstUser) > basedef.ADDRESS_LENGTH {
						continue
					}
					wrapperTransactions = append(wrapperTransactions, &models.WrapperTransaction{
						Hash:         event.TxHash,
						User:         states[2].(string),
						DstChainId:   toChain.Uint64(),
						DstUser:      states[4].(string),
						FeeTokenHash: basedef.HexStringReverse(states[1].(string)),
						FeeAmount:    models.NewBigInt(amount),
						ServerId:     serverId.Uint64(),
						Status:       basedef.STATE_SOURCE_DONE,
						Time:         tt,
						BlockHeight:  height,
						SrcChainId:   this.GetChainId(),
					})
				}
			} else if notify.ContractAddress == this.ontCfg.CCMContract {
				states := notify.States.([]interface{})
				contractMethod, _ := states[0].(string)
				switch contractMethod {
				case _ont_crosschainlock:
					logs.Info("(lock) from chain: %s, height: %d, txhash: %s", this.GetChainName(), height, event.TxHash)
					if len(states) < 6 {
						continue
					}
					srcTransfer := &models.SrcTransfer{}
					for _, notifyNew := range event.Notify {
						statesNew := notifyNew.States.([]interface{})
						method, ok := statesNew[0].(string)
						if !ok {
							continue
						}
						method = this.parseOntolofyMethod(method)
						if method == _ont_lock {
							if len(statesNew) < 7 {
								continue
							}
							srcTransfer.ChainId = this.GetChainId()
							srcTransfer.TxHash = event.TxHash
							srcTransfer.Time = tt
							srcTransfer.From = statesNew[2].(string)
							srcTransfer.To = states[5].(string)
							srcTransfer.Asset = basedef.HexStringReverse(statesNew[1].(string))
							if len(srcTransfer.Asset) < 20 {
								continue
							}
							amount, _ := new(big.Int).SetString(basedef.HexStringReverse(statesNew[6].(string)), 16)
							srcTransfer.Amount = models.NewBigInt(amount)
							toChain, _ := new(big.Int).SetString(basedef.HexStringReverse(statesNew[3].(string)), 16)
							srcTransfer.DstChainId = toChain.Uint64()
							srcTransfer.DstAsset = statesNew[4].(string)

							if srcTransfer.DstChainId == basedef.APTOS_CROSSCHAIN_ID {
								aptosAsset, err := hex.DecodeString(srcTransfer.DstAsset)
								if err == nil {
									srcTransfer.DstAsset = string(aptosAsset)
								}
							}
							srcTransfer.DstAsset = models.FormatAssert(srcTransfer.DstAsset)

							srcTransfer.DstUser = statesNew[5].(string)
							if len(srcTransfer.From) > basedef.ADDRESS_LENGTH {
								srcTransfer.From = ""
							}
							if len(srcTransfer.To) > basedef.ADDRESS_LENGTH {
								srcTransfer.To = ""
							}
							if len(srcTransfer.DstUser) > basedef.ADDRESS_LENGTH {
								srcTransfer.DstUser = ""
							}
							break
						}
					}
					srcTransaction := &models.SrcTransaction{}
					srcTransaction.ChainId = this.GetChainId()
					srcTransaction.Hash = event.TxHash
					srcTransaction.State = uint64(event.State)
					srcTransaction.Fee = models.NewBigIntFromInt(int64(event.GasConsumed))
					srcTransaction.Time = tt
					srcTransaction.Height = height
					srcTransaction.User = srcTransfer.From
					dstChainId, err := strconv.ParseUint(states[4].(string), 10, 32)
					if err != nil {
						logs.Error("invalid onto event log, fail to parse dstChainId err, %s, height %d", err, height)
					}
					srcTransaction.DstChainId = dstChainId
					contract, err := utils.AddressFromBase58(states[3].(string))
					if err != nil {
						logs.Error("invalid onto event log, fail to parse contractArr err, %s, height %d", err, height)
					}
					srcTransaction.Contract = contract.ToHexString()
					srcTransaction.Key = states[2].(string)
					srcTransaction.Param = states[5].(string)
					srcTransaction.SrcTransfer = srcTransfer
					srcTransactions = append(srcTransactions, srcTransaction)
				case _ont_crosschainunlock:
					logs.Info("(unlock) to chain: %s, height: %d, txhash: %s", this.GetChainName(), height, event.TxHash)
					if len(states) < 5 {
						continue
					}
					dstTransfer := &models.DstTransfer{}
					for _, notifyNew := range event.Notify {
						statesNew := notifyNew.States.([]interface{})
						method, ok := statesNew[0].(string)
						if !ok {
							continue
						}
						method = this.parseOntolofyMethod(method)
						if method == _ont_unlock {
							if len(statesNew) < 4 {
								continue
							}
							dstTransfer.ChainId = this.GetChainId()
							dstTransfer.TxHash = event.TxHash
							dstTransfer.Time = tt
							dstTransfer.From = states[4].(string)
							dstTransfer.To = statesNew[2].(string)
							dstTransfer.Asset = basedef.HexStringReverse(statesNew[1].(string))
							if len(dstTransfer.Asset) < 20 {
								continue
							}
							amount, _ := new(big.Int).SetString(basedef.HexStringReverse(statesNew[3].(string)), 16)
							dstTransfer.Amount = models.NewBigInt(amount)
							break
						}
					}
					dstTransaction := &models.DstTransaction{}
					dstTransaction.ChainId = this.GetChainId()
					dstTransaction.Hash = event.TxHash
					dstTransaction.State = uint64(event.State)
					dstTransaction.Fee = models.NewBigIntFromInt(int64(event.GasConsumed))
					dstTransaction.Time = tt
					dstTransaction.Height = height
					chainId, err := strconv.ParseInt(states[3].(string), 10, 64)
					if err != nil {
						logs.Error("invalid onto event log, fail to parse chainId err, %s, height %d", err, height)
					}
					dstTransaction.SrcChainId = uint64(chainId)
					dstTransaction.Contract = states[4].(string)
					dstTransaction.PolyHash = states[1].(string)
					dstTransaction.DstTransfer = dstTransfer
					dstTransactions = append(dstTransactions, dstTransaction)
				default:
					logs.Warn("ignore method: %s", contractMethod)
				}
			}
		}
	}
	return wrapperTransactions, srcTransactions, nil, dstTransactions, nil, nil, len(srcTransactions), len(dstTransactions), nil
}

func (this *OntologyChainListen) GetExtendLatestHeight() (uint64, error) {
	if len(this.ontCfg.ExtendNodes) == 0 {
		return this.GetLatestHeight()
	}
	return this.GetLatestHeight()
}
