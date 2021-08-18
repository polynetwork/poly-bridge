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

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/crosschaindao"
	"poly-bridge/crosschainlisten"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/urfave/cli"
)

const (
	FETCH_BLOCK = "fetch_block"
)

func executeMethod(method string, ctx *cli.Context) {
	configFile := ctx.GlobalString(getFlagName(configPathFlag))
	fmt.Println("configFile", configFile)
	config := conf.NewConfig(configFile)
	if config == nil {
		logs.Error("startServer - read config failed!")
		return
	}

	switch method {
	case FETCH_BLOCK:
		fetchBlock(config)
	case "bingfaSWTH":
		bingfaSWTH(config)
	default:
		fmt.Printf("Available methods: \n %s", strings.Join([]string{FETCH_BLOCK}, "\n"))
	}
}

func fetchBlock(config *conf.Config) {
	height, _ := strconv.Atoi(os.Getenv("BR_HEIGHT"))
	chain, _ := strconv.Atoi(os.Getenv("BR_CHAIN"))
	save := os.Getenv("BR_SAVE")
	if height == 0 || chain == 0 {
		panic(fmt.Sprintf("Invalid param chain %d height %d", chain, height))
	}

	dao := crosschaindao.NewCrossChainDao(basedef.SERVER_POLY_BRIDGE, false, config.DBConfig)
	if dao == nil {
		panic("server is not valid")
	}

	var handle crosschainlisten.ChainHandle
	for _, cfg := range config.ChainListenConfig {
		if int(cfg.ChainId) == chain {
			handle = crosschainlisten.NewChainHandle(cfg)
			break
		}
	}
	if handle == nil {
		panic(fmt.Sprintf("chain %d handler is invalid", chain))
	}
	wrapperTransactions, srcTransactions, polyTransactions, dstTransactions, locks, unlocks, err := handle.HandleNewBlock(uint64(height))
	if err != nil {
		panic(fmt.Sprintf("HandleNewBlock %d err: %v", height, err))
	}
	if save == "true" {
		err = dao.UpdateEvents(nil, wrapperTransactions, srcTransactions, polyTransactions, dstTransactions)
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf(
		"Fetch block events success chain %d height %d wrapper %d src %d poly %d dst %d  locks %d unlocks %d \n",
		chain, height, len(wrapperTransactions), len(srcTransactions), len(polyTransactions), len(dstTransactions), locks, unlocks,
	)
	for i, tx := range wrapperTransactions {
		fmt.Printf("wrapper %d: %+v\n", i, *tx)
	}
	for i, tx := range srcTransactions {
		fmt.Printf("src %d: %+v srcTransfer:%+v\n", i, *tx, tx.SrcTransfer)
	}
	for i, tx := range polyTransactions {
		fmt.Printf("poly %d: %+v\n", i, *tx)
	}
	for i, tx := range dstTransactions {
		fmt.Printf("dst %d: %+v\n", i, *tx)
	}
}

func bingfaSWTH(config *conf.Config) {
	dao := crosschaindao.NewCrossChainDao(basedef.SERVER_POLY_BRIDGE, false, config.DBConfig)
	if dao == nil {
		panic("server is not valid")
	}
	var handle crosschainlisten.ChainHandle
	for _, cfg := range config.ChainListenConfig {
		if cfg.ChainId == basedef.SWITCHEO_CROSSCHAIN_ID {
			handle = crosschainlisten.NewChainHandle(cfg)
			break
		}
	}
	if handle == nil {
		panic(fmt.Sprintf("chain %d handler is invalid", basedef.SWITCHEO_CROSSCHAIN_ID))
	}
	Logger := logger.Default
	dbCfg := config.DBConfig
	if dbCfg.Debug == true {
		Logger = Logger.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{Logger: Logger})
	srcHeights := make([]int, 0)
	dstHeights := make([]int, 0)
	err = db.Table("src_transactions").
		Select("height").
		Where("chain_id = ?", basedef.SWITCHEO_CROSSCHAIN_ID).
		Find(&srcHeights).Error
	if err != nil {
		panic(fmt.Sprintf("bingfaSWTH db Find(&inHeights) err:%v", err))
	}
	fmt.Println("bingfaSWTH Find(&srcHeights)", srcHeights[:3])
	err = db.Table("dst_transactions").
		Select("height").
		Where("chain_id = ?", basedef.SWITCHEO_CROSSCHAIN_ID).
		Find(&dstHeights).Error
	if err != nil {
		panic(fmt.Sprintf("bingfaSWTH db Find(&dstHeights) err:%v", err))
	}
	fmt.Println("bingfaSWTH Find(&dstHeights)", dstHeights[:3])
	heights := srcHeights
	heights = append(heights, dstHeights...)

	for _, height := range heights {
		wrapperTransactions, srcTransactions, polyTransactions, dstTransactions, _, _, err := handle.HandleNewBlock(uint64(height))
		if err != nil {
			panic(fmt.Sprintf("bingfaSWTH HandleNewBlock %d err: %v", height, err))
		}
		err = dao.UpdateEvents(nil, wrapperTransactions, srcTransactions, polyTransactions, dstTransactions)
		if err != nil {
			panic(fmt.Sprintf("bingfaSWTH bingfaSWTH panic panicHeight:%v,flagerr is:%v", height, err))
		}
		fmt.Printf("bingfaSWTH ing.....nowHeight:%v /n", height)
	}
}
