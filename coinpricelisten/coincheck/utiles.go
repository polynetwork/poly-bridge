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

package coincheck

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"io/ioutil"
	"math/big"
	"net/http"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/models"
	"reflect"
)

type CoincheckSdk struct {
	client *http.Client
	nodes  []*conf.Restful
}

func NewCoincheckSdk(cfg *conf.CoinPriceListenConfig) *CoincheckSdk {
	client := &http.Client{}
	sdk := &CoincheckSdk{
		client: client,
		nodes:  cfg.Nodes,
	}
	return sdk
}

func (c *CoincheckSdk) QuotesLatest() (*Rate, error) {
	for i := 0; i < len(c.nodes); i++ {
		quotes, err := c.quotesLatest(i)
		if err != nil {
			logs.Error("Coincheck QuotesLatest err: %s", err.Error())
			continue
		} else {
			return quotes, nil
		}
	}
	return nil, fmt.Errorf("Cannot get Coincheck QuotesLatest!")
}

func (c *CoincheckSdk) quotesLatest(node int) (*Rate, error) {
	req, err := http.NewRequest("GET", c.nodes[node].Url+"api/rate/all", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accepts", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response status code: %d", resp.StatusCode)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	rates := new(Rate)
	err = json.Unmarshal(respBody, rates)
	if err != nil {
		return nil, err
	}
	return rates, nil
}

func (c *CoincheckSdk) GetMarketName() string {
	return basedef.MARKET_COINCHECK
}

func (c *CoincheckSdk) GetCoinPriceAndRank(coins []models.NameAndmarketId) (map[string]float64, map[string]int, error) {
	rates, err := c.QuotesLatest()
	if err != nil {
		return nil, nil, err
	}

	var (
		supportedCoinsMap = make(map[string]struct{}, 0)
		coinPrice         = make(map[string]float64, 0)
		coinRink          = make(map[string]int, 0)
		typeOfRates       = reflect.TypeOf(rates.Jpy)
		valueOfRate       = reflect.ValueOf(rates.Jpy)
		usdJpyRate        = big.NewFloat(rates.Jpy.USD)
	)

	for i := 0; i < typeOfRates.NumField(); i++ {
		fieldType := typeOfRates.Field(i)
		supportedCoinsMap[fieldType.Name] = struct{}{}
	}
	for _, coin := range coins {
		if _, ok := supportedCoinsMap[coin.PriceMarketName]; !ok {
			logs.Warn("%s price is not available in Coincheck!", coin)
			continue
		}
		coinJpyRate := valueOfRate.FieldByName(coin.PriceMarketName).Float()
		coinPrice[coin.PriceMarketName], _ = new(big.Float).Quo(big.NewFloat(coinJpyRate), usdJpyRate).Float64()
		coinRink[coin.PriceMarketName] = 0
	}
	return coinPrice, coinRink, nil
}
