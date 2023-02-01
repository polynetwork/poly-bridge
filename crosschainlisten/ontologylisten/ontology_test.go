package ontologylisten

import (
	"fmt"
	"os"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"testing"
)

func Test_parse(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current directory: %s\n", dir)
	config := conf.NewConfig("../../config_testnet_local_zion.json")
	if config == nil {
		panic("read config failed!")
	}
	ListenConfig := config.GetChainListenConfig(basedef.ONT_CROSSCHAIN_ID)
	if ListenConfig == nil {
		panic("config is not valid")
	}
	ont := NewOntologyChainListen(ListenConfig)
	fmt.Println(ont.parseOntolofyMethod("6c6f636b"))
	fmt.Println(basedef.HexStringReverse("b6326b756ff2f2820d4cea745c202aa286126cbb"))
}
