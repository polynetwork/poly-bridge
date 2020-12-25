package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"poly-swap/conf"
	"poly-swap/models"
)

func startDeploy(cfg *conf.DeployConfig) {
	dbCfg := cfg.DBConfig
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.Chain{}, &models.WrapperTransaction{}, &models.TokenBasic{}, &models.ChainFee{}, &models.Token{}, &models.TokenMap{},
		&models.SrcTransaction{}, &models.SrcTransfer{}, &models.PolyTransaction{}, &models.DstTransaction{}, &models.DstTransfer{})
	if err != nil {
		panic(err)
	}
	//
	db.Save(cfg.Chains)
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
