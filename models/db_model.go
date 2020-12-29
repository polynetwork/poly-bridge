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
	"database/sql/driver"
	"fmt"
	"math/big"
)

type Chain struct {
	ChainId uint64 `gorm:"primaryKey;type:bigint(20);not null"`
	Name    string `gorm:"size:64"`
	Height  uint64 `gorm:"type:bigint(20);not null"`
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
	SrcTransfer *SrcTransfer `gorm:"foreignKey:Hash;references:Hash"`
}

type SrcTransfer struct {
	Hash       string  `gorm:"primaryKey;size:66;not null"`
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
	SrcTransaction  *SrcTransaction `gorm:"foreignKey:Hash;references:SrcHash"`
	PolyHash        string
	PolyTransaction *PolyTransaction `gorm:"foreignKey:Hash;references:PolyHash"`
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
	DstTransfer *DstTransfer `gorm:"foreignKey:Hash;references:Hash"`
}

type DstTransfer struct {
	Hash    string  `gorm:"primaryKey;size:66;not null"`
	ChainId uint64  `gorm:"type:bigint(20);not null"`
	Time    uint64  `gorm:"type:bigint(20);not null"`
	Asset   string  `gorm:"type:varchar(66);not null"`
	From    string  `gorm:"type:varchar(66);not null"`
	To      string  `gorm:"type:varchar(66);not null"`
	Amount  *BigInt `gorm:"type:varchar(64);not null"`
}

type WrapperTransaction struct {
	Hash         string  `gorm:"primaryKey;size:66;not null"`
	User         string  `gorm:"size:64"`
	SrcChainId   uint64  `gorm:"type:bigint(20);not null"`
	BlockHeight  uint64  `gorm:"type:bigint(20);not null"`
	Time         uint64  `gorm:"type:bigint(20);not null"`
	DstChainId   uint64  `gorm:"type:bigint(20);not null"`
	FeeTokenHash string  `gorm:"size:66;not null"`
	FeeToken     *Token  `gorm:"foreignKey:FeeTokenHash;references:Hash"`
	FeeAmount    *BigInt `gorm:"type:varchar(64);not null"`
	Status       uint64  `gorm:"type:bigint(20);not null"`
}

type TokenBasic struct {
	Name      string   `gorm:"primaryKey;size:64;not null"`
	Precision uint64   `gorm:"type:bigint(20);not null"`
	CmcName   string   `gorm:"size:64;not null"`
	CmcPrice  int64    `gorm:"size:64;not null"`
	CmcInd    uint64   `gorm:"type:bigint(20);not null"`
	BinName   string   `gorm:"size:64;not null"`
	BinPrice  int64    `gorm:"size:64;not null"`
	BinInd    uint64   `gorm:"type:bigint(20);not null"`
	AvgPrice  int64    `gorm:"size:64;not null"`
	AvgInd    uint64   `gorm:"type:bigint(20);not null"`
	Time      uint64   `gorm:"type:bigint(20);not null"`
	Tokens    []*Token `gorm:"foreignKey:TokenBasicName;references:Name"`
}

type ChainFee struct {
	ChainId        uint64      `gorm:"primaryKey;type:bigint(20);not null"`
	TokenBasicName string      `gorm:"size:64;not null"`
	TokenBasic     *TokenBasic `gorm:"foreignKey:TokenBasicName;references:Name"`
	MaxFee         *BigInt     `gorm:"type:varchar(64);not null"`
	MinFee         *BigInt     `gorm:"type:varchar(64);not null"`
	ProxyFee       *BigInt     `gorm:"type:varchar(64);not null"`
}

type Token struct {
	Hash           string      `gorm:"primaryKey;size:66;not null"`
	ChainId        uint64      `gorm:"type:bigint(20);not null"`
	Name           string      `gorm:"size:64;not null"`
	Precision      uint64      `gorm:"type:bigint(20);not null"`
	TokenBasicName string      `gorm:"size:64;not null"`
	TokenBasic     *TokenBasic `gorm:"foreignKey:TokenBasicName;references:Name"`
	TokenMaps      []*TokenMap `gorm:"foreignKey:SrcTokenHash;references:Hash"`
}

type TokenMap struct {
	SrcTokenHash string `gorm:"primaryKey;size:66;not null"`
	SrcToken     *Token `gorm:"foreignKey:SrcTokenHash;references:Hash"`
	DstTokenHash string `gorm:"primaryKey;size:66;not null"`
	DstToken     *Token `gorm:"foreignKey:DstTokenHash;references:Hash"`
}

type BigInt struct {
	big.Int
}

func NewBigIntFromInt(value int64) *BigInt {
	x := new(big.Int).SetInt64(value)
	return NewBigInt(x)
}

func NewBigInt(value *big.Int) *BigInt {
	return &BigInt{Int: *value}
}

func (bigInt *BigInt) Value() (driver.Value, error) {
	if bigInt == nil {
		return "null", nil
	}
	return bigInt.String(), nil
}

func (bigInt *BigInt) Scan(v interface{}) error {
	value, ok := v.([]byte)
	if !ok {
		return fmt.Errorf("type error, %v", v)
	}
	if string(value) == "null" {
		return nil
	}
	data, ok := new(big.Int).SetString(string(value), 10)
	if !ok {
		return fmt.Errorf("not a valid big integer: %s", value)
	}
	bigInt.Int = *data
	return nil
}

func (bigInt *BigInt) MarshalJSON() ([]byte, error) {
	if bigInt == nil {
		return []byte("null"), nil
	}
	return []byte(bigInt.String()), nil
}

func (bigInt *BigInt) UnmarshalJSON(p []byte) error {
	if string(p) == "null" {
		return nil
	}
	data, ok := new(big.Int).SetString(string(p), 10)
	if !ok {
		return fmt.Errorf("not a valid big integer: %s", p)
	}
	bigInt.Int = *data
	return nil
}
