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

package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"poly-bridge/models"
)

var (
	db   = newDB()
	sdks = make(map[uint64]*chainsdk.EthereumSdkPro)
	assets = make([]*models.Token, 0)
	wrapperAddrs = make(map[uint64]common.Address)
)

func newDB() *gorm.DB {
	user := beego.AppConfig.String("mysqluser")
	password := beego.AppConfig.String("mysqlpass")
	url := beego.AppConfig.String("mysqlurls")
	scheme := beego.AppConfig.String("mysqldb")
	mode := beego.AppConfig.String("runmode")
	Logger := logger.Default
	if mode == "dev" {
		Logger = Logger.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(user+":"+password+"@tcp("+url+")/"+scheme+"?charset=utf8"), &gorm.Config{Logger: Logger})
	if err != nil {
		panic(err)
	}

	err = db.Debug().AutoMigrate(&models.Chain{}, &models.WrapperTransaction{}, &models.ChainFee{}, &models.TokenBasic{}, &models.Token{}, &models.PriceMarket{},
		&models.TokenMap{}, &models.SrcTransaction{}, &models.SrcTransfer{}, &models.PolyTransaction{}, &models.DstTransaction{}, &models.DstTransfer{})
	if err != nil {
		panic(err)
	}

	//db.Model(&models.Token{}).
	//	Where("standard = ?", models.TokenTypeErc721).
	//	Preload("TokenBasic").
	//	Find(&assets)

	db.Where("standard = ?", models.TokenTypeErc721).
		Preload("TokenBasic").
		//Preload("TokenMaps").
		//Preload("TokenMaps.DstToken").
		Find(&assets)

	for _, v := range assets {
		logs.Info("load asset %s, chainid %d", v.TokenBasicName, v.ChainId)
	}
	return db
}

func Initialize(c *conf.Config) {
	//var err error
	//Logger := logger.Default
	//if c.DBConfig.Debug {
	//	Logger = Logger.LogMode(logger.Info)
	//}
	//link := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
	//	c.DBConfig.User,
	//	c.DBConfig.Password,
	//	c.DBConfig.URL,
	//	c.DBConfig.Scheme,
	//)
	//if db, err = gorm.Open(mysql.Open(link), &gorm.Config{Logger: Logger}); err != nil {
	//	panic(err)
	//}

	for _, v := range c.ChainListenConfig {
		pro := chainsdk.NewEthereumSdkPro(v.GetNodesUrl(), v.ListenSlot, v.ChainId)
		sdks[v.ChainId] = pro
		wrapperAddrs[v.ChainId] = common.HexToAddress(v.NFTWrapperContract)
		logs.Info("load chain id %d, contract %s", v.ChainId, v.NFTWrapperContract)
	}
}

func selectNode(chainID uint64) *chainsdk.EthereumSdkPro {
	pro, ok := sdks[chainID]
	if !ok {
		return nil
	}
	return pro
}

var emptyAddr = common.Address{}
func selectWrapper(chainID uint64) common.Address {
	addr, ok := wrapperAddrs[chainID]
	if !ok {
		return emptyAddr
	}
	return addr
}

const (
	ErrCodeRequest     int = 400
	ErrCodeNotExist    int = 404
	ErrCodeNodeInvalid int = 500
)

var errMap = map[int]string{
	ErrCodeRequest:     "request parameter is invalid!",
	ErrCodeNotExist:    "not found",
	ErrCodeNodeInvalid: "blockchain node exception",
}

func input(c *beego.Controller, req interface{}) bool {
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		code := ErrCodeRequest
		customInput(c, code, errMap[code])
		return false
	} else {
		return true
	}
}

func customInput(c *beego.Controller, code int, msg string) {
	c.Data["json"] = models.MakeErrorRsp(msg)
	c.Ctx.ResponseWriter.WriteHeader(code)
	c.ServeJSON()
}

func notExist(c *beego.Controller) {
	code := ErrCodeNotExist
	c.Data["json"] = models.MakeErrorRsp(errMap[code])
	c.Ctx.ResponseWriter.WriteHeader(code)
	c.ServeJSON()
}

func nodeInvalid(c *beego.Controller) {
	code := ErrCodeNodeInvalid
	c.Data["json"] = models.MakeErrorRsp(errMap[code])
	c.Ctx.ResponseWriter.WriteHeader(code)
	c.ServeJSON()
}

func output(c *beego.Controller, data interface{}) {
	c.Data["json"] = data
	c.ServeJSON()
}

func customOutput(c *beego.Controller, code int, msg string) {
	c.Data["json"] = models.MakeErrorRsp(msg)
	c.Ctx.ResponseWriter.WriteHeader(code)
	c.ServeJSON()
}

func getPageNo(totalNo, pageSize int) int {
	return (int(totalNo) + pageSize - 1) / pageSize
}

func findFeeToken(cid uint64, hash string) *models.Token{
	feeTokens := make([]*models.Token, 0)
	db.Model(&models.Token{}).
		Where("hash = ?", nativeHash).
		Preload("TokenBasic").
		Find(&feeTokens)

	for _, v := range feeTokens {
		if cid == v.ChainId && hash == v.Hash {
			return v
		}
	}
	return nil
}

func findAsset(cid uint64, hash string) *models.Token {
	for _, v := range assets {
		if v.ChainId == cid && hash == v.Hash {
			return v
		}
	}
	return nil
}