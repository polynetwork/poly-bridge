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

package ethereumlisten

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"poly-swap/chainsdk"
	"poly-swap/conf"
	"poly-swap/crosschainlisten/ethereumlisten/eccm_abi"
	"poly-swap/crosschainlisten/ethereumlisten/lock_proxy_abi"
	"poly-swap/crosschainlisten/ethereumlisten/wrapper_abi"
	"poly-swap/models"
	"poly-swap/utils"
	"strings"
)

const (
	_eth_crosschainlock   = "CrossChainLockEvent"
	_eth_crosschainunlock = "CrossChainUnlockEvent"
	_eth_lock             = "LockEvent"
	_eth_unlock           = "UnlockEvent"
)

type EthereumChainListen struct {
	ethCfg *conf.ChainListenConfig
	ethSdk *chainsdk.EthereumSdkPro
}

func NewEthereumChainListen(cfg *conf.ChainListenConfig) *EthereumChainListen {
	ethListen := &EthereumChainListen{}
	ethListen.ethCfg = cfg
	//
	urls := cfg.GetNodesUrl()
	sdk := chainsdk.NewEthereumSdkPro(urls, cfg.ListenSlot, cfg.ChainId)
	ethListen.ethSdk = sdk
	return ethListen
}

func (this *EthereumChainListen) GetLatestHeight() (uint64, error) {
	return this.ethSdk.GetLatestHeight()
}

func (this *EthereumChainListen) GetBackwardBlockNumber() uint64 {
	return this.ethCfg.BackwardBlockNumber
}

func (this *EthereumChainListen) GetChainListenSlot() uint64 {
	return this.ethCfg.ListenSlot
}

func (this *EthereumChainListen) GetChainId() uint64 {
	return this.ethCfg.ChainId
}

func (this *EthereumChainListen) GetChainName() string {
	return this.ethCfg.ChainName
}

func (this *EthereumChainListen) HandleNewBlock(height uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, error) {
	blockHeader, err := this.ethSdk.GetHeaderByNumber(height)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	if blockHeader == nil {
		return nil, nil, nil, nil, fmt.Errorf("there is no ethereum block!")
	}
	tt := blockHeader.Time
	wrapperTransactions, err := this.getWapperEventByBlockNumber(this.ethCfg.WrapperContract, height)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	for _, item := range wrapperTransactions {
		item.Time = tt
		item.BlockHeight = height
		item.SrcChainId = this.GetChainId()
		item.Status = conf.STATE_SOURCE_DONE
	}
	eccmLockEvents, eccmUnLockEvents, err := this.getECCMEventByBlockNumber(this.ethCfg.CCMContract, height)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	proxyLockEvents, proxyUnlockEvents, err := this.getProxyEventByBlockNumber(this.ethCfg.ProxyContract, height)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	//
	srcTransactions := make([]*models.SrcTransaction, 0)
	dstTransactions := make([]*models.DstTransaction, 0)
	for _, lockEvent := range eccmLockEvents {
		if lockEvent.Method == _eth_crosschainlock {
			logs.Info("from chain %s: txhash: %s, txid: %s\n", this.GetChainName(), lockEvent.TxHash, lockEvent.Txid)
			srcTransaction := &models.SrcTransaction{}
			srcTransaction.ChainId = this.GetChainId()
			srcTransaction.Hash = lockEvent.TxHash
			srcTransaction.State = 1
			srcTransaction.Fee = models.NewBigIntFromInt(int64(lockEvent.Fee))
			srcTransaction.Time = tt
			srcTransaction.Height = height
			srcTransaction.User = utils.Hash2Address(this.GetChainId(), lockEvent.User)
			srcTransaction.DstChainId = uint64(lockEvent.Tchain)
			srcTransaction.Contract = lockEvent.Contract
			srcTransaction.Key = lockEvent.Txid
			srcTransaction.Param = hex.EncodeToString(lockEvent.Value)
			for _, v := range proxyLockEvents {
				if v.TxHash == lockEvent.TxHash {
					toAssetHash := v.ToAssetHash
					srcTransfer := &models.SrcTransfer{}
					srcTransfer.Time = tt
					srcTransfer.ChainId = this.GetChainId()
					srcTransfer.Hash = lockEvent.TxHash
					srcTransfer.From = utils.Hash2Address(this.GetChainId(), v.FromAddress)
					srcTransfer.To = utils.Hash2Address(this.GetChainId(), lockEvent.Contract)
					srcTransfer.Asset = strings.ToLower(v.FromAssetHash)
					srcTransfer.Amount = models.NewBigInt(v.Amount)
					srcTransfer.DstChainId = uint64(v.ToChainId)
					srcTransfer.DstAsset = toAssetHash
					srcTransfer.DstUser = utils.Hash2Address(uint64(v.ToChainId), v.ToAddress)
					srcTransaction.SrcTransfer = srcTransfer
					break
				}
			}
			srcTransactions = append(srcTransactions, srcTransaction)
		}
	}
	// save unLockEvent to db
	for _, unLockEvent := range eccmUnLockEvents {
		if unLockEvent.Method == _eth_crosschainunlock {
			logs.Info("to chain %s: txhash: %s\n", this.GetChainName(), unLockEvent.TxHash)
			dstTransaction := &models.DstTransaction{}
			dstTransaction.ChainId = this.GetChainId()
			dstTransaction.Hash = unLockEvent.TxHash
			dstTransaction.State = 1
			dstTransaction.Fee = &models.BigInt{*big.NewInt(int64(unLockEvent.Fee))}
			dstTransaction.Time = tt
			dstTransaction.Height = height
			dstTransaction.SrcChainId = uint64(unLockEvent.FChainId)
			dstTransaction.Contract = unLockEvent.Contract
			dstTransaction.PolyHash = unLockEvent.RTxHash
			for _, v := range proxyUnlockEvents {
				if v.TxHash == unLockEvent.TxHash {
					dstTransfer := &models.DstTransfer{}
					dstTransfer.Hash = unLockEvent.TxHash
					dstTransfer.Time = tt
					dstTransfer.ChainId = this.GetChainId()
					dstTransfer.From = utils.Hash2Address(this.GetChainId(), unLockEvent.Contract)
					dstTransfer.To = utils.Hash2Address(this.GetChainId(), v.ToAddress)
					dstTransfer.Asset = strings.ToLower(v.ToAssetHash)
					dstTransfer.Amount = &models.BigInt{*v.Amount}
					dstTransaction.DstTransfer = dstTransfer
					break
				}
			}
			dstTransactions = append(dstTransactions, dstTransaction)
		}
	}
	return wrapperTransactions, srcTransactions, nil, dstTransactions, nil
}

