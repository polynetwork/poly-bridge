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

import "github.com/beego/beego/v2/server/web"

type AssetMapController struct {
	web.Controller
}

//func (c *AssetMapController) AssetMap() {
//	var req NFTAssetMapReq
//	if !input(&c.Controller, &req) {
//		return
//	}
//
//	assetMaps := make([]*models.TokenMap, 0)
//	res := db.Where("src_chain_id = ? and src_token_hash = ?", req.ChainId, req.Hash).
//		Preload("SrcToken").
//		Preload("DstToken").
//		Find(&assetMaps)
//	if res.RowsAffected == 0 {
//		notExist(&c.Controller)
//		return
//	}
//
//	data := MakeNFTAssetMapsRsp(assetMaps)
//	output(&c.Controller, data)
//}
//
//func (c *AssetMapController) AssetMapReverse() {
//	var req NFTAssetMapReq
//	if !input(&c.Controller, &req) {
//		return
//	}
//
//	assetMaps := make([]*models.TokenMap, 0)
//	res := db.Where("dst_chain_id = ? and dst_asset_hash = ?", req.ChainId, req.Hash).
//		Preload("SrcAsset").
//		Preload("DstAsset").
//		Find(&assetMaps)
//	if res.RowsAffected == 0 {
//		notExist(&c.Controller)
//		return
//	}
//
//	data := MakeNFTAssetMapsRsp(assetMaps)
//	output(&c.Controller, data)
//}
