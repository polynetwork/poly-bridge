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

package coinmarketcap

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"io/ioutil"
	"net/http"
	"net/url"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/models"
	"strings"
)

var apiKeyCount = 0

type CoinMarketCapSdk struct {
	client *http.Client
	nodes  []*conf.Restful
}

func DefaultCoinMarketCapSdk() *CoinMarketCapSdk {
	client := &http.Client{}
	sdk := &CoinMarketCapSdk{
		client: client,
		nodes: []*conf.Restful{
			{
				Url: "https://pro-api.coinmarketcap.com/v1/cryptocurrency/",
				Key: "8efe5156-8b37-4c77-8e1d-a140c97bf466",
			},
		},
	}
	return sdk
}

func NewCoinMarketCapSdk(cfg *conf.CoinPriceListenConfig) *CoinMarketCapSdk {
	client := &http.Client{}
	sdk := &CoinMarketCapSdk{
		client: client,
		nodes:  cfg.Nodes,
	}
	return sdk
}

type ListingsMedia struct {
	Data []*Listing `json:"data"`
}

func (sdk *CoinMarketCapSdk) ListingsLatest() ([]*Listing, error) {
	for i := 0; i < len(sdk.nodes); i++ {
		listings, err := sdk.listingsLatest(i)
		if err != nil {
			logs.Error("CoinMarketCap ListingsLatest err: %s", err.Error())
			continue
		} else {
			return listings, nil
		}
	}
	return nil, fmt.Errorf("Cannot get CoinMarketCap ListingsLatest!")
}

func (sdk *CoinMarketCapSdk) listingsLatest(node int) ([]*Listing, error) {
	allListing := make([]*Listing, 0)
	aa, err := sdk.listingsLatest1(node, 1, 5000)
	if err != nil {
		return nil, err
	}
	allListing = append(allListing, aa...)
	bb, err := sdk.listingsLatest1(node, 5000, 5000)
	if err != nil {
		return nil, err
	}
	allListing = append(allListing, bb...)
	return allListing, nil
}

func (sdk *CoinMarketCapSdk) listingsLatest1(node int, start int, limit int) ([]*Listing, error) {
	req, err := http.NewRequest("GET", sdk.nodes[node].Url+"listings/latest", nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("start", fmt.Sprintf("%d", start))
	q.Add("limit", fmt.Sprintf("%d", limit))
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", sdk.nodes[node].Key)
	req.URL.RawQuery = q.Encode()

	resp, err := sdk.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response status code: %d", resp.StatusCode)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	var body ListingsMedia
	err = json.Unmarshal(respBody, &body)
	if err != nil {
		return nil, err
	}
	return body.Data, nil
}

type QuotesLatestMedia struct {
	Data map[string]*Ticker `json:"data"`
}

func (sdk *CoinMarketCapSdk) QuotesLatest(coins string) (map[string]*Ticker, error) {
	if apiKeyCount == len(sdk.nodes) {
		apiKeyCount = 0
	}
	for i := apiKeyCount; i < len(sdk.nodes); i++ {
		quotes, err := sdk.quotesLatest(coins, i)
		if err != nil {
			logs.Error("CoinMarketCap QuotesLatest err: %s, apiKey: %s", err.Error(), sdk.nodes[i].Key)
			continue
		} else {
			apiKeyCount = i + 1
			return quotes, nil
		}
	}
	return nil, fmt.Errorf("Cannot get CoinMarketCap QuotesLatest!")
}

func (sdk *CoinMarketCapSdk) quotesLatest(coins string, node int) (map[string]*Ticker, error) {
	req, err := http.NewRequest("GET", sdk.nodes[node].Url+"quotes/latest", nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("id", coins)
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", sdk.nodes[node].Key)
	req.URL.RawQuery = q.Encode()

	resp, err := sdk.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response status code: %d", resp.StatusCode)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	var body QuotesLatestMedia
	err = json.Unmarshal(respBody, &body)
	if err != nil {
		return nil, err
	}
	return body.Data, nil
}

func (sdk *CoinMarketCapSdk) GetMarketName() string {
	return basedef.MARKET_COINMARKETCAP
}

func (sdk *CoinMarketCapSdk) GetCoinPriceAndRank(coins []models.NameAndmarketId) (map[string]float64, map[string]int, error) {
	coinIds := make([]string, 0)
	coinInId := make(map[int]bool, 0)
	for _, coin := range coins {
		if coin.CoinMarketId <= 0 {
			continue
		}
		if _, ok := coinInId[coin.CoinMarketId]; ok {
			continue
		}
		coinInId[coin.CoinMarketId] = true
		coinIds = append(coinIds, fmt.Sprintf("%d", coin.CoinMarketId))
	}
	//
	requestCoinIds := strings.Join(coinIds, ",")
	quotes, err := sdk.QuotesLatest(requestCoinIds)
	if err != nil {
		return nil, nil, err
	}
	coinId2Price := make(map[int]float64)
	coinId2Rank := make(map[int]int)
	for _, v := range quotes {
		name := v.Name
		if v.Quote == nil || v.Quote["USD"] == nil {
			logs.Warn(" There is no price for coin %s in CoinMarketCap!", name)
			continue
		}
		coinId2Price[v.ID] = v.Quote["USD"].Price
		coinId2Rank[v.ID] = v.Rank
	}

	coinName2Price := make(map[string]float64)
	coinName2Rank := make(map[string]int)
	for _, coin := range coins {
		if coin.CoinMarketId <= 0 {
			continue
		}
		if price, ok := coinId2Price[coin.CoinMarketId]; ok {
			coinName2Price[coin.PriceMarketName] = price
			coinName2Rank[coin.PriceMarketName] = coinId2Rank[coin.CoinMarketId]
		}
	}
	return coinName2Price, coinName2Rank, nil
}
