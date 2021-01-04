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
	"github.com/polynetwork/poly-go-sdk"
	"github.com/polynetwork/poly-go-sdk/common"
	"github.com/polynetwork/poly/core/types"
)

type PolySDK struct {
	sdk *poly_go_sdk.PolySdk
	url string
}

func NewPolySDK(url string) *PolySDK {
	rawsdk := poly_go_sdk.NewPolySdk()
	rawsdk.NewRpcClient().SetAddress(url)
	return &PolySDK{
		sdk: rawsdk,
		url: url,
	}
}

func (sdk *PolySDK) GetCurrentBlockHeight() (uint64, error) {
	height, err := sdk.sdk.GetCurrentBlockHeight()
	return uint64(height), err
}

func (sdk *PolySDK) GetBlockByHeight(height uint64) (*types.Block, error) {
	block, err := sdk.sdk.GetBlockByHeight(uint32(height))
	return block, err
}

func (sdk *PolySDK) GetSmartContractEventByBlock(height uint64) ([]*common.SmartContactEvent, error) {
	event, err := sdk.sdk.GetSmartContractEventByBlock(uint32(height))
	return event, err
}
