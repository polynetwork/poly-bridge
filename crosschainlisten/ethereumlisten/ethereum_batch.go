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
	"fmt"
	"strings"

	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"poly-bridge/go_abi/eccm_abi"
	"poly-bridge/go_abi/lock_proxy_abi"
	"poly-bridge/go_abi/swapper_abi"
	"poly-bridge/go_abi/wrapper_abi"
	"poly-bridge/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type EthereumChainListenBatch struct {
	ethCfg *conf.ChainListenConfig
	ethSdk *chainsdk.EthereumSdkPro
}

func NewEthereumChainListenBatch(cfg *conf.ChainListenConfig) *EthereumChainListenBatch {
	ethListen := &EthereumChainListenBatch{}
	ethListen.ethCfg = cfg
	//
	urls := cfg.GetNodesUrl()
	sdk := chainsdk.NewEthereumSdkPro(urls, cfg.ListenSlot, cfg.ChainId)
	ethListen.ethSdk = sdk
	return ethListen
}

func (this *EthereumChainListenBatch) GetLatestHeight() (uint64, error) {
	return this.ethSdk.GetLatestHeight()
}

func (this *EthereumChainListenBatch) GetChainListenSlot() uint64 {
	return this.ethCfg.ListenSlot
}

func (this *EthereumChainListenBatch) GetChainId() uint64 {
	return this.ethCfg.ChainId
}

func (this *EthereumChainListenBatch) GetChainName() string {
	return this.ethCfg.ChainName
}

func (this *EthereumChainListenBatch) GetDefer() uint64 {
	return this.ethCfg.Defer
}

func (this *EthereumChainListenBatch) GetBatchSize() uint64 {
	return this.ethCfg.BatchSize
}

