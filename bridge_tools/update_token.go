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
	"poly-bridge/bridge_tools/conf"
	"poly-bridge/cacheRedis"
	serverconf "poly-bridge/conf"
	"poly-bridge/crosschaindao"
	"poly-bridge/models"
	"strings"
	"time"
)

func startUpdateToken(cfg *conf.DeployConfig, servercfg *serverconf.Config) {
	dbCfg := cfg.DBConfig
	Logger := logger.Default
	if dbCfg.Debug == true {
		Logger = Logger.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{Logger: Logger})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.Chain{}, &models.WrapperTransaction{}, &models.ChainFee{}, &models.TokenBasic{}, &models.Token{}, &models.PriceMarket{},
		&models.TokenMap{}, &models.SrcTransaction{}, &models.SrcTransfer{}, &models.PolyTransaction{}, &models.DstTransaction{}, &models.DstTransfer{})
	if err != nil {
		panic(err)
	}
	//
	db.Where("1 = 1").Delete(&models.Chain{})
	db.Where("1 = 1").Delete(&models.ChainFee{})
	db.Where("1 = 1").Delete(&models.PriceMarket{})
	db.Where("1 = 1").Delete(&models.TokenMap{})
	db.Where("1 = 1").Delete(&models.Token{})
	db.Where("1 = 1").Delete(&models.TokenBasic{})
	//
	dao := crosschaindao.NewCrossChainDao(cfg.Server, cfg.Backup, cfg.DBConfig)
	if dao == nil {
		panic("server is invalid")
	}
	for _, tokenBasic := range cfg.TokenBasics {
		for _, token := range tokenBasic.Tokens {
			token.Hash = strings.ToLower(token.Hash)
		}
	}
	dao.AddTokens(cfg.TokenBasics, cfg.TokenMaps, servercfg)
	dao.AddChains(cfg.Chains, cfg.ChainFees)
}

func SetDyingToken(tokenBasicName string, proxyFee int) {
	if ok, err := cacheRedis.Redis.Set(cacheRedis.MarkTokenAsDying+tokenBasicName, proxyFee, 24*time.Hour); err == nil && ok {
		fmt.Printf("set dying token successfully, %v : %v", tokenBasicName, proxyFee)
	} else {
		panic(err)
	}
}

func RemoveDyingToken(tokenBasicName string) {
	if _, err := cacheRedis.Redis.Del(cacheRedis.MarkTokenAsDying + tokenBasicName); err == nil {
		fmt.Printf("remove dying token successfully, %v ", tokenBasicName)
	} else {
		panic(err)
	}
}
