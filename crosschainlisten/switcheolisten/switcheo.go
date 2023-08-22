package switcheolisten

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"poly-bridge/models"

	"github.com/beego/beego/v2/core/logs"
)

const (
	_switcheo_crosschainlock   = "make_from_cosmos_proof"
	_switcheo_crosschainunlock = "verify_to_cosmos_proof"
	_switcheo_lock             = "Switcheo.carbon.lockproxy.LockEvent"
	_switcheo_unlock           = "Switcheo.carbon.lockproxy.UnlockEvent"
)

type SwitcheoChainListen struct {
	swthCfg *conf.ChainListenConfig
	swthSdk *chainsdk.SwitcheoSdkPro
}

func NewSwitcheoChainListen(cfg *conf.ChainListenConfig) *SwitcheoChainListen {
	swthListen := &SwitcheoChainListen{}
	swthListen.swthCfg = cfg
	urls := cfg.GetNodesUrl()
	sdk := chainsdk.NewSwitcheoSdkPro(urls, cfg.ListenSlot, cfg.ChainId)
	swthListen.swthSdk = sdk
	return swthListen
}

func (this *SwitcheoChainListen) GetLatestHeight() (uint64, error) {
	return this.swthSdk.GetLatestHeight()
}

func (this *SwitcheoChainListen) GetChainListenSlot() uint64 {
	return this.swthCfg.ListenSlot
}

func (this *SwitcheoChainListen) GetChainId() uint64 {
	return this.swthCfg.ChainId
}

func (this *SwitcheoChainListen) GetChainName() string {
	return this.swthCfg.ChainName
}

func (this *SwitcheoChainListen) GetDefer() uint64 {
	return this.swthCfg.Defer
}

func (this *SwitcheoChainListen) GetBatchSize() uint64 {
	return this.swthCfg.BatchSize
}
func (this *SwitcheoChainListen) GetBatchLength() (uint64, uint64) {
	return this.swthCfg.MinBatchLength, this.swthCfg.MaxBatchLength
}

func (this *SwitcheoChainListen) HandleNewBatchBlock(start, end uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, int, int, error) {
	return nil, nil, nil, nil, 0, 0, nil
}

