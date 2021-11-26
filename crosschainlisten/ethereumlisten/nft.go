package ethereumlisten

import (
	"bytes"
	"context"
	"fmt"
	"poly-bridge/go_abi/eccm_abi"
	nftlp "poly-bridge/go_abi/nft_lock_proxy_abi"
	nftwp "poly-bridge/go_abi/nft_wrap_abi"
	"poly-bridge/models"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func allErr(errs ...error) error {
	allError := true
	var lastErr error
	for _, v := range errs {
		if v == nil {
			allError = false
		} else {
			lastErr = v
		}
	}
	if !allError {
		return nil
	} else {
		return lastErr
	}
}

func isContract(addr string) bool {
	if strings.Trim(addr, " ") == "" {
		return false
	}
	if addr == "0000000000000000000000000000000000000000" {
		return false
	}
	if len(addr) < 40 {
		return false
	}
	return true
}

func (e *EthereumChainListen) isNFTECCMLockEvent(event *models.ECCMLockEvent) bool {
	addr1 := common.HexToAddress(event.Contract)
	for _, contract := range e.ethCfg.NFTProxyContract {
		addr2 := common.HexToAddress(contract)
		if bytes.Equal(addr1.Bytes(), addr2.Bytes()) {
			return true
		}
	}
	return false
}

func (e *EthereumChainListen) isNFTECCMUnlockEvent(event *models.ECCMUnlockEvent) bool {
	addr1 := common.HexToAddress(event.Contract)
	for _, contract := range e.ethCfg.NFTProxyContract {
		addr2 := common.HexToAddress(contract)
		if bytes.Equal(addr1.Bytes(), addr2.Bytes()) {
			return true
		}
	}
	return false
}

func (e *EthereumChainListen) NFTWrapperAddress(nftWrapperContract string) common.Address {
	return common.HexToAddress(nftWrapperContract)
}

//func (e *EthereumChainListen) ECCMAddress() common.Address {
//	return common.HexToAddress(e.ethCfg.CCMContract)
//}

func (e *EthereumChainListen) NFTProxyAddress(nftProxyContract string) common.Address {
	return common.HexToAddress(nftProxyContract)
}

//func (e *EthereumChainListen) HandleNFTNewBlock(
//	height uint64, tt uint64,
//	eccmLockEvents []*models.ECCMLockEvent,
//	eccmUnLockEvents []*models.ECCMUnlockEvent) (
//	[]*models.WrapperTransaction,
//	[]*models.SrcTransaction,
//	[]*models.DstTransaction,
//	error,
//) {
//
//	wrapAddr := e.NFTWrapperAddress()
//	proxyAddr := e.NFTProxyAddress()
//	chainName := e.GetChainName()
//	chainID := e.GetChainId()
//
//	wrapperTransactions, err := e.getNFTWrapperEventByBlockNumber(wrapAddr, height, height)
//	if err != nil {
//		return nil, nil, nil, err
//	}
//	for _, wtx := range wrapperTransactions {
//		logs.Info("(wrapper) from chain: %s, txhash: %s", chainName, wtx.Hash)
//		wtx.Time = tt
//		wtx.SrcChainId = e.GetChainId()
//		wtx.Status = basedef.STATE_SOURCE_DONE
//	}
//	proxyLockEvents, proxyUnlockEvents, err := e.getNFTProxyEventByBlockNumber(proxyAddr, height, height)
//	if err != nil {
//		return nil, nil, nil, err
//	}
//
//	srcTransactions := make([]*models.SrcTransaction, 0)
//	dstTransactions := make([]*models.DstTransaction, 0)
//	for _, lockEvent := range eccmLockEvents {
//		if lockEvent.Method == _eth_crosschainlock {
//			logs.Info("(lock) from chain: %s, txhash: %s, txid: %s",
//				chainName, lockEvent.TxHash, lockEvent.Txid)
//			srcTransaction := assembleSrcTransaction(lockEvent, proxyLockEvents, chainID, tt)
//			srcTransactions = append(srcTransactions, srcTransaction)
//		}
//	}
//	// save unLockEvent to db
//	for _, unLockEvent := range eccmUnLockEvents {
//		if unLockEvent.Method == _eth_crosschainunlock {
//			logs.Info("(unlock) to chain: %s, txhash: %s", chainName, unLockEvent.TxHash)
//			dstTransaction := assembleDstTransaction(unLockEvent, proxyUnlockEvents, chainID, tt)
//			dstTransactions = append(dstTransactions, dstTransaction)
//		}
//	}
//
//	addERC721TokenStandard(wrapperTransactions, srcTransactions, dstTransactions)
//
//	return wrapperTransactions, srcTransactions, dstTransactions, nil
//}

//func (e *EthereumChainListen) HandleNFTBlockBatch(
//	startHeight, endHeight uint64,
//	eccmLockEvents []*models.ECCMLockEvent,
//	eccmUnLockEvents []*models.ECCMUnlockEvent,
//) (
//	[]*models.WrapperTransaction,
//	[]*models.SrcTransaction,
//	[]*models.DstTransaction,
//	error,
//) {
//
//	wrapAddr := e.NFTWrapperAddress()
//	proxyAddr := e.NFTProxyAddress()
//	chainName := e.GetChainName()
//	chainID := e.GetChainId()
//
//	wrapperTransactions, err := e.getNFTWrapperEventByBlockNumber(wrapAddr, startHeight, endHeight)
//	if err != nil {
//		return nil, nil, nil, err
//	}
//	for _, wtx := range wrapperTransactions {
//		logs.Info("(wrapper) from chain: %s, txhash: %s", chainName, wtx.Hash)
//		wtx.SrcChainId = e.GetChainId()
//		wtx.Status = basedef.STATE_SOURCE_DONE
//	}
//	proxyLockEvents, proxyUnlockEvents, err := e.getNFTProxyEventByBlockNumber(proxyAddr, startHeight, endHeight)
//	if err != nil {
//		return nil, nil, nil, err
//	}
//	//
//
//	srcTransactions := make([]*models.SrcTransaction, 0)
//	dstTransactions := make([]*models.DstTransaction, 0)
//	for _, lockEvent := range eccmLockEvents {
//		if lockEvent.Method == _eth_crosschainlock {
//			logs.Info("(lock) from chain: %s, txhash: %s, txid: %s", chainName, lockEvent.TxHash, lockEvent.Txid)
//			srcTransaction := assembleSrcTransaction(lockEvent, proxyLockEvents, chainID, 0)
//			srcTransactions = append(srcTransactions, srcTransaction)
//		}
//	}
//	// save unLockEvent to db
//	for _, unLockEvent := range eccmUnLockEvents {
//		if unLockEvent.Method == _eth_crosschainunlock {
//			logs.Info("(unlock) to chain: %s, txhash: %s", chainName, unLockEvent.TxHash)
//			dstTransaction := assembleDstTransaction(unLockEvent, proxyUnlockEvents, chainID, 0)
//			dstTransactions = append(dstTransactions, dstTransaction)
//		}
//	}
//
//	addERC721TokenStandard(wrapperTransactions, srcTransactions, dstTransactions)
//
//	return wrapperTransactions, srcTransactions, dstTransactions, nil
//}

func (e *EthereumChainListen) getNFTWrapperEventByBlockNumber(
	wrapAddrStrs []string,
	startHeight, endHeight uint64) (
	[]*models.WrapperTransaction,
	error,
) {
	wrapperTransactions := make([]*models.WrapperTransaction, 0)
	for _, wrapAddrStr := range wrapAddrStrs {
		if !isContract(wrapAddrStr) {
			return nil, nil
		}
		wrapAddr := common.HexToAddress(wrapAddrStr)
		wrapperContract, err := nftwp.NewPolyNFTWrapper(wrapAddr, e.ethSdk.GetClient())
		if err != nil {
			return nil, fmt.Errorf("GetSmartContractEventByBlock, error: %s", err.Error())
		}
		opt := &bind.FilterOpts{
			Start:   startHeight,
			End:     &endHeight,
			Context: context.Background(),
		}

		// get ethereum lock events from given block
		lockEvents, err := wrapperContract.FilterPolyWrapperLock(opt, nil, nil)
		if err != nil {
			return nil, fmt.Errorf("GetSmartContractEventByBlock, filter lock events :%s", err.Error())
		}
		for lockEvents.Next() {
			evt := lockEvents.Event
			wtx := wrapLockEvent2WrapTx(evt)
			wrapperTransactions = append(wrapperTransactions, wtx)
		}
		speedupEvents, err := wrapperContract.FilterPolyWrapperSpeedUp(opt, nil, nil, nil)
		if err != nil {
			return nil, fmt.Errorf("GetSmartContractEventByBlock, filter lock events :%s", err.Error())
		}
		for speedupEvents.Next() {
			evt := speedupEvents.Event
			wtx := wrapSpeedUpEvent2WrapTx(evt)
			wrapperTransactions = append(wrapperTransactions, wtx)
		}
	}
	return wrapperTransactions, nil
}

func (e *EthereumChainListen) getNFTECCMEventByBlockNumber(
	eccmAddr common.Address,
	startHeight, endHeight uint64) (
	[]*models.ECCMLockEvent,
	[]*models.ECCMUnlockEvent,
	error,
) {

	eccmContract, err := eccm_abi.NewEthCrossChainManagerImplemetation(eccmAddr, e.ethSdk.GetClient())
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
		Fee := e.GetConsumeGas(evt.Raw.TxHash)
		eccmLockEvent := crossChainEvent2ProxyLockEvent(evt, Fee)
		eccmLockEvents = append(eccmLockEvents, eccmLockEvent)
	}
	// ethereum unlock events from given block
	eccmUnlockEvents := make([]*models.ECCMUnlockEvent, 0)
	executeTxEvent, err := eccmContract.FilterVerifyHeaderAndExecuteTxEvent(opt)
	if err != nil {
		return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, filter unlock events :%s", err.Error())
	}

	for executeTxEvent.Next() {
		evt := executeTxEvent.Event
		Fee := e.GetConsumeGas(evt.Raw.TxHash)
		eccmUnlockEvent := verifyAndExecuteEvent2ProxyUnlockEvent(evt, Fee)
		eccmUnlockEvents = append(eccmUnlockEvents, eccmUnlockEvent)
	}
	return eccmLockEvents, eccmUnlockEvents, nil
}

