package bfclisten

import (
	"encoding/hex"
	"fmt"
	suimodels "github.com/block-vision/sui-go-sdk/models"
	polycommon "github.com/polynetwork/poly/common"
	common2 "github.com/polynetwork/poly/native/service/cross_chain_manager/common"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"poly-bridge/models"
	"strconv"
	"strings"
)

type BfcChainListen struct {
	BfcCfg *conf.ChainListenConfig
	BfcSdk *chainsdk.BfcSdkPro
}

func NewBfcChainListen(cfg *conf.ChainListenConfig) *BfcChainListen {
	bfcListen := &BfcChainListen{}
	bfcListen.BfcCfg = cfg
	urls := cfg.GetNodesUrl()
	sdk := chainsdk.NewBfcSdkPro(urls, cfg.ListenSlot, cfg.ChainId)
	bfcListen.BfcSdk = sdk
	return bfcListen
}

func (b *BfcChainListen) GetExtendLatestHeight() (uint64, error) {
	return b.GetLatestHeight()
}

func (b *BfcChainListen) GetLatestHeight() (uint64, error) {
	return b.BfcSdk.GetTotalTransactionBlocks()
}

func (b *BfcChainListen) GetChainListenSlot() uint64 {
	return b.BfcCfg.ListenSlot
}

func (b *BfcChainListen) GetChainId() uint64 {
	return b.BfcCfg.ChainId
}

func (b *BfcChainListen) GetChainName() string {
	return b.BfcCfg.ChainName
}

func (b *BfcChainListen) GetDefer() uint64 {
	return b.BfcCfg.Defer
}

func (b *BfcChainListen) GetBatchSize() uint64 {
	return b.BfcCfg.BatchSize
}

func (b *BfcChainListen) GetBatchLength() (uint64, uint64) {
	return b.BfcCfg.MinBatchLength, b.BfcCfg.MaxBatchLength
}

func (b *BfcChainListen) HandleNewBatchBlock(start, end uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, int, int, error) {
	return nil, nil, nil, nil, 0, 0, nil
}

func (b *BfcChainListen) HandleNewBlock(height uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, []*models.WrapperDetail, []*models.PolyDetail, int, int, error) {
	return nil, nil, nil, nil, nil, nil, 0, 0, nil
}

func (b *BfcChainListen) HandleSrcEvent(crossChainEvent suimodels.PaginatedEventsResponse, crossChainEventCursor string) ([]*models.SrcTransaction, error) {
	srcTxs := make([]*models.SrcTransaction, 0)
	if len(crossChainEvent.Data) > 0 {
		srcTransaction, err := b.handleSrcTransaction(crossChainEvent)
		if err != nil {
			return srcTxs, fmt.Errorf("handleSrcTransaction crossChainEventCursor: %v err: %v", crossChainEventCursor, err)
		}
		checkpoint, err := b.BfcSdk.GetCheckpoint(srcTransaction.Hash)
		if err != nil {
			return srcTxs, fmt.Errorf("GetCheckpoint hash: %v err: %v", srcTransaction.Hash, err)
		}
		srcTransaction.Height = checkpoint
		lockEvent, err := b.QueryLockEvent(crossChainEventCursor, b.BfcCfg.CCMContract)
		if err != nil {
			return srcTxs, fmt.Errorf("QueryLockEvent crossChainEventCursor: %v err: %v", crossChainEventCursor, err)
		}
		if len(lockEvent.Data) > 0 {
			srcTransfer, err := b.handleSrcTransfer(lockEvent)
			if err != nil {
				return srcTxs, fmt.Errorf("handleSrcTransfer crossChainEventCursor: %v srcTransaction.hash: %v err: %v", crossChainEventCursor, srcTransaction.Hash, err)
			}
			if srcTransaction.Hash == srcTransfer.TxHash {
				srcTransaction.SrcTransfer = srcTransfer
			}
		}
		srcTxs = append(srcTxs, srcTransaction)
	}
	return srcTxs, nil
}

