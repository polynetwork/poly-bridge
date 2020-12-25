package test

import (
	"fmt"
	"os"
	"poly-swap/chainlisten"
	"poly-swap/chainlisten/ethereumlisten"
	"poly-swap/conf"
	"poly-swap/dao/stake_dao"
	"testing"
)

func TestEthereumListen(t *testing.T) {
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
	ethereumListen := ethereumlisten.NewEthereumChainListen(config.ChainListenConfig.EthereumChainListenConfig)
	chainListen := chainlisten.NewChainListen(ethereumListen, dao)
	chainListen.ListenChain()
}
