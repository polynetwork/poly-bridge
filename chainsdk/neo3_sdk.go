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

package chainsdk

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/nep17"
	"github.com/joeqian10/neo3-gogogo/rpc"
	"github.com/joeqian10/neo3-gogogo/rpc/models"
	"math/big"
	"strconv"
)

type Neo3Sdk struct {
	client *rpc.RpcClient
	url    string
}

func NewNeo3Sdk(url string) *Neo3Sdk {
	return &Neo3Sdk{
		client: rpc.NewClient(url),
		url:    url,
	}
}

func (sdk *Neo3Sdk) GetClient() *rpc.RpcClient {
	return sdk.client
}

func (sdk *Neo3Sdk) GetUrl() string {
	return sdk.url
}

func (sdk *Neo3Sdk) GetBlockCount() (uint64, error) {
	res := sdk.client.GetBlockCount()
	if res.ErrorResponse.Error.Message != "" {
		return 0, fmt.Errorf("%s", res.ErrorResponse.Error.Message)
	}
	return uint64(res.Result), nil
}

func (sdk *Neo3Sdk) GetBlockByIndex(index uint64) (*models.RpcBlock, error) {
	res := sdk.client.GetBlock(strconv.Itoa(int(index)))
	if res.ErrorResponse.Error.Message != "" {
		return nil, fmt.Errorf("%s", res.ErrorResponse.Error.Message)
	}
	return &res.Result, nil
}

func (sdk *Neo3Sdk) GetApplicationLog(txId string) (*models.RpcApplicationLog, error) {
	res := sdk.client.GetApplicationLog(txId)
	if res.ErrorResponse.Error.Message != "" {
		return nil, fmt.Errorf("%s", res.ErrorResponse.Error.Message)
	}
	return &res.Result, nil
}

func (sdk *Neo3Sdk) GetTransactionHeight(hash string) (uint64, error) {
	res := sdk.client.GetTransactionHeight(hash)
	if res.ErrorResponse.Error.Message != "" {
		return 0, fmt.Errorf("%s", res.ErrorResponse.Error.Message)
	}
	return uint64(res.Result), nil
}

func (sdk *Neo3Sdk) SendRawTransaction(txHex string) (bool, error) {
	res := sdk.client.SendRawTransaction(txHex)
	if res.HasError() {
		return false, fmt.Errorf("%s", res.ErrorResponse.Error.Message)
	}
	return true, nil
}

func (sdk *Neo3Sdk) Nep17Info(hash string) (string, string, int64, error) {
	scriptHash, err := helper.UInt160FromString(hash)
	if err != nil {
		return "", "", 0, err
	}
	nep17 := nep17.NewNep17Helper(scriptHash, sdk.client)
	decimal, err := nep17.Decimals()
	if err != nil {
		return "", "", 0, err
	}
	symbol, err := nep17.Symbol()
	if err != nil {
		return "", "", 0, err
	}
	return hash, symbol, int64(decimal), nil
}

func (sdk *Neo3Sdk) Nep17Balance(hash string, addr string) (*big.Int, error) {
	scriptHash, err := helper.UInt160FromString(hash)
	if err != nil {
		return new(big.Int).SetUint64(0), err
	}
	nep17 := nep17.NewNep17Helper(scriptHash, sdk.client)
	addrHash, err := helper.UInt160FromString(addr)
	if err != nil {
		return new(big.Int).SetUint64(0), err
	}
	return nep17.BalanceOf(addrHash)
}

func (sdk *Neo3Sdk) Nep17TotalSupply(hash string) (*big.Int, error) {
	scriptHash, err := helper.UInt160FromString(hash)
	if err != nil {
		return new(big.Int).SetUint64(0), err
	}
	logs.Info("hash: %s", hash)
	nep17 := nep17.NewNep17Helper(scriptHash, sdk.client)
	if err != nil {
		return new(big.Int).SetUint64(0), err
	}
	return nep17.TotalSupply()
}
