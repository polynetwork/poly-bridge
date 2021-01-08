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
	var checkFeesReq models.CheckFeesReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &checkFeesReq); err != nil {
		panic(err)
	}
	db := newDB()
	wrapperTransactionWithTokens := make([]*models.WrapperTransactionWithToken, 0)
	res := db.Table("wrapper_transactions").Where("hash in ?", checkFeesReq.Hashs).Preload("FeeToken").Preload("FeeToken.TokenBasic").Find(&wrapperTransactionWithTokens)
	if res.RowsAffected == 0 {
		c.Data["json"] = models.MakeCheckFeesRsp(nil)
		c.ServeJSON()
		return
	}
	chainFees := make([]*models.ChainFee, 0)
	db.Preload("TokenBasic").Find(&chainFees)
	chainFeesMap := make(map[uint64]*models.ChainFee, 0)
	for _, chainFee := range chainFees {
		chainFeesMap[chainFee.ChainId] = chainFee
	}
	checkFees := make([]*models.CheckFee, 0)
	for _, wrapperTransactionWithToken := range wrapperTransactionWithTokens {
		hasPay := false
		x := new(big.Int).Mul(&wrapperTransactionWithToken.FeeAmount.Int, big.NewInt(wrapperTransactionWithToken.FeeToken.TokenBasic.Price))
		feePay := new(big.Float).Quo(new(big.Float).SetInt(x), new(big.Float).SetInt64(utils.Int64FromFigure(int(wrapperTransactionWithToken.FeeToken.Precision))))
		feePay = new(big.Float).Quo(feePay, new(big.Float).SetInt64(conf.PRICE_PRECISION))
		chainFee := chainFeesMap[wrapperTransactionWithToken.DstChainId]
		x = new(big.Int).Mul(&chainFee.MinFee.Int, big.NewInt(chainFee.TokenBasic.Price))
		feeMin := new(big.Float).Quo(new(big.Float).SetInt(x), new(big.Float).SetInt64(conf.PRICE_PRECISION))
		feeMin = new(big.Float).Quo(feeMin, new(big.Float).SetInt64(conf.FEE_PRECISION))
		feeMin = new(big.Float).Quo(feeMin, new(big.Float).SetInt64(utils.Int64FromFigure(int(chainFee.TokenBasic.Precision))))
		if feePay.Cmp(feeMin) >= 0 {
			hasPay = true
		}
		checkFees = append(checkFees, &models.CheckFee{
			Hash:  wrapperTransactionWithToken.Hash,
			HasPay: hasPay,
			Amount: feePay.String(),
		})
	}
	c.Data["json"] = models.MakeCheckFeesRsp(checkFees)
	c.ServeJSON()
}
