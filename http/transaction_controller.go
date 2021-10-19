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
	"poly-bridge/basedef"
	"poly-bridge/models"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"gorm.io/gorm"
)

type TransactionController struct {
	web.Controller
}

func (c *TransactionController) Transactions() {
	var transactionsReq models.WrapperTransactionsReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &transactionsReq); err != nil || transactionsReq.PageSize == 0 {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
	}
	transactions := make([]*models.WrapperTransaction, 0)
	db.Limit(transactionsReq.PageSize).Offset(transactionsReq.PageSize * transactionsReq.PageNo).Order("time asc").Find(&transactions)
	var transactionNum int64
	db.Model(&models.WrapperTransaction{}).Count(&transactionNum)
	c.Data["json"] = models.MakeWrapperTransactionsRsp(transactionsReq.PageSize, transactionsReq.PageNo, (int(transactionNum)+transactionsReq.PageSize-1)/transactionsReq.PageSize,
		int(transactionNum), transactions)
	c.ServeJSON()
}

func (c *TransactionController) TransactionsWithFilter() {
	var transactionsReq models.WrapperTransactionsWithFilterReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &transactionsReq); err != nil || transactionsReq.PageSize == 0 {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
	}
	transactions := make([]*models.WrapperTransaction, 0)

	query := db.Model(&models.WrapperTransaction{}).Joins("LEFT JOIN src_transfers ON wrapper_transactions.hash = src_transfers.tx_hash").Where("src_transfers.chain_id = ? and src_transfers.dst_chain_id = ? and src_transfers.asset in ?", transactionsReq.SrcChainId, transactionsReq.DstChainId, transactionsReq.Assets)
	query.Limit(transactionsReq.PageSize).Offset(transactionsReq.PageSize * transactionsReq.PageNo).Order("time asc").Find(&transactions)
	var transactionNum int64
	query.Count(&transactionNum)
	c.Data["json"] = models.MakeWrapperTransactionsRsp(transactionsReq.PageSize, transactionsReq.PageNo, (int(transactionNum)+transactionsReq.PageSize-1)/transactionsReq.PageSize,
		int(transactionNum), transactions)
	c.ServeJSON()
}

func (c *TransactionController) PolyTransactions() {
	var transactionsReq models.PolyTransactionsReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &transactionsReq); err != nil {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
	}
	transactions := make([]*models.PolyTransaction, 0)
	db.Limit(transactionsReq.PageSize).Offset(transactionsReq.PageSize * transactionsReq.PageNo).Order("time asc").Find(&transactions)
	var transactionNum int64
	db.Model(&models.PolyTransaction{}).Count(&transactionNum)
	c.Data["json"] = models.MakePolyTransactionsRsp(transactionsReq.PageSize, transactionsReq.PageNo, (int(transactionNum)+transactionsReq.PageSize-1)/transactionsReq.PageSize,
		int(transactionNum), transactions)
	c.ServeJSON()
}

