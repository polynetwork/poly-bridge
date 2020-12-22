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

package neolisten

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/joeqian10/neo-gogogo/rpc"
	"github.com/joeqian10/neo-gogogo/rpc/models"
)

type NeoSdk struct {
	client  *rpc.RpcClient
	urls     []string
	node     int
}

func NewNeoSdk(urls []string) *NeoSdk {
	return &NeoSdk{
		client: rpc.NewClient(urls[0]),
		urls: urls,
		node: 0,
	}
}

func (sdk *NeoSdk) NextClient() int {
	sdk.node++
	sdk.node = sdk.node % len(sdk.urls)
	sdk.client = rpc.NewClient(sdk.urls[sdk.node])
	return sdk.node
}

func (sdk *NeoSdk) GetBlockCount() (uint64, error) {
	cur := sdk.node
	res := sdk.client.GetBlockCount()
	for res.ErrorResponse.Error.Message != "" {
		logs.Error("NeoClient.GetBlockCount err:%s, url: %s", res.ErrorResponse.Error.Message, sdk.urls[sdk.node])
		next := sdk.NextClient()
		if next == cur {
			return 0, fmt.Errorf("all node is not working!")
		}
		res = sdk.client.GetBlockCount()
	}
	return uint64(res.Result), nil
}

func (sdk *NeoSdk) GetBlockByIndex(index uint64) (*models.RpcBlock, error) {
	cur := sdk.node
	res := sdk.client.GetBlockByIndex(uint32(index))
	for res.ErrorResponse.Error.Message != "" {
		logs.Error("NeoClient.GetBlockByIndex err:%s, url: %s", res.ErrorResponse.Error.Message, sdk.urls[sdk.node])
		next := sdk.NextClient()
		if next == cur {
			return nil, fmt.Errorf("all node is not working!")
		}
		res = sdk.client.GetBlockByIndex(uint32(index))
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
