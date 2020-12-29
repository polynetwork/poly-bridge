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

import "math/big"

type ECCMLockEvent struct {
	Method   string
	Txid     string
	TxHash   string
	User     string
	Tchain   uint32
	Contract string
	Height   uint64
	Value    []byte
	Fee      uint64
}
type ECCMUnlockEvent struct {
	Method   string
	TxHash   string
	RTxHash  string
	FChainId uint32
	Contract string
	Height   uint64
	Fee      uint64
}
type ProxyLockEvent struct {
	Method        string
	TxHash        string
	FromAddress   string
	FromAssetHash string
	ToChainId     uint32
	ToAssetHash   string
	ToAddress     string
	Amount        *big.Int
}
type ProxyUnlockEvent struct {
	Method      string
	TxHash      string
	ToAssetHash string
	ToAddress   string
	Amount      *big.Int
}
