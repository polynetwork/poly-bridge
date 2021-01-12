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
	"fmt"
	"github.com/astaxie/beego"
	"math/big"
	"poly-bridge/conf"
	"poly-bridge/models"
	"poly-bridge/utils"
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
	res := db.Where("hash = ? and chain_id = ?", getFeeReq.Hash, getFeeReq.SrcChainId).Preload("TokenBasic").First(token)
	if res.RowsAffected == 0 {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("chain: %d does not have token: %s", getFeeReq.SrcChainId, getFeeReq.Hash))
		c.ServeJSON()
		return
	}
	chainFee := new(models.ChainFee)
	res = db.Where("chain_id = ?", getFeeReq.DstChainId).Preload("TokenBasic").First(chainFee)
	if res.RowsAffected == 0 {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("chain: %d does not have fee", getFeeReq.DstChainId))
		c.ServeJSON()
		return
	}
	proxyFee := new(big.Float).SetInt(&chainFee.ProxyFee.Int)
	proxyFee = new(big.Float).Quo(proxyFee, new(big.Float).SetInt64(conf.FEE_PRECISION))
	proxyFee = new(big.Float).Quo(proxyFee, new(big.Float).SetInt64(utils.Int64FromFigure(int(chainFee.TokenBasic.Precision))))
	proxyFee = new(big.Float).Mul(proxyFee, new(big.Float).SetInt64(chainFee.TokenBasic.Price))
	usdtFee := new(big.Float).Quo(proxyFee, new(big.Float).SetInt64(conf.PRICE_PRECISION))
	tokenFee := new(big.Float).Quo(proxyFee, new(big.Float).SetInt64(token.TokenBasic.Price))
	tokenFeeWithPrecision := new(big.Float).Mul(tokenFee, new(big.Float).SetInt64(utils.Int64FromFigure(int(token.Precision))))
	c.Data["json"] = models.MakeGetFeeRsp(getFeeReq.SrcChainId, getFeeReq.Hash, getFeeReq.DstChainId, usdtFee, tokenFee, tokenFeeWithPrecision)
	c.ServeJSON()
}

func (c *FeeController) CheckFee() {
	var checkFeesReq models.CheckFeesReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &checkFeesReq); err != nil {
		panic(err)
	}
	db := newDB()
	srcTransactions := make([]*models.SrcTransaction, 0)
	db.Model(&models.SrcTransaction{}).Where("key in ?", checkFeesReq.Hashs).Find(&srcTransactions)
	keyTxhashMap := make(map[string]string, 0)
	for _, srcTransaction := range srcTransactions {
		keyTxhashMap[srcTransaction.Key] = srcTransaction.Hash
	}
	checkHashs := make([]string, 0)
	for _, hash := range checkFeesReq.Hashs {
		newHash, ok := keyTxhashMap[hash]
		if ok {
			checkHashs = append(checkHashs, newHash)
		} else {
			checkHashs = append(checkHashs, hash)
		}
	}
	wrapperTransactionWithTokens := make([]*models.WrapperTransactionWithToken, 0)
	res := db.Table("wrapper_transactions").Where("hash in ?", checkHashs).Preload("FeeToken").Preload("FeeToken.TokenBasic").Find(&wrapperTransactionWithTokens)
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
		chainFee, ok := chainFeesMap[wrapperTransactionWithToken.DstChainId]
		if !ok {
			continue
		}
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
			MinProxyFee: feeMin.String(),
		})
	}
	c.Data["json"] = models.MakeCheckFeesRsp(checkFees)
	c.ServeJSON()
}
