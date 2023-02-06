package o3listen

import (
	"encoding/hex"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/crosschainlisten/o3listen/swapper_abi"
	"poly-bridge/go_abi/eccm_abi"
	"poly-bridge/models"
	"poly-bridge/utils/addr"
	"strings"
)

func (this *O3ChainListen) HandleNewBatchBlock(start, end uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, []*models.WrapperDetail, []*models.PolyDetail, int, int, error) {
	if start*2 > end+1 {
		start = start*2 - end - 1
	}
	contractLogs, err := this.ethSdk.FilterLog(big.NewInt(int64(start)), big.NewInt(int64(end)), this.filterContracts, this.filterTopics)
	if err != nil {
		logs.Error("fail to filter log, %v", err)
		return nil, nil, nil, nil, nil, nil, 0, 0, err
	}
	eccmLockEvents, eccmUnLockEvents, err := this.getBatchECCMEventsByLogAndContractAddr(contractLogs, this.contractAddr.ccmContractAddr)
	if err != nil {
		logs.Error("fail to get eccm event, %v", err)
		return nil, nil, nil, nil, nil, nil, 0, 0, err
	}
	proxyLockEvents, proxyUnlockEvents, swapUnlockEvents, err := this.getSwapProxyEventByLog(contractLogs, this.contractAddr.swapContract)
	if err != nil {
		logs.Error("fail to get proxy event, %v", err)
		return nil, nil, nil, nil, nil, nil, 0, 0, err
	}
	blockTimer := make(map[uint64]uint64, 0)
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
	for _, v := range swapUnlockEvents {
		blockTimer[v.BlockNumber] = 0
	}
	for k := range blockTimer {
		timestamp, err := this.ethSdk.GetBlockTimeByNumber(k)
		if err != nil {
			logs.Error("fail to get block time, %v", err)
			return nil, nil, nil, nil, nil, nil, 0, 0, err
		}
		blockTimer[k] = timestamp
	}
	srcTransactions := make([]*models.SrcTransaction, 0)
	dstTransactions := make([]*models.DstTransaction, 0)
	for _, lockEvent := range eccmLockEvents {
		if lockEvent.Method == _eth_crosschainlock {
			logs.Info("(lock) from chain: %s, height: %d, txhash: %s, txid: %s", this.GetChainName(), lockEvent.Height, lockEvent.TxHash, lockEvent.Txid)
			srcTransaction := &models.SrcTransaction{}
			srcTransaction.ChainId = this.GetChainId()
			srcTransaction.Hash = lockEvent.TxHash
			srcTransaction.State = 1
			srcTransaction.Fee = models.NewBigIntFromInt(int64(lockEvent.Fee))
			srcTransaction.Time = blockTimer[lockEvent.Height]
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
					srcTransfer.ChainId = this.GetChainId()
					srcTransfer.Time = blockTimer[v.BlockNumber]
					srcTransfer.TxHash = lockEvent.TxHash
					srcTransfer.From = lockEvent.User
					srcTransfer.To = lockEvent.Contract
					srcTransfer.Asset = v.FromAssetHash
					srcTransfer.Amount = models.NewBigInt(v.Amount)
					srcTransfer.DstChainId = uint64(v.ToChainId)
					if srcTransfer.DstChainId == basedef.APTOS_CROSSCHAIN_ID {
						aptosAsset, err := hex.DecodeString(toAssetHash)
						if err == nil {
							toAssetHash = string(aptosAsset)
						}
					}
					srcTransfer.DstAsset = models.FormatAssert(toAssetHash)
					srcTransfer.DstUser = v.ToAddress
					srcTransaction.SrcTransfer = srcTransfer
					break
				}
			}
			srcTransactions = append(srcTransactions, srcTransaction)
			/*
				if srcTransaction.SrcTransfer != nil {
					srcTransactions = append(srcTransactions, srcTransaction)
				}
			*/
		}
	}
	// save unLockEvent to db
	for _, unLockEvent := range eccmUnLockEvents {
		if unLockEvent.Method == _eth_crosschainunlock {
			logs.Info("(unlock) to chain: %s, height: %d, txhash: %s", this.GetChainName(), unLockEvent.Height, unLockEvent.TxHash)
			dstTransaction := &models.DstTransaction{}
			dstTransaction.ChainId = this.GetChainId()
			dstTransaction.Hash = unLockEvent.TxHash
			dstTransaction.State = 1
			dstTransaction.Fee = models.NewBigIntFromInt(int64(unLockEvent.Fee))
			dstTransaction.Time = blockTimer[unLockEvent.Height]
			dstTransaction.Height = unLockEvent.Height
			dstTransaction.SrcChainId = uint64(unLockEvent.FChainId)
			dstTransaction.Contract = unLockEvent.Contract
			dstTransaction.PolyHash = unLockEvent.RTxHash
			for _, v := range proxyUnlockEvents {
				if v.TxHash == unLockEvent.TxHash && v.Method == _eth_unlock {
					dstTransfer := &models.DstTransfer{}
					dstTransfer.TxHash = unLockEvent.TxHash
					dstTransfer.Time = blockTimer[v.BlockNumber]
					dstTransfer.ChainId = this.GetChainId()
					dstTransfer.From = unLockEvent.Contract
					dstTransfer.To = v.ToAddress
					dstTransfer.Asset = v.ToAssetHash
					dstTransfer.Amount = models.NewBigInt(v.Amount)
					dstTransaction.DstTransfer = dstTransfer
					break
				}
			}
			for _, v := range swapUnlockEvents {
				if v.TxHash == unLockEvent.TxHash {
					dstTransfer := &models.DstSwap{}
					dstTransfer.TxHash = unLockEvent.TxHash
					dstTransfer.Time = blockTimer[v.BlockNumber]
					dstTransfer.ChainId = this.GetChainId()
					dstTransfer.PoolId = v.ToPoolId
					dstTransfer.InAsset = v.InAssetHash
					dstTransfer.InAmount = models.NewBigInt(v.InAmount)
					dstTransfer.OutAsset = v.OutAssetHash
					dstTransfer.OutAmount = models.NewBigInt(v.OutAmount)
					dstTransfer.DstChainId = v.ToChainId
					dstTransfer.DstAsset = v.ToAssetHash
					dstTransfer.DstUser = v.ToAddress
					dstTransfer.Type = v.Type
					dstTransaction.DstSwap = dstTransfer
					break
				}
			}
			dstTransactions = append(dstTransactions, dstTransaction)
			/*
				if dstTransaction.DstTransfer != nil || dstTransaction.DstSwap != nil {
					dstTransactions = append(dstTransactions, dstTransaction)
				}
			*/
		}
	}
	return nil, srcTransactions, nil, dstTransactions, nil, nil, len(srcTransactions), len(dstTransactions), nil
}

