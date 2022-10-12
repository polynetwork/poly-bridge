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
	ARBITRUM_CROSSCHAIN_ID   = uint64(215)
	XDAI_CROSSCHAIN_ID       = uint64(206)
	ZILLIQA_CROSSCHAIN_ID    = uint64(111)
	FANTOM_CROSSCHAIN_ID     = uint64(208)
	AVAX_CROSSCHAIN_ID       = uint64(209)
	OPTIMISTIC_CROSSCHAIN_ID = uint64(210)
	METIS_CROSSCHAIN_ID      = uint64(300)
	BOBA_CROSSCHAIN_ID       = uint64(400)
	RINKEBY_CROSSCHAIN_ID    = uint64(402)
	OASIS_CROSSCHAIN_ID      = uint64(500)
	HARMONY_CROSSCHAIN_ID    = uint64(800)
	KCC_CROSSCHAIN_ID        = uint64(900)
	BYTOM_CROSSCHAIN_ID      = uint64(701)
	HSC_CROSSCHAIN_ID        = uint64(603)
	STARCOIN_CROSSCHAIN_ID   = uint64(318)
	KAVA_CROSSCHAIN_ID       = uint64(920)
	CUBE_CROSSCHAIN_ID       = uint64(930)
	ZKSYNC_CROSSCHAIN_ID     = uint64(940)
	CELO_CROSSCHAIN_ID       = uint64(960)
	CLOVER_CROSSCHAIN_ID     = uint64(970)
	GOERLI_CROSSCHAIN_ID     = uint64(1000002)
	CONFLUX_CROSSCHAIN_ID    = uint64(980)
	RIPPLE_CROSSCHAIN_ID     = uint64(223)
	ASTAR_CROSSCHAIN_ID      = uint64(990)
	ENV                      = "testnet"
)

const (
	BSC_NORMAL_GASPRICE = 5000000000
)

var ETH_CHAINS = []uint64{
	ETHEREUM_CROSSCHAIN_ID, BSC_CROSSCHAIN_ID, HECO_CROSSCHAIN_ID, OK_CROSSCHAIN_ID, MATIC_CROSSCHAIN_ID,
	O3_CROSSCHAIN_ID, PLT_CROSSCHAIN_ID, ARBITRUM_CROSSCHAIN_ID, XDAI_CROSSCHAIN_ID, OPTIMISTIC_CROSSCHAIN_ID,
	FANTOM_CROSSCHAIN_ID, AVAX_CROSSCHAIN_ID, METIS_CROSSCHAIN_ID, BOBA_CROSSCHAIN_ID, RINKEBY_CROSSCHAIN_ID,
	OASIS_CROSSCHAIN_ID, HARMONY_CROSSCHAIN_ID, KCC_CROSSCHAIN_ID, BYTOM_CROSSCHAIN_ID, HSC_CROSSCHAIN_ID,
	KAVA_CROSSCHAIN_ID, CUBE_CROSSCHAIN_ID, ZKSYNC_CROSSCHAIN_ID, CELO_CROSSCHAIN_ID, CLOVER_CROSSCHAIN_ID,
	CONFLUX_CROSSCHAIN_ID, ASTAR_CROSSCHAIN_ID,
}
