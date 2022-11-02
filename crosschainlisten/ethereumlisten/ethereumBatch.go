package ethereumlisten

import (
	"encoding/hex"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/go_abi/eccm_abi"
	"poly-bridge/go_abi/lock_proxy_abi"
	"poly-bridge/go_abi/nft_lock_proxy_abi"
	nftwp "poly-bridge/go_abi/nft_wrap_abi"
	"poly-bridge/go_abi/swapper_abi"
	"poly-bridge/go_abi/wrapper_abi"
	"poly-bridge/models"
	"strings"
)

//HandleNewBatchBlock
func (this *EthereumChainListen) HandleNewBatchBlock(start, end uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, int, int, error) {
	backStart := start*2 - end - 1
	if backStart > 0 {
		start = backStart
	}

	wrapperContracts := make([]common.Address, 0)
	var wrapperV1Contract common.Address
	nftWrapperContracts := make([]common.Address, 0)
	ccmContract := common.HexToAddress(this.ethCfg.CCMContract)
	lockProxyContracts := make([]common.Address, 0)
	nftLockProxyContracts := make([]common.Address, 0)
	swapContract := common.HexToAddress(this.ethCfg.SwapContract)

	for i, v := range this.ethCfg.WrapperContract {
		if len(strings.TrimSpace(v)) == 0 {
			continue
		}
		if i == 0 {
			wrapperV1Contract = common.HexToAddress(v)
		}
		wrapperContracts = append(wrapperContracts, common.HexToAddress(v))
	}

	for _, v := range this.ethCfg.NFTWrapperContract {
		if len(strings.TrimSpace(v)) == 0 {
			continue
		}
		nftWrapperContracts = append(nftWrapperContracts, common.HexToAddress(v))
	}

	for _, v := range this.ethCfg.ProxyContract {
		if len(strings.TrimSpace(v)) == 0 {
			continue
		}
		lockProxyContracts = append(lockProxyContracts, common.HexToAddress(v))
	}

	for _, v := range this.ethCfg.NFTProxyContract {
		if len(strings.TrimSpace(v)) == 0 {
			continue
		}
		nftLockProxyContracts = append(nftLockProxyContracts, common.HexToAddress(v))
	}

	filterContracts := make([]common.Address, 0)
	filterContracts = append(filterContracts, wrapperContracts...)
	filterContracts = append(filterContracts, nftWrapperContracts...)
	if ccmContract != common.HexToAddress("") {
		filterContracts = append(filterContracts, ccmContract)
	}
	filterContracts = append(filterContracts, lockProxyContracts...)
	filterContracts = append(filterContracts, nftLockProxyContracts...)
	if swapContract != common.HexToAddress("") {
		filterContracts = append(filterContracts, swapContract)
	}

	contractlogs, err := this.ethSdk.FilterLog(big.NewInt(int64(start)), big.NewInt(int64(end)), filterContracts)
	if err != nil {
		return nil, nil, nil, nil, 0, 0, err
	}
	if len(contractlogs) == 0 {
		return nil, nil, nil, nil, 0, 0, nil
	}

	wrapperTransactions, err := this.getWrapperTransactions(contractlogs, wrapperContracts, nftWrapperContracts, wrapperV1Contract)
	if err != nil {
		return nil, nil, nil, nil, 0, 0, err
	}
	eccmLockEvents, eccmUnLockEvents, err := this.getECCMEvents(contractlogs, ccmContract)
	if err != nil {
		return nil, nil, nil, nil, 0, 0, err
	}
	proxyLockEvents, proxyUnlockEvents, swapEvents, err := this.getProxyEvents(contractlogs, lockProxyContracts, nftLockProxyContracts, swapContract)
	if err != nil {
		return nil, nil, nil, nil, 0, 0, err
	}

	blockTimer := make(map[uint64]uint64, 0)
	for _, v := range wrapperTransactions {
		blockTimer[v.BlockHeight] = 0
	}
	for _, v := range eccmLockEvents {
		blockTimer[v.Height] = 0
	}
	for _, v := range eccmUnLockEvents {
		blockTimer[v.Height] = 0
	}
	for _, v := range proxyLockEvents {
		blockTimer[v.BlockNumber] = 0
	}
	for _, v := range proxyUnlockEvents {
		blockTimer[v.BlockNumber] = 0
	}
	for _, v := range swapEvents {
		blockTimer[v.BlockNumber] = 0
	}

	for k, _ := range blockTimer {
		timestamp, err := this.ethSdk.GetBlockTimeByNumber(k)
		if err != nil {
			return nil, nil, nil, nil, 0, 0, err
		}
		blockTimer[k] = timestamp
	}

	for _, item := range wrapperTransactions {
		logs.Info("(wrapper) from chain: %s, height: %d, txhash: %s", this.GetChainName(), item.BlockHeight, item.Hash)
		item.Time = blockTimer[item.BlockHeight]
		item.SrcChainId = this.GetChainId()
		item.Status = basedef.STATE_SOURCE_DONE
	}

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
	return wrapperTransactions, srcTransactions, nil, dstTransactions, len(proxyLockEvents), len(proxyUnlockEvents), nil

}

