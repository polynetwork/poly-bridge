package cubescan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"poly-bridge/basedef"
	"poly-bridge/models"
	"time"
)

type CubescanSdk struct {
	client string
}

func NewCubescanSdk() *CubescanSdk {
	return &CubescanSdk{"https://api.cubescan.link/api/trx/volume"}
}

func (c *CubescanSdk) GetCoinPriceAndRank(coins []models.NameAndmarketId) (map[string]float64, map[string]int, error) {
	coinPrice := make(map[string]float64, 0)
	coinRank := make(map[string]int, 0)
	apiURL := fmt.Sprintf("%s?end_timestamp=%d&limit=1&source=coingecko", c.client, time.Now().UnixMilli())

	response, err := http.Get(apiURL)
	if err != nil {
		return coinPrice, coinRank, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return coinPrice, coinRank, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return coinPrice, coinRank, err
	}

	data := result["data"].([]interface{})
	if len(data) > 0 {
		openValue := data[0].(map[string]interface{})["open"].(float64)
		for _, v := range coins {
			coinPrice[v.PriceMarketName] = openValue
			coinRank[v.PriceMarketName] = 0
		}
		return coinPrice, coinRank, nil
	}
	return coinPrice, coinRank, fmt.Errorf("no data")
}

func (c *CubescanSdk) GetMarketName() string {
	return basedef.MARKET_CUBESCAN
}
