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

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/utils/decimal"
	"strings"
	"time"
)

type PolyBridgeResp struct {
	Version string
	URL     string
}

type ErrorRsp struct {
	Message string
}

func MakeErrorRsp(messgae string) *ErrorRsp {
	errorRsp := &ErrorRsp{
		Message: messgae,
	}
	return errorRsp
}

type TokenBasicReq struct {
	Name string
}

type TokenBasicRsp struct {
	Name         string
	Precision    uint64
	Price        string
	Ind          uint64
	Time         int64
	Property     int64
	Meta         string
	PriceMarkets []*PriceMarketRsp
	Tokens       []*TokenRsp
}

type TxHashChainIdPair struct {
	SrcHash    string
	PolyHash   string
	DstHash    string
	SrcChainId uint64
	DstChainId uint64
	WrapperId  uint64
}

func MakeTokenBasicRsp(tokenBasic *TokenBasic) *TokenBasicRsp {
	price := new(big.Float).Quo(new(big.Float).SetInt64(tokenBasic.Price), new(big.Float).SetInt64(basedef.PRICE_PRECISION))
	tokenBasicRsp := &TokenBasicRsp{
		Name:      tokenBasic.Name,
		Time:      tokenBasic.Time,
		Precision: tokenBasic.Precision,
		Meta:      tokenBasic.Meta,
		Price:     price.String(),
		Ind:       tokenBasic.Ind,
		Property:  tokenBasic.Property,
		Tokens:    nil,
	}
	if tokenBasic.Tokens != nil {
		for _, token := range tokenBasic.Tokens {
			if token.Property == 1 {
				tokenBasicRsp.Tokens = append(tokenBasicRsp.Tokens, MakeTokenRsp(token))
			}
		}
	}
	if tokenBasic.PriceMarkets != nil {
		for _, priceMarket := range tokenBasic.PriceMarkets {
			tokenBasicRsp.PriceMarkets = append(tokenBasicRsp.PriceMarkets, MakePriceMarketRsp(priceMarket))
		}
	}
	return tokenBasicRsp
}

type TokenBasicsReq struct {
}

type TokenBasicsRsp struct {
	TotalCount  uint64
	TokenBasics []*TokenBasicRsp
}

func MakeTokenBasicsRsp(tokenBasics []*TokenBasic) *TokenBasicsRsp {
	tokenBasicsRsp := &TokenBasicsRsp{
		TotalCount: uint64(len(tokenBasics)),
	}
	for _, tokenBasic := range tokenBasics {
		if tokenBasic.Property == 1 {
			tokenBasicsRsp.TokenBasics = append(tokenBasicsRsp.TokenBasics, MakeTokenBasicRsp(tokenBasic))
		}
	}
	return tokenBasicsRsp
}

type TokenBasicsInfoReq struct {
	PageSize int
	PageNo   int
	Order    string
}

type TokenBasicsInfoRsp struct {
	PageSize    int
	PageNo      int
	TotalPage   uint
	TotalCount  uint64
	TokenBasics []*TokenBasicInfoRsp
}

type TokenBasicInfoRsp struct {
	TokenBasicRsp  `json:",inline"`
	TotalAmount    string
	TotalVolume    string
	TotalCount     uint64
	SocialTwitter  string
	SocialTelegram string
	SocialWebsite  string
	SocialOther    string
}

func MakeTokenBasicsInfoRsp(req *TokenBasicsInfoReq, count uint64, tokenBasics []*TokenBasic) *TokenBasicsInfoRsp {
	pages := int(count) / req.PageSize
	if int(count)%req.PageSize != 0 {
		pages++
	}
	tokenBasicsRsp := &TokenBasicsInfoRsp{
		TotalCount: count,
		PageSize:   len(tokenBasics),
		PageNo:     req.PageNo,
		TotalPage:  uint(pages),
	}
	for _, tokenBasic := range tokenBasics {
		info := &TokenBasicInfoRsp{TokenBasicRsp: *MakeTokenBasicRsp(tokenBasic)}
		if tokenBasic.TotalAmount == nil {
			info.TotalAmount = "0"
			info.TotalVolume = "0"
		} else {
			info.TotalAmount = tokenBasic.TotalAmount.String()
			volume := new(big.Int).Mul(&tokenBasic.TotalAmount.Int, big.NewInt(tokenBasic.Price))
			if tokenBasic.Precision > 0 {
				volume = new(big.Int).Quo(volume, new(big.Int).SetInt64(basedef.Int64FromFigure(int(tokenBasic.Precision))))
			}
			info.TotalVolume = new(big.Int).Quo(volume, big.NewInt(basedef.PRICE_PRECISION)).String()
		}
		info.TotalCount = tokenBasic.TotalCount
		info.SocialTwitter = tokenBasic.SocialTwitter
		info.SocialTelegram = tokenBasic.SocialTelegram
		info.SocialWebsite = tokenBasic.SocialWebsite
		info.SocialOther = tokenBasic.SocialOther
		tokenBasicsRsp.TokenBasics = append(tokenBasicsRsp.TokenBasics, info)
	}
	return tokenBasicsRsp
}

type TokenReq struct {
	ChainId uint64
	Hash    string
}

type TokenRsp struct {
	Hash            string
	ChainId         uint64
	Name            string
	Property        int64
	TokenBasicName  string
	Precision       uint64
	AvailableAmount string
	TokenBasic      *TokenBasicRsp
	TokenMaps       []*TokenMapRsp
}

func MakeTokenRsp(token *Token) *TokenRsp {
	tokenRsp := &TokenRsp{
		Hash:           token.Hash,
		ChainId:        token.ChainId,
		Name:           token.Name,
		TokenBasicName: token.TokenBasicName,
		Property:       token.Property,
		Precision:      token.Precision,
	}
	if token.AvailableAmount != nil {
		tokenRsp.AvailableAmount = token.AvailableAmount.String()
	}
	if token.TokenBasic != nil {
		tokenRsp.TokenBasic = MakeTokenBasicRsp(token.TokenBasic)
	}
	if token.TokenMaps != nil {
		for _, tokenmap := range token.TokenMaps {
			if tokenmap.Property == 1 {
				tokenRsp.TokenMaps = append(tokenRsp.TokenMaps, MakeTokenMapRsp(tokenmap))
			}
		}
	}
	return tokenRsp
}

type PriceMarketRsp struct {
	TokenBasicName string
	MarketName     string
	Name           string
	Price          string
	Ind            uint64
	Time           int64
	TokenBasic     *TokenBasicRsp
}

