package pltelfrate

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/models"
)

type ElfRateSdk struct {
	client string
}

func NewElfRateSdk(cfg *conf.CoinPriceListenConfig) *ElfRateSdk {
	for _, v := range cfg.Nodes {
		return &ElfRateSdk{
			client: v.Url,
		}
	}
	return nil
}

type ElfRateResponse struct {
	Data struct {
		CurrencyCode string  `json:"currency_code"`
		Rate         float64 `json:"rate"`
	}
}

func (e *ElfRateSdk) GetCoinPriceAndRank(coins []models.NameAndmarketId) (map[string]float64, map[string]int, error) {
	coinPrice := make(map[string]float64, 0)
	coinRank := make(map[string]int, 0)
	apiURL := e.client

	response, err := http.Get(apiURL)
	if err != nil {
		return coinPrice, coinRank, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return coinPrice, coinRank, err
	}

	result := new(ElfRateResponse)

	err = json.Unmarshal(body, result)
	if err != nil {
		return coinPrice, coinRank, err
	}
	for _, v := range coins {
		coinPrice[v.PriceMarketName] = result.Data.Rate
		coinRank[v.PriceMarketName] = 0
	}
	return coinPrice, coinRank, nil
}

func (e *ElfRateSdk) GetMarketName() string {
	return basedef.MARKET_ELFRATE
}
