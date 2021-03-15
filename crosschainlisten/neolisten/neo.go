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

package neolisten

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"math/big"
	"net/http"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"poly-bridge/models"
	"strings"
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

type NeoChainListen struct {
	neoCfg *conf.ChainListenConfig
	neoSdk *chainsdk.NeoSdkPro
}

func NewNeoChainListen(cfg *conf.ChainListenConfig) *NeoChainListen {
	ethListen := &NeoChainListen{}
	ethListen.neoCfg = cfg
	urls := cfg.GetNodesUrl()
	sdk := chainsdk.NewNeoSdkPro(urls, cfg.ListenSlot, cfg.ChainId)
	ethListen.neoSdk = sdk
	return ethListen
}

func (this *NeoChainListen) GetLatestHeight() (uint64, error) {
	return this.neoSdk.GetBlockCount()
}

func (this *NeoChainListen) GetChainListenSlot() uint64 {
	return this.neoCfg.ListenSlot
}

func (this *NeoChainListen) GetChainId() uint64 {
	return this.neoCfg.ChainId
}

func (this *NeoChainListen) GetChainName() string {
	return this.neoCfg.ChainName
}

func (this *NeoChainListen) parseNeoMethod(v string) string {
	xx, _ := hex.DecodeString(v)
	return string(xx)
}

func (this *NeoChainListen) GetDefer() uint64 {
	return this.neoCfg.Defer
}

