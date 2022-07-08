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

type Chain struct {
	Id                  int64  `gorm:"primaryKey;autoIncrement"`
	ChainId             uint64 `gorm:"uniqueIndex;type:bigint(20);not null"`
	Name                string `gorm:"type:varchar(32)"`
	Height              uint64 `gorm:"type:bigint(20);not null"`
	HeightSwap          uint64 `gorm:"type:bigint(20);not null"`
	BackwardBlockNumber uint64 `gorm:"type:bigint(20);not null"`
	ChainLogo           string `gorm:"type:varchar(128)"`
	ChainExplorerUrl    string `gorm:"type:varchar(128)"`
}

type ChainStatistic struct {
	Id             int64  `gorm:"primaryKey;autoIncrement"`
	ChainId        uint64 `gorm:"uniqueIndex;type:bigint(20);not null"`
	Addresses      int64  `gorm:"type:bigint(20);not null"`
	In             int64  `gorm:"type:bigint(20);not null"`
	Out            int64  `gorm:"type:bigint(20);not null"`
	LastInCheckId  int64  `gorm:"type:int"`
	LastOutCheckId int64  `gorm:"type:int"`
}

type SrcTransaction struct {
	Id          int64        `gorm:"primaryKey;autoIncrement"`
	Hash        string       `gorm:"uniqueIndex;size:66;not null"`
	ChainId     uint64       `gorm:"type:bigint(20);not null"`
	Standard    uint8        `gorm:"type:int(8);not null"`
	State       uint64       `gorm:"type:bigint(20);not null"`
	Time        uint64       `gorm:"type:bigint(20);not null"`
	Fee         *BigInt      `gorm:"type:varchar(64);not null"`
	Height      uint64       `gorm:"type:bigint(20);not null"`
	User        string       `gorm:"type:varchar(66);not null"`
	DstChainId  uint64       `gorm:"type:bigint(20);not null"`
	Contract    string       `gorm:"type:varchar(66);not null"`
	Key         string       `gorm:"index;size:128;not null"`
	Param       string       `gorm:"type:varchar(8192);not null"`
	SrcTransfer *SrcTransfer `gorm:"foreignKey:TxHash;references:Hash"`
	SrcSwap     *SrcSwap     `gorm:"foreignKey:TxHash;references:Hash"`
}

type SrcTransfer struct {
	Id         int64   `gorm:"primaryKey;autoIncrement"`
	TxHash     string  `gorm:"uniqueIndex;size:66;not null"`
	ChainId    uint64  `gorm:"type:bigint(20);not null"`
	Standard   uint8   `gorm:"type:int(8);not null"`
	Time       uint64  `gorm:"type:bigint(20);not null"`
	Asset      string  `gorm:"type:varchar(120);not null"`
	From       string  `gorm:"type:varchar(66);not null"`
	To         string  `gorm:"type:varchar(66);not null"`
	Amount     *BigInt `gorm:"type:varchar(80);not null"`
	DstChainId uint64  `gorm:"type:bigint(20);not null"`
	DstAsset   string  `gorm:"type:varchar(120);not null"`
	DstUser    string  `gorm:"type:varchar(66);not null"`
	Token      *Token  `gorm:"foreignKey:Hash,ChainId;references:Asset,ChainId"`
}

type SrcSwap struct {
	Id         int64   `gorm:"primaryKey;autoIncrement"`
	TxHash     string  `gorm:"uniqueIndex;size:66;not null"`
	ChainId    uint64  `gorm:"type:bigint(20);not null"`
	Time       uint64  `gorm:"type:bigint(20);not null"`
	Asset      string  `gorm:"type:varchar(120);not null"`
	From       string  `gorm:"type:varchar(66);not null"`
	To         string  `gorm:"type:varchar(66);not null"`
	Amount     *BigInt `gorm:"type:varchar(64);not null"`
	PoolId     uint64  `gorm:"type:bigint(20);not null"`
	DstChainId uint64  `gorm:"type:bigint(20);not null"`
	DstAsset   string  `gorm:"type:varchar(120);not null"`
	DstUser    string  `gorm:"type:varchar(66);not null"`
	Type       uint64  `gorm:"type:bigint(20);not null"`
}

