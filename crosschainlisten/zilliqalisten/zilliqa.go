package zilliqalisten

import (
	"encoding/hex"
	"fmt"
	"github.com/Zilliqa/gozilliqa-sdk/bech32"
	"github.com/Zilliqa/gozilliqa-sdk/util"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"poly-bridge/models"
	"strconv"
	"strings"
)

type ZilliqaChainListen struct {
	zliCfg *conf.ChainListenConfig
	zliSdk *chainsdk.ZilliqaSdkPro
}

func NewZilliqaChainListen(cfg *conf.ChainListenConfig) *ZilliqaChainListen {
	zilListen := &ZilliqaChainListen{}
	zilListen.zliCfg = cfg
	urls := cfg.GetNodesUrl()
	sdk := chainsdk.NewZilliqaSdkPro(urls, cfg.ListenSlot, cfg.ChainId)
	zilListen.zliSdk = sdk
	return zilListen
}

func (this *ZilliqaChainListen) GetLatestHeight() (uint64, error) {
	return this.zliSdk.GetLatestHeight()
}

func (this *ZilliqaChainListen) GetChainListenSlot() uint64 {
	return this.zliCfg.ListenSlot
}

func (this *ZilliqaChainListen) GetChainId() uint64 {
	return this.zliCfg.ChainId
}

func (this *ZilliqaChainListen) GetChainName() string {
	return this.zliCfg.ChainName
}

func (this *ZilliqaChainListen) GetDefer() uint64 {
	return this.zliCfg.Defer
}

func (this *ZilliqaChainListen) GetBatchSize() uint64 {
	return this.zliCfg.BatchSize
}

func (this *ZilliqaChainListen) HandleNewBlock(height uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, int, int, error) {
	block, err := this.zliSdk.GetBlockByHeight(height)
	if err != nil {
		return nil, nil, nil, nil, 0, 0, err
	}
	if block == nil {
		return nil, nil, nil, nil, 0, 0, fmt.Errorf("there is no zilliqa block!")
	}
	tt := block.Timestamp
	srcTransactions := make([]*models.SrcTransaction, 0)
	dstTransactions := make([]*models.DstTransaction, 0)

	ccmLockEvent, lockEvents, err := this.getzilliqaCCMLockEventByBlockNumber(height, block)
	if err != nil {
		return nil, nil, nil, nil, 0, 0, err
	}

	ccmUnlockEvent, unlockEvents, err := this.getCosmosCCMUnlockEventByBlockNumber(height)
	if err != nil {
		return nil, nil, nil, nil, 0, 0, err
	}

	for _, lockEvent := range ccmLockEvent {
		if lockEvent.Method == _switcheo_crosschainlock {
			logs.Info("from chain: %d, txhash: %s\n", this.GetChainName(), lockEvent.TxHash)
			srcTransfer := &models.SrcTransfer{}
			for _, v := range lockEvents {
				if v.TxHash == lockEvent.TxHash {
					srcTransfer.ChainId = this.GetChainId()
					srcTransfer.TxHash = lockEvent.TxHash
					srcTransfer.Time = tt
					srcTransfer.From, _ = basedef.Address2Hash(srcTransfer.ChainId, v.FromAddress)
					srcTransfer.To = v.ToAddress
					srcTransfer.Asset = v.FromAssetHash
					srcTransfer.Amount = models.NewBigInt(v.Amount)
					srcTransfer.DstChainId = uint64(v.ToChainId)
					srcTransfer.DstAsset = v.ToAssetHash
					srcTransfer.DstUser = v.DstUser
					break
				}
			}
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
			srcTransaction.Key = lockEvent.TxHash
			srcTransaction.Param = hex.EncodeToString(lockEvent.Value)
			srcTransaction.SrcTransfer = srcTransfer
			srcTransactions = append(srcTransactions, srcTransaction)
		}
	}
	for _, unLockEvent := range ccmUnlockEvent {
		if unLockEvent.Method == _switcheo_crosschainunlock {
			logs.Info("to chain: %s, txhash: %s\n", this.GetChainName(), unLockEvent.TxHash)
			dstTransfer := &models.DstTransfer{}
			for _, v := range unlockEvents {
				if v.TxHash == unLockEvent.TxHash {
					dstTransfer.ChainId = this.GetChainId()
					dstTransfer.TxHash = unLockEvent.TxHash
					dstTransfer.Time = tt
					dstTransfer.From, _ = basedef.Address2Hash(dstTransfer.ChainId, unLockEvent.Contract)
					dstTransfer.To, _ = basedef.Address2Hash(dstTransfer.ChainId, v.ToAddress)
					dstTransfer.Asset = v.ToAssetHash
					dstTransfer.Amount = models.NewBigInt(v.Amount)
					break
				}
			}
			dstTransaction := &models.DstTransaction{}
			dstTransaction.ChainId = this.GetChainId()
			dstTransaction.Hash = unLockEvent.TxHash
			dstTransaction.State = 1
			dstTransaction.Fee = models.NewBigIntFromInt(int64(unLockEvent.Fee))
			dstTransaction.Time = tt
			dstTransaction.Height = height
			dstTransaction.SrcChainId = uint64(unLockEvent.FChainId)
			dstTransaction.Contract = unLockEvent.Contract
			dstTransaction.PolyHash = unLockEvent.RTxHash
			dstTransaction.DstTransfer = dstTransfer
			dstTransactions = append(dstTransactions, dstTransaction)
		}
	}
	return nil, srcTransactions, nil, dstTransactions, len(srcTransactions), len(dstTransactions), nil
}

