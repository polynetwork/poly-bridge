package test

import (
	"encoding/json"
	"fmt"
	"os"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/crosschainlisten"
	"testing"
)

func TestEthereumChainListen_HandleNewBatchBlock(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current directory: %s\n", dir)
	config := conf.NewConfig("./config.json")
	if config == nil {
		panic("read config failed!")
	}
	ethListenConfig := config.GetChainListenConfig(basedef.ZIONMAIN_CROSSCHAIN_ID)
	if ethListenConfig == nil {
		panic("config is not valid")
	}
	chainHandle := crosschainlisten.NewChainHandle(ethListenConfig)
	_, _, polyTransactions, _, _, _, err := chainHandle.HandleNewBatchBlock(29492, 29492)
	if err != nil {
		fmt.Println("err", err)
	}
	a, _ := json.MarshalIndent(polyTransactions, "", "	")
	fmt.Println(string(a))
}
