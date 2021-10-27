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
	"poly-bridge/bridge_tools/conf"
	serverconf "poly-bridge/conf"
	"poly-bridge/crosschaindao"
	"poly-bridge/models"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func startDeploy(cfg *conf.DeployConfig, servercfg *serverconf.Config) {
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
	err = db.Debug().AutoMigrate(
		&models.ChainFee{},
		&models.Chain{},
		&models.DstSwap{},
		&models.DstTransaction{},
		&models.DstTransfer{},
		&models.NFTProfile{},
		&models.PolyTransaction{},
		&models.PriceMarket{},
		&models.SrcSwap{},
		&models.SrcTransaction{},
		&models.SrcTransfer{},
		&models.TimeStatistic{},
		&models.TokenBasic{},
		&models.TokenMap{},
		&models.Token{},
		&models.WrapperTransaction{},
	)
	if err != nil {
		panic(err)
	}
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