func (this *EthereumChainListen) getWapperEventByBlockNumber(contractAddr string, height uint64) ([]*models.WrapperTransaction, error) {
	if len(contractAddr) == 0 {
		return nil, nil
	}
	wrapperAddress := common.HexToAddress(contractAddr)
	wrapperContract, err := polywrapper.NewIPolyWrapper(wrapperAddress, this.ethSdk.GetClient())
	if err != nil {
		return nil, fmt.Errorf("GetSmartContractEventByBlock, error: %s", err.Error())
	}
	opt := &bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: context.Background(),
	}
	// get ethereum lock events from given block
	wrapperTransactions := make([]*models.WrapperTransaction, 0)
	lockEvents, err := wrapperContract.FilterPolyWrapperLock(opt, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("GetSmartContractEventByBlock, filter lock events :%s", err.Error())
	}
	for lockEvents.Next() {
		evt := lockEvents.Event
		wrapperTransactions = append(wrapperTransactions, &models.WrapperTransaction{
			Hash:         evt.Raw.TxHash.String()[2:],
			User:         strings.ToLower(evt.Sender.String()[2:]),
			DstChainId:   evt.ToChainId,
			FeeTokenHash: evt.FromAsset.String()[2:],
			FeeAmount:    &models.BigInt{*evt.Fee},
		})
	}
	speedupEvents, err := wrapperContract.FilterPolyWrapperSpeedUp(opt, nil, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("GetSmartContractEventByBlock, filter lock events :%s", err.Error())
	}
	for speedupEvents.Next() {
		evt := speedupEvents.Event
		wrapperTransactions = append(wrapperTransactions, &models.WrapperTransaction{
			Hash:         evt.TxHash.String(),
			User:         evt.Sender.String(),
			FeeTokenHash: evt.FromAsset.String(),
			FeeAmount:    &models.BigInt{*evt.Efee},
		})
	}
	return wrapperTransactions, nil
}

func (this *EthereumChainListen) getECCMEventByBlockNumber(contractAddr string, height uint64) ([]*models.ECCMLockEvent, []*models.ECCMUnlockEvent, error) {
	eccmContractAddress := common.HexToAddress(contractAddr)
	eccmContract, err := eccm_abi.NewEthCrossChainManager(eccmContractAddress, this.ethSdk.GetClient())
	if err != nil {
		return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, error: %s", err.Error())
	}
	opt := &bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: context.Background(),
	}
	// get ethereum lock events from given block
	eccmLockEvents := make([]*models.ECCMLockEvent, 0)
	crossChainEvents, err := eccmContract.FilterCrossChainEvent(opt, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, filter lock events :%s", err.Error())
	}
	for crossChainEvents.Next() {
		evt := crossChainEvents.Event
		Fee := this.GetConsumeGas(evt.Raw.TxHash)
		eccmLockEvents = append(eccmLockEvents, &models.ECCMLockEvent{
			Method:   _eth_crosschainlock,
			Txid:     hex.EncodeToString(evt.TxId),
			TxHash:   evt.Raw.TxHash.String()[2:],
			User:     strings.ToLower(evt.Sender.String()[2:]),
			Tchain:   uint32(evt.ToChainId),
			Contract: strings.ToLower(evt.ProxyOrAssetContract.String()[2:]),
			Value:    evt.Rawdata,
			Height:   height,
			Fee:      Fee,
		})
	}
	// ethereum unlock events from given block
	eccmUnlockEvents := make([]*models.ECCMUnlockEvent, 0)
	executeTxEvent, err := eccmContract.FilterVerifyHeaderAndExecuteTxEvent(opt)
	if err != nil {
		return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, filter unlock events :%s", err.Error())
	}

	for executeTxEvent.Next() {
		evt := executeTxEvent.Event
		Fee := this.GetConsumeGas(evt.Raw.TxHash)
		eccmUnlockEvents = append(eccmUnlockEvents, &models.ECCMUnlockEvent{
			Method:   _eth_crosschainunlock,
			TxHash:   evt.Raw.TxHash.String()[2:],
			RTxHash:  utils.HexStringReverse(hex.EncodeToString(evt.CrossChainTxHash)),
			Contract: hex.EncodeToString(evt.ToContract),
			FChainId: uint32(evt.FromChainID),
			Height:   height,
			Fee:      Fee,
		})
	}
	return eccmLockEvents, eccmUnlockEvents, nil
}

