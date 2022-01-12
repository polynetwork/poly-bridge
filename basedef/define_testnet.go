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
	POLY_CROSSCHAIN_ID       = uint64(0)
	BTC_CROSSCHAIN_ID        = uint64(1)
	ETHEREUM_CROSSCHAIN_ID   = uint64(2)
	ONT_CROSSCHAIN_ID        = uint64(3)
	NEO_CROSSCHAIN_ID        = uint64(5)
	HECO_CROSSCHAIN_ID       = uint64(7)
	BSC_CROSSCHAIN_ID        = uint64(79)
	O3_CROSSCHAIN_ID         = uint64(82)
	NEO3_CROSSCHAIN_ID       = uint64(88)
	PLT_CROSSCHAIN_ID        = uint64(107)
	OK_CROSSCHAIN_ID         = uint64(200)
	MATIC_CROSSCHAIN_ID      = uint64(202)
	SWITCHEO_CROSSCHAIN_ID   = uint64(1000) // No testnet for cosmos
	ARBITRUM_CROSSCHAIN_ID   = uint64(205)
	XDAI_CROSSCHAIN_ID       = uint64(206)
	ZILLIQA_CROSSCHAIN_ID    = uint64(111)
	FANTOM_CROSSCHAIN_ID     = uint64(208)
	AVAX_CROSSCHAIN_ID       = uint64(209)
	OPTIMISTIC_CROSSCHAIN_ID = uint64(210)
	METIS_CROSSCHAIN_ID      = uint64(300)
	PIXIE_CROSSCHAIN_ID      = uint64(316)
	RINKEBY_CROSSCHAIN_ID    = uint64(402)
	BOBA_CROSSCHAIN_ID       = uint64(400)

	ENV = "testnet"
)
