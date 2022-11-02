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
	"io/ioutil"
	"net/http"
	"net/url"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"poly-bridge/go_abi/eccm_abi"
	"poly-bridge/go_abi/lock_proxy_abi"
	"poly-bridge/go_abi/swapper_abi"
	"poly-bridge/go_abi/wrapper_abi"
	"poly-bridge/models"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	_eth_crosschainlock   = "CrossChainLockEvent"
	_eth_crosschainunlock = "CrossChainUnlockEvent"
	_eth_lock             = "LockEvent"
	_eth_unlock           = "UnlockEvent"
)

type EthereumChainListen struct {
	ethCfg                               *conf.ChainListenConfig
	ethSdk                               *chainsdk.EthereumSdkPro
	eventPolyWrapperLockId               common.Hash
	eventNftPolyWrapperLockId            common.Hash
	eventCrossChainEventId               common.Hash
	eventVerifyHeaderAndExecuteTxEventId common.Hash
	eventLockEventId                     common.Hash
	eventUnlockEventId                   common.Hash
	eventNftLockEventId                  common.Hash
	eventNftUnlockEventId                common.Hash
	eventAddLiquidityEventId             common.Hash
	eventRemoveLiquidityEventId          common.Hash
	eventSwapEventId                     common.Hash
	eventSwapperLockEventId              common.Hash
}

func NewEthereumChainListen(cfg *conf.ChainListenConfig) *EthereumChainListen {
	ethListen := &EthereumChainListen{}
	ethListen.ethCfg = cfg
	//
	urls := cfg.GetNodesUrl()
	sdk := chainsdk.NewEthereumSdkPro(urls, cfg.ListenSlot, cfg.ChainId)
	ethListen.ethSdk = sdk
	ethListen.eventPolyWrapperLockId = common.HexToHash("0x2b0591052cc6602e870d3994f0a1b173fdac98c215cb3b0baf84eaca5a0aa81e")
	ethListen.eventNftPolyWrapperLockId = common.HexToHash("0x3a15d8cf4b167dd8963989f8038f2333a4889f74033bb53bfb767a5cced072e2")
	ethListen.eventCrossChainEventId = common.HexToHash("0x6ad3bf15c1988bc04bc153490cab16db8efb9a3990215bf1c64ea6e28be88483")
	ethListen.eventVerifyHeaderAndExecuteTxEventId = common.HexToHash("0x8a4a2663ce60ce4955c595da2894de0415240f1ace024cfbff85f513b656bdae")
	ethListen.eventLockEventId = common.HexToHash("0x8636abd6d0e464fe725a13346c7ac779b73561c705506044a2e6b2cdb1295ea5")
	ethListen.eventUnlockEventId = common.HexToHash("0xd90288730b87c2b8e0c45bd82260fd22478aba30ae1c4d578b8daba9261604df")
	ethListen.eventNftLockEventId = common.HexToHash("0x98081b3037dc78e7a7ffa56932222cfc7ea9325ad6a3e7b0b3b4e3e678d7fd13")
	ethListen.eventNftUnlockEventId = common.HexToHash("0xd90288730b87c2b8e0c45bd82260fd22478aba30ae1c4d578b8daba9261604df")
	ethListen.eventAddLiquidityEventId = common.HexToHash("0x7b634860445c375b3604695e3d36b0ca94d7342cacaae46d96b8727e86522d32")
	ethListen.eventRemoveLiquidityEventId = common.HexToHash("0x7ee445799431a22b707efdb3f751a430c4f01f12d902e952200041d81255a41e")
	ethListen.eventSwapEventId = common.HexToHash("0x9e37e0e96b266241aa70174d3c6d60151148a5b4181a57fb3d9475aa39ed0672")
	ethListen.eventSwapperLockEventId = common.HexToHash("0x8636abd6d0e464fe725a13346c7ac779b73561c705506044a2e6b2cdb1295ea5")

	return ethListen
}

