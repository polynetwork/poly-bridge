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
	"encoding/json"
	"github.com/astaxie/beego"
	"poly-bridge/models"
)

type TokenMapController struct {
	beego.Controller
}

func (c *TokenMapController) TokenMap() {
	var tokenMapReq models.TokenMapReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &tokenMapReq); err != nil {
		panic(err)
	}
	db := newDB()
	tokenMaps := make([]*models.TokenMap, 0)
	db.Where("src_chain_id = ? and src_token_hash = ?", tokenMapReq.ChainId, tokenMapReq.Hash).Preload("SrcToken").Preload("DstToken").Find(&tokenMaps)
	c.Data["json"] = models.MakeTokenMapsRsp(tokenMaps)
	c.ServeJSON()
}

func (c *TokenMapController) TokenMapReverse() {
	var tokenMapReq models.TokenMapReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &tokenMapReq); err != nil {
		panic(err)
	}
	db := newDB()
	tokenMaps := make([]*models.TokenMap, 0)
	db.Where("dst_chain_id = ? and dst_token_hash = ?", tokenMapReq.ChainId, tokenMapReq.Hash).Preload("SrcToken").Preload("DstToken").Find(&tokenMaps)
	c.Data["json"] = models.MakeTokenMapsRsp(tokenMaps)
	c.ServeJSON()
}
