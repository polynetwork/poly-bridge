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

package self

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"io/ioutil"
	"net/http"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/models"
	"strings"
)

type SelfSdk struct {
	client *http.Client
	nodes  []*conf.Restful
}

func NewSelfSdk(cfg *conf.CoinPriceListenConfig) *SelfSdk {
	client := &http.Client{}
	sdk := &SelfSdk{
		client: client,
		nodes:  cfg.Nodes,
	}
	return sdk
}

func (sdk *SelfSdk) QuotesLatest(url string) ([]*Ticker, error) {
	for i := 0; i < len(sdk.nodes); i++ {
		quotes, err := sdk.quotesLatest(i, url)
		if err != nil {
			logs.Error("CoinMarketCap QuotesLatest err: %s", err.Error())
			continue
		} else {
			return quotes, nil
		}
	}
	return nil, fmt.Errorf("Cannot get Binance QuotesLatest!")
}

func (sdk *SelfSdk) quotesLatest(node int, url string) ([]*Ticker, error) {
	req, err := http.NewRequest("GET", sdk.nodes[node].Url+"price/"+url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accepts", "application/json")

	resp, err := sdk.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response status code: %d", resp.StatusCode)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	var rsp Rsp
	err = json.Unmarshal(respBody, &rsp)
	if err != nil {
		return nil, err
	}
	return rsp.Prices, nil
}

func (sdk *SelfSdk) GetMarketName() string {
	return basedef.MARKET_SELF
}

func (sdk *SelfSdk) GetCoinPriceAndRank(coins []models.NameAndmarketId) (map[string]float64, map[string]int, error) {
	coinIds := make([]string, 0)
	for _, coin := range coins {
		coinIds = append(coinIds, coin.PriceMarketName)
	}
	//
	requestCoinIds := strings.Join(coinIds, ",")
	quotes, _ := sdk.QuotesLatest(requestCoinIds)

	coinSymbol2Price := make(map[string]float64, 0)
	for _, v := range quotes {
		coinSymbol2Price[v.Symbol] = v.Price
	}
	coinPrice := make(map[string]float64, 0)
	coinRank := make(map[string]int, 0)
	for _, coin := range coins {
		price, ok := coinSymbol2Price[coin.PriceMarketName]
		if !ok {
			logs.Warn("There is no coin price %s in self!", coin)
			continue
		}
		coinPrice[coin.PriceMarketName] = price
		coinRank[coin.PriceMarketName] = 0
	}
	return coinPrice, coinRank, nil
}