func (this *EthereumChainListen) GetLatestHeight() (uint64, error) {
	return this.ethSdk.GetLatestHeight()
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

func (this *EthereumChainListen) GetDefer() uint64 {
	return this.ethCfg.Defer
}

func (this *EthereumChainListen) GetBatchSize() uint64 {
	return this.ethCfg.BatchSize
}

func (this *EthereumChainListen) GetBatchLength() (uint64, uint64) {
	return this.ethCfg.MinBatchLength, this.ethCfg.MaxBatchLength
}

func (this *EthereumChainListen) getPLTUnlock(tx common.Hash) *models.ProxyUnlockEvent {
	address, asset, amount, err := this.GetPaletteLockProxyUnlockEvent(tx)
	if err != nil {
		logs.Error("Get palette lock proxy event error %v", err)
		return nil
	}
	return &models.ProxyUnlockEvent{
		Amount:      amount,
		ToAddress:   strings.ToLower(address.String()[2:]),
		ToAssetHash: strings.ToLower(asset.String()[2:]),
	}
}

func (this *EthereumChainListen) HandleNewBlock(height uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, []*models.WrapperDetail, []*models.PolyDetail, int, int, error) {
	startHeight := height - 2
	endHeight := height

	blockTimer := make(map[uint64]uint64)
	for i := startHeight; i <= endHeight; i++ {
		timestamp, err := this.ethSdk.GetBlockTimeByNumber(i)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 0, 0, err
		}
		blockTimer[i] = timestamp
	}

	wrapperTransactions := make([]*models.WrapperTransaction, 0)
	erc20WrapperTransactions, err := this.getWrapperEventByBlockNumber(this.ethCfg.WrapperContract, startHeight, endHeight)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, 0, 0, err
	}
	nftWrapperTransactions, err := this.getNFTWrapperEventByBlockNumber(this.ethCfg.NFTWrapperContract, startHeight, endHeight)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, 0, 0, err
	}
	wrapperTransactions = append(wrapperTransactions, erc20WrapperTransactions...)
	wrapperTransactions = append(wrapperTransactions, nftWrapperTransactions...)

	for _, item := range wrapperTransactions {
		logs.Info("(wrapper) from chain: %s, height: %d, txhash: %s", this.GetChainName(), item.BlockHeight, item.Hash)
		item.Time = blockTimer[item.BlockHeight]
		item.SrcChainId = this.GetChainId()
		item.Status = basedef.STATE_SOURCE_DONE
	}
	eccmLockEvents, eccmUnLockEvents, err := this.getECCMEventByBlockNumber(this.ethCfg.CCMContract, startHeight, endHeight)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, 0, 0, err
	}

	proxyLockEvents, proxyUnlockEvents := make([]*models.ProxyLockEvent, 0), make([]*models.ProxyUnlockEvent, 0)
	erc20ProxyLockEvents, erc20ProxyUnlockEvents, err := this.getProxyEventByBlockNumber(startHeight, endHeight)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, 0, 0, err
	}

	nftProxyLockEvents, nftProxyUnlockEvents, err := this.getNFTProxyEventByBlockNumber(this.ethCfg.NFTProxyContract, startHeight, endHeight)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, 0, 0, err
	}

	proxyLockEvents = append(proxyLockEvents, erc20ProxyLockEvents...)
	proxyUnlockEvents = append(proxyUnlockEvents, erc20ProxyUnlockEvents...)
	proxyLockEvents = append(proxyLockEvents, nftProxyLockEvents...)
	proxyUnlockEvents = append(proxyUnlockEvents, nftProxyUnlockEvents...)

	swapLockEvents, swapEvents, err := this.getSwapEventByBlockNumber(this.ethCfg.SwapContract, startHeight, endHeight)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, 0, 0, err
	}
	proxyLockEvents = append(proxyLockEvents, swapLockEvents...)

	//
	srcTransactions := make([]*models.SrcTransaction, 0)
	dstTransactions := make([]*models.DstTransaction, 0)
	for _, lockEvent := range eccmLockEvents {
		logs.Info("(lock, method: %s) from chain: %s, height: %d, txhash: %s, txid: %s", lockEvent.Method, this.GetChainName(), lockEvent.Height, lockEvent.TxHash, lockEvent.Txid)
		if lockEvent.Method == _eth_crosschainlock {
			srcTransaction := &models.SrcTransaction{}
			srcTransaction.ChainId = this.GetChainId()
			srcTransaction.Hash = lockEvent.TxHash
			srcTransaction.State = 1
			srcTransaction.Fee = models.NewBigIntFromInt(int64(lockEvent.Fee))
			srcTransaction.Time = blockTimer[lockEvent.Height]
			srcTransaction.Height = lockEvent.Height
			srcTransaction.User = models.FormatString(lockEvent.User)
			srcTransaction.DstChainId = uint64(lockEvent.Tchain)
			srcTransaction.Contract = models.FormatString(lockEvent.Contract)
			srcTransaction.Key = lockEvent.Txid
			srcTransaction.Param = hex.EncodeToString(lockEvent.Value)
			var lock *models.ProxyLockEvent
			if srcTransaction.ChainId == basedef.PLT_CROSSCHAIN_ID && !this.isNFTECCMLockEvent(lockEvent) {
				// TODO: with retry later
				lock, _ = this.GetPaletteLockProxyLockEvent(common.HexToHash("0x" + lockEvent.TxHash))
			} else {
				for _, v := range proxyLockEvents {
					if v.TxHash == lockEvent.TxHash {
						lock = v
						break
					}
				}
			}
			if lock != nil {
				toAssetHash := lock.ToAssetHash
				srcTransfer := &models.SrcTransfer{}
				srcTransfer.Time = blockTimer[lock.BlockNumber]
				srcTransfer.ChainId = this.GetChainId()
				srcTransfer.TxHash = lockEvent.TxHash
				srcTransfer.From = models.FormatString(lockEvent.User)
				srcTransfer.To = models.FormatString(lockEvent.Contract)
				srcTransfer.Asset = models.FormatString(lock.FromAssetHash)
				srcTransfer.Amount = models.NewBigInt(lock.Amount)
				srcTransfer.DstChainId = uint64(lock.ToChainId)
				srcTransfer.DstAsset = models.FormatString(toAssetHash)
				srcTransfer.DstUser = models.FormatString(lock.ToAddress)
				srcTransaction.SrcTransfer = srcTransfer
				if this.isNFTECCMLockEvent(lockEvent) {
					srcTransaction.Standard = models.TokenTypeErc721
					srcTransaction.SrcTransfer.Standard = models.TokenTypeErc721
				}
			}

			for _, v := range swapEvents {
				if v.TxHash == lockEvent.TxHash {
					srcSwapTransfer := &models.SrcSwap{}
					srcSwapTransfer.Time = blockTimer[v.BlockNumber]
					srcSwapTransfer.ChainId = this.GetChainId()
					srcSwapTransfer.TxHash = lockEvent.TxHash
					srcSwapTransfer.From = models.FormatString(lockEvent.User)
					srcSwapTransfer.To = models.FormatString(lockEvent.Contract)
					srcSwapTransfer.Asset = models.FormatString(v.FromAssetHash)
					srcSwapTransfer.Amount = models.NewBigInt(v.Amount)
					srcSwapTransfer.DstChainId = v.ToChainId
					srcSwapTransfer.DstUser = models.FormatString(v.ToAddress)
					srcSwapTransfer.PoolId = v.ToPoolId
					srcSwapTransfer.Type = v.Type
					srcTransaction.SrcSwap = srcSwapTransfer

					wrapperTransaction := &models.WrapperTransaction{}
					wrapperTransaction.Hash = lockEvent.TxHash
					wrapperTransaction.User = models.FormatString(lockEvent.User)
					wrapperTransaction.SrcChainId = this.GetChainId()
					wrapperTransaction.BlockHeight = v.BlockNumber
					wrapperTransaction.Time = blockTimer[v.BlockNumber]
					wrapperTransaction.DstChainId = v.ToChainId
					wrapperTransaction.DstUser = models.FormatString(v.ToAddress)
					wrapperTransaction.ServerId = v.ServerId.Uint64()
					wrapperTransaction.FeeTokenHash = models.FormatString(v.FeeAssetHash)
					wrapperTransaction.FeeAmount = models.NewBigInt(v.Fee)
					wrapperTransaction.Status = basedef.STATE_SOURCE_DONE
					wrapperTransactions = append(wrapperTransactions, wrapperTransaction)
					break
				}
			}
			//opensrcTransactions
			//if srcTransaction.SrcTransfer != nil || srcTransaction.SrcSwap != nil {
			srcTransactions = append(srcTransactions, srcTransaction)
			//}
		}
	}
	// save unLockEvent to db
	for _, unLockEvent := range eccmUnLockEvents {
		logs.Info("(unlock, method: %s) to chain: %s, height: %d, txhash: %s", this.GetChainName(), unLockEvent.Method, unLockEvent.Height, unLockEvent.TxHash)
		if unLockEvent.Method == _eth_crosschainunlock {
			dstTransaction := &models.DstTransaction{}
			dstTransaction.ChainId = this.GetChainId()
			dstTransaction.Hash = unLockEvent.TxHash
			dstTransaction.State = 1
			dstTransaction.Fee = models.NewBigIntFromInt(int64(unLockEvent.Fee))
			dstTransaction.Time = blockTimer[unLockEvent.Height]
			dstTransaction.Height = unLockEvent.Height
			dstTransaction.SrcChainId = uint64(unLockEvent.FChainId)
			dstTransaction.Contract = models.FormatString(unLockEvent.Contract)
			dstTransaction.PolyHash = unLockEvent.RTxHash
			var unlock *models.ProxyUnlockEvent
			if dstTransaction.ChainId == basedef.PLT_CROSSCHAIN_ID && !this.isNFTECCMUnlockEvent(unLockEvent) {
				unlock = this.getPLTUnlock(common.HexToHash("0x" + unLockEvent.TxHash))
			} else {
				for _, v := range proxyUnlockEvents {
					if v.TxHash == unLockEvent.TxHash {
						unlock = v
						break
					}
				}
			}
			if unlock != nil {
				dstTransfer := &models.DstTransfer{}
				dstTransfer.TxHash = unLockEvent.TxHash
				dstTransfer.Time = blockTimer[unlock.BlockNumber]
				dstTransfer.ChainId = this.GetChainId()
				dstTransfer.From = models.FormatString(unLockEvent.Contract)
				dstTransfer.To = models.FormatString(unlock.ToAddress)
				dstTransfer.Asset = models.FormatString(unlock.ToAssetHash)
				dstTransfer.Amount = models.NewBigInt(unlock.Amount)
				dstTransaction.DstTransfer = dstTransfer
				if this.isNFTECCMUnlockEvent(unLockEvent) {
					dstTransaction.Standard = models.TokenTypeErc721
					dstTransaction.DstTransfer.Standard = models.TokenTypeErc721
				}
			}
			//opendstTransactions
			//if dstTransaction.DstTransfer != nil {
			dstTransactions = append(dstTransactions, dstTransaction)
			//}
		}
	}
	return wrapperTransactions, srcTransactions, nil, dstTransactions, nil, nil, len(proxyLockEvents), len(proxyUnlockEvents), nil
}

