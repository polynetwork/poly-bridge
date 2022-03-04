package chainsdk

import (
	"context"
	"fmt"
	"github.com/starcoinorg/starcoin-go/client"
	"math/big"
)

type StarCoinSdk struct {
	client *client.StarcoinClient
	url    string
}

type LockTreasuryResource struct {
	Raw  string `json:"raw"`
	Json struct {
		Token struct {
			Value big.Int `json:"value"`
		} `json:"token"`
	} `json:"json"`
}

func NewStarCoinSdk(url string) *StarCoinSdk {
	starcoinClient := client.NewStarcoinClient(url)
	return &StarCoinSdk{
		client: &starcoinClient,
		url:    url,
	}
}

func (sdk *StarCoinSdk) GetClient() *client.StarcoinClient {
	return sdk.client
}

func (sdk *StarCoinSdk) GetUrl() string {
	return sdk.url
}

func (sdk *StarCoinSdk) GetCurrentBlockHeight() (uint64, error) {
	nodeInfo, err := sdk.client.GetNodeInfo(context.Background())
	if err != nil {
		return 0, err
	}
	currentHeight, err := nodeInfo.GetBlockNumber()
	if err != nil {
		return 0, err
	}
	return currentHeight, nil
}

func (sdk *StarCoinSdk) GetBlockByNumber(number uint64) (*client.Block, error) {
	block, err := sdk.client.GetBlockByNumber(context.Background(), int(number))
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (sdk *StarCoinSdk) GetEvents(filter *client.EventFilter) ([]client.Event, error) {
	events, err := sdk.client.GetEvents(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (sdk *StarCoinSdk) GetTransactionInfoByHash(hash string) (*client.TransactionInfo, error) {
	tx, err := sdk.client.GetTransactionInfoByHash(context.Background(), hash)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (sdk *StarCoinSdk) GetGasPrice() (int, error) {
	return sdk.client.GetGasUnitPrice(context.Background())
}

func (sdk *StarCoinSdk) GetBalance(tokenHash string, genesisAccountAddress string) (*big.Int, error) {
	lockTreasuryTag := "::LockProxy::LockTreasury"
	resType := fmt.Sprintf("%s%s<%s>", genesisAccountAddress, lockTreasuryTag, tokenHash)
	getResOption := client.GetResourceOption{
		Decode: true,
	}
	lockRes := new(LockTreasuryResource)
	r, err := sdk.client.GetResource(context.Background(), genesisAccountAddress, resType, getResOption, lockRes)
	if err != nil {
		return new(big.Int).SetUint64(0), err
	}
	if lockRes, ok := r.(*LockTreasuryResource); ok && lockRes != nil {
		return &lockRes.Json.Token.Value, nil
	}
	return new(big.Int).SetUint64(0), nil
}