func (this *EthereumChainListen) getWrapperTransactions(contractlogs []types.Log, wrapperContracts []common.Address, nftWrapperContracts []common.Address, wrapperV1Contract common.Address) ([]*models.WrapperTransaction, error) {
	wrapperTransactions := make([]*models.WrapperTransaction, 0)
	erc20WrapperTransactions, err := this.ParseWrapperEventByLog(contractlogs, wrapperContracts, wrapperV1Contract)
	if err != nil {
		return nil, err
	}
	nftWrapperTransactions, err := this.ParseNFTWrapperEventByLog(contractlogs, nftWrapperContracts)
	if err != nil {
		return nil, err
	}
	wrapperTransactions = append(wrapperTransactions, erc20WrapperTransactions...)
	wrapperTransactions = append(wrapperTransactions, nftWrapperTransactions...)

	return wrapperTransactions, nil
}

func (this *EthereumChainListen) ParseWrapperEventByLog(contractlogs []types.Log, wrapperContracts []common.Address, wrapperV1Contract common.Address) ([]*models.WrapperTransaction, error) {
	if len(wrapperContracts) == 0 {
		return nil, nil
	}
	wrapperContractAbi, err := wrapper_abi.NewPolyWrapper(wrapperContracts[0], nil)
	if err != nil {
		return nil, fmt.Errorf("ParseWrapperEventByLog NewPolyWrapper, error: %s", err.Error())
	}

	wrapperTransactions := make([]*models.WrapperTransaction, 0)
	for _, v := range contractlogs {
		if !inSlice(v.Address, wrapperContracts...) {
			continue
		}

		switch v.Topics[0] {
		case this.eventPolyWrapperLockId:
			evt, err := wrapperContractAbi.ParsePolyWrapperLock(v)
			if err == nil {
				wrapperTransactions = append(wrapperTransactions, &models.WrapperTransaction{
					Hash:       evt.Raw.TxHash.String()[2:],
					User:       models.FormatString(strings.ToLower(evt.Sender.String()[2:])),
					DstChainId: evt.ToChainId,
					DstUser:    models.FormatString(hex.EncodeToString(evt.ToAddress)),
					FeeTokenHash: func() string {
						if !strings.EqualFold(v.Address.String(), wrapperV1Contract.String()) {
							switch this.GetChainId() {
							case basedef.METIS_CROSSCHAIN_ID:
								return "deaddeaddeaddeaddeaddeaddeaddeaddead0000"
							default:
								return "0000000000000000000000000000000000000000"
							}
						}
						return models.FormatString(strings.ToLower(evt.FromAsset.String()[2:]))
					}(),
					FeeAmount:   models.NewBigInt(evt.Fee),
					ServerId:    evt.Id.Uint64(),
					BlockHeight: evt.Raw.BlockNumber,
				})
			}
		}
	}

	return wrapperTransactions, nil
}

