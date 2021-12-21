package zionmainlisten

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"poly-bridge/go_abi/main_chain_lock_proxy_abi"
	"poly-bridge/go_abi/wrapper_abi"
	"poly-bridge/models"
	"strings"
)

type ZionMainListen struct {
	zionmainCfg *conf.ChainListenConfig
	zionmainSdk *chainsdk.EthereumSdkPro
}

func NewZionMainChainListen(cfg *conf.ChainListenConfig) *ZionMainListen {
	zionMainListen := &ZionMainListen{}
	zionMainListen.zionmainCfg = cfg
	urls := cfg.GetNodesUrl()
	sdk := chainsdk.NewEthereumSdkPro(urls, cfg.ListenSlot, cfg.ChainId)
	zionMainListen.zionmainSdk = sdk
	return zionMainListen
}

func (this *ZionMainListen) GetLatestHeight() (uint64, error) {
	return this.zionmainSdk.GetLatestHeight()
}

func (this *ZionMainListen) GetChainListenSlot() uint64 {
	return this.zionmainCfg.ListenSlot
}

func (this *ZionMainListen) GetChainId() uint64 {
	return this.zionmainCfg.ChainId
}

func (this *ZionMainListen) GetChainName() string {
	return this.zionmainCfg.ChainName
}

func (this *ZionMainListen) GetDefer() uint64 {
	return this.zionmainCfg.Defer
}

func (this *ZionMainListen) GetBatchSize() uint64 {
	return this.zionmainCfg.BatchSize
}

func (this *ZionMainListen) getCCMandLockEventByBlockNumber(startHeight,endHeight uint64,blockTT map[uint64]uint64) ([]*models.SrcTransaction, error) {
	eccmContractAddress := common.HexToAddress(this.zionmainCfg.CCMContract)
	lockContractAddress := common.HexToAddress(this.zionmainCfg.ProxyContract)
	client := this.zionmainSdk.GetClient()
	if client == nil {
		return nil, fmt.Errorf("ZionMain getECCMEventByBlockNumber GetClient error: nil")
	}
	eccmContract, err := main_chain_lock_proxy_abi.NewIMainChainLockProxy(eccmContractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("ZionMain NewIMainChainLockProxy eccmContract error: %s",err.Error())
	}
	lockContract, err := main_chain_lock_proxy_abi.NewIMainChainLockProxy(lockContractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("ZionMain NewIMainChainLockProxy lockContract error: %s",err.Error())
	}
	opt := &bind.FilterOpts{
		Start:   startHeight,
		End:     &endHeight,
		Context: context.Background(),
	}
	srcTransactions := make([]*models.SrcTransaction, 0)
	srcTransfers := make([]*models.SrcTransfer, 0)
	crossChainEvents, err := eccmContract.FilterCrossChainEvent(opt)
	if err != nil {
		return nil,fmt.Errorf("ZionMainListen FilterCrossChainEvent err: %s", err.Error())
	}
	lockEvents, err := lockContract.FilterLockEvent(opt)
	if err != nil {
		return nil,fmt.Errorf("ZionMainListen FilterLockEvent err: %s", err.Error())
	}
	for crossChainEvents.Next() {
		evt := crossChainEvents.Event
		fee := this.GetConsumeGas(evt.Raw.TxHash)
		srcTransactions = append(srcTransactions, &models.SrcTransaction{
			Hash:       evt.Raw.TxHash.String()[2:],
			ChainId:    this.zionmainCfg.ChainId,
			Standard:   models.TokenTypeErc20,
			State:      1,
			Fee:        models.NewBigIntFromInt(int64(fee)),
			Height:     evt.Raw.BlockNumber,
			User:       strings.ToLower(evt.Sender.String()[2:]),
			DstChainId: evt.ToChainId,
			Contract:   strings.ToLower(evt.ProxyOrAssetContract.String()[2:]),
			Key:        hex.EncodeToString(evt.TxId),
			Param:      hex.EncodeToString(evt.Rawdata),
		})
	}
	for lockEvents.Next(){
		evt := lockEvents.Event
		srcTransfers = append(srcTransfers, &models.SrcTransfer{
			TxHash:     evt.Raw.TxHash.String()[2:],
			ChainId:    this.zionmainCfg.ChainId,
			Standard:   models.TokenTypeErc20,
			Asset:      strings.ToLower(evt.FromAssetHash.String()[2:]),
			From:		strings.ToLower(evt.FromAddress.String()[2:]),
			To: 		strings.ToLower(evt.Raw.Address.String()[2:]),
			Amount:  	models.NewBigInt(evt.Amount),
			DstChainId: evt.ToChainId,
			DstAsset: strings.ToLower(hex.EncodeToString(evt.ToAssetHash)),
			DstUser: strings.ToLower(hex.EncodeToString(evt.ToAddress)),
		})
	}
	for _,srcTransaction:=range srcTransactions{
		for _,srcTransfer :=range srcTransfers{
			if srcTransaction.Hash==srcTransfer.TxHash{
				tt:=blockTT[srcTransaction.Height]
				srcTransfer.Time=tt
				srcTransaction.Time=tt
				srcTransaction.SrcTransfer=srcTransfer
				break
			}
		}
	}
	return srcTransactions,nil
}

