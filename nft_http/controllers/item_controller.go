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
	"math/big"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum/common"
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
	wrapper := selectWrapper(req.ChainId)
	if wrapper == emptyAddr {
		customInput(&c.Controller, ErrCodeRequest, "chain id not exist")
		return
	}

	owner := common.HexToAddress(req.Address)
	asset := common.HexToAddress(req.Asset)
	totalCnt, err := sdk.NFTBalance(asset, owner)
	if err != nil {
		logs.Error("get nft balance err: %v", err)
		nodeInvalid(&c.Controller)
		return
	}

	totalPage := getPageNo(totalCnt, req.PageSize)
	start := req.PageNo * req.PageSize
	empty := func() {
		data := new(ItemsOfAddressRsp).instance(req.PageSize, req.PageNo, totalPage, totalCnt, []*Item{})
		output(&c.Controller, data)
	}

	tokenIdStr := strings.Trim(req.TokenId, " ")
	if tokenIdStr != "" {
		tokenId, ok := new(big.Int).SetString(tokenIdStr, 10)
		if !ok {
			input(&c.Controller, req)
			return
		}
		url, err := sdk.GetAndCheckTokenUrl(wrapper, asset, owner, tokenId)
		if err != nil {
			empty()
			return
		}
		item := &Item{TokenId: tokenId.String(), Url: url}
		data := new(ItemsOfAddressRsp).instance(req.PageSize, req.PageNo, totalPage, totalCnt, []*Item{item})
		output(&c.Controller, data)
		return
	}

	if start >= totalCnt {
		empty()
		return
	}
	res, err := sdk.GetTokensByIndex(wrapper, asset, owner, start, req.PageSize)
	if err != nil {
		logs.Error("batch get NFT token infos err: %v", err)
		empty()
		return
	}
	if len(res) == 0 {
		empty()
		return
	}

	items := make([]*Item, 0)
	for tokenId, url := range res {
		items = append(items, &Item{
			TokenId: tokenId.String(),
			Url:     url,
		})
	}

	data := new(ItemsOfAddressRsp).instance(req.PageSize, req.PageNo, totalPage, totalCnt, items)
	output(&c.Controller, data)
}