func (this *ZilliqaChainListen) getzilliqaCCMLockEventByBlockNumber(height uint64, block *chainsdk.ZilBlock) ([]*models.ECCMLockEvent, []*models.ProxyLockEvent, error) {
	client := this.zliSdk
	ccmLockEvents := make([]*models.ECCMLockEvent, 0)
	lockEvents := make([]*models.ProxyLockEvent, 0)
	for _, transaction := range block.Transactions {
		if !transaction.Receipt.Success {
			continue
		}
		events := transaction.Receipt.EventLogs
		for _, event := range events {
			// 1. contract address should be cross chain manager
			// 2. event name should be CrossChainEvent
			toAddr, _ := bech32.ToBech32Address(event.Address)
			if toAddr == s.crossChainManagerAddress {
				if event.EventName != "CrossChainEvent" {
					continue
				}
				log.Infof("ZilliqaSyncManager found event on cross chain manager: %+v\n", event)
				// todo parse event to struct CrossTransfer
				crossTx := &CrossTransfer{}
				for _, param := range event.Params {
					switch param.VName {
					case "txId":
						index := big.NewInt(0)
						index.SetBytes(util.DecodeHex(param.Value.(string)))
						crossTx.txIndex = tools.EncodeBigInt(index)
					case "toChainId":
						toChainId, _ := strconv.ParseUint(param.Value.(string), 10, 32)
						crossTx.toChain = uint32(toChainId)
					case "rawdata":
						crossTx.value = util.DecodeHex(param.Value.(string))
					}
				}
				crossTx.height = height
				crossTx.txId = util.DecodeHex(transaction.ID)
				log.Infof("ZilliqaSyncManager parsed cross tx is: %+v\n", crossTx)
				sink := common.NewZeroCopySink(nil)
				crossTx.Serialization(sink)
				err1 := s.db.PutRetry(sink.Bytes())
				if err1 != nil {
					log.Errorf("ZilliqaSyncManager fetchLockDepositEvents - this.db.PutRetry error: %s", err)
				}
				log.Infof("ZilliqaSyncManager fetchLockDepositEvent -  height: %d", height)
			} else {
				log.Infof("ZilliqaSyncManager found event but not on cross chain manager, ignore: %+v\n", event)
			}
		}
	}

	if res.TotalCount != 0 {
		pages := ((res.TotalCount - 1) / 100) + 1
		for p := 1; p <= pages; p++ {
			if p > 1 {
				res, err = client.TxSearch(height, query, false, p, 100, "asc")
				if err != nil {
					return ccmLockEvents, lockEvents, err
				}
			}
			for _, tx := range res.Txs {
				for _, e := range tx.TxResult.Events {
					if e.Type == _switcheo_crosschainlock {
						tchainId, _ := strconv.ParseUint(string(e.Attributes[5].Value), 10, 32)
						value, _ := hex.DecodeString(string(e.Attributes[6].Value))
						ccmLockEvents = append(ccmLockEvents, &models.ECCMLockEvent{
							Method:   _switcheo_crosschainlock,
							Txid:     string(e.Attributes[1].Value),
							TxHash:   strings.ToLower(tx.Hash.String()),
							User:     string(e.Attributes[3].Value),
							Tchain:   uint32(tchainId),
							Contract: string(e.Attributes[4].Value),
							Height:   height,
							Value:    value,
							Fee:      uint64(tx.TxResult.GasUsed),
						})
					} else if e.Type == _switcheo_lock {
						tchainId, _ := strconv.ParseUint(string(e.Attributes[1].Value), 10, 32)
						amount, _ := decimal.NewFromString(string(e.Attributes[6].Value))
						lockEvents = append(lockEvents, &models.ProxyLockEvent{
							Method:        _switcheo_lock,
							TxHash:        strings.ToLower(tx.Hash.String()),
							FromAddress:   string(e.Attributes[4].Value),
							FromAssetHash: string(e.Attributes[0].Value),
							ToChainId:     uint32(tchainId),
							ToAssetHash:   string(e.Attributes[3].Value),
							ToAddress:     string(e.Attributes[7].Value),
							Amount:        amount.BigInt(),
							DstUser:       string(e.Attributes[5].Value),
						})
					}
				}
			}
		}
	}

	return ccmLockEvents, lockEvents, nil
}

