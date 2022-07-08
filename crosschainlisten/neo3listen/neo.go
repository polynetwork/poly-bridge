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

package neo3listen

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"

	neo3_models "github.com/joeqian10/neo3-gogogo/rpc/models"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"poly-bridge/models"

	"github.com/beego/beego/v2/core/logs"
)

const (
	_neo_crosschainlock   = "CrossChainLockEvent"
	_neo_crosschainunlock = "CrossChainUnlockEvent"
	_neo_lock             = "Lock"
	_neo_lock2            = "LockEvent"
	_neo_unlock           = "UnlockEvent"
	_neo_unlock2          = "Unlock"
	_poly_wrapper_lock    = "PolyWrapperLock"
)

type Neo3ChainListen struct {
	neoCfg *conf.ChainListenConfig
	neoSdk *chainsdk.Neo3SdkPro
}

func NewNeo3ChainListen(cfg *conf.ChainListenConfig) *Neo3ChainListen {
	ethListen := &Neo3ChainListen{}
	ethListen.neoCfg = cfg
	urls := cfg.GetNodesUrl()
	sdk := chainsdk.NewNeo3SdkPro(urls, cfg.ListenSlot, cfg.ChainId)
	ethListen.neoSdk = sdk
	return ethListen
}

func (this *Neo3ChainListen) GetLatestHeight() (uint64, error) {
	return this.neoSdk.GetBlockCount()
}

func (this *Neo3ChainListen) GetChainListenSlot() uint64 {
	return this.neoCfg.ListenSlot
}

func (this *Neo3ChainListen) GetChainId() uint64 {
	return this.neoCfg.ChainId
}

func (this *Neo3ChainListen) GetChainName() string {
	return this.neoCfg.ChainName
}

func (this *Neo3ChainListen) parseNeoMethod(v string) string {
	xx, _ := hex.DecodeString(v)
	return string(xx)
}

func (this *Neo3ChainListen) GetDefer() uint64 {
	return this.neoCfg.Defer
}

func (this *Neo3ChainListen) GetBatchSize() uint64 {
	return this.neoCfg.BatchSize
}

func (this *Neo3ChainListen) isListeningContract(contract string, contracts []string) bool {
	for _, item := range contracts {
		if contract == item {
			return true
		}
	}
	return false
}

