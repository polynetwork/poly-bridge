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
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/models"

	"github.com/beego/beego/v2/server/web"
)

type FeeController struct {
	web.Controller
}

func (c *FeeController) GetFee() {
	var req models.GetFeeReq
	if !input(&c.Controller, &req) {
		return
	}
	token := new(models.Token)
	res := db.Where("hash = ? and chain_id = ?", req.Hash, req.SrcChainId).
		Preload("TokenBasic").
		First(token)
	if res.RowsAffected == 0 {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("chain: %d does not have token: %s", req.SrcChainId, req.Hash))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	chainFee := new(models.ChainFee)
	res = db.Where("chain_id = ?", req.DstChainId).Preload("TokenBasic").First(chainFee)
	if res.RowsAffected == 0 {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("chain: %d does not have fee", req.DstChainId))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	fzero := new(big.Float).SetUint64(0)
	proxyFee := new(big.Float).SetInt(&chainFee.ProxyFee.Int)
	proxyFee = new(big.Float).Quo(proxyFee, new(big.Float).SetInt64(basedef.FEE_PRECISION))
	proxyFee = new(big.Float).Quo(proxyFee, new(big.Float).SetInt64(basedef.Int64FromFigure(int(chainFee.TokenBasic.Precision))))
	usdtFee := new(big.Float).Mul(proxyFee, new(big.Float).SetInt64(chainFee.TokenBasic.Price))
	usdtFee = new(big.Float).Quo(usdtFee, new(big.Float).SetInt64(basedef.PRICE_PRECISION))
	tokenFee := new(big.Float).Mul(usdtFee, new(big.Float).SetInt64(basedef.PRICE_PRECISION))
	tokenFee = new(big.Float).Quo(tokenFee, new(big.Float).SetInt64(token.TokenBasic.Price))
	tokenFeeWithPrecision := new(big.Float).Mul(tokenFee, new(big.Float).SetInt64(basedef.Int64FromFigure(int(token.Precision))))
	c.Data["json"] = models.MakeGetFeeRsp(req.SrcChainId, req.Hash, req.DstChainId, usdtFee, tokenFee, tokenFeeWithPrecision, "", fzero, fzero, false, fzero)
	c.ServeJSON()
}
