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
	sdk "github.com/polynetwork/poly-go-sdk"
	"github.com/polynetwork/poly-go-sdk/common"
	"github.com/polynetwork/poly/core/types"
)

type PolySDK struct {
	sdk  *sdk.PolySdk
	urls []string
	node int
}

func NewPolySDK(urls []string) *PolySDK {
	rawsdk := sdk.NewPolySdk()
	rawsdk.NewRpcClient().SetAddress(urls[0])
	return &PolySDK{
		sdk:  rawsdk,
		urls: urls,
		node: 0,
	}
}

func (ps *PolySDK) NextClient() int {
	ps.node++
	ps.node = ps.node % len(ps.urls)
	ps.sdk = sdk.NewPolySdk()
	ps.sdk.NewRpcClient().SetAddress(ps.urls[ps.node])
	return ps.node
}

func (ps *PolySDK) GetLatestHeight() (uint64, error) {
	latestHeight := uint32(0)
	for i, url := range ps.urls {
		sdk := sdk.NewPolySdk()
		sdk.NewRpcClient().SetAddress(url)
		height, err := sdk.GetCurrentBlockHeight()
		if err != nil {
			continue
		}
		if height > latestHeight {
			ps.node = i
			latestHeight = height
		}
	}
	ps.sdk = sdk.NewPolySdk()
	ps.sdk.NewRpcClient().SetAddress(ps.urls[ps.node])
	return uint64(latestHeight), nil
}

func (client *PolySDK) GetCurrentBlockHeight() (uint64, error) {
	cur := client.node
	height, err := client.sdk.GetCurrentBlockHeight()
	for err != nil {
		logs.Error("PolySDK.GetCurrentBlockHeight err:%s, url: %s", err.Error(), client.urls[client.node])
		next := client.NextClient()
		if next == cur {
			return 0, fmt.Errorf("all node is not working!")
		}
		height, err = client.sdk.GetCurrentBlockHeight()
	}
	return uint64(height), err
}

func (client *PolySDK) GetBlockByHeight(height uint64) (*types.Block, error) {
	cur := client.node
	block, err := client.sdk.GetBlockByHeight(uint32(height))
	for err != nil {
		logs.Error("PolySDK.GetBlockByHeight err:%s, url: %s", err.Error(), client.urls[client.node])
		next := client.NextClient()
		if next == cur {
			return nil, fmt.Errorf("all node is not working!")
		}
		block, err = client.sdk.GetBlockByHeight(uint32(height))
	}
	return block, err
}

func (client *PolySDK) GetSmartContractEventByBlock(height uint64) ([]*common.SmartContactEvent, error) {
	cur := client.node
	event, err := client.sdk.GetSmartContractEventByBlock(uint32(height))
	for err != nil {
		logs.Error("PolySDK.GetSmartContractEventByBlock err:%s, url: %s", err.Error(), client.urls[client.node])
		next := client.NextClient()
		if next == cur {
			return nil, fmt.Errorf("all node is not working!")
		}
		event, err = client.sdk.GetSmartContractEventByBlock(uint32(height))
	}
	return event, err
}