func (c *TransactionController) TransactionsOfAddressWithFilter() {
	var req models.TransactionsOfAddressWithFilterReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil || req.PageSize == 0 {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
	}
	srcPolyDstRelations := make([]*models.SrcPolyDstRelation, 0)
	query := func(tx *gorm.DB) *gorm.DB {
		return tx.Table("(?) as u", db.Model(&models.SrcTransfer{}).Select("tx_hash as hash, asset as asset, fee_token_hash as fee_token_hash, src_transfers.chain_id as chain_id").Joins("inner join wrapper_transactions on src_transfers.tx_hash = wrapper_transactions.hash").
			Where("`from` in ? or src_transfers.dst_user in ?", req.Addresses, req.Addresses).
			Where("src_transfers.chain_id = ? and src_transfers.dst_chain_id = ? and src_transfers.asset in ?",
				req.SrcChainId,
				req.DstChainId,
				req.Assets,
			),
		).
			Where("src_transactions.standard = ?", 0).
			Select("src_transactions.hash as src_hash, poly_transactions.hash as poly_hash, dst_transactions.hash as dst_hash, src_transactions.chain_id as chain_id, u.asset as token_hash, u.fee_token_hash as fee_token_hash").
			Joins("inner join tokens on u.chain_id = tokens.chain_id and u.asset = tokens.hash").
			Joins("left join src_transactions on u.hash = src_transactions.hash").
			Joins("left join poly_transactions on src_transactions.hash = poly_transactions.src_hash").
			Joins("left join dst_transactions on poly_transactions.hash = dst_transactions.poly_hash")
	}

	err = query(db).
		Preload("WrapperTransaction").
		Preload("SrcTransaction").
		Preload("SrcTransaction.SrcTransfer").
		Preload("PolyTransaction").
		Preload("DstTransaction").
		Preload("DstTransaction.DstTransfer").
		Preload("Token").
		Preload("Token.TokenBasic").
		Preload("FeeToken").
		Limit(req.PageSize).Offset(req.PageSize * req.PageNo).
		Order("src_transactions.time desc").
		Find(&srcPolyDstRelations).Error

	if err == nil {
		var transactionNum int64
		err = query(db).Count(&transactionNum).Error
		if err == nil {
			chains := make([]*models.Chain, 0)
			db.Model(&models.Chain{}).Find(&chains)
			chainsMap := make(map[uint64]*models.Chain)
			for _, chain := range chains {
				chainsMap[chain.ChainId] = chain
			}
			c.Data["json"] = models.MakeTransactionsOfUserRsp(req.PageSize, req.PageNo,
				(int(transactionNum)+req.PageSize-1)/req.PageSize, int(transactionNum), srcPolyDstRelations, chainsMap)
			c.ServeJSON()
			return
		}
	}
	logs.Error("Load data error %v", err)
	c.Data["json"] = models.MakeErrorRsp("service error!")
	c.Ctx.ResponseWriter.WriteHeader(500)
	c.ServeJSON()
}

func (c *TransactionController) TransactionsOfAddress() {
	var transactionsOfAddressReq models.TransactionsOfAddressReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &transactionsOfAddressReq); err != nil {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
	}
	srcPolyDstRelations := make([]*models.SrcPolyDstRelation, 0)
	db.Table("(?) as u", db.Model(&models.SrcTransfer{}).Select("tx_hash as hash, asset as asset, fee_token_hash as fee_token_hash, src_transfers.chain_id as chain_id").Joins("inner join wrapper_transactions on src_transfers.tx_hash = wrapper_transactions.hash").
		Where("`from` in ? or src_transfers.dst_user in ?", transactionsOfAddressReq.Addresses, transactionsOfAddressReq.Addresses)).
		Where("src_transactions.standard = ?", 0).
		Select("src_transactions.hash as src_hash, poly_transactions.hash as poly_hash, dst_transactions.hash as dst_hash, src_transactions.chain_id as chain_id, u.asset as token_hash, u.fee_token_hash as fee_token_hash").
		Joins("inner join tokens on u.chain_id = tokens.chain_id and u.asset = tokens.hash").
		Joins("left join src_transactions on u.hash = src_transactions.hash").
		Joins("left join poly_transactions on src_transactions.hash = poly_transactions.src_hash").
		Joins("left join dst_transactions on poly_transactions.hash = dst_transactions.poly_hash").
		Preload("WrapperTransaction").
		Preload("SrcTransaction").
		Preload("SrcTransaction.SrcTransfer").
		Preload("PolyTransaction").
		Preload("DstTransaction").
		Preload("DstTransaction.DstTransfer").
		Preload("Token").
		Preload("Token.TokenBasic").
		Preload("FeeToken").
		Limit(transactionsOfAddressReq.PageSize).Offset(transactionsOfAddressReq.PageSize * transactionsOfAddressReq.PageNo).
		Order("src_transactions.time desc").
		Find(&srcPolyDstRelations)
	var transactionNum int64
	db.Model(&models.SrcTransfer{}).Joins("inner join wrapper_transactions on src_transfers.tx_hash = wrapper_transactions.hash").Where("`from` in ? or src_transfers.dst_user in ?", transactionsOfAddressReq.Addresses, transactionsOfAddressReq.Addresses).Count(&transactionNum)
	chains := make([]*models.Chain, 0)
	db.Model(&models.Chain{}).Find(&chains)
	chainsMap := make(map[uint64]*models.Chain)
	for _, chain := range chains {
		chainsMap[chain.ChainId] = chain
	}
	c.Data["json"] = models.MakeTransactionsOfUserRsp(transactionsOfAddressReq.PageSize, transactionsOfAddressReq.PageNo,
		(int(transactionNum)+transactionsOfAddressReq.PageSize-1)/transactionsOfAddressReq.PageSize, int(transactionNum), srcPolyDstRelations, chainsMap)
	c.ServeJSON()
}

