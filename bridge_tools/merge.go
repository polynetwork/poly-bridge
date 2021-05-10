package main

import (
	"encoding/json"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"poly-bridge/conf"
	"poly-bridge/crosschaindao/explorerdao"
	"poly-bridge/models"
	"time"
)

func merge() {
	csdbCfg := new(conf.DBConfig)
	csdbCfg.User = "root"
	csdbCfg.Debug = true
	csdbCfg.Scheme = "cross_chain_explorer"
	csdbCfg.URL = "127.0.0.1:3306"
	csdbCfg.Password = "Onchain@2019"
	Logger := logger.Default
	if csdbCfg.Debug == true {
		Logger = Logger.LogMode(logger.Info)
	}
	csdb, err := gorm.Open(mysql.Open(csdbCfg.User+":"+csdbCfg.Password+"@tcp("+csdbCfg.URL+")/"+
		csdbCfg.Scheme+"?charset=utf8"), &gorm.Config{Logger: Logger})
	if err != nil {
		panic(err)
	}
	newswapdbCfg := new(conf.DBConfig)
	newswapdbCfg.User = "root"
	newswapdbCfg.Debug = true
	newswapdbCfg.Scheme = "polybridge_v2"
	newswapdbCfg.URL = "10.203.0.11:3306"
	newswapdbCfg.Password = "PAIGWICQFKDNzdL5aTw0pIPrBeoYinXu4A=="
	newswapdb, err := gorm.Open(mysql.Open(newswapdbCfg.User+":"+newswapdbCfg.Password+"@tcp("+newswapdbCfg.URL+")/"+
		newswapdbCfg.Scheme+"?charset=utf8"), &gorm.Config{Logger: Logger})
	if err != nil {
		panic(err)
	}
	//selectNum := 10
	insertNum := 1000
	{
		srcTransactions := make([]*explorerdao.SrcTransaction, 0)
		//csdb.Preload("SrcTransfer").Order("tt asc").Limit(selectNum).Find(&srcTransactions)
		csdb.Preload("SrcTransfer").Order("tt asc").Find(&srcTransactions)
		for i := 0;i < len(srcTransactions); {
			j := i + insertNum
			if j > len(srcTransactions) {
				j = len(srcTransactions)
			}
			xx := srcTransactions[i:j]
			srcTransactionsJson, err := json.Marshal(xx)
			if err != nil {
				panic(err)
			}
			newSrcTransactions := make([]*models.SrcTransaction, 0)
			err = json.Unmarshal(srcTransactionsJson, &newSrcTransactions)
			if err != nil {
				panic(err)
			}
			newswapdb.Save(newSrcTransactions)
			time.Sleep(time.Second * 1)
			i = j
		}
	}
	{
		polyTransactions := make([]*explorerdao.PolyTransaction, 0)
		//csdb.Order("tt asc").Limit(selectNum).Find(&polyTransactions)
		csdb.Order("tt asc").Find(&polyTransactions)
		for i := 0;i < len(polyTransactions); {
			j := i + insertNum
			if j > len(polyTransactions) {
				j = len(polyTransactions)
			}
			xx := polyTransactions[i:j]
			polyTransactionsJson, err := json.Marshal(xx)
			if err != nil {
				panic(err)
			}
			newPolyTransactions := make([]*models.PolyTransaction, 0)
			err = json.Unmarshal(polyTransactionsJson, &newPolyTransactions)
			if err != nil {
				panic(err)
			}
			newswapdb.Save(newPolyTransactions)
			time.Sleep(time.Second * 1)
			i = j
		}
	}
	{
		dstTransactions := make([]*explorerdao.DstTransaction, 0)
		//csdb.Preload("DstTransfer").Order("tt asc").Limit(selectNum).Find(&dstTransactions)
		csdb.Preload("DstTransfer").Order("tt asc").Find(&dstTransactions)
		for i := 0;i < len(dstTransactions); {
			j := i + insertNum
			if j > len(dstTransactions) {
				j = len(dstTransactions)
			}
			xx := dstTransactions[i:j]
			dstTransactionsJson, err := json.Marshal(xx)
			if err != nil {
				panic(err)
			}
			newDstTransactions := make([]*models.DstTransaction, 0)
			err = json.Unmarshal(dstTransactionsJson, &newDstTransactions)
			if err != nil {
				panic(err)
			}
			newswapdb.Save(newDstTransactions)
			time.Sleep(time.Second * 1)
			i = j
		}
	}
}