func (this *EthereumChainListenBatch) HandleNewBlock(heightStart uint64, heightEnd uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, error) {
	erc20WrapperTransactions, err := this.getWrapperEventByBlockNumber(this.ethCfg.WrapperContract, heightStart, heightEnd)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	eccmLockEvents, eccmUnLockEvents, err := this.getECCMEventByBlockNumber(this.ethCfg.CCMContract, heightStart, heightEnd)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	proxyLockEvents, proxyUnlockEvents := make([]*models.ProxyLockEvent, 0), make([]*models.ProxyUnlockEvent, 0)
	erc20ProxyLockEvents, erc20ProxyUnlockEvents, err := this.getProxyEventByBlockNumber(this.ethCfg.ProxyContract, heightStart, heightEnd)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	proxyLockEvents = append(proxyLockEvents, erc20ProxyLockEvents...)
	proxyUnlockEvents = append(proxyUnlockEvents, erc20ProxyUnlockEvents...)
	swapLockEvents, swapEvents, err := this.getSwapEventByBlockNumber(this.ethCfg.SwapContract, heightStart, heightEnd)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	proxyLockEvents = append(proxyLockEvents, swapLockEvents...)

	//
	wrapperTransactions := make([]*models.WrapperTransaction, 0)
	srcTransactions := make([]*models.SrcTransaction, 0)
	dstTransactions := make([]*models.DstTransaction, 0)
	for _, lockEvent := range eccmLockEvents {
		if lockEvent.Method == _eth_crosschainlock {
			logs.Info("(lock) from chain: %s, txhash: %s, txid: %s", this.GetChainName(), lockEvent.TxHash, lockEvent.Txid)
			blockHeader, err := this.ethSdk.GetHeaderByNumber(lockEvent.Height)
			if err != nil {
				logs.Error("get header by number err: %v", err)
			}
			if blockHeader == nil {
				logs.Error("there is no ethereum block!")
			}
			tt := blockHeader.Time
			srcTransaction := &models.SrcTransaction{}
			srcTransaction.ChainId = this.GetChainId()
			srcTransaction.Hash = lockEvent.TxHash
			srcTransaction.State = 1
			srcTransaction.Fee = models.NewBigIntFromInt(int64(lockEvent.Fee))
			srcTransaction.Time = tt
			srcTransaction.Height = lockEvent.Height
			srcTransaction.User = lockEvent.User
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
					srcTransfer.TxHash = lockEvent.TxHash
					srcTransfer.From = lockEvent.User
					srcTransfer.To = lockEvent.Contract
					srcTransfer.Asset = v.FromAssetHash
					srcTransfer.Amount = models.NewBigInt(v.Amount)
					srcTransfer.DstChainId = uint64(v.ToChainId)
					srcTransfer.DstAsset = toAssetHash
					srcTransfer.DstUser = v.ToAddress
					srcTransaction.SrcTransfer = srcTransfer
					break
				}
			}
			for _, v := range erc20WrapperTransactions {
				if v.Hash == lockEvent.TxHash {
					logs.Info("(wrapper) from chain: %s, txhash: %s", this.GetChainName(), v.Hash)
					v.Time = tt
					v.SrcChainId = this.GetChainId()
					v.Status = basedef.STATE_SOURCE_DONE
					wrapperTransactions = append(wrapperTransactions, v)
					break
				}
			}
			for _, v := range swapEvents {
				if v.TxHash == lockEvent.TxHash {
					srcSwapTransfer := &models.SrcSwap{}
					srcSwapTransfer.Time = tt
					srcSwapTransfer.ChainId = this.GetChainId()
					srcSwapTransfer.TxHash = lockEvent.TxHash
					srcSwapTransfer.From = lockEvent.User
					srcSwapTransfer.To = lockEvent.Contract
					srcSwapTransfer.Asset = v.FromAssetHash
					srcSwapTransfer.Amount = models.NewBigInt(v.Amount)
					srcSwapTransfer.DstChainId = v.ToChainId
					srcSwapTransfer.DstUser = v.ToAddress
					srcSwapTransfer.PoolId = v.ToPoolId
					srcTransaction.SrcSwap = srcSwapTransfer

					wrapperTransaction := &models.WrapperTransaction{}
					wrapperTransaction.Hash = lockEvent.TxHash
					wrapperTransaction.User = lockEvent.User
					wrapperTransaction.SrcChainId = this.GetChainId()
					wrapperTransaction.BlockHeight = blockHeader.Number.Uint64()
					wrapperTransaction.Time = tt
					wrapperTransaction.DstChainId = v.ToChainId
					wrapperTransaction.DstUser = v.ToAddress
					wrapperTransaction.ServerId = v.ServerId.Uint64()
					wrapperTransaction.FeeTokenHash = v.FeeAssetHash
					wrapperTransaction.FeeAmount = models.NewBigInt(v.Fee)
					wrapperTransaction.Status = basedef.STATE_SOURCE_DONE
					wrapperTransactions = append(wrapperTransactions, wrapperTransaction)
					break
				}
			}
			if srcTransaction.SrcTransfer != nil || srcTransaction.SrcSwap != nil {
				srcTransactions = append(srcTransactions, srcTransaction)
			}
		}
	}
	// save unLockEvent to db
	for _, unLockEvent := range eccmUnLockEvents {
		if unLockEvent.Method == _eth_crosschainunlock {
			logs.Info("(unlock) to chain: %s, txhash: %s", this.GetChainName(), unLockEvent.TxHash)
			blockHeader, err := this.ethSdk.GetHeaderByNumber(unLockEvent.Height)
			if err != nil {
				logs.Error("get header by number err: %v", err)
			}
			if blockHeader == nil {
				logs.Error("there is no ethereum block!")
			}
			tt := blockHeader.Time
			dstTransaction := &models.DstTransaction{}
			dstTransaction.ChainId = this.GetChainId()
			dstTransaction.Hash = unLockEvent.TxHash
			dstTransaction.State = 1
			dstTransaction.Fee = models.NewBigIntFromInt(int64(unLockEvent.Fee))
			dstTransaction.Time = tt
			dstTransaction.Height = unLockEvent.Height
			dstTransaction.SrcChainId = uint64(unLockEvent.FChainId)
			dstTransaction.Contract = unLockEvent.Contract
			dstTransaction.PolyHash = unLockEvent.RTxHash
			for _, v := range proxyUnlockEvents {
				if v.TxHash == unLockEvent.TxHash {
					dstTransfer := &models.DstTransfer{}
					dstTransfer.TxHash = unLockEvent.TxHash
					dstTransfer.Time = tt
					dstTransfer.ChainId = this.GetChainId()
					dstTransfer.From = unLockEvent.Contract
					dstTransfer.To = v.ToAddress
					dstTransfer.Asset = v.ToAssetHash
					dstTransfer.Amount = models.NewBigInt(v.Amount)
					dstTransaction.DstTransfer = dstTransfer
					break
				}
			}
			if dstTransaction.DstTransfer != nil {
				dstTransactions = append(dstTransactions, dstTransaction)
			}
		}
	}

	return wrapperTransactions, srcTransactions, nil, dstTransactions, nil
}

