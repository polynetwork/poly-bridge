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
	"poly-bridge/basedef"
	"poly-bridge/models"
	nftdb "poly-bridge/nft_http/db"

	"github.com/beego/beego/v2/server/web"
)

type TransactionController struct {
	web.Controller
}

func (c *TransactionController) TransactionsOfAddress() {
	var req models.TransactionsOfAddressReq
	if !input(&c.Controller, &req) {
		return
	}
	if !checkPageSize(&c.Controller, req.PageSize, 10) {
		return
	}
	wrapTxs := make([]*models.WrapperTransaction, 0)
	//convert for neo3
	for _, addr := range req.Addresses {
		potentialNeoAddr := basedef.HexStringReverse(addr)
		req.Addresses = append(req.Addresses, potentialNeoAddr)
	}
	if req.State == -1 {
		db.Model(&models.WrapperTransaction{}).
			Where("standard = ? and (user in ? or dst_user in ? )", models.TokenTypeErc721, req.Addresses, req.Addresses).
			Limit(req.PageSize).Offset(req.PageSize * req.PageNo).
			Order("time desc").
			Find(&wrapTxs)
	} else {
		db.Model(&models.WrapperTransaction{}).
			Where("standard = ? and status = ? and (user in ? or dst_user in ?)", models.TokenTypeErc721, req.State, req.Addresses, req.Addresses).
			Limit(req.PageSize).Offset(req.PageSize * req.PageNo).
			Order("time desc").
			Find(&wrapTxs)
	}

	list := findSrcPolyDstRelation(wrapTxs)

	// get transaction number
	var transactionNum int64
	db.Model(&models.SrcTransfer{}).
		Joins("inner join wrapper_transactions on src_transfers.tx_hash = wrapper_transactions.hash").
		Where("src_transfers.standard = ? and (`from` in ? or src_transfers.dst_user in ?)", models.TokenTypeErc721, req.Addresses, req.Addresses).
		Count(&transactionNum)

	// get chains
	chains := make([]*models.Chain, 0)
	db.Model(&models.Chain{}).Find(&chains)
	chainsMap := make(map[uint64]*models.Chain)
	for _, chain := range chains {
		chainsMap[chain.ChainId] = chain
	}

	totalPage := (int(transactionNum) + req.PageSize - 1) / req.PageSize
	totalCnt := int(transactionNum)

	data := new(TransactionsOfAddressRsp).instance(req.PageSize, req.PageNo, totalPage, totalCnt, list, chainsMap)
	output(&c.Controller, data)
}

func (c *TransactionController) TransactionOfHash() {
	var req models.TransactionOfHashReq
	if !input(&c.Controller, &req) {
		return
	}

	wrapTx := new(models.WrapperTransaction)
	res := db.Model(&models.WrapperTransaction{}).
		Where("standard = ? and hash = ?", models.TokenTypeErc721, req.Hash).
		Find(&wrapTx)

	if res.RowsAffected == 0 {
		notExist(&c.Controller)
		return
	}

	wrapTxs := []*models.WrapperTransaction{wrapTx}
	list := findSrcPolyDstRelation(wrapTxs)
	if len(list) == 0 {
		notExist(&c.Controller)
		return
	}

	chains := make([]*models.Chain, 0)
	db.Model(&models.Chain{}).Find(&chains)
	chainsMap := make(map[uint64]*models.Chain)
	for _, chain := range chains {
		chainsMap[chain.ChainId] = chain
	}

	data := new(TransactionRsp).instance(list[0], chainsMap)
	output(&c.Controller, data)
}