func (this *O3ChainListen) getBatchECCMEventsByLogAndContractAddr(contractLogs []types.Log, ccmContract common.Address) ([]*models.ECCMLockEvent, []*models.ECCMUnlockEvent, error) {
	if ccmContract == common.HexToAddress("") {
		return nil, nil, nil
	}

	ccmContractAbi, err := eccm_abi.NewEthCrossChainManager(ccmContract, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("getECCMEvents NewEthCrossChainManager, error: %s", err.Error())
	}

	eccmLockEvents := make([]*models.ECCMLockEvent, 0)
	eccmUnlockEvents := make([]*models.ECCMUnlockEvent, 0)
	for _, v := range contractLogs {
		if !addr.InSlice(v.Address, ccmContract) {
			continue
		}
		switch v.Topics[0] {
		case this.o3EventTopicIds.eventCrossChainEventId:
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
		case this.o3EventTopicIds.eventVerifyHeaderAndExecuteTxEventId:
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

func (this *O3ChainListen) getSwapProxyEventByLog(contractLogs []types.Log, swapContract common.Address) ([]*models.ProxyLockEvent, []*models.ProxyUnlockEvent, []*models.SwapUnlockEvent, error) {
	if swapContract == common.HexToAddress("") {
		return nil, nil, nil, nil
	}
	backend := this.ethSdk.GetClient()
	if backend == nil {
		return nil, nil, nil, fmt.Errorf("GetSmartContractEventByBlock, error: %s", "GetClient() return nil")
	}
	swapProxyContractAbi, err := swapper_abi.NewSwapProxy(swapContract, nil)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("ParseSwapProxyEventByLog NewSwapper, error: %s", err.Error())
	}
	proxyLockEvents := make([]*models.ProxyLockEvent, 0)
	proxyUnlockEvents := make([]*models.ProxyUnlockEvent, 0)
	swapUnLockEvents := make([]*models.SwapUnlockEvent, 0)
	for _, v := range contractLogs {
		if !addr.InSlice(v.Address, swapContract) {
			continue
		}
		switch v.Topics[0] {
		case this.o3EventTopicIds.eventLockEventId:
			evt, err := swapProxyContractAbi.ParseLockEvent(v)
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
			} else {
				logs.Error("fail to ParseLockEvent, chain: %s, contractAddr: %s, height: %d,  err: %v", basedef.GetChainName(this.ethCfg.ChainId), v.Address, v.BlockNumber, err)
			}
		case this.o3EventTopicIds.eventUnlockEventId:
			evt, err := swapProxyContractAbi.ParseUnlockEvent(v)
			if err == nil {
				proxyUnlockEvents = append(proxyUnlockEvents, &models.ProxyUnlockEvent{
					BlockNumber: evt.Raw.BlockNumber,
					Method:      _eth_unlock,
					TxHash:      evt.Raw.TxHash.String()[2:],
					ToAssetHash: strings.ToLower(evt.ToAssetHash.String()[2:]),
					ToAddress:   strings.ToLower(evt.ToAddress.String()[2:]),
					Amount:      evt.Amount,
				})
			} else {
				logs.Error("fail to ParseUnlockEvent, chain: %s, contractAddr: %s, height: %d,  err: %v", basedef.GetChainName(this.ethCfg.ChainId), v.Address, v.BlockNumber, err)
			}
		case this.o3EventTopicIds.eventAddLiquidityEventId:
			evt, err := swapProxyContractAbi.ParseAddLiquidityEvent(v)
			if err == nil {
				swapUnLockEvents = append(swapUnLockEvents, &models.SwapUnlockEvent{
					BlockNumber:  evt.Raw.BlockNumber,
					Type:         basedef.SWAP_ADDLIQUIDITY,
					TxHash:       evt.Raw.TxHash.String()[2:],
					ToPoolId:     evt.ToPoolId,
					InAssetHash:  strings.ToLower(evt.InAssetAddress.String()[2:]),
					InAmount:     evt.InAmount,
					OutAssetHash: strings.ToLower(evt.PoolTokenAddress.String()[2:]),
					OutAmount:    evt.OutLPAmount,
					ToChainId:    evt.ToChainId,
					ToAssetHash:  hex.EncodeToString(evt.ToAssetHash),
					ToAddress:    hex.EncodeToString(evt.ToAddress),
				})
			} else {
				logs.Error("fail to ParseAddLiquidityEvent, chain: %s, contractAddr: %s, height: %d,  err: %v", basedef.GetChainName(this.ethCfg.ChainId), v.Address, v.BlockNumber, err)
			}
		case this.o3EventTopicIds.eventRemoveLiquidityEventId:
			evt, err := swapProxyContractAbi.ParseRemoveLiquidityEvent(v)
			if err == nil {
				swapUnLockEvents = append(swapUnLockEvents, &models.SwapUnlockEvent{
					BlockNumber:  evt.Raw.BlockNumber,
					Type:         basedef.SWAP_REMOVELIQUIDITY,
					TxHash:       evt.Raw.TxHash.String()[2:],
					ToPoolId:     evt.ToPoolId,
					InAssetHash:  strings.ToLower(evt.PoolTokenAddress.String()[2:]),
					InAmount:     evt.InLPAmount,
					OutAssetHash: strings.ToLower(evt.OutAssetAddress.String()[2:]),
					OutAmount:    evt.OutAmount,
					ToChainId:    evt.ToChainId,
					ToAssetHash:  hex.EncodeToString(evt.ToAssetHash),
					ToAddress:    hex.EncodeToString(evt.ToAddress),
				})
			} else {
				logs.Error("fail to ParseRemoveLiquidityEvent, chain: %s, contractAddr: %s, height: %d,  err: %v", basedef.GetChainName(this.ethCfg.ChainId), v.Address, v.BlockNumber, err)
			}
		case this.o3EventTopicIds.eventSwapEventId:
			evt, err := swapProxyContractAbi.ParseSwapEvent(v)
			if err == nil {
				swapUnLockEvents = append(swapUnLockEvents, &models.SwapUnlockEvent{
					BlockNumber:  evt.Raw.BlockNumber,
					Type:         basedef.SWAP_SWAP,
					TxHash:       evt.Raw.TxHash.String()[2:],
					ToPoolId:     evt.ToPoolId,
					InAssetHash:  strings.ToLower(evt.InAssetAddress.String()[2:]),
					InAmount:     evt.InAmount,
					OutAssetHash: strings.ToLower(evt.OutAssetAddress.String()[2:]),
					OutAmount:    evt.OutAmount,
					ToChainId:    evt.ToChainId,
					ToAssetHash:  hex.EncodeToString(evt.ToAssetHash),
					ToAddress:    hex.EncodeToString(evt.ToAddress),
				})
			} else {
				logs.Error("fail to ParseSwapEvent, chain: %s, contractAddr: %s, height: %d,  err: %v", basedef.GetChainName(this.ethCfg.ChainId), v.Address, v.BlockNumber, err)
			}
		case this.o3EventTopicIds.eventRollBackEventId:
			evt, err := swapProxyContractAbi.ParseRollBackEvent(v)
			if err == nil {
				swapUnLockEvents = append(swapUnLockEvents, &models.SwapUnlockEvent{
					Type:         basedef.SWAP_ROLLBACK,
					TxHash:       evt.Raw.TxHash.String()[2:],
					ToPoolId:     0,
					InAssetHash:  strings.ToLower("0000000000000000000000000000000000000000"),
					InAmount:     new(big.Int).SetUint64(0),
					OutAssetHash: strings.ToLower("0000000000000000000000000000000000000000"),
					OutAmount:    new(big.Int).SetUint64(0),
					ToChainId:    evt.BackChainId,
					ToAssetHash:  hex.EncodeToString(evt.BackAssetHash),
					ToAddress:    hex.EncodeToString(evt.BackAddress),
				})
			} else {
				logs.Error("fail to ParseRollBackEvent, chain: %s, contractAddr: %s, height: %d,  err: %v", basedef.GetChainName(this.ethCfg.ChainId), v.Address, v.BlockNumber, err)
			}
		}
	}
	return proxyLockEvents, proxyUnlockEvents, swapUnLockEvents, nil
}