func (b *BfcChainListen) HandleDstEvent(verifyHeaderAndExecuteTxEvent suimodels.PaginatedEventsResponse, verifyHeaderAndExecuteTxEventCursor string) ([]*models.DstTransaction, error) {
	dstTxs := make([]*models.DstTransaction, 0)

	if len(verifyHeaderAndExecuteTxEvent.Data) > 0 {
		dstTransaction, err := b.handleDstTransaction(verifyHeaderAndExecuteTxEvent)
		if err != nil {
			return dstTxs, fmt.Errorf("handleDstTransaction verifyHeaderAndExecuteTxEventCursor: %v err: %v", verifyHeaderAndExecuteTxEventCursor, err)
		}
		checkpoint, err := b.BfcSdk.GetCheckpoint(dstTransaction.Hash)
		if err != nil {
			return dstTxs, fmt.Errorf("GetCheckpoint hash: %v err: %v", dstTransaction.Hash, err)
		}
		dstTransaction.Height = checkpoint
		unlockEvent, err := b.QueryUnlockEvent(verifyHeaderAndExecuteTxEventCursor, b.BfcCfg.CCMContract)
		if err != nil {
			return dstTxs, fmt.Errorf("QueryUnlockEvent verifyHeaderAndExecuteTxEventCursor: %v err: %v", verifyHeaderAndExecuteTxEventCursor, err)
		}
		if len(unlockEvent.Data) > 0 {
			dstTransfer, err := b.handleDstTransfer(unlockEvent)
			if err != nil {
				return dstTxs, fmt.Errorf("handleDstTransfer verifyHeaderAndExecuteTxEventCursor: %v dstTransaction.hash: %v err: %v", verifyHeaderAndExecuteTxEventCursor, dstTransaction.Hash, err)
			}
			if dstTransaction.Hash == dstTransfer.TxHash {
				dstTransaction.DstTransfer = dstTransfer
			}
		}
		dstTxs = append(dstTxs, dstTransaction)
	}
	return dstTxs, nil
}

//ccmAddress has 0x and low string
func (b *BfcChainListen) QueryCrossChainEvent(crossChainEventCursor, ccmAddress string) (suimodels.PaginatedEventsResponse, error) {
	crossChainEventType := ccmAddress + "::events::CrossChainEvent"
	return b.BfcSdk.QueryEvents(crossChainEventType, crossChainEventCursor, 1)
}

func (b *BfcChainListen) QueryLockEvent(crossChainEventCursor, ccmAddress string) (suimodels.PaginatedEventsResponse, error) {
	lockEventType := ccmAddress + "::events::LockEvent"
	return b.BfcSdk.QueryEvents(lockEventType, crossChainEventCursor, 1)
}

func (b *BfcChainListen) QueryVerifyHeaderAndExecuteTxEvent(verifyHeaderAndExecuteTxEventCursor, ccmAddress string) (suimodels.PaginatedEventsResponse, error) {
	verifyHeaderAndExecuteTxEventType := ccmAddress + "::events::VerifyHeaderAndExecuteTxEvent"
	return b.BfcSdk.QueryEvents(verifyHeaderAndExecuteTxEventType, verifyHeaderAndExecuteTxEventCursor, 1)
}

func (b *BfcChainListen) QueryUnlockEvent(verifyHeaderAndExecuteTxEventCursor, ccmAddress string) (suimodels.PaginatedEventsResponse, error) {
	unlockEventType := ccmAddress + "::events::UnlockEvent"
	return b.BfcSdk.QueryEvents(unlockEventType, verifyHeaderAndExecuteTxEventCursor, 1)
}

func (b *BfcChainListen) handleSrcTransaction(crossChainEvent suimodels.PaginatedEventsResponse) (*models.SrcTransaction, error) {
	event := crossChainEvent.Data[0]

	if strings.EqualFold(event.PackageId[3:67], strings.TrimPrefix(b.BfcCfg.CCMContract, "0x")) {
		raw_data, ok := event.ParsedJson["raw_data"]
		if !ok {
			return nil, fmt.Errorf("raw_data err")
		}
		data, ok := raw_data.([]interface{})
		if !ok {
			return nil, fmt.Errorf("raw_data interface err")
		}
		rawData := make([]byte, len(data))
		for i, v := range data {
			rawData[i] = uint8(v.(float64))
		}
		param := &common2.MakeTxParam{}
		_ = param.Deserialization(polycommon.NewZeroCopySource(rawData))

		srcTransaction := &models.SrcTransaction{}
		srcTransaction.ChainId = b.GetChainId()
		srcTransaction.Hash = event.Id.TxDigest
		srcTransaction.State = 1
		srcTransaction.User = event.Sender
		srcTransaction.DstChainId = func() uint64 {
			dstChainId, _ := strconv.Atoi(event.ParsedJson["to_chain_id"].(string))
			return uint64(dstChainId)
		}()
		srcTransaction.Contract = strings.TrimPrefix(b.BfcCfg.CCMContract, "0x")
		srcTransaction.Key = hex.EncodeToString(param.CrossChainID)
		srcTransaction.Param = hex.EncodeToString(rawData)
		return srcTransaction, nil
	}
	return nil, fmt.Errorf("PackageId not EqualFold ccmAddress")
}

