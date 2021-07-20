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
	tokenBinds = append(tokenBinds, ChainTokenBind{6, "8PAY"})
	tokenBinds = append(tokenBinds, ChainTokenBind{2, "BET"})
	tokenBinds = append(tokenBinds, ChainTokenBind{2, "BKC"})
	tokenBinds = append(tokenBinds, ChainTokenBind{6, "BNB"})
	tokenBinds = append(tokenBinds, ChainTokenBind{2, "CART"})
	tokenBinds = append(tokenBinds, ChainTokenBind{2, "DETO"})
	tokenBinds = append(tokenBinds, ChainTokenBind{7, "DOG"})
	tokenBinds = append(tokenBinds, ChainTokenBind{2, "DOI"})
	tokenBinds = append(tokenBinds, ChainTokenBind{2, "ECELL"})
	tokenBinds = append(tokenBinds, ChainTokenBind{2, "ETH"})
	tokenBinds = append(tokenBinds, ChainTokenBind{2, "FILE"})
	tokenBinds = append(tokenBinds, ChainTokenBind{4, "FLM"})
	tokenBinds = append(tokenBinds, ChainTokenBind{2, "FLUX"})
	tokenBinds = append(tokenBinds, ChainTokenBind{2, "FREL"})
	tokenBinds = append(tokenBinds, ChainTokenBind{7, "HAI"})
	tokenBinds = append(tokenBinds, ChainTokenBind{7, "HDT"})
	tokenBinds = append(tokenBinds, ChainTokenBind{7, "HKR"})
	tokenBinds = append(tokenBinds, ChainTokenBind{6, "ISM"})
	tokenBinds = append(tokenBinds, ChainTokenBind{6, "KEL"})
	tokenBinds = append(tokenBinds, ChainTokenBind{2, "KISHU"})
	tokenBinds = append(tokenBinds, ChainTokenBind{6, "LKT"})
	tokenBinds = append(tokenBinds, ChainTokenBind{2, "MBTC"})
	tokenBinds = append(tokenBinds, ChainTokenBind{6, "MIX"})
	tokenBinds = append(tokenBinds, ChainTokenBind{6, "NEST"})
	tokenBinds = append(tokenBinds, ChainTokenBind{3, "ONG"})
	tokenBinds = append(tokenBinds, ChainTokenBind{6, "PKR"})
	tokenBinds = append(tokenBinds, ChainTokenBind{7, "PLF"})
	tokenBinds = append(tokenBinds, ChainTokenBind{7, "PLUT"})
	tokenBinds = append(tokenBinds, ChainTokenBind{2, "ROSN"})
	tokenBinds = append(tokenBinds, ChainTokenBind{7, "SHARE"})
	tokenBinds = append(tokenBinds, ChainTokenBind{2, "SHIB"})
	tokenBinds = append(tokenBinds, ChainTokenBind{6, "SLD"})
	tokenBinds = append(tokenBinds, ChainTokenBind{6, "SOFA"})
	tokenBinds = append(tokenBinds, ChainTokenBind{6, "SPAY"})
	tokenBinds = append(tokenBinds, ChainTokenBind{2, "SPHRI"})
	tokenBinds = append(tokenBinds, ChainTokenBind{2, "sUSD"})
	tokenBinds = append(tokenBinds, ChainTokenBind{4, "SWTH"})
	tokenBinds = append(tokenBinds, ChainTokenBind{6, "Tribe"})
	tokenBinds = append(tokenBinds, ChainTokenBind{2, "UNI"})
	tokenBinds = append(tokenBinds, ChainTokenBind{2, "USDC"})
	tokenBinds = append(tokenBinds, ChainTokenBind{2, "WBTC"})
	tokenBinds = append(tokenBinds, ChainTokenBind{3, "WING"})
	tokenBinds = append(tokenBinds, ChainTokenBind{6, "XTF"})
	tokenBinds = append(tokenBinds, ChainTokenBind{7, "XTM"})
	fmt.Println(tokenBinds)
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
