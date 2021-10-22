package chainsdk

import (
	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

type ZilliqaSdk struct {
	client *provider.Provider
}

func NewZilliqaSdk(url string) (*ZilliqaSdk, error) {
	zilClient := provider.NewProvider(url)
	return &ZilliqaSdk{
		client: zilClient,
	}, nil
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