func (this *EthereumChainListenBatch) getWrapperEventByBlockNumber(contractAddrs []string, startHeight uint64, endHeight uint64) ([]*models.WrapperTransaction, error) {
	txs := make([]*models.WrapperTransaction, 0)
	for _, contract := range contractAddrs {
		aaa, err := this.getWrapperEventByBlockNumber1(contract, startHeight, endHeight)
		if err != nil {
			return nil, err
		}
		txs = append(txs, aaa...)
	}
	return txs, nil
}

func (this *EthereumChainListenBatch) getWrapperEventByBlockNumber1(contractAddr string, startHeight uint64, endHeight uint64) ([]*models.WrapperTransaction, error) {
	if len(contractAddr) == 0 {
		return nil, nil
	}
	wrapperAddress := common.HexToAddress(contractAddr)
	wrapperContract, err := wrapper_abi.NewIPolyWrapper(wrapperAddress, this.ethSdk.GetClient())
	if err != nil {
		return nil, fmt.Errorf("GetSmartContractEventByBlock, error: %s", err.Error())
	}
	opt := &bind.FilterOpts{
		Start:   startHeight,
		End:     &endHeight,
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
			DstUser:      hex.EncodeToString(evt.ToAddress),
			FeeTokenHash: strings.ToLower(evt.FromAsset.String()[2:]),
			FeeAmount:    models.NewBigInt(evt.Fee),
			ServerId:     evt.Id.Uint64(),
			BlockHeight:  evt.Raw.BlockNumber,
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
			FeeAmount:    models.NewBigInt(evt.Efee),
		})
	}
	return wrapperTransactions, nil
}

func (this *EthereumChainListenBatch) getECCMEventByBlockNumber(contractAddr string, startHeight uint64, endHeight uint64) ([]*models.ECCMLockEvent, []*models.ECCMUnlockEvent, error) {
	eccmContractAddress := common.HexToAddress(contractAddr)
	eccmContract, err := eccm_abi.NewEthCrossChainManager(eccmContractAddress, this.ethSdk.GetClient())
	if err != nil {
		return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, error: %s", err.Error())
	}
	opt := &bind.FilterOpts{
		Start:   startHeight,
		End:     &endHeight,
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
			Height:   evt.Raw.BlockNumber,
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
			RTxHash:  basedef.HexStringReverse(hex.EncodeToString(evt.CrossChainTxHash)),
			Contract: hex.EncodeToString(evt.ToContract),
			FChainId: uint32(evt.FromChainID),
			Height:   evt.Raw.BlockNumber,
			Fee:      Fee,
		})
	}
	return eccmLockEvents, eccmUnlockEvents, nil
}