func (this *ZionMainListen) getCCMandUnLockEventByBlockNumber(startHeight,endHeight uint64,blockTT map[uint64]uint64) ([]*models.DstTransaction, error) {
	eccmContractAddress := common.HexToAddress(this.zionmainCfg.CCMContract)
	lockContractAddress := common.HexToAddress(this.zionmainCfg.ProxyContract)
	client := this.zionmainSdk.GetClient()
	if client == nil {
		return nil, fmt.Errorf("ZionMain getCCMandUnLockEventByBlockNumber GetClient error: nil")
	}
	eccmContract, err := main_chain_lock_proxy_abi.NewIMainChainLockProxy(eccmContractAddress, client)
	lockContract, err := main_chain_lock_proxy_abi.NewIMainChainLockProxy(lockContractAddress, client)
	opt := &bind.FilterOpts{
		Start:   startHeight,
		End:     &endHeight,
		Context: context.Background(),
	}
	dstTransactions := make([]*models.DstTransaction, 0)
	dstTransfers := make([]*models.DstTransfer, 0)
	crossChainEvents, err := eccmContract.FilterVerifyHeaderAndExecuteTxEvent(opt)
	if err != nil {
		return nil,fmt.Errorf("ZionMainListen FilterVerifyHeaderAndExecuteTxEvent err: %s", err.Error())
	}
	unLockEvents, err := lockContract.FilterUnlockEvent(opt)
	if err != nil {
		return nil,fmt.Errorf("ZionMainListen FilterUnlockEvent err: %s", err.Error())
	}
	for crossChainEvents.Next() {
		evt := crossChainEvents.Event
		fee := this.GetConsumeGas(evt.Raw.TxHash)
		dstTransactions = append(dstTransactions, &models.DstTransaction{
			Hash:       evt.Raw.TxHash.String()[2:],
			ChainId:    this.zionmainCfg.ChainId,
			Standard:   models.TokenTypeErc20,
			State:      1,
			Fee:        models.NewBigIntFromInt(int64(fee)),
			Height:     evt.Raw.BlockNumber,
			SrcChainId: evt.FromChainID,
			Contract:   strings.ToLower(hex.EncodeToString(evt.ToContract)),
			PolyHash: 	strings.ToLower(hex.EncodeToString(evt.CrossChainTxHash)),
		})
	}
	for unLockEvents.Next(){
		evt := unLockEvents.Event
		dstTransfers = append(dstTransfers, &models.DstTransfer{
			TxHash:     evt.Raw.TxHash.String()[2:],
			ChainId:    this.zionmainCfg.ChainId,
			Standard:   models.TokenTypeErc20,
			Asset:      strings.ToLower(evt.ToAssetHash.String()[2:]),
			From:		strings.ToLower(evt.Raw.Address.String()[2:]),
			To: 		strings.ToLower(evt.ToAddress.String()[2:]),
			Amount:  	models.NewBigInt(evt.Amount),
		})
	}
	for _,dstTransaction:=range dstTransactions{
		for _,dstTransfer :=range dstTransfers{
			if dstTransaction.Hash==dstTransfer.TxHash{
				tt:=blockTT[dstTransaction.Height]
				dstTransfer.Time=tt
				dstTransaction.Time=tt
				dstTransaction.DstTransfer=dstTransfer
				break
			}
		}
	}
	return dstTransactions,nil
}

