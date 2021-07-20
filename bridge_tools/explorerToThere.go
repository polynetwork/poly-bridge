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
	ChainId uint64
	Name    string
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
	err = db.
		Find(&tokenBasics).Error
	if err != nil {
		panic(err)
	}
	chainTokenBinds := make([]*ChainTokenBind, 0)
	err = exp.Raw(`SELECT b.id as chain_id ,b.xname as name from chain_token_bind a join chain_token b on a.hash_src=b.hash Where a.hash_src=a.hash_dest and  b.hash != '0000000000000000000000000000000000000000'`).
		Scan(&chainTokenBinds).Error
	for _, x := range chainTokenBinds {
		fmt.Println(*x)
	}
	if err != nil {
		panic(fmt.Sprint("select exp ChainTokenBind err :", err))
	}
	mapTokenBinds := make(map[ChainTokenBind]int)
	for _, tokenBind := range chainTokenBinds {
		mapTokenBinds[*tokenBind]++
	}
	tokenBinds := make([]ChainTokenBind, 0)
	for k, v := range mapTokenBinds {
		if v == 1 {
			tokenBinds = append(tokenBinds, k)
		}
	}
	for _, basic := range tokenBasics {
		for _, tokenBind := range tokenBinds {
			if tokenBind.Name == basic.Name {
				err := db.Debug().Model(&models.TokenBasic{}).
					Where("name = ?", basic.Name).
					Update("chain_id", tokenBind.ChainId).
					Error
				if err != nil {
					panic(fmt.Sprint("Update tokenBasic err :", err))
				}
				fmt.Println(tokenBind)
				break
			}
		}
	}
}
