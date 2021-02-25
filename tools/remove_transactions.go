package main

import (
	"encoding/json"
	"poly-bridge/conf"
	"poly-bridge/crosschaindao"
)

func startRemoveTransactions(cfg *conf.Config, path string) {
	dao := crosschaindao.NewCrossChainDao(cfg.Server, cfg.DBConfig)
	if dao == nil {
		panic("server is invalid")
	}
	removeTransactions := make([]string, 0)
	{
		removeTransactionsData := readFile(path + "/remove_transactions.json")
		if len(removeTransactionsData) > 0 {
			err := json.Unmarshal(removeTransactionsData, &removeTransactions)
			if err != nil {
				panic(err)
			}
		} else {
			removeTransactions = nil
		}
	}
	err := dao.RemoveEvents(removeTransactions, removeTransactions, removeTransactions)
	if err != nil {
		panic(err)
	}
}