func (this *ZionMainListen) getWrapperEventByBlockNumber(startHeight,endHeight uint64,blockTT map[uint64]uint64) ([]*models.WrapperTransaction, error) {
	wrapperTransactions := make([]*models.WrapperTransaction, 0)
	for index, contract := range this.zionmainCfg.WrapperContract {
		if contract == "" {
			continue
		}
		aaa, err := this.getWrapperEventByBlockNumber1(contract, startHeight, endHeight, index)
		if err != nil {
			return nil, err
		}
		wrapperTransactions = append(wrapperTransactions, aaa...)
	}
	for _,wrapperTransaction :=range wrapperTransactions{
		wrapperTransaction.Time=blockTT[wrapperTransaction.BlockHeight]
	}
	return wrapperTransactions, nil
}

func (this *ZionMainListen) getWrapperEventByBlockNumber1(contract string,startHeight,endHeight uint64,index int) ([]*models.WrapperTransaction, error) {
	if len(contract) == 0 {
		return nil, nil
	}
	wrapperAddress := common.HexToAddress(contract)
	client := this.zionmainSdk.GetClient()
	if client == nil {
		return nil, fmt.Errorf("ZionMain getWrapperEventByBlockNumber1 GetClient error: nil")
	}
	wrapperContract, err := wrapper_abi.NewPolyWrapper(wrapperAddress, client)
	if err != nil {
		return nil, fmt.Errorf("ZionMain GetSmartContractEventByBlock(wrapper_ab_NewPolyWrapper), error: %s", err.Error())
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
		return nil, fmt.Errorf("getWrapperEventByBlockNumber1, filter lock events :%s", err.Error())
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
	if index == 1 {
		for _, tx := range wrapperTransactions {
			tx.FeeTokenHash = "0000000000000000000000000000000000000000"
		}
	}
	return wrapperTransactions, nil
}




func (this *ZionMainListen) GetConsumeGas(hash common.Hash) uint64 {
	tx, err := this.zionmainSdk.GetTransactionByHash(hash)
	if err != nil {
		return 0
	}
	receipt, err := this.zionmainSdk.GetTransactionReceipt(hash)
	if err != nil {
		return 0
	}
	return tx.GasPrice().Uint64() * receipt.GasUsed
}

func (this *ZionMainListen) HandleNewBlock(height uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, int, int, error) {
	startHeight := height
	if startHeight > 2 {
		startHeight -= 2
	}
	endHeight := height

	blockTT := make(map[uint64]uint64)
	for h := startHeight; h <= endHeight; h++ {
		blockHeader, err := this.zionmainSdk.GetHeaderByNumber(h)
		if err != nil {
			return nil, nil, nil, nil, 0, 0, err
		}
		if blockHeader == nil {
			return nil, nil, nil, nil, 0, 0, fmt.Errorf("there is no zionmain block on height: %d!", h)
		}
		blockTT[h] = blockHeader.Time
	}
	srcTransactions, err := this.getCCMandLockEventByBlockNumber(startHeight,endHeight,blockTT)
	if err != nil {
		return nil, nil, nil, nil, 0, 0, err
	}
	dstTransactions,err:=this.getCCMandUnLockEventByBlockNumber(startHeight,endHeight,blockTT)
	if err != nil {
		return nil, nil, nil, nil, 0, 0, err
	}
	wrapperTransactions,err:= this.getWrapperEventByBlockNumber(startHeight,endHeight,blockTT)
	if err != nil {
		return nil, nil, nil, nil, 0, 0, err
	}
	return wrapperTransactions, srcTransactions, nil, dstTransactions, len(srcTransactions), len(dstTransactions), nil
}

func (this *ZionMainListen) GetExtendLatestHeight() (uint64, error) {
	if len(this.zionmainCfg.ExtendNodes) == 0 {
		return this.GetLatestHeight()
	}
	return this.GetLatestHeight()
}
