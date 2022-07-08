package starcoinlisten

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/serde"
	"github.com/starcoinorg/starcoin-go/client"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"poly-bridge/models"
	"strconv"
	"strings"
)

const (
	starcoin_ccm_lock_event_tag     = "::CrossChainManager::CrossChainEvent"
	starcoin_ccm_execute_Event_tag  = "::CrossChainManager::VerifyHeaderAndExecuteTxEvent"
	starcoin_proxy_lock_event_tag   = "::LockProxy::LockEvent"
	starcoin_proxy_unlock_event_tag = "::LockProxy::UnlockEvent"
	starcoin_proxy_wrap_event_tag   = "::LockProxy::CrossChainFeeLockEvent"
)

type AccountAddress [16]uint8

type StarcoinChainListen struct {
	starcoinCfg *conf.ChainListenConfig
	starcoinSdk *chainsdk.StarcoinSdkPro
}

func (s *StarcoinChainListen) GetExtendLatestHeight() (uint64, error) {
	return s.GetLatestHeight()
}

type StarcoinEvents struct {
	crossChainEvent        *client.Event
	lockEvent              *client.Event
	crossChainFeeLockEvent *client.Event
	executeTxEvent         *client.Event
	unlockEvent            *client.Event
}

type LockEvent struct {
	FromAssetHash TokenCode
	FromAddress   []byte
	ToChainId     uint64
	ToAssetHash   []byte
	ToAddress     []byte
	Amount        serde.Uint128
}

type CrossChainEvent struct {
	Sender               []byte
	TxId                 []byte
	ProxyOrAssetContract []byte
	ToChainId            uint64
	ToContract           []byte
	RawData              []byte
}

type CrossChainFeeLockEvent struct {
	FromAssetHash TokenCode
	Sender        AccountAddress
	ToChainId     uint64
	ToAddress     []byte
	Net           serde.Uint128
	Fee           serde.Uint128
	Id            serde.Uint128
}

type TokenCode struct {
	Address AccountAddress
	Module  string
	Name    string
}

type VerifyHeaderAndExecuteTxEvent struct {
	FromChainId      uint64
	ToContract       []byte
	CrossChainTxHash []byte
	FromChainTxHash  []byte
}

type UnlockEvent struct {
	ToAssetHash []byte
	ToAddress   []byte
	Amount      serde.Uint128
}

func NewStarcoinChainListen(cfg *conf.ChainListenConfig) *StarcoinChainListen {
	starcoinListen := &StarcoinChainListen{}
	starcoinListen.starcoinCfg = cfg
	urls := cfg.GetNodesUrl()
	sdk := chainsdk.NewStarcoinSdkPro(urls, cfg.ListenSlot, cfg.ChainId)
	starcoinListen.starcoinSdk = sdk
	return starcoinListen
}

func (s *StarcoinChainListen) GetLatestHeight() (uint64, error) {
	return s.starcoinSdk.GetBlockCount()
}

func (s *StarcoinChainListen) GetChainListenSlot() uint64 {
	return s.starcoinCfg.ListenSlot
}

func (s *StarcoinChainListen) GetChainId() uint64 {
	return s.starcoinCfg.ChainId
}

func (s *StarcoinChainListen) GetChainName() string {
	return s.starcoinCfg.ChainName
}

func (s *StarcoinChainListen) GetDefer() uint64 {
	return s.starcoinCfg.Defer
}

func (s *StarcoinChainListen) GetBatchSize() uint64 {
	return s.starcoinCfg.BatchSize
}

func (s *StarcoinChainListen) HandleNewBlock(height uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, []*models.WrapperDetail, []*models.PolyDetail, int, int, error) {
	block, err := s.starcoinSdk.GetBlockByIndex(height)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, 0, 0, err
	}
	if block == nil {
		return nil, nil, nil, nil, nil, nil, 0, 0, fmt.Errorf("can not get starcoin block, height=%d", height)
	}
	blockTime, err := strconv.Atoi(block.BlockHeader.Timestamp)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, 0, 0, fmt.Errorf("parse block time failed. time=%s, err=%s", block.BlockHeader.Timestamp, err)
	}
	blockTime = blockTime / 1000

	wrapperTransactions, srcTransactions, dstTransactions := s.getStarcoinTxs(height, blockTime)
	return wrapperTransactions, srcTransactions, nil, dstTransactions, nil, nil, len(srcTransactions), len(dstTransactions), nil
}

