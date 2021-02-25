package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"poly-bridge/conf"
	"poly-bridge/crosschaindao"
	"poly-bridge/models"
)

func readFile(fileName string) []byte {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("File %s close error %s", fileName, err)
		}
	}()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return data
}

func startAddTransactions(cfg *conf.Config, path string) {
	dao := crosschaindao.NewCrossChainDao(cfg.Server, cfg.DBConfig)
	if dao == nil {
		panic("server is invalid")
	}

	wrapperTransactions := make([]*models.WrapperTransaction, 0)
	{
		wrapperTransactionsData := readFile(path + "/wrapper_transactions.json")
		//wrapperTransactionsData := []byte(`[{"Hash":"336cd94f1ec80280c684606b8c9358f1ad0e9e7e7ce69f0da35c21a66fa0c729","User":"ad79c606bd4ef330ac45df9d2ace4e7e7c6db13f","SrcChainId":2,"BlockHeight":9329385,"Time":1608885420,"DstChainId":4,"FeeTokenHash":"0000000000000000000000000000000000000000","FeeToken":null,"FeeAmount":1000000000000000000000000000000,"Status":1}]`)
		if len(wrapperTransactionsData) > 0 {
			err := json.Unmarshal(wrapperTransactionsData, &wrapperTransactions)
			if err != nil {
				panic(err)
			}
		} else {
			wrapperTransactions = nil
		}
	}
	srcTransactions := make([]*models.SrcTransaction, 0)
	{
		srcTransactionsData := readFile(path + "/src_transactions.json")
		//srcTransactionsData := []byte(`[{"Hash":"336cd94f1ec80280c684606b8c9358f1ad0e9e7e7ce69f0da35c21a66fa0c729","ChainId":2,"State":1,"Time":1608885420,"Fee":11370800000000000000000000,"Height":9329385,"User":"ad79c606bd4ef330ac45df9d2ace4e7e7c6db13f","DstChainId":4,"Contract":"d8ae73e06552e270340b63a8bcabf9277a1aac99","Key":"0000000000000000000000000000000000000000000000000000000000000abe","Param":"200000000000000000000000000000000000000000000000000000000000000abe20e9ef3fe2112e936772ea145dad804d2a33786fe6953ba56f294de9fab4406b0614d8ae73e06552e270340b63a8bcabf9277a1aac99040000000000000014961a23e727ea1beb5f2e2863ec427b7a99cc6f0c06756e6c6f636b4a14bf9c0fd26055ff19245c8080df06d97ae32db3d7146e43f9988f2771f1a2b140cb3faad424767d39fc0000c9ed85be3f01000000000000000000000000000000000000000000000000","SrcTransfer":{"Hash":"336cd94f1ec80280c684606b8c9358f1ad0e9e7e7ce69f0da35c21a66fa0c729","ChainId":2,"Time":1608885420,"Asset":"0000000000000000000000000000000000000000","From":"8bc7e7304120b88d111431f6a4853589d10e8132","To":"d8ae73e06552e270340b63a8bcabf9277a1aac99","Amount":9000000000000000000000000000000,"DstChainId":4,"DstAsset":"bf9c0fd26055ff19245c8080df06d97ae32db3d7","DstUser":"ARpuQar5CPtxEoqfcg1fxGWnwDdp7w3jj8"}}]`)
		if len(srcTransactionsData) > 0 {
			err := json.Unmarshal(srcTransactionsData, &srcTransactions)
			if err != nil {
				panic(err)
			}
		} else {
			srcTransactions = nil
		}
	}

	polyTransactions := make([]*models.PolyTransaction, 0)
	{
		polyTransactionsData := readFile(path + "/poly_transactions.json")
		//polyTransactionsData := []byte(`[{"Hash":"d2e8e325265ed314d9f538c2cb3f8e0a71ca2adad8b31db98278a4af6aecc1df","ChainId":0,"State":1,"Time":1609143919,"Fee":0,"Height":1641497,"SrcChainId":2,"SrcHash":"0000000000000000000000000000000000000000000000000000000000000abe","DstChainId":3,"Key":"","SrcTransaction":null}]`)
		if len(polyTransactionsData) > 0 {
			err := json.Unmarshal(polyTransactionsData, &polyTransactions)
			if err != nil {
				panic(err)
			}
		} else {
			polyTransactions = nil
		}
	}


	dstTransactions := make([]*models.DstTransaction, 0)
	{
		dstTransactionsData := readFile(path + "/dst_transactions.json")
		//dstTransactionsData := []byte(`[{"Hash":"a1d5eb3a3bf5f90438aae3b3092e8a79d5c60d46c146013015b3957f9077b399","ChainId":2,"State":1,"Time":1614143082,"Fee":237670000000000,"Height":9723051,"SrcChainId":79,"Contract":"d8ae73e06552e270340b63a8bcabf9277a1aac99","PolyHash":"8df512cc3742086c2951e8537dd9bce6c7153e16467f462d30ebf227cb6df368","DstTransfer":{"TxHash":"a1d5eb3a3bf5f90438aae3b3092e8a79d5c60d46c146013015b3957f9077b399","ChainId":2,"Time":1614143082,"Asset":"7f8f2a4ae259b3655539a58636f35dad0a1d9989","From":"d8ae73e06552e270340b63a8bcabf9277a1aac99","To":"735adb7d290ee47a8e677e1c5f3b7e922c3366c4","Amount":499991626530000000000}}]`)
		if len(dstTransactionsData) > 0 {
			err := json.Unmarshal(dstTransactionsData, &dstTransactions)
			if err != nil {
				panic(err)
			}
		} else {
			dstTransactions = nil
		}
	}

	chain := new(models.Chain)
	{
		chainData := readFile(path + "/chains.json")
		//chainData := []byte(`{"ChainId":2,"Name":"Ethereum","Height":9329385}`)
		if len(chainData) > 0 {
			err := json.Unmarshal(chainData, chain)
			if err != nil {
				panic(err)
			}
		} else {
			chain = nil
		}
	}
	err := dao.UpdateEvents(chain, wrapperTransactions, srcTransactions, polyTransactions, dstTransactions)
	if err != nil {
		panic(err)
	}
}
