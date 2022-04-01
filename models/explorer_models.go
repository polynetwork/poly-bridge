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

// Package classification User API.
//
// The purpose of this service is to provide an application
// that is using plain go code to define an API
//
//      Host: localhost
//      Version: 0.0.1
//
// swagger:meta

package models

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/utils/decimal"
	"sort"
	"strconv"
	"strings"
)

type ExplorerInfoResp struct {
	Chains        []*ChainInfoResp       `json:"chains"`
	CrossTxNumber int64                  `json:"crosstxnumber"`
	Tokens        []*CrossChainTokenResp `json:"tokens"`
}

func getChainStatistic(chainId uint64, statistics []*ChainStatistic) *ChainStatistic {
	for _, statistic := range statistics {
		if statistic.ChainId == chainId {
			return statistic
		}
	}
	return nil
}

func MakeExplorerInfoResp(chains []*Chain, statistics []*ChainStatistic, tokenBasics []*TokenBasic) *ExplorerInfoResp {
	chainInfoResps := make([]*ChainInfoResp, 0)
	for _, chain := range chains {
		chainInfoResp := MakeChainInfoResp(chain)
		for _, statistic := range statistics {
			if statistic.ChainId == chain.ChainId {
				chainInfoResp.Addresses = statistic.Addresses
				chainInfoResp.In = statistic.In
				chainInfoResp.Out = statistic.Out
			}
		}
		for _, tokenBasic := range tokenBasics {
			for _, token := range tokenBasic.Tokens {
				if token.ChainId == chain.ChainId && token.Property == int64(1) {
					chainInfoResp.Tokens = append(chainInfoResp.Tokens, MakeChainTokenResp(token))
				}
			}
		}
		chainInfoResps = append(chainInfoResps, chainInfoResp)
	}
	crossTxNumber := getChainStatistic(basedef.POLY_CROSSCHAIN_ID, statistics).In
	crossChainTokenResp := make([]*CrossChainTokenResp, 0)
	for _, tokenBasic := range tokenBasics {
		crossChainTokenResp = append(crossChainTokenResp, MakeTokenBasicResp(tokenBasic))
	}
	explorerInfoResp := &ExplorerInfoResp{
		Chains:        chainInfoResps,
		CrossTxNumber: crossTxNumber,
		Tokens:        crossChainTokenResp,
	}
	return explorerInfoResp
}

type ChainInfoResp struct {
	Id        uint32 `json:"chainid"`
	Name      string `json:"chainname"`
	ChainLogo string `json:"chainlogo"`
	Height    uint32 `json:"blockheight"`
	In        int64  `json:"in"`
	//InCrossChainTxStatus []*CrossChainTxStatus    `json:"incrosschaintxstatus"`
	Out int64 `json:"out"`
	//OutCrossChainTxStatus []*CrossChainTxStatus    `json:"outcrosschaintxstatus"`
	Addresses int64 `json:"addresses"`
	//Contracts []*ChainContractResp `json:"contracts"`
	Tokens []*ChainTokenResp `json:"tokens"`
}

func MakeChainInfoResp(chain *Chain) *ChainInfoResp {
	chainInfoResp := &ChainInfoResp{
		Id:        uint32(chain.ChainId),
		Name:      chain.Name,
		Height:    uint32(chain.Height),
		ChainLogo: chain.ChainLogo,
	}
	return chainInfoResp
}

type CrossChainTxStatus struct {
	TT       uint32 `json:"timestamp"`
	TxNumber uint32 `json:"txnumber"`
}

type ChainContractResp struct {
	Id       uint32 `json:"chainid"`
	Contract string `json:"contract"`
}

