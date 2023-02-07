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
	ethListenConfig := config.GetChainListenConfig(basedef.ZION_CROSSCHAIN_ID)
	if ethListenConfig == nil {
		panic("config is not valid")
	}
	chainHandle := crosschainlisten.NewChainHandle(ethListenConfig)
	_, _, polyTransactions, _, _, polyDetails, _, _, err := chainHandle.HandleNewBatchBlock(582378, 582384)
	if err != nil {
		fmt.Println("handle err", err)
	}
	a, err := json.MarshalIndent(polyTransactions, "", "	")
	fmt.Println("tx", string(a))
	if err != nil {
		fmt.Println("marshal err", err)
	}
	a, _ = json.MarshalIndent(polyDetails, "", "	")

	fmt.Println("detail", string(a))
}

func TestEthereumChainListen_HandleNewBatchBlock2(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current directory: %s\n", dir)
	config := conf.NewConfig("./config.json")
	if config == nil {
		panic("read config failed!")
	}
	ethListenConfig := config.GetChainListenConfig(basedef.BSC_CROSSCHAIN_ID)
	if ethListenConfig == nil {
		panic("config is not valid")
	}
	chainHandle := crosschainlisten.NewChainHandle(ethListenConfig)
	a, b, c, d, z, x, _, _, err := chainHandle.HandleNewBatchBlock(26852504, 26852504)
	if err != nil {
		fmt.Println("err", err)
	}
	e, _ := json.MarshalIndent(a, "", "	")
	fmt.Println("a", string(e))
	e, _ = json.MarshalIndent(b, "", "	")
	fmt.Println("b", string(e))
	e, _ = json.MarshalIndent(c, "", "	")
	fmt.Println("c", string(e))
	e, _ = json.MarshalIndent(d, "", "	")
	fmt.Println("d", string(e))
	e, _ = json.MarshalIndent(z, "", "	")
	fmt.Println("z", string(e))
	e, _ = json.MarshalIndent(x, "", "	")
	fmt.Println("x", string(e))
}

func TestEthereumChainListen_HandleNewBlock3(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current directory: %s\n", dir)
	config := conf.NewConfig("../../../config_testnet_local_zion.json")
	if config == nil {
		panic("read config failed!")
	}
	ListenConfig := config.GetChainListenConfig(basedef.NEO3_CROSSCHAIN_ID)
	if ListenConfig == nil {
		panic("config is not valid")
	}
	chainHandle := crosschainlisten.NewChainHandle(ListenConfig)
	a, b, c, d, z, x, _, _, err := chainHandle.HandleNewBlock(147736)
	if err != nil {
		fmt.Println("err", err)
	}
	e, _ := json.MarshalIndent(a, "", "	")
	fmt.Println("a", string(e))
	e, _ = json.MarshalIndent(b, "", "	")
	fmt.Println("b", string(e))
	e, _ = json.MarshalIndent(c, "", "	")
	fmt.Println("c", string(e))
	e, _ = json.MarshalIndent(d, "", "	")
	fmt.Println("d", string(e))
	e, _ = json.MarshalIndent(z, "", "	")
	fmt.Println("z", string(e))
	e, _ = json.MarshalIndent(x, "", "	")
	fmt.Println("x", string(e))
}
