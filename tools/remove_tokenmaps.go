package main

import (
	"encoding/json"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"poly-bridge/conf"
	"poly-bridge/models"
	"strings"
)

func startRemoveTokenMaps(cfg *conf.Config, path string) {
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
	tokenMaps := make([]*models.TokenMap, 0)
	{
		tokenMapsData := readFile(path + "/remove_tokenmaps.json")
		if len(tokenMapsData) > 0 {
			err := json.Unmarshal(tokenMapsData, &tokenMaps)
			if err != nil {
				panic(err)
			}
		} else {
			tokenMaps = nil
		}
	}
	for _, tokenMap := range tokenMaps {
		db.Where("src_chain_id = ? and src_token_hash = ? and dst_chain_id = ? and dst_token_hash = ?",
			tokenMap.SrcChainId, strings.ToLower(tokenMap.SrcTokenHash), tokenMap.DstChainId, strings.ToLower(tokenMap.DstTokenHash)).Delete(&models.TokenMap{})
	}
}
