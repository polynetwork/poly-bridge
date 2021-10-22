package chainsdk

import (
	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)


type ZilliqaInfo struct {
	sdk *provider.Provider
	latestHeight uint64
}

func NewZilliqaInfo(url string) (*ZilliqaInfo, error) {
	sdk := provider.NewProvider(url)
	if sdk.GetNetworkId
	return &ZilliqaInfo{
		sdk:          sdk,
		latestHeight: 0,
	},nil
}

func (client *ZilliqaSdk) GetCurrentBlockHeight() (uint64, error) {
	client.client.
	var result hexutil.Big
	err := ec.rpcClient.CallContext(context.Background(), &result, "eth_blockNumber")
	for err != nil {
		return 0, err
	}
	return (*big.Int)(&result).Uint64(), err
}
