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
	"os"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/crosschaindao"
	"poly-bridge/models"
	"testing"
)

func TestSwapEffect_AddSrcTransaction(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current directory: %s\n", dir)
	config := conf.NewConfig("./../../conf/config_testnet.json")
	if config == nil {
		panic("read config failed!")
	}
	dao := crosschaindao.NewCrossChainDao(basedef.SERVER_POLY_SWAP, config.DBConfig)
	if dao == nil {
		panic("server is not valid")
	}
	wrapperTransactions := make([]*models.WrapperTransaction, 0)
	wrapperTransactionsData := []byte(`[{"Hash":"336cd94f1ec80280c684606b8c9358f1ad0e9e7e7ce69f0da35c21a66fa0c729","User":"ad79c606bd4ef330ac45df9d2ace4e7e7c6db13f","SrcChainId":2,"BlockHeight":9329385,"Time":1608885420,"DstChainId":4,"FeeTokenHash":"0000000000000000000000000000000000000000","FeeToken":null,"FeeAmount":1000000000000000000000000000000,"Status":1}]`)
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

func TestSwapEffect_AddPolyTransaction(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current directory: %s\n", dir)
	config := conf.NewConfig("./../../conf/config_testnet.json")
	if config == nil {
		panic("read config failed!")
	}
	dao := crosschaindao.NewCrossChainDao(basedef.SERVER_POLY_SWAP, config.DBConfig)
	if dao == nil {
		panic("server is not valid")
	}
	polyTransactions := make([]*models.PolyTransaction, 0)
	polyTransactionsData := []byte(`[{"Hash":"d2e8e325265ed314d9f538c2cb3f8e0a71ca2adad8b31db98278a4af6aecc1df","ChainId":0,"State":1,"Time":1609143919,"Fee":0,"Height":1641497,"SrcChainId":2,"SrcHash":"0000000000000000000000000000000000000000000000000000000000000abe","DstChainId":3,"Key":"","SrcTransaction":null}]`)
	err = json.Unmarshal(polyTransactionsData, &polyTransactions)
	if err != nil {
		panic(err)
	}
	chain := new(models.Chain)
	chainData := []byte(`{"ChainId":0,"Name":"Poly","Height":1641497}`)
	err = json.Unmarshal(chainData, chain)
	if err != nil {
		panic(err)
	}
	err = dao.UpdateEvents(chain, nil, nil, polyTransactions, nil)
	if err != nil {
		panic(err)
	}
}

func TestSwapEffect_AddDstTransaction(t *testing.T) {
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
	dstTransactions := make([]*models.DstTransaction, 0)
	dstTransactionsData := []byte(`[{"Hash":"a1d5eb3a3bf5f90438aae3b3092e8a79d5c60d46c146013015b3957f9077b399","ChainId":2,"State":1,"Time":1614143082,"Fee":237670000000000,"Height":9723051,"SrcChainId":79,"Contract":"d8ae73e06552e270340b63a8bcabf9277a1aac99","PolyHash":"8df512cc3742086c2951e8537dd9bce6c7153e16467f462d30ebf227cb6df368","DstTransfer":{"TxHash":"a1d5eb3a3bf5f90438aae3b3092e8a79d5c60d46c146013015b3957f9077b399","ChainId":2,"Time":1614143082,"Asset":"7f8f2a4ae259b3655539a58636f35dad0a1d9989","From":"d8ae73e06552e270340b63a8bcabf9277a1aac99","To":"735adb7d290ee47a8e677e1c5f3b7e922c3366c4","Amount":499991626530000000000}}]`)
	err = json.Unmarshal(dstTransactionsData, &dstTransactions)
	if err != nil {
		panic(err)
	}
	chain := new(models.Chain)
	chainData := []byte(`{"ChainId":4,"Name":"Neo","Height":4810052}`)
	err = json.Unmarshal(chainData, chain)
	if err != nil {
		panic(err)
	}
	err = dao.UpdateEvents(chain, nil, nil, nil, dstTransactions)
	if err != nil {
		panic(err)
	}
}
