package chainsdk

import (
	"context"
	"github.com/starcoinorg/starcoin-go/client"
)

type StarCoinSdk struct {
	client *client.StarcoinClient
	url    string
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