type ChainTokenResp struct {
	Chain     int32  `json:"chainid"`
	ChainName string `json:"chainname"`
	Hash      string `json:"hash"`
	Token     string `json:"token"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Precision uint64 `json:"precision"`
	Desc      string `json:"desc"`
	Meta      string `json:"meta"`
}

func MakeChainTokenResp(token *Token) *ChainTokenResp {
	chainTokenResp := &ChainTokenResp{
		Chain:     int32(token.ChainId),
		ChainName: ChainId2Name(token.ChainId),
		Hash:      token.Hash,
		Token:     token.TokenBasicName,
		Name:      token.Name,
		Type:      token.TokenType,
		Precision: token.Precision,
	}
	return chainTokenResp
}

type CrossChainTokenResp struct {
	Name      string            `json:"name"`
	TokenLogo string            `json:"tokenlogo"`
	Tokens    []*ChainTokenResp `json:"tokens"`
}

func MakeTokenBasicResp(tokenBasic *TokenBasic) *CrossChainTokenResp {
	crossChainTokenResp := &CrossChainTokenResp{
		Name:      tokenBasic.Name,
		TokenLogo: tokenBasic.Meta,
	}
	for _, token := range tokenBasic.Tokens {
		if token.Property == int64(1) {
			crossChainTokenResp.Tokens = append(crossChainTokenResp.Tokens, MakeChainTokenResp(token))
		}
	}
	return crossChainTokenResp
}

type FChainTxResp struct {
	ChainId          uint32              `json:"chainid"`
	ChainName        string              `json:"chainname"`
	ChainLogo        string              `json:"chainlogo"`
	ChainExplorerUrl string              `json:"chainexplorerurl"`
	TxHash           string              `json:"txhash"`
	State            byte                `json:"state"`
	TT               uint32              `json:"timestamp"`
	Fee              string              `json:"fee"`
	FeeName          string              `json:"feename"`
	FeeLogo          string              `json:"feelogo"`
	Height           uint32              `json:"blockheight"`
	User             string              `json:"user"`
	TChainId         uint32              `json:"tchainid"`
	TChainName       string              `json:"tchainname"`
	TChainLogo       string              `json:"tchainlogo"`
	Contract         string              `json:"contract"`
	Key              string              `json:"key"`
	Param            string              `json:"param"`
	Transfer         *FChainTransferResp `json:"transfer"`
}

func makeFChainTxResp(fChainTx *SrcTransaction, token, toToken *Token) *FChainTxResp {
	fChainTxResp := &FChainTxResp{
		ChainId:          uint32(fChainTx.ChainId),
		ChainName:        ChainId2Name(fChainTx.ChainId),
		ChainLogo:        ChainId2ChainCache(fChainTx.ChainId).ChainFeeLogo,
		ChainExplorerUrl: ChainId2ChainCache(fChainTx.ChainId).ChainExplorerUrl,
		TxHash:           fChainTx.Hash,
		State:            byte(fChainTx.State),
		TT:               uint32(fChainTx.Time),
		Fee:              FormatFee(fChainTx.ChainId, fChainTx.Fee),
		FeeName:          ChainId2ChainCache(fChainTx.ChainId).ChainFeeName,
		FeeLogo:          ChainId2ChainCache(fChainTx.ChainId).ChainFeeLogo,
		Height:           uint32(fChainTx.Height),
		User:             basedef.Hash2Address(fChainTx.ChainId, fChainTx.User),
		TChainId:         uint32(fChainTx.DstChainId),
		TChainName:       ChainId2Name(fChainTx.DstChainId),
		TChainLogo:       ChainId2ChainCache(fChainTx.DstChainId).ChainLogo,
		Contract:         fChainTx.Contract,
		Key:              fChainTx.Key,
		Param:            fChainTx.Param,
	}
	if fChainTx.SrcTransfer.Amount == nil {
		fChainTx.SrcTransfer.Amount = NewBigIntFromInt(0)
	}
	fChainTxResp.Transfer = &FChainTransferResp{
		From:        basedef.Hash2Address(fChainTx.SrcTransfer.ChainId, fChainTx.SrcTransfer.From),
		To:          basedef.Hash2Address(fChainTx.SrcTransfer.ChainId, fChainTx.SrcTransfer.To),
		Amount:      strconv.FormatUint(fChainTx.SrcTransfer.Amount.Uint64(), 10),
		ToChain:     uint32(fChainTx.SrcTransfer.DstChainId),
		ToChainName: ChainId2Name(fChainTx.SrcTransfer.DstChainId),
		ToUser:      basedef.Hash2Address(fChainTx.SrcTransfer.DstChainId, fChainTx.SrcTransfer.DstUser),
	}
	fChainTxResp.Transfer.TokenHash = fChainTx.SrcTransfer.Asset
	if token != nil {
		fChainTxResp.Transfer.TokenHash = token.Hash
		fChainTxResp.Transfer.TokenName = token.Name
		if token.TokenBasic != nil {
			fChainTxResp.Transfer.TokenLogo = token.TokenBasic.Meta
		}
		fChainTxResp.Transfer.TokenType = token.TokenType
		fChainTxResp.Transfer.Amount = FormatAmount(token.Precision, fChainTx.SrcTransfer.Amount)
	} else {
		fChainTxResp.Transfer.TokenName = fChainTx.SrcTransfer.Asset
		fChainTxResp.Transfer.Amount = fChainTx.SrcTransfer.Amount.String()
	}
	fChainTxResp.Transfer.ToTokenHash = fChainTx.SrcTransfer.DstAsset
	if toToken != nil {
		fChainTxResp.Transfer.ToTokenHash = toToken.Hash
		fChainTxResp.Transfer.ToTokenName = toToken.Name
		if toToken.TokenBasic != nil {
			fChainTxResp.Transfer.TokenLogo = toToken.TokenBasic.Meta
		}
		fChainTxResp.Transfer.ToTokenType = toToken.TokenType
	} else {
		fChainTxResp.Transfer.ToTokenName = fChainTx.SrcTransfer.DstAsset
	}
	if fChainTx.ChainId == basedef.ETHEREUM_CROSSCHAIN_ID || fChainTx.ChainId == basedef.ZILLIQA_CROSSCHAIN_ID {
		fChainTxResp.TxHash = "0x" + fChainTxResp.TxHash
	} else if fChainTx.ChainId == basedef.SWITCHEO_CROSSCHAIN_ID {
		fChainTxResp.TxHash = strings.ToUpper(fChainTxResp.TxHash)
	}
	return fChainTxResp
}

type FChainTransferResp struct {
	TokenHash   string `json:"tokenhash"`
	TokenName   string `json:"tokenname"`
	TokenLogo   string `json:"tokenlogo"`
	TokenType   string `json:"tokentype"`
	From        string `json:"from"`
	To          string `json:"to"`
	Amount      string `json:"amount"`
	ToChain     uint32 `json:"tchainid"`
	ToChainName string `json:"tchainname"`
	ToTokenHash string `json:"totokenhash"`
	ToTokenName string `json:"totokenname"`
	ToTokenLogo string `json:"totokenlogo"`
	ToTokenType string `json:"totokentype"`
	ToUser      string `json:"tuser"`
}

type MChainTxResp struct {
	ChainId    uint32 `json:"chainid"`
	ChainName  string `json:"chainname"`
	TxHash     string `json:"txhash"`
	State      byte   `json:"state"`
	TT         uint32 `json:"timestamp"`
	Fee        string `json:"fee"`
	Height     uint32 `json:"blockheight"`
	FChainId   uint32 `json:"fchainid"`
	FChainName string `json:"fchainname"`
	FChainLogo string `json:"fchainlogo"`
	FTxHash    string `json:"ftxhash"`
	TChainId   uint32 `json:"tchainid"`
	TChainName string `json:"tchainname"`
	TChainLogo string `json:"tchainlogo"`
	Key        string `json:"key"`
}

func makeMChainTxResp(mChainTx *PolyTransaction) *MChainTxResp {
	mChainTxResp := &MChainTxResp{
		ChainId:    uint32(mChainTx.ChainId),
		ChainName:  ChainId2Name(mChainTx.ChainId),
		TxHash:     mChainTx.Hash,
		State:      byte(mChainTx.State),
		TT:         uint32(mChainTx.Time),
		Fee:        FormatFee(mChainTx.ChainId, mChainTx.Fee),
		Height:     uint32(mChainTx.Height),
		FChainId:   uint32(mChainTx.SrcChainId),
		FChainName: ChainId2Name(mChainTx.SrcChainId),
		FChainLogo: ChainId2ChainCache(mChainTx.SrcChainId).ChainLogo,
		FTxHash:    mChainTx.SrcHash,
		TChainId:   uint32(mChainTx.DstChainId),
		TChainName: ChainId2Name(mChainTx.DstChainId),
		TChainLogo: ChainId2ChainCache(mChainTx.DstChainId).ChainLogo,
		Key:        mChainTx.Key,
	}
	return mChainTxResp
}

type TChainTxResp struct {
	ChainId          uint32              `json:"chainid"`
	ChainName        string              `json:"chainname"`
	ChainLogo        string              `json:"chainlogo"`
	TxHash           string              `json:"txhash"`
	State            byte                `json:"state"`
	TT               uint32              `json:"timestamp"`
	Fee              string              `json:"fee"`
	FeeName          string              `json:"feename"`
	FeeLogo          string              `json:"feelogo"`
	Height           uint32              `json:"blockheight"`
	FChainId         uint32              `json:"fchainid"`
	FChainName       string              `json:"fchainname"`
	FChainLogo       string              `json:"fchainlogo"`
	Contract         string              `json:"contract"`
	RTxHash          string              `json:"mtxhash"`
	Transfer         *TChainTransferResp `json:"transfer"`
	ChainExplorerUrl string              `json:"chainExplorerUrl"`
}

func makeTChainTxResp(tChainTx *DstTransaction, toToken *Token) *TChainTxResp {
	tChainTxResp := &TChainTxResp{
		ChainId:          uint32(tChainTx.ChainId),
		ChainName:        ChainId2Name(tChainTx.ChainId),
		ChainLogo:        ChainId2ChainCache(tChainTx.ChainId).ChainLogo,
		ChainExplorerUrl: ChainId2ChainCache(tChainTx.ChainId).ChainExplorerUrl,
		TxHash:           tChainTx.Hash,
		State:            byte(tChainTx.State),
		TT:               uint32(tChainTx.Time),
		Fee:              FormatFee(tChainTx.ChainId, tChainTx.Fee),
		FeeName:          ChainId2ChainCache(tChainTx.ChainId).ChainFeeName,
		FeeLogo:          ChainId2ChainCache(tChainTx.ChainId).ChainFeeLogo,
		Height:           uint32(tChainTx.Height),
		FChainId:         uint32(tChainTx.SrcChainId),
		FChainName:       ChainId2Name(tChainTx.SrcChainId),
		FChainLogo:       ChainId2ChainCache(tChainTx.SrcChainId).ChainLogo,
		Contract:         tChainTx.Contract,
		RTxHash:          tChainTx.PolyHash,
	}
	if tChainTx.DstTransfer.Amount == nil {
		tChainTx.DstTransfer.Amount = NewBigIntFromInt(0)
	}
	tChainTxResp.Transfer = &TChainTransferResp{
		From:   basedef.Hash2Address(tChainTx.DstTransfer.ChainId, tChainTx.DstTransfer.From),
		To:     basedef.Hash2Address(tChainTx.DstTransfer.ChainId, tChainTx.DstTransfer.To),
		Amount: tChainTx.DstTransfer.Amount.String(),
	}
	tChainTxResp.Transfer.TokenHash = tChainTx.DstTransfer.Asset
	if toToken != nil {
		tChainTxResp.Transfer.TokenHash = toToken.Hash
		tChainTxResp.Transfer.TokenName = toToken.Name
		tChainTxResp.Transfer.TokenType = toToken.TokenType
		if toToken.TokenBasic != nil {
			tChainTxResp.Transfer.TokenLogo = toToken.TokenBasic.Meta
		}
		tChainTxResp.Transfer.Amount = FormatAmount(toToken.Precision, tChainTx.DstTransfer.Amount)
	} else {
		tChainTxResp.Transfer.TokenName = tChainTx.DstTransfer.Asset
		tChainTxResp.Transfer.Amount = tChainTx.DstTransfer.Amount.String()
	}

	if tChainTx.ChainId == basedef.ETHEREUM_CROSSCHAIN_ID || tChainTx.ChainId == basedef.ZILLIQA_CROSSCHAIN_ID {
		tChainTxResp.TxHash = "0x" + tChainTxResp.TxHash
	} else if tChainTx.ChainId == basedef.SWITCHEO_CROSSCHAIN_ID {
		tChainTxResp.TxHash = strings.ToUpper(tChainTxResp.TxHash)
	}
	return tChainTxResp
}

type TChainTransferResp struct {
	TokenHash string `json:"tokenhash"`
	TokenName string `json:"tokenname"`
	TokenLogo string `json:"tokenlogo"`
	TokenType string `json:"tokentype"`
	From      string `json:"from"`
	To        string `json:"to"`
	Amount    string `json:"amount"`
}

type CrossTransferResp struct {
	CrossTxType   uint32 `json:"crosstxtype"`
	CrossTxName   string `json:"crosstxname"`
	FromChainId   uint32 `json:"fromchainid"`
	FromChain     string `json:"fromchainname"`
	FromChainLogo string `json:"fromchainlogo"`
	FromAddress   string `json:"fromaddress"`
	ToChainId     uint32 `json:"tochainid"`
	ToChain       string `json:"tochainname"`
	ToChainLogo   string `json:"tochainlogo"`
	ToAddress     string `json:"toaddress"`
	TokenHash     string `json:"tokenhash"`
	TokenName     string `json:"tokenname"`
	TokenType     string `json:"tokentype"`
	Amount        string `json:"amount"`
}

func makeCrossTransfer(chainid uint64, user string, transfer *SrcTransfer, token *Token) *CrossTransferResp {
	if transfer.Amount == nil {
		transfer.Amount = NewBigIntFromInt(0)
	}
	crossTransfer := new(CrossTransferResp)
	crossTransfer.CrossTxType = 1
	crossTransfer.CrossTxName = TxType2Name(crossTransfer.CrossTxType)
	crossTransfer.FromChainId = uint32(chainid)
	crossTransfer.FromChain = ChainId2Name(uint64(crossTransfer.FromChainId))
	crossTransfer.FromChainLogo = ChainId2ChainCache(uint64(crossTransfer.FromChainId)).ChainLogo
	crossTransfer.FromAddress = basedef.Hash2Address(chainid, user)
	crossTransfer.ToChainId = uint32(transfer.DstChainId)
	crossTransfer.ToChain = ChainId2Name(uint64(crossTransfer.ToChainId))
	crossTransfer.ToChainLogo = ChainId2ChainCache(uint64(crossTransfer.ToChainId)).ChainLogo
	crossTransfer.ToAddress = basedef.Hash2Address(transfer.DstChainId, transfer.DstUser)
	jsonToken, _ := json.Marshal(token)
	logs.Info("yuan crossTransfer.Amount:", transfer.Amount, string(jsonToken))
	if token != nil {
		crossTransfer.TokenHash = token.Hash
		crossTransfer.TokenName = token.Name
		crossTransfer.TokenType = token.TokenType
		crossTransfer.Amount = FormatAmount(token.Precision, transfer.Amount)
	} else {
		crossTransfer.TokenName = transfer.Asset
		crossTransfer.Amount = transfer.Amount.String()
	}
	logs.Info("xian crossTransfer.Amount:", crossTransfer.Amount)
	return crossTransfer
}

// swagger:parameters CrossTxReq
type CrossTxReq struct {
	// in: query
	TxHash string `json:"txhash"`
}

type CrossTxResp struct {
	Transfer        *CrossTransferResp `json:"crosstransfer"`
	Fchaintx        *FChainTxResp      `json:"fchaintx"`
	Fchaintx_valid  bool               `json:"fchaintx_valid"`
	Mchaintx        *MChainTxResp      `json:"mchaintx"`
	Mchaintx_valid  bool               `json:"mchaintx_valid"`
	Tchaintx        *TChainTxResp      `json:"tchaintx"`
	Tchaintx_valid  bool               `json:"tchaintx_valid"`
	RelatedPolyHash string             `json:"relatedPolyHash"`
}

type TransferStatisticReq struct {
	Chain uint64 `json:"chain"`
}

func MakeCrossTxResp(srcPolyDst *PolyTxRelation) *CrossTxResp {
	crosstx := &CrossTxResp{
		Fchaintx_valid: false,
		Mchaintx_valid: false,
		Tchaintx_valid: false,
		Transfer: &CrossTransferResp{
			CrossTxType: 0,
		},
		RelatedPolyHash: srcPolyDst.RelatedPolyHash,
	}
	tx := srcPolyDst
	if tx.SrcTransaction != nil {
		crosstx.Fchaintx_valid = true
		crosstx.Fchaintx = makeFChainTxResp(tx.SrcTransaction, tx.Token, tx.ToToken)
		crosstx.Transfer = makeCrossTransfer(tx.SrcTransaction.ChainId, tx.SrcTransaction.User, tx.SrcTransaction.SrcTransfer, tx.Token)
	}
	if tx.PolyTransaction != nil && crosstx.Fchaintx_valid == true {
		crosstx.Mchaintx_valid = true
		crosstx.Mchaintx = makeMChainTxResp(tx.PolyTransaction)
	}
	if tx.DstTransaction != nil && crosstx.Mchaintx_valid == true {
		crosstx.Tchaintx_valid = true
		crosstx.Tchaintx = makeTChainTxResp(tx.DstTransaction, tx.DstToken)
	}
	return crosstx
}

type CrossTxListReq struct {
	PageSize int
	PageNo   int
}

type CrossTxOutlineResp struct {
	TxHash     string `json:"txhash"`
	State      byte   `json:"state"`
	TT         uint32 `json:"timestamp"`
	Fee        string `json:"fee"`
	Height     uint32 `json:"blockheight"`
	FChainId   uint32 `json:"fchainid"`
	FChainName string `json:"fchainname"`
	FChainLogo string `json:"fchainlogo"`
	TChainId   uint32 `json:"tchainid"`
	TChainName string `json:"tchainname"`
	TChainLogo string `json:"tchainlogo"`
}

type CrossTxListResp struct {
	CrossTxList []*CrossTxOutlineResp `json:"crosstxs"`
	Total       int64                 `json:"total"`
}

func MakeCrossTxListResp(txs []*SrcPolyDstRelation, counter int64) *CrossTxListResp {
	logs.Info("jinru MakeCrossTxListResp")
	crossTxListResp := &CrossTxListResp{}
	crossTxListResp.CrossTxList = make([]*CrossTxOutlineResp, 0)
	for _, tx := range txs {
		crossTxListResp.CrossTxList = append(crossTxListResp.CrossTxList, &CrossTxOutlineResp{
			TxHash:     tx.PolyHash,
			State:      byte(tx.PolyTransaction.State),
			TT:         uint32(tx.PolyTransaction.Time),
			Fee:        FormatFee(tx.PolyTransaction.ChainId, tx.PolyTransaction.Fee),
			Height:     uint32(tx.PolyTransaction.Height),
			FChainId:   uint32(tx.PolyTransaction.SrcChainId),
			FChainName: ChainId2Name(tx.PolyTransaction.SrcChainId),
			FChainLogo: ChainId2ChainCache(tx.PolyTransaction.SrcChainId).ChainLogo,
			TChainId:   uint32(tx.PolyTransaction.DstChainId),
			TChainName: ChainId2Name(tx.PolyTransaction.DstChainId),
			TChainLogo: ChainId2ChainCache(tx.PolyTransaction.DstChainId).ChainLogo,
		})
	}
	crossTxListResp.Total = counter
	logs.Info("jieshu MakeCrossTxListResp")
	return crossTxListResp
}

type TokenTxListReq struct {
	ChainId  uint64 `json:"chain"`
	Token    string `json:"token"`
	PageSize int
	PageNo   int
}

type TokenTxResp struct {
	TxHash string `json:"txhash"`
	From   string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"`
	TT     uint32 `json:"timestamp"`
	Height uint32 `json:"blockheight"`
	Direct uint32 `json:"direct"`
}

