package ethereumlisten

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"net/http"
	"poly-swap/chainlisten/ethereumlisten/eccm_abi"
	"poly-swap/chainlisten/ethereumlisten/lock_proxy_abi"
	"poly-swap/chainlisten/ethereumlisten/wrapper_abi"
	"poly-swap/conf"
	"poly-swap/models"
	"poly-swap/utils"
	"strings"
)
const (
	_eth_crosschainlock   = "CrossChainLockEvent"
	_eth_crosschainunlock = "CrossChainUnlockEvent"
	_eth_lock = "LockEvent"
	_eth_unlock = "UnlockEvent"
)


type EthereumChainListen struct {
	ethCfg *conf.EthereumChainListenConfig
	ethSdk *EthereumSdk
}

func NewEthereumChainListen(cfg *conf.EthereumChainListenConfig) *EthereumChainListen {
	ethListen := &EthereumChainListen{}
	ethListen.ethCfg = cfg
	//
	sdk, err := NewEthereumSdk(cfg.RestURL)
	if err != nil {
		panic(err)
	}
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
			fctx := &models.SrcTransaction{}
			fctx.ChainId = this.GetChainId()
			fctx.Hash = lockEvent.TxHash
			fctx.State = 1
			fctx.Fee = lockEvent.Fee
			fctx.Time = tt
			fctx.Height = height
			fctx.User = utils.Hash2Address(this.GetChainId(), lockEvent.User)
			fctx.DstChainId = uint64(lockEvent.Tchain)
			fctx.Contract = lockEvent.Contract
			fctx.Key = lockEvent.Txid
			fctx.Param = hex.EncodeToString(lockEvent.Value)
			for _, v := range proxyLockEvents {
				if v.TxHash == lockEvent.TxHash {
					toAssetHash := v.ToAssetHash
					fctransfer := &models.SrcTransfer{}
					fctransfer.Hash = lockEvent.TxHash
					fctransfer.From = utils.Hash2Address(this.GetChainId(), v.FromAddress)
					fctransfer.To = utils.Hash2Address(this.GetChainId(), lockEvent.Contract)
					fctransfer.Asset = strings.ToLower(v.FromAssetHash)
					fctransfer.Amount = v.Amount.Uint64()
					fctransfer.DstChainId = uint64(v.ToChainId)
					fctransfer.DstAsset = toAssetHash
					fctransfer.DstUser = utils.Hash2Address(uint64(v.ToChainId), v.ToAddress)
					fctx.SrcTransfer = fctransfer
					break
				}
			}
			srcTransactions = append(srcTransactions, fctx)
		}
	}
	// save unLockEvent to db
	for _, unLockEvent := range eccmUnLockEvents {
		if unLockEvent.Method == _eth_crosschainunlock {
			logs.Info("to chain: txhash: %s\n", unLockEvent.TxHash)
			tctx := &models.DstTransaction{}
			tctx.ChainId = this.GetChainId()
			tctx.Hash = unLockEvent.TxHash
			tctx.State = 1
			tctx.Fee = unLockEvent.Fee
			tctx.Time = uint64(tt)
			tctx.Height = height
			tctx.SrcChainId = uint64(unLockEvent.FChainId)
			tctx.Contract = unLockEvent.Contract
			tctx.PolyHash = unLockEvent.RTxHash
			for _, v := range proxyUnlockEvents {
				if v.TxHash == unLockEvent.TxHash {
					tctransfer := &models.DstTransfer{}
					tctransfer.Hash = unLockEvent.TxHash
					tctransfer.From = utils.Hash2Address(this.GetChainId(), unLockEvent.Contract)
					tctransfer.To = utils.Hash2Address(this.GetChainId(), v.ToAddress)
					tctransfer.Asset = strings.ToLower(v.ToAssetHash)
					tctransfer.Amount = v.Amount.Uint64()
					tctx.DstTransfer = tctransfer
					break
				}
			}
			dstTransactions = append(dstTransactions, tctx)
		}
	}
	return wrapperTransactions, srcTransactions, nil, dstTransactions, nil
}

func (this *EthereumChainListen) getWapperEventByBlockNumber(contractAddr string, height uint64) ([]*models.WrapperTransaction, error) {
	wrapperAddress := common.HexToAddress(contractAddr)
	wrapperContract, err := polywrapper.NewIPolyWrapper(wrapperAddress, this.ethSdk.rawClient)
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
			User:         evt.Sender.String(),
			DstChainId:   uint64(evt.ToChainId),
			FeeTokenHash: evt.FromAsset.String(),
			FeeAmount:    evt.Fee.Uint64(),
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
			FeeAmount:    evt.Efee.Uint64(),
		})
	}
	return wrapperTransactions, nil
}

func (this *EthereumChainListen) getECCMEventByBlockNumber(contractAddr string, height uint64) ([]*models.ECCMLockEvent, []*models.ECCMUnlockEvent, error) {
	eccmContractAddress := common.HexToAddress(contractAddr)
	eccmContract, err := eccm_abi.NewEthCrossChainManager(eccmContractAddress, this.ethSdk.rawClient)
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
			Fee: Fee,
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
			Method:    _eth_crosschainunlock,
			TxHash:    evt.Raw.TxHash.String()[2:],
			RTxHash:   utils.HexStringReverse(hex.EncodeToString(evt.CrossChainTxHash)),
			Contract:  hex.EncodeToString(evt.ToContract),
			FChainId:  uint32(evt.FromChainID),
			Height:    height,
			Fee: Fee,
		})
	}
	return eccmLockEvents, eccmUnlockEvents, nil
}

func (this *EthereumChainListen) getProxyEventByBlockNumber(contractAddr string, height uint64) ([]*models.ProxyLockEvent, []*models.ProxyUnlockEvent, error) {
	proxyAddress := common.HexToAddress(contractAddr)
	proxyContract, err := lock_proxy_abi.NewLockProxy(proxyAddress, this.ethSdk.rawClient)
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
			Method:   _eth_lock,
			TxHash:         evt.Raw.TxHash.String()[2:],
			FromAddress:     evt.FromAddress.String()[2:],
			FromAssetHash:   evt.FromAssetHash.String()[2:],
			ToChainId:     uint32(evt.ToChainId),
			ToAssetHash:   hex.EncodeToString(evt.ToAssetHash),
			ToAddress:  hex.EncodeToString(evt.ToAddress),
			Amount:    evt.Amount,
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
			Method:    _eth_unlock,
			TxHash:    evt.Raw.TxHash.String()[2:],
			ToAssetHash:    evt.ToAssetHash.String()[2:],
			ToAddress:   evt.ToAddress.String()[2:],
			Amount:  evt.Amount,
		})
	}
	return proxyLockEvents, proxyUnlockEvents, nil
}
func (this *EthereumChainListen) GetConsumeGas(hash common.Hash) uint64 {
	tx, err := this.ethSdk.GetTransactionByHash(hash)
	if err != nil {
		return 0
	}
	receipt, err :=  this.ethSdk.GetTransactionReceipt(hash)
	if err != nil {
		return 0
	}
	return tx.GasPrice().Uint64() * receipt.GasUsed
}

type ExtendHeight struct {
	last_block_height  uint64  `json:"last_block_height,string"`
}

func (this *EthereumChainListen) GetExtendLatestHeight() (uint64, error) {
	req, err := http.NewRequest("GET", this.ethCfg.ExtendNodeURL, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Accepts", "application/json")
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
	extendHeight := new(ExtendHeight)
	err = json.Unmarshal(respBody, extendHeight)
	if err != nil {
		return 0, err
	}
	return extendHeight.last_block_height, nil
}



