package test

import (
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/crosschainlisten/ethereumlisten"
	"testing"
)

func TestNewEthereumChainListen(t *testing.T) {
	config := conf.NewConfig("./config.json")
	if config == nil {
		panic("read config failed!")
	}
	ethListenConfig := config.GetChainListenConfig(basedef.ZION_CROSSCHAIN_ID)
	if ethListenConfig == nil {
		panic("config is not valid")
	}
	ethereumlisten.NewEthereumChainListen(ethListenConfig)
}
