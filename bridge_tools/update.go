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
	"gorm.io/gorm/logger"
	"poly-bridge/basedef"
	"poly-bridge/bridge_tools/conf"
	serverconf "poly-bridge/conf"
	"poly-bridge/crosschaindao"
	"poly-bridge/crosschaindao/swapdao"
	"poly-bridge/models"
	"strings"
)

func startUpdate(cfg *conf.UpdateConfig, servercfg *serverconf.Config) {
	dbCfg := cfg.DBConfig
	Logger := logger.Default
	if dbCfg.Debug == true {
		Logger = Logger.LogMode(logger.Info)
	}
	/*db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{Logger: Logger})
	if err != nil {
		panic(err)
	}
	err = db.Debug().AutoMigrate(&models.Chain{}, &models.WrapperTransaction{}, &models.ChainFee{}, &models.TokenBasic{}, &models.Token{}, &models.PriceMarket{},
		&models.TokenMap{}, &models.SrcTransaction{}, &models.SrcTransfer{}, &models.PolyTransaction{}, &models.DstTransaction{}, &models.DstTransfer{},
		&models.NFTProfile{}, &models.TimeStatistic{})
	if err != nil {
		panic(err)
	}*/
	dao := crosschaindao.NewCrossChainDao(cfg.Server, cfg.Backup, cfg.DBConfig)
	if dao == nil {
		panic("server is invalid")
	}
	//
	for _, tokenBasic := range cfg.TokenBasics {
		for _, token := range tokenBasic.Tokens {
			if token.ChainId != basedef.STARCOIN_CROSSCHAIN_ID {
				token.Hash = strings.ToLower(token.Hash)
			}
		}
	}
	dao.RemoveTokens(cfg.RemoveTokens)
	dao.AddTokens(cfg.TokenBasics, cfg.TokenMaps, servercfg)
	dao.AddChains(cfg.Chains, cfg.ChainFees)
	dao.RemoveTokenMaps(cfg.RemoveTokenMaps)

	if len(cfg.RemoveTokens) > 0 && len(cfg.RemoveTokens) == len(cfg.TokenBasics) {
		spdao, ok := dao.(*swapdao.SwapDao)
		if ok {
			for idx, asset := range cfg.TokenBasics {
				if asset.Standard != models.TokenTypeErc721 {
					continue
				}
				oldName := cfg.RemoveTokens[idx]
				newName := asset.Name
				spdao.UpdateNFTProfileTokenName(oldName, newName)
			}
		}
	}
}