func (e *EthereumChainListen) getNFTProxyEventByBlockNumber(
	proxyAddrStrs []string,
	startHeight, endHeight uint64) (
	[]*models.ProxyLockEvent,
	[]*models.ProxyUnlockEvent,
	error,
) {
	proxyLockEvents := make([]*models.ProxyLockEvent, 0)
	proxyUnlockEvents := make([]*models.ProxyUnlockEvent, 0)
	for _, proxyAddrStr := range proxyAddrStrs {
		if !isContract(proxyAddrStr) {
			continue
		}
		proxyAddr := common.HexToAddress(proxyAddrStr)
		proxyContract, err := nftlp.NewPolyNFTLockProxy(proxyAddr, e.ethSdk.GetClient())
		if err != nil {
			return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, error: %s", err.Error())
		}
		opt := &bind.FilterOpts{
			Start:   startHeight,
			End:     &endHeight,
			Context: context.Background(),
		}
		// get ethereum lock events from given block
		lockEvents, err := proxyContract.FilterLockEvent(opt)
		if err != nil {
			return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, filter lock events :%s", err.Error())
		}
		for lockEvents.Next() {
			proxyLockEvent := convertLockProxyEvent(lockEvents.Event)
			proxyLockEvents = append(proxyLockEvents, proxyLockEvent)
		}

		// ethereum unlock events from given block
		unlockEvents, err := proxyContract.FilterUnlockEvent(opt)
		if err != nil {
			return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, filter unlock events :%s", err.Error())
		}
		for unlockEvents.Next() {
			proxyUnlockEvent := convertUnlockProxyEvent(unlockEvents.Event)
			proxyUnlockEvents = append(proxyUnlockEvents, proxyUnlockEvent)
		}
	}
	return proxyLockEvents, proxyUnlockEvents, nil
}
