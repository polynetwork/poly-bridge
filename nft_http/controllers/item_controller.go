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
	"poly-bridge/models"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum/common"
	mcm "poly-bridge/nft_http/meta/common"
)

type ItemController struct {
	beego.Controller
}

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
	nftAsset := selectNFTAsset(req.Address)
	if nftAsset == nil {
		customInput(&c.Controller, ErrCodeRequest, "NFT Asset not exist")
		return
	}

	owner := common.HexToAddress(req.Address)
	asset := common.HexToAddress(req.Asset)
	bigTotalCnt, err := sdk.NFTBalance(asset, owner)
	if err != nil {
		logs.Error("get nft balance err: %v", err)
		nodeInvalid(&c.Controller)
		return
	}
	totalCnt := int(bigTotalCnt.Uint64())

	totalPage := getPageNo(totalCnt, req.PageSize)
	start := req.PageNo * req.PageSize
	empty := func() {
		data := new(ItemsOfAddressRsp).instance(req.PageSize, req.PageNo, totalPage, totalCnt, []*Item{})
		output(&c.Controller, data)
	}
	if totalCnt == 0 {
		empty()
		return
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
			logs.Error("getAndCheckTokenUrl err: %v", err)
			empty()
			return
		}
		profile, err := fetcher.Fetch(nftAsset.TokenBasicName, &mcm.FetchRequestParams{
			TokenId: models.NewBigInt(tokenId),
			Url:     url,
		})
		item := new(Item).instance(tokenId, profile)
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
		logs.Error("GetTokensByIndex err: %v", err)
		empty()
		return
	}
	if len(res) == 0 {
		empty()
		return
	}

	items := getProfileItemsWithChainData(res, nftAsset)
	data := new(ItemsOfAddressRsp).instance(req.PageSize, req.PageNo, totalPage, totalCnt, items)
	output(&c.Controller, data)
}

func getProfileItemsWithChainData(data map[*big.Int]string, nftAsset *models.Token) []*Item {
	profileReqs := make([]*mcm.FetchRequestParams, 0)
	for tokenId, url := range data {
		profileReqs = append(profileReqs, &mcm.FetchRequestParams{
			TokenId: models.NewBigInt(tokenId),
			Url:     url,
		})
	}
	profiles, _ := fetcher.BatchFetch(nftAsset.TokenBasicName, profileReqs)
	profileMap := make(map[string]*models.NFTProfile)
	if profiles != nil {
		for _, v := range profiles {
			profileMap[v.NftTokenId.String()] = v
		}
	}
	items := make([]*Item, 0)
	for tokenId, _ := range data {
		profile := profileMap[tokenId.String()]
		items = append(items, new(Item).instance(tokenId, profile))
	}
	return items
}
