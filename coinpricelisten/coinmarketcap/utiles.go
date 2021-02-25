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
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"net/url"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"strings"
)

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
	req, err := http.NewRequest("GET", sdk.nodes[node].Url+"listings/latest", nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "5000")
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
	for i := 0; i < len(sdk.nodes); i++ {
		quotes, err := sdk.quotesLatest(coins, i)
		if err != nil {
			logs.Error("CoinMarketCap QuotesLatest err: %s", err.Error())
			continue
		} else {
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

func (sdk *CoinMarketCapSdk) GetCoinPrice(coins []string) (map[string]float64, error) {
	listings, err := sdk.ListingsLatest()
	if err != nil {
		return nil, err
	}
	//
	coinName2Id := make(map[string]string, 0)
	for _, listing := range listings {
		coinName2Id[listing.Name] = fmt.Sprintf("%d", listing.ID)
	}
	//
	coinIds := make([]string, 0)
	for _, coin := range coins {
		coinId, ok := coinName2Id[coin]
		if !ok {
			logs.Warn("There is no coin %s in CoinMarketCap!", coin)
			continue
		}
		coinIds = append(coinIds, coinId)
	}
	//
	requestCoinIds := strings.Join(coinIds, ",")
	quotes, err := sdk.QuotesLatest(requestCoinIds)
	if err != nil {
		return nil, err
	}
	//
	coinName2Price := make(map[string]float64)
	for _, v := range quotes {
		name := v.Name
		if v.Quote == nil || v.Quote["USD"] == nil {
			logs.Warn(" There is no price for coin %s in CoinMarketCap!", name)
			continue
		}
		coinName2Price[name] = v.Quote["USD"].Price
	}
	return coinName2Price, nil
}
