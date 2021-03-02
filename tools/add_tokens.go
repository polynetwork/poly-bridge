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
	"encoding/json"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"poly-bridge/conf"
	"poly-bridge/crosschaindao"
	"poly-bridge/models"
	"strings"
)

func startAddToken(cfg *conf.DeployConfig) {
	dbCfg := cfg.DBConfig
	Logger := logger.Default
	if dbCfg.Debug == true {
		Logger = Logger.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{Logger:Logger})
	if err != nil {
		panic(err)
	}
	//
	tokenBasics := make([]*models.TokenBasic, 0)
	db.Model(&models.TokenBasic{}).Preload("PriceMarkets").Preload("Tokens").Find(&tokenBasics)
	name2TokenBasic := make(map[string]*models.TokenBasic, 0)
	for _, tokenBasic := range tokenBasics {
		name2TokenBasic[tokenBasic.Name] = tokenBasic
	}
	//
	addTokenBasics := make([]*models.TokenBasic, 0)
	for _, tokenBasic := range cfg.TokenBasics {
		_, ok := name2TokenBasic[tokenBasic.Name]
		if !ok {
			for _, token := range tokenBasic.Tokens {
				token.Hash = strings.ToLower(token.Hash)
			}
			addTokenBasics = append(addTokenBasics, tokenBasic)
		}
	}
	db.Save(addTokenBasics)
	tokenMaps := getTokenMapsFromToken(addTokenBasics)
	db.Save(tokenMaps)
}

func startAddToken2(cfg *conf.Config, path string) {
	dao := crosschaindao.NewCrossChainDao(cfg.Server, cfg.DBConfig)
	if dao == nil {
		panic("server is invalid")
	}
	//
	tokenBasics := make([]*models.TokenBasic, 0)
	{
		tokenBasicsData := readFile(path + "/add_tokens.json")
		if len(tokenBasicsData) > 0 {
			err := json.Unmarshal(tokenBasicsData, &tokenBasics)
			if err != nil {
				panic(err)
			}
		} else {
			tokenBasics = nil
		}
	}
	for _, tokenBasic := range tokenBasics {
		for _, token := range tokenBasic.Tokens {
			token.Hash = strings.ToLower(token.Hash)
		}
	}
	err := dao.AddTokens(tokenBasics)
	if err != nil {
		panic(err)
	}
}
