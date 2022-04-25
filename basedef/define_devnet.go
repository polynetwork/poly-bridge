//go:build devnet
// +build devnet

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
	ETHEREUM_CROSSCHAIN_ID   = uint64(2)
	ONT_CROSSCHAIN_ID        = uint64(3)
	NEO_CROSSCHAIN_ID        = uint64(4)
	BSC_CROSSCHAIN_ID        = uint64(6)
	HECO_CROSSCHAIN_ID       = uint64(7)
	O3_CROSSCHAIN_ID         = uint64(80)
	NEO3_CROSSCHAIN_ID       = uint64(88)
	OK_CROSSCHAIN_ID         = uint64(90)
	MATIC_CROSSCHAIN_ID      = uint64(13)
	PLT_CROSSCHAIN_ID        = uint64(107)
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
	OASIS_CROSSCHAIN_ID      = uint64(500)
	STARCOIN_CROSSCHAIN_ID   = uint64(318)
	HARMONY_CROSSCHAIN_ID    = uint64(800)
	HSC_CROSSCHAIN_ID        = uint64(603)
	BCSPALETTE_CROSSCHAIN_ID = uint64(1001)
	KCC_CROSSCHAIN_ID        = uint64(900)
	ONTEVM_CROSSCHAIN_ID     = uint64(333)

	ENV = "devnet"
)

const (
	BSC_NORMAL_GASPRICE = 5000000000
)

var ETH_CHAINS = []uint64{
	ETHEREUM_CROSSCHAIN_ID, BSC_CROSSCHAIN_ID, HECO_CROSSCHAIN_ID, OK_CROSSCHAIN_ID, MATIC_CROSSCHAIN_ID,
	O3_CROSSCHAIN_ID, PLT_CROSSCHAIN_ID, ARBITRUM_CROSSCHAIN_ID, XDAI_CROSSCHAIN_ID, OPTIMISTIC_CROSSCHAIN_ID,
	FANTOM_CROSSCHAIN_ID, AVAX_CROSSCHAIN_ID, METIS_CROSSCHAIN_ID, BOBA_CROSSCHAIN_ID, RINKEBY_CROSSCHAIN_ID,
	OASIS_CROSSCHAIN_ID, KCC_CROSSCHAIN_ID,
}