func (this *EthereumChainListen) getWrapperEventByBlockNumber(contractAddrs []string, startHeight uint64, endHeight uint64) ([]*models.WrapperTransaction, error) {
	txs := make([]*models.WrapperTransaction, 0)
	for i, contract := range contractAddrs {
		if contract == "" {
			continue
		}
		aaa, err := this.getWrapperEventByBlockNumber1(contract, startHeight, endHeight, i)
		if err != nil {
			return nil, err
		}
		txs = append(txs, aaa...)
	}
	return txs, nil
}

func (this *EthereumChainListen) getWrapperEventByBlockNumber1(contractAddr string, startHeight uint64, endHeight uint64, index int) ([]*models.WrapperTransaction, error) {
	if len(contractAddr) == 0 {
		return nil, nil
	}
	wrapperAddress := common.HexToAddress(contractAddr)
	client := this.ethSdk.GetClient()
	if client == nil {
		return nil, fmt.Errorf("getWrapperEventByBlockNumber1 GetClient error: nil")
	}
	wrapperContract, err := wrapper_abi.NewPolyWrapper(wrapperAddress, client)
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
			User:         models.FormatString(strings.ToLower(evt.Sender.String()[2:])),
			DstChainId:   evt.ToChainId,
			DstUser:      models.FormatString(hex.EncodeToString(evt.ToAddress)),
			FeeTokenHash: models.FormatString(strings.ToLower(evt.FromAsset.String()[2:])),
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
			User:         models.FormatString(evt.Sender.String()),
			FeeTokenHash: models.FormatString(evt.FromAsset.String()),
			FeeAmount:    models.NewBigInt(evt.Efee),
		})
	}
	if index != 0 {
		for _, tx := range wrapperTransactions {
			if this.GetChainId() == basedef.METIS_CROSSCHAIN_ID {
				tx.FeeTokenHash = "deaddeaddeaddeaddeaddeaddeaddeaddead0000"
			} else {
				tx.FeeTokenHash = "0000000000000000000000000000000000000000"
			}
		}
	}
	return wrapperTransactions, nil
}

