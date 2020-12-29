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
)

type CoinMarketCapSdk struct {
	client *http.Client
	url    []string
	key    []string
}

func DefaultCoinMarketCapSdk() *CoinMarketCapSdk {
	//return NewCoinMarketCapSdk("https://api.coinmarketcap.com/v2")
	return NewCoinMarketCapSdk([]string{"https://pro-api.coinmarketcap.com/v1/cryptocurrency/"}, []string{"8efe5156-8b37-4c77-8e1d-a140c97bf466"})
}
func NewCoinMarketCapSdk(url []string, key []string) *CoinMarketCapSdk {
	client := &http.Client{}
	sdk := &CoinMarketCapSdk{
		client: client,
		url:    url,
		key:    key,
	}
	return sdk
}

type ListingsMedia struct {
	Data []*Listing `json:"data"`
}

func (sdk *CoinMarketCapSdk) ListingsLatest() ([]*Listing, error) {
	for i := 0; i < len(sdk.url); i++ {
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
	req, err := http.NewRequest("GET", sdk.url[node]+"listings/latest", nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "5000")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", sdk.key[node])
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
	for i := 0; i < len(sdk.url); i++ {
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
	req, err := http.NewRequest("GET", sdk.url[node]+"quotes/latest", nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("id", coins)
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", sdk.key[node])
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