type TokenTxListResp struct {
	TokenTxList []*TokenTxResp `json:"tokentxs"`
	Total       int64          `json:"total"`
}

func MakeTokenTxList(transactoins []*TransactionOnToken, counter int64, token *Token) *TokenTxListResp {
	tokenTxListResp := &TokenTxListResp{}
	tokenTxListResp.Total = counter
	tokenTxListResp.TokenTxList = make([]*TokenTxResp, 0)
	for _, transactoin := range transactoins {
		amount := transactoin.Amount.String()
		if token != nil {
			amount = FormatAmount(token.TokenBasic.Precision, transactoin.Amount)
		}
		tokenTxListResp.TokenTxList = append(tokenTxListResp.TokenTxList, &TokenTxResp{
			TxHash: transactoin.Hash,
			From:   basedef.Hash2Address(transactoin.ChainId, transactoin.From),
			To:     basedef.Hash2Address(transactoin.ChainId, transactoin.To),
			Amount: amount,
			Height: uint32(transactoin.Height),
			TT:     uint32(transactoin.Time),
			Direct: transactoin.Direct,
		})
	}
	return tokenTxListResp
}

type AddressTxListReq struct {
	PageSize int
	PageNo   int
	Address  string `json:"address"`
	ChainId  uint64 `json:"chain"`
}

type AddressTxResp struct {
	TxHash    string `json:"txhash"`
	ChainName string `json:"chainname"`
	ChainLogo string `json:"chainlogo"`
	From      string `json:"from"`
	To        string `json:"to"`
	Amount    string `json:"amount"`
	TT        uint32 `json:"timestamp"`
	Height    uint32 `json:"blockheight"`
	TokenHash string `json:"tokenhash"`
	TokenName string `json:"tokenname"`
	TokenLogo string `json:"tokenlogo"`
	TokenType string `json:"tokentype"`
	Direct    uint32 `json:"direct"`
}

type AddressTxListResp struct {
	AddressTxList []*AddressTxResp `json:"addresstxs"`
	Total         int64            `json:"total"`
}

func MakeAddressTxList(transactoins []*TransactionOnAddress, counter int64) *AddressTxListResp {
	addressTxListResp := &AddressTxListResp{}
	addressTxListResp.Total = counter
	addressTxListResp.AddressTxList = make([]*AddressTxResp, 0)
	for _, transactoin := range transactoins {
		addressTxListResp.AddressTxList = append(addressTxListResp.AddressTxList, &AddressTxResp{
			TxHash:    transactoin.Hash,
			ChainName: ChainId2Name(transactoin.ChainId),
			ChainLogo: ChainId2ChainCache(transactoin.ChainId).ChainLogo,
			From:      basedef.Hash2Address(transactoin.ChainId, transactoin.From),
			To:        basedef.Hash2Address(transactoin.ChainId, transactoin.To),
			Amount:    FormatAmount(transactoin.Precision, transactoin.Amount),
			Height:    uint32(transactoin.Height),
			TT:        uint32(transactoin.Time),
			Direct:    transactoin.Direct,
			TokenHash: transactoin.TokenHash,
			TokenName: transactoin.TokenName,
			TokenType: transactoin.TokenType,
			TokenLogo: transactoin.Meta,
		})
	}
	return addressTxListResp
}

type TransferStatisticResp struct {
	Name             string
	ChainId          uint64
	SourceName       string
	Hash             string
	Amount           *big.Int
	AmountBtc        *big.Int
	AmountBtcPrecent string   `json:"amount_btc_precent"`
	AmountUsd        *big.Int `json:"amount_usd"`
	AmountUsdPrecent string   `json:"Amount_usd_precent"`
}

type AllTransferStatisticResp struct {
	ChainTransferStatistics []*ChainTransferStatisticResp `json:"chain_transfer_statistics"`
	AmountUsd1              *big.Int
	AmountUsd               string `json:"amounts_usd"`
	Addresses               uint32 `json:"addresses"`
	Transactions            uint32 `json:"transactions"`
}
type ChainTransferStatisticResp struct {
	Chain                   uint32 `json:"chainid"`
	ChainName               string `json:"chainname"`
	ChainLogo               string `json:"chainlogo"`
	AmountBtc               string `json:"amount_btc"`
	AmountUsd               string `json:"amount_usd"`
	TxAmount                int64  `json:"txAmount"`
	AmountUsd1              *big.Int
	In                      uint32                        `json:"in"`
	Out                     uint32                        `json:"out"`
	Addresses               uint32                        `json:"addresses"`
	Height                  uint32                        `json:"blockheight"`
	AssetTransferStatistics []*AssetTransferStatisticResp `json:"asset_transfer_statistics"`
}
type AssetTransferStatisticResp struct {
	Name             string `json:"name"`
	Token            string `json:"token"`
	TokenLogo        string `json:"tokenlogo"`
	Hash             string `json:"hash"`
	Amount           string `json:"amount"`
	Amount1          *big.Int
	AmountBtc        string `json:"amount_btc"`
	AmountUsd        string `json:"amount_usd"`
	AmountUsdPrecent string `json:"Amount_usd_precent"`
	AmountUsd1       *big.Int
	SourceName       string `json:"source_name"`
	SourceChain      uint32 `json:"source_chainid"`
	SourceChainName  string `json:"source_chainname"`
	SourceChainLogo  string `json:"source_chainlogo"`
}