func (c *TransactionController) getTransactionByHash(hash string) (*models.SrcPolyDstRelation, error) {
	srcPolyDstRelation := new(models.SrcPolyDstRelation)
	res := db.Table("src_transactions").
		Select("src_transactions.hash as src_hash, poly_transactions.hash as poly_hash, dst_transactions.hash as dst_hash, src_transactions.chain_id as chain_id, src_transfers.asset as token_hash, wrapper_transactions.fee_token_hash as fee_token_hash").
		Where("src_transactions.hash = ?", hash).
		Where("src_transactions.standard = ?", 0).
		Joins("left join wrapper_transactions on src_transactions.hash = wrapper_transactions.hash").
		Joins("left join src_transfers on src_transactions.hash = src_transfers.tx_hash").
		Joins("left join poly_transactions on src_transactions.hash = poly_transactions.src_hash").
		Joins("left join dst_transactions on poly_transactions.hash = dst_transactions.poly_hash").
		Preload("WrapperTransaction").
		Preload("SrcTransaction").
		Preload("SrcTransaction.SrcTransfer").
		Preload("PolyTransaction").
		Preload("DstTransaction").
		Preload("DstTransaction.DstTransfer").
		Preload("Token").
		Preload("Token.TokenBasic").
		Preload("FeeToken").
		Order("src_transactions.time desc").
		Find(&srcPolyDstRelation)
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("transacion: %s does not exist", hash)
	}
	return srcPolyDstRelation, nil
}

func (c *TransactionController) getTransactionByDstHash(hash string) (*models.SrcPolyDstRelation, error) {
	srcPolyDstRelation := new(models.SrcPolyDstRelation)
	res := db.Table("dst_transactions").
		Select("src_transactions.hash as src_hash, poly_transactions.hash as poly_hash, dst_transactions.hash as dst_hash, src_transactions.chain_id as chain_id, src_transfers.asset as token_hash, wrapper_transactions.fee_token_hash as fee_token_hash").
		Where("dst_transactions.hash = ?", hash).
		Where("dst_transactions.standard = ?", 0).
		Joins("inner join poly_transactions on dst_transactions.poly_hash = poly_transactions.hash").
		Joins("inner join src_transactions on poly_transactions.src_hash = src_transactions.hash").
		Joins("left join wrapper_transactions on src_transactions.hash = wrapper_transactions.hash").
		Joins("left join src_transfers on src_transactions.hash = src_transfers.tx_hash").
		Preload("WrapperTransaction").
		Preload("SrcTransaction").
		Preload("SrcTransaction.SrcTransfer").
		Preload("PolyTransaction").
		Preload("DstTransaction").
		Preload("DstTransaction.DstTransfer").
		Preload("Token").
		Preload("Token.TokenBasic").
		Preload("FeeToken").
		Order("src_transactions.time desc").
		Find(&srcPolyDstRelation)
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("transacion: %s does not exist", hash)
	}
	return srcPolyDstRelation, nil
}