type PolyTransaction struct {
	Id          int64   `gorm:"primaryKey;autoIncrement"`
	Hash        string  `gorm:"uniqueIndex;size:66;not null"`
	ChainId     uint64  `gorm:"type:bigint(20);not null"`
	State       uint64  `gorm:"type:bigint(20);not null"`
	Time        uint64  `gorm:"type:bigint(20);not null"`
	Fee         *BigInt `gorm:"type:varchar(64);not null"`
	Height      uint64  `gorm:"type:bigint(20);not null"`
	SrcChainId  uint64  `gorm:"type:bigint(20);not null"`
	SrcHash     string  `gorm:"index;size:66;not null"`
	DstChainId  uint64  `gorm:"type:bigint(20);not null"`
	Key         string  `gorm:"type:varchar(8192);not null"`
	DstSequence uint64  `gorm:"type:bigint(20);not null"`
}

type PolyDetail struct {
	Id          int64   `gorm:"primaryKey;autoIncrement"`
	Hash        string  `gorm:"uniqueIndex;size:66;not null"`
	ChainId     uint64  `gorm:"type:bigint(20);not null"`
	State       uint64  `gorm:"type:bigint(20);not null"`
	Time        uint64  `gorm:"type:bigint(20);not null"`
	Fee         *BigInt `gorm:"type:varchar(64);not null"`
	Height      uint64  `gorm:"type:bigint(20);not null"`
	SrcChainId  uint64  `gorm:"type:bigint(20);not null"`
	SrcHash     string  `gorm:"index;size:66;not null"`
	DstChainId  uint64  `gorm:"type:bigint(20);not null"`
	Key         string  `gorm:"type:varchar(8192);not null"`
	DstSequence uint64  `gorm:"type:bigint(20);not null"`
}

type PolySrcRelation struct {
	SrcHash         string
	SrcTransaction  *SrcTransaction `gorm:"foreignKey:SrcHash;references:Hash"`
	PolyHash        string
	PolyTransaction *PolyTransaction `gorm:"foreignKey:PolyHash;references:Hash"`
}

type DstPolyRelation struct {
	PolyHash        string
	PolyTransaction *PolyTransaction `gorm:"foreignKey:PolyHash;references:Hash"`
	DstHash         string
	DstTransaction  *DstTransaction `gorm:"foreignKey:DstHash;references:Hash"`
}

type DstTransaction struct {
	Id          int64        `gorm:"primaryKey;autoIncrement"`
	Hash        string       `gorm:"uniqueIndex;size:66;not null"`
	ChainId     uint64       `gorm:"type:bigint(20);not null"`
	Standard    uint8        `gorm:"type:int(8);not null"`
	State       uint64       `gorm:"type:bigint(20);not null"`
	Time        uint64       `gorm:"type:bigint(20);not null"`
	Fee         *BigInt      `gorm:"type:varchar(64);not null"`
	Height      uint64       `gorm:"type:bigint(20);not null"`
	SrcChainId  uint64       `gorm:"type:bigint(20);not null"`
	Contract    string       `gorm:"type:varchar(66);not null"`
	PolyHash    string       `gorm:"index;size:66;not null"`
	Sequence    uint64       `gorm:"type:bigint(20);not null"`
	DstTransfer *DstTransfer `gorm:"foreignKey:TxHash;references:Hash"`
	DstSwap     *DstSwap     `gorm:"foreignKey:TxHash;references:Hash"`
}

type DstTransfer struct {
	Id       int64   `gorm:"primaryKey;autoIncrement"`
	TxHash   string  `gorm:"uniqueIndex;size:66;not null"`
	ChainId  uint64  `gorm:"type:bigint(20);not null"`
	Standard uint8   `gorm:"type:int(8);not null"`
	Time     uint64  `gorm:"type:bigint(20);not null"`
	Asset    string  `gorm:"type:varchar(120);not null"`
	From     string  `gorm:"type:varchar(66);not null"`
	To       string  `gorm:"type:varchar(66);not null"`
	Amount   *BigInt `gorm:"type:varchar(80);not null"`
}