func MakeTransferInfoResp(tokenStatistics []*TokenStatistic, chainStatistics []*ChainStatistic, chains []*Chain) *AllTransferStatisticResp {
	allTransferStatistic := new(AllTransferStatisticResp)
	allTransferStatistic.ChainTransferStatistics = make([]*ChainTransferStatisticResp, 0)

	allAmountUsdTotal := new(big.Int).SetInt64(0)
	for _, tokenStatistic := range tokenStatistics {
		if tokenStatistic.Token != nil && tokenStatistic.Token.TokenBasic != nil &&
			tokenStatistic.ChainId == tokenStatistic.Token.TokenBasic.ChainId {
			allAmountUsdTotal.Add(allAmountUsdTotal, &tokenStatistic.InAmountUsd.Int)
		}
	}
	allAddress := uint32(0)
	allTransactions := uint32(0)
	for _, chainStatistic := range chainStatistics {
		if chainStatistic.ChainId == uint64(0) {
			continue
		}
		amountBtcTotal := new(big.Int).SetInt64(0)
		amountUsdTotal := new(big.Int).SetInt64(0)
		totalHeight := uint64(0)
		for _, chain := range chains {
			if chainStatistic.ChainId == chain.ChainId {
				totalHeight += chain.Height
			}
		}
		assetTransferStatisticResps := make([]*AssetTransferStatisticResp, 0)
		for _, tokenStatistic := range tokenStatistics {
			if tokenStatistic.ChainId == chainStatistic.ChainId && tokenStatistic.ChainId != tokenStatistic.Token.TokenBasic.ChainId {
				amount := new(big.Int).Sub(&tokenStatistic.InAmount.Int, &tokenStatistic.OutAmount.Int)
				amountBtc := new(big.Int).Sub(&tokenStatistic.InAmountBtc.Int, &tokenStatistic.OutAmountBtc.Int)
				amountUsd := new(big.Int).Sub(&tokenStatistic.InAmountUsd.Int, &tokenStatistic.OutAmountUsd.Int)
				if amountUsd.Cmp(big.NewInt(0)) == 1 {
					amountBtcTotal = new(big.Int).Add(amountBtcTotal, amountBtc)
					amountUsdTotal = new(big.Int).Add(amountUsdTotal, amountUsd)
				}
				assetTransferStatisticResp := &AssetTransferStatisticResp{
					Amount:     FormatAmount(2, NewBigInt(amount)),
					AmountBtc:  FormatAmount(4, NewBigInt(amountBtc)),
					AmountUsd:  FormatAmount(4, NewBigInt(amountUsd)),
					AmountUsd1: amountUsd,
				}
				if tokenStatistic.Token != nil {
					assetTransferStatisticResp.Name = tokenStatistic.Token.Name
					assetTransferStatisticResp.Hash = tokenStatistic.Hash
					if tokenStatistic.Token.TokenBasic != nil {
						assetTransferStatisticResp.TokenLogo = tokenStatistic.Token.TokenBasic.Meta
						assetTransferStatisticResp.SourceChainName = ChainId2Name(tokenStatistic.Token.TokenBasic.ChainId)
						assetTransferStatisticResp.SourceChainLogo = ChainId2ChainCache(tokenStatistic.Token.TokenBasic.ChainId).ChainLogo
					} else {
						logs.Info("MakeTransferInfoResp tokenStatistic_Token_TokenBasic!=nil nil %v,%v", tokenStatistic.ChainId, tokenStatistic.Hash)
					}
				} else {
					logs.Info("MakeTransferInfoResp tokenStatistic_Token nil %v,%v", tokenStatistic.ChainId, tokenStatistic.Hash)
				}
				assetTransferStatisticResps = append(assetTransferStatisticResps, assetTransferStatisticResp)
			}
		}

		for _, assetTransferStatisticResp := range assetTransferStatisticResps {
			if assetTransferStatisticResp.AmountUsd1.Cmp(big.NewInt(0)) == 1 {
				assetTransferStatisticResp.AmountUsdPrecent = Precent(assetTransferStatisticResp.AmountUsd1.Uint64(), amountUsdTotal.Uint64())
			}
		}
		allAddress += uint32(chainStatistic.Addresses)

		chainTransferStatisticResp := &ChainTransferStatisticResp{
			Chain:                   uint32(chainStatistic.ChainId),
			ChainName:               ChainId2Name(chainStatistic.ChainId),
			ChainLogo:               ChainId2ChainCache(chainStatistic.ChainId).ChainLogo,
			AmountBtc:               FormatAmount(4, NewBigInt(amountBtcTotal)),
			AmountUsd:               FormatAmount(4, NewBigInt(amountUsdTotal)),
			AmountUsd1:              amountUsdTotal,
			In:                      uint32(chainStatistic.In),
			Out:                     uint32(chainStatistic.Out),
			TxAmount:                chainStatistic.In + chainStatistic.Out,
			Addresses:               uint32(chainStatistic.Addresses),
			Height:                  uint32(totalHeight),
			AssetTransferStatistics: assetTransferStatisticResps,
		}
		allTransactions += uint32(chainStatistic.In) + uint32(chainStatistic.Out)
		allTransferStatistic.ChainTransferStatistics = append(allTransferStatistic.ChainTransferStatistics, chainTransferStatisticResp)
	}
	sort.Slice(allTransferStatistic.ChainTransferStatistics, func(i, j int) bool {
		return allTransferStatistic.ChainTransferStatistics[i].AmountUsd1.Cmp(allTransferStatistic.ChainTransferStatistics[j].AmountUsd1) == 1
	})
	for _, statis := range allTransferStatistic.ChainTransferStatistics {
		sort.Slice(statis.AssetTransferStatistics, func(i, j int) bool {
			return statis.AssetTransferStatistics[i].AmountUsd1.Cmp(statis.AssetTransferStatistics[j].AmountUsd1) == 1
		})
	}
	allTransferStatistic.AmountUsd = FormatAmount(4, NewBigInt(allAmountUsdTotal))
	allTransferStatistic.Addresses = allAddress
	allTransferStatistic.Transactions = allTransactions
	return allTransferStatistic
}

