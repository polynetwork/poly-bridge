package aptoslisten

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"poly-bridge/crosschaindao"
	"poly-bridge/models"
	"strconv"
)

type AptosChainListen struct {
	aptosCfg *conf.ChainListenConfig
	aptosSdk *chainsdk.AptosSdkPro
}

func NewAptosChainListen(cfg *conf.ChainListenConfig) *AptosChainListen {
	aptosListen := &AptosChainListen{}
	aptosListen.aptosCfg = cfg
	sdk := chainsdk.NewAptosSdkPro(cfg.Nodes, cfg.ListenSlot, cfg.ChainId)
	aptosListen.aptosSdk = sdk
	return aptosListen
}

func (a *AptosChainListen) GetExtendLatestHeight() (uint64, error) {
	return a.GetLatestHeight()
}

func (a *AptosChainListen) GetLatestHeight() (uint64, error) {
	return a.aptosSdk.GetBlockCount()
}

func (a *AptosChainListen) GetChainListenSlot() uint64 {
	return a.aptosCfg.ListenSlot
}

func (a *AptosChainListen) GetChainId() uint64 {
	return a.aptosCfg.ChainId
}

func (a *AptosChainListen) GetChainName() string {
	return a.aptosCfg.ChainName
}

func (a *AptosChainListen) GetDefer() uint64 {
	return a.aptosCfg.Defer
}

func (a *AptosChainListen) GetBatchSize() uint64 {
	return a.aptosCfg.BatchSize
}

func (a *AptosChainListen) GetBatchLength() (uint64, uint64) {
	return a.aptosCfg.MinBatchLength, a.aptosCfg.MaxBatchLength
}

func (a *AptosChainListen) HandleNewBatchBlock(start, end uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, []*models.WrapperDetail, []*models.PolyDetail, int, int, error) {
	return nil, nil, nil, nil, nil, nil, 0, 0, nil
}

func (a *AptosChainListen) HandleNewBlock(height uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, []*models.WrapperDetail, []*models.PolyDetail, int, int, error) {

	return nil, nil, nil, nil, nil, nil, 0, 0, nil
}

