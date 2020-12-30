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
	"math/big"
	"poly-swap/conf"
	"poly-swap/models"
	"poly-swap/utils"
)

type FeeController struct {
	beego.Controller
}

func (c *FeeController) GetFee() {
	var getFeeReq models.GetFeeReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &getFeeReq); err != nil {
		panic(err)
	}
	db := newDB()
	token := new(models.Token)
	db.Where("hash = ?", getFeeReq.Hash).Preload("TokenBasic").First(token)
	chainFee := new(models.ChainFee)
	db.Where("chain_id = ?", getFeeReq.ChainId).Preload("TokenBasic").First(chainFee)
	x := new(big.Int).Mul(&chainFee.ProxyFee.Int, big.NewInt(chainFee.TokenBasic.Price))
	y := new(big.Float).Quo(new(big.Float).SetInt(x), new(big.Float).SetInt64(utils.Int64FromFigure(int(chainFee.TokenBasic.Precision))))
	y = new(big.Float).Quo(y, new(big.Float).SetInt64(conf.FEE_PRECISION))
	y = new(big.Float).Quo(y, new(big.Float).SetInt64(token.TokenBasic.Price))
	y = new(big.Float).Mul(y, new(big.Float).SetInt64(utils.Int64FromFigure(int(token.Precision))))
	z, _ := y.Float64()
	c.Data["json"] = models.MakeGetFeeRsp(getFeeReq.ChainId, getFeeReq.Hash, z)
	c.ServeJSON()
}

func (c *FeeController) CheckFee() {
	var checkFeeReq models.CheckFeeReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &checkFeeReq); err != nil {
		panic(err)
	}
	db := newDB()
	wrapperTransactionWithToken := new(models.WrapperTransactionWithToken)
	db.Model(&models.WrapperTransaction{}).Where("hash = ?", checkFeeReq.Hash).Preload("FeeToken").Preload("FeeToken.TokenBasic").First(wrapperTransactionWithToken)
	chainFee := new(models.ChainFee)
	db.Where("chain_id = ?", wrapperTransactionWithToken.DstChainId).Preload("TokenBasic").First(chainFee)
	hasPay := false
	x := new(big.Int).Mul(&wrapperTransactionWithToken.FeeAmount.Int, big.NewInt(wrapperTransactionWithToken.FeeToken.TokenBasic.Price))
	feePay := new(big.Float).Quo(new(big.Float).SetInt(x), new(big.Float).SetInt64(utils.Int64FromFigure(int(wrapperTransactionWithToken.FeeToken.Precision))))
	feePay = new(big.Float).Quo(feePay, new(big.Float).SetInt64(conf.PRICE_PRECISION))
	x = new(big.Int).Mul(&chainFee.MinFee.Int, big.NewInt(chainFee.TokenBasic.Price))
	feeMin := new(big.Float).Quo(new(big.Float).SetInt(x), new(big.Float).SetInt64(conf.PRICE_PRECISION))
	feeMin = new(big.Float).Quo(feeMin, new(big.Float).SetInt64(conf.FEE_PRECISION))
	feeMin = new(big.Float).Quo(feeMin, new(big.Float).SetInt64(utils.Int64FromFigure(int(chainFee.TokenBasic.Precision))))
	if feePay.Cmp(feeMin) >= 0 {
		hasPay = true
	}
	z, _ := feePay.Float64()
	c.Data["json"] = models.MakeCheckFeeRsp(hasPay, z)
	c.ServeJSON()
}
