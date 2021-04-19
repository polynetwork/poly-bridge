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
	"github.com/ethereum/go-ethereum/common"
	"poly-bridge/nft_http/meta"

	"github.com/astaxie/beego"
	"poly-bridge/models"
	"poly-bridge/utils/net"
)

var (
	mode string
	port int
)

type InfoController struct {
	beego.Controller
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
	c.Data["json"] = explorer
	c.ServeJSON()
}

func (c *InfoController) Home() {
	var req HomeReq
	if !input(&c.Controller, &req) {
		return
	}

	sdk := selectNode(req.ChainId)
	if sdk == nil {
		customInput(&c.Controller, ErrCodeRequest, "chain id not exist")
		return
	}

	if cache, ok := GetHomePageCache(req.ChainId); ok {
		output(&c.Controller, cache)
		return
	}

	chainAssets := selectAssetsByChainId(req.ChainId)
	totalCnt := len(chainAssets)
	assetItems := make([]*AssetItems, 0)
	for _, v := range chainAssets {
		if v.TokenBasic.MetaFetcherType != meta.FetcherTypeUnknown {
			addr := common.HexToAddress(v.Hash)
			tokenIds, _ := sdk.GetAssetNFTs(addr, 0, req.Size)
			chainData, _ := sdk.GetNFTURLs(addr, tokenIds)
			items := getItemsWithChainData(v.TokenBasicName, v.Hash, v.ChainId, chainData)
			assetItem := &AssetItems{
				Asset: new(AssetRsp).instance(v),
				Items: items,
			}
			assetItems = append(assetItems, assetItem)
			break
		}
	}
	data := new(HomeRsp).instance(totalCnt, assetItems)
	SetHomePageCache(req.ChainId, data)

	output(&c.Controller, data)
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
	return "", fmt.Errorf("beego running mode invalid")
}