func (c *TransactionController) TransactionOfHash() {
	var transactionOfHashReq models.TransactionOfHashReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &transactionOfHashReq); err != nil {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
	}
	srcPolyDstRelation, err := c.getTransactionByHash(transactionOfHashReq.Hash)
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	if srcPolyDstRelation.SrcTransaction.DstChainId != basedef.O3_CROSSCHAIN_ID || srcPolyDstRelation.DstTransaction == nil {
		chains := make([]*models.Chain, 0)
		db.Model(&models.Chain{}).Find(&chains)
		chainsMap := make(map[uint64]*models.Chain)
		for _, chain := range chains {
			chainsMap[chain.ChainId] = chain
		}
		resp := models.MakeTransactionRsp(srcPolyDstRelation, chainsMap)
		if resp == nil {
			c.Data["json"] = "transaction does not exist"
			c.Ctx.ResponseWriter.WriteHeader(400)
		} else {
			c.Data["json"] = resp
		}
		c.ServeJSON()
		return
	}
	srcPolyDstRelation2, err := c.getTransactionByHash(srcPolyDstRelation.DstHash)
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	srcPolyDstRelation.DstHash = srcPolyDstRelation2.DstHash
	srcPolyDstRelation.DstTransaction = srcPolyDstRelation2.DstTransaction
	chains := make([]*models.Chain, 0)
	db.Model(&models.Chain{}).Find(&chains)
	chainsMap := make(map[uint64]*models.Chain)
	for _, chain := range chains {
		chainsMap[chain.ChainId] = chain
	}
	c.Data["json"] = models.MakeTransactionRsp(srcPolyDstRelation, chainsMap)
	c.ServeJSON()
}

func (c *TransactionController) TransactionOfCurve() {
	var transactionOfHashReq models.TransactionOfHashReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &transactionOfHashReq); err != nil {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	srcPolyDstRelation1, err := c.getTransactionByDstHash(transactionOfHashReq.Hash)
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	if srcPolyDstRelation1.SrcTransaction.DstChainId != basedef.O3_CROSSCHAIN_ID || srcPolyDstRelation1.DstTransaction == nil {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	srcPolyDstRelation2, err := c.getTransactionByHash(srcPolyDstRelation1.DstHash)
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	chains := make([]*models.Chain, 0)
	db.Model(&models.Chain{}).Find(&chains)
	chainsMap := make(map[uint64]*models.Chain)
	for _, chain := range chains {
		chainsMap[chain.ChainId] = chain
	}
	c.Data["json"] = models.MakeCurveTransactionRsp(srcPolyDstRelation1, srcPolyDstRelation2, chainsMap)
	c.ServeJSON()
}

func (c *TransactionController) TransactionsOfState() {
	var transactionsOfStateReq models.TransactionsOfStateReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &transactionsOfStateReq); err != nil {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
	}
	transactions := make([]*models.WrapperTransaction, 0)
	db.Where("status = ?", transactionsOfStateReq.State).Limit(transactionsOfStateReq.PageSize).Offset(transactionsOfStateReq.PageSize * transactionsOfStateReq.PageNo).Order("time asc").Find(&transactions)
	var transactionNum int64
	db.Model(&models.WrapperTransaction{}).Where("status = ?", transactionsOfStateReq.State).Count(&transactionNum)
	c.Data["json"] = models.MakeTransactionsOfStateRsp(transactionsOfStateReq.PageSize, transactionsOfStateReq.PageNo,
		(int(transactionNum)+transactionsOfStateReq.PageSize-1)/transactionsOfStateReq.PageSize, int(transactionNum), transactions)
	c.ServeJSON()
}