func MakePriceMarketRsp(priceMarket *PriceMarket) *PriceMarketRsp {
	price := new(big.Float).Quo(new(big.Float).SetInt64(priceMarket.Price), new(big.Float).SetInt64(basedef.PRICE_PRECISION))
	priceMarketRsp := &PriceMarketRsp{
		TokenBasicName: priceMarket.TokenBasicName,
		MarketName:     priceMarket.MarketName,
		Name:           priceMarket.Name,
		Price:          price.String(),
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
	ChainId uint64
	Hash    string
}

type TokenMapRsp struct {
	SrcTokenHash string
	SrcToken     *TokenRsp
	DstTokenHash string
	DstToken     *TokenRsp
	Property     int64
}

func MakeTokenMapRsp(tokenMap *TokenMap) *TokenMapRsp {
	tokenMapRsp := &TokenMapRsp{
		SrcTokenHash: tokenMap.SrcTokenHash,
		DstTokenHash: tokenMap.DstTokenHash,
		Property:     tokenMap.Property,
	}
	if tokenMap.SrcToken != nil {
		tokenMapRsp.SrcToken = MakeTokenRsp(tokenMap.SrcToken)
	}
	if tokenMap.DstToken != nil {
		tokenMapRsp.DstToken = MakeTokenRsp(tokenMap.DstToken)
	}
	return tokenMapRsp
}

type TokenMapsReq struct {
	ChainId uint64
	Hash    string
}

type TokenMapsRsp struct {
	TotalCount uint64
	TokenMaps  []*TokenMapRsp
}

func MakeTokenMapsRsp(tokenMaps []*TokenMap) *TokenMapsRsp {
	tokenMapsRsp := &TokenMapsRsp{
		TotalCount: uint64(len(tokenMaps)),
	}
	for _, tokenMap := range tokenMaps {
		if tokenMap.Property == 1 {
			tokenMapsRsp.TokenMaps = append(tokenMapsRsp.TokenMaps, MakeTokenMapRsp(tokenMap))
		}
	}
	return tokenMapsRsp
}

type GetFeeReq struct {
	SrcChainId    uint64
	Hash          string
	DstChainId    uint64
	SwapTokenHash string
}

type GetFeeRsp struct {
	SrcChainId               uint64
	Hash                     string
	DstChainId               uint64
	UsdtAmount               string
	TokenAmount              string
	TokenAmountWithPrecision string
	SwapTokenHash            string
	Balance                  string
	BalanceWithPrecision     string
	IsNative                 bool
	NativeTokenAmount        string
}

func MakeGetFeeRsp(srcChainId uint64, hash string, dstChainId uint64, usdtAmount *big.Float, tokenAmount *big.Float, tokenAmountWithPrecision *big.Float,
	swapTokenHash string, balance *big.Float, balanceWithoutPrecision *big.Float, isNative bool, nativeTokenAmount *big.Float) *GetFeeRsp {
	getFeeRsp := &GetFeeRsp{
		SrcChainId:               srcChainId,
		Hash:                     hash,
		DstChainId:               dstChainId,
		UsdtAmount:               fmt.Sprintf("%v", usdtAmount),
		TokenAmount:              tokenAmount.String(),
		TokenAmountWithPrecision: fmt.Sprintf("%v", tokenAmountWithPrecision),
		SwapTokenHash:            swapTokenHash,
		Balance:                  fmt.Sprintf("%v", balanceWithoutPrecision),
		BalanceWithPrecision:     fmt.Sprintf("%v", balance),
		IsNative:                 isNative,
		NativeTokenAmount:        fmt.Sprintf("%v", nativeTokenAmount),
	}
	{
		precision := decimal.NewFromInt(basedef.PRICE_PRECISION)
		aaa := new(big.Float).Mul(tokenAmount, new(big.Float).SetInt64(basedef.PRICE_PRECISION))
		bbb, _ := aaa.Float64()
		ccc := decimal.NewFromFloat(bbb)
		tokenAmount := ccc.Div(precision)
		getFeeRsp.TokenAmount = tokenAmount.String()
	}
	if getFeeRsp.DstChainId == basedef.BSC_CROSSCHAIN_ID {
		logs.Info("tobscgetfee srcChain:%v swapTokenHash:%v feeAmount:%v", srcChainId, swapTokenHash, getFeeRsp.TokenAmount)
	}
	return getFeeRsp
}

type CheckFeeReq struct {
	Hash    string
	ChainId uint64
}

type CheckFeeRsp struct {
	ChainId     uint64
	Hash        string
	PayState    int
	Amount      string
	MinProxyFee string
}

type CheckFeesReq struct {
	Checks []*CheckFeeReq
}

type CheckFeesRsp struct {
	TotalCount uint64
	CheckFees  []*CheckFeeRsp
}

func MakeCheckFeesRsp(checkFees []*CheckFee) *CheckFeesRsp {
	checkFeesRsp := &CheckFeesRsp{
		TotalCount: uint64(len(checkFees)),
	}
	for _, checkFee := range checkFees {
		checkFeesRsp.CheckFees = append(checkFeesRsp.CheckFees, MakeCheckFeeRsp(checkFee))
	}
	return checkFeesRsp
}

func MakeCheckFeeRsp(checkFee *CheckFee) *CheckFeeRsp {
	checkFeeRsp := &CheckFeeRsp{
		ChainId:     checkFee.ChainId,
		Hash:        checkFee.Hash,
		PayState:    checkFee.PayState,
		Amount:      checkFee.Amount.String(),
		MinProxyFee: checkFee.MinProxyFee.String(),
	}
	{
		aaa, _ := checkFee.Amount.Float64()
		bbb := decimal.NewFromFloat(aaa)
		checkFeeRsp.Amount = bbb.String()
	}
	{
		aaa, _ := checkFee.MinProxyFee.Float64()
		bbb := decimal.NewFromFloat(aaa)
		checkFeeRsp.MinProxyFee = bbb.String()
	}
	return checkFeeRsp
}

type WrapperTransactionReq struct {
	Hash string
}

type WrapperTransactionRsp struct {
	Hash         string
	User         string
	SrcChainId   uint64
	BlockHeight  uint64
	Time         uint64
	DstChainId   uint64
	DstUser      string
	ServerId     uint64
	FeeTokenHash string
	FeeAmount    string
	State        uint64
}

func MakeWrapperTransactionRsp(transaction *WrapperTransaction) *WrapperTransactionRsp {
	transactionRsp := &WrapperTransactionRsp{
		Hash:         transaction.Hash,
		User:         transaction.User,
		SrcChainId:   transaction.SrcChainId,
		BlockHeight:  transaction.BlockHeight,
		Time:         transaction.Time,
		DstChainId:   transaction.DstChainId,
		DstUser:      transaction.DstUser,
		ServerId:     transaction.ServerId,
		FeeTokenHash: transaction.FeeTokenHash,
		FeeAmount:    transaction.FeeAmount.String(),
		State:        transaction.Status,
	}
	return transactionRsp
}

type WrapperTransactionsReq struct {
	PageSize int
	PageNo   int
}

type WrapperTransactionsWithFilterReq struct {
	PageSize   int
	PageNo     int
	SrcChainId int
	DstChainId int
	Assets     []string
}

type WrapperTransactionsRsp struct {
	PageSize     int
	PageNo       int
	TotalPage    int
	TotalCount   int
	Transactions []*WrapperTransactionRsp
}

func MakeWrapperTransactionsRsp(pageSize int, pageNo int, totalPage int, totalCount int, transactions []*WrapperTransaction) *WrapperTransactionsRsp {
	transactionsRsp := &WrapperTransactionsRsp{
		PageSize:   pageSize,
		PageNo:     pageNo,
		TotalPage:  totalPage,
		TotalCount: totalCount,
	}
	for _, transaction := range transactions {
		transactionsRsp.Transactions = append(transactionsRsp.Transactions, MakeWrapperTransactionRsp(transaction))
	}
	return transactionsRsp
}

type TransactionOfHashReq struct {
	Hash string
}

type TransactionStateRsp struct {
	Hash       string
	ChainId    uint64
	Blocks     uint64
	NeedBlocks uint64
	Time       uint64
}

type TransactionRsp struct {
	Hash        string
	User        string
	SrcChainId  uint64
	BlockHeight uint64
	Time        uint64
	DstChainId  uint64
	//Amount           string
	FeeAmount        string
	TransferAmount   string
	DstUser          string
	ServerId         uint64
	State            uint64
	Token            *TokenRsp
	FeeToken         *TokenRsp
	TransactionState []*TransactionStateRsp
}

func MakeTransactionRsp(transaction *SrcPolyDstRelation, chainsMap map[uint64]*Chain) *TransactionRsp {
	if transaction.SrcTransaction == nil {
		return nil
	}
	if transaction.WrapperTransaction == nil {
		return MakeTransactionRspWithoutWrapper(transaction, chainsMap)
	}

	feeAmount := ""
	if transaction.WrapperTransaction != nil {
		aaa := new(big.Int).Set(&transaction.WrapperTransaction.FeeAmount.Int)
		feeAmount = aaa.String()
	}
	transferAmount := ""
	dstUser := ""
	if transaction.SrcTransaction.SrcTransfer != nil {
		aaa := new(big.Int).Set(&transaction.SrcTransaction.SrcTransfer.Amount.Int)
		transferAmount = aaa.String()
		dstUser = transaction.SrcTransaction.SrcTransfer.DstUser
	}
	transactionRsp := &TransactionRsp{
		Hash:           transaction.WrapperTransaction.Hash,
		User:           transaction.WrapperTransaction.User,
		SrcChainId:     transaction.WrapperTransaction.SrcChainId,
		BlockHeight:    transaction.WrapperTransaction.BlockHeight,
		Time:           transaction.WrapperTransaction.Time,
		DstChainId:     transaction.WrapperTransaction.DstChainId,
		ServerId:       transaction.WrapperTransaction.ServerId,
		FeeAmount:      feeAmount,
		TransferAmount: transferAmount,
		DstUser:        dstUser,
		State:          transaction.WrapperTransaction.Status,
	}
	if transaction.Token != nil {
		transactionRsp.Token = MakeTokenRsp(transaction.Token)
		precision := decimal.NewFromInt(basedef.Int64FromFigure(int(transaction.Token.Precision)))
		if transaction.SrcTransaction.SrcTransfer != nil {
			bbb := decimal.NewFromBigInt(&transaction.SrcTransaction.SrcTransfer.Amount.Int, 0)
			transferAmount := bbb.Div(precision)
			transactionRsp.TransferAmount = transferAmount.String()
		}
	}
	if transaction.FeeToken != nil {
		transactionRsp.FeeToken = MakeTokenRsp(transaction.FeeToken)
		precision := decimal.NewFromInt(basedef.Int64FromFigure(int(transaction.FeeToken.Precision)))
		if transaction.WrapperTransaction != nil {
			bbb := decimal.NewFromBigInt(&transaction.WrapperTransaction.FeeAmount.Int, 0)
			feeAmount := bbb.Div(precision)
			transactionRsp.FeeAmount = feeAmount.String()
		}
	}

	srcTransactionState := &TransactionStateRsp{
		Hash:    "",
		ChainId: transaction.WrapperTransaction.SrcChainId,
		Blocks:  0,
		Time:    0,
	}
	polyTransactionState := &TransactionStateRsp{
		Hash:    "",
		ChainId: 0,
		Blocks:  0,
		Time:    0,
	}
	dstTransactionState := &TransactionStateRsp{
		Hash:    "",
		ChainId: transaction.WrapperTransaction.DstChainId,
		Blocks:  0,
		Time:    0,
	}
	transactionRsp.TransactionState = append(transactionRsp.TransactionState, srcTransactionState)
	transactionRsp.TransactionState = append(transactionRsp.TransactionState, polyTransactionState)
	transactionRsp.TransactionState = append(transactionRsp.TransactionState, dstTransactionState)

	if transaction.SrcTransaction != nil {
		height := transaction.SrcTransaction.Height
		srcTransactionState.Hash = transaction.SrcTransaction.Hash
		srcTransactionState.ChainId = transaction.SrcTransaction.ChainId
		srcTransactionState.Time = transaction.SrcTransaction.Time

		srcChain, ok := chainsMap[srcTransactionState.ChainId]
		if ok {
			srcTransactionState.NeedBlocks = srcChain.BackwardBlockNumber
			srcTransactionState.Blocks = srcChain.Height - height
			if srcTransactionState.Blocks > srcTransactionState.NeedBlocks {
				srcTransactionState.Blocks = srcTransactionState.NeedBlocks
			}
		}

		switch transaction.ChainId {
		case basedef.ARBITRUM_CROSSCHAIN_ID:
			if l1BlockNumber, err := GetL1BlockNumberOfArbitrumTx(transaction.SrcTransaction.Hash); err == nil {
				height = l1BlockNumber
				ethChain, ok := chainsMap[basedef.ETHEREUM_CROSSCHAIN_ID]
				if ok {
					srcTransactionState.NeedBlocks = ethChain.BackwardBlockNumber
					if ethChain.Height < height {
						srcTransactionState.Blocks = 0
					} else {
						srcTransactionState.Blocks = ethChain.Height - height
					}
					if srcTransactionState.Blocks > srcTransactionState.NeedBlocks {
						srcTransactionState.Blocks = srcTransactionState.NeedBlocks
					}
				}
			} else {
				logs.Error("GetL1BlockNumberOfArbitrumTx failed. hash=%s, err:", transaction.SrcTransaction.Hash, err)
			}
		case basedef.ZKSYNC_CROSSCHAIN_ID:
			if zkChain, ok := chainsMap[basedef.ZKSYNC_CROSSCHAIN_ID]; ok {
				l1ChainId := basedef.ETHEREUM_CROSSCHAIN_ID
				if basedef.ENV == basedef.TESTNET {
					l1ChainId = basedef.GOERLI_CROSSCHAIN_ID
				}
				l1Chain, _ := chainsMap[l1ChainId]
				if l1Height, err := GetZkSyncL1Height(zkChain, l1Chain); err != nil {
					srcTransactionState.NeedBlocks = srcChain.BackwardBlockNumber
					if l1Height < height {
						srcTransactionState.Blocks = 0
					} else {
						srcTransactionState.Blocks = l1Height - height
					}
					if srcTransactionState.Blocks > srcTransactionState.NeedBlocks {
						srcTransactionState.Blocks = srcTransactionState.NeedBlocks
					}
				}
			}
		}
	}
	if transaction.PolyTransaction != nil {
		polyTransactionState.Hash = transaction.PolyTransaction.Hash
		polyTransactionState.ChainId = transaction.PolyTransaction.ChainId
		polyTransactionState.Time = transaction.PolyTransaction.Time

		polyChain, ok := chainsMap[polyTransactionState.ChainId]
		if ok {
			polyTransactionState.NeedBlocks = polyChain.BackwardBlockNumber
			polyTransactionState.Blocks = polyChain.Height - transaction.PolyTransaction.Height
			if polyTransactionState.Blocks > polyTransactionState.NeedBlocks {
				polyTransactionState.Blocks = polyTransactionState.NeedBlocks
			}
		}

		srcChain, ok := chainsMap[srcTransactionState.ChainId]
		if ok {
			srcTransactionState.NeedBlocks = srcChain.BackwardBlockNumber
			srcTransactionState.Blocks = srcTransactionState.NeedBlocks
		}
	}
	if transaction.DstTransaction != nil {
		dstTransactionState.Hash = transaction.DstTransaction.Hash
		dstTransactionState.ChainId = transaction.DstTransaction.ChainId
		dstTransactionState.Time = transaction.DstTransaction.Time
		dstTransactionState.NeedBlocks = 1

		dstTransactionState.Blocks = transaction.DstTransaction.Height
		dstChain, ok := chainsMap[dstTransactionState.ChainId]
		if ok {
			dstTransactionState.Blocks = dstChain.Height - transaction.DstTransaction.Height
		}
		if dstTransactionState.Blocks > dstTransactionState.NeedBlocks {
			dstTransactionState.Blocks = dstTransactionState.NeedBlocks
		}

		polyChain, ok := chainsMap[polyTransactionState.ChainId]
		if ok {
			polyTransactionState.NeedBlocks = polyChain.BackwardBlockNumber
			polyTransactionState.Blocks = polyTransactionState.NeedBlocks
		}

		srcChain, ok := chainsMap[srcTransactionState.ChainId]
		if ok {
			srcTransactionState.NeedBlocks = srcChain.BackwardBlockNumber
			srcTransactionState.Blocks = srcTransactionState.NeedBlocks
		}

	}
	return transactionRsp
}

func MakeTransactionRspWithoutWrapper(transaction *SrcPolyDstRelation, chainsMap map[uint64]*Chain) *TransactionRsp {
	transferAmount := ""
	dstUser := ""
	if transaction.SrcTransaction.SrcTransfer != nil {
		aaa := new(big.Int).Set(&transaction.SrcTransaction.SrcTransfer.Amount.Int)
		transferAmount = aaa.String()
		dstUser = transaction.SrcTransaction.SrcTransfer.DstUser
	}
	transactionRsp := &TransactionRsp{
		Hash:           transaction.SrcHash,
		User:           transaction.SrcTransaction.User,
		SrcChainId:     transaction.SrcTransaction.ChainId,
		BlockHeight:    transaction.SrcTransaction.Height,
		Time:           transaction.SrcTransaction.Time,
		DstChainId:     transaction.SrcTransaction.DstChainId,
		TransferAmount: transferAmount,
		DstUser:        dstUser,
	}
	switch {
	case transaction.PolyTransaction == nil && transaction.DstTransaction == nil:
		transactionRsp.State = basedef.STATE_SOURCE_CONFIRMED
	case transaction.DstTransaction == nil:
		transactionRsp.State = basedef.STATE_POLY_CONFIRMED
	default:
		transactionRsp.State = basedef.STATE_FINISHED
	}
	switch transaction.SrcTransaction.ChainId {
	case basedef.RIPPLE_CROSSCHAIN_ID:
		transactionRsp.State = basedef.STATE_WITHOUT_WRAPPER
	}
	if transaction.Token != nil {
		transactionRsp.Token = MakeTokenRsp(transaction.Token)
		precision := decimal.NewFromInt(basedef.Int64FromFigure(int(transaction.Token.Precision)))
		if transaction.SrcTransaction.SrcTransfer != nil {
			bbb := decimal.NewFromBigInt(&transaction.SrcTransaction.SrcTransfer.Amount.Int, 0)
			transferAmount := bbb.Div(precision)
			transactionRsp.TransferAmount = transferAmount.String()
		}
	}

	srcTransactionState := &TransactionStateRsp{
		Hash:    "",
		ChainId: transaction.SrcTransaction.ChainId,
		Blocks:  0,
		Time:    0,
	}
	polyTransactionState := &TransactionStateRsp{
		Hash:    "",
		ChainId: 0,
		Blocks:  0,
		Time:    0,
	}
	dstTransactionState := &TransactionStateRsp{
		Hash:    "",
		ChainId: transaction.SrcTransaction.DstChainId,
		Blocks:  0,
		Time:    0,
	}
	transactionRsp.TransactionState = append(transactionRsp.TransactionState, srcTransactionState)
	transactionRsp.TransactionState = append(transactionRsp.TransactionState, polyTransactionState)
	transactionRsp.TransactionState = append(transactionRsp.TransactionState, dstTransactionState)

	if transaction.SrcTransaction != nil {
		height := transaction.SrcTransaction.Height
		srcTransactionState.Hash = transaction.SrcTransaction.Hash
		srcTransactionState.ChainId = transaction.SrcTransaction.ChainId
		srcTransactionState.Time = transaction.SrcTransaction.Time

		srcChain, ok := chainsMap[srcTransactionState.ChainId]
		if ok {
			srcTransactionState.NeedBlocks = srcChain.BackwardBlockNumber
			srcTransactionState.Blocks = srcChain.Height - height
			if srcTransactionState.Blocks > srcTransactionState.NeedBlocks {
				srcTransactionState.Blocks = srcTransactionState.NeedBlocks
			}
		}
		if transaction.ChainId == basedef.ARBITRUM_CROSSCHAIN_ID {
			if l1BlockNumber, err := GetL1BlockNumberOfArbitrumTx(transaction.SrcTransaction.Hash); err == nil {
				height = l1BlockNumber
				ethChain, ok := chainsMap[basedef.ETHEREUM_CROSSCHAIN_ID]
				if ok {
					srcTransactionState.NeedBlocks = ethChain.BackwardBlockNumber
					srcTransactionState.Blocks = ethChain.Height - height
					if srcTransactionState.Blocks > srcTransactionState.NeedBlocks {
						srcTransactionState.Blocks = srcTransactionState.NeedBlocks
					}
				}
			} else {
				logs.Error("GetL1BlockNumberOfArbitrumTx failed. hash=%s, err:", transaction.SrcTransaction.Hash, err)
			}
		}
	}
	if transaction.PolyTransaction != nil {
		polyTransactionState.Hash = transaction.PolyTransaction.Hash
		polyTransactionState.ChainId = transaction.PolyTransaction.ChainId
		polyTransactionState.Time = transaction.PolyTransaction.Time

		polyChain, ok := chainsMap[polyTransactionState.ChainId]
		if ok {
			polyTransactionState.NeedBlocks = polyChain.BackwardBlockNumber
			polyTransactionState.Blocks = polyChain.Height - transaction.PolyTransaction.Height
			if polyTransactionState.Blocks > polyTransactionState.NeedBlocks {
				polyTransactionState.Blocks = polyTransactionState.NeedBlocks
			}
		}

		srcChain, ok := chainsMap[srcTransactionState.ChainId]
		if ok {
			srcTransactionState.NeedBlocks = srcChain.BackwardBlockNumber
			srcTransactionState.Blocks = srcTransactionState.NeedBlocks
		}
	}
	if transaction.DstTransaction != nil {
		dstTransactionState.Hash = transaction.DstTransaction.Hash
		dstTransactionState.ChainId = transaction.DstTransaction.ChainId
		dstTransactionState.Time = transaction.DstTransaction.Time
		dstTransactionState.NeedBlocks = 1

		dstTransactionState.Blocks = transaction.DstTransaction.Height
		dstChain, ok := chainsMap[dstTransactionState.ChainId]
		if ok {
			dstTransactionState.Blocks = dstChain.Height - transaction.DstTransaction.Height
		}
		if dstTransactionState.Blocks > dstTransactionState.NeedBlocks {
			dstTransactionState.Blocks = dstTransactionState.NeedBlocks
		}

		polyChain, ok := chainsMap[polyTransactionState.ChainId]
		if ok {
			polyTransactionState.NeedBlocks = polyChain.BackwardBlockNumber
			polyTransactionState.Blocks = polyTransactionState.NeedBlocks
		}

		srcChain, ok := chainsMap[srcTransactionState.ChainId]
		if ok {
			srcTransactionState.NeedBlocks = srcChain.BackwardBlockNumber
			srcTransactionState.Blocks = srcTransactionState.NeedBlocks
		}
	}
	return transactionRsp
}

func MakeCurveTransactionRsp(transaction1 *SrcPolyDstRelation, transaction2 *SrcPolyDstRelation, chainsMap map[uint64]*Chain) *TransactionRsp {
	feeAmount := ""
	if transaction1.WrapperTransaction != nil {
		aaa := new(big.Int).Set(&transaction1.WrapperTransaction.FeeAmount.Int)
		feeAmount = aaa.String()
	}
	transferAmount := ""
	dstUser := ""
	if transaction1.SrcTransaction.SrcTransfer != nil {
		aaa := new(big.Int).Set(&transaction1.SrcTransaction.SrcTransfer.Amount.Int)
		transferAmount = aaa.String()
		dstUser = transaction1.SrcTransaction.SrcTransfer.DstUser
	}
	transactionRsp := &TransactionRsp{
		Hash:           transaction1.WrapperTransaction.Hash,
		User:           transaction1.WrapperTransaction.User,
		SrcChainId:     transaction1.WrapperTransaction.SrcChainId,
		BlockHeight:    transaction1.WrapperTransaction.BlockHeight,
		Time:           transaction1.WrapperTransaction.Time,
		DstChainId:     transaction1.WrapperTransaction.DstChainId,
		ServerId:       transaction1.WrapperTransaction.ServerId,
		FeeAmount:      feeAmount,
		TransferAmount: transferAmount,
		//Amount:         amount.String(),
		DstUser: dstUser,
		State:   transaction1.WrapperTransaction.Status,
	}
	if transaction1.Token != nil {
		transactionRsp.Token = MakeTokenRsp(transaction1.Token)
		precision := decimal.NewFromInt(basedef.Int64FromFigure(int(transaction1.Token.Precision)))
		if transaction1.SrcTransaction.SrcTransfer != nil {
			bbb := decimal.NewFromBigInt(&transaction1.SrcTransaction.SrcTransfer.Amount.Int, 0)
			transferAmount := bbb.Div(precision)
			transactionRsp.TransferAmount = transferAmount.String()
		}
	}
	if transaction1.FeeToken != nil {
		transactionRsp.FeeToken = MakeTokenRsp(transaction1.FeeToken)
		precision := decimal.NewFromInt(basedef.Int64FromFigure(int(transaction1.FeeToken.Precision)))
		if transaction1.WrapperTransaction != nil {
			bbb := decimal.NewFromBigInt(&transaction1.WrapperTransaction.FeeAmount.Int, 0)
			feeAmount := bbb.Div(precision)
			transactionRsp.FeeAmount = feeAmount.String()
		}
	}
	if transaction1.SrcTransaction != nil {
		transactionRsp.TransactionState = append(transactionRsp.TransactionState, &TransactionStateRsp{
			Hash:    transaction1.SrcTransaction.Hash,
			ChainId: transaction1.SrcTransaction.ChainId,
			Blocks:  transaction1.SrcTransaction.Height,
			Time:    transaction1.SrcTransaction.Time,
		})
	} else {
		transactionRsp.TransactionState = append(transactionRsp.TransactionState, &TransactionStateRsp{
			Hash:    "",
			ChainId: transaction1.WrapperTransaction.SrcChainId,
			Blocks:  0,
			Time:    0,
		})
	}
	if transaction1.PolyTransaction != nil {
		transactionRsp.TransactionState = append(transactionRsp.TransactionState, &TransactionStateRsp{
			Hash:    transaction1.PolyTransaction.Hash,
			ChainId: transaction1.PolyTransaction.ChainId,
			Blocks:  transaction1.PolyTransaction.Height,
			Time:    transaction1.PolyTransaction.Time,
		})
	} else {
		transactionRsp.TransactionState = append(transactionRsp.TransactionState, &TransactionStateRsp{
			Hash:    "",
			ChainId: 0,
			Blocks:  0,
			Time:    0,
		})
	}
	if transaction1.DstTransaction != nil {
		transactionRsp.TransactionState = append(transactionRsp.TransactionState, &TransactionStateRsp{
			Hash:    transaction1.DstTransaction.Hash,
			ChainId: transaction1.DstTransaction.ChainId,
			Blocks:  transaction1.DstTransaction.Height,
			Time:    transaction1.DstTransaction.Time,
		})
	} else {
		transactionRsp.TransactionState = append(transactionRsp.TransactionState, &TransactionStateRsp{
			Hash:    "",
			ChainId: basedef.O3_CROSSCHAIN_ID,
			Blocks:  0,
			Time:    0,
		})
	}
	if transaction2.PolyTransaction != nil {
		transactionRsp.TransactionState = append(transactionRsp.TransactionState, &TransactionStateRsp{
			Hash:    transaction2.PolyTransaction.Hash,
			ChainId: transaction2.PolyTransaction.ChainId,
			Blocks:  transaction2.PolyTransaction.Height,
			Time:    transaction2.PolyTransaction.Time,
		})
	} else {
		transactionRsp.TransactionState = append(transactionRsp.TransactionState, &TransactionStateRsp{
			Hash:    "",
			ChainId: 0,
			Blocks:  0,
			Time:    0,
		})
	}
	if transaction2.DstTransaction != nil {
		transactionRsp.TransactionState = append(transactionRsp.TransactionState, &TransactionStateRsp{
			Hash:    transaction2.DstTransaction.Hash,
			ChainId: transaction2.DstTransaction.ChainId,
			Blocks:  transaction2.DstTransaction.Height,
			Time:    transaction2.DstTransaction.Time,
		})
	} else {
		transactionRsp.TransactionState = append(transactionRsp.TransactionState, &TransactionStateRsp{
			Hash:    "",
			ChainId: transaction1.WrapperTransaction.DstChainId,
			Blocks:  0,
			Time:    0,
		})
	}
	for i, state := range transactionRsp.TransactionState {
		chain, ok := chainsMap[state.ChainId]
		if ok {
			if i == 0 {
				state.NeedBlocks = chain.BackwardBlockNumber
			} else if state.ChainId == basedef.O3_CROSSCHAIN_ID || state.ChainId == transaction1.WrapperTransaction.DstChainId {
				state.NeedBlocks = 1
			} else {
				state.NeedBlocks = chain.BackwardBlockNumber
			}
			if state.Blocks <= 1 {
				continue
			}
			state.Blocks = chain.Height - state.Blocks
			if state.Blocks > state.NeedBlocks {
				state.Blocks = state.NeedBlocks
			}
		}
	}
	return transactionRsp
}

type TransactionsOfAddressReq struct {
	State     int // -1 表示查全部
	Addresses []string
	PageSize  int
	PageNo    int
}

type TransactionsOfAddressWithFilterReq struct {
	State      int // -1 表示查全部
	Addresses  []string
	SrcChainId int
	DstChainId int
	Assets     []string
	PageSize   int
	PageNo     int
}

type TransactionsOfAddressRsp struct {
	PageSize     int
	PageNo       int
	TotalPage    int
	TotalCount   int
	Transactions []*TransactionRsp
}

func MakeTransactionsOfUserRsp(pageSize int, pageNo int, totalPage int, totalCount int, transactions []*SrcPolyDstRelation, chainsMap map[uint64]*Chain) *TransactionsOfAddressRsp {
	transactionsRsp := &TransactionsOfAddressRsp{
		PageSize:   pageSize,
		PageNo:     pageNo,
		TotalPage:  totalPage,
		TotalCount: totalCount,
	}
	for _, transaction := range transactions {
		rsp := MakeTransactionRsp(transaction, chainsMap)
		transactionsRsp.Transactions = append(transactionsRsp.Transactions, rsp)
	}
	return transactionsRsp
}

type TransactionsOfStateReq struct {
	State    uint64
	PageSize int
	PageNo   int
}

type TransactionsOfUnfinishedReq struct {
	PageSize int
	PageNo   int
}

type TransactionsOfAssetReq struct {
	Asset    string
	Chain    int
	PageSize int
	PageNo   int
}

type CrossChainTransactionRsp struct {
	WrapperTransaction *WrapperTransactionRsp
	SrcTransaction     *SrcTransactionRsp
	PolyTransaction    *PolyTransactionRsp
	DstTransaction     *DstTransactionRsp
	Token              *TokenRsp
}

type TransactionOfUnfinishedRsp struct {
	PageSize     int
	PageNo       int
	TotalPage    int
	TotalCount   int
	Transactions []*CrossChainTransactionRsp
}

type SrcTransactionRsp struct {
	Hash        string
	ChainId     uint64
	Standard    uint8
	State       uint64
	Time        uint64
	Height      uint64
	DstChainId  uint64
	SrcTransfer *SrcTransferRsp
}

type SrcTransferRsp struct {
	TxHash     string
	ChainId    uint64
	Standard   uint8
	Time       uint64
	Asset      string
	Amount     string
	DstChainId uint64
	DstAsset   string
}

func MakeSrcTransferRsp(transaction *SrcTransfer) *SrcTransferRsp {
	transactionRsp := &SrcTransferRsp{
		TxHash:     transaction.TxHash,
		ChainId:    transaction.ChainId,
		Standard:   transaction.Standard,
		Time:       transaction.Time,
		DstChainId: transaction.DstChainId,
		Asset:      transaction.Asset,
		DstAsset:   transaction.DstAsset,
	}
	return transactionRsp
}

func MakeSrcTransactionRsp(transaction *SrcTransaction) *SrcTransactionRsp {
	transactionRsp := &SrcTransactionRsp{
		Hash:       transaction.Hash,
		ChainId:    transaction.ChainId,
		Standard:   transaction.Standard,
		State:      transaction.State,
		Time:       transaction.Time,
		DstChainId: transaction.DstChainId,
		Height:     transaction.Height,
	}
	if transaction.SrcTransfer != nil {
		transactionRsp.SrcTransfer = MakeSrcTransferRsp(transaction.SrcTransfer)
	}
	return transactionRsp
}

type DstTransactionRsp struct {
	Hash        string
	ChainId     uint64
	Standard    uint8
	State       uint64
	Time        uint64
	Height      uint64
	SrcChainId  uint64
	DstTransfer *DstTransferRsp
}

type DstTransferRsp struct {
	TxHash   string
	ChainId  uint64
	Standard uint8
	Time     uint64
	Asset    string
	Amount   string
}

func MakeDstTransferRsp(transaction *DstTransfer) *DstTransferRsp {
	transactionRsp := &DstTransferRsp{
		TxHash:   transaction.TxHash,
		ChainId:  transaction.ChainId,
		Standard: transaction.Standard,
		Time:     transaction.Time,
		Asset:    transaction.Asset,
	}
	return transactionRsp
}

func MakeDstTransactionRsp(transaction *DstTransaction) *DstTransactionRsp {
	transactionRsp := &DstTransactionRsp{
		Hash:       transaction.Hash,
		ChainId:    transaction.ChainId,
		Standard:   transaction.Standard,
		State:      transaction.State,
		Time:       transaction.Time,
		SrcChainId: transaction.SrcChainId,
		Height:     transaction.Height,
	}
	if transaction.DstTransfer != nil {
		transactionRsp.DstTransfer = MakeDstTransferRsp(transaction.DstTransfer)
	}
	return transactionRsp
}

func MakeCrossChainTransactionRsp(transaction *SrcPolyDstRelation) *CrossChainTransactionRsp {
	crossChainTransactionRsp := new(CrossChainTransactionRsp)
	if transaction.WrapperTransaction != nil {
		crossChainTransactionRsp.WrapperTransaction = MakeWrapperTransactionRsp(transaction.WrapperTransaction)
		if transaction.Token != nil {
			precision := decimal.NewFromInt(basedef.Int64FromFigure(int(transaction.Token.Precision)))
			{
				bbb := decimal.NewFromBigInt(&transaction.WrapperTransaction.FeeAmount.Int, 0)
				feeAmount := bbb.Div(precision)
				crossChainTransactionRsp.WrapperTransaction.FeeAmount = feeAmount.String()
			}
		}
	}
	if transaction.SrcTransaction != nil {
		crossChainTransactionRsp.SrcTransaction = MakeSrcTransactionRsp(transaction.SrcTransaction)
		if transaction.Token != nil && transaction.SrcTransaction.SrcTransfer != nil {
			precision := decimal.NewFromInt(basedef.Int64FromFigure(int(transaction.Token.Precision)))
			{
				bbb := decimal.NewFromBigInt(&transaction.SrcTransaction.SrcTransfer.Amount.Int, 0)
				transferAmount := bbb.Div(precision)
				crossChainTransactionRsp.SrcTransaction.SrcTransfer.Amount = transferAmount.String()
			}
		}
	}
	if transaction.PolyTransaction != nil {
		crossChainTransactionRsp.PolyTransaction = MakePolyTransactionRsp(transaction.PolyTransaction)
	}
	if transaction.DstTransaction != nil {
		crossChainTransactionRsp.DstTransaction = MakeDstTransactionRsp(transaction.DstTransaction)
	}
	if transaction.Token != nil {
		crossChainTransactionRsp.Token = MakeTokenRsp(transaction.Token)
	}
	return crossChainTransactionRsp
}

func MakeTransactionOfUnfinishedRsp(pageSize int, pageNo int, totalPage int, totalCount int, transactions []*SrcPolyDstRelation) *TransactionOfUnfinishedRsp {
	transactionOfUnfinishedRsp := &TransactionOfUnfinishedRsp{
		PageSize:   pageSize,
		PageNo:     pageNo,
		TotalPage:  totalPage,
		TotalCount: totalCount,
	}
	for _, transaction := range transactions {
		transactionOfUnfinishedRsp.Transactions = append(transactionOfUnfinishedRsp.Transactions, MakeCrossChainTransactionRsp(transaction))
	}
	return transactionOfUnfinishedRsp
}

type TransactionsOfStateRsp struct {
	PageSize     int
	PageNo       int
	TotalPage    int
	TotalCount   int
	Transactions []*WrapperTransactionRsp
}

func MakeTransactionsOfStateRsp(pageSize int, pageNo int, totalPage int, totalCount int, transactions []*WrapperTransaction) *WrapperTransactionsRsp {
	transactionsRsp := &WrapperTransactionsRsp{
		PageSize:   pageSize,
		PageNo:     pageNo,
		TotalPage:  totalPage,
		TotalCount: totalCount,
	}
	for _, transaction := range transactions {
		transactionsRsp.Transactions = append(transactionsRsp.Transactions, MakeWrapperTransactionRsp(transaction))
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

type PolyTransactionReq struct {
	Hash string
}

type PolyTransactionRsp struct {
	Hash       string
	ChainId    uint64
	State      uint64
	Time       uint64
	Fee        string
	Height     uint64
	SrcChainId uint64
	SrcHash    string
	DstChainId uint64
	Key        string
}

func MakePolyTransactionRsp(transaction *PolyTransaction) *PolyTransactionRsp {
	transactionRsp := &PolyTransactionRsp{
		Hash:       transaction.Hash,
		ChainId:    transaction.ChainId,
		State:      transaction.State,
		Time:       transaction.Time,
		Fee:        transaction.Fee.String(),
		Height:     transaction.Height,
		SrcChainId: transaction.SrcChainId,
		SrcHash:    transaction.SrcHash,
		DstChainId: transaction.DstChainId,
		Key:        transaction.Key,
	}
	return transactionRsp
}

type PolyTransactionsReq struct {
	PageSize int
	PageNo   int
}

type PolyTransactionsRsp struct {
	PageSize     int
	PageNo       int
	TotalPage    int
	TotalCount   int
	Transactions []*PolyTransactionRsp
}

func MakePolyTransactionsRsp(pageSize int, pageNo int, totalPage int, totalCount int, transactions []*PolyTransaction) *PolyTransactionsRsp {
	transactionsRsp := &PolyTransactionsRsp{
		PageSize:   pageSize,
		PageNo:     pageNo,
		TotalPage:  totalPage,
		TotalCount: totalCount,
	}
	for _, transaction := range transactions {
		transactionsRsp.Transactions = append(transactionsRsp.Transactions, MakePolyTransactionRsp(transaction))
	}
	return transactionsRsp
}

type ExpectTimeReq struct {
	SrcChainId uint64
	DstChainId uint64
}

type ExpectTimeRsp struct {
	SrcChainId uint64
	DstChainId uint64
	Time       uint64
}

func MakeExpectTimeRsp(srcchainId uint64, dstchainid uint64, time uint64) *ExpectTimeRsp {
	expectTimeRsp := &ExpectTimeRsp{
		Time:       time,
		SrcChainId: srcchainId,
		DstChainId: dstchainid,
	}
	return expectTimeRsp
}

type TokenAssetReq struct {
	NameOrHash string
}

type AssetDetailRes struct {
	BasicName  string
	TokenAsset []*DstChainAssetRes
	Precision  uint64
}
type DstChainAssetRes struct {
	ChainName   string
	Hash        string
	TotalSupply *big.Int
	Balance     *big.Int
	ErrReason   string
}

type CheckFeeResult struct {
	Pass bool
	Paid float64
	Min  float64
}

type BotTx struct {
	Asset        string
	Hash         string
	PolyHash     string
	SrcChainId   uint64
	SrcChainName string
	DstChainId   uint64
	DstChainName string
	Amount       string
	Time         string
	Duration     string
	Status       string
	FeeToken     string
	FeePaid      float64
	FeeMin       float64
	FeePass      string
	ProxyProject string
}

func ParseBotTx(tx *SrcPolyDstRelation, fees map[string]CheckFeeResult) BotTx {
	// in case src transaction is missing
	hash := tx.SrcHash
	if hash == "" {
		hash = tx.PolyHash
	}
	v := BotTx{Hash: hash, PolyHash: tx.PolyHash}
	if c := tx.WrapperTransaction; c != nil {
		v.SrcChainId = c.SrcChainId
		v.DstChainId = c.DstChainId
		v.SrcChainName = basedef.GetChainName(v.SrcChainId)
		v.DstChainName = basedef.GetChainName(v.DstChainId)
		v.Status = basedef.GetStateName(int(c.Status))
		tsp := time.Unix(int64(c.Time), 0)
		v.Time = tsp.Format(time.RFC3339)
		v.Duration = time.Now().Sub(tsp).String()
	} else if s := tx.SrcTransaction; s != nil {
		v.SrcChainId = s.ChainId
		v.DstChainId = s.DstChainId
		v.SrcChainName = basedef.GetChainName(v.SrcChainId)
		v.DstChainName = basedef.GetChainName(v.DstChainId)
		//v.Status = basedef.GetStateName(int(c.Status))
		tsp := time.Unix(int64(s.Time), 0)
		v.Time = tsp.Format(time.RFC3339)
		v.Duration = time.Now().Sub(tsp).String()
	}
	v.ProxyProject = "POLY"
	if s := tx.SrcTransaction; s != nil {
		if _, in := conf.EstimateProxy[strings.ToUpper(s.Contract)]; in {
			v.ProxyProject = "O3V2"
		}
	}
	if fee, ok := fees[v.Hash]; ok {
		v.FeePaid = fee.Paid
		v.FeeMin = fee.Min
		v.FeePass = "NotPass"
		if fee.Pass {
			v.FeePass = "Pass"
		}
	} else {
		v.FeePass = "Unknown"
	}
	if token := tx.Token; token != nil {
		v.Asset = token.Name
	} else {
		v.Asset = tx.TokenHash
	}
	if token := tx.FeeToken; token != nil {
		v.FeeToken = token.Name
	}
	if tx.SrcTransaction != nil && tx.SrcTransaction.SrcTransfer != nil {
		if amount := tx.SrcTransaction.SrcTransfer.Amount; amount != nil {
			v.Amount = amount.String()
			if tx.Token != nil {
				tokenAmount := new(big.Float).Quo(new(big.Float).SetInt(&amount.Int), new(big.Float).SetInt64(basedef.Int64FromFigure(int(tx.Token.Precision))))
				v.Amount = tokenAmount.String()
			}
		}
	}
	return v
}

func MakeBottxsRsp(pageSize int, pageNo int, totalPage int, totalCount int, transactions []*SrcPolyDstRelation, fees map[string]CheckFeeResult) map[string]interface{} {
	rsp := map[string]interface{}{}
	rsp["PageSize"] = pageSize
	rsp["PageNo"] = pageNo
	rsp["TotalPage"] = totalPage
	rsp["TotalCount"] = totalCount
	txs := make([]BotTx, len(transactions))
	for i, tx := range transactions {
		txs[i] = ParseBotTx(tx, fees)
	}
	rsp["Transactions"] = txs
	return rsp
}

type ManualTxDataReq struct {
	PolyHash string
}

type ManualTxDataResp struct {
	Data   string `json:"data"`
	DstCCM string `json:"dst_ccm"`
}

type ChainHealthReq struct {
	ChainIds []uint64
}

type ChainHealthRsp struct {
	Result map[uint64]bool
}

type WrapperCheckReq struct {
	ChainId uint64
}

type WrapperCheckRsp struct {
	ChainId uint64
	Wrapper []string
}

type TxWithoutWrapperReq struct {
	ChainId  uint64
	User     string
	PageSize int
	PageNo   int
}

type TxWithoutWrapperRes struct {
	Total    int64
	PageSize int
	PageNo   int
	Txs      []*TxWithoutWrappertx
}

type TxWithoutWrappertx struct {
	TxHash     string
	ChainId    uint64
	Time       uint64
	From       string
	Amount     string
	DstChainId uint64
	DstUser    string
	TokenName  string
}

func MakeTxWithoutWrapperRsp(pageSize int, pageNo int, srcTransfers []*SrcTransfer, count int64) *TxWithoutWrapperRes {
	txWithoutWrapperRes := new(TxWithoutWrapperRes)
	txWithoutWrapperRes.Total = count
	txWithoutWrapperRes.PageSize = pageSize
	txWithoutWrapperRes.PageNo = pageNo
	txWithoutWrapperRes.Txs = make([]*TxWithoutWrappertx, 0)
	for _, v := range srcTransfers {
		txWithoutWrappertx := &TxWithoutWrappertx{
			v.TxHash,
			v.ChainId,
			v.Time,
			v.From,
			v.Amount.String(),
			v.DstChainId,
			v.DstUser,
			v.Asset,
		}

		if v.Token != nil {
			txWithoutWrappertx.Amount = FormatAmount(v.Token.Precision, v.Amount)
			txWithoutWrappertx.TokenName = v.Token.TokenBasicName
		}
		txWithoutWrapperRes.Txs = append(txWithoutWrapperRes.Txs, txWithoutWrappertx)
	}
	return txWithoutWrapperRes
}
