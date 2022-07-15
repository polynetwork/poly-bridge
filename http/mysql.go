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

package http

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"poly-bridge/basedef"
	"poly-bridge/conf"
)

var db *gorm.DB
var relayUrl string
var contractCheck map[uint64]([]string)

func Init() {
	config := conf.GlobalConfig.DBConfig
	Logger := logger.Default
	if conf.GlobalConfig.RunMode == "dev" {
		Logger = Logger.LogMode(logger.Info)
	}

	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", config.User, config.Password, config.URL, config.Scheme)
	var err error
	db, err = gorm.Open(mysql.Open(conn), &gorm.Config{Logger: Logger})
	if err != nil {
		panic(err)
	}

	relayUrl = conf.GlobalConfig.RelayUrl
	if relayUrl == "" {
		panic("relayUrl is null")
	}

	contractCheck = make(map[uint64]([]string), 0)
	for _, chainListen := range conf.GlobalConfig.ChainListenConfig {
		contract := make([]string, 0)
		if len(chainListen.WrapperContract) > 0 {
			for _, v := range chainListen.WrapperContract {
				if v != "" {
					contract = append(contract, v)
				}
			}
		}
		if len(chainListen.NFTWrapperContract) > 0 {
			for _, v := range chainListen.NFTWrapperContract {
				if v != "" {
					contract = append(contract, v)
				}
			}
		}
		if chainListen.ChainId == basedef.RIPPLE_CROSSCHAIN_ID && len(chainListen.ProxyContract) > 0 {
			for _, v := range chainListen.ProxyContract {
				if v != "" {
					contract = append(contract, v)
				}
			}
		}
		contractCheck[chainListen.ChainId] = contract
	}
}