func (e *EthereumChainListen) ParseNFTWrapperEventByLog(contractlogs []types.Log, nftWrapperContracts []common.Address) ([]*models.WrapperTransaction, error) {
	if len(nftWrapperContracts) == 0 {
		return nil, nil
	}
	nftWrapperContractAbi, err := nftwp.NewPolyNFTWrapper(nftWrapperContracts[0], nil)
	if err != nil {
		return nil, fmt.Errorf("ParseNFTWrapperEventByLog NewPolyNFTWrapper, error: %s", err.Error())
	}

	wrapperTransactions := make([]*models.WrapperTransaction, 0)
	for _, v := range contractlogs {
		if !inSlice(v.Address, nftWrapperContracts...) {
			continue
		}
		switch v.Topics[0] {
		case e.eventNftPolyWrapperLockId:
			evt, err := nftWrapperContractAbi.ParsePolyWrapperLock(v)
			if err == nil {
				wtx := wrapLockEvent2WrapTx(evt)
				wrapperTransactions = append(wrapperTransactions, wtx)
			}
		}
	}
	return wrapperTransactions, nil
}

func (this *EthereumChainListen) getECCMEvents(contractlogs []types.Log, ccmContract common.Address) ([]*models.ECCMLockEvent, []*models.ECCMUnlockEvent, error) {
	if ccmContract == common.HexToAddress("") {
		return nil, nil, nil
	}

	ccmContractAbi, err := eccm_abi.NewEthCrossChainManager(ccmContract, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("getECCMEvents NewEthCrossChainManager, error: %s", err.Error())
	}

	eccmLockEvents := make([]*models.ECCMLockEvent, 0)
	eccmUnlockEvents := make([]*models.ECCMUnlockEvent, 0)

	for _, v := range contractlogs {
		if !inSlice(v.Address, ccmContract) {
			continue
		}
		switch v.Topics[0] {
		case this.eventCrossChainEventId:
			evt, err := ccmContractAbi.ParseCrossChainEvent(v)
			if err == nil {
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
		case this.eventVerifyHeaderAndExecuteTxEventId:
			evt, err := ccmContractAbi.ParseVerifyHeaderAndExecuteTxEvent(v)
			if err == nil {
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
		}
	}
	return eccmLockEvents, eccmUnlockEvents, nil
}
func (this *EthereumChainListen) getProxyEvents(contractlogs []types.Log, lockProxyContracts []common.Address, nftLockProxyContracts []common.Address, swapContract common.Address) ([]*models.ProxyLockEvent, []*models.ProxyUnlockEvent, []*models.SwapLockEvent, error) {

	proxyLockEvents, proxyUnlockEvents := make([]*models.ProxyLockEvent, 0), make([]*models.ProxyUnlockEvent, 0)
	lockProxyLockEvents, lockProxyUnlockEvents, err := this.ParseLockProxyEventByLog(contractlogs, lockProxyContracts)
	if err != nil {
		return nil, nil, nil, err
	}
	nftProxyLockEvents, nftProxyUnlockEvents, err := this.ParseNftProxyEventByLog(contractlogs, nftLockProxyContracts)
	if err != nil {
		return nil, nil, nil, err
	}
	swapProxyLockEvents, swapEvents, err := this.ParseSwapProxyEventByLog(contractlogs, swapContract)
	if err != nil {
		return nil, nil, nil, err
	}

	proxyLockEvents = append(proxyLockEvents, lockProxyLockEvents...)
	proxyLockEvents = append(proxyLockEvents, nftProxyLockEvents...)
	proxyLockEvents = append(proxyLockEvents, swapProxyLockEvents...)

	proxyUnlockEvents = append(proxyUnlockEvents, lockProxyUnlockEvents...)
	proxyUnlockEvents = append(proxyUnlockEvents, nftProxyUnlockEvents...)

	return proxyLockEvents, proxyUnlockEvents, swapEvents, nil

}

func (this *EthereumChainListen) ParseLockProxyEventByLog(contractlogs []types.Log, lockProxyContracts []common.Address) ([]*models.ProxyLockEvent, []*models.ProxyUnlockEvent, error) {
	if len(lockProxyContracts) == 0 {
		return nil, nil, nil
	}

	lockProxyContractAbi, err := lock_proxy_abi.NewLockProxy(lockProxyContracts[0], nil)
	if err != nil {
		return nil, nil, fmt.Errorf("ParseLockProxyEventByLog NewLockProxy, error: %s", err.Error())
	}

	proxyLockEvents := make([]*models.ProxyLockEvent, 0)
	proxyUnlockEvents := make([]*models.ProxyUnlockEvent, 0)
	for _, v := range contractlogs {
		if !inSlice(v.Address, lockProxyContracts...) {
			continue
		}
		switch v.Topics[0] {
		case this.eventLockEventId:
			evt, err := lockProxyContractAbi.ParseLockEvent(v)
			if err == nil {
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
		case this.eventUnlockEventId:
			evt, err := lockProxyContractAbi.ParseUnlockEvent(v)
			if err == nil {
				proxyUnlockEvents = append(proxyUnlockEvents, &models.ProxyUnlockEvent{
					BlockNumber: evt.Raw.BlockNumber,
					Method:      _eth_unlock,
					TxHash:      evt.Raw.TxHash.String()[2:],
					ToAssetHash: strings.ToLower(evt.ToAssetHash.String()[2:]),
					ToAddress:   strings.ToLower(evt.ToAddress.String()[2:]),
					Amount:      evt.Amount,
				})
			}
		}
	}
	return proxyLockEvents, proxyUnlockEvents, nil
}

func (this *EthereumChainListen) ParseNftProxyEventByLog(contractlogs []types.Log, nftProxyContracts []common.Address) ([]*models.ProxyLockEvent, []*models.ProxyUnlockEvent, error) {
	if len(nftProxyContracts) == 0 {
		return nil, nil, nil
	}

	nftLockProxyContractAbi, err := nft_lock_proxy_abi.NewPolyNFTLockProxy(nftProxyContracts[0], nil)
	if err != nil {
		return nil, nil, fmt.Errorf("ParseNftProxyEventByLog NewPolyNFTLockProxy, error: %s", err.Error())
	}

	proxyLockEvents := make([]*models.ProxyLockEvent, 0)
	proxyUnlockEvents := make([]*models.ProxyUnlockEvent, 0)
	for _, v := range contractlogs {
		if !inSlice(v.Address, nftProxyContracts...) {
			continue
		}
		switch v.Topics[0] {
		case this.eventNftLockEventId:
			evt, err := nftLockProxyContractAbi.ParseLockEvent(v)
			if err == nil {
				proxyLockEvent := convertLockProxyEvent(evt)
				proxyLockEvents = append(proxyLockEvents, proxyLockEvent)
			}
		case this.eventNftUnlockEventId:
			evt, err := nftLockProxyContractAbi.ParseUnlockEvent(v)
			if err == nil {
				proxyUnlockEvent := convertUnlockProxyEvent(evt)
				proxyUnlockEvents = append(proxyUnlockEvents, proxyUnlockEvent)
			}
		}
	}
	return proxyLockEvents, proxyUnlockEvents, nil
}

func (this *EthereumChainListen) ParseSwapProxyEventByLog(contractlogs []types.Log, swapContract common.Address) ([]*models.ProxyLockEvent, []*models.SwapLockEvent, error) {
	if swapContract == common.HexToAddress("") {
		return nil, nil, nil
	}

	swapperContractAbi, err := swapper_abi.NewSwapper(swapContract, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("ParseSwapProxyEventByLog NewSwapper, error: %s", err.Error())
	}

	swapLockEvents := make([]*models.SwapLockEvent, 0)
	proxyLockEvents := make([]*models.ProxyLockEvent, 0)

	for _, v := range contractlogs {
		if !inSlice(v.Address, swapContract) {
			continue
		}
		switch v.Topics[0] {
		case this.eventAddLiquidityEventId:
			evt, err := swapperContractAbi.ParseAddLiquidityEvent(v)
			if err == nil {
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
		case this.eventRemoveLiquidityEventId:
			evt, err := swapperContractAbi.ParseRemoveLiquidityEvent(v)
			if err == nil {
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
		case this.eventSwapEventId:
			evt, err := swapperContractAbi.ParseSwapEvent(v)
			if err == nil {
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
		case this.eventSwapperLockEventId:
			evt, err := swapperContractAbi.ParseLockEvent(v)
			if err == nil {
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
		}
	}
	return proxyLockEvents, swapLockEvents, nil
}

func inSlice(a common.Address, b ...common.Address) bool {
	for _, v := range b {
		if strings.EqualFold(v.String(), a.String()) {
			return true
		}
	}
	return false
}
