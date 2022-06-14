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
	"context"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/urfave/cli"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"poly-bridge/basedef"
	"poly-bridge/bridge_tools/toolsmethod"
	"poly-bridge/coinpricelisten/coinmarketcap"
	"poly-bridge/conf"
	"poly-bridge/crosschaindao"
	"poly-bridge/crosschainlisten"
	"poly-bridge/models"
	"strconv"
	"strings"
	"time"

	"github.com/devfans/cogroup"
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
	case "initcoinmarketid":
		initcoinmarketid(config)
	case "migrateLockTokenStatisticTable":
		migrateLockTokenStatisticTable(config)
	case "updateZilliqaPolyOldData":
		updateZilliqaPolyOldData(config)
	case "nft":
		toolsmethod.Nft(config)
	default:
		fmt.Printf("Available methods: \n %s", strings.Join([]string{FETCH_BLOCK}, "\n"))
	}
}

func retry(f func() error, count int, duration time.Duration) func(context.Context) error {
	return func(context.Context) error {
		i := 0
		for {
			i++
			if i > count && count != 0 {
				return nil
			}
			err := f()
			if err == nil {
				return nil
			}
			time.Sleep(duration)
		}
		return nil
	}
}

func fetchSingleBlock(chainId, height uint64, handle crosschainlisten.ChainHandle, dao crosschaindao.CrossChainDao, save bool) error {
	wrapperTransactions, srcTransactions, polyTransactions, dstTransactions, locks, unlocks, err := handle.HandleNewBlock(height)
	if err != nil {
		logs.Error(fmt.Sprintf("HandleNewBlock %d err: %v", height, err))
		return err
	}
	if save {
		err = dao.UpdateEvents(wrapperTransactions, srcTransactions, polyTransactions, dstTransactions)
		if err != nil {
			return nil
		}
	}
	fmt.Printf(
		"Fetch block events success chain %d height %d wrapper %d src %d poly %d dst %d  locks %d unlocks %d \n",
		chainId, height, len(wrapperTransactions), len(srcTransactions), len(polyTransactions), len(dstTransactions), locks, unlocks,
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
	return nil
}

func fetchBlock(config *conf.Config) {
	height, _ := strconv.Atoi(os.Getenv("BR_HEIGHT"))
	chain, _ := strconv.Atoi(os.Getenv("BR_CHAIN"))
	save := os.Getenv("BR_SAVE")
	endheight, _ := strconv.Atoi(os.Getenv("END_HEIGHT"))
	if endheight < height {
		endheight = height
	}
	if height == 0 {
		panic(fmt.Sprintf("Invalid param chain %d height %d", chain, height))
	}

	dao := crosschaindao.NewCrossChainDaoWithCheckFeeConfig(basedef.SERVER_POLY_BRIDGE, false, config.DBConfig, config.ChainListenConfig, config.FeeListenConfig)
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

	g := cogroup.Start(context.Background(), 4, 8, false)
	for h := height; h <= endheight; h++ {
		block := uint64(h)
		g.Insert(retry(func() error {
			return fetchSingleBlock(uint64(chain), block, handle, dao, save == "true")
		}, 0, 2*time.Second))
	}
	g.Wait()
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
		err = dao.UpdateEvents(wrapperTransactions, srcTransactions, polyTransactions, dstTransactions)
		if err != nil {
			panic(fmt.Sprintf("bingfaSWTH bingfaSWTH panic panicHeight:%v,flagerr is:%v", height, err))
		}
		fmt.Printf("bingfaSWTH ing.....nowHeight:%v /n", height)
	}
}

func initcoinmarketid(config *conf.Config) {
	Logger := logger.Default
	dbCfg := config.DBConfig
	if dbCfg.Debug == true {
		Logger = Logger.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{Logger: Logger})

	var coinmarketsdk *coinmarketcap.CoinMarketCapSdk
	for _, coinconfig := range config.CoinPriceListenConfig {
		if coinconfig.MarketName == basedef.MARKET_COINMARKETCAP {
			coinmarketsdk = coinmarketcap.NewCoinMarketCapSdk(coinconfig)
			break
		}
	}
	listings, err := coinmarketsdk.ListingsLatest()
	if err != nil {
		logs.Error("coinmarketsdk.ListingsLatest error")
		return
	}
	coinName2Id := make(map[string]int, 0)
	for _, listing := range listings {
		coinName2Id[listing.Name] = listing.ID
	}
	priceMarkets := make([]*models.PriceMarket, 0)
	err = db.Find(&priceMarkets).
		Error
	if err != nil {
		logs.Error("db.Find(&priceMarkets) err", err)
	}
	for _, priceMarket := range priceMarkets {
		if priceMarket.MarketName == basedef.MARKET_COINMARKETCAP {
			if marketid, ok := coinName2Id[priceMarket.Name]; ok {
				err := db.Model(&models.PriceMarket{}).Where("name = ? and market_name= ?", priceMarket.Name, basedef.MARKET_COINMARKETCAP).Update("coin_market_id", marketid).
					Error
				if err != nil {
					logs.Error("Update marketid err %v,Name %v,marketid", err, priceMarket.Name, marketid)
				}
			}
		}
	}
}

func migrateLockTokenStatisticTable(config *conf.Config) {
	Logger := logger.Default
	dbCfg := config.DBConfig
	if dbCfg.Debug == true {
		Logger = Logger.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{Logger: Logger})
	if err != nil {
		logs.Error("Open mysql err", err)
	}
	err = db.Debug().AutoMigrate(
		&models.LockTokenStatistic{},
	)
	checkError(err, "Creating tables")
}

func updateZilliqaPolyOldData(config *conf.Config) {
	tt, err := strconv.ParseInt(os.Getenv("END_TIME"), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("ParseInt err,%v", err))
	}
	chainId, err := strconv.Atoi(os.Getenv("CHAINID"))
	if err != nil {
		panic(fmt.Sprintf("ParseInt err,%v", err))
	}

	Logger := logger.Default
	dbCfg := config.DBConfig
	if dbCfg.Debug == true {
		Logger = Logger.LogMode(logger.Warn)
	}
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{Logger: Logger})
	if err != nil {
		panic(fmt.Sprintf("db err,%v", err))
	}

	var count int
	err = db.Raw("select count(1) from poly_transactions where  src_chain_id=? and `key` ='' and src_hash like '%00000000' and time<?", chainId, tt).
		Find(&count).Error
	if err != nil {
		panic(fmt.Sprintf("db.Raw err,%v", err))
	}
	x := count/100 + 1
	type y struct {
		Id      int64
		SrcHash string
	}
	flag := 0
	for i := 0; i < x; i++ {
		srcs := make([]*y, 0)
		err = db.Raw("select id,src_hash from poly_transactions where  src_chain_id=? and `key` ='' and src_hash like '%00000000' and time<? limit 100", chainId, tt).
			Find(&srcs).Error
		if err != nil {
			panic(fmt.Sprintf("Find srcs err,%v", err))
		}
		if len(srcs) > 0 {
			for _, v := range srcs {
				rightSrcHash := basedef.HexStringReverse(v.SrcHash)
				err = db.Exec("update poly_transactions set src_hash=? where id=? and src_hash=?", rightSrcHash, v.Id, v.SrcHash).
					Error
				if err != nil {
					panic(fmt.Sprintf("update poly_transactions err,%v", err))
				}
			}
			logs.Info("success update id :", srcs[0].Id)
		} else {
			logs.Info("len(srcs)=0,end")
			flag++
		}
		if flag >= 2 {
			logs.Info("flag >=2")
			break
		}
	}
}
