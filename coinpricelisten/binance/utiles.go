/*
 * Copyright (C) 2020 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package binance

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"io/ioutil"
	"net/http"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/models"
)

type BinanceSdk struct {
	client *http.Client
	nodes  []*conf.Restful
}

func DefaultBinanceSdk() *BinanceSdk {
	client := &http.Client{}
	sdk := &BinanceSdk{
		client: client,
		nodes: []*conf.Restful{
			{
				Url: "https://api1.binance.com/",
			},
		},
	}
	return sdk
}

func NewBinanceSdk(cfg *conf.CoinPriceListenConfig) *BinanceSdk {
	client := &http.Client{}
	sdk := &BinanceSdk{
		client: client,
		nodes:  cfg.Nodes,
	}
	return sdk
}

func (sdk *BinanceSdk) QuotesLatest() ([]*Ticker, error) {
	for i := 0; i < len(sdk.nodes); i++ {
		quotes, err := sdk.quotesLatest(i)
		if err != nil {
			logs.Error("Binance QuotesLatest err: %s", err.Error())
			continue
		} else {
			return quotes, nil
		}
	}
	return nil, fmt.Errorf("Cannot get Binance QuotesLatest!")
}

func (sdk *BinanceSdk) quotesLatest(node int) ([]*Ticker, error) {
	req, err := http.NewRequest("GET", sdk.nodes[node].Url+"api/v3/ticker/price", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "8efe5156-8b37-4c77-8e1d-a140c97bf466")

	resp, err := sdk.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response status code: %d", resp.StatusCode)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	tickers := make([]*Ticker, 0)
	err = json.Unmarshal(respBody, &tickers)
	if err != nil {
		return nil, err
	}
	return tickers, nil
}

func (sdk *BinanceSdk) GetMarketName() string {
	return basedef.MARKET_BINANCE
}

func (this *BinanceSdk) GetCoinPriceAndRank(coins []models.NameAndmarketId) (map[string]float64, map[string]int, error) {
	quotes, err := this.QuotesLatest()
	if err != nil {
		return nil, nil, err
	}
	coinSymbol2Price := make(map[string]float64, 0)
	for _, v := range quotes {
		coinSymbol2Price[v.Symbol] = v.Price
	}
	coinPrice := make(map[string]float64, 0)
	coinRank := make(map[string]int, 0)
	for _, coin := range coins {
		price, ok := coinSymbol2Price[coin.PriceMarketName]
		if !ok {
			logs.Warn("There is no coin price %s in Binance!", coin)
			continue
		}
		coinPrice[coin.PriceMarketName] = price
		coinRank[coin.PriceMarketName] = 0
	}
	return coinPrice, coinRank, nil
}
