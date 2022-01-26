//go:build testnet
// +build testnet

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

package basedef

const (
	ZION_CROSSCHAIN_ID      = uint64(0)
	ZIONMAIN_CROSSCHAIN_ID  = uint64(1)
	SIDECHAIN_CROSSCHAIN_ID = uint64(77)
	ETHEREUM_CROSSCHAIN_ID  = uint64(2)
	BSC_CROSSCHAIN_ID       = uint64(79)
	HECO_CROSSCHAIN_ID      = uint64(7)
	OK_CROSSCHAIN_ID        = uint64(1012)
	PLT_CROSSCHAIN_ID       = uint64(107)
	MATIC_CROSSCHAIN_ID     = uint64(20016)
	KOVAN_CROSSCHAIN_ID     = uint64(302)
	RINKEBY_CROSSCHAIN_ID   = uint64(402)
	GOERLI_CROSSCHAIN_ID    = uint64(502)
	// not support on zion
	BTC_CROSSCHAIN_ID        = uint64(10000001)
	ONT_CROSSCHAIN_ID        = uint64(10000002)
	NEO_CROSSCHAIN_ID        = uint64(10000003)
	O3_CROSSCHAIN_ID         = uint64(10000004)
	NEO3_CROSSCHAIN_ID       = uint64(10000005)
	SWITCHEO_CROSSCHAIN_ID   = uint64(10000006) // No testnet for cosmos
	ARBITRUM_CROSSCHAIN_ID   = uint64(10000007)
	XDAI_CROSSCHAIN_ID       = uint64(10000008)
	ZILLIQA_CROSSCHAIN_ID    = uint64(10000009)
	FANTOM_CROSSCHAIN_ID     = uint64(10000010)
	AVAX_CROSSCHAIN_ID       = uint64(10000011)
	OPTIMISTIC_CROSSCHAIN_ID = uint64(10000012)

	ENV = "testnet"
)
