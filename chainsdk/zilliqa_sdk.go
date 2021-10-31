package chainsdk

import (
	"github.com/Zilliqa/gozilliqa-sdk/bech32"
	"github.com/Zilliqa/gozilliqa-sdk/core"
	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/Zilliqa/gozilliqa-sdk/util"
	"github.com/beego/beego/v2/core/logs"
	"github.com/shopspring/decimal"
	"math/big"
	"strconv"
	"strings"
)

type ZilliqaSdk struct {
	client *provider.Provider
}

func NewZilliqaSdk(url string) *ZilliqaSdk {
	zilClient := provider.NewProvider(url)
	return &ZilliqaSdk{
		client: zilClient,
	}
}

func (zs *ZilliqaSdk) GetCurrentBlockHeight() (uint64, error) {
	txBlock, err := zs.client.GetLatestTxBlock()
	if err != nil {
		logs.Error("ZilliqaSdk GetCurrentBlockHeight - cannot getLatestTxBlock, err: %s\n", err.Error())
	}
	blockNumber, err1 := strconv.ParseUint(txBlock.Header.BlockNum, 10, 32)
	if err1 != nil {
		logs.Error("ZilliqaSdk GetCurrentBlockHeight - cannot parse block height, err: %s\n", err1.Error())
	}
	return blockNumber, err
}

type ZilBlock struct {
	Timestamp    uint64
	Transactions []core.Transaction
}

func (zs *ZilliqaSdk) GetBlock(height uint64) (*ZilBlock, error) {
	zilBlock := new(ZilBlock)
	txBlockT, err := zs.client.GetTxBlock(strconv.FormatUint(height, 10))
	if err != nil {
		return nil, err
	}
	timestamp, err := decimal.NewFromString(txBlockT.Header.Timestamp)
	if err != nil {
		return nil, err
	}
	tt := timestamp.Div(decimal.New(1, 6)).BigInt().Uint64()
	zilBlock.Timestamp = tt
	transactions, err := zs.client.GetTxnBodiesForTxBlock(strconv.FormatUint(height, 10))
	if err != nil {
		if strings.Contains(err.Error(), "TxBlock has no transactions") {
			logs.Error("ZilliqaSdk no transaction in block %d\n", height)
			return nil, err
		} else {
			logs.Error("ZilliqaSdk get transactions for tx block %d failed: %s\n", height, err.Error())
			return nil, err
		}
	}
	zilBlock.Transactions = transactions
	return zilBlock, nil
}

func (s *ZilliqaSyncManager) fetchLockDepositEvents(height uint64) bool {
	transactions, err := s.zilSdk.GetTxnBodiesForTxBlock(strconv.FormatUint(height, 10))
	if err != nil {
		if strings.Contains(err.Error(), "TxBlock has no transactions") {
			log.Infof("ZilliqaSyncManager no transaction in block %d\n", height)
			return true
		} else {
			log.Infof("ZilliqaSyncManager get transactions for tx block %d failed: %s\n", height, err.Error())
			return false
		}
	}

	for _, transaction := range transactions {
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

	return true
}
