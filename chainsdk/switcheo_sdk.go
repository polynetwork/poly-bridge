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
	"context"

	"github.com/beego/beego/v2/core/logs"
	tmclient "github.com/tendermint/tendermint/rpc/client/http"
	ctypes "github.com/tendermint/tendermint/rpc/coretypes"
)

type SwitcheoSDK struct {
	client *tmclient.HTTP
	url    string
}

func NewSwitcheoSDK(url string) *SwitcheoSDK {
	rawClient, err := tmclient.New(url)
	if err != nil {
		panic(err)
	}
	return &SwitcheoSDK{
		client: rawClient,
		url:    url,
	}
}

func (client *SwitcheoSDK) Status() (*ctypes.ResultStatus, error) {
	return client.client.Status(context.Background())
}

func (client *SwitcheoSDK) Block(height *int64) (*ctypes.ResultBlock, error) {
	return client.client.Block(context.Background(), height)
}

func (client *SwitcheoSDK) TxSearch(query string, prove bool, page, perPage int, orderBy string) (*ctypes.ResultTxSearch, error) {
	return client.client.TxSearch(context.Background(), query, prove, &page, &perPage, orderBy)
}

func (sdk *SwitcheoSDK) GetCurrentBlockHeight() (uint64, error) {
	status, err := sdk.Status()
	if err != nil {
		logs.Error("GetLatestBlockHeight switcheo: get current block status， err: %v", err)
		return 0, err
	}
	res := status.SyncInfo.LatestBlockHeight
	return uint64(res), nil
}
