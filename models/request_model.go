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

package models

type PolySwapResp struct {
	Version string
	URL     string
}

type TokenBasicReq struct {
	Name string
}

type TokenBasicRsp struct {
	Name         string
	Precision    uint64
	Price        int64
	Ind          uint64
	Time         int64
	PriceMarkets []*PriceMarketRsp
	Tokens       []*TokenRsp
}

func MakeTokenBasicRsp(tokenBasic *TokenBasic) *TokenBasicRsp {
	tokenBasicRsp := &TokenBasicRsp{
		Name:   tokenBasic.Name,
		Time:   tokenBasic.Time,
		Tokens: nil,
	}
	if tokenBasic.Tokens != nil {
		for _, token := range tokenBasic.Tokens {
			tokenBasicRsp.Tokens = append(tokenBasicRsp.Tokens, MakeTokenRsp(token))
		}
	}
	if tokenBasic.PriceMarkets != nil {
		for _, priceMarket := range tokenBasic.PriceMarkets {
			tokenBasicRsp.PriceMarkets = append(tokenBasicRsp.PriceMarkets, MakePriceMarketRsp(priceMarket))
		}
	}
	return tokenBasicRsp
}

type TokenReq struct {
	Hash string
}

type TokenRsp struct {
	Hash           string
	ChainId        uint64
	Name           string
	TokenBasicName string
	TokenBasic     *TokenBasicRsp
	TokenMaps      []*TokenMapRsp
}

func MakeTokenRsp(token *Token) *TokenRsp {
	tokenRsp := &TokenRsp{
		Hash:           token.Hash,
		ChainId:        token.ChainId,
		Name:           token.Name,
		TokenBasicName: token.TokenBasicName,
	}
	if token.TokenBasic != nil {
		tokenRsp.TokenBasic = MakeTokenBasicRsp(token.TokenBasic)
	}
	if token.TokenMaps != nil {
		for _, tokenmap := range token.TokenMaps {
			tokenRsp.TokenMaps = append(tokenRsp.TokenMaps, MakeTokenMapRsp(tokenmap))
		}
	}
	return tokenRsp
}

type PriceMarketRsp struct {
	TokenBasicName string
	MarketName     string
	Name           string
	Price          int64
	Ind            uint64
	Time           int64
	TokenBasic     *TokenBasicRsp
}

func MakePriceMarketRsp(priceMarket *PriceMarket) *PriceMarketRsp {
	priceMarketRsp := &PriceMarketRsp{
		TokenBasicName: priceMarket.TokenBasicName,
		MarketName:     priceMarket.MarketName,
		Name:           priceMarket.Name,
		Price:          priceMarket.Price,
		Ind:            priceMarket.Ind,
		Time:           priceMarket.Time,
	}
	if priceMarket.TokenBasic != nil {
		priceMarketRsp.TokenBasic = MakeTokenBasicRsp(priceMarket.TokenBasic)
	}
	return priceMarketRsp
}

type TokensReq struct {
	ChainId uint64
}

type TokensRsp struct {
	TotalCount uint64
	Tokens     []*TokenRsp
}

func MakeTokensRsp(tokens []*Token) *TokensRsp {
	tokensRsp := &TokensRsp{
		TotalCount: uint64(len(tokens)),
	}
	for _, token := range tokens {
		tokensRsp.Tokens = append(tokensRsp.Tokens, MakeTokenRsp(token))
	}
	return tokensRsp
}

type TokenMapReq struct {
	Hash string
}

type TokenMapRsp struct {
	SrcTokenHash string
	SrcToken     *TokenRsp
	DstTokenHash string
	DstToken     *TokenRsp
}

func MakeTokenMapRsp(tokenMap *TokenMap) *TokenMapRsp {
	tokenMapRsp := &TokenMapRsp{
		SrcTokenHash: tokenMap.SrcTokenHash,
		DstTokenHash: tokenMap.DstTokenHash,
	}
	if tokenMap.SrcToken != nil {
		tokenMapRsp.SrcToken = MakeTokenRsp(tokenMap.SrcToken)
	}
	if tokenMap.DstToken != nil {
		tokenMapRsp.DstToken = MakeTokenRsp(tokenMap.DstToken)
	}
	return tokenMapRsp
}

type GetFeeReq struct {
	ChainId uint64
	Hash    string
}

