package healthmonitor

import (
	"context"
	"github.com/beego/beego/v2/core/logs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"poly-bridge/go_abi/eccm_abi"
	"testing"
)

func TestEthNodeMonitor(t *testing.T) {
	//config := conf.NewConfig("../conf/config_mainnet.json")
	config := conf.NewConfig("../prod.json")
	type args struct {
		config *conf.Config
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{config: config},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EthNodeMonitor(tt.args.config)
		})
	}
}

func EthNodeMonitor(config *conf.Config) {
	logs.Info("EthNodeMonitor")
	chainId := basedef.MATIC_CROSSCHAIN_ID
	var ccmContractAddr string
	for _, listenConfig := range config.ChainListenConfig {
		if listenConfig.ChainId == chainId {
			ccmContractAddr = listenConfig.CCMContract
			break
		}
	}

	for _, chainNodeConfig := range config.ChainNodes {
		if chainNodeConfig.ChainId == chainId {
			for _, node := range chainNodeConfig.Nodes {
				sdk, err := chainsdk.NewEthereumSdk(node.Url)
				if err != nil || sdk == nil || sdk.GetClient() == nil {
					logs.Info("node: %s,NewEthereumSdk error: %s", node.Url, err)
					continue
				}
				height, err := sdk.GetCurrentBlockHeight()
				if err != nil || height == 0 || height == math.MaxUint64 {
					logs.Error("node: %s, get current block height err: %s, ", sdk.GetUrl(), err)
					continue
				}
				height -= 1
				//height = 13881338

				logs.Info("node: %s, height: %d", node.Url, height)

				eccmContractAddress := common.HexToAddress(ccmContractAddr)
				client := sdk.GetClient()
				eccmContract, err := eccm_abi.NewEthCrossChainManager(eccmContractAddress, client)
				if err != nil {
					logs.Error("node: %s, NewEthCrossChainManager error: %s", sdk.GetUrl(), err)
					continue
				}
				opt := &bind.FilterOpts{
					Start:   height,
					End:     &height,
					Context: context.Background(),
				}
				// get ethereum lock events from given block
				_, err = eccmContract.FilterCrossChainEvent(opt, nil)
				if err != nil {
					logs.Error("node: %s, FilterCrossChainEvent error: %s", sdk.GetUrl(), err)
					continue
				}
				// ethereum unlock events from given block
				_, err = eccmContract.FilterVerifyHeaderAndExecuteTxEvent(opt)
				if err != nil {
					logs.Error("node: %s, FilterVerifyHeaderAndExecuteTxEvent error: %s", sdk.GetUrl(), err)
					continue
				}
			}
		}
	}
}