type DstSwap struct {
	Id         int64   `gorm:"primaryKey;autoIncrement"`
	TxHash     string  `gorm:"uniqueIndex;size:66;not null"`
	ChainId    uint64  `gorm:"type:bigint(20);not null"`
	Time       uint64  `gorm:"type:bigint(20);not null"`
	PoolId     uint64  `gorm:"type:bigint(20);not null"`
	InAsset    string  `gorm:"type:varchar(66);not null"`
	InAmount   *BigInt `gorm:"type:varchar(64);not null"`
	OutAsset   string  `gorm:"type:varchar(120);not null"`
	OutAmount  *BigInt `gorm:"type:varchar(64);not null"`
	DstChainId uint64  `gorm:"type:bigint(20);not null"`
	DstAsset   string  `gorm:"type:varchar(120);not null"`
	DstUser    string  `gorm:"type:varchar(66);not null"`
	Type       uint64  `gorm:"type:bigint(20);not null"`
}

type WrapperDetail struct {
	Id           int64   `gorm:"primaryKey;autoIncrement"`
	WrapperHash  string  `gorm:"index:wrapper_details_wrapper_hash;size:66;not null"`
	Hash         string  `gorm:"uniqueIndex;size:66;not null"`
	User         string  `gorm:"type:varchar(66);not null"`
	SrcChainId   uint64  `gorm:"type:bigint(20);not null"`
	Standard     uint8   `gorm:"type:int(8);not null"`
	BlockHeight  uint64  `gorm:"type:bigint(20);not null"`
	Time         uint64  `gorm:"type:bigint(20);not null"`
	DstChainId   uint64  `gorm:"type:bigint(20);not null"`
	DstUser      string  `gorm:"type:varchar(66);not null"`
	ServerId     uint64  `gorm:"type:bigint(20);not null"`
	FeeTokenHash string  `gorm:"size:66;not null"`
	FeeAmount    *BigInt `gorm:"type:varchar(64);not null"`
	Status       uint64  `gorm:"type:bigint(20);not null"`
}

type WrapperTransaction struct {
	Id           int64   `gorm:"primaryKey;autoIncrement"`
	Hash         string  `gorm:"uniqueIndex;size:66;not null"`
	User         string  `gorm:"type:varchar(66);not null"`
	SrcChainId   uint64  `gorm:"type:bigint(20);not null"`
	Standard     uint8   `gorm:"type:int(8);not null"`
	BlockHeight  uint64  `gorm:"type:bigint(20);not null"`
	Time         uint64  `gorm:"type:bigint(20);not null"`
	DstChainId   uint64  `gorm:"type:bigint(20);not null"`
	DstUser      string  `gorm:"type:varchar(66);not null"`
	ServerId     uint64  `gorm:"type:bigint(20);not null"`
	FeeTokenHash string  `gorm:"size:66;not null"`
	FeeAmount    *BigInt `gorm:"type:varchar(64);not null"`
	Status       uint64  `gorm:"type:bigint(20);not null"`
}

type SrcPolyDstRelation struct {
	SrcHash            string
	WrapperTransaction *WrapperTransaction `gorm:"foreignKey:SrcHash;references:Hash"`
	SrcTransaction     *SrcTransaction     `gorm:"foreignKey:SrcHash;references:Hash"`
	PolyHash           string
	PolyTransaction    *PolyTransaction `gorm:"foreignKey:PolyHash;references:Hash"`
	DstHash            string
	DstTransaction     *DstTransaction `gorm:"foreignKey:DstHash;references:Hash"`
	ChainId            uint64
	TokenHash          string
	FeeTokenHash       string
	Token              *Token `gorm:"foreignKey:TokenHash,ChainId;references:Hash,ChainId"`
	FeeToken           *Token `gorm:"foreignKey:FeeTokenHash,ChainId;references:Hash,ChainId"`
}

type PolyTxRelation struct {
	SrcHash            string
	WrapperTransaction *WrapperTransaction `gorm:"foreignKey:SrcHash;references:Hash"`
	SrcTransaction     *SrcTransaction     `gorm:"foreignKey:SrcHash;references:Hash"`
	PolyHash           string
	RelatedPolyHash    string
	PolyTransaction    *PolyTransaction `gorm:"foreignKey:PolyHash;references:Hash"`
	DstHash            string
	DstTransaction     *DstTransaction `gorm:"foreignKey:DstHash;references:Hash"`
	ChainId            uint64          `gorm:"type:bigint(20);not null"`
	ToChainId          uint64          `gorm:"type:bigint(20);not null"`
	DstChainId         uint64          `gorm:"type:bigint(20);not null"`
	TokenHash          string          `gorm:"type:varchar(66);not null"`
	ToTokenHash        string          `gorm:"type:varchar(66);not null"`
	DstTokenHash       string          `gorm:"type:varchar(66);not null"`
	FeeTokenHash       string          `gorm:"type:varchar(66);not null"`
	Token              *Token          `gorm:"foreignKey:TokenHash,ChainId;references:Hash,ChainId"`
	FeeToken           *Token          `gorm:"foreignKey:FeeTokenHash,ChainId;references:Hash,ChainId"`
	ToToken            *Token          `gorm:"foreignKey:ToTokenHash,ToChainId;references:Hash,ChainId"`
	DstToken           *Token          `gorm:"foreignKey:DstTokenHash,DstChainId;references:Hash,ChainId"`
}

