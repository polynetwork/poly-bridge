package test

import (
	"encoding/json"
	"fmt"
	"os"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/crosschaindao"
	"poly-bridge/crosschainlisten"
	"poly-bridge/crosschainlisten/aptoslisten"
	"testing"
)

func TestAptosChainListen_HandleNewEvent(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current directory: %s\n", dir)
	config := conf.NewConfig("../../../config_testnet_local_zion.json")
	if config == nil {
		panic("read config failed!")
	}
	dao := crosschaindao.NewCrossChainDao(basedef.SERVER_POLY_BRIDGE, false, config.DBConfig)
	ListenConfig := config.GetChainListenConfig(basedef.APTOS_CROSSCHAIN_ID)
	if ListenConfig == nil {
		panic("config is not valid")
	}
	chainHandle := crosschainlisten.NewChainHandle(ListenConfig)
	a, b, c, d, z, x, err := chainHandle.(*aptoslisten.AptosChainListen).HandleEvent(dao, 0, 0, 0)
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
