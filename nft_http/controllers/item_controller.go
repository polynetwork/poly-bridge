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

package controllers

import (
	"github.com/astaxie/beego"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

type ItemController struct {
	beego.Controller
}

// todo: cache url and token ids
func (c *ItemController) Items() {
	var req ItemsOfAddressReq
	if !input(&c.Controller, &req) {
		return
	}

	sdk := selectNode(req.ChainId)
	if sdk == nil {
		customInput(&c.Controller, ErrCodeRequest, "chain id not exist")
		return
	}

	owner := common.HexToAddress(req.Address)
	asset := common.HexToAddress(req.Asset)

	items := make([]*Item, 0)
	totalCnt, err := sdk.NFTBalance(asset, owner)
	if err != nil {
		nodeInvalid(&c.Controller)
		return
	}

	totalPage := getPageNo(totalCnt, req.PageSize)
	start := req.PageNo * req.PageSize
	end := start + req.PageSize

	empty := func() {
		data := new(ItemsOfAddressRsp).instance(req.PageSize, req.PageNo, totalPage, totalCnt, []*Item{})
		output(&c.Controller, data)
	}

	if start >= totalCnt {
		empty()
		return
	}
	if end > totalCnt {
		end = totalCnt
	}

	var list []*big.Int
	tokenIdStr := strings.Trim(req.TokenId, " ")
	if tokenIdStr != "" {
		tokenId, ok := new(big.Int).SetString(tokenIdStr, 10)
		if !ok {
			customInput(&c.Controller, ErrCodeRequest, "token id string invalid")
			return
		}
		list = []*big.Int{tokenId}
	} else {
		if list, err = sdk.GetNFTs(asset, owner, start, end); err != nil {
			nodeInvalid(&c.Controller)
			return
		}
	}
	if len(list) == 0 {
		empty()
		return
	}

	urlmap, err := sdk.GetNFTURLs(asset, list)
	if err != nil {
		nodeInvalid(&c.Controller)
		return
	}

	for _, v := range list {
		items = append(items, &Item{
			TokenId: v.Uint64(),
			Url:     urlmap[v.Uint64()],
		})
	}

	data := new(ItemsOfAddressRsp).instance(req.PageSize, req.PageNo, totalPage, totalCnt, items)
	output(&c.Controller, data)
}
