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
	"github.com/astaxie/beego/logs"
	"github.com/joeqian10/neo-gogogo/rpc"
	"github.com/joeqian10/neo-gogogo/rpc/models"
)

type NeoSdk struct {
	client *rpc.RpcClient
	url   string
}

func NewNeoSdk(url string) *NeoSdk {
	return &NeoSdk{
		client: rpc.NewClient(url),
		url:   url,
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
	cur := sdk.node
	res := sdk.client.GetApplicationLog(txId)
	for res.ErrorResponse.Error.Message != "" {
		logs.Error("NeoClient.GetApplicationLog err:%s, url: %s", res.ErrorResponse.Error.Message, sdk.urls[sdk.node])
		next := sdk.NextClient()
		if next == cur {
			return nil, fmt.Errorf("all node is not working!")
		}
		res = sdk.client.GetApplicationLog(txId)
	}
	return &res.Result, nil
}
