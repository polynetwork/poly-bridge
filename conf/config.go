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
	POLY_CHAIN_NAME        = "Poly"
	ETHEREUM_CROSSCHAIN_ID = uint64(2)
	ETHEREUM_CHAIN_NAME    = "Ethereum"
	NEO_CROSSCHAIN_ID      = uint64(4)
	NEO_CHAIN_NAME         = "NEO"
	BSC_CROSSCHAIN_ID      = uint64(8)
	BSC_CHAIN_NAME         = "BSC"
)

var (
	PRICE_PRECISION = int64(100000000)
	FEE_PRECISION   = int64(100000000)
)

const (
	FINISHED = iota
	PENDDING
	SOURCE_DONE
	SOURCE_CONFIRMED
	POLY_CONFIRMED
	TAGET_CONFIRMED
)

const (
	SERVER_POLY_SWAP = "polyswap"
	SERVER_EXPLORER  = "explorer"
	SERVER_ADDRESS   = "address"
)

type DBConfig struct {
	URL      string
	User     string
	Password string
	Scheme   string
}

type ChainListenConfig struct {
	PolyChainListenConfig     *PolyChainListenConfig
	EthereumChainListenConfig *EthereumChainListenConfig
	NeoChainListenConfig      *NeoChainListenConfig
	BscChainListenConfig      *BscChainListenConfig
}

type EthereumChainListenConfig struct {
	ChainName           string
	ChainId             uint64
	ListenSlot          uint64
	BackwardBlockNumber uint64
	RestURL             []string
	ExtendNodeURL       string
	ExtendNodeApiKey    string
	WrapperContract     string
	CCMContract         string
	ProxyContract       string
}

type NeoChainListenConfig struct {
	ChainName           string
	ChainId             uint64
	ListenSlot          uint64
	BackwardBlockNumber uint64
	RestURL             []string
	ExtendNodeURL       string
	WrapperContract     string
	ProxyContract       string
}

type BscChainListenConfig struct {
	ChainName           string
	ChainId             uint64
	ListenSlot          uint64
	BackwardBlockNumber uint64
	RestURL             []string
	ExtendNodeURL       string
	ExtendNodeApiKey    string
	WrapperContract     string
	CCMContract         string
	ProxyContract       string
}

type PolyChainListenConfig struct {
	ChainName           string
	ChainId             uint64
	ListenSlot          uint64
	BackwardBlockNumber uint64
	RestURL             []string
	Contract            string
}

type CoinPriceListenConfig struct {
	PriceUpdateSlot                int64
	CoinMarketCapPriceListenConfig *CoinMarketCapPriceListenConfig
	BinancePriceListenConfig       *BinancePriceListenConfig
}

type CoinMarketCapPriceListenConfig struct {
	RestURL []string
	Key     []string
}

type BinancePriceListenConfig struct {
	RestURL []string
}

type FeeListenConfig struct {
	FeeUpdateSlot           int64
	EthereumFeeListenConfig *EthereumFeeListenConfig
	NeoFeeListenConfig      *NeoFeeListenConfig
	BscFeeListenConfig      *BscFeeListenConfig
}

type EthereumFeeListenConfig struct {
	RestURL  []string
	ProxyFee uint64
	GasLimit int64
}

type NeoFeeListenConfig struct {
	RestURL  []string
	ProxyFee uint64
}

type BscFeeListenConfig struct {
	RestURL  []string
	ProxyFee uint64
	GasLimit int64
}

type TxStatusListenConfig struct {
	UpdateSlot        int64
	RestURL           []string
	EthereumConfirmed uint64
	NeoConfirmed      uint64
	BscConfirmed      uint64
}

type Config struct {
	Server                string
	ChainListenConfig     *ChainListenConfig
	CoinPriceListenConfig *CoinPriceListenConfig
	FeeListenConfig       *FeeListenConfig
	TxStatusListenConfig  *TxStatusListenConfig
	DBConfig              *DBConfig
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