type AssetStatisticResp struct {
	Name              string `json:"name"`
	TokenLogo         string `json:"tokenlogo"`
	AddressNum        uint64 `json:"addressnumber"`
	AddressNumPrecent string `json:"addressnumber_precent"`
	Amount            string `json:"amount"`
	AmountBtc         string `json:"amount_btc"`
	AmountBtcPrecent  string `json:"amount_btc_precent"`
	AmountUsd         string `json:"amount_usd"`
	AmountUsd1        *big.Int
	AmountUsdPrecent  string `json:"Amount_usd_precent"`
	TxNum             uint64 `json:"txnumber"`
	TxNumPrecent      string `json:"txnumber_precent"`
}
type AssetInfoResp struct {
	AmountBtcTotal  string                `json:"amount_btc_total"`
	AmountUsdTotal  string                `json:"amount_usd_total"`
	AssetStatistics []*AssetStatisticResp `json:"asset_statistics"`
}

func MakeAssetInfoResp(assetStatistics []*AssetStatistic) *AssetInfoResp {
	assetInfo := new(AssetInfoResp)
	amountBtcTotal := new(big.Int)
	amountUsdTotal := new(big.Int)
	addressNumberTotal := uint64(0)
	txNumTotal := uint64(0)

	for _, assetStatistic := range assetStatistics {
		amountBtcTotal.Add(amountBtcTotal, &assetStatistic.AmountBtc.Int)
		amountUsdTotal.Add(amountUsdTotal, &assetStatistic.AmountUsd.Int)
		addressNumberTotal += assetStatistic.Addressnum
		txNumTotal += assetStatistic.Txnum
	}

	assetInfo.AmountBtcTotal = FormatAmount(4, NewBigInt(amountBtcTotal))
	assetInfo.AmountUsdTotal = FormatAmount(4, NewBigInt(amountUsdTotal))

	for _, assetStatistic := range assetStatistics {
		assetStatisticResp := &AssetStatisticResp{
			Name:              assetStatistic.TokenBasicName,
			TokenLogo:         assetStatistic.TokenBasic.Meta,
			AddressNum:        assetStatistic.Addressnum,
			AddressNumPrecent: Precent(assetStatistic.Addressnum, addressNumberTotal),
			Amount:            FormatAmount(2, assetStatistic.Amount),
			AmountBtc:         FormatAmount(4, assetStatistic.AmountBtc),
			AmountBtcPrecent:  Precent(assetStatistic.AmountBtc.Uint64(), amountBtcTotal.Uint64()),
			AmountUsd1:        &assetStatistic.AmountUsd.Int,
			AmountUsd:         FormatAmount(4, assetStatistic.AmountUsd),
			AmountUsdPrecent:  Precent(assetStatistic.AmountUsd.Uint64(), amountUsdTotal.Uint64()),
			TxNum:             assetStatistic.Txnum,
			TxNumPrecent:      Precent(assetStatistic.Txnum, txNumTotal),
		}
		assetInfo.AssetStatistics = append(assetInfo.AssetStatistics, assetStatisticResp)
	}
	sort.Slice(assetInfo.AssetStatistics, func(i, j int) bool {
		return assetInfo.AssetStatistics[i].AmountUsd1.Cmp(assetInfo.AssetStatistics[j].AmountUsd1) == 1
	})
	return assetInfo
}