type GetFeeRsp struct {
	ChainId uint64
	Hash    string
	Amount  float64
}

func MakeGetFeeRsp(chainId uint64, hash string, amount float64) *GetFeeRsp {
	getFeeRsp := &GetFeeRsp{
		ChainId: chainId,
		Hash:    hash,
		Amount:  amount,
	}
	return getFeeRsp
}

type CheckFeeReq struct {
	Hash string
}

type CheckFeeRsp struct {
	HasPay bool
	Amount float64
}

func MakeCheckFeeRsp(hashPay bool, amount float64) *CheckFeeRsp {
	checkFeeRsp := &CheckFeeRsp{
		HasPay: hashPay,
		Amount: amount,
	}
	return checkFeeRsp
}

type TransactionReq struct {
	Hash string
}

type TransactionRsp struct {
	Hash         string
	User         string
	SrcChainId   uint64
	BlockHeight  uint64
	Time         uint64
	DstChainId   uint64
	FeeTokenHash string
	FeeAmount    uint64
}

func MakeTransactionRsp(transaction *WrapperTransaction) *TransactionRsp {
	transactionRsp := &TransactionRsp{
		Hash:         transaction.Hash,
		User:         transaction.User,
		SrcChainId:   transaction.SrcChainId,
		BlockHeight:  transaction.BlockHeight,
		Time:         transaction.Time,
		DstChainId:   transaction.DstChainId,
		FeeTokenHash: transaction.FeeTokenHash,
		FeeAmount:    transaction.FeeAmount.Uint64(),
	}
	return transactionRsp
}

type TransactionsReq struct {
	PageSize int
	PageNo   int
}

type TransactionsRsp struct {
	PageSize     int
	PageNo       int
	TotalPage    int
	TotalCount   int
	Transactions []*TransactionRsp
}

func MakeTransactionsRsp(pageSize int, pageNo int, totalPage int, totalCount int, transactions []*WrapperTransaction) *TransactionsRsp {
	transactionsRsp := &TransactionsRsp{
		PageSize:   pageSize,
		PageNo:     pageNo,
		TotalPage:  totalPage,
		TotalCount: totalCount,
	}
	for _, transaction := range transactions {
		transactionsRsp.Transactions = append(transactionsRsp.Transactions, MakeTransactionRsp(transaction))
	}
	return transactionsRsp
}

type TransactionsOfUserReq struct {
	User     string
	PageSize int
	PageNo   int
}

type TransactionsOfUserRsp struct {
	PageSize     int
	PageNo       int
	TotalPage    int
	TotalCount   int
	Transactions []*TransactionRsp
}

func MakeTransactionsOfUserRsp(pageSize int, pageNo int, totalPage int, totalCount int, transactions []*WrapperTransaction) *TransactionsRsp {
	transactionsRsp := &TransactionsRsp{
		PageSize:   pageSize,
		PageNo:     pageNo,
		TotalPage:  totalPage,
		TotalCount: totalCount,
	}
	for _, transaction := range transactions {
		transactionsRsp.Transactions = append(transactionsRsp.Transactions, MakeTransactionRsp(transaction))
	}
	return transactionsRsp
}

type TransactionsOfStateReq struct {
	State    string
	PageSize int
	PageNo   int
}

type TransactionsOfStateRsp struct {
	PageSize     int
	PageNo       int
	TotalPage    int
	TotalCount   int
	Transactions []*TransactionRsp
}

func MakeTransactionsOfStateRsp(pageSize int, pageNo int, totalPage int, totalCount int, transactions []*WrapperTransaction) *TransactionsRsp {
	transactionsRsp := &TransactionsRsp{
		PageSize:   pageSize,
		PageNo:     pageNo,
		TotalPage:  totalPage,
		TotalCount: totalCount,
	}
	for _, transaction := range transactions {
		transactionsRsp.Transactions = append(transactionsRsp.Transactions, MakeTransactionRsp(transaction))
	}
	return transactionsRsp
}

type AddressReq struct {
	ChainId     uint64
	AddressHash string
}

type AddressRsp struct {
	AddressHash string
	Address     string
	ChainId     uint64
}

func MakeAddressRsp(addressHash string, chainId uint64, address string) *AddressRsp {
	addressRsp := &AddressRsp{
		AddressHash: addressHash,
		Address:     address,
		ChainId:     chainId,
	}
	return addressRsp
}
