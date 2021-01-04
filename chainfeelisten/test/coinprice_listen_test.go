package test

import (
	"fmt"
	"os"
	"poly-swap/chainfeedao"
	"poly-swap/chainfeelisten"
	"poly-swap/conf"
	"testing"
)

func TestListenFee(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current directory: %s\n", dir)
	config := conf.NewConfig("./../../conf/config_testnet.json")
	if config == nil {
		panic("read config failed!")
	}
	dao := chainfeedao.NewChainFeeDao(config.Server, config.DBConfig)
	if dao == nil {
		panic("server is not valid")
	}
	feeListenCfgs := config.FeeListenConfig
	chainFees := make([]chainfeelisten.ChainFee, 0)
	for _, cfg := range feeListenCfgs {
		chainFee := chainfeelisten.NewChainFee(cfg)
		chainFees = append(chainFees, chainFee)
	}
	feeListen := chainfeelisten.NewFeeListen(config.FeeUpdateSlot, chainFees, dao)
	feeListen.ListenFee()
}
