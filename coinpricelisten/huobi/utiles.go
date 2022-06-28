package huobi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/models"
	"strings"
)

type HuobiSdk struct {
	client *http.Client
	nodes  []*conf.Restful
}
type HuobiRecentTradeRecord struct {
	Ch     string `json:"ch"`
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   []struct {
		Id   int64 `json:"id"`
		Ts   int64 `json:"ts"`
		Data []struct {
			Id        float64 `json:"id"`
			Ts        int64   `json:"ts"`
			TradeId   int64   `json:"trade-id"`
			Amount    float64 `json:"amount"`
			Price     float64 `json:"price"`
			Direction string  `json:"direction"`
		} `json:"data"`
	} `json:"data"`
}

func DefaultBinanceSdk() *HuobiSdk {
	client := &http.Client{}
	sdk := &HuobiSdk{
		client: client,
		nodes: []*conf.Restful{
			{
				Url: "https://api.huobi.pro/",
			},
		},
	}
	return sdk
}

func NewHuobiSdk() *HuobiSdk {
	client := &http.Client{}

	sdk := &HuobiSdk{
		client: client,
		nodes: []*conf.Restful{
			{
				Url: "https://api.huobi.pro/",
			},
		},
	}
	return sdk
}

func (sdk *HuobiSdk) GetMarketName() string {
	return basedef.MARKET_HUOBI
}

func (this *HuobiSdk) GetCoinPriceAndRank(coins []models.NameAndmarketId) (map[string]float64, map[string]int, error) {
	coinPrice := make(map[string]float64, 0)
	coinRank := make(map[string]int, 0)
	for _, coin := range coins {
		resp, err := this.quotesLatest(coin.PriceMarketName, 0)
		if err != nil {
			return nil, nil, err
		}
		var total float64
		for _, d := range resp.Data {
			total += d.Data[0].Price
		}
		avgPrice := total / float64(len(resp.Data))
		coinPrice[coin.PriceMarketName] = avgPrice
		coinRank[coin.PriceMarketName] = 0
	}
	return coinPrice, coinRank, nil
}

func (sdk *HuobiSdk) quotesLatest(coins string, node int) (HuobiRecentTradeRecord, error) {
	req, err := http.NewRequest("GET", sdk.nodes[node].Url+"market/history/trade", nil)
	if err != nil {
		return HuobiRecentTradeRecord{}, err
	}
	q := url.Values{}
	coins = strings.ToLower(coins)
	q.Add("symbol", coins+"usdt")
	q.Add("size", "10")
	req.Header.Set("Accepts", "application/json")
	req.URL.RawQuery = q.Encode()
	resp, err := sdk.client.Do(req)
	if err != nil {
		return HuobiRecentTradeRecord{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return HuobiRecentTradeRecord{}, fmt.Errorf("response status code: %d", resp.StatusCode)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	var body HuobiRecentTradeRecord
	err = json.Unmarshal(respBody, &body)
	if err != nil {
		return HuobiRecentTradeRecord{}, err
	}
	return body, nil
}