func (this *EthereumChainListen) getProxyEventByBlockNumber(contractAddr string, height uint64) ([]*models.ProxyLockEvent, []*models.ProxyUnlockEvent, error) {
	proxyAddress := common.HexToAddress(contractAddr)
	proxyContract, err := lock_proxy_abi.NewLockProxy(proxyAddress, this.ethSdk.GetClient())
	if err != nil {
		return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, error: %s", err.Error())
	}
	opt := &bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: context.Background(),
	}
	// get ethereum lock events from given block
	proxyLockEvents := make([]*models.ProxyLockEvent, 0)
	lockEvents, err := proxyContract.FilterLockEvent(opt)
	if err != nil {
		return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, filter lock events :%s", err.Error())
	}
	for lockEvents.Next() {
		evt := lockEvents.Event
		proxyLockEvents = append(proxyLockEvents, &models.ProxyLockEvent{
			Method:        _eth_lock,
			TxHash:        evt.Raw.TxHash.String()[2:],
			FromAddress:   evt.FromAddress.String()[2:],
			FromAssetHash: evt.FromAssetHash.String()[2:],
			ToChainId:     uint32(evt.ToChainId),
			ToAssetHash:   hex.EncodeToString(evt.ToAssetHash),
			ToAddress:     hex.EncodeToString(evt.ToAddress),
			Amount:        evt.Amount,
		})
	}

	// ethereum unlock events from given block
	proxyUnlockEvents := make([]*models.ProxyUnlockEvent, 0)
	unlockEvents, err := proxyContract.FilterUnlockEvent(opt)
	if err != nil {
		return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, filter unlock events :%s", err.Error())
	}
	for unlockEvents.Next() {
		evt := unlockEvents.Event
		proxyUnlockEvents = append(proxyUnlockEvents, &models.ProxyUnlockEvent{
			Method:      _eth_unlock,
			TxHash:      evt.Raw.TxHash.String()[2:],
			ToAssetHash: evt.ToAssetHash.String()[2:],
			ToAddress:   evt.ToAddress.String()[2:],
			Amount:      evt.Amount,
		})
	}
	return proxyLockEvents, proxyUnlockEvents, nil
}
func (this *EthereumChainListen) GetConsumeGas(hash common.Hash) uint64 {
	tx, err := this.ethSdk.GetTransactionByHash(hash)
	if err != nil {
		return 0
	}
	receipt, err := this.ethSdk.GetTransactionReceipt(hash)
	if err != nil {
		return 0
	}
	return tx.GasPrice().Uint64() * receipt.GasUsed
}

type ExtendHeightRsp struct {
	Status  uint64 `json:"status,string"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

func (this *EthereumChainListen) GetExtendLatestHeight() (uint64, error) {
	if len(this.ethCfg.ExtendNodes) == 0 {
		return this.GetLatestHeight()
	}
	for i, _ := range this.ethCfg.ExtendNodes {
		height, err := this.getExtendLatestHeight(i)
		if err == nil {
			return height, nil
		}
	}
	return 0, fmt.Errorf("all extend node is not working")
}

func (this *EthereumChainListen) getExtendLatestHeight(node int) (uint64, error) {
	req, err := http.NewRequest("GET", this.ethCfg.ExtendNodes[node].Url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Accepts", "application/json")
	q := url.Values{}
	q.Add("module", "proxy")
	q.Add("action", "eth_blockNumber")
	q.Add("apikey", this.ethCfg.ExtendNodes[node].Key)
	req.URL.RawQuery = q.Encode()
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
	//fmt.Printf("resp body: %s\n", string(respBody))
	extendHeight := new(ExtendHeightRsp)
	extendHeight.Status = 1
	err = json.Unmarshal(respBody, extendHeight)
	if err != nil {
		return 0, err
	}
	if extendHeight.Status == 0 {
		return 0, fmt.Errorf(extendHeight.Result)
	}
	height, err := hexutil.DecodeBig(extendHeight.Result)
	if err != nil {
		return 0, err
	}
	return height.Uint64(), nil
}
