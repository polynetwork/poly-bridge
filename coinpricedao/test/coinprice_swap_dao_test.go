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
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"poly-bridge/basedef"
	"poly-bridge/coinpricedao"
	"poly-bridge/conf"
	"poly-bridge/models"
	"testing"
)

func TestSavePrice_SwapDao(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current directory: %s\n", dir)
	config := conf.NewConfig("./../../conf/config_testnet.json")
	if config == nil {
		panic("read config failed!")
	}
	db := coinpricedao.NewCoinPriceDao(basedef.SERVER_POLY_SWAP, config.DBConfig)
	if db == nil {
		panic("dao is invalid")
	}
	tokenBasics := make([]*models.TokenBasic, 0)
	tokenBasicsJson := []byte(`[{"Name":"Ethereum","Precision":0,"AvgPrice":73080095858,"AvgInd":1,"Time":0,"PriceMarkets":[{"TokenBasicName":"Ethereum","MarketName":"binance","Name":"ETHUSDT","Price":73080000000,"Ind":1,"Time":1609308634,"TokenBasic":null},{"TokenBasicName":"Ethereum","MarketName":"coinmarketcap","Name":"Ethereum","Price":73080191717,"Ind":1,"Time":1609308634,"TokenBasic":null}],"Tokens":[{"Hash":"0000000000000000000000000000000000000000","ChainId":2,"Name":"Ethereum","Precision":18,"TokenBasicName":"Ethereum","TokenBasic":null,"TokenMaps":null},{"Hash":"0000000000000000000000000000000000000005","ChainId":4,"Name":"Ethereum","Precision":18,"TokenBasicName":"Ethereum","TokenBasic":null,"TokenMaps":null}]},{"Name":"Neo","Precision":0,"AvgPrice":1485333999,"AvgInd":1,"Time":0,"PriceMarkets":[{"TokenBasicName":"Neo","MarketName":"binance","Name":"NEOUSDT","Price":1485000000,"Ind":1,"Time":1609308634,"TokenBasic":null},{"TokenBasicName":"Neo","MarketName":"coinmarketcap","Name":"Neo","Price":1485667998,"Ind":1,"Time":1609308634,"TokenBasic":null}],"Tokens":[{"Hash":"0000000000000000000000000000000000000001","ChainId":2,"Name":"Neo","Precision":9,"TokenBasicName":"Neo","TokenBasic":null,"TokenMaps":null},{"Hash":"0000000000000000000000000000000000000006","ChainId":4,"Name":"Neo","Precision":9,"TokenBasicName":"Neo","TokenBasic":null,"TokenMaps":null}]}]`)
	err = json.Unmarshal(tokenBasicsJson, &tokenBasics)
	if err != nil {
		panic(err)
	}
	err = db.SavePrices(tokenBasics)
	if err != nil {
		panic(err)
	}
}

func TestQueryTokens_SwapDao(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current directory: %s\n", dir)
	config := conf.NewConfig("./../../conf/config_testnet.json")
	if config == nil {
		panic("read config failed!")
	}
	dbCfg := config.DBConfig
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	tokenBasics := make([]*models.TokenBasic, 0)
	db.Debug().Model(&models.TokenBasic{}).Preload("PriceMarkets").Preload("Tokens").Find(&tokenBasics)
	json, _ := json.Marshal(tokenBasics)
	fmt.Printf("src Transaction: %s\n", json)
}
