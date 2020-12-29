package test

import (
	"encoding/json"
	"fmt"
	"os"
	"poly-swap/chainlisten"
	"poly-swap/chainlisten/ethereumlisten"
	"poly-swap/conf"
	"poly-swap/dao/stake_dao"
	"testing"
)

func TestBscListen(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current directory: %s\n", dir)
	config := conf.NewConfig("./../../../conf/config_testnet.json")
	if config == nil {
		panic("read config failed!")
	}
	dao := stake_dao.NewStakeDao()
	ethereumListenConfig := new(conf.EthereumChainListenConfig)
	chainJson, err := json.Marshal(config.ChainListenConfig.BscChainListenConfig)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(chainJson, ethereumListenConfig)
	if err != nil {
		panic(err)
	}
	ethereumListen := ethereumlisten.NewEthereumChainListen(ethereumListenConfig)
	chainListen := chainlisten.NewChainListen(ethereumListen, dao)
	chainListen.ListenChain()
}