func (c *TransactionController) TransactionsOfUnfinished() {
	var transactionsOfUnfinishedReq models.TransactionsOfUnfinishedReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &transactionsOfUnfinishedReq); err != nil {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
	}
	srcPolyDstRelations := make([]*models.SrcPolyDstRelation, 0)
	tt := time.Now().Unix()
	res := db.Table("src_transactions").
		Select("src_transactions.hash as src_hash, poly_transactions.hash as poly_hash, dst_transactions.hash as dst_hash, src_transactions.chain_id as chain_id, src_transfers.asset as token_hash, wrapper_transactions.fee_token_hash as fee_token_hash").
		Where("dst_transactions.hash is null").
		Where("src_transactions.standard = ?", 0).
		Where("src_transactions.time > ?", tt-24*60*60*28).
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
		Limit(transactionsOfUnfinishedReq.PageSize).Offset(transactionsOfUnfinishedReq.PageSize * transactionsOfUnfinishedReq.PageNo).
		Order("src_transactions.time desc").
		Find(&srcPolyDstRelations)
	if res.Error != nil {
		c.Data["json"] = res.Error.Error()
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	var transactionNum int64
	db.Table("src_transactions").
		Where("dst_transactions.hash is null").
		Where("src_transactions.standard = ?", 0).
		Where("src_transactions.time > ?", tt-24*60*60*28).
		Joins("left join poly_transactions on src_transactions.hash = poly_transactions.src_hash").
		Joins("left join dst_transactions on poly_transactions.hash = dst_transactions.poly_hash").
		Joins("inner join wrapper_transactions on src_transactions.hash = wrapper_transactions.hash").
		Count(&transactionNum)
	c.Data["json"] = models.MakeTransactionOfUnfinishedRsp(transactionsOfUnfinishedReq.PageSize, transactionsOfUnfinishedReq.PageNo,
		(int(transactionNum)+transactionsOfUnfinishedReq.PageSize-1)/transactionsOfUnfinishedReq.PageSize, int(transactionNum), srcPolyDstRelations)
	c.ServeJSON()
}

func (c *TransactionController) TransactionsOfAsset() {
	var transactionsOfAssetReq models.TransactionsOfAssetReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &transactionsOfAssetReq); err != nil {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
	}
	srcPolyDstRelations := make([]*models.SrcPolyDstRelation, 0)
	res := db.Table("src_transactions").
		Select("src_transactions.hash as src_hash, poly_transactions.hash as poly_hash, dst_transactions.hash as dst_hash, src_transactions.chain_id as chain_id, src_transfers.asset as token_hash, wrapper_transactions.fee_token_hash as fee_token_hash").
		Where("src_transfers.asset = ?", transactionsOfAssetReq.Asset).
		Where("src_transfers.chain_id = ?", transactionsOfAssetReq.Chain).
		Where("src_transactions.standard = ?", 0).
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
		Limit(transactionsOfAssetReq.PageSize).Offset(transactionsOfAssetReq.PageSize * transactionsOfAssetReq.PageNo).
		Order("src_transactions.time desc").
		Find(&srcPolyDstRelations)
	if res.Error != nil {
		c.Data["json"] = res.Error.Error()
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}
	var transactionNum int64
	db.Table("src_transactions").
		Where("src_transfers.asset = ?", transactionsOfAssetReq.Asset).
		Where("src_transfers.chain_id = ?", transactionsOfAssetReq.Chain).
		Where("src_transactions.standard = ?", 0).
		Joins("left join src_transfers on src_transactions.hash = src_transfers.tx_hash").
		Joins("left join poly_transactions on src_transactions.hash = poly_transactions.src_hash").
		Joins("left join dst_transactions on poly_transactions.hash = dst_transactions.poly_hash").
		Count(&transactionNum)
	c.Data["json"] = models.MakeTransactionOfUnfinishedRsp(transactionsOfAssetReq.PageSize, transactionsOfAssetReq.PageNo,
		(int(transactionNum)+transactionsOfAssetReq.PageSize-1)/transactionsOfAssetReq.PageSize, int(transactionNum), srcPolyDstRelations)
	c.ServeJSON()
}