func (this *ZilliqaChainListen) getCosmosCCMUnlockEventByBlockNumber(height uint64) ([]*models.ECCMUnlockEvent, []*models.ProxyUnlockEvent, error) {
	client := this.swthSdk
	ccmUnlockEvents := make([]*models.ECCMUnlockEvent, 0)
	unlockEvents := make([]*models.ProxyUnlockEvent, 0)
	query := fmt.Sprintf("tx.height=%d", height)
	res, err := client.TxSearch(height, query, false, 1, 100, "asc")
	if err != nil {
		return ccmUnlockEvents, unlockEvents, err
	}
	if res.TotalCount != 0 {
		pages := ((res.TotalCount - 1) / 100) + 1
		for p := 1; p <= pages; p++ {
			if p > 1 {
				res, err = client.TxSearch(height, query, false, p, 100, "asc")
				if err != nil {
					return ccmUnlockEvents, unlockEvents, err
				}
			}
			for _, tx := range res.Txs {
				for _, e := range tx.TxResult.Events {
					if e.Type == _switcheo_crosschainunlock {
						fchainId, _ := strconv.ParseUint(string(e.Attributes[2].Value), 10, 32)
						ccmUnlockEvents = append(ccmUnlockEvents, &models.ECCMUnlockEvent{
							Method:   _switcheo_crosschainunlock,
							TxHash:   strings.ToLower(tx.Hash.String()),
							RTxHash:  basedef.HexStringReverse(string(e.Attributes[0].Value)),
							FChainId: uint32(fchainId),
							Contract: string(e.Attributes[3].Value),
							Height:   height,
							Fee:      uint64(tx.TxResult.GasUsed),
						})
					} else if e.Type == _switcheo_unlock {
						amount, _ := decimal.NewFromString(string(e.Attributes[2].Value))
						unlockEvents = append(unlockEvents, &models.ProxyUnlockEvent{
							Method:      _switcheo_unlock,
							TxHash:      strings.ToLower(tx.Hash.String()),
							ToAssetHash: string(e.Attributes[0].Value),
							ToAddress:   string(e.Attributes[1].Value),
							Amount:      amount.BigInt(),
						})
					}
				}
			}
		}
	}

	return ccmUnlockEvents, unlockEvents, nil
}

func (this *ZilliqaChainListen) GetExtendLatestHeight() (uint64, error) {
	if len(this.swthCfg.ExtendNodes) == 0 {
		return this.GetLatestHeight()
	}
	return this.GetLatestHeight()
}
