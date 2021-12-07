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
	"fmt"
	"poly-bridge/models"
	"poly-bridge/nft_http/meta"
	"poly-bridge/utils/net"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/ethereum/go-ethereum/common"
)

var (
	mode string
	port int
)

type InfoController struct {
	web.Controller
}

func (c *InfoController) Get() {
	url, err := captureUrl()
	if err != nil {
		c.Data["json"] = models.MakeErrorRsp(err.Error())
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	explorer := &PolyBridgeInfoResp{
		Version: "v1",
		URL:     url,
	}
	explorer.Entrance = make([]*ContractEntrance, 0)
	for _, v := range chainConfig {
		explorer.Entrance = append(explorer.Entrance, &ContractEntrance{
			ChainId:         v.ChainId,
			ChainName:       v.ChainName,
			WrapperContract: v.NFTWrapperContract,
		})
	}
	c.Data["json"] = explorer
	c.ServeJSON()
}

func (c *InfoController) Home() {
	var req HomeReq
	if !input(&c.Controller, &req) {
		return
	}
	//if !checkPageSize(&c.Controller, req.PageSize) {
	//	return
	//}

	chainAssets := selectAssetsByChainId(req.ChainId)
	totalCnt := len(chainAssets)
	list := make([]*AssetItems, 0)
	start := req.PageSize * req.PageNo
	end := start + req.PageSize

	for _, asset := range chainAssets {
		items := &AssetItems{
			Asset:   new(AssetRsp).instance(asset),
			Items:   nil,
			HasMore: false,
		}
		if cache, exist := GetHomePageItemsCache(req.ChainId, asset.TokenBasicName); exist {
			if total := len(cache); total > start {
				if end > total && end > start {
					end = total
				}
				items.Items = cache[start:end]
				if len(cache) > end {
					items.HasMore = true
				}
				list = append(list, items)
			}
		}
	}

	data := new(HomeRsp).instance(totalCnt, list)
	output(&c.Controller, data)
}

func prepareHomepageItems(asset *models.Token, maxNum int) (bool, error) {
	sdk, inquirer, lockProxies, err := selectNodeAndWrapper(asset.ChainId)
	if err != nil {
		return false, err
	}

	if asset.TokenBasic.MetaFetcherType == meta.FetcherTypeUnknown {
		return false, fmt.Errorf("invalid fetcher type")
	}

	chainId := asset.ChainId
	assetName := asset.TokenBasicName
	assetAddr := common.HexToAddress(asset.Hash)
	pageSize := 10

	list := make([]*Item, 0)
	for start := 0; start < maxNum; start += pageSize {
		tokenUrls, _ := sdk.GetUnCrossChainNFTsByIndex(inquirer, assetAddr, lockProxies, start, pageSize)
		if len(tokenUrls) == 0 {
			break
		}
		items := getItemsWithChainData(assetName, asset.Hash, chainId, tokenUrls)
		list = append(list, items...)
	}

	if len(list) == 0 {
		return false, nil
	}

	SetHomePageItemsCache(asset.ChainId, assetName, list)
	cache, _ := GetHomePageItemsCache(asset.ChainId, asset.TokenBasicName)
	logs.Info("prepare chain %d asset %s home page items, total %d ", chainId, assetName, len(cache))
	return true, nil
}

func SetBaseInfo(_mode string, _port int) {
	mode = _mode
	port = _port
}

func captureUrl() (string, error) {
	switch mode {
	case "dev", "test":
		ips, err := net.GetLocalIPv4s()
		if err != nil {
			return "", err
		}
		if len(ips) == 0 {
			return "", fmt.Errorf("local IPv4s reading error")
		}
		return fmt.Sprintf("http://%s:%d/nft", ips[0], port), nil
	case "prod":
		return "http://bridge.poly.network/nft", nil
	}
	return "", fmt.Errorf("web running mode invalid")
}
