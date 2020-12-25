package test

import (
	"fmt"
	"os"
	"poly-swap/chainlisten"
	"poly-swap/chainlisten/neolisten"
	"poly-swap/conf"
	"poly-swap/dao/stake_dao"
	"testing"
)

func TestNeoListen(t *testing.T) {
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
	neoListen := neolisten.NewNeoChainListen(config.ChainListenConfig.NeoChainListenConfig)
	chainListen := chainlisten.NewChainListen(neoListen, dao)
	chainListen.ListenChain()
}
