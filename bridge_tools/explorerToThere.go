package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"poly-bridge/conf"
	"poly-bridge/models"
)

type ChainTokenBind struct {
	HashSrc  string
	HashDest string
}

func startExploerToThere(expcfg *conf.ExpConfig, dbcfg *conf.DBConfig) {
	Logger := logger.Default
	if expcfg.Debug == true {
		Logger = Logger.LogMode(logger.Info)
	}
	exp, err := gorm.Open(mysql.Open(expcfg.User+":"+expcfg.Password+"@tcp("+expcfg.URL+")/"+
		expcfg.Scheme+"?charset=utf8"), &gorm.Config{Logger: Logger})
	if err != nil {
		panic(err)
	}
	if dbcfg.Debug == true {
		Logger = Logger.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(dbcfg.User+":"+dbcfg.Password+"@tcp("+dbcfg.URL+")/"+
		dbcfg.Scheme+"?charset=utf8"), &gorm.Config{Logger: Logger})
	if err != nil {
		panic(err)
	}

	tokenBasics := make([]*models.TokenBasic, 0)
	err = db.Preload("Tokens").
		Find(&tokenBasics).Error
	if err != nil {
		panic(err)
	}
	chainTokenBinds := make([]*ChainTokenBind, 0)
	err = exp.Raw(`select * from chain_token_bind`).
		Scan(&chainTokenBinds).Error
	if err != nil {
		panic(fmt.Sprint("select exp ChainTokenBind err :", err))
	}
	for _, basic := range tokenBasics {
		for _, token := range basic.Tokens {
			for _, tokenBind := range chainTokenBinds {
				if token.Hash == tokenBind.HashSrc {
					basic.ChainId = token.ChainId
					break
				}
			}
		}
	}
	err = db.Save(tokenBasics).Error
	if err != nil {
		panic(fmt.Sprint("Save tokenBasics err :", err))
	}

}
