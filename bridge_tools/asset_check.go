package main

import (
	"encoding/json"
	"fmt"
	"github.com/polynetwork/poly-io-test/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"math/big"
	"poly-bridge/common"
	"poly-bridge/conf"
	"poly-bridge/models"
)

type AssetDetail struct {
	BasicName  string
	TokenAsset []*DstChainAsset
	Difference *big.Int
}
type DstChainAsset struct {
	ChainId     uint64
	TotalSupply *big.Int
	Balance     *big.Int
	flow        *big.Int
}

func startCheckAsset(dbCfg *conf.DBConfig) {
	log.Info("q-w-e-r-t start startCheckAsset")
	Logger := logger.Default
	if dbCfg.Debug == true {
		Logger = Logger.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{Logger: Logger})
	if err != nil {
		panic(err)
	}

	resAssetDetails := make([]*AssetDetail, 0)
	tokenBasics := make([]*models.TokenBasic, 0)
	res := db.
		Where("property = ?", 1).
		Preload("Tokens").
		Find(&tokenBasics)
	if res.Error != nil {
		panic(fmt.Errorf("load chainBasics faild, err: %v", res.Error))
	}
	log.Info("q-w-e-r-t start to foreach tokenBasics")
	for _, basic := range tokenBasics {
		log.Info(fmt.Sprintf("	for basicname: %v", basic.Name))
		assetDetail := new(AssetDetail)
		dstChainAssets := make([]*DstChainAsset, 0)
		totalFlow := big.NewInt(0)
		for _, token := range basic.Tokens {
			chainAsset := new(DstChainAsset)
			chainAsset.ChainId = token.ChainId
			balance, err := common.GetBalance(token.ChainId, token.Hash)
			if err != nil {
				panic(fmt.Errorf("q-w-e-r-t In CheckAsset Chain: %v,hash: %v , GetBalance faild, err: %v", token.ChainId, token.Hash, res.Error))
			}
			log.Info(fmt.Sprintf("	chainId: %v, Hash: %v, balance: %v", token.ChainId, token.Hash, balance.String()))
			chainAsset.Balance = balance
			totalSupply, err := common.GetTotalSupply(token.ChainId, token.Hash)
			if err != nil {
				panic(fmt.Errorf("q-w-e-r-t In CheckAsset Chain: %v,hash: %v , GetTotalSupply faild, err: %v", token.ChainId, token.Hash, res.Error))
			}
			log.Info(fmt.Sprintf("	chainId: %v, Hash: %v, totalSupply: %v", token.ChainId, token.Hash, totalSupply.String()))
			chainAsset.TotalSupply = totalSupply
			chainAsset.flow = new(big.Int).Sub(totalSupply, balance)
			log.Info(fmt.Sprintf("	chainId: %v, Hash: %v, flow: %v", token.ChainId, token.Hash, chainAsset.flow.String()))
			totalFlow = new(big.Int).Add(totalFlow, chainAsset.flow)
			dstChainAssets = append(dstChainAssets, chainAsset)
		}
		assetDetail.TokenAsset = dstChainAssets
		log.Info(fmt.Sprintf("	basic: %v,totalFlow: %v", basic.Name, totalFlow.String()))
		assetDetail.Difference = totalFlow
		assetDetail.BasicName = basic.Name
		resAssetDetails = append(resAssetDetails, assetDetail)
	}
	jsonCheckAsset, _ := json.Marshal(resAssetDetails)
	log.Info(fmt.Sprintf("q-w-e-r-t" + string(jsonCheckAsset)))
	fmt.Println("q-w-e-r-t" + string(jsonCheckAsset))
}
