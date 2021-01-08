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
	"poly-bridge/conf"
	"poly-bridge/crosschaineffect"
	"poly-bridge/models"
	"testing"
)

func TestSwapEffect_Clear(t *testing.T) {
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
	db.Where("1 = 1").Delete(models.WrapperTransaction{})
	db.Where("1 = 1").Delete(models.SrcTransfer{})
	db.Where("1 = 1").Delete(models.SrcTransaction{})
	db.Where("1 = 1").Delete(models.PolyTransaction{})
	db.Where("1 = 1").Delete(models.DstTransfer{})
	db.Where("1 = 1").Delete(models.DstTransaction{})
}

func TestSwapEffect_WrapperTransactions(t *testing.T) {
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
	wrapperTransactions := make([]*models.WrapperTransaction, 0)
	db.Model(&models.WrapperTransaction{}).Find(&wrapperTransactions)
	json, _ := json.Marshal(wrapperTransactions)
	fmt.Printf("src Transaction: %s\n", json)
}

func TestSwapEffect_Run(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current directory: %s\n", dir)
	config := conf.NewConfig("./../../conf/config_testnet.json")
	if config == nil {
		panic("read config failed!")
	}
	effect := crosschaineffect.NewEffect(config.EventEffectConfig, config.DBConfig)
	if effect == nil {
		panic("monitor is not valid")
	}
	crossChainEffect := crosschaineffect.NewCrossChainEffect(effect)
	crossChainEffect.Check()
}
