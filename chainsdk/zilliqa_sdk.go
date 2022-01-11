package chainsdk

import (
	"encoding/json"
	"fmt"
	"github.com/Zilliqa/gozilliqa-sdk/core"
	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/beego/beego/v2/core/logs"
	"io/ioutil"
	"math/big"
	"net/http"
	"poly-bridge/utils/decimal"
	"strconv"
	"strings"
)

type ZilliqaSdk struct {
	client *provider.Provider
	url    string
}

func NewZilliqaSdk(url string) *ZilliqaSdk {
	zilClient := provider.NewProvider(url)
	return &ZilliqaSdk{
		client: zilClient,
		url:    url,
	}
}

func (zs *ZilliqaSdk) GetUrl() string {
	return zs.url
}

func (zs *ZilliqaSdk) GetCurrentBlockHeight() (uint64, error) {
	txBlock, err := zs.client.GetLatestTxBlock()
	if err != nil {
		logs.Error("ZilliqaSdk GetCurrentBlockHeight - cannot getLatestTxBlock, err: %s\n", err.Error())
		return 0, err
	}
	if txBlock.Header.BlockNum != "" {
		blockNumber, err1 := strconv.ParseUint(txBlock.Header.BlockNum, 10, 32)
		if err1 != nil {
			logs.Error("ZilliqaSdk GetCurrentBlockHeight - cannot parse block height, err: %s\n", err1.Error())
			return 0, err1
		}
		return blockNumber, nil
	}
	return 0, err
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
		if strings.Contains(err.Error(), "TxBlock has no transactions") || strings.Contains(err.Error(), "Failed to get Microblock") {
			return &ZilBlock{
				tt,
				[]core.Transaction{},
			}, nil
		} else {
			logs.Error("ZilliqaSdk get transactions for tx block %d failed: %s\n", height, err.Error())
			return nil, err
		}
	}
	zilBlock.Transactions = transactions
	return zilBlock, nil
}

func (zs *ZilliqaSdk) GetMinimumGasPrice() (string, error) {
	return zs.client.GetMinimumGasPrice()
}

func (zs *ZilliqaSdk) GetTokenBalance(tokenhash, addrhash string) (*big.Int, error) {
	/*	curl -X POST \
		https://api.zilliqa.com/ \
			-H 'cache-control: no-cache' \
			-H 'content-type: application/json' \
			-H 'postman-token: cbe304e8-0db3-11bb-2fe3-ecbbbe4d4a35' \
			-d '{
			"id": "1",
				"jsonrpc": "2.0",
				"method": "GetSmartContractSubState",
				"params": ["75fa7d8ba6bed4a68774c758a5e43cfb6633d9d6","balances",["0x7772ba52a3474203e3bd41a6821bc51d56d9895f"]]
		}'
	*/
	addrhash = "0x" + addrhash
	url := zs.url
	type zilBalancereq struct {
		Id      string      `json:"id"`
		Jsonrpc string      `json:"jsonrpc"`
		Method  string      `json:"method"`
		Params  interface{} `json:"params"`
	}

	b := [3]interface{}{}
	b[0] = tokenhash
	b[1] = "balances"
	b[2] = []string{addrhash}

	x := &zilBalancereq{
		"1",
		"2.0",
		"GetSmartContractSubState",
		b,
	}
	requestJson, err := json.Marshal(x)
	if err != nil {
		return new(big.Int).SetUint64(0), fmt.Errorf("zil GetTokenBalance err %v", err)
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(string(requestJson)))
	if err != nil {
		return new(big.Int).SetUint64(0), fmt.Errorf("zil GetTokenBalance err %v", err)
	}
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("postman-token", "cbe304e8-0db3-11bb-2fe3-ecbbbe4d4a35")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return new(big.Int).SetUint64(0), fmt.Errorf("zil GetTokenBalance err %v", err)
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return new(big.Int).SetUint64(0), fmt.Errorf("zil GetTokenBalance err %v", err)
	}

	type respblance struct {
		Balances map[string]string
	}
	type zilBalanceresp struct {
		Id      string
		Jsonrpc string
		Result  *respblance
	}
	zilBalance := new(zilBalanceresp)
	err = json.Unmarshal(respBody, zilBalance)
	if err != nil {
		return new(big.Int).SetUint64(0), fmt.Errorf("zil GetTokenBalance err %v", err)
	}
	if zilBalance.Result == nil || zilBalance.Result.Balances == nil {
		return new(big.Int).SetUint64(0), fmt.Errorf("zil GetTokenBalance err %v", err)
	}
	if _, ok := (zilBalance.Result.Balances)[addrhash]; !ok {
		return new(big.Int).SetUint64(0), fmt.Errorf("zil GetTokenBalance err %v", err)
	}
	amount, err := decimal.NewFromString((zilBalance.Result.Balances)[addrhash])
	if err != nil {
		return new(big.Int).SetUint64(0), fmt.Errorf("zil GetTokenBalance err %v", err)
	}
	return amount.BigInt(), nil
}
