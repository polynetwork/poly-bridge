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
	"poly-bridge/models"
)

type AssetController struct {
	beego.Controller
}

func (c *AssetController) Assets() {
	var req AssetsReq
	if !input(&c.Controller, &req) {
		return
	}

	assets := make([]*models.Token, 0)
	db.Where("chain_id = ? and standard = ? and property = ?", req.ChainId, models.TokenTypeErc721, 1).
		Preload("TokenBasic").
		Preload("TokenMaps").
		Preload("TokenMaps.DstToken").
		Preload("TokenMaps.DstToken.TokenBasic").
		Find(&assets)
	data := new(AssetsRsp).instance(assets)
	output(&c.Controller, data)
}

func (c *AssetController) Asset() {
	var req AssetReq
	if !input(&c.Controller, &req) {
		return
	}

	asset := new(models.Token)
	res := db.Where("hash = ? and chain_id = ? and standard = ? and property = ?", req.Hash, req.ChainId, models.TokenTypeErc721, 1).
		Preload("TokenBasic").
		Preload("TokenMaps").
		Preload("TokenMaps.DstToken").
		Preload("TokenMaps.DstToken.TokenBasic").
		First(asset)
	if res.RowsAffected == 0 {
		notExist(&c.Controller)
		return
	}
	data := new(AssetMap).instance(asset)
	output(&c.Controller, data)
}

//
//func (c *AssetController) AssetBasics() {
//	assetBasics := make([]*models.TokenBasic, 0)
//	db.Model(&models.TokenBasic{}).Preload("Assets").Find(&assetBasics)
//	data := MakeNFTAssetBasicsRsp(assetBasics)
//	output(&c.Controller, data)
//}
