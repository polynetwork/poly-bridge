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

package http

import (
	"encoding/json"
	"fmt"
	"poly-bridge/models"

	"github.com/beego/beego/v2/server/web"
)

type TokenController struct {
	web.Controller
}

func (c *TokenController) Tokens() {
	var tokensReq models.TokensReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &tokensReq); err != nil {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
	}
	tokens := make([]*models.Token, 0)
	db.Where("chain_id = ? and standard = 0", tokensReq.ChainId).Preload("TokenBasic").Preload("TokenMaps").Preload("TokenMaps.DstToken").Find(&tokens)
	c.Data["json"] = models.MakeTokensRsp(tokens)
	c.ServeJSON()
}

func (c *TokenController) Token() {
	var tokenReq models.TokenReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &tokenReq); err != nil {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
	}
	token := new(models.Token)
	res := db.Where("hash = ? and chain_id = ? and standard = 0", tokenReq.Hash, tokenReq.ChainId).Preload("TokenBasic").Preload("TokenMaps").Preload("TokenMaps.DstToken").First(token)
	if res.RowsAffected == 0 {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("token: (%s,%d) does not exist", tokenReq.Hash, tokenReq.ChainId))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	c.Data["json"] = models.MakeTokenRsp(token)
	c.ServeJSON()
}

func (c *TokenController) TokenBasics() {
	var tokenBasicReq models.TokenBasicReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &tokenBasicReq); err != nil {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
	}
	tokenBasics := make([]*models.TokenBasic, 0)
	db.Model(&models.TokenBasic{}).Where("standard = 0").Preload("Tokens").Find(&tokenBasics)
	c.Data["json"] = models.MakeTokenBasicsRsp(tokenBasics)
	c.ServeJSON()
}

func (c *TokenController) TokenBasicsInfo() {
	var tokenBasicReq models.TokenBasicsInfoReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &tokenBasicReq); err != nil || tokenBasicReq.PageSize < 1 {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
	}
	tokenBasics := make([]*models.TokenBasic, 0)
	query := db.Model(&models.TokenBasic{}).Where("standard = 0 AND property = 1")
	var totalCount int64
	query.Count(&totalCount)
	query.Limit(tokenBasicReq.PageSize).Offset(tokenBasicReq.PageSize * tokenBasicReq.PageNo).Order("total_amount, name").Preload("Tokens").Find(&tokenBasics)
	c.Data["json"] = models.MakeTokenBasicsInfoRsp(&tokenBasicReq, uint64(totalCount), tokenBasics)
	c.ServeJSON()
}