func (this *EthereumChainListen) getECCMEventByBlockNumber(contractAddr string, startHeight uint64, endHeight uint64) ([]*models.ECCMLockEvent, []*models.ECCMUnlockEvent, error) {
	eccmContractAddress := common.HexToAddress(contractAddr)
	client := this.ethSdk.GetClient()
	if client == nil {
		return nil, nil, fmt.Errorf("getECCMEventByBlockNumber GetClient error: nil")
	}
	eccmContract, err := eccm_abi.NewEthCrossChainManager(eccmContractAddress, client)
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
		this.ethSdk.SetClientHeightZero(client)
		return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, filter lock events :%s", err.Error())
	}
	for crossChainEvents.Next() {
		evt := crossChainEvents.Event

		user := evt.Sender
		if evt.Sender.String() == "0x0000000000000000000000000000000000000000" {
			sender, err := this.getTxSenderByTxHash(evt.Raw.TxHash)
			if err != nil {
				logs.Error("getTxSenderByTxHash error： vv")
			} else {
				user = sender
			}
		}

		Fee := this.GetConsumeGas(evt.Raw.TxHash)
		eccmLockEvents = append(eccmLockEvents, &models.ECCMLockEvent{
			Method:   _eth_crosschainlock,
			Txid:     hex.EncodeToString(evt.TxId),
			TxHash:   evt.Raw.TxHash.String()[2:],
			User:     strings.ToLower(user.String()[2:]),
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

func (this *EthereumChainListen) getTxSenderByTxHash(txHash common.Hash) (common.Address, error) {
	client := this.ethSdk.GetClient()
	if client == nil {
		return common.Address{}, fmt.Errorf("getTxSenderByTxHash GetClient error: nil")
	}
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return common.Address{}, fmt.Errorf("TransactionReceipt error: %v", err)
	}
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return common.Address{}, fmt.Errorf("TransactionByHash error: %v", err)
	}
	sender, err := client.TransactionSender(context.Background(), tx, receipt.BlockHash, receipt.TransactionIndex)
	if err != nil {
		return common.Address{}, fmt.Errorf("TransactionSender error: %v", err)
	}
	return sender, nil
}

func (this *EthereumChainListen) getProxyEventByBlockNumber(startHeight uint64, endHeight uint64) ([]*models.ProxyLockEvent, []*models.ProxyUnlockEvent, error) {
	lockProxies := make(map[string]struct{}, 0)
	for _, proxy := range this.ethCfg.ProxyContract {
		lockProxies[proxy] = struct{}{}
	}
	for _, other := range this.ethCfg.OtherProxyContract {
		lockProxies[other.ItemProxy] = struct{}{}
	}

	erc20ProxyLockEvents, erc20ProxyUnlockEvents := make([]*models.ProxyLockEvent, 0), make([]*models.ProxyUnlockEvent, 0)
	for lockContract, _ := range lockProxies {
		if len(strings.TrimSpace(lockContract)) == 0 {
			continue
		}
		erc20ProxyLockEvents1, erc20ProxyUnlockEvents1, err := this.getProxyEventByBlockNumber1(lockContract, startHeight, endHeight)
		if err != nil {
			return nil, nil, err
		}
		erc20ProxyLockEvents = append(erc20ProxyLockEvents, erc20ProxyLockEvents1...)
		erc20ProxyUnlockEvents = append(erc20ProxyUnlockEvents, erc20ProxyUnlockEvents1...)
	}
	return erc20ProxyLockEvents, erc20ProxyUnlockEvents, nil
}

func (this *EthereumChainListen) getProxyEventByBlockNumber1(contractAddr string, startHeight uint64, endHeight uint64) ([]*models.ProxyLockEvent, []*models.ProxyUnlockEvent, error) {
	proxyAddress := common.HexToAddress(contractAddr)
	client := this.ethSdk.GetClient()
	if client == nil {
		return nil, nil, fmt.Errorf("getProxyEventByBlockNumber GetClient error: nil")
	}
	proxyContract, err := lock_proxy_abi.NewLockProxy(proxyAddress, client)
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
			BlockNumber:   evt.Raw.BlockNumber,
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
			BlockNumber: evt.Raw.BlockNumber,
			Method:      _eth_unlock,
			TxHash:      evt.Raw.TxHash.String()[2:],
			ToAssetHash: strings.ToLower(evt.ToAssetHash.String()[2:]),
			ToAddress:   strings.ToLower(evt.ToAddress.String()[2:]),
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

func (this *EthereumChainListen) getSwapEventByBlockNumber(contractAddr string, startHeight uint64, endHeight uint64) ([]*models.ProxyLockEvent, []*models.SwapLockEvent, error) {
	if len(contractAddr) == 0 {
		return nil, nil, nil
	}
	swapperContractAddress := common.HexToAddress(contractAddr)
	client := this.ethSdk.GetClient()
	if client == nil {
		return nil, nil, fmt.Errorf("getSwapEventByBlockNumber GetClient error: nil")
	}
	swapperContract, err := swapper_abi.NewSwapper(swapperContractAddress, client)
	if err != nil {
		return nil, nil, fmt.Errorf("getSwapEventByBlockNumber, error: %s", err.Error())
	}
	opt := &bind.FilterOpts{
		Start:   startHeight,
		End:     &endHeight,
		Context: context.Background(),
	}
	// get ethereum lock events from given block
	swapLockEvents := make([]*models.SwapLockEvent, 0)
	{
		lockEvents, err := swapperContract.FilterAddLiquidityEvent(opt)
		if err != nil {
			return nil, nil, fmt.Errorf("getSwapEventByBlockNumber, filter lock events :%s", err.Error())
		}
		for lockEvents.Next() {
			evt := lockEvents.Event
			swapLockEvents = append(swapLockEvents, &models.SwapLockEvent{
				BlockNumber:   evt.Raw.BlockNumber,
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
			swapLockEvents = append(swapLockEvents, &models.SwapLockEvent{
				BlockNumber:   evt.Raw.BlockNumber,
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
			swapLockEvents = append(swapLockEvents, &models.SwapLockEvent{
				BlockNumber:   evt.Raw.BlockNumber,
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
			BlockNumber:   evt.Raw.BlockNumber,
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