func (this *Neo3ChainListen) HandleNewBlock(height uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, []*models.WrapperDetail, []*models.PolyDetail, int, int, error) {
	block, err := this.neoSdk.GetBlockByIndex(height)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, 0, 0, err
	}
	if block == nil {
		return nil, nil, nil, nil, nil, nil, 0, 0, fmt.Errorf("can not get neo block!")
	}
	tt := block.Time / 1000
	wrapperTransactions := make([]*models.WrapperTransaction, 0)
	srcTransactions := make([]*models.SrcTransaction, 0)
	dstTransactions := make([]*models.DstTransaction, 0)
	for _, tx := range block.Tx {
		appLog, err := this.neoSdk.GetApplicationLog(tx.Hash)
		if err != nil || appLog == nil {
			continue
		}
		for _, exeitem := range appLog.Executions {
			if exeitem.VMState == "FAULT" {
				continue
			}
			for _, notify := range exeitem.Notifications {
				if this.isListeningContract(notify.Contract[2:], this.neoCfg.WrapperContract) {
					if notify.State.Type != "Array" {
						continue
					}
					notify.State.Convert()
					states := notify.State.Value.([]neo3_models.InvokeStack)
					if len(states) < 0 {
						continue
					}
					eventName := notify.EventName
					switch eventName {
					case _poly_wrapper_lock:
						logs.Info("(wrapper) from chain: %s, txhash: %s", this.GetChainName(), tx.Hash[2:])
						if len(states) < 7 {
							continue
						}
						tchainId := big.NewInt(0)
						if states[2].Type == "Integer" {
							tchainId, _ = new(big.Int).SetString(states[2].Value.(string), 10)
						} else {
							tchainId, _ = new(big.Int).SetString(basedef.HexStringReverse(states[2].Value.(string)), 16)
						}
						serverId := big.NewInt(0)
						if states[6].Type == "Integer" {
							serverId, _ = new(big.Int).SetString(states[6].Value.(string), 10)
						} else {
							serverId, _ = new(big.Int).SetString(basedef.HexStringReverse(states[6].Value.(string)), 16)
						}
						if serverId == nil {
							serverId = new(big.Int).SetUint64(0)
						}

						encodeUserString := states[1].Value.(string)
						decodeUserBytes, err := base64.StdEncoding.DecodeString(encodeUserString)
						if err != nil {
							logs.Error("txhash: %s decode wrapper user: %s err: %s", tx.Hash[2:], encodeUserString, err)
							continue
						}
						user := hex.EncodeToString(basedef.HexReverse(decodeUserBytes))

						encodeDstUserString := states[3].Value.(string)
						decodeDstUserBytes, err := base64.StdEncoding.DecodeString(encodeDstUserString)
						if err != nil {
							logs.Error("txhash: %s decode wrapper dst user: %s err: %s", tx.Hash[2:], encodeDstUserString, err)
							continue
						}
						dstUser := hex.EncodeToString(basedef.HexReverse(decodeDstUserBytes))

						amount := big.NewInt(0)
						if states[5].Type == "Integer" {
							amount, _ = new(big.Int).SetString(states[5].Value.(string), 10)
						} else {
							amount, _ = new(big.Int).SetString(basedef.HexStringReverse(states[5].Value.(string)), 16)
						}
						wrapperTransactions = append(wrapperTransactions, &models.WrapperTransaction{
							Hash:         tx.Hash[2:],
							User:         user,
							DstChainId:   tchainId.Uint64(),
							DstUser:      dstUser,
							FeeTokenHash: "d2a4cff31913016155e38e474a2c06d08be276cf",
							FeeAmount:    models.NewBigInt(amount),
							ServerId:     serverId.Uint64(),
							Status:       basedef.STATE_SOURCE_DONE,
							Time:         uint64(tt),
							BlockHeight:  height,
							SrcChainId:   this.GetChainId(),
						})
					}
				} else if notify.Contract[2:] == this.neoCfg.CCMContract {
					if notify.State.Type != "Array" {
						continue
					}
					notify.State.Convert()
					states := notify.State.Value.([]neo3_models.InvokeStack)
					eventName := notify.EventName
					switch eventName {
					case _neo_crosschainlock:
						logs.Info("(lock) from chain: %s, height:%d, txhash: %s", this.GetChainName(), height, tx.Hash[2:])
						if len(states) < 5 {
							continue
						}
						fctransfer := &models.SrcTransfer{}
						contract, _ := states[1].ToParameter()
						toChainId, _ := states[2].ToParameter()
						key, _ := states[3].ToParameter()
						param, _ := states[4].ToParameter()
						for _, notifyNew := range exeitem.Notifications {
							if notifyNew.State.Type != "Array" {
								continue
							}
							notifyNew.State.Convert()
							statesNew := notifyNew.State.Value.([]neo3_models.InvokeStack)
							eventNameNew := notifyNew.EventName
							if eventNameNew == _neo_lock || eventNameNew == _neo_lock2 {
								if len(statesNew) < 6 {
									continue
								}
								fromAddress, _ := statesNew[1].ToParameter()
								toAddress := contract
								asset, _ := statesNew[0].ToParameter()
								//TODO nft
								amount, err := statesNew[5].ToParameter()
								if err != nil {
									logs.Error("%v is neo3 nft", tx.Hash[2:])
									continue
								}
								toChainId, _ := statesNew[2].ToParameter()
								dstUser, _ := statesNew[4].ToParameter()
								dstAsset, _ := statesNew[3].ToParameter()
								fctransfer.ChainId = this.GetChainId()
								fctransfer.TxHash = tx.Hash[2:]
								fctransfer.Time = uint64(tt)
								fctransfer.From = hex.EncodeToString(fromAddress.Value.([]byte))
								fctransfer.To = hex.EncodeToString(toAddress.Value.([]byte))
								fctransfer.Asset = basedef.HexStringReverse(hex.EncodeToString(asset.Value.([]byte)))
								//TODO nft
								if _, ok := amount.Value.(*big.Int); !ok {
									continue
								}
								fctransfer.Amount = models.NewBigInt(amount.Value.(*big.Int))
								fctransfer.DstChainId = toChainId.Value.(*big.Int).Uint64()
								fctransfer.DstUser = hex.EncodeToString(dstUser.Value.([]byte))
								fctransfer.DstAsset = hex.EncodeToString(dstAsset.Value.([]byte))
								break
							}
						}
						fctx := &models.SrcTransaction{}
						fctx.ChainId = this.GetChainId()
						fctx.Hash = tx.Hash[2:]
						fctx.State = 1
						fctx.Fee = models.NewBigInt(big.NewInt(int64(basedef.String2Float64(exeitem.GasConsumed))))
						fctx.Time = uint64(tt)
						fctx.Height = height
						fctx.User = fctransfer.From
						fctx.DstChainId = toChainId.Value.(*big.Int).Uint64()
						fctx.Contract = hex.EncodeToString(contract.Value.([]byte))
						fctx.Key = hex.EncodeToString(key.Value.([]byte))
						fctx.Param = hex.EncodeToString(param.Value.([]byte))
						fctx.SrcTransfer = fctransfer
						srcTransactions = append(srcTransactions, fctx)
					case _neo_crosschainunlock:
						logs.Info("(unlock) to chain: %s, height:%d, txhash: %s", this.GetChainName(), height, tx.Hash[2:])
						if len(states) < 3 {
							continue
						}
						aaa, _ := states[0].ToParameter()
						fromChainId := new(big.Int).SetBytes(basedef.HexReverse(aaa.Value.([]byte)))
						contract, _ := states[1].ToParameter()
						polyHash, _ := states[2].ToParameter()
						tctransfer := &models.DstTransfer{}
						for _, notifyNew := range exeitem.Notifications {
							if notifyNew.State.Type != "Array" {
								continue
							}
							notifyNew.State.Convert()
							statesNew := notifyNew.State.Value.([]neo3_models.InvokeStack)
							eventNameNew := notifyNew.EventName
							if eventNameNew == _neo_unlock || eventNameNew == _neo_unlock2 {
								if len(statesNew) < 3 {
									continue
								}
								fromAddress := contract
								toAddress, _ := statesNew[1].ToParameter()
								amount, _ := statesNew[2].ToParameter()
								asset, _ := statesNew[0].ToParameter()
								tctransfer.ChainId = this.GetChainId()
								tctransfer.TxHash = tx.Hash[2:]
								tctransfer.Time = uint64(tt)
								tctransfer.From = hex.EncodeToString(fromAddress.Value.([]byte))
								tctransfer.To = hex.EncodeToString(toAddress.Value.([]byte))
								tctransfer.Asset = basedef.HexStringReverse(hex.EncodeToString(asset.Value.([]byte)))
								if x, ok := amount.Value.(*big.Int); !ok {
									tctransfer.Amount = models.NewBigIntFromInt(0)
								} else {
									tctransfer.Amount = models.NewBigInt(x)
								}
								break
							}
						}
						tctx := &models.DstTransaction{}
						tctx.ChainId = this.GetChainId()
						tctx.Hash = tx.Hash[2:]
						tctx.State = 1
						tctx.Fee = models.NewBigInt(big.NewInt(int64(basedef.String2Float64(exeitem.GasConsumed))))
						tctx.Time = uint64(tt)
						tctx.Height = height
						tctx.SrcChainId = fromChainId.Uint64()
						tctx.Contract = basedef.HexStringReverse(hex.EncodeToString(contract.Value.([]byte)))
						tctx.PolyHash = basedef.HexStringReverse(hex.EncodeToString(polyHash.Value.([]byte)))
						tctx.DstTransfer = tctransfer
						dstTransactions = append(dstTransactions, tctx)
					default:
						logs.Warn("ignore method: %s", eventName)
					}
				}
			}
		}
	}
	return wrapperTransactions, srcTransactions, nil, dstTransactions, nil, nil, len(srcTransactions), len(dstTransactions), nil
}

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}
type ExtendHeight struct {
	LastHeight uint64 `json:"result"`
	Error      *Error `json:"error"`
}

func (this *Neo3ChainListen) GetExtendLatestHeight() (uint64, error) {
	if len(this.neoCfg.ExtendNodes) == 0 {
		return this.GetLatestHeight()
	}
	return this.GetLatestHeight()
}