type LockTokenResp struct {
	ChainId      uint64  `json:"chainId"`
	ChainName    string  `json:"chainname"`
	ChainLogo    string  `json:"chainlogo"`
	InAmountUsd  *BigInt `json:"-"`
	InAmountUsd1 string  `json:"amountUsd"`
	TokenNum     int     `json:"tokenNum"`
	ProxyNum     int     `json:"proxyNum"`
}

func MakeLockTokenListResp(lockTokenResps []*LockTokenResp) []*LockTokenResp {
	sort.Slice(lockTokenResps, func(i, j int) bool {
		return (&lockTokenResps[i].InAmountUsd.Int).Cmp(&lockTokenResps[j].InAmountUsd.Int) == 1
	})
	for _, lockTokenResp := range lockTokenResps {
		lockTokenResp.ChainName = ChainId2Name(lockTokenResp.ChainId)
		lockTokenResp.ChainLogo = ChainId2ChainCache(lockTokenResp.ChainId).ChainLogo
		lockTokenResp.InAmountUsd1 = decimal.NewFromBigInt(&lockTokenResp.InAmountUsd.Int, -4).String()
	}
	return lockTokenResps
}

type LockTokenInfoReq struct {
	ChainId uint64 `json:"chainId"`
}
type LockTokenInfoResp struct {
	ChainId          uint64
	ChainName        string
	TokenName        string
	TokenLogo        string `json:"tokenlogo"`
	Hash             string
	ItemProxy        string
	AmountUsd        string
	ChainLogo        string `json:"chainlogo"`
	ChainExplorerUrl string `json:"chainexplorerurl"`
}

func MakeLockTokenInfoResp(lockTokenStatistics []*LockTokenStatistic) []*LockTokenInfoResp {
	sort.Slice(lockTokenStatistics, func(i, j int) bool {
		return (&lockTokenStatistics[i].InAmountUsd.Int).Cmp(&lockTokenStatistics[j].InAmountUsd.Int) == 1
	})
	lockTokenInfoResps := make([]*LockTokenInfoResp, 0)

	for _, lockTokenStatistic := range lockTokenStatistics {
		resp := new(LockTokenInfoResp)
		resp.ChainId = lockTokenStatistic.ChainId
		resp.ChainName = ChainId2Name(resp.ChainId)
		resp.ChainLogo = ChainId2ChainCache(resp.ChainId).ChainLogo
		resp.ChainExplorerUrl = ChainId2ChainCache(resp.ChainId).ChainExplorerUrl
		resp.Hash = lockTokenStatistic.Hash
		resp.TokenName = lockTokenStatistic.Token.Name
		resp.TokenLogo = lockTokenStatistic.Token.TokenBasic.Meta
		resp.ItemProxy = basedef.Proxy2Address(resp.ChainId, lockTokenStatistic.ItemProxy)
		resp.AmountUsd = decimal.NewFromBigInt(&lockTokenStatistic.InAmountUsd.Int, -4).String()
		lockTokenInfoResps = append(lockTokenInfoResps, resp)
	}
	return lockTokenInfoResps
}

type NftSignReq struct {
	Address string `json:"address"`
}

type EvmosEthNftInfoReq struct {
	PageSize int
	PageNo   int
}

type EvmosEthNftInfoResp struct {
	ContractWithExploit string
	HeightWithExploit   uint64
	EvmosEthNftInfos    []*EvmosEthNftInfo
	TotalNum            int64
}

type EvmosEthNftInfo struct {
	Address    string
	Amount_ETH string
}

func MakeEvmosEthNftInfoResp(nftUsers []NftUser, total int64) *EvmosEthNftInfoResp {
	evmosEthNftInfoResp := new(EvmosEthNftInfoResp)
	evmosEthNftInfoResp.HeightWithExploit = 12996659
	evmosEthNftInfoResp.ContractWithExploit = "0x838bf9E95CB12Dd76a54C9f9D2E3082EAF928270"
	evmosEthNftInfoResp.TotalNum = total
	for _, v := range nftUsers {
		evmosEthNftInfo := new(EvmosEthNftInfo)
		amount_u := new(big.Float).Quo(new(big.Float).SetInt(&v.EffectAmountUsd.Int), new(big.Float).SetInt64(10000))
		amount_eth := new(big.Float).Quo(amount_u, new(big.Float).SetFloat64(2607.47))
		evmosEthNftInfo.Amount_ETH = fmt.Sprintf("%.8f", amount_eth)
		evmosEthNftInfo.Address = common.HexToAddress(v.DfAddress).Hex()
		evmosEthNftInfoResp.EvmosEthNftInfos = append(evmosEthNftInfoResp.EvmosEthNftInfos, evmosEthNftInfo)
	}
	return evmosEthNftInfoResp
}