func (a *AptosChainListen) HandleEvent(db crosschaindao.CrossChainDao, crossChainSequenceNumber, executeTxSequenceNumber, limit uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, []*models.WrapperDetail, []*models.PolyDetail, error) {
	var dbChain *models.Chain
	if db != nil {
		chain, err := db.GetChain(a.GetChainId())
		if err != nil {
			return nil, nil, nil, nil, nil, nil, fmt.Errorf("get chain %d err: %v", a.GetChainId(), err)
		}
		crossChainSequenceNumber = chain.CrossChainSequenceNumber
		executeTxSequenceNumber = chain.ExecuteTxSequenceNumber
		dbChain = chain

		height, err := a.aptosSdk.GetBlockCount()
		if err != nil {
			return nil, nil, nil, nil, nil, nil, fmt.Errorf("get aptos latest err: %v", a.GetChainId(), err)
		}
		dbChain.Height = height
	}

	if limit == 0 {
		limit = a.aptosCfg.BatchSize
	}

	wrapperTransactions := make([]*models.WrapperTransaction, 0)
	srcTransactions := make([]*models.SrcTransaction, 0)
	dstTransactions := make([]*models.DstTransaction, 0)

	crossChainEventFilter := &chainsdk.AptosEventFilter{Address: a.aptosCfg.CCMContract, CreationNumber: a.aptosCfg.CrossChainEventCreationNumber, Query: make(map[string]interface{})}
	crossChainEventFilter.Query["limit"] = limit
	crossChainEventFilter.Query["start"] = crossChainSequenceNumber

	executeTxEventFilter := &chainsdk.AptosEventFilter{Address: a.aptosCfg.CCMContract, CreationNumber: a.aptosCfg.ExecuteTxEventCreationNumber, Query: make(map[string]interface{})}
	executeTxEventFilter.Query["limit"] = limit
	executeTxEventFilter.Query["start"] = executeTxSequenceNumber
	crossChainEvents, err := a.aptosSdk.GetEvents(crossChainEventFilter)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("aptos get crossChainEvents failed. filter: %+v, err: %v", *crossChainEventFilter, err)
	}
	if len(crossChainEvents) != 0 {
		logs.Info("aptos crossChainEvents=%+v", crossChainEvents)
	}

	executeTxEvents, err := a.aptosSdk.GetEvents(executeTxEventFilter)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("aptos get executeTxEvents failed. filter: %+v, err: %v", *crossChainEventFilter, err)
	}
	if len(executeTxEvents) != 0 {
		logs.Info("aptos executeTxEvents=%+v", executeTxEvents)
	}

	var nextCrossChainSequenceNumber uint64
	for _, event := range crossChainEvents {
		block, err := a.aptosSdk.GetBlockByVersion(uint64(event.Version))
		if err != nil {
			return nil, nil, nil, nil, nil, nil, fmt.Errorf("GetBlockByVersion failed. version:%s, err: %v", event.Version, err)
		}
		tx, err := a.aptosSdk.GetTxByVersion(uint64(event.Version))
		if err != nil {
			return nil, nil, nil, nil, nil, nil, fmt.Errorf("GetTxByVersion failed. version:%s, err: %v", event.Version, err)
		}
		txTime, err := strconv.ParseUint(tx.Timestamp[:10], 0, 32)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, fmt.Errorf("parse tx Timestamp failed. version:%s, err: %v", event.Version, err)
		}

		// source transaction
		srcTx := &models.SrcTransaction{}
		srcTx.ChainId = a.GetChainId()
		srcTx.DstChainId, _ = strconv.ParseUint(event.Data["to_chain_id"].(string), 0, 32)

		srcTx.Hash = tx.Hash[2:]
		srcTx.State = 1
		gasUsed, _ := strconv.ParseInt(tx.GasUsed, 0, 32)
		srcTx.Fee = models.NewBigIntFromInt(gasUsed)
		srcTx.Time = txTime

		srcTx.Height, _ = strconv.ParseUint(block.BlockHeight, 0, 32)
		srcTx.User = tx.Sender
		srcTx.Contract = event.GUID.AccountAddress[2:]
		srcTx.Key = event.Data["tx_id"].(string)[2:]
		srcTx.Param = event.Data["raw_data"].(string)

		// source transfer
		if lockEvent := a.aptosSdk.GetLatest().Sdk.GetLockEvent(tx.Events); lockEvent != nil {
			srcTransfer := &models.SrcTransfer{}
			srcTransfer.Time = txTime
			srcTransfer.ChainId = a.GetChainId()
			srcTransfer.DstChainId, _ = strconv.ParseUint(lockEvent.Data["to_chain_id"].(string), 0, 32)
			srcTransfer.TxHash = tx.Hash[2:]
			srcTransfer.From = tx.Sender
			srcTransfer.To = event.GUID.AccountAddress
			srcTransfer.Asset = tx.Payload.TypeArguments[0]
			amount, _ := strconv.ParseInt(tx.Payload.Arguments[0].(string), 0, 32)
			srcTransfer.Amount = models.NewBigIntFromInt(amount)

			srcTransfer.DstAsset = lockEvent.Data["to_asset_hash"].(string)
			if basedef.Has0xPrefix(srcTransfer.DstAsset) {
				srcTransfer.DstAsset = srcTransfer.DstAsset[2:]
			}
			srcTransfer.DstAsset = models.FormatAssert(srcTransfer.DstAsset)

			srcTransfer.DstUser = lockEvent.Data["to_address"].(string)
			if basedef.Has0xPrefix(srcTransfer.DstUser) {
				srcTransfer.DstUser = srcTransfer.DstUser[2:]
			}
			srcTransfer.DstUser = models.FormatString(srcTransfer.DstUser)

			srcTx.SrcTransfer = srcTransfer
		}
		srcTransactions = append(srcTransactions, srcTx)

		// wrapper transaction
		wrapperTx := &models.WrapperTransaction{}
		if lockWithFeeEvent := a.aptosSdk.GetLatest().Sdk.GetLockWithFeeEvent(tx.Events); lockWithFeeEvent != nil {
			wrapperTx.Hash = tx.Hash[2:]
			wrapperTx.SrcChainId = a.GetChainId()
			wrapperTx.BlockHeight = srcTx.Height
			wrapperTx.Time = txTime
			wrapperTx.DstChainId, _ = strconv.ParseUint(lockWithFeeEvent.Data["to_chain_id"].(string), 0, 32)
			wrapperTx.DstUser = models.FormatString(lockWithFeeEvent.Data["to_address"].(string))
			wrapperTx.FeeTokenHash = "0x1::aptos_coin::AptosCoin"
			feeAmount, _ := strconv.ParseInt(lockWithFeeEvent.Data["fee_amount"].(string), 0, 32)
			wrapperTx.FeeAmount = models.NewBigIntFromInt(feeAmount)
			wrapperTx.Status = basedef.STATE_SOURCE_DONE
			wrapperTransactions = append(wrapperTransactions, wrapperTx)
		}
		nextCrossChainSequenceNumber = uint64(event.SequenceNumber)
	}

	var nextExecuteTxSequenceNumber uint64
	for _, event := range executeTxEvents {
		block, err := a.aptosSdk.GetBlockByVersion(uint64(event.Version))
		if err != nil {
			return nil, nil, nil, nil, nil, nil, fmt.Errorf("GetBlockByVersion failed. version:%s, err: %v", event.Version, err)
		}
		tx, err := a.aptosSdk.GetTxByVersion(uint64(event.Version))
		if err != nil {
			return nil, nil, nil, nil, nil, nil, fmt.Errorf("GetTxByVersion failed. version:%s, err: %v", event.Version, err)
		}
		txTime, err := strconv.ParseUint(tx.Timestamp[:10], 0, 32)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, fmt.Errorf("parse tx Timestamp failed. version:%s, err: %v", event.Version, err)
		}

		// dst transaction
		dstTx := &models.DstTransaction{}
		//dstTx.DstTransfer = dstTransfers
		dstTx.ChainId = a.GetChainId()
		dstTx.Hash = tx.Hash[2:]
		dstTx.State = 1
		gasUsed, _ := strconv.ParseInt(tx.GasUsed, 0, 32)
		dstTx.Fee = models.NewBigIntFromInt(gasUsed)
		dstTx.Time = txTime
		dstTx.Height, _ = strconv.ParseUint(block.BlockHeight, 0, 32)
		fmt.Println("aptos dst from_chain_id", event.Data["from_chain_id"])
		//dstTx.SrcChainId, _ = strconv.ParseUint(event.Data["from_chain_id"].(string), 0, 32)
		dstTx.Contract = event.GUID.AccountAddress
		dstTx.PolyHash = event.Data["cross_chain_tx_hash"].(string)[2:]

		// dst transfer
		if unLockEvent := a.aptosSdk.GetLatest().Sdk.GetUnLockEvent(tx.Events); unLockEvent != nil {
			dstTransfer := &models.DstTransfer{}
			dstTransfer.TxHash = tx.Hash[2:]
			dstTransfer.Time = txTime
			dstTransfer.ChainId = a.GetChainId()
			dstTransfer.From = event.GUID.AccountAddress
			dstTransfer.To = models.FormatString(unLockEvent.Data["to_address"].(string))
			dstTransfer.Asset = tx.Payload.TypeArguments[0]
			amount, _ := strconv.ParseInt(unLockEvent.Data["amount"].(string), 0, 32)
			dstTransfer.Amount = models.NewBigIntFromInt(amount)
			dstTx.DstTransfer = dstTransfer
		}
		nextExecuteTxSequenceNumber = uint64(event.SequenceNumber)
		dstTransactions = append(dstTransactions, dstTx)
	}

	// update chain
	if dbChain != nil {
		if len(srcTransactions) > 0 && nextCrossChainSequenceNumber >= dbChain.CrossChainSequenceNumber {
			dbChain.CrossChainSequenceNumber = nextCrossChainSequenceNumber + 1
		}
		if len(executeTxEvents) > 0 && nextExecuteTxSequenceNumber >= dbChain.ExecuteTxSequenceNumber {
			dbChain.ExecuteTxSequenceNumber = nextExecuteTxSequenceNumber + 1
		}
		err = db.UpdateChain(dbChain)
		if err != nil {
			logs.Error("Aptos listen update chain err: %v", err)
		}
	}

	return wrapperTransactions, srcTransactions, nil, dstTransactions, nil, nil, nil
}
