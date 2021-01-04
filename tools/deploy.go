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
	"poly-swap/models"
	conf2 "poly-swap/tools/conf"
)

func startDeploy(cfg *conf2.DeployConfig) {
	dbCfg := cfg.DBConfig
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.Chain{}, &models.WrapperTransaction{}, &models.TokenBasic{}, &models.ChainFee{}, &models.Token{}, &models.PriceMarket{},
		&models.TokenMap{}, &models.SrcTransaction{}, &models.SrcTransfer{}, &models.PolyTransaction{}, &models.DstTransaction{}, &models.DstTransfer{})
	if err != nil {
		panic(err)
	}
	//
	db.Debug().Save(cfg.Chains)
	db.Save(cfg.TokenBasics)
	db.Save(cfg.ChainFees)
	tokenMaps := getTokenMapsFromToken(cfg.TokenBasics)
	tokenMaps = append(tokenMaps, cfg.TokenMaps...)
	db.Save(tokenMaps)
}

func getTokenMapsFromToken(tokenBasics []*models.TokenBasic) []*models.TokenMap {
	tokenMaps := make([]*models.TokenMap, 0)
	for _, tokenBasic := range tokenBasics {
		for _, tokenSrc := range tokenBasic.Tokens {
			for _, tokenDst := range tokenBasic.Tokens {
				if tokenDst.ChainId != tokenSrc.ChainId {
					tokenMaps = append(tokenMaps, &models.TokenMap{
						SrcTokenHash: tokenSrc.Hash,
						DstTokenHash: tokenDst.Hash,
					})
				}
			}
		}
	}
	return tokenMaps
}
