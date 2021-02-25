/*
 * Copyright (C) 2020 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package test

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"poly-bridge/crosschaindao"
	"poly-bridge/crosschaindao/explorerdao"
	"poly-bridge/models"
	"testing"
)

func TestCrossChain_ExplorerDao(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current directory: %s\n", dir)
	config := conf.NewConfig("./../../conf/config_testnet.json")
	if config == nil {
		panic("read config failed!")
	}
	dao := explorerdao.NewExplorerDao(config.DBConfig)
	wrapperTransactions := make([]*models.WrapperTransaction, 0)
	wrapperTransactionsData := []byte(`[{"Hash":"336cd94f1ec80280c684606b8c9358f1ad0e9e7e7ce69f0da35c21a66fa0c729","User":"ad79c606bd4ef330ac45df9d2ace4e7e7c6db13f","SrcChainId":2,"BlockHeight":9329385,"Time":1608885420,"DstChainId":4,"FeeTokenHash":"0000000000000000000000000000000000000000","FeeToken":null,"FeeAmount":1000000000000000000000000000000,"Status":0}]`)
	err = json.Unmarshal(wrapperTransactionsData, &wrapperTransactions)
	if err != nil {
		panic(err)
	}
	srcTransactions := make([]*models.SrcTransaction, 0)
	srcTransactionsData := []byte(`[{"Hash":"336cd94f1ec80280c684606b8c9358f1ad0e9e7e7ce69f0da35c21a66fa0c729","ChainId":2,"State":1,"Time":1608885420,"Fee":11370800000000,"Height":9329385,"User":"ad79c606bd4ef330ac45df9d2ace4e7e7c6db13f","DstChainId":4,"Contract":"d8ae73e06552e270340b63a8bcabf9277a1aac99","Key":"0000000000000000000000000000000000000000000000000000000000000abe","Param":"200000000000000000000000000000000000000000000000000000000000000abe20e9ef3fe2112e936772ea145dad804d2a33786fe6953ba56f294de9fab4406b0614d8ae73e06552e270340b63a8bcabf9277a1aac99040000000000000014961a23e727ea1beb5f2e2863ec427b7a99cc6f0c06756e6c6f636b4a14bf9c0fd26055ff19245c8080df06d97ae32db3d7146e43f9988f2771f1a2b140cb3faad424767d39fc0000c9ed85be3f01000000000000000000000000000000000000000000000000","SrcTransfer":{"Hash":"336cd94f1ec80280c684606b8c9358f1ad0e9e7e7ce69f0da35c21a66fa0c729","ChainId":2,"Time":1608885420,"Asset":"0000000000000000000000000000000000000000","From":"8bc7e7304120b88d111431f6a4853589d10e8132","To":"d8ae73e06552e270340b63a8bcabf9277a1aac99","Amount":9000000000000000000000000000000,"DstChainId":4,"DstAsset":"bf9c0fd26055ff19245c8080df06d97ae32db3d7","DstUser":"ARpuQar5CPtxEoqfcg1fxGWnwDdp7w3jj8"}}]`)
	err = json.Unmarshal(srcTransactionsData, &srcTransactions)
	if err != nil {
		panic(err)
	}
	chain := new(models.Chain)
	chainData := []byte(`{"ChainId":2,"Name":"Ethereum","Height":9329385}`)
	err = json.Unmarshal(chainData, chain)
	if err != nil {
		panic(err)
	}
	err = dao.UpdateEvents(chain, wrapperTransactions, srcTransactions, nil, nil)
	if err != nil {
		panic(err)
	}
}

func TestCrossChainSrc_ExplorerDao(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current directory: %s\n", dir)
	config := conf.NewConfig("./../../conf/config_testnet.json")
	if config == nil {
		panic("read config failed!")
	}
	dao := crosschaindao.NewCrossChainDao(basedef.SERVER_EXPLORER, config.DBConfig)
	if dao == nil {
		panic("server is not valid")
	}
	srcTransactions := make([]*models.SrcTransaction, 0)
	srcTransactionsData := []byte(`[{"Hash":"74b469fb26f1229db5cf516b6ac2b9722eb68573e3e1f5687849cab9feb3b10c","ChainId":7,"State":1,"Time":1611028940,"Fee":2729461364730000,"Height":1379405,"User":"6b0f370aa682cd43066c134e4b3e2e0922832408","DstChainId":5,"Contract":"4a76e52600c6285029c8f7c52183cf86282ca5b8","Key":"0000000000000000000000000000000000000000000000000000000000000018","Param":"20000000000000000000000000000000000000000000000000000000000000001820b082c060ef6652d93b5702a5c70eefcd15e5a18b4f27c764322635a8bbbb26b7144a76e52600c6285029c8f7c52183cf86282ca5b80500000000000000144f5f702b3f459f222d371052940bb9ce2d86d2ed06756e6c6f636b4a146cf6e87ab27a492647277686d29bc4a451ac01bb14ec796ad7d3f70013cba8b2499b7e36bdc74abbc1c43d000000000000000000000000000000000000000000000000000000000000","SrcTransfer":{"TxHash":"74b469fb26f1229db5cf516b6ac2b9722eb68573e3e1f5687849cab9feb3b10c","ChainId":7,"Time":1611028940,"Asset":"faddf0cfb08f92779560db57be6b2c7303aad266","From":"6b0f370aa682cd43066c134e4b3e2e0922832408","To":"4a76e52600c6285029c8f7c52183cf86282ca5b8","Amount":15812,"DstChainId":5,"DstAsset":"6cf6e87ab27a492647277686d29bc4a451ac01bb","DstUser":"ec796ad7d3f70013cba8b2499b7e36bdc74abbc1"}}]`)
	err = json.Unmarshal(srcTransactionsData, &srcTransactions)
	if err != nil {
		panic(err)
	}
	err = dao.UpdateEvents(nil, nil, srcTransactions, nil, nil)
	if err != nil {
		panic(err)
	}
}

func TestQuerySrcTransaction_ExplorerDao(t *testing.T) {
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

func TestQuerySrcTransfer_ExplorerDao(t *testing.T) {
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

func TestQueryPolySrcRelation_ExplorerDao(t *testing.T) {
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
	polySrcRelations := make([]*explorerdao.PolySrcRelation, 0)
	db.Debug().Table("mchain_tx").Where("left(mchain_tx.ftxhash, 8) = ? and fchain = ?", "00000000", basedef.ETHEREUM_CROSSCHAIN_ID).Select("mchain_tx.txhash as poly_hash, fchain_tx.txhash as src_hash").Joins("left join fchain_tx on mchain_tx.ftxhash = fchain_tx.xkey and mchain_tx.fchain = fchain_tx.chain_id").Preload("SrcTransaction").Preload("PolyTransaction").Find(&polySrcRelations)
	json, _ := json.Marshal(polySrcRelations)
	fmt.Printf("src Transaction: %s\n", json)
}

func TestUpdateTokenInfo_ExplorerDao(t *testing.T) {
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
	var ethSdk *chainsdk.EthereumSdkPro
	var bscSdk *chainsdk.EthereumSdkPro
	var hecoSdk *chainsdk.EthereumSdkPro
	var neoSdk *chainsdk.NeoSdkPro
	for _, listenConfig := range config.ChainListenConfig {
		urls := listenConfig.GetNodesUrl()
		if listenConfig.ChainId == basedef.ETHEREUM_CROSSCHAIN_ID {
			ethSdk = chainsdk.NewEthereumSdkPro(urls, listenConfig.ListenSlot, listenConfig.ChainId)
		} else if listenConfig.ChainId == basedef.BSC_CROSSCHAIN_ID {
			bscSdk = chainsdk.NewEthereumSdkPro(urls, listenConfig.ListenSlot, listenConfig.ChainId)
		} else if listenConfig.ChainId == basedef.HECO_CROSSCHAIN_ID {
			hecoSdk = chainsdk.NewEthereumSdkPro(urls, listenConfig.ListenSlot, listenConfig.ChainId)
		} else if listenConfig.ChainId == basedef.NEO_CROSSCHAIN_ID {
			neoSdk = chainsdk.NewNeoSdkPro(urls, listenConfig.ListenSlot, listenConfig.ChainId)
		}
	}
	{
		tokens := make([]*explorerdao.Token, 0)
		newTokens := make([]*explorerdao.Token, 0)
		db.Debug().Table("fchain_transfer").Distinct("chain_id as id", "asset as hash").Find(&tokens)
		for _, token := range tokens {
			fmt.Printf("chain: %d, token hash: %s\n", token.Id, token.Hash)
			if token.Hash[0:4] == "0000" || token.Hash[28:32] == "0000" {
				continue
			}
			if token.Id == basedef.ETHEREUM_CROSSCHAIN_ID {
				hash, name, decimal, symbol, err := ethSdk.Erc20Info(token.Hash)
				if err != nil {
					panic(err)
				}
				token.Name = name
				token.Type = "ERC20"
				token.Desc = symbol
				token.Precision = fmt.Sprintf("%d", basedef.Int64FromFigure(int(decimal)))
				newTokens = append(newTokens, token)
				fmt.Printf("erc20: %s, name: %s, decimal: %d, symbol: %s\n",
					hash, name, decimal, symbol)
			} else if token.Id == basedef.BSC_CROSSCHAIN_ID {
				hash, name, decimal, symbol, err := bscSdk.Erc20Info(token.Hash)
				if err != nil {
					panic(err)
				}
				token.Name = name
				token.Type = "ERC20"
				token.Desc = symbol
				token.Precision = fmt.Sprintf("%d", basedef.Int64FromFigure(int(decimal)))
				fmt.Printf("bsc: %s, name: %s, decimal: %d, symbol: %s\n",
					hash, name, decimal, symbol)
				newTokens = append(newTokens, token)
			} else if token.Id == basedef.HECO_CROSSCHAIN_ID {
				hash, name, decimal, symbol, err := hecoSdk.Erc20Info(token.Hash)
				if err != nil {
					panic(err)
				}
				token.Name = name
				token.Type = "ERC20"
				token.Desc = symbol
				token.Precision = fmt.Sprintf("%d", basedef.Int64FromFigure(int(decimal)))
				fmt.Printf("heco: %s, name: %s, decimal: %d, symbol: %s\n",
					hash, name, decimal, symbol)
				newTokens = append(newTokens, token)
			} else if token.Id == basedef.NEO_CROSSCHAIN_ID {
				hash, name, decimal, err := neoSdk.Nep5Info(token.Hash)
				if err != nil {
					panic(err)
				}
				token.Name = name
				token.Type = "NEP5"
				token.Desc = name
				token.Precision = fmt.Sprintf("%d", basedef.Int64FromFigure(int(decimal)))
				fmt.Printf("nep5: %s, decimal: %d, name: %s\n", hash, decimal, name)
				newTokens = append(newTokens, token)
			}
		}
		data, _ := json.Marshal(newTokens)
		fmt.Printf("%s\n", data)
		db.Debug().Save(newTokens)
	}
}
