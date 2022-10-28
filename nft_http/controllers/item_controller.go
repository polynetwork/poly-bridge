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
	"poly-bridge/basedef"
	"poly-bridge/models"
	mcm "poly-bridge/nft_http/meta/common"
	"poly-bridge/nft_http/nft_sdk"
	"sort"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type ItemController struct {
	web.Controller
}

func (c *ItemController) Items() {
	var req ItemsOfAddressReq
	if !input(&c.Controller, &req) {
		return
	}

	if strings.Trim(req.TokenId, " ") != "" {
		c.fetchSingleNFTItem(&req)
	} else {
		c.batchFetchNFTItems(&req)
	}
}

func (c *ItemController) fetchSingleNFTItem(req *ItemsOfAddressReq) {
	// check params
	//tokenId, err := checkNumString(req.TokenId)
	//if err != nil {
	//	customInput(&c.Controller, ErrCodeRequest, err.Error())
	//	return
	//}
	if !checkPageSize(&c.Controller, req.PageSize, 20) {
		return
	}
	sdk, wrapper, _, err := selectNodeAndWrapper(req.ChainId)
	if err != nil {
		customInput(&c.Controller, ErrCodeRequest, err.Error())
		return
	}
	token := selectNFTAsset(req.Asset)
	if token == nil {
		customInput(&c.Controller, ErrCodeRequest, "NFT Asset not exist")
		return
	}
	item, err := getSingleItem(sdk, wrapper, token, req.TokenId, req.Address)
	if err != nil {
		logs.Error("get single item err: %v", err)
	}

	items := make([]*Item, 0)
	if item != nil {
		items = append(items, item)
	}
	data := new(ItemsOfAddressRsp).instance(req.PageSize, req.PageNo, 0, len(items), items)
	output(&c.Controller, data)
}

func (c *ItemController) batchFetchNFTItems(req *ItemsOfAddressReq) {
	// check size in contract
	if !checkPageSize(&c.Controller, req.PageSize, 20) {
		return
	}
	sdk, wrapper, _, err := selectNodeAndWrapper(req.ChainId)
	if err != nil {
		customInput(&c.Controller, ErrCodeRequest, err.Error())
		return
	}
	token := selectNFTAsset(req.Asset)
	if token == nil {
		customInput(&c.Controller, ErrCodeRequest, "NFT Asset not exist")
		return
	}

	// get user balance and format page attribute
	asset := token.Hash
	owner := req.Address
	bigTotalCnt, err := sdk.NFTBalance(asset, owner)
	if err != nil {
		customInput(&c.Controller, ErrCodeRequest, err.Error())
		return
	}
	totalCnt := int(bigTotalCnt.Uint64())
	totalPage := getPageNo(totalCnt, req.PageSize)

	// define empty output
	response := func(list []*Item) {
		data := new(ItemsOfAddressRsp).instance(req.PageSize, req.PageNo, totalPage, totalCnt, list)
		output(&c.Controller, data)
	}

	// check user balance and query index
	if totalCnt == 0 {
		response(nil)
		return
	}
	start := req.PageNo * req.PageSize
	if start >= totalCnt {
		customInput(&c.Controller, ErrCodeRequest, "start out of range")
		return
	}
	length := req.PageSize
	if start+length > totalCnt {
		length = totalCnt - start
	}

	// get token id list from contract, order by index
	tokenIdUrlMap, err := sdk.GetTokensByIndex(wrapper, asset, owner, start, length)
	if err != nil {
		logs.Error("GetOwnerNFTsByIndex err: %v", err)
		response(nil)
		return
	}
	if len(tokenIdUrlMap) == 0 {
		response(nil)
		return
	}

	items := getItemsWithChainData(token.TokenBasicName, token.Hash, token.ChainId, tokenIdUrlMap, token.TokenBasic)
	response(items)
}

func getSingleItem(
	sdk nft_sdk.INftSdkPro,
	wrapper string,
	asset *models.Token,
	tokenId string,
	ownerHash string,
) (item *Item, err error) {
	// get and output cache if exist
	cache, ok := GetItemCache(asset.ChainId, asset.Hash, tokenId)
	if ok {
		return cache, nil
	}

	// fetch url from wrapper contract
	// do not need to check user address if ownerHash is empty
	var url string
	assetAddr := asset.Hash
	if ownerHash == "" {
		url, err = sdk.GetNFTUrl(assetAddr, tokenId)
	} else {
		url, err = sdk.GetAndCheckTokenUrl(wrapper, assetAddr, ownerHash, tokenId)
	}
	if err != nil {
		return
	}

	profile, err := fetcher.Fetch(asset.ChainId, asset.TokenBasicName, &mcm.FetchRequestParams{
		TokenId: tokenId,
		Url:     url,
	})
	if err != nil {
		return
	}
	item = new(Item).instance(asset.TokenBasicName, tokenId, profile)
	if item != nil {
		if asset.TokenBasic.ImageDisplayType == 0 {
			// display asset image instead of specific token image
			item.Image = asset.TokenBasic.Meta
		}
	}
	SetItemCache(asset.ChainId, asset.Hash, tokenId, item)
	return
}

func getItemsWithChainData(name string, asset string, chainId uint64, tokenIdUrlMap map[string]string, tokenBasic *models.TokenBasic) []*Item {
	list := make([]*Item, 0)
	// get cache if exist
	profileReqs := make([]*mcm.FetchRequestParams, 0)
	for tokenId, url := range tokenIdUrlMap {
		cache, ok := GetItemCache(chainId, asset, tokenId)
		if ok {
			list = append(list, cache)
			delete(tokenIdUrlMap, tokenId)
			continue
		}

		req := &mcm.FetchRequestParams{
			TokenId: tokenId,
			Url:     url,
		}
		profileReqs = append(profileReqs, req)
	}
	// fetch metadata list and show rpc time
	tBeforeBatchFetch := time.Now().UnixNano()
	profiles, err := fetcher.BatchFetch(chainId, name, profileReqs)
	if err != nil {
		logs.Error("batch fetch NFT profiles err: %v", err)
	}
	tAfterBatchFetch := time.Now().UnixNano()
	debugBatchFetchTime := (tAfterBatchFetch - tBeforeBatchFetch) / int64(time.Microsecond)
	logs.Info("batchFetchNFTItems - batchFetchTime: %d microsecond， profiles %d", debugBatchFetchTime, len(profiles))
	// convert to items
	for _, v := range profiles {
		tokenId := v.NftTokenId
		item := new(Item).instance(name, tokenId, v)
		if item != nil {
			// display asset image instead of specific token image
			if tokenBasic.ImageDisplayType == 0 {
				item.Image = tokenBasic.Meta
			}
		}
		list = append(list, item)
		SetItemCache(chainId, asset, tokenId, item)
		delete(tokenIdUrlMap, tokenId)
	}
	for tokenId := range tokenIdUrlMap {
		item := new(Item).instance(name, tokenId, nil)
		list = append(list, item)
	}
	// sort items with token id
	if len(list) < 2 {
		return list
	}
	if chainId == basedef.NEO3_CROSSCHAIN_ID {
		sort.Slice(list, func(i, j int) bool {
			itemi := list[i].TokenId
			itemj := list[j].TokenId
			return itemi < itemj
		})
	} else {
		sort.Slice(list, func(i, j int) bool {
			itemi, _ := string2Big(list[i].TokenId)
			itemj, _ := string2Big(list[j].TokenId)
			return itemi.Cmp(itemj) < 0
		})
	}
	return list
}

func string2Big(str string) (*big.Int, bool) {
	return new(big.Int).SetString(str, 10)
}