func (s *StarcoinChainListen) getStarcoinTxs(height uint64, blockTime int) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.DstTransaction) {
	wrapperTransactions := make([]*models.WrapperTransaction, 0)
	srcTransactions := make([]*models.SrcTransaction, 0)
	dstTransactions := make([]*models.DstTransaction, 0)
	for _, contract := range s.starcoinCfg.ProxyContract {
		typeTags := []string{
			contract + starcoin_ccm_lock_event_tag,
			contract + starcoin_proxy_lock_event_tag,
			contract + starcoin_proxy_wrap_event_tag,
			contract + starcoin_ccm_execute_Event_tag,
			contract + starcoin_proxy_unlock_event_tag,
		}
		eventFilter := &client.EventFilter{
			Address:   []string{contract},
			TypeTags:  typeTags,
			FromBlock: height,
			ToBlock:   &height,
		}
		starcoinEvents, err := s.starcoinSdk.GetEvents(eventFilter)
		if err != nil {
			logs.Error("starcoin height:%d fetch source events error=%s", height, err)
			continue
		} else if starcoinEvents == nil {
			logs.Info("starcoin height:%d no source events found", height)
		}

		starcoinEventsMap := make(map[string]*StarcoinEvents, 0)
		for i, _ := range starcoinEvents {
			event := starcoinEvents[i]
			srcHash := event.TransactionHash[2:]
			var starcoinEvents *StarcoinEvents
			if evts, ok := starcoinEventsMap[srcHash]; ok {
				starcoinEvents = evts
			} else {
				starcoinEvents = &StarcoinEvents{}
				starcoinEventsMap[srcHash] = starcoinEvents
			}

			if strings.Contains(event.TypeTag, starcoin_ccm_lock_event_tag) {
				starcoinEvents.crossChainEvent = &event
			} else if strings.Contains(event.TypeTag, starcoin_proxy_lock_event_tag) {
				starcoinEvents.lockEvent = &event
			} else if strings.Contains(event.TypeTag, starcoin_proxy_wrap_event_tag) {
				starcoinEvents.crossChainFeeLockEvent = &event
			} else if strings.Contains(event.TypeTag, starcoin_ccm_execute_Event_tag) {
				starcoinEvents.executeTxEvent = &event
			} else if strings.Contains(event.TypeTag, starcoin_proxy_unlock_event_tag) {
				starcoinEvents.unlockEvent = &event
			}
		}

		for hash, evts := range starcoinEventsMap {
			txInfo, err := s.starcoinSdk.GetTransactionInfoByHash(hash)
			var txFee *models.BigInt
			if err == nil && txInfo != nil {
				gasUsed, _ := strconv.ParseInt(txInfo.GasUsed, 0, 64)
				txFee = models.NewBigIntFromInt(gasUsed)
			}

			var ccEvent *CrossChainEvent
			var lockEvent *LockEvent
			var feeEvent *CrossChainFeeLockEvent
			var executeTxEvent *VerifyHeaderAndExecuteTxEvent
			var unlockEvent *UnlockEvent
			if evts.crossChainEvent != nil {
				// ccm lock event
				crossChainEventData, err := HexToBytes(evts.crossChainEvent.Data)
				if err != nil {
					logs.Error("starcoin crossChainEvent.Data HexToBytes err=%s", err)
					continue
				}
				ccEvent, err = BcsDeserializeCrossChainEvent(crossChainEventData)
				if err != nil {
					logs.Error("starcoin BcsDeserializeCrossChainEvent err=%s", err)
				}
			}
			if evts.lockEvent != nil {
				// lock event
				lockEventData, err := HexToBytes(evts.lockEvent.Data)
				if err != nil {
					logs.Error("starcoin lockEvent.Data HexToBytes err=%s", err)
					continue
				}
				lockEvent, err = BcsDeserializeLockEvent(lockEventData)
				if err != nil {
					logs.Error("starcoin BcsDeserializeLockEvent err=%s", err)
				}
			}
			if evts.crossChainFeeLockEvent != nil {
				// fee event
				feeEventData, err := HexToBytes(evts.crossChainFeeLockEvent.Data)
				if err != nil {
					logs.Error("starcoin crossChainFeeLockEvent.Data HexToBytes err=%s", err)
					continue
				}
				feeEvent, err = BcsDeserializeCrossChainFeeLockEvent(feeEventData)
				if err != nil {
					logs.Error("starcoin BcsDeserializeCrossChainFeeLockEvent err=%s", err)
				}
			}
			if evts.executeTxEvent != nil {
				// ccm execute event
				executeTxEventData, err := HexToBytes(evts.executeTxEvent.Data)
				if err != nil {
					logs.Error("starcoin executeTxEvent.Data HexToBytes err=%s", err)
					continue
				}
				executeTxEvent, err = BcsDeserializeVerifyHeaderAndExecuteTxEvent(executeTxEventData)
				if err != nil {
					logs.Error("starcoin BcsDeserializeVerifyHeaderAndExecuteTxEvent err=%s", err)
				}
			}
			if evts.unlockEvent != nil {
				// unlock event
				unlockEventData, err := HexToBytes(evts.unlockEvent.Data)
				if err != nil {
					logs.Error("starcoin unlockEvent.Data HexToBytes err=%s", err)
					continue
				}
				unlockEvent, err = BcsDeserializeUnlockEvent(unlockEventData)
				if err != nil {
					logs.Error("starcoin BcsDeserializeUnlockEvent err=%s", err)
				}
			}
			ccEventJson, _ := json.Marshal(ccEvent)
			lockEventJson, _ := json.Marshal(lockEvent)
			feeEventJson, _ := json.Marshal(feeEvent)
			executeTxEventJson, _ := json.Marshal(executeTxEvent)
			unlockEventJson, _ := json.Marshal(unlockEvent)
			logs.Info("starcoin height=%d, ccEvent=%s, lockEvent=%s, feeEvent=%s, executeTxEvent=%s, unlockEvent=%s",
				height, ccEventJson, lockEventJson, feeEventJson, executeTxEventJson, unlockEventJson)

			if ccEvent != nil && lockEvent != nil {
				// source transfer
				srcTransfer := &models.SrcTransfer{}
				srcTransfer.Time = uint64(blockTime)
				srcTransfer.ChainId = s.GetChainId()
				srcTransfer.DstChainId = ccEvent.ToChainId
				srcTransfer.TxHash = hash
				srcTransfer.From = models.FormatString(hex.EncodeToString(lockEvent.FromAddress))
				srcTransfer.To = models.FormatString(contract)
				srcTransfer.Asset = models.FormatString(GetTokenCodeString(&lockEvent.FromAssetHash))
				srcTransfer.Amount = models.NewBigInt(Uint128ToBigInt(&lockEvent.Amount))
				srcTransfer.DstAsset = models.FormatString(hex.EncodeToString(lockEvent.ToAssetHash))
				srcTransfer.DstUser = models.FormatString(hex.EncodeToString(lockEvent.ToAddress))

				// source transaction
				srcTx := &models.SrcTransaction{}
				srcTx.SrcTransfer = srcTransfer
				srcTx.ChainId = s.GetChainId()
				srcTx.DstChainId = ccEvent.ToChainId
				srcTx.Hash = hash
				srcTx.State = 1
				srcTx.Fee = txFee
				srcTx.Time = uint64(blockTime)
				srcTx.Height = height
				srcTx.User = models.FormatString(hex.EncodeToString(lockEvent.FromAddress))
				srcTx.Contract = models.FormatString(contract)
				srcTx.Key = hex.EncodeToString(ccEvent.TxId)
				srcTx.Param = hex.EncodeToString(ccEvent.RawData)
				srcTransactions = append(srcTransactions, srcTx)

				if feeEvent != nil {
					// wrapper transaction
					wrapperTx := &models.WrapperTransaction{}
					wrapperTx.Hash = hash
					wrapperTx.User = models.FormatString(hex.EncodeToString(lockEvent.FromAddress))
					wrapperTx.SrcChainId = s.GetChainId()
					wrapperTx.BlockHeight = height
					wrapperTx.Time = uint64(blockTime)
					wrapperTx.DstChainId = ccEvent.ToChainId
					wrapperTx.DstUser = models.FormatString(hex.EncodeToString(lockEvent.ToAddress))
					wrapperTx.FeeTokenHash = "0x00000000000000000000000000000001::STC::STC"
					wrapperTx.FeeAmount = models.NewBigInt(Uint128ToBigInt(&feeEvent.Fee))
					wrapperTx.Status = basedef.STATE_SOURCE_DONE
					wrapperTransactions = append(wrapperTransactions, wrapperTx)
				}
			}

			if unlockEvent != nil && executeTxEvent != nil {
				// dst transfer
				dstTransfer := &models.DstTransfer{}
				dstTransfer.TxHash = hash
				dstTransfer.Time = uint64(blockTime)
				dstTransfer.ChainId = s.GetChainId()
				dstTransfer.From = models.FormatString(contract)
				dstTransfer.To = models.FormatString(hex.EncodeToString(unlockEvent.ToAddress))
				dstTransfer.Asset = models.FormatString(string(unlockEvent.ToAssetHash))
				dstTransfer.Amount = models.NewBigInt(Uint128ToBigInt(&unlockEvent.Amount))

				// dst transaction
				dstTx := &models.DstTransaction{}
				dstTx.DstTransfer = dstTransfer
				dstTx.ChainId = s.GetChainId()
				dstTx.Hash = hash
				dstTx.State = 1
				dstTx.Fee = txFee
				dstTx.Time = uint64(blockTime)
				dstTx.Height = height
				dstTx.SrcChainId = executeTxEvent.FromChainId
				dstTx.Contract = models.FormatString(hex.EncodeToString(executeTxEvent.ToContract))
				dstTx.PolyHash = basedef.HexStringReverse(hex.EncodeToString(executeTxEvent.CrossChainTxHash))
				dstTransactions = append(dstTransactions, dstTx)
			}
		}
	}
	return wrapperTransactions, srcTransactions, dstTransactions
}