func (b *BfcChainListen) handleSrcTransfer(lockEvent suimodels.PaginatedEventsResponse) (*models.SrcTransfer, error) {
	event := lockEvent.Data[0]
	if strings.EqualFold(event.PackageId[3:67], strings.TrimPrefix(b.BfcCfg.CCMContract, "0x")) {
		srcTransfer := &models.SrcTransfer{}
		srcTransfer.ChainId = b.GetChainId()
		srcTransfer.TxHash = event.Id.TxDigest
		srcTransfer.From = event.ParsedJson["from_address"].(string)
		srcTransfer.To = strings.TrimPrefix(b.BfcCfg.CCMContract, "0x")
		srcTransfer.Asset = func() string {
			asset := (event.ParsedJson["from_asset"].(map[string]interface{}))["name"].(string)
			if len(asset) > 76 {
				return asset[76:]
			}
			return asset
		}()
		srcTransfer.Amount = func() *models.BigInt {
			amount, _ := new(big.Int).SetString(event.ParsedJson["amount"].(string), 10)
			return models.NewBigInt(amount)
		}()
		srcTransfer.DstChainId = func() uint64 {
			dstChainId, _ := strconv.Atoi(event.ParsedJson["to_chain_id"].(string))
			return uint64(dstChainId)
		}()

		srcTransfer.DstAsset = convertToHex(event.ParsedJson, "to_asset_hash")
		srcTransfer.DstUser = strings.ToLower(strings.TrimPrefix(convertToString(event.ParsedJson, "to_address"), "0x"))
		return srcTransfer, nil
	}
	return nil, fmt.Errorf("PackageId not EqualFold ccmAddress")
}

func (b *BfcChainListen) handleDstTransaction(verifyHeaderAndExecuteTxEvent suimodels.PaginatedEventsResponse) (*models.DstTransaction, error) {
	event := verifyHeaderAndExecuteTxEvent.Data[0]
	if strings.EqualFold(event.PackageId[3:67], strings.TrimPrefix(b.BfcCfg.CCMContract, "0x")) {
		dstTransaction := &models.DstTransaction{}
		dstTransaction.ChainId = b.GetChainId()
		dstTransaction.Hash = event.Id.TxDigest
		dstTransaction.State = 1
		dstTransaction.SrcChainId = func() uint64 {
			fromchainid, _ := strconv.Atoi(event.ParsedJson["from_chain_id"].(string))
			return uint64(fromchainid)
		}()
		dstTransaction.Contract = strings.TrimPrefix(b.BfcCfg.CCMContract, "0x")
		dstTransaction.PolyHash = basedef.HexStringReverse(convertToHex(event.ParsedJson, "cross_chain_tx_hash"))
		return dstTransaction, nil
	}
	return nil, fmt.Errorf("PackageId not EqualFold ccmAddress")
}

func (b *BfcChainListen) handleDstTransfer(unlockEvent suimodels.PaginatedEventsResponse) (*models.DstTransfer, error) {
	event := unlockEvent.Data[0]
	if strings.EqualFold(event.PackageId[3:67], strings.TrimPrefix(b.BfcCfg.CCMContract, "0x")) {
		dstTransfer := &models.DstTransfer{}
		dstTransfer.TxHash = event.Id.TxDigest
		dstTransfer.ChainId = b.GetChainId()
		dstTransfer.From = strings.TrimPrefix(b.BfcCfg.CCMContract, "0x")
		dstTransfer.To = event.ParsedJson["to_address"].(string)
		dstTransfer.Asset = func() string {
			asset := (event.ParsedJson["to_asset"].(map[string]interface{}))["name"].(string)
			if len(asset) > 76 {
				return asset[76:]
			}
			return asset
		}()
		dstTransfer.Amount = func() *models.BigInt {
			amount, _ := new(big.Int).SetString(event.ParsedJson["amount"].(string), 10)
			return models.NewBigInt(amount)
		}()
		return dstTransfer, nil
	}
	return nil, fmt.Errorf("PackageId not EqualFold ccmAddress")
}

func convertToHex(parsedJson map[string]interface{}, key string) string {
	value, ok := parsedJson[key]
	if !ok {
		fmt.Println("convertToHex parsedJson err key:", key)
		return ""
	}
	data, ok := value.([]interface{})
	if !ok {
		fmt.Println("convertToHex interface err key:", key)
		return ""
	}
	byteData := make([]byte, len(data))
	for i, v := range data {
		byteData[i] = uint8(v.(float64))
	}
	return hex.EncodeToString(byteData)
}

func convertToString(parsedJson map[string]interface{}, key string) string {
	value, ok := parsedJson[key]
	if !ok {
		fmt.Println("convertToString parsedJson err key:", key)
		return ""
	}
	data, ok := value.([]interface{})
	if !ok {
		fmt.Println("convertToString interface err key:", data)
		return ""
	}
	byteData := make([]byte, len(data))
	for i, v := range data {
		byteData[i] = uint8(v.(float64))
	}
	return string(byteData)
}
