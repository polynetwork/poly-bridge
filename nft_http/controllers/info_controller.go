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
	"time"

	"github.com/astaxie/beego"
	"github.com/ethereum/go-ethereum/common"
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

	cache, ok := GetHomePageCache(req.ChainId)
	if ok && cache != nil && cache.Time.Add(600*time.Second).After(time.Now()) {
		output(&c.Controller, cache.Rsp)
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

	chainAssets := selectAssetsByChainId(req.ChainId)
	totalCnt := len(chainAssets)
	list := make([]*AssetItems, 0)
	for _, v := range chainAssets {
		if v.TokenBasic.MetaFetcherType != meta.FetcherTypeUnknown {
			addr := common.HexToAddress(v.Hash)
			tokenUrls, _ := sdk.GetUnCrossChainNFTsByIndex(wrapper, addr, 0, req.Size)
			if len(tokenUrls) == 0 {
				continue
			}
			items := getItemsWithChainData(v.TokenBasicName, v.Hash, v.ChainId, tokenUrls)
			assets := &AssetItems{
				Asset: new(AssetRsp).instance(v),
				Items: items,
			}
			list = append(list, assets)
			break
		}
	}

	data := new(HomeRsp).instance(totalCnt, list)
	SetHomePageCache(req.ChainId, &CacheHomeRsp{
		Rsp:  data,
		Time: time.Now(),
	})
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
