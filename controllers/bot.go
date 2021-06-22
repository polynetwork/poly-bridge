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
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/models"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type BotController struct {
	beego.Controller
	Conf conf.Config
}

type CheckFeeResult struct {
	Pass bool
	Paid float64
	Min  float64
}

func (c *BotController) CheckFees() {
	hashes := []string{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &hashes)
	if err != nil {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}

	result, err := c.checkFees(hashes)
	if err == nil {
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
	c.Data["json"] = err.Error()
	c.Ctx.ResponseWriter.WriteHeader(400)
	c.ServeJSON()
}

func (c *BotController) checkFees(hashes []string) (fees map[string]CheckFeeResult, err error) {
	wrapperTransactionWithTokens := make([]*models.WrapperTransactionWithToken, 0)
	err = db.Table("wrapper_transactions").Where("hash in ?", hashes).Preload("FeeToken").Preload("FeeToken.TokenBasic").Find(&wrapperTransactionWithTokens).Error
	if err != nil {
		return
	}
	o3Hashes := []string{}
	for _, tx := range wrapperTransactionWithTokens {
		if tx.DstChainId == basedef.O3_CROSSCHAIN_ID {
			o3Hashes = append(o3Hashes, tx.Hash)
		}
	}
	if len(o3Hashes) > 0 {
		srcHashes, err := getSwapSrcTransactions(o3Hashes)
		o3srcs := []string{}
		for _, v := range srcHashes {
			o3srcs = append(o3srcs, v)
		}

		o3txs := make([]*models.WrapperTransactionWithToken, 0)
		err = db.Table("wrapper_transactions").Where("hash in ?", hashes).Preload("FeeToken").Preload("FeeToken.TokenBasic").Find(&o3txs).Error
		if err != nil {
			return nil, err
		}
		wrapperTransactionWithTokens = append(wrapperTransactionWithTokens, o3txs...)
	}

	chainFees := make([]*models.ChainFee, 0)
	db.Preload("TokenBasic").Find(&chainFees)
	chain2Fees := make(map[uint64]*models.ChainFee, 0)
	for _, chainFee := range chainFees {
		chain2Fees[chainFee.ChainId] = chainFee
	}

	fees = make(map[string]CheckFeeResult, 0)
	for _, tx := range wrapperTransactionWithTokens {
		if tx.DstChainId == basedef.O3_CROSSCHAIN_ID {
			continue
		}
		chainId := tx.DstChainId
		if chainId == basedef.O3_CROSSCHAIN_ID {
			chainId = tx.SrcChainId
		}

		chainFee, ok := chain2Fees[chainId]
		if !ok {
			logs.Error("Failed to find chain fee for %d", tx.DstChainId)
			continue
		}

		x := new(big.Int).Mul(&tx.FeeAmount.Int, big.NewInt(tx.FeeToken.TokenBasic.Price))
		feePay := new(big.Float).Quo(new(big.Float).SetInt(x), new(big.Float).SetInt64(basedef.Int64FromFigure(int(tx.FeeToken.Precision))))
		feePay = new(big.Float).Quo(feePay, new(big.Float).SetInt64(basedef.PRICE_PRECISION))
		x = new(big.Int).Mul(&chainFee.MinFee.Int, big.NewInt(chainFee.TokenBasic.Price))
		feeMin := new(big.Float).Quo(new(big.Float).SetInt(x), new(big.Float).SetInt64(basedef.PRICE_PRECISION))
		feeMin = new(big.Float).Quo(feeMin, new(big.Float).SetInt64(basedef.FEE_PRECISION))
		feeMin = new(big.Float).Quo(feeMin, new(big.Float).SetInt64(basedef.Int64FromFigure(int(chainFee.TokenBasic.Precision))))
		res := CheckFeeResult{}
		if feePay.Cmp(feeMin) >= 0 {
			res.Pass = true
		}
		res.Paid, _ = feePay.Float64()
		res.Min, _ = feeMin.Float64()
		fees[tx.Hash] = res
	}

	return
}

func (c *BotController) GetTxs() {
	var err error
	var transactionsOfUnfinishedReq models.TransactionsOfUnfinishedReq
	transactionsOfUnfinishedReq.PageNo, _ = strconv.Atoi(c.Ctx.Input.Query("page_no"))
	transactionsOfUnfinishedReq.PageSize, _ = strconv.Atoi(c.Ctx.Input.Query("page_size"))
	if transactionsOfUnfinishedReq.PageSize == 0 {
		transactionsOfUnfinishedReq.PageSize = 10
	}

	txs, count, err := c.getTxs(transactionsOfUnfinishedReq.PageSize, transactionsOfUnfinishedReq.PageNo)
	if err == nil {
		// Check fee
		hashes := make([]string, len(txs))
		for i, tx := range txs {
			hashes[i] = tx.SrcHash
		}
		fees, checkFeeError := c.checkFees(hashes)
		if checkFeeError != nil {
			err = checkFeeError
		} else {
			resp := models.MakeTransactionOfUnfinishedRsp(transactionsOfUnfinishedReq.PageSize, transactionsOfUnfinishedReq.PageNo,
				(count+transactionsOfUnfinishedReq.PageSize-1)/transactionsOfUnfinishedReq.PageSize, count, txs)
			c.Data["json"] = struct {
				models.TransactionOfUnfinishedRsp `json:",inline"`
				CheckFeeResult                    map[string]CheckFeeResult
			}{
				TransactionOfUnfinishedRsp: *resp,
				CheckFeeResult:             fees,
			}
			c.ServeJSON()
			return
		}
	}

	c.Data["json"] = err.Error()
	c.Ctx.ResponseWriter.WriteHeader(400)
	c.ServeJSON()
}

func (c *BotController) getTxs(pageSize, pageNo int) ([]*models.SrcPolyDstRelation, int, error) {
	srcPolyDstRelations := make([]*models.SrcPolyDstRelation, 0)
	tt := time.Now().Unix()
	from := tt - c.Conf.EventEffectConfig.HowOld2
	res := db.Table("src_transactions").
		Select("src_transactions.hash as src_hash, poly_transactions.hash as poly_hash, dst_transactions.hash as dst_hash, src_transactions.chain_id as chain_id, src_transfers.asset as token_hash, wrapper_transactions.fee_token_hash as fee_token_hash").
		Where("dst_transactions.hash is null").
		Where("src_transactions.standard = ?", 0).
		Where("src_transactions.time > ?", tt-24*60*60*28).
		Where("wrapper_transactions.time < ?", from).
		Joins("left join src_transfers on src_transactions.hash = src_transfers.tx_hash").
		Joins("left join poly_transactions on src_transactions.hash = poly_transactions.src_hash").
		Joins("left join dst_transactions on poly_transactions.hash = dst_transactions.poly_hash").
		Joins("inner join wrapper_transactions on src_transactions.hash = wrapper_transactions.hash").
		Preload("WrapperTransaction").
		Preload("SrcTransaction").
		Preload("SrcTransaction.SrcTransfer").
		Preload("PolyTransaction").
		Preload("DstTransaction").
		Preload("DstTransaction.DstTransfer").
		Preload("Token").
		Preload("Token.TokenBasic").
		Preload("FeeToken").
		Limit(pageSize).Offset(pageSize * pageNo).
		Order("src_transactions.time desc").
		Find(&srcPolyDstRelations)
	if res.Error != nil {
		return nil, 0, res.Error
	}
	var transactionNum int64
	err := db.Table("src_transactions").
		Where("dst_transactions.hash is null").
		Where("src_transactions.standard = ?", 0).
		Where("src_transactions.time > ?", tt-24*60*60*28).
		Where("wrapper_transactions.time < ?", from).
		Joins("left join poly_transactions on src_transactions.hash = poly_transactions.src_hash").
		Joins("left join dst_transactions on poly_transactions.hash = dst_transactions.poly_hash").
		Joins("inner join wrapper_transactions on src_transactions.hash = wrapper_transactions.hash").
		Count(&transactionNum).Error
	if err != nil {
		return nil, 0, err
	}
	return srcPolyDstRelations, int(transactionNum), nil
}
