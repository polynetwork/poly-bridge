package test

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"poly-swap/conf"
	"poly-swap/dao/swap_dao"
	"poly-swap/models"
	"testing"
)

func TestCrossChain_SwapDao(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current directory: %s\n", dir)
	config := conf.NewConfig("./../../conf/config_testnet.json")
	if config == nil {
		panic("read config failed!")
	}
	dao := swap_dao.NewSwapDao(config.DBConfig)
	wrapperTransactions := make([]*models.WrapperTransaction, 0)
	wrapperTransactionsData := []byte(`[{"Hash":"336cd94f1ec80280c684606b8c9358f1ad0e9e7e7ce69f0da35c21a66fa0c729","User":"ad79c606bd4ef330ac45df9d2ace4e7e7c6db13f","SrcChainId":2,"BlockHeight":9329385,"Time":1608885420,"DstChainId":4,"FeeTokenHash":"0000000000000000000000000000000000000000","FeeToken":null,"FeeAmount":1000000000000000000000000000000,"Status":0}]`)
	err = json.Unmarshal(wrapperTransactionsData, &wrapperTransactions)
	if err != nil {
		panic(err)
	}
	srcTransactions := make([]*models.SrcTransaction, 0)
	srcTransactionsData := []byte(`[{"Hash":"336cd94f1ec80280c684606b8c9358f1ad0e9e7e7ce69f0da35c21a66fa0c729","ChainId":2,"State":1,"Time":1608885420,"Fee":11370800000000000000000000,"Height":9329385,"User":"ad79c606bd4ef330ac45df9d2ace4e7e7c6db13f","DstChainId":4,"Contract":"d8ae73e06552e270340b63a8bcabf9277a1aac99","Key":"0000000000000000000000000000000000000000000000000000000000000abe","Param":"200000000000000000000000000000000000000000000000000000000000000abe20e9ef3fe2112e936772ea145dad804d2a33786fe6953ba56f294de9fab4406b0614d8ae73e06552e270340b63a8bcabf9277a1aac99040000000000000014961a23e727ea1beb5f2e2863ec427b7a99cc6f0c06756e6c6f636b4a14bf9c0fd26055ff19245c8080df06d97ae32db3d7146e43f9988f2771f1a2b140cb3faad424767d39fc0000c9ed85be3f01000000000000000000000000000000000000000000000000","SrcTransfer":{"Hash":"336cd94f1ec80280c684606b8c9358f1ad0e9e7e7ce69f0da35c21a66fa0c729","ChainId":2,"Time":1608885420,"Asset":"0000000000000000000000000000000000000000","From":"8bc7e7304120b88d111431f6a4853589d10e8132","To":"d8ae73e06552e270340b63a8bcabf9277a1aac99","Amount":9000000000000000000000000000000,"DstChainId":4,"DstAsset":"bf9c0fd26055ff19245c8080df06d97ae32db3d7","DstUser":"ARpuQar5CPtxEoqfcg1fxGWnwDdp7w3jj8"}}]`)
	err = json.Unmarshal(srcTransactionsData, &srcTransactions)
	if err != nil {
		panic(err)
	}
	polyTransactions := make([]*models.PolyTransaction, 0)
	polyTransactionsData := []byte(`[{"Hash":"d2e8e325265ed314d9f538c2cb3f8e0a71ca2adad8b31db98278a4af6aecc1df","ChainId":0,"State":1,"Time":1609143919,"Fee":0,"Height":1641497,"SrcChainId":2,"SrcHash":"0000000000000000000000000000000000000000000000000000000000000abe","DstChainId":3,"Key":"","SrcTransaction":null}]`)
	err = json.Unmarshal(polyTransactionsData, &polyTransactions)
	if err != nil {
		panic(err)
	}
	chain := new(models.Chain)
	chainData := []byte(`{"ChainId":2,"Name":"Ethereum","Height":9329385}`)
	err = json.Unmarshal(chainData, chain)
	if err != nil {
		panic(err)
	}
	err = dao.UpdateEvents(chain, wrapperTransactions, srcTransactions, polyTransactions, nil)
	if err != nil {
		panic(err)
	}
}

func TestQuerySrcTransaction_SwapDao(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current directory: %s\n", dir)
	config := conf.NewConfig("./../../conf/config_testnet.json")
	if config == nil {
		panic("read config failed!")
	}
	dbCfg := config.DBConfig
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	srcTransaction := new(models.SrcTransaction)
	db.Model(&models.SrcTransaction{}).Preload("SrcTransfer").First(srcTransaction)
	json, _ := json.Marshal(srcTransaction)
	fmt.Printf("src Transaction: %s\n", json)
}

func TestQuerySrcTransfer_SwapDao(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current directory: %s\n", dir)
	config := conf.NewConfig("./../../conf/config_testnet.json")
	if config == nil {
		panic("read config failed!")
	}
	dbCfg := config.DBConfig
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	srcTransaction := new(models.SrcTransaction)
	db.Model(&models.SrcTransaction{}).Preload("SrcTransfer", "chain_id = ?", 2).First(srcTransaction)
	json, _ := json.Marshal(srcTransaction)
	fmt.Printf("src Transaction: %s\n", json)
}

func TestQueryPolySrcRelation_SwapDao(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current directory: %s\n", dir)
	config := conf.NewConfig("./../../conf/config_testnet.json")
	if config == nil {
		panic("read config failed!")
	}
	dbCfg := config.DBConfig
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	polySrcRelations := make([]*models.PolySrcRelation, 0)
	db.Debug().Table("poly_transactions").Where("left(poly_transactions.src_hash, 8) = ?", "00000000").Select("poly_transactions.hash as poly_hash, src_transactions.hash as src_hash").Joins("left join src_transactions on poly_transactions.src_hash = src_transactions.key").Preload("SrcTransaction").Preload("PolyTransaction").Find(&polySrcRelations)
	json, _ := json.Marshal(polySrcRelations)
	fmt.Printf("src Transaction: %s\n", json)
}