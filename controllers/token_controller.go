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
	"github.com/polynetwork/poly-bridge/models"
)

type TokenController struct {
	beego.Controller
}

func (c *TokenController) Tokens() {
	var tokensReq models.TokensReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &tokensReq); err != nil {
		panic(err)
	}
	db := newDB()
	tokens := make([]*models.Token, 0)
	db.Where("chain_id = ?", tokensReq.ChainId).Preload("TokenBasic").Preload("TokenMaps").Preload("TokenMaps.DstToken").Find(&tokens)
	c.Data["json"] = models.MakeTokensRsp(tokens)
	c.ServeJSON()
}

func (c *TokenController) Token() {
	var tokenReq models.TokenReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &tokenReq); err != nil {
		panic(err)
	}
	db := newDB()
	token := new(models.Token)
	db.Where("hash = ?", tokenReq.Hash).Preload("TokenBasic").Preload("TokenMaps").Preload("TokenMaps.DstToken").First(token)
	c.Data["json"] = models.MakeTokenRsp(token)
	c.ServeJSON()
}

func (c *TokenController) TokenBasics() {
	var tokenBasicReq models.TokenBasicReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &tokenBasicReq); err != nil {
		panic(err)
	}
	db := newDB()
	tokenBasics := make([]*models.TokenBasic, 0)
	db.Model(&models.TokenBasic{}).Preload("Tokens").Find(&tokenBasics)
	c.Data["json"] = models.MakeTokenBasicsRsp(tokenBasics)
	c.ServeJSON()
}