func (this *NeoChainListen) HandleNewBlock(height uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, error) {
	block, err := this.neoSdk.GetBlockByIndex(height)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	if block == nil {
		return nil, nil, nil, nil, fmt.Errorf("can not get neo block!")
	}
	tt := block.Time
	wrapperTransactions := make([]*models.WrapperTransaction, 0)
	srcTransactions := make([]*models.SrcTransaction, 0)
	dstTransactions := make([]*models.DstTransaction, 0)
	for _, tx := range block.Tx {
		if tx.Type != "InvocationTransaction" {
			continue
		}
		appLog, err := this.neoSdk.GetApplicationLog(tx.Txid)
		if err != nil || appLog == nil {
			continue
		}
		for _, exeitem := range appLog.Executions {
			for _, notify := range exeitem.Notifications {
				if notify.Contract[2:] == this.neoCfg.WrapperContract {
					if len(notify.State.Value) < 0 {
						continue
					}
					contractMethod := this.parseNeoMethod(notify.State.Value[0].Value)
					switch contractMethod {
					case _poly_wrapper_lock:
						logs.Info("(wrapper) from chain: %s, txhash: %s", this.GetChainName(), tx.Txid[2:])
						if len(notify.State.Value) < 8 {
							continue
						}
						value := notify.State.Value
						tchainId := big.NewInt(0)
						if value[3].Type == "Integer" {
							tchainId, _ = new(big.Int).SetString(value[3].Value, 10)
						} else {
							tchainId, _ = new(big.Int).SetString(basedef.HexStringReverse(value[3].Value), 16)
						}
						serverId := big.NewInt(0)
						if value[7].Type == "Integer" {
							serverId, _ = new(big.Int).SetString(value[7].Value, 10)
						} else {
							serverId, _ = new(big.Int).SetString(basedef.HexStringReverse(value[7].Value), 16)
						}
						if serverId == nil {
							serverId = new(big.Int).SetUint64(0)
						}
						asset := basedef.HexStringReverse(value[1].Value)
						amount := big.NewInt(0)
						if value[6].Type == "Integer" {
							amount, _ = new(big.Int).SetString(value[6].Value, 10)
						} else {
							amount, _ = new(big.Int).SetString(basedef.HexStringReverse(value[6].Value), 16)
						}
						wrapperTransactions = append(wrapperTransactions, &models.WrapperTransaction{
							Hash:         tx.Txid[2:],
							User:         notify.State.Value[2].Value,
							DstChainId:   tchainId.Uint64(),
							DstUser:      notify.State.Value[4].Value,
							FeeTokenHash: asset,
							FeeAmount:    models.NewBigInt(amount),
							ServerId:     serverId.Uint64(),
							Status:       basedef.STATE_SOURCE_DONE,
							Time:         uint64(tt),
							BlockHeight:  height,
							SrcChainId:   this.GetChainId(),
						})
					}
				} else if notify.Contract[2:] == this.neoCfg.CCMContract {
					if len(notify.State.Value) <= 0 {
						continue
					}
					contractMethod := this.parseNeoMethod(notify.State.Value[0].Value)
					switch contractMethod {
					case _neo_crosschainlock:
						logs.Info("(lock) from chain: %s, txhash: %s", this.GetChainName(), tx.Txid[2:])
						if len(notify.State.Value) < 6 {
							continue
						}
						fctransfer := &models.SrcTransfer{}
						for _, notifynew := range exeitem.Notifications {
							contractMethodNew := this.parseNeoMethod(notifynew.State.Value[0].Value)
							if contractMethodNew == _neo_lock || contractMethodNew == _neo_lock2 {
								if len(notifynew.State.Value) < 7 {
									continue
								}
								fctransfer.ChainId = this.GetChainId()
								fctransfer.TxHash = tx.Txid[2:]
								fctransfer.Time = uint64(tt)
								fctransfer.From = notifynew.State.Value[2].Value
								fctransfer.To = notify.State.Value[2].Value
								fctransfer.Asset = basedef.HexStringReverse(notifynew.State.Value[1].Value)
								amount := big.NewInt(0)
								if notifynew.State.Value[6].Type == "Integer" {
									amount, _ = new(big.Int).SetString(notifynew.State.Value[6].Value, 10)
								} else {
									amount, _ = new(big.Int).SetString(basedef.HexStringReverse(notifynew.State.Value[6].Value), 16)
								}
								fctransfer.Amount = models.NewBigInt(amount)
								tChainId := big.NewInt(0)
								if notifynew.State.Value[3].Type == "Integer" {
									tChainId, _ = new(big.Int).SetString(notifynew.State.Value[3].Value, 10)
								} else {
									tChainId, _ = new(big.Int).SetString(basedef.HexStringReverse(notifynew.State.Value[3].Value), 16)
								}
								fctransfer.DstChainId = tChainId.Uint64()
								if len(notifynew.State.Value[5].Value) != 40 {
									continue
								}
								fctransfer.DstUser = notifynew.State.Value[5].Value
								fctransfer.DstAsset = notifynew.State.Value[4].Value
								break
							}
						}
						fctx := &models.SrcTransaction{}
						fctx.ChainId = this.GetChainId()
						fctx.Hash = tx.Txid[2:]
						fctx.State = 1
						fctx.Fee = models.NewBigInt(big.NewInt(int64(basedef.String2Float64(exeitem.GasConsumed))))
						fctx.Time = uint64(tt)
						fctx.Height = height
						fctx.User = fctransfer.From
						toChainId := big.NewInt(0)
						if notify.State.Value[3].Type == "Integer" {
							toChainId, _ = new(big.Int).SetString(notify.State.Value[3].Value, 10)
						} else {
							toChainId, _ = new(big.Int).SetString(basedef.HexStringReverse(notify.State.Value[3].Value), 16)
						}
						fctx.DstChainId = toChainId.Uint64()
						fctx.Contract = notify.State.Value[2].Value
						fctx.Key = notify.State.Value[4].Value
						fctx.Param = notify.State.Value[5].Value
						fctx.SrcTransfer = fctransfer
						srcTransactions = append(srcTransactions, fctx)
					case _neo_crosschainunlock:
						logs.Info("(unlock) to chain: %s, txhash: %s", this.GetChainName(), tx.Txid[2:])
						if len(notify.State.Value) < 4 {
							continue
						}
						tctransfer := &models.DstTransfer{}
						for _, notifynew := range exeitem.Notifications {
							contractMethodNew := this.parseNeoMethod(notifynew.State.Value[0].Value)
							if contractMethodNew == _neo_unlock || contractMethodNew == _neo_unlock2 {
								if len(notifynew.State.Value) < 4 {
									continue
								}
								tctransfer.ChainId = this.GetChainId()
								tctransfer.TxHash = tx.Txid[2:]
								tctransfer.Time = uint64(tt)
								tctransfer.From = notify.State.Value[2].Value
								tctransfer.To = notifynew.State.Value[2].Value
								tctransfer.Asset = basedef.HexStringReverse(notifynew.State.Value[1].Value)
								amount := big.NewInt(0)
								if notifynew.State.Value[3].Type == "Integer" {
									amount, _ = new(big.Int).SetString(notifynew.State.Value[3].Value, 10)
								} else {
									amount, _ = new(big.Int).SetString(basedef.HexStringReverse(notifynew.State.Value[3].Value), 16)
								}
								tctransfer.Amount = models.NewBigInt(amount)
								break
							}
						}
						tctx := &models.DstTransaction{}
						tctx.ChainId = this.GetChainId()
						tctx.Hash = tx.Txid[2:]
						tctx.State = 1
						tctx.Fee = models.NewBigInt(big.NewInt(int64(basedef.String2Float64(exeitem.GasConsumed))))
						tctx.Time = uint64(tt)
						tctx.Height = height
						fChainId := big.NewInt(0)
						if notify.State.Value[1].Type == "Integer" {
							fChainId, _ = new(big.Int).SetString(notify.State.Value[1].Value, 10)
						} else {
							fChainId, _ = new(big.Int).SetString(basedef.HexStringReverse(notify.State.Value[1].Value), 16)
						}
						tctx.SrcChainId = fChainId.Uint64()
						tctx.Contract = basedef.HexStringReverse(notify.State.Value[2].Value)
						tctx.PolyHash = basedef.HexStringReverse(notify.State.Value[3].Value)
						tctx.DstTransfer = tctransfer
						dstTransactions = append(dstTransactions, tctx)
					default:
						logs.Warn("ignore method: %s", contractMethod)
					}
				}
			}
		}
	}
	return wrapperTransactions, srcTransactions, nil, dstTransactions, nil
}

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}
type ExtendHeight struct {
	LastHeight uint64 `json:"result"`
	Error      *Error `json:"error"`
}

func (this *NeoChainListen) GetExtendLatestHeight() (uint64, error) {
	if len(this.neoCfg.ExtendNodes) == 0 {
		return this.GetLatestHeight()
	}
	for i, _ := range this.neoCfg.ExtendNodes {
		height, err := this.getExtendLatestHeight(i)
		if err == nil {
			return height, nil
		}
	}
	return 0, fmt.Errorf("all extend node is not working")
}

func (this *NeoChainListen) getExtendLatestHeight(node int) (uint64, error) {
	requestJson := `{"jsonrpc": "2.0", "method": "getblockcount", "params": [], "id": 1}`
	req, err := http.NewRequest("POST", this.neoCfg.ExtendNodes[node].Url, strings.NewReader(requestJson))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Accepts", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("response status code: %d", resp.StatusCode)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	extendHeight := new(ExtendHeight)
	err = json.Unmarshal(respBody, extendHeight)
	if err != nil {
		return 0, err
	}
	if extendHeight.Error != nil {
		return 0, fmt.Errorf("%s", extendHeight.Error.Message)
	}
	return extendHeight.LastHeight, nil
}
