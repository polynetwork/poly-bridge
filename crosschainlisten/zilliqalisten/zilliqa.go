package zilliqalisten

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"poly-bridge/models"
	"poly-bridge/utils/decimal"
	"strconv"
	"strings"
)

const (
	zilliqa_cross_chain                       = "CrossChainEvent"
	zilliqa_lock                              = "Lock"
	zilliqa_transfer_to_proxy                 = "TransferZRC2ToLockProxy"
	ziliqa_verify_header_and_execute_tx_event = "VerifyHeaderAndExecuteTxEvent"
	zilliqa_unlock                            = "Unlock"
)

type ZilliqaChainListen struct {
	zliCfg *conf.ChainListenConfig
	zliSdk *chainsdk.ZilliqaSdkPro
}

func NewZilliqaChainListen(cfg *conf.ChainListenConfig) *ZilliqaChainListen {
	zilListen := &ZilliqaChainListen{}
	zilListen.zliCfg = cfg
	sdk := chainsdk.NewZilliqaSdkPro(cfg.Nodes, cfg.ListenSlot, cfg.ChainId)
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

func (this *ZilliqaChainListen) GetBatchLength() (uint64, uint64) {
	return this.zliCfg.MinBatchLength, this.zliCfg.MaxBatchLength
}

func (this *ZilliqaChainListen) GetExtendLatestHeight() (uint64, error) {
	if len(this.zliCfg.ExtendNodes) == 0 {
		return this.GetLatestHeight()
	}
	//for _, v := range this.zliCfg.ExtendNodes {
	//	height, err := this.getExtendLatestHeight(v.Url)
	//	if err == nil {
	//		return height, nil
	//	}
	//}
	return this.GetLatestHeight()
}

func (this *ZilliqaChainListen) getExtendLatestHeight(url string) (uint64, error) {
	info := chainsdk.NewZilliqaInfo(url)
	return info.GetLastHeight()
}

func (this *ZilliqaChainListen) HandleNewBatchBlock(start, end uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, []*models.WrapperDetail, []*models.PolyDetail, int, int, error) {
	return nil, nil, nil, nil, nil, nil, 0, 0, nil
}

func (this *ZilliqaChainListen) HandleNewBlock(height uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, []*models.WrapperDetail, []*models.PolyDetail, int, int, error) {
	block, err := this.zliSdk.GetBlockByHeight(height)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, 0, 0, err
	}
	if block == nil {
		return nil, nil, nil, nil, nil, nil, 0, 0, fmt.Errorf("there is no zilliqa block!")
	}
	srcTransactions := this.getzilliqaSrcTransactionByBlockNumber(height, block)
	dstTransactions := this.getzilliqaDstTransactionByBlockNumber(height, block)

	return nil, srcTransactions, nil, dstTransactions, nil, nil, len(srcTransactions), len(dstTransactions), nil
}

