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

package main

type Config struct {
	Ethereum *ChainConfig
	Bsc      *ChainConfig
	Heco     *ChainConfig
	Poly     *PolyConfig

	// leveldb direction
	LevelDB string

	// oss
	OSS string
}

type ChainConfig struct {
	SideChainID   uint64 // 注册在poly上的侧链ID，这个id同时也必须是genesis.json中的chainId，尤其是bsc，会根据这个校验header.
	SideChainName string
	RPC           string
	Admin         string
	Keystore      string

	ECCD string
	ECCM string
	CCMP string

	NFTLockProxy string
	NFTWrap      string
	FeeToken     string
	FeeCollector string
}

type PolyConfig struct {
	RPC        string
	Keystore   string
	Passphrase string
}
