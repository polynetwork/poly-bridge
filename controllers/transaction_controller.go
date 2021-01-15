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

type TransactionController struct {
	beego.Controller
}

func (c *TransactionController) Transactions() {
	var transactionsReq models.WrapperTransactionsReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &transactionsReq); err != nil {
		panic(err)
	}
	transactions := make([]*models.WrapperTransaction, 0)
	db.Limit(transactionsReq.PageSize).Offset(transactionsReq.PageSize * transactionsReq.PageNo).Order("time asc").Find(&transactions)
	var transactionNum int64
	db.Model(&models.WrapperTransaction{}).Count(&transactionNum)
	c.Data["json"] = models.MakeTransactionsRsp(transactionsReq.PageSize, transactionsReq.PageNo, (int(transactionNum)+transactionsReq.PageSize-1)/transactionsReq.PageSize,
		int(transactionNum), transactions)
	c.ServeJSON()
}

func (c *TransactionController) TransactionsOfAddress() {
	var transactionsOfAddressReq models.TransactionsOfAddressReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &transactionsOfAddressReq); err != nil {
		panic(err)
	}
	srcPolyDstRelations := make([]*models.SrcPolyDstRelation, 0)
	db.Table("(?) as u", db.Model(&models.SrcTransfer{}).Select("tx_hash as hash, asset as asset").Where("`from` in ? or dst_user in ?", transactionsOfAddressReq.Addresses, transactionsOfAddressReq.Addresses)).
		Select("src_transactions.hash as src_hash, poly_transactions.hash as poly_hash, dst_transactions.hash as dst_hash, src_transactions.chain_id as chain_id, u.asset as token_hash").
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
		Limit(transactionsOfAddressReq.PageSize).Offset(transactionsOfAddressReq.PageSize * transactionsOfAddressReq.PageNo).
		Order("src_transactions.time desc").
		Find(&srcPolyDstRelations)
	var transactionNum int64
	db.Model(&models.SrcTransfer{}).Where("`from` in ? or dst_user in ?", transactionsOfAddressReq.Addresses, transactionsOfAddressReq.Addresses).Count(&transactionNum)
	chains := make([]*models.Chain, 0)
	db.Model(&models.Chain{}).Find(&chains)
	chainsMap := make(map[uint64]*models.Chain)
	for _, chain := range chains {
		chainsMap[*chain.ChainId] = chain
	}
	c.Data["json"] = models.MakeTransactionsOfUserRsp(transactionsOfAddressReq.PageSize, transactionsOfAddressReq.PageNo,
		(int(transactionNum)+transactionsOfAddressReq.PageSize-1)/transactionsOfAddressReq.PageSize, int(transactionNum), srcPolyDstRelations, chainsMap)
	c.ServeJSON()
}

func (c *TransactionController) TransactionOfHash() {
	var transactionOfHashReq models.TransactionOfHashReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &transactionOfHashReq); err != nil {
		panic(err)
	}
	srcPolyDstRelation := new(models.SrcPolyDstRelation)
	db.Table("src_transactions").
		Select("src_transactions.hash as src_hash, poly_transactions.hash as poly_hash, dst_transactions.hash as dst_hash, src_transactions.chain_id as chain_id, src_transfers.asset as token_hash").
		Where("src_transactions.hash = ?", transactionOfHashReq.Hash).
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
		Order("src_transactions.time desc").
		Find(srcPolyDstRelation)
	chains := make([]*models.Chain, 0)
	db.Model(&models.Chain{}).Find(&chains)
	chainsMap := make(map[uint64]*models.Chain)
	for _, chain := range chains {
		chainsMap[*chain.ChainId] = chain
	}
	c.Data["json"] = models.MakeTransactionRsp(srcPolyDstRelation, chainsMap)
	c.ServeJSON()
}

func (c *TransactionController) TransactionsOfState() {
	var transactionsOfStateReq models.TransactionsOfStateReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &transactionsOfStateReq); err != nil {
		panic(err)
	}
	transactions := make([]*models.WrapperTransaction, 0)
	db.Where("status = ?", transactionsOfStateReq.State).Limit(transactionsOfStateReq.PageSize).Offset(transactionsOfStateReq.PageSize * transactionsOfStateReq.PageNo).Order("time asc").Find(&transactions)
	var transactionNum int64
	db.Model(&models.WrapperTransaction{}).Where("status = ?", transactionsOfStateReq.State).Count(&transactionNum)
	c.Data["json"] = models.MakeTransactionsOfStateRsp(transactionsOfStateReq.PageSize, transactionsOfStateReq.PageNo,
		(int(transactionNum)+transactionsOfStateReq.PageSize-1)/transactionsOfStateReq.PageSize, int(transactionNum), transactions)
	c.ServeJSON()
}
