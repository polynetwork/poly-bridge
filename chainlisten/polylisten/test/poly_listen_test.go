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

package test

import (
	"fmt"
	"os"
	"poly-swap/chainlisten"
	"poly-swap/chainlisten/polylisten"
	"poly-swap/conf"
	"poly-swap/dao/stake_dao"
	"testing"
)

func TestPolyListen(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current directory: %s\n", dir)
	config := conf.NewConfig("./../../../conf/config_testnet.json")
	if config == nil {
		panic("read config failed!")
	}
	dao := stake_dao.NewStakeDao()
	polyListen := polylisten.NewPolyChainListen(config.ChainListenConfig.PolyChainListenConfig)
	chainListen := chainlisten.NewChainListen(polyListen, dao)
	chainListen.ListenChain()
}
