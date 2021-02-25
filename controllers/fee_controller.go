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
	"github.com/astaxie/beego/logs"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/models"
)

type FeeController struct {
	beego.Controller
}

func (c *FeeController) GetFee() {
	var getFeeReq models.GetFeeReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &getFeeReq); err != nil {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
	}
	token := new(models.Token)
	res := db.Where("hash = ? and chain_id = ?", getFeeReq.Hash, getFeeReq.SrcChainId).Preload("TokenBasic").First(token)
	if res.RowsAffected == 0 {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("chain: %d does not have token: %s", getFeeReq.SrcChainId, getFeeReq.Hash))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	chainFee := new(models.ChainFee)
	res = db.Where("chain_id = ?", getFeeReq.DstChainId).Preload("TokenBasic").First(chainFee)
	if res.RowsAffected == 0 {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("chain: %d does not have fee", getFeeReq.DstChainId))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	proxyFee := new(big.Float).SetInt(&chainFee.ProxyFee.Int)
	proxyFee = new(big.Float).Quo(proxyFee, new(big.Float).SetInt64(basedef.FEE_PRECISION))
	proxyFee = new(big.Float).Quo(proxyFee, new(big.Float).SetInt64(basedef.Int64FromFigure(int(chainFee.TokenBasic.Precision))))
	usdtFee := new(big.Float).Mul(proxyFee, new(big.Float).SetInt64(chainFee.TokenBasic.Price))
	usdtFee = new(big.Float).Quo(usdtFee, new(big.Float).SetInt64(basedef.PRICE_PRECISION))
	tokenFee := new(big.Float).Mul(usdtFee, new(big.Float).SetInt64(basedef.PRICE_PRECISION))
	tokenFee = new(big.Float).Quo(tokenFee, new(big.Float).SetInt64(token.TokenBasic.Price))
	tokenFeeWithPrecision := new(big.Float).Mul(tokenFee, new(big.Float).SetInt64(basedef.Int64FromFigure(int(token.Precision))))
	c.Data["json"] = models.MakeGetFeeRsp(getFeeReq.SrcChainId, getFeeReq.Hash, getFeeReq.DstChainId, usdtFee, tokenFee, tokenFeeWithPrecision)
	c.ServeJSON()
}

func (c *FeeController) CheckFee() {
	logs.Debug("check fee request: %s", string(c.Ctx.Input.RequestBody))
	var checkFeesReq models.CheckFeesReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &checkFeesReq); err != nil {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
	}
	hash2ChainId := make(map[string]uint64, 0)
	requestHashs := make([]string, 0)
	for _, check := range checkFeesReq.Checks {
		hash2ChainId[check.Hash] = check.ChainId
		requestHashs = append(requestHashs, check.Hash)
		requestHashs = append(requestHashs, basedef.HexStringReverse(check.Hash))
	}
	srcTransactions := make([]*models.SrcTransaction, 0)
	db.Model(&models.SrcTransaction{}).Where("(`key` in ? or `hash` in ?)", requestHashs, requestHashs).Find(&srcTransactions)
	key2Txhash := make(map[string]string, 0)
	for _, srcTransaction := range srcTransactions {
		prefix := srcTransaction.Key[0:8]
		if prefix == "00000000" {
			chainId, ok := hash2ChainId[srcTransaction.Key]
			if ok && chainId == srcTransaction.ChainId {
				key2Txhash[srcTransaction.Key] = srcTransaction.Hash
			}
		} else {
			key2Txhash[srcTransaction.Hash] = srcTransaction.Hash
			key2Txhash[basedef.HexStringReverse(srcTransaction.Hash)] = srcTransaction.Hash
		}
	}
	checkHashes := make([]string, 0)
	for _, check := range checkFeesReq.Checks {
		newHash, ok := key2Txhash[check.Hash]
		if ok {
			checkHashes = append(checkHashes, newHash)
		}
	}
	wrapperTransactionWithTokens := make([]*models.WrapperTransactionWithToken, 0)
	db.Table("wrapper_transactions").Where("hash in ?", checkHashes).Preload("FeeToken").Preload("FeeToken.TokenBasic").Find(&wrapperTransactionWithTokens)
	txHash2WrapperTransaction := make(map[string]*models.WrapperTransactionWithToken, 0)
	for _, wrapperTransactionWithToken := range wrapperTransactionWithTokens {
		txHash2WrapperTransaction[wrapperTransactionWithToken.Hash] = wrapperTransactionWithToken
	}
	chainFees := make([]*models.ChainFee, 0)
	db.Preload("TokenBasic").Find(&chainFees)
	chain2Fees := make(map[uint64]*models.ChainFee, 0)
	for _, chainFee := range chainFees {
		chain2Fees[chainFee.ChainId] = chainFee
	}
	checkFees := make([]*models.CheckFee, 0)
	for _, check := range checkFeesReq.Checks {
		checkFee := &models.CheckFee{}
		checkFee.Hash = check.Hash
		checkFee.ChainId = check.ChainId
		checkFee.Amount = new(big.Float).SetInt64(0)
		checkFee.MinProxyFee = new(big.Float).SetInt64(0)
		_, ok := chain2Fees[check.ChainId]
		if !ok {
			checkFee.PayState = -1
			checkFees = append(checkFees, checkFee)
			continue
		}
		newHash, ok := key2Txhash[check.Hash]
		if !ok {
			checkFee.PayState = 0
			checkFees = append(checkFees, checkFee)
			continue
		}
		wrapperTransactionWithToken, ok := txHash2WrapperTransaction[newHash]
		if !ok {
			checkFee.PayState = -1
			checkFees = append(checkFees, checkFee)
			continue
		}
		chainFee, ok := chain2Fees[wrapperTransactionWithToken.DstChainId]
		if !ok {
			checkFee.PayState = -1
			checkFees = append(checkFees, checkFee)
			continue
		}
		x := new(big.Int).Mul(&wrapperTransactionWithToken.FeeAmount.Int, big.NewInt(wrapperTransactionWithToken.FeeToken.TokenBasic.Price))
		feePay := new(big.Float).Quo(new(big.Float).SetInt(x), new(big.Float).SetInt64(basedef.Int64FromFigure(int(wrapperTransactionWithToken.FeeToken.Precision))))
		feePay = new(big.Float).Quo(feePay, new(big.Float).SetInt64(basedef.PRICE_PRECISION))
		x = new(big.Int).Mul(&chainFee.MinFee.Int, big.NewInt(chainFee.TokenBasic.Price))
		feeMin := new(big.Float).Quo(new(big.Float).SetInt(x), new(big.Float).SetInt64(basedef.PRICE_PRECISION))
		feeMin = new(big.Float).Quo(feeMin, new(big.Float).SetInt64(basedef.FEE_PRECISION))
		feeMin = new(big.Float).Quo(feeMin, new(big.Float).SetInt64(basedef.Int64FromFigure(int(chainFee.TokenBasic.Precision))))
		if feePay.Cmp(feeMin) >= 0 {
			checkFee.PayState = 1
		} else {
			checkFee.PayState = -1
		}
		checkFee.Amount = feePay
		checkFee.MinProxyFee = feeMin
		checkFees = append(checkFees, checkFee)
	}
	c.Data["json"] = models.MakeCheckFeesRsp(checkFees)
	c.ServeJSON()
}