func (this *EthereumChainListenBatch) getProxyEventByBlockNumber(contractAddr string, startHeight uint64, endHeight uint64) ([]*models.ProxyLockEvent, []*models.ProxyUnlockEvent, error) {
	proxyAddress := common.HexToAddress(contractAddr)
	proxyContract, err := lock_proxy_abi.NewLockProxy(proxyAddress, this.ethSdk.GetClient())
	if err != nil {
		return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, error: %s", err.Error())
	}
	opt := &bind.FilterOpts{
		Start:   startHeight,
		End:     &endHeight,
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
			FromAssetHash: strings.ToLower(evt.FromAssetHash.String()[2:]),
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
			ToAssetHash: strings.ToLower(evt.ToAssetHash.String()[2:]),
			ToAddress:   strings.ToLower(evt.ToAddress.String()[2:]),
			Amount:      evt.Amount,
		})
	}
	return proxyLockEvents, proxyUnlockEvents, nil
}
func (this *EthereumChainListenBatch) GetConsumeGas(hash common.Hash) uint64 {
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

func (this *EthereumChainListenBatch) getSwapEventByBlockNumber(contractAddr string, startHeight uint64, endHeight uint64) ([]*models.ProxyLockEvent, []*models.SwapEvent, error) {
	if len(contractAddr) == 0 {
		return nil, nil, nil
	}
	swapperContractAddress := common.HexToAddress(contractAddr)
	swapperContract, err := swapper_abi.NewSwapper(swapperContractAddress, this.ethSdk.GetClient())
	if err != nil {
		return nil, nil, fmt.Errorf("getSwapEventByBlockNumber, error: %s", err.Error())
	}
	opt := &bind.FilterOpts{
		Start:   startHeight,
		End:     &endHeight,
		Context: context.Background(),
	}
	// get ethereum lock events from given block
	swapLockEvents := make([]*models.SwapEvent, 0)
	{
		lockEvents, err := swapperContract.FilterAddLiquidityEvent(opt)
		if err != nil {
			return nil, nil, fmt.Errorf("getSwapEventByBlockNumber, filter lock events :%s", err.Error())
		}
		for lockEvents.Next() {
			evt := lockEvents.Event
			swapLockEvents = append(swapLockEvents, &models.SwapEvent{
				Type:          basedef.SWAP_ADDLIQUIDITY,
				TxHash:        evt.Raw.TxHash.String()[2:],
				FromAssetHash: strings.ToLower(evt.FromAssetHash.String()[2:]),
				FromAddress:   strings.ToLower(evt.FromAddress.String()[2:]),
				ToChainId:     evt.ToChainId,
				ToPoolId:      evt.ToPoolId,
				ToAddress:     hex.EncodeToString(evt.ToAddress),
				Amount:        evt.Amount,
				FeeAssetHash:  "0000000000000000000000000000000000000000",
				Fee:           evt.Fee,
				ServerId:      evt.Id,
			})
		}
	}
	{
		lockEvents, err := swapperContract.FilterRemoveLiquidityEvent(opt)
		if err != nil {
			return nil, nil, fmt.Errorf("getSwapEventByBlockNumber, filter lock events :%s", err.Error())
		}
		for lockEvents.Next() {
			evt := lockEvents.Event
			swapLockEvents = append(swapLockEvents, &models.SwapEvent{
				Type:          basedef.SWAP_REMOVELIQUIDITY,
				TxHash:        evt.Raw.TxHash.String()[2:],
				FromAssetHash: strings.ToLower(evt.FromAssetHash.String()[2:]),
				FromAddress:   strings.ToLower(evt.FromAddress.String()[2:]),
				ToChainId:     evt.ToChainId,
				ToPoolId:      evt.ToPoolId,
				ToAddress:     hex.EncodeToString(evt.ToAddress),
				Amount:        evt.Amount,
				FeeAssetHash:  "0000000000000000000000000000000000000000",
				Fee:           evt.Fee,
				ServerId:      evt.Id,
			})
		}
	}
	{
		lockEvents, err := swapperContract.FilterSwapEvent(opt)
		if err != nil {
			return nil, nil, fmt.Errorf("getSwapEventByBlockNumber, filter lock events :%s", err.Error())
		}
		for lockEvents.Next() {
			evt := lockEvents.Event
			swapLockEvents = append(swapLockEvents, &models.SwapEvent{
				Type:          basedef.SWAP_SWAP,
				TxHash:        evt.Raw.TxHash.String()[2:],
				FromAssetHash: strings.ToLower(evt.FromAssetHash.String()[2:]),
				FromAddress:   strings.ToLower(evt.FromAddress.String()[2:]),
				ToChainId:     evt.ToChainId,
				ToPoolId:      evt.ToPoolId,
				ToAddress:     hex.EncodeToString(evt.ToAddress),
				Amount:        evt.Amount,
				FeeAssetHash:  "0000000000000000000000000000000000000000",
				Fee:           evt.Fee,
				ServerId:      evt.Id,
			})
		}
	}
	proxyLockEvents := make([]*models.ProxyLockEvent, 0)
	lockEvents, err := swapperContract.FilterLockEvent(opt)
	if err != nil {
		return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, filter lock events :%s", err.Error())
	}
	for lockEvents.Next() {
		evt := lockEvents.Event
		proxyLockEvents = append(proxyLockEvents, &models.ProxyLockEvent{
			Method:        _eth_lock,
			TxHash:        evt.Raw.TxHash.String()[2:],
			FromAddress:   evt.FromAddress.String()[2:],
			FromAssetHash: strings.ToLower(evt.FromAssetHash.String()[2:]),
			ToChainId:     uint32(evt.ToChainId),
			ToAssetHash:   hex.EncodeToString(evt.ToAssetHash),
			ToAddress:     hex.EncodeToString(evt.ToAddress),
			Amount:        evt.Amount,
		})
	}
	return proxyLockEvents, swapLockEvents, nil
}
