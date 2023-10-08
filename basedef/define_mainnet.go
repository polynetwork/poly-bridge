//go:build mainnet
// +build mainnet

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
	NEO_CROSSCHAIN_ID        = uint64(4)
	SWITCHEO_CROSSCHAIN_ID   = uint64(5)
	BSC_CROSSCHAIN_ID        = uint64(6)
	HECO_CROSSCHAIN_ID       = uint64(7)
	PLT_CROSSCHAIN_ID        = uint64(8)
	O3_CROSSCHAIN_ID         = uint64(10)
	OK_CROSSCHAIN_ID         = uint64(12)
	NEO3_CROSSCHAIN_ID       = uint64(14)
	MATIC_CROSSCHAIN_ID      = uint64(17)
	ZILLIQA_CROSSCHAIN_ID    = uint64(18)
	ARBITRUM_CROSSCHAIN_ID   = uint64(19)
	XDAI_CROSSCHAIN_ID       = uint64(20)
	AVAX_CROSSCHAIN_ID       = uint64(21)
	FANTOM_CROSSCHAIN_ID     = uint64(22)
	OPTIMISTIC_CROSSCHAIN_ID = uint64(23)
	METIS_CROSSCHAIN_ID      = uint64(24)
	BOBA_CROSSCHAIN_ID       = uint64(25)
	RINKEBY_CROSSCHAIN_ID    = uint64(1000001)
	GOERLI_CROSSCHAIN_ID     = uint64(1000002)
	OASIS_CROSSCHAIN_ID      = uint64(26)
	HARMONY_CROSSCHAIN_ID    = uint64(27)
	HSC_CROSSCHAIN_ID        = uint64(28)
	BYTOM_CROSSCHAIN_ID      = uint64(29)
	KCC_CROSSCHAIN_ID        = uint64(30)
	STARCOIN_CROSSCHAIN_ID   = uint64(31)
	KAVA_CROSSCHAIN_ID       = uint64(32)
	CUBE_CROSSCHAIN_ID       = uint64(35)
	CELO_CROSSCHAIN_ID       = uint64(36)
	CLOVER_CROSSCHAIN_ID     = uint64(37)
	CONFLUX_CROSSCHAIN_ID    = uint64(38)
	RIPPLE_CROSSCHAIN_ID     = uint64(39)
	ASTAR_CROSSCHAIN_ID      = uint64(40)
	APTOS_CROSSCHAIN_ID      = uint64(41)
	BRISE_CROSSCHAIN_ID      = uint64(42)
	DEXIT_CROSSCHAIN_ID      = uint64(43)
	CLOUDTX_CROSSCHAIN_ID    = uint64(44)
	ZKSYNC_CROSSCHAIN_ID     = uint64(45)
	XINFIN_CROSSCHAIN_ID     = uint64(46)
	ONTEVM_CROSSCHAIN_ID     = uint64(47)

	ENV = "mainnet"
)

const (
	BSC_NORMAL_GASPRICE   = 5000000000
	ASTAR_NORMAL_GASPRICE = 60000000000
)

var ETH_CHAINS = []uint64{
	ETHEREUM_CROSSCHAIN_ID, BSC_CROSSCHAIN_ID, HECO_CROSSCHAIN_ID, OK_CROSSCHAIN_ID, MATIC_CROSSCHAIN_ID,
	O3_CROSSCHAIN_ID, PLT_CROSSCHAIN_ID, ARBITRUM_CROSSCHAIN_ID, XDAI_CROSSCHAIN_ID, OPTIMISTIC_CROSSCHAIN_ID,
	FANTOM_CROSSCHAIN_ID, AVAX_CROSSCHAIN_ID, METIS_CROSSCHAIN_ID, BOBA_CROSSCHAIN_ID, RINKEBY_CROSSCHAIN_ID,
	OASIS_CROSSCHAIN_ID, HARMONY_CROSSCHAIN_ID, KCC_CROSSCHAIN_ID, BYTOM_CROSSCHAIN_ID, HSC_CROSSCHAIN_ID,
	KAVA_CROSSCHAIN_ID, CUBE_CROSSCHAIN_ID, ZKSYNC_CROSSCHAIN_ID, CELO_CROSSCHAIN_ID, CLOVER_CROSSCHAIN_ID,
	CONFLUX_CROSSCHAIN_ID, ASTAR_CROSSCHAIN_ID, BRISE_CROSSCHAIN_ID, DEXIT_CROSSCHAIN_ID, CLOUDTX_CROSSCHAIN_ID,
	XINFIN_CROSSCHAIN_ID, ONTEVM_CROSSCHAIN_ID,
}
