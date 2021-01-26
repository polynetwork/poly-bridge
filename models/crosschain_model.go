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
	ChainId             *uint64 `gorm:"primaryKey;type:bigint(20);not null"`
	Height              uint64  `gorm:"type:bigint(20);not null"`
	BackwardBlockNumber uint64  `gorm:"type:bigint(20);not null"`
}

type SrcTransaction struct {
	Hash        string       `gorm:"primaryKey;size:66;not null"`
	ChainId     uint64       `gorm:"type:bigint(20);not null"`
	State       uint64       `gorm:"type:bigint(20);not null"`
	Time        uint64       `gorm:"type:bigint(20);not null"`
	Fee         *BigInt      `gorm:"type:varchar(64);not null"`
	Height      uint64       `gorm:"type:bigint(20);not null"`
	User        string       `gorm:"type:varchar(66);not null"`
	DstChainId  uint64       `gorm:"type:bigint(20);not null"`
	Contract    string       `gorm:"type:varchar(66);not null"`
	Key         string       `gorm:"type:varchar(8192);not null"`
	Param       string       `gorm:"type:varchar(8192);not null"`
	SrcTransfer *SrcTransfer `gorm:"foreignKey:TxHash;references:Hash"`
}

type SrcTransfer struct {
	TxHash     string  `gorm:"primaryKey;size:66;not null"`
	ChainId    uint64  `gorm:"type:bigint(20);not null"`
	Time       uint64  `gorm:"type:bigint(20);not null"`
	Asset      string  `gorm:"type:varchar(66);not null"`
	From       string  `gorm:"type:varchar(66);not null"`
	To         string  `gorm:"type:varchar(66);not null"`
	Amount     *BigInt `gorm:"type:varchar(64);not null"`
	DstChainId uint64  `gorm:"type:bigint(20);not null"`
	DstAsset   string  `gorm:"type:varchar(66);not null"`
	DstUser    string  `gorm:"type:varchar(66);not null"`
}

type PolyTransaction struct {
	Hash       string  `gorm:"primaryKey;size:66;not null"`
	ChainId    uint64  `gorm:"type:bigint(20);not null"`
	State      uint64  `gorm:"type:bigint(20);not null"`
	Time       uint64  `gorm:"type:bigint(20);not null"`
	Fee        *BigInt `gorm:"type:varchar(64);not null"`
	Height     uint64  `gorm:"type:bigint(20);not null"`
	SrcChainId uint64  `gorm:"type:bigint(20);not null"`
	SrcHash    string  `gorm:"size:66;not null"`
	DstChainId uint64  `gorm:"type:bigint(20);not null"`
	Key        string  `gorm:"type:varchar(8192);not null"`
}

type PolySrcRelation struct {
	SrcHash         string
	SrcTransaction  *SrcTransaction `gorm:"foreignKey:SrcHash;references:Hash"`
	PolyHash        string
	PolyTransaction *PolyTransaction `gorm:"foreignKey:PolyHash;references:Hash"`
}

type DstTransaction struct {
	Hash        string       `gorm:"primaryKey;size:66;not null"`
	ChainId     uint64       `gorm:"type:bigint(20);not null"`
	State       uint64       `gorm:"type:bigint(20);not null"`
	Time        uint64       `gorm:"type:bigint(20);not null"`
	Fee         *BigInt      `gorm:"type:varchar(64);not null"`
	Height      uint64       `gorm:"type:bigint(20);not null"`
	SrcChainId  uint64       `gorm:"type:bigint(20);not null"`
	Contract    string       `gorm:"type:varchar(66);not null"`
	PolyHash    string       `gorm:"size:66;not null"`
	DstTransfer *DstTransfer `gorm:"foreignKey:TxHash;references:Hash"`
}

type DstTransfer struct {
	TxHash  string  `gorm:"primaryKey;size:66;not null"`
	ChainId uint64  `gorm:"type:bigint(20);not null"`
	Time    uint64  `gorm:"type:bigint(20);not null"`
	Asset   string  `gorm:"type:varchar(66);not null"`
	From    string  `gorm:"type:varchar(66);not null"`
	To      string  `gorm:"type:varchar(66);not null"`
	Amount  *BigInt `gorm:"type:varchar(64);not null"`
}

type WrapperTransaction struct {
	Hash         string  `gorm:"primaryKey;size:66;not null"`
	User         string  `gorm:"type:varchar(66);not null"`
	SrcChainId   uint64  `gorm:"type:bigint(20);not null"`
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
	ChainId            uint64          `gorm:"type:bigint(20);not null"`
	TokenHash          string          `gorm:"type:varchar(66);not null"`
	Token              *Token          `gorm:"foreignKey:TokenHash,ChainId;references:Hash,ChainId"`
}
