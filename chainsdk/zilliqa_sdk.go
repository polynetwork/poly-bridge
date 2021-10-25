package chainsdk

import (
	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/beego/beego/v2/core/logs"
	"strconv"
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