func (this *SwitcheoChainListen) HandleNewBlock(height uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, []*models.WrapperDetail, []*models.PolyDetail, int, int, error) {
	block, err := this.swthSdk.GetBlockByHeight(height)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, 0, 0, err
	}
	if block == nil {
		return nil, nil, nil, nil, nil, nil, 0, 0, fmt.Errorf("there is no switcheo block!")
	}
	parsedTime, _ := time.Parse(time.RFC3339Nano, block.Result.Block.Header.Time)
	tt := uint64(parsedTime.Unix())
	srcTransactions := make([]*models.SrcTransaction, 0)
	dstTransactions := make([]*models.DstTransaction, 0)

	ccmLockEvent, lockEvents, err := this.getCosmosCCMLockEventByBlockNumber(height)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, 0, 0, err
	}

	ccmUnlockEvent, unlockEvents, err := this.getCosmosCCMUnlockEventByBlockNumber(height)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, 0, 0, err
	}

	for _, lockEvent := range ccmLockEvent {
		if lockEvent.Method == _switcheo_crosschainlock {
			logs.Info("from chain: %v, height: %d, txhash: %s\n", this.GetChainName(), height, lockEvent.TxHash)
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
					if srcTransfer.DstChainId == basedef.APTOS_CROSSCHAIN_ID {
						aptosAsset, err := hex.DecodeString(v.ToAssetHash)
						if err == nil {
							v.ToAssetHash = string(aptosAsset)
						}
					}
					srcTransfer.DstAsset = v.ToAssetHash
					srcTransfer.DstAsset = models.FormatAssert(srcTransfer.DstAsset)
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
			logs.Info("to chain: %s, height: %d, txhash: %s\n", this.GetChainName(), height, unLockEvent.TxHash)
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
	return nil, srcTransactions, nil, dstTransactions, nil, nil, len(srcTransactions), len(dstTransactions), nil
}

func (this *SwitcheoChainListen) getCosmosCCMLockEventByBlockNumber(height uint64) ([]*models.ECCMLockEvent, []*models.ProxyLockEvent, error) {
	client := this.swthSdk
	ccmLockEvents := make([]*models.ECCMLockEvent, 0)
	lockEvents := make([]*models.ProxyLockEvent, 0)
	query := fmt.Sprintf("tx.height=%d AND make_from_cosmos_proof.status='1'", height)
	res, err := client.TxSearch(height, query, false, 1, 100, "asc")
	if err != nil {
		return ccmLockEvents, lockEvents, err
	}
	totalCount, _ := strconv.ParseInt(res.Result.TotalCount, 10, 32)
	if totalCount > 0 {
		pages := int((totalCount-1)/100) + 1
		for p := 1; p <= pages; p++ {
			if p > 1 {
				res, err = client.TxSearch(height, query, false, p, 100, "asc")
				if err != nil {
					return ccmLockEvents, lockEvents, err
				}
			}
			for _, tx := range res.Result.Txs {
				for _, e := range tx.TxResult.Events {
					if e.Type == _switcheo_crosschainlock {
						tchainId, _ := strconv.ParseUint(string(e.Attributes[5].Value), 10, 32)
						value, _ := hex.DecodeString(string(e.Attributes[6].Value))
						ccmLockEvents = append(ccmLockEvents, &models.ECCMLockEvent{
							Method:   _switcheo_crosschainlock,
							Txid:     string(e.Attributes[1].Value),
							TxHash:   strings.ToLower(tx.Hash),
							User:     string(e.Attributes[3].Value),
							Tchain:   uint32(tchainId),
							Contract: string(e.Attributes[4].Value),
							Height:   height,
							Value:    value,
							Fee: func() uint64 {
								x, _ := strconv.ParseUint(tx.TxResult.GasUsed, 10, 64)
								return x
							}(),
						})
					} else if e.Type == _switcheo_lock {
						tchainId, _ := strconv.ParseUint(strings.Trim(string(e.Attributes[10].Value), "\\\""), 10, 32)
						amount, _ := new(big.Int).SetString(strings.Trim(string(e.Attributes[0].Value), "\\\""), 10)
						lockEvents = append(lockEvents, &models.ProxyLockEvent{
							Method:        _switcheo_lock,
							TxHash:        strings.ToLower(tx.Hash),
							FromAddress:   strings.Trim(string(e.Attributes[4].Value), "\\\""),
							FromAssetHash: strings.Trim(string(e.Attributes[5].Value), "\\\""),
							ToChainId:     uint32(tchainId),
							ToAssetHash:   strings.Trim(string(e.Attributes[9].Value), "\\\""),
							ToAddress:     strings.Trim(string(e.Attributes[6].Value), "\\\""),
							Amount:        amount,
							DstUser:       strings.Trim(string(e.Attributes[8].Value), "\\\""),
						})
					}
				}
			}
		}
	}

	return ccmLockEvents, lockEvents, nil
}

func (this *SwitcheoChainListen) getCosmosCCMUnlockEventByBlockNumber(height uint64) ([]*models.ECCMUnlockEvent, []*models.ProxyUnlockEvent, error) {
	client := this.swthSdk
	ccmUnlockEvents := make([]*models.ECCMUnlockEvent, 0)
	unlockEvents := make([]*models.ProxyUnlockEvent, 0)
	query := fmt.Sprintf("tx.height=%d AND verify_to_cosmos_proof.from_chain_id>=0", height)
	res, err := client.TxSearch(height, query, false, 1, 100, "asc")
	if err != nil {
		return ccmUnlockEvents, unlockEvents, err
	}
	totalCount, _ := strconv.ParseInt(res.Result.TotalCount, 10, 32)
	if totalCount > 0 {
		pages := int((totalCount-1)/100) + 1
		for p := 1; p <= pages; p++ {
			if p > 1 {
				res, err = client.TxSearch(height, query, false, p, 100, "asc")
				if err != nil {
					return ccmUnlockEvents, unlockEvents, err
				}
			}
			for _, tx := range res.Result.Txs {
				for _, e := range tx.TxResult.Events {
					if e.Type == _switcheo_crosschainunlock {
						fchainId, _ := strconv.ParseUint(string(e.Attributes[2].Value), 10, 32)
						ccmUnlockEvents = append(ccmUnlockEvents, &models.ECCMUnlockEvent{
							Method:   _switcheo_crosschainunlock,
							TxHash:   strings.ToLower(tx.Hash),
							RTxHash:  basedef.HexStringReverse(string(e.Attributes[0].Value)),
							FChainId: uint32(fchainId),
							Contract: string(e.Attributes[3].Value),
							Height:   height,
							Fee: func() uint64 {
								x, _ := strconv.ParseUint(tx.TxResult.GasUsed, 10, 64)
								return x
							}(),
						})
					} else if e.Type == _switcheo_unlock {
						amount, _ := new(big.Int).SetString(strings.Trim(string(e.Attributes[0].Value), "\\\""), 10)
						unlockEvents = append(unlockEvents, &models.ProxyUnlockEvent{
							BlockNumber: height,
							Method:      _switcheo_unlock,
							TxHash:      strings.ToLower(tx.Hash),
							ToAssetHash: strings.Trim(string(e.Attributes[7].Value), "\\\""),
							ToAddress:   strings.Trim(string(e.Attributes[6].Value), "\\\""),
							Amount:      amount,
						})
					}
				}
			}
		}
	}

	return ccmUnlockEvents, unlockEvents, nil
}

func (this *SwitcheoChainListen) GetExtendLatestHeight() (uint64, error) {
	if len(this.swthCfg.ExtendNodes) == 0 {
		return this.GetLatestHeight()
	}
	return this.GetLatestHeight()
}