type txData struct {
	Tag    string `json:"_tag"`
	Params []struct {
		Vname string `json:"vname"`
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"params"`
}

func (this *ZilliqaChainListen) getzilliqaSrcTransactionByBlockNumber(height uint64, block *chainsdk.ZilBlock) []*models.SrcTransaction {
	srcTransactions := make([]*models.SrcTransaction, 0)
	for _, transaction := range block.Transactions {
		if !transaction.Receipt.Success {
			continue
		}
		events := transaction.Receipt.EventLogs
		if len(events) == 0 {
			continue
		}
		srcTransaction := new(models.SrcTransaction)
		srcTransfer := new(models.SrcTransfer)
		for _, event := range events {
			switch event.EventName {
			case zilliqa_cross_chain:
				// 1. contract address should be cross chain manager
				// 2. event name should be CrossChainEvent
				// zilliqa address bech32.ToBech32Address(event.Address)
				addr := event.Address[2:]
				if strings.EqualFold(this.zliCfg.CCMContract, addr) {
					logs.Info("ZilliqaChainListen found src event on cross chain: %+v\n", event)
					srcTransaction.Hash = transaction.ID
					srcTransaction.ChainId = this.GetChainId()
					srcTransaction.Height = height
					srcTransaction.Time = block.Timestamp
					srcTransaction.State = 1
					srcTransaction.Standard = models.TokenTypeErc20
					gasPrice, _ := decimal.NewFromString(transaction.Receipt.CumulativeGas)
					srcTransaction.Fee = models.NewBigInt(gasPrice.BigInt())
					for _, param := range event.Params {
						switch param.VName {
						case "txId":
							srcTransaction.Key = param.Value.(string)[2:]
						case "toChainId":
							toChainId, _ := strconv.ParseUint(param.Value.(string), 10, 64)
							srcTransaction.DstChainId = toChainId
						case "rawdata":
							srcTransaction.Param = param.Value.(string)[2:]
						case "sender":
							srcTransaction.User = param.Value.(string)[2:]
						case "proxyOrAssetContract":
							srcTransaction.Contract = param.Value.(string)[2:]
						}
					}

				}
			case zilliqa_lock:
				srcTransfer.TxHash = transaction.ID
				srcTransfer.ChainId = this.GetChainId()
				srcTransfer.Time = block.Timestamp
				srcTransfer.To = event.Address[2:]
				for _, param := range event.Params {
					switch param.VName {
					case "fromAssetDenom":
						srcTransfer.Asset = param.Value.(string)[2:]
					case "toChainId":
						toChainId, _ := strconv.ParseUint(param.Value.(string), 10, 64)
						srcTransfer.DstChainId = toChainId
					}
				}
				for _, contract := range this.zliCfg.NFTProxyContract {
					if strings.EqualFold(contract, event.Address[2:]) {
						srcTransfer.Standard = models.TokenTypeErc721
						break
					}
				}
			}
		}
		if srcTransaction.Hash != "" {
			data := transaction.Data.(string)
			txD := txData{}
			err := json.Unmarshal([]byte(data), &txD)
			if err != nil {
				logs.Error("fail to marshal tx data, %v", err)
				continue
			}
			txDataMap := extractZilliqatxData(txD)
			for k, v := range txDataMap {
				switch k {
				case "toAddress":
					srcTransfer.DstUser = v[2:]
				case "toAssetDenom":
					srcTransfer.DstAsset = v[2:]
					if srcTransfer.DstChainId == basedef.APTOS_CROSSCHAIN_ID {
						aptosAsset, err := hex.DecodeString(srcTransfer.DstAsset)
						if err == nil {
							srcTransfer.DstAsset = string(aptosAsset)
						}
					}
					srcTransfer.DstAsset = models.FormatAssert(srcTransfer.DstAsset)
				case "amount":
					amount, _ := decimal.NewFromString(v)
					srcTransfer.Amount = models.NewBigInt(amount.BigInt())
				}
			}
			if srcTransfer.TxHash == srcTransaction.Hash {
				srcTransfer.From = srcTransaction.User
				srcTransaction.Standard = srcTransfer.Standard
				srcTransaction.SrcTransfer = srcTransfer
			}
			srcTransactions = append(srcTransactions, srcTransaction)
		}
	}
	return srcTransactions
}

func extractZilliqatxData(data txData) map[string]string {
	if data.Tag != zilliqa_lock {
		return nil
	}
	m := make(map[string]string)
	for _, v := range data.Params {
		m[v.Vname] = v.Value
	}
	return m
}

func (this *ZilliqaChainListen) getzilliqaDstTransactionByBlockNumber(height uint64, block *chainsdk.ZilBlock) []*models.DstTransaction {
	dstTransactions := make([]*models.DstTransaction, 0)
	for _, transaction := range block.Transactions {
		if !transaction.Receipt.Success {
			continue
		}
		events := transaction.Receipt.EventLogs
		if len(events) == 0 {
			continue
		}
		dstTransaction := new(models.DstTransaction)
		dstTransfer := new(models.DstTransfer)
		for _, event := range events {
			if event.EventName == ziliqa_verify_header_and_execute_tx_event {
				addr := event.Address[2:]
				if strings.EqualFold(this.zliCfg.CCMContract, addr) {
					logs.Info("ZilliqaChainListen found dst event on cross chain: %+v\n", event)
					dstTransaction.Hash = transaction.ID
					dstTransaction.ChainId = this.GetChainId()
					dstTransaction.Height = height
					dstTransaction.Time = block.Timestamp
					dstTransaction.State = 1
					dstTransaction.Standard = models.TokenTypeErc20
					gasPrice, _ := decimal.NewFromString(transaction.Receipt.CumulativeGas)
					dstTransaction.Fee = models.NewBigInt(gasPrice.BigInt())
					for _, param := range event.Params {
						switch param.VName {
						case "crossChainTxHash":
							dstTransaction.PolyHash = basedef.HexStringReverse(param.Value.(string)[2:])
						case "fromChainId":
							srcChainId, _ := strconv.ParseUint(param.Value.(string), 10, 64)
							dstTransaction.SrcChainId = srcChainId
						case "toContractAddr":
							dstTransaction.Contract = param.Value.(string)[2:]
						}
					}
				}
			} else if event.EventName == zilliqa_unlock {
				dstTransfer.TxHash = transaction.ID
				dstTransfer.ChainId = this.GetChainId()
				dstTransfer.Time = block.Timestamp
				dstTransfer.From = event.Address[2:]
				for _, param := range event.Params {
					switch param.VName {
					case "toAssetHash":
						dstTransfer.Asset = param.Value.(string)[2:]
					case "toAddressHash":
						dstTransfer.To = param.Value.(string)[2:]
					case "amount":
						amount, _ := decimal.NewFromString(param.Value.(string))
						dstTransfer.Amount = models.NewBigInt(amount.BigInt())
					}
				}
				for _, contract := range this.zliCfg.NFTProxyContract {
					if strings.EqualFold(contract, event.Address[2:]) {
						dstTransfer.Standard = models.TokenTypeErc721
						break
					}
				}
			}
		}
		if dstTransaction.Hash != "" {
			if dstTransfer.TxHash == dstTransaction.Hash {
				dstTransaction.Standard = dstTransfer.Standard
				dstTransaction.DstTransfer = dstTransfer
			}
			dstTransactions = append(dstTransactions, dstTransaction)
		}
	}
	return dstTransactions
}
