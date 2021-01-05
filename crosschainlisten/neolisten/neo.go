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
	"poly-swap/chainsdk"
	"poly-swap/conf"
	"poly-swap/models"
	"poly-swap/utils"
	"strconv"
	"strings"
)

const (
	_neo_crosschainlock   = "CrossChainLockEvent"
	_neo_crosschainunlock = "CrossChainUnlockEvent"
	_neo_lock             = "Lock"
	_neo_lock2            = "LockEvent"
	_neo_unlock           = "UnlockEvent"
	_neo_unlock2          = "Unlock"
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

func (this *NeoChainListen) GetBackwardBlockNumber() uint64 {
	return this.neoCfg.BackwardBlockNumber
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
		if err != nil {
			continue
		}
		for _, exeitem := range appLog.Executions {
			for _, notify := range exeitem.Notifications {
				if notify.Contract[2:] == this.neoCfg.WrapperContract {
					if len(notify.State.Value) <= 0 {
						continue
					}
					contractMethod := this.parseNeoMethod(notify.State.Value[0].Value)
					switch contractMethod {
					case _neo_crosschainlock:
						logs.Info("from chain: %s, txhash: %s\n", this.GetChainName(), tx.Txid[2:])
						if len(notify.State.Value) < 6 {
							continue
						}
						xx, _ := strconv.ParseUint(notify.State.Value[3].Value, 10, 64)
						wrapperTransactions = append(wrapperTransactions, &models.WrapperTransaction{
							Hash:         tx.Txid[2:],
							User:         notify.State.Value[4].Value,
							SrcChainId:   xx,
							DstChainId:   xx,
							FeeTokenHash: notify.State.Value[4].Value,
							FeeAmount:    &models.BigInt{*big.NewInt(int64(xx))},
						})
					}
				} else if notify.Contract[2:] == this.neoCfg.ProxyContract {
					if len(notify.State.Value) <= 0 {
						continue
					}
					contractMethod := this.parseNeoMethod(notify.State.Value[0].Value)
					switch contractMethod {
					case _neo_crosschainlock:
						logs.Info("from chain: %s, txhash: %s\n", this.GetChainName(), tx.Txid[2:])
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
								fctransfer.Hash = tx.Txid[2:]
								fctransfer.From = utils.Hash2Address(this.GetChainId(), notifynew.State.Value[2].Value)
								fctransfer.To = utils.Hash2Address(this.GetChainId(), notify.State.Value[2].Value)
								fctransfer.Asset = utils.HexStringReverse(notifynew.State.Value[1].Value)
								amount := big.NewInt(0)
								if notifynew.State.Value[6].Type == "Integer" {
									amount, _ = new(big.Int).SetString(notifynew.State.Value[6].Value, 10)
								} else {
									amount, _ = new(big.Int).SetString(utils.HexStringReverse(notifynew.State.Value[6].Value), 16)
								}
								fctransfer.Amount = &models.BigInt{*amount}
								tchainId, _ := strconv.ParseUint(notifynew.State.Value[3].Value, 10, 32)
								fctransfer.DstChainId = uint64(tchainId)
								if len(notifynew.State.Value[5].Value) != 40 {
									continue
								}
								fctransfer.DstUser = utils.Hash2Address(uint64(tchainId), notifynew.State.Value[5].Value)
								fctransfer.DstAsset = notifynew.State.Value[4].Value
								break
							}
						}
						fctx := &models.SrcTransaction{}
						fctx.ChainId = this.GetChainId()
						fctx.Hash = tx.Txid[2:]
						fctx.State = 1
						fctx.Fee = &models.BigInt{*big.NewInt(int64(utils.String2Float64(exeitem.GasConsumed)))}
						fctx.Time = uint64(tt)
						fctx.Height = height
						fctx.User = fctransfer.From
						toChainId, _ := strconv.ParseInt(notify.State.Value[3].Value, 10, 64)
						fctx.DstChainId = uint64(toChainId)
						fctx.Contract = notify.State.Value[2].Value
						fctx.Key = notify.State.Value[4].Value
						fctx.Param = notify.State.Value[5].Value
						fctx.SrcTransfer = fctransfer
						srcTransactions = append(srcTransactions, fctx)
					case _neo_crosschainunlock:
						logs.Info("to chain: %s, txhash: %s\n", this.GetChainName(), tx.Txid[2:])
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
								tctransfer.Hash = tx.Txid[2:]
								tctransfer.From = utils.Hash2Address(this.GetChainId(), notify.State.Value[2].Value)
								tctransfer.To = utils.Hash2Address(this.GetChainId(), notifynew.State.Value[2].Value)
								tctransfer.Asset = utils.HexStringReverse(notifynew.State.Value[1].Value)
								//amount, _ := strconv.ParseUint(common.HexStringReverse(notifynew.State.Value[3].Value), 16, 64)
								amount := big.NewInt(0)
								if notifynew.State.Value[3].Type == "Integer" {
									amount, _ = new(big.Int).SetString(notifynew.State.Value[3].Value, 10)
								} else {
									amount, _ = new(big.Int).SetString(utils.HexStringReverse(notifynew.State.Value[3].Value), 16)
								}
								tctransfer.Amount = &models.BigInt{*amount}
								break
							}
						}
						tctx := &models.DstTransaction{}
						tctx.ChainId = this.GetChainId()
						tctx.Hash = tx.Txid[2:]
						tctx.State = 1
						tctx.Fee = &models.BigInt{*big.NewInt(int64(utils.String2Float64(exeitem.GasConsumed)))}
						tctx.Time = uint64(tt)
						tctx.Height = height
						fchainId, _ := strconv.ParseUint(notify.State.Value[1].Value, 10, 32)
						tctx.SrcChainId = uint64(fchainId)
						tctx.Contract = utils.HexStringReverse(notify.State.Value[2].Value)
						tctx.PolyHash = utils.HexStringReverse(notify.State.Value[3].Value)
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
	Code int64 `json:"code"`
	Message string `json:"message"`
}
type ExtendHeight struct {
	LastHeight uint64 `json:"result"`
	Error *Error `json:"error"`
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
