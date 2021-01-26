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
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/nep5"
	"github.com/joeqian10/neo-gogogo/rpc"
	"github.com/joeqian10/neo-gogogo/rpc/models"
)

type NeoSdk struct {
	client *rpc.RpcClient
	url    string
}

func NewNeoSdk(url string) *NeoSdk {
	return &NeoSdk{
		client: rpc.NewClient(url),
		url:    url,
	}
}

func (sdk *NeoSdk) GetBlockCount() (uint64, error) {
	res := sdk.client.GetBlockCount()
	if res.ErrorResponse.Error.Message != "" {
		return 0, fmt.Errorf("%s", res.ErrorResponse.Error.Message)
	}
	return uint64(res.Result), nil
}

func (sdk *NeoSdk) GetBlockByIndex(index uint64) (*models.RpcBlock, error) {
	res := sdk.client.GetBlockByIndex(uint32(index))
	if res.ErrorResponse.Error.Message != "" {
		return nil, fmt.Errorf("%s", res.ErrorResponse.Error.Message)
	}
	return &res.Result, nil
}

func (sdk *NeoSdk) GetApplicationLog(txId string) (*models.RpcApplicationLog, error) {
	res := sdk.client.GetApplicationLog(txId)
	if res.ErrorResponse.Error.Message != "" {
		return nil, fmt.Errorf("%s", res.ErrorResponse.Error.Message)
	}
	return &res.Result, nil
}

func (sdk *NeoSdk) GetTransactionHeight(hash string) (uint64, error) {
	res := sdk.client.GetTransactionHeight(hash)
	if res.ErrorResponse.Error.Message != "" {
		return 0, fmt.Errorf("%s", res.ErrorResponse.Error.Message)
	}
	return uint64(res.Result), nil
}

func (sdk *NeoSdk) SendRawTransaction(txHex string) (bool, error) {
	res := sdk.client.SendRawTransaction(txHex)
	if res.HasError() {
		return false, fmt.Errorf("%s", res.ErrorResponse.Error.Message)
	}
	return res.Result, nil
}

func (sdk *NeoSdk) Nep5Info(hash string) (string, string, int64, error) {
	scriptHash, err := helper.UInt160FromString(hash)
	if err != nil {
		return "", "", 0, err
	}
	nep5 := nep5.NewNep5Helper(scriptHash, sdk.url)
	decimal, err := nep5.Decimals()
	if err != nil {
		return "", "", 0, err
	}
	name, err := nep5.Name()
	if err != nil {
		return "", "", 0, err
	}
	return hash, name, int64(decimal), nil
}
