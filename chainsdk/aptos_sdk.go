package chainsdk

import (
	"context"
	"fmt"
	"github.com/portto/aptos-go-sdk/client"
	"github.com/portto/aptos-go-sdk/models"
	"net/http"
	"strconv"
	"strings"
)

type AptosSdk struct {
	client client.AptosClient
	url    string
}

type AptosEventFilter struct {
	Address        string
	CreationNumber string
	Query          map[string]interface{}
}

func NewAptosSdk(url string) *AptosSdk {
	aptosClient := client.NewAptosClient(url)
	return &AptosSdk{
		client: aptosClient,
		url:    url,
	}
}

func (sdk *AptosSdk) GetClient() client.AptosClient {
	return sdk.client
}

func (sdk *AptosSdk) GetUrl() string {
	return sdk.url
}

func (sdk *AptosSdk) GetCurrentBlockHeight() (uint64, error) {
	url := sdk.url + "/v1/-/healthy"
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	res, err := httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	heightStr := res.Header.Get("x-aptos-block-height")
	return strconv.ParseUint(heightStr, 10, 64)
}

func (sdk *AptosSdk) GetBlockByNumber(number uint64) (*client.Block, error) {
	block, err := sdk.client.GetBlocksByHeight(context.Background(), number, true)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (sdk *AptosSdk) GetEvents(ctx context.Context, filter *AptosEventFilter) ([]models.Event, error) {
	return sdk.client.GetEventsByCreationNumber(ctx, filter.Address, filter.CreationNumber, filter.Query)
}

func (sdk *AptosSdk) GetTxByVersion(ctx context.Context, version uint64) (*client.TransactionResp, error) {
	return sdk.client.GetTransactionByVersion(ctx, version)
}

func (sdk *AptosSdk) GetBlockByVersion(ctx context.Context, version uint64) (*client.Block, error) {
	return sdk.client.GetBlocksByVersion(ctx, version, true)
}

func (sdk *AptosSdk) GetLockEvent(events []models.Event) *models.Event {
	for _, event := range events {
		if strings.HasSuffix(event.Type, "lock_proxy::LockEvent") {
			return &event
		}
	}
	return nil
}

func (sdk *AptosSdk) GetLockWithFeeEvent(events []models.Event) *models.Event {
	for _, event := range events {
		if strings.HasSuffix(event.Type, "wrapper_v1::LockWithFeeEvent") {
			return &event
		}
	}
	return nil
}

func (sdk *AptosSdk) GetUnLockEvent(events []models.Event) *models.Event {
	for _, event := range events {
		if strings.HasSuffix(event.Type, "lock_proxy::UnlockEvent") {
			return &event
		}
	}
	return nil
}

func (sdk *AptosSdk) GetGasPrice(ctx context.Context) (uint64, error) {
	return sdk.client.EstimateGasPrice(ctx)
}

func (sdk *AptosSdk) GetBalance(ctx context.Context, token, address string) (*client.AccountResource, error) {
	return sdk.client.GetResourceByAccountAddressAndResourceType(ctx, address, fmt.Sprintf("%s::lock_proxy::Treasury<%s>", "0x"+strings.TrimPrefix(address, "0x"), token))
}

//func (sdk *AptosSdk) GetTransactionInfoByHash(hash string) (*client.TransactionInfo, error) {
//	tx, err := sdk.client.GetTransactionInfoByHash(context.Background(), hash)
//	if err != nil {
//		return nil, err
//	}
//	return tx, nil
//}
//
//func (sdk *AptosSdk) GetGasPrice() (int, error) {
//	return sdk.client.GetGasUnitPrice(context.Background())
//}
//
//func (sdk *AptosSdk) GetBalance(tokenHash string, genesisAccountAddress string) (*big.Int, error) {
//	lockTreasuryTag := "::LockProxy::LockTreasury"
//	resType := fmt.Sprintf("%s%s<%s>", genesisAccountAddress, lockTreasuryTag, tokenHash)
//	getResOption := client.GetResourceOption{
//		Decode: true,
//	}
//	lockRes := new(LockTreasuryResource)
//	logs.Info("genesisAccountAddress=%s", genesisAccountAddress)
//	logs.Info("resType=%s", resType)
//	r, err := sdk.client.GetResource(context.Background(), genesisAccountAddress, resType, getResOption, lockRes)
//	if err != nil {
//		return new(big.Int).SetUint64(0), err
//	}
//	if lockRes, ok := r.(*LockTreasuryResource); ok && lockRes != nil {
//		return &lockRes.Json.Token.Value, nil
//	}
//	return new(big.Int).SetUint64(0), nil
//}
