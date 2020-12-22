package main

import (
	"gorm.io/gorm"
	"poly-swap/conf"
	"poly-swap/models"
	"gorm.io/driver/mysql"
)

func startDeploy(cfg *conf.DeployConfig) {
	dbCfg := cfg.DBConfig
	db, err := gorm.Open(mysql.Open(dbCfg.User + ":" + dbCfg.Password + "@tcp(" + dbCfg.URL + ")/" +
		dbCfg.Scheme + "?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.Chain{}, &models.Transaction{}, &models.TokenBasic{}, &models.ChainFee{}, &models.Token{}, &models.TokenMap{})
	if err != nil {
		panic(err)
	}
	//
	db.Save(cfg.Chains)
	db.Save(cfg.TokenBasics)
	db.Save(cfg.ChainFees)
	db.Save(cfg.TokenMaps)
}
/*
func updateChains() []*models.Chain {
	// todo, add your chain
	chains := []*models.Chain {
		{
			ChainId: conf.ETHEREUM_CROSSCHAIN_ID,
			Name:    conf.ETHEREUM_CHAIN_NAME,
		},
		{
			ChainId: conf.NEO_CROSSCHAIN_ID,
			Name:    conf.NEO_CHAIN_NAME,
		},
		{
			ChainId: conf.BSC_CROSSCHAIN_ID,
			Name:    conf.BSC_CHAIN_NAME,
		},
	}
	return chains
}

func updateChainGases() []*models.ChainGas {
	// todo, add your chain gas
	chainGases := []*models.ChainGas {
		{
			ChainId:  conf.ETHEREUM_CROSSCHAIN_ID,
		},
		{
			ChainId:  conf.NEO_CROSSCHAIN_ID,
		},
		{
			ChainId:  conf.BSC_CROSSCHAIN_ID,
		},
	}
	return chainGases
}

func updateTokens() []*models.TokenBasic {
	// todo, add your token
	tokens := []*models.TokenBasic {
		{
			Name:     "Bitcoin",
			Tokens:   []*models.Token {
				{
					ChainId:      conf.ETHEREUM_CROSSCHAIN_ID,
					Hash:         "",
					Name:         "",
				},
				{
					ChainId:      conf.NEO_CROSSCHAIN_ID,
					Hash:         "",
					Name:         "",
				},
			},
		},
		{
			Name:     "Ethereum",
			Tokens:   []*models.Token {
				{
					ChainId:      conf.ETHEREUM_CROSSCHAIN_ID,
					Hash:         "",
					Name:         "",
				},
				{
					ChainId:      conf.NEO_CROSSCHAIN_ID,
					Hash:         "",
					Name:         "",
				},
			},
		},
	}
	return tokens
}

func updateTokenMap(tokens []*models.Token) []*models.TokenMap {
	hash2Token := make(map[string]*models.Token, 0)
	for _, token := range tokens {
		hash2Token[token.Hash] = token
	}
	// todo, add your token map
	crosschainTokens := map[string]string {
		"aaa" : "bbb",
	}
	tokenMaps := make([]*models.TokenMap, 0)
	for k, v := range crosschainTokens {
		srcTokenId, ok := hash2Token[k]
		if !ok {
			continue
		}
		dstTokenId, ok := hash2Token[v]
		if !ok {
			continue
		}
		tokenMaps = append(tokenMaps, &models.TokenMap {
			SrcTokenId: srcTokenId.Id,
			DstTokenId: dstTokenId.Id,
		})
	}
	return tokenMaps
}
*/
