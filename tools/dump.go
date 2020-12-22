package main

import (
	"fmt"
	"gorm.io/gorm"
	"poly-swap/conf"
	"poly-swap/models"
	"gorm.io/driver/mysql"
)

func dumpStatus(dbCfg *conf.DBConfig) {
	db, err := gorm.Open(mysql.Open(dbCfg.User + ":" + dbCfg.Password + "@tcp(" + dbCfg.URL + ")/" +
		dbCfg.Scheme + "?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	{
		chains := make([]*models.Chain, 0)
		db.Find(&chains)
		fmt.Printf("chain info:\nchainid\t\t\t\tname\t\t\t\theight\t\t\t\t\n")
		for _, chain := range chains {
			fmt.Printf("%d\t\t\t\t%s\t\t\t\t%d\t\t\t\t\n", chain.ChainId, chain.Name, chain.Height)
		}
	}
	{
		tokens := make([]*models.Token, 0)
		db.Find(&tokens)
		fmt.Printf("token info:\nChainId\t\t\t\tHash\t\t\t\tName\t\t\t\tTokenBasicName\t\t\t\t\n")
		for _, token := range tokens {
			fmt.Printf("%d\t\t\t\t%s\t\t\t\t%s\t\t\t\t%s\t\t\t\t\n",
				token.ChainId, token.Hash, token.Name, token.TokenBasicName)
		}
	}
	{
		tokenBasics := make([]*models.TokenBasic, 0)
		db.Find(&tokenBasics)
		fmt.Printf("token basic info:\nName\t\t\t\tCMCPrice\t\t\t\tCMCInd\t\t\t\tBinPrice\t\t\t\tBinInd\t\t\t\tAvgPrice\t\t\t\tAvgInd\t\t\t\tTime\t\t\t\t\n")
		for _, tokenBasic := range tokenBasics {
			fmt.Printf("%s\t\t\t\t%d\t\t\t\t%d\t\t\t\t%d\t\t\t\t%d\t\t\t\t%d\t\t\t\t%d\t\t\t\t%d\t\t\t\t\n",
				tokenBasic.Name, tokenBasic.CmcPrice, tokenBasic.CmcInd, tokenBasic.BinPrice, tokenBasic.BinInd, tokenBasic.AvgPrice, tokenBasic.AvgInd,
				tokenBasic.Time)
		}
	}
	{
		TokenMaps := make([]*models.TokenMap, 0)
		db.Find(&TokenMaps)
		fmt.Printf("token map info:\nSrcTokenHash\t\t\t\tDstTokenHash\t\t\t\t\n")
		for _, TokenMap := range TokenMaps {
			fmt.Printf("%s\t\t\t\t%s\t\t\t\t\n", TokenMap.SrcTokenHash, TokenMap.DstTokenHash)
		}
	}
	{
		chainFees := make([]*models.ChainFee, 0)
		db.Find(&chainFees)
		fmt.Printf("chain fee info:\nchainid\t\t\t\tTokenBasicName\t\t\t\tmaxfee\t\t\t\tminfee\t\t\t\tproxyfee\t\t\t\t\n")
		for _, chainFee := range chainFees {
			fmt.Printf("%d\t\t\t\t%s\t\t\t\t%d\t\t\t\t%d\t\t\t\t%d\t\t\t\t\n",
				chainFee.ChainId, chainFee.TokenBasicName, chainFee.MaxFee, chainFee.MinFee, chainFee.ProxyFee)
		}
	}
}