func (c *TransactionController) TransactionsOfState() {
	var req models.TransactionsOfStateReq
	if !input(&c.Controller, &req) {
		return
	}
	if !checkPageSize(&c.Controller, req.PageSize, 10) {
		return
	}

	transactions := make([]*models.WrapperTransaction, 0)
	db.Where("standard = ? and status = ?", models.TokenTypeErc721, req.State).
		Limit(req.PageSize).
		Offset(req.PageSize * req.PageNo).
		Order("time asc").
		Find(&transactions)

	var transactionNum int64
	db.Model(&models.WrapperTransaction{}).
		Where("standard = ? and status = ?", models.TokenTypeErc721, req.State).
		Count(&transactionNum)

	totalPage := (int(transactionNum) + req.PageSize - 1) / req.PageSize
	totalCount := int(transactionNum)
	data := models.MakeTransactionsOfStateRsp(req.PageSize, req.PageNo, totalPage, totalCount, transactions)
	output(&c.Controller, data)
}

func findSrcPolyDstRelation(wrapTxs []*models.WrapperTransaction) []*SrcPolyDstRelation {
	list := make([]*SrcPolyDstRelation, len(wrapTxs))
	srcTxHashs := make([]string, 0)
	for i, v := range wrapTxs {
		data := new(SrcPolyDstRelation)
		data.SrcHash = v.Hash
		data.WrapperTransaction = v
		data.FeeTokenHash = v.FeeTokenHash
		data.ChainId = v.SrcChainId
		list[i] = data
		srcTxHashs = append(srcTxHashs, v.Hash)
	}

	srcTxs := make([]*models.SrcTransaction, 0)
	srcTxsMap := make(map[string]*models.SrcTransaction)
	db.Model(&models.SrcTransaction{}).
		Where("src_transactions.hash in ?", srcTxHashs).
		Joins("left join src_transfers on src_transfers.tx_hash = src_transactions.hash").
		Preload("SrcTransfer").
		Find(&srcTxs)
	for _, v := range srcTxs {
		srcTxsMap[v.Hash] = v
	}

	polyTxs := make([]*models.PolyTransaction, 0)
	polyMap := make(map[string]*models.PolyTransaction)
	polyDstTxHashs := make([]string, 0)
	db.Model(&models.PolyTransaction{}).
		Where("src_hash in ?", srcTxHashs).
		Find(&polyTxs)
	for _, v := range polyTxs {
		polyMap[v.SrcHash] = v
		polyDstTxHashs = append(polyDstTxHashs, v.Hash)
	}

	dstTxs := make([]*models.DstTransaction, 0)
	dstTxsMap := make(map[string]*models.DstTransaction)
	db.Model(&models.DstTransaction{}).
		Where("dst_transactions.poly_hash in ?", polyDstTxHashs).
		Joins("left join dst_transfers on dst_transfers.tx_hash = dst_transactions.hash").
		Preload("DstTransfer").
		Find(&dstTxs)
	for _, v := range dstTxs {
		dstTxsMap[v.PolyHash] = v
	}

	for _, v := range list {
		if srcTx, ok := srcTxsMap[v.SrcHash]; ok {
			v.SrcTransaction = srcTx
		}
		if polyTx, ok := polyMap[v.SrcHash]; ok {
			v.PolyHash = polyTx.Hash
			v.PolyTransaction = polyTx
			v.DstHash = polyTx.Hash
		}
		if dstTx, ok := dstTxsMap[v.DstHash]; ok {
			v.DstTransaction = dstTx
			v.DstHash = dstTx.Hash
		}
		feeToken := nftdb.FindFeeToken(v.WrapperTransaction.SrcChainId, v.WrapperTransaction.FeeTokenHash)
		if feeToken != nil {
			v.FeeTokenHash = feeToken.Hash
			v.FeeToken = feeToken
		}

		srcAssetHash := v.SrcTransaction.SrcTransfer.Asset
		srcAsset := findAsset(v.WrapperTransaction.SrcChainId, srcAssetHash)
		if srcAsset != nil {
			v.SrcAsset = srcAsset
		}

		if v.DstTransaction != nil && v.DstTransaction.DstTransfer != nil {
			dstAssetHash := v.DstTransaction.DstTransfer.Asset
			dstAsset := findAsset(v.WrapperTransaction.DstChainId, dstAssetHash)
			if dstAsset != nil {
				v.DstAsset = dstAsset
			}
		}
	}

	return list
}
