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

package coinpricelisten

import (
	"github.com/beego/beego/v2/core/logs"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/coinpricedao"
	"poly-bridge/coinpricelisten/binance"
	"poly-bridge/coinpricelisten/coincheck"
	"poly-bridge/coinpricelisten/coinmarketcap"
	"poly-bridge/coinpricelisten/gateio"
	"poly-bridge/coinpricelisten/huobi"
	"poly-bridge/coinpricelisten/self"
	"poly-bridge/conf"
	"poly-bridge/models"
	"runtime/debug"
	"strings"
	"time"
)

var cpListen *CoinPriceListen

func StartCoinPriceListen(server string, priceUpdateSlot int64, coinPricecfg []*conf.CoinPriceListenConfig, dbCfg *conf.DBConfig) {
	dao := coinpricedao.NewCoinPriceDao(server, dbCfg)
	if dao == nil {
		panic("server is not valid")
	}
	priceMarkets := make([]PriceMarket, 0)
	for _, cfg := range coinPricecfg {
		priceMarket := NewPriceMarket(cfg)
		if priceMarket == nil {
			panic("price market is not valid")
		}
		priceMarkets = append(priceMarkets, priceMarket)
	}
	cpListen = NewCoinPriceListen(priceUpdateSlot, priceMarkets, dao)
	cpListen.Start()
}

func StopCoinPriceListen() {
	if cpListen != nil {
		cpListen.Stop()
	}
}

type PriceMarket interface {
	GetCoinPriceAndRank(coins []models.NameAndmarketId) (map[string]float64, map[string]int, error)
	GetMarketName() string
}

func NewPriceMarket(cfg *conf.CoinPriceListenConfig) PriceMarket {
	switch cfg.MarketName {
	case basedef.MARKET_COINMARKETCAP:
		return coinmarketcap.NewCoinMarketCapSdk(cfg)
	case basedef.MARKET_BINANCE:
		return binance.NewBinanceSdk(cfg)
	case basedef.MARKET_COINCHECK:
		return coincheck.NewCoincheckSdk(cfg)
	case basedef.MARKET_GATEIO:
		return gateio.NewGateioSdk()
	case basedef.MARKET_SELF:
		return self.NewSelfSdk(cfg)
	case basedef.MARKET_HUOBI:
		return huobi.NewHuobiSdk()
	default:
		return nil
	}
}

type CoinPriceListen struct {
	priceUpdateSlot int64
	priceMarket     map[string]PriceMarket
	db              coinpricedao.CoinPriceDao
	exit            chan bool
}

func NewCoinPriceListen(priceUpdateSlot int64, priceMarkets []PriceMarket, db coinpricedao.CoinPriceDao) *CoinPriceListen {
	cpListen := &CoinPriceListen{}
	cpListen.priceUpdateSlot = priceUpdateSlot
	cpListen.db = db
	cpListen.exit = make(chan bool, 0)
	cpListen.priceMarket = make(map[string]PriceMarket)
	for _, market := range priceMarkets {
		cpListen.priceMarket[market.GetMarketName()] = market
	}
	//
	tokenBasics, err := db.GetTokens()
	if err != nil {
		panic(err)
	}
	err = cpListen.updateCoinPrice(tokenBasics)
	if err != nil {
		panic(err)
	}
	err = db.SavePrices(tokenBasics)
	if err != nil {
		panic(err)
	}
	return cpListen
}

func (cpl *CoinPriceListen) RegisterPriceQuery(priceMarket PriceMarket) {
	cpl.priceMarket[priceMarket.GetMarketName()] = priceMarket
}

func (cpl *CoinPriceListen) Start() {
	logs.Info("start coin price listen.")
	go cpl.ListenPrice()
}

func (cpl *CoinPriceListen) Stop() {
	cpl.exit <- true
	logs.Info("stop coin price listen.")
}

func (cpl *CoinPriceListen) ListenPrice() {
	for {
		exit := cpl.listenPrice()
		if exit {
			close(cpl.exit)
			break
		}
		time.Sleep(time.Second * 5)
	}
}

