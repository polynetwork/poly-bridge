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
	"io/ioutil"
	"net/http"
	"poly-bridge/basedef"
	"poly-bridge/cacheRedis"
	"poly-bridge/models"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"gorm.io/gorm"
)

type TransactionController struct {
	web.Controller
}

func (c *TransactionController) return400(message string) {
	c.Data["json"] = models.MakeErrorRsp(message)
	c.Ctx.ResponseWriter.WriteHeader(400)
	c.ServeJSON()
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
		u := db.Model(&models.SrcTransfer{}).Select("tx_hash as hash, asset as asset, fee_token_hash as fee_token_hash, src_transfers.chain_id as chain_id").Joins("left join wrapper_transactions on src_transfers.tx_hash = wrapper_transactions.hash").
			Where("`from` in ? or src_transfers.dst_user in ?", req.Addresses, req.Addresses).
			Where("src_transfers.asset in ?", req.Assets)

		if req.SrcChainId > 0 {
			u = u.Where("src_transfers.chain_id = ?", req.SrcChainId)
		}
		if req.DstChainId > 0 {
			u = u.Where("src_transfers.dst_chain_id = ?", req.DstChainId)
		}
		return tx.Table("(?) as u", u).
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
	db.Debug().Table("(?) as u", db.Model(&models.SrcTransfer{}).
		Select("tx_hash as hash, asset as asset, fee_token_hash as fee_token_hash, src_transfers.chain_id as chain_id").
		Joins("left join wrapper_transactions on src_transfers.tx_hash = wrapper_transactions.hash").
		Where("`from` in ? or src_transfers.dst_user in ?", transactionsOfAddressReq.Addresses, transactionsOfAddressReq.Addresses).
		Where("wrapper_transactions.hash is NOT NULL or src_transfers.chain_id = ?", basedef.RIPPLE_CROSSCHAIN_ID)).
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
	db.Model(&models.SrcTransfer{}).
		Joins("left join wrapper_transactions on src_transfers.tx_hash = wrapper_transactions.hash").
		Where("src_transfers.`from` in ? or src_transfers.dst_user in ?", transactionsOfAddressReq.Addresses, transactionsOfAddressReq.Addresses).
		Where("wrapper_transactions.hash is NOT NULL or src_transfers.chain_id = ?", basedef.RIPPLE_CROSSCHAIN_ID).
		Count(&transactionNum)
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

func (c *TransactionController) GetManualTxData() {
	var manualTxDataReq models.ManualTxDataReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &manualTxDataReq); err != nil {
		c.return400("request parameter is invalid!")
		return
	}
	if manualTxDataReq.PolyHash[:2] == "0x" || manualTxDataReq.PolyHash[:2] == "0X" {
		manualTxDataReq.PolyHash = manualTxDataReq.PolyHash[2:]
	}
	polyTransaction := new(models.PolyTransaction)
	res := db.Where("hash = ?", manualTxDataReq.PolyHash).First(polyTransaction)
	if res.RowsAffected == 0 {
		c.return400(fmt.Sprintf("%v is not polyhash", manualTxDataReq.PolyHash))
		return
	}
	if polyTransaction.DstChainId == basedef.ONT_CROSSCHAIN_ID || polyTransaction.DstChainId == basedef.NEO_CROSSCHAIN_ID || polyTransaction.DstChainId == basedef.NEO3_CROSSCHAIN_ID {
		c.return400(fmt.Sprintf("%v can not submit to dst chain", manualTxDataReq.PolyHash))
		return
	}
	var x string
	res = db.Model(&models.DstTransaction{}).Select("hash").Where("poly_hash = ?", manualTxDataReq.PolyHash).First(&x)
	if res.RowsAffected != 0 {
		c.return400(fmt.Sprintf("%v was submitted to dst chain", manualTxDataReq.PolyHash))
		return
	}
	manualTxDataResp := new(models.ManualTxDataResp)
	manualData, err := cacheRedis.Redis.GetManualTx(manualTxDataReq.PolyHash)
	if err == nil {
		if manualData == "" {
			c.return400(fmt.Sprintf("%v getManualData loading", manualTxDataReq.PolyHash))
			return
		}
		json.Unmarshal([]byte(manualData), manualTxDataResp)
		c.Data["json"] = manualTxDataResp
		c.ServeJSON()
		return
	}
	url := relayUrl + "/api/v1/composetx?hash=" + manualTxDataReq.PolyHash
	for i := 0; i < 2; i++ {
		resp, err := http.Get(url)
		if err != nil {
			logs.Error("getManualData polyhash:", manualTxDataReq.PolyHash, "err:", err)
			time.Sleep(time.Millisecond * 500)
			continue
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, manualTxDataResp)
		manualData = string(body)
		cacheRedis.Redis.SetManualTx(manualTxDataReq.PolyHash, manualData)

		c.Data["json"] = manualTxDataResp
		c.ServeJSON()
		return
	}
	c.return400(fmt.Sprintf("%v getManualData timeout", manualTxDataReq.PolyHash))
	return
}

func (c *TransactionController) TransactionsWithoutWrapper() {
	var txWithoutWrapperReq models.TxWithoutWrapperReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &txWithoutWrapperReq); err != nil {
		c.return400("request parameter is invalid!")
		return
	}
	srcTransfers := make([]*models.SrcTransfer, 0)
	db.Table("src_transfers").
		Where("src_transfers.from = ?", txWithoutWrapperReq.User).
		Where("src_transfers.chain_id = ?", txWithoutWrapperReq.ChainId).
		Where("src_transfers.standard = ?", 0).
		Where("wrapper_transactions.hash is NULL").
		Joins("left join wrapper_transactions on src_transfers.tx_hash = wrapper_transactions.hash and src_transfers.chain_id = wrapper_transactions.src_chain_id").
		Limit(txWithoutWrapperReq.PageSize).
		Offset(txWithoutWrapperReq.PageSize * txWithoutWrapperReq.PageNo).
		Order("src_transfers.id desc").
		Preload("Token").
		Find(&srcTransfers)
	var count int64
	db.Table("src_transfers").
		Where("src_transfers.from = ?", txWithoutWrapperReq.User).
		Where("src_transfers.chain_id = ?", txWithoutWrapperReq.ChainId).
		Where("src_transfers.standard = ?", 0).
		Where("wrapper_transactions.hash is NULL").
		Joins("left join wrapper_transactions on src_transfers.tx_hash = wrapper_transactions.hash and src_transfers.chain_id = wrapper_transactions.src_chain_id").
		Count(&count)
	c.Data["json"] = models.MakeTxWithoutWrapperRsp(txWithoutWrapperReq.PageSize, txWithoutWrapperReq.PageNo, srcTransfers, count)
	c.ServeJSON()
	return
}
