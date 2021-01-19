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

package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	POLY_CROSSCHAIN_ID     = uint64(0)
	ETHEREUM_CROSSCHAIN_ID = uint64(2)
	NEO_CROSSCHAIN_ID      = uint64(5)
	BSC_CROSSCHAIN_ID      = uint64(79)
	HECO_CROSSCHAIN_ID      = uint64(7)
)

var (
	PRICE_PRECISION = int64(100000000)
	FEE_PRECISION   = int64(100000000)
)

var (
	MARKET_COINMARKETCAP = "coinmarketcap"
	MARKET_BINANCE       = "binance"
	MARKET_HUOBI         = "huobi"
)

const (
	STATE_FINISHED = iota
	STATE_PENDDING
	STATE_SOURCE_DONE
	STATE_SOURCE_CONFIRMED
	STATE_POLY_CONFIRMED
)

const (
	SERVER_POLY_SWAP = "polyswap"
	SERVER_EXPLORER  = "explorer"
	SERVER_ADDRESS   = "address"
	SERVER_STAKE     = "stake"
)

func StateCode2Name(code int) string {
	if code == STATE_FINISHED {
		return "finished"
	} else if code == STATE_PENDDING {
		return "pendding"
	} else if code == STATE_SOURCE_DONE {
		return "source done"
	} else if code == STATE_SOURCE_CONFIRMED {
		return "source confirmed"
	} else if code == STATE_POLY_CONFIRMED {
		return "poly confirmed"
	} else {
		return "unknown"
	}
}

func StateName2Code(state string) int {
	if state == "finished" {
		return STATE_FINISHED
	} else if state == "pendding" {
		return STATE_PENDDING
	} else if state == "source done" {
		return STATE_SOURCE_DONE
	} else if state == "source confirmed" {
		return STATE_SOURCE_CONFIRMED
	} else if state == "poly confirmed" {
		return STATE_POLY_CONFIRMED
	} else {
		return -1
	}
}

type DBConfig struct {
	URL      string
	User     string
	Password string
	Scheme   string
}

type Restful struct {
	Url string
	Key string
}

type ChainListenConfig struct {
	ChainName           string
	ChainId             uint64
	ListenSlot          uint64
	BackwardBlockNumber uint64
	Nodes               []*Restful
	ExtendNodes         []*Restful
	WrapperContract     string
	CCMContract         string
	ProxyContract       string
}

func (cfg *ChainListenConfig) GetNodesUrl() []string {
	urls := make([]string, 0)
	for _, node := range cfg.Nodes {
		urls = append(urls, node.Url)
	}
	return urls
}

func (cfg *ChainListenConfig) GetNodesKey() []string {
	keys := make([]string, 0)
	for _, node := range cfg.Nodes {
		keys = append(keys, node.Key)
	}
	return keys
}

func (cfg *ChainListenConfig) GetExtendNodesUrl() []string {
	urls := make([]string, 0)
	for _, node := range cfg.ExtendNodes {
		urls = append(urls, node.Url)
	}
	return urls
}

func (cfg *ChainListenConfig) GetExtendNodesKey() []string {
	keys := make([]string, 0)
	for _, node := range cfg.ExtendNodes {
		keys = append(keys, node.Key)
	}
	return keys
}

type CoinPriceListenConfig struct {
	MarketName string
	Nodes      []*Restful
}

func (cfg *CoinPriceListenConfig) GetNodesUrl() []string {
	urls := make([]string, 0)
	for _, node := range cfg.Nodes {
		urls = append(urls, node.Url)
	}
	return urls
}

func (cfg *CoinPriceListenConfig) GetNodesKey() []string {
	keys := make([]string, 0)
	for _, node := range cfg.Nodes {
		keys = append(keys, node.Key)
	}
	return keys
}

type FeeListenConfig struct {
	ChainId   uint64
	ChainName string
	Nodes     []*Restful
	ProxyFee  int64
	MinFee    int64
	GasLimit  int64
}

func (cfg *FeeListenConfig) GetNodesUrl() []string {
	urls := make([]string, 0)
	for _, node := range cfg.Nodes {
		urls = append(urls, node.Url)
	}
	return urls
}

func (cfg *FeeListenConfig) GetNodesKey() []string {
	keys := make([]string, 0)
	for _, node := range cfg.Nodes {
		keys = append(keys, node.Key)
	}
	return keys
}

type EventEffectConfig struct {
	HowOld int64
	Server string
}

type Config struct {
	Server                string
	ChainListenConfig     []*ChainListenConfig
	CoinPriceUpdateSlot   int64
	CoinPriceListenConfig []*CoinPriceListenConfig
	FeeUpdateSlot         int64
	FeeListenConfig       []*FeeListenConfig
	EventEffectConfig     *EventEffectConfig
	DBConfig              *DBConfig
}

func (cfg *Config) GetChainListenConfig(chainId uint64) *ChainListenConfig {
	for _, chainListenConfig := range cfg.ChainListenConfig {
		if chainListenConfig.ChainId == chainId {
			return chainListenConfig
		}
	}
	return nil
}

func (cfg *Config) GetCoinPriceListenConfig(market string) *CoinPriceListenConfig {
	for _, coinPriceListenConfig := range cfg.CoinPriceListenConfig {
		if coinPriceListenConfig.MarketName == market {
			return coinPriceListenConfig
		}
	}
	return nil
}

func (cfg *Config) GetFeeListenConfig(chainId uint64) *FeeListenConfig {
	for _, feeListenConfig := range cfg.FeeListenConfig {
		if feeListenConfig.ChainId == chainId {
			return feeListenConfig
		}
	}
	return nil
}

func ReadFile(fileName string) ([]byte, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("ReadFile: open file %s error %s", fileName, err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Errorf("ReadFile: File %s close error %s", fileName, err)
		}
	}()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ReadFile: ioutil.ReadAll %s error %s", fileName, err)
	}
	return data, nil
}

func NewConfig(filePath string) *Config {
	fileContent, err := ReadFile(filePath)
	if err != nil {
		fmt.Errorf("NewServiceConfig: failed, err: %s", err)
		return nil
	}
	config := &Config{}
	err = json.Unmarshal(fileContent, config)
	if err != nil {
		fmt.Errorf("NewServiceConfig: failed, err: %s", err)
		return nil
	}
	return config
}