type AssetStatistic struct {
	Id             int64       `gorm:"primaryKey;autoIncrement"`
	Amount         *BigInt     `gorm:"type:varchar(64);not null"`
	Txnum          uint64      `gorm:"type:bigint(20);not null"`
	Addressnum     uint64      `gorm:"type:bigint(20);not null"`
	TokenBasicName string      `gorm:"uniqueIndex;size:64;not null"`
	AmountBtc      *BigInt     `gorm:"type:varchar(64);not null"`
	AmountUsd      *BigInt     `gorm:"type:varchar(64);not null"`
	LastCheckId    int64       `gorm:"type:int"`
	TokenBasic     *TokenBasic `gorm:"foreignKey:TokenBasicName;references:Name"`
}

type LockTokenStatistic struct {
	Id          int64   `gorm:"primaryKey;autoIncrement"`
	Hash        string  `gorm:"uniqueIndex:idx_locktoken;size:66;not null"`
	ChainId     uint64  `gorm:"uniqueIndex:idx_locktoken;type:bigint(20);not null"`
	ItemProxy   string  `gorm:"uniqueIndex:idx_locktoken;type:varchar(66);not null"`
	ItemName    string  `gorm:"type:varchar(32);not null"`
	InAmount    *BigInt `gorm:"type:varchar(64);not null"`
	InAmountBtc *BigInt `gorm:"type:varchar(64);not null"`
	InAmountUsd *BigInt `gorm:"type:varchar(64);not null"`
	UpdateTime  uint64  `gorm:"type:bigint(20);not null"`
	Token       *Token  `gorm:"foreignKey:Hash,ChainId;references:Hash,ChainId"`
}

type AssetInfo struct {
	Amount         *BigInt
	Txnum          uint64
	Price          int64
	TokenBasicName string
	Precision      uint64
}

type TransactionOnToken struct {
	Hash    string
	ChainId uint64
	Time    uint64
	Height  uint64
	From    string
	To      string
	Amount  *BigInt
	Direct  uint32
}

type TransactionOnAddress struct {
	Hash          string
	ChainId       uint64
	Time          uint64
	Height        uint64
	From          string
	To            string
	Amount        *BigInt
	TokenHash     string
	TokenName     string
	TokenStandard uint8
	Direct        uint32
	Precision     uint64
	Meta          string
}

type NftUser struct {
	Id              int64   `gorm:"primaryKey;autoIncrement"`
	ColChainId      uint64  `gorm:"type:bigint(20);not null"`
	DfChainId       uint64  `gorm:"type:bigint(20)"`
	AddrHash        string  `gorm:"type:varchar(66);not null"`
	ColAddress      string  `gorm:"uniqueIndex:nftusers_coladdress;type:varchar(66);not null"`
	DfAddress       string  `gorm:"index:nftusers_dfaddress;type:varchar(66)"`
	Txnum           uint64  `gorm:"type:bigint(20);not null"`
	FirstTime       uint64  `gorm:"type:bigint(20);not null"`
	TxAmountUsd     *BigInt `gorm:"type:varchar(64);not null"`
	EffectAmountUsd *BigInt `gorm:"type:varchar(64);not null"`
	NftColId        int     `gorm:"index:nftusers_nftcolid;type:int;not null"`
	NftDfId         int     `gorm:"index:nftusers_nftdfid;type:int"`
	NftColsig       string  `gorm:"size:132;not null"`
	NftDfsig        string  `gorm:"size:132"`
	IsClaimCol      uint64  `gorm:"type:bigint(20);not null"`
	IsClaimDf       uint64  `gorm:"type:bigint(20);not null"`
}