func (cpl *CoinPriceListen) listenPrice() (exit bool) {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("service start, recover info: %s", string(debug.Stack()))
			exit = false
		}
	}()

	logs.Debug("coin price listen, market: %s, dao: %s......", cpl.GetPriceMarket(), cpl.db.Name())
	ticker := time.NewTicker(time.Minute)
	for {
		select {
		case <-ticker.C:
			now := time.Now().Unix() / 60
			if now%cpl.priceUpdateSlot != 0 {
				continue
			}
			counter := 0
			for counter < 5 {
				logs.Info("do price update at time: %s", time.Now().Format("2006-01-02 15:04:05"))
				time.Sleep(time.Second * 5)
				counter++
				tokenBasics, err := cpl.db.GetTokens()
				if err != nil {
					logs.Error("get token basic err: %v", err)
					continue
				}
				err = cpl.updateCoinPrice(tokenBasics)
				if err != nil {
					logs.Error("updateCoinPrice err: %v", err)
					continue
				}
				err = cpl.db.SavePrices(tokenBasics)
				if err != nil {
					logs.Error("save price err: %v", err)
					continue
				}
				break
			}
		case <-cpl.exit:
			logs.Info("coin price listen exit, market: %s, dao: %s......", cpl.GetPriceMarket(), cpl.db.Name())
			return true
		}
	}
}

func (cpl *CoinPriceListen) updateCoinPrice(tokenBasics []*models.TokenBasic) error {

	marketCoins := make(map[string][]models.NameAndmarketId)
	marketCoinPrices := make(map[string][]*models.PriceMarket)
	for _, tokenBasic := range tokenBasics {
		for _, priceMarket := range tokenBasic.PriceMarkets {
			_, ok := marketCoins[priceMarket.MarketName]
			if !ok {
				marketCoins[priceMarket.MarketName] = make([]models.NameAndmarketId, 0)
			}
			nameAndmarketId := models.NameAndmarketId{
				priceMarket.Name,
				priceMarket.CoinMarketId,
			}
			marketCoins[priceMarket.MarketName] = append(marketCoins[priceMarket.MarketName], nameAndmarketId)

			_, ok = marketCoinPrices[priceMarket.MarketName+priceMarket.Name]
			if !ok {
				marketCoinPrices[priceMarket.MarketName+priceMarket.Name] = make([]*models.PriceMarket, 0)
			}
			marketCoinPrices[priceMarket.MarketName+priceMarket.Name] = append(marketCoinPrices[priceMarket.MarketName+priceMarket.Name], priceMarket)
			priceMarket.Ind = 0
			tokenBasic.Ind = 0
		}
	}
	for market, query := range cpl.priceMarket {
		coins, ok := marketCoins[market]
		if !ok {
			logs.Error("there is no coins of market: %s", market)
			continue
		}
		coinPrices, coinRanks, err := query.GetCoinPriceAndRank(coins)
		if err != nil {
			logs.Error("get coin price of market: %s err: %v", market, err)
			continue
		}
		logs.Info("get coin price of market: %s successful", market)
		for name, price := range coinPrices {
			tokenPrices, ok := marketCoinPrices[market+name]
			if !ok {
				logs.Error("there is no coins of market: %s and token: %s", market, name)
				continue
			}
			price, _ := new(big.Float).Mul(big.NewFloat(price), big.NewFloat(float64(basedef.PRICE_PRECISION))).Int64()
			rank := coinRanks[name]
			for _, tokenPrice := range tokenPrices {
				tokenPrice.Price = price
				tokenPrice.Time = time.Now().Unix()
				tokenPrice.Ind = 1
				tokenPrice.Rank = rank
			}
		}
	}
	for _, tokenBasic := range tokenBasics {
		price := int64(0)
		rank := 0
		counter := int64(0)
		for _, tokenPrice := range tokenBasic.PriceMarkets {
			if tokenPrice.Ind == 1 {
				price += tokenPrice.Price
				counter++
			}
			if tokenPrice.Rank != 0 {
				rank = tokenPrice.Rank
			}
		}
		if counter > 0 {
			price = price / counter
			tokenBasic.Price = price
			tokenBasic.Rank = rank
			tokenBasic.Ind = 1
			tokenBasic.Time = time.Now().Unix()
		}
		if tokenBasic.Ind == 0 {
			logs.Error("Price of token %s is not update", tokenBasic.Name)
		}
	}
	return nil
}

func (cpl *CoinPriceListen) GetPriceMarket() string {
	priceMarkets := make([]string, 0)
	for _, priceMarket := range cpl.priceMarket {
		priceMarkets = append(priceMarkets, priceMarket.GetMarketName())
	}
	return strings.Join(priceMarkets, ",")
}
