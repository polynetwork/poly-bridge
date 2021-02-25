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
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"poly-bridge/conf"
	"poly-bridge/models"
	"strings"
)

func startUpdateToken(cfg *conf.DeployConfig) {
	dbCfg := cfg.DBConfig
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.Debug().AutoMigrate(&models.Chain{}, &models.WrapperTransaction{}, &models.ChainFee{}, &models.TokenBasic{}, &models.Token{}, &models.PriceMarket{},
		&models.TokenMap{}, &models.SrcTransaction{}, &models.SrcTransfer{}, &models.PolyTransaction{}, &models.DstTransaction{}, &models.DstTransfer{})
	if err != nil {
		panic(err)
	}
	//
	db.Where("1 = 1").Delete(&models.ChainFee{})
	db.Where("1 = 1").Delete(&models.PriceMarket{})
	db.Where("1 = 1").Delete(&models.TokenMap{})
	db.Where("1 = 1").Delete(&models.Token{})
	db.Where("1 = 1").Delete(&models.TokenBasic{})
	//
	for _, tokenBasic := range cfg.TokenBasics {
		for _, token := range tokenBasic.Tokens {
			token.Hash = strings.ToLower(token.Hash)
		}
	}
	for _, tokenMap := range cfg.TokenMaps {
		tokenMap.SrcToken.Hash = strings.ToLower(tokenMap.SrcToken.Hash)
		tokenMap.DstToken.Hash = strings.ToLower(tokenMap.DstToken.Hash)
	}
	db.Save(cfg.TokenBasics)
	db.Save(cfg.ChainFees)
	tokenMaps := getTokenMapsFromToken(cfg.TokenBasics)
	tokenMaps = append(tokenMaps, cfg.TokenMaps...)
	db.Save(tokenMaps)
}
