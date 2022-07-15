package healthmonitor

import (
	"context"
	"github.com/beego/beego/v2/core/logs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joeqian10/neo-gogogo/tx"
	"github.com/joeqian10/neo-gogogo/wallet"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/keys"
	tx2 "github.com/joeqian10/neo3-gogogo/tx"
	wallet2 "github.com/joeqian10/neo3-gogogo/wallet"
	ontologygosdk "github.com/ontio/ontology-go-sdk"
	common2 "github.com/ontio/ontology/common"
	"math"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"poly-bridge/go_abi/eccm_abi"
	"testing"
)

func TestEthNodeMonitor(t *testing.T) {
	//config := conf.NewConfig("../conf/config_mainnet.json")
	config := conf.NewConfig("../../prod.json")
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
	chainId := basedef.ZKSYNC_CROSSCHAIN_ID
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

				timestamp, err := sdk.GetBlockTimeByNumber(chainId, height)
				if err != nil {
					logs.Error("node: %s, GetHeaderTimeByBlockNumber err: %s, ", sdk.GetUrl(), err)
					continue
				}
				logs.Info("node: %s, GetHeaderTimeByBlockNumber: %d, time:%d", node.Url, height, timestamp)

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

func TestEthRelayerMonitor(t *testing.T) {
	config := conf.NewConfig("../../prod.json")
	relayerConfig := conf.NewRelayerConfig("../../relayer_prod.json")
	type args struct {
		config      *conf.Config
		relayConfig *conf.RelayerConfig
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{config: config, relayConfig: relayerConfig},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EthRelayerMonitor(tt.args.config, tt.args.relayConfig)
		})
	}
}

func EthRelayerMonitor(config *conf.Config, relayConfig *conf.RelayerConfig) {
	chainId := basedef.ETHEREUM_CROSSCHAIN_ID
	var precision float64 = 1000000000000000000
	var relayAccount *conf.RelayAccountConfig
	for _, cfg := range relayConfig.RelayAccountConfig {
		if cfg.ChainId == chainId {
			relayAccount = cfg
		}
	}

	for _, chainNodeConfig := range config.ChainNodes {
		if chainNodeConfig.ChainId == chainId {
			for _, node := range chainNodeConfig.Nodes {
				balanceSuccessMap := make(map[string]float64, 0)
				balanceFailedMap := make(map[string]string, 0)
				sdk, err := chainsdk.NewEthereumSdk(node.Url)
				if err != nil || sdk == nil || sdk.GetClient() == nil {
					logs.Info("node: %s,NewEthereumSdk error: %s", node.Url, err)
					continue
				}
				for _, address := range relayAccount.Address {
					if _, ok := balanceSuccessMap[address]; ok {
						continue
					}
					balance, err := sdk.GetNativeBalance(common.HexToAddress(address))
					if err == nil {
						balanceSuccessMap[address] = float64(balance.Uint64()) / precision
						delete(balanceFailedMap, address)
					} else {
						balanceFailedMap[address] = err.Error()
					}
				}
				logs.Info("balanceSuccessMap=%+v", balanceSuccessMap)
				logs.Info("balanceFailedMap=%+v", balanceFailedMap)
			}
		}
	}
}

func TestNeoRelayerMonitor(t *testing.T) {
	config := conf.NewConfig("../../prod.json")
	relayerConfig := conf.NewRelayerConfig("../../relayer_prod.json")
	type args struct {
		config      *conf.Config
		relayConfig *conf.RelayerConfig
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{config: config, relayConfig: relayerConfig},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NeoRelayerMonitor(tt.args.config, tt.args.relayConfig)
		})
	}
}

func NeoRelayerMonitor(config *conf.Config, relayConfig *conf.RelayerConfig) {
	chainId := basedef.NEO_CROSSCHAIN_ID
	var relayAccount *conf.RelayAccountConfig
	for _, cfg := range relayConfig.RelayAccountConfig {
		if cfg.ChainId == chainId {
			relayAccount = cfg
		}
	}

	for _, chainNodeConfig := range config.ChainNodes {
		if chainNodeConfig.ChainId == chainId {
			for _, node := range chainNodeConfig.Nodes {
				balanceSuccessMap := make(map[string]float64, 0)
				balanceFailedMap := make(map[string]string, 0)
				sdk := chainsdk.NewNeoSdk(node.Url)
				if sdk.GetClient() == nil {
					logs.Info("node: %s,NewNeoSdk error: %s", node.Url)
					continue
				}
				for _, address := range relayAccount.Address {
					if _, ok := balanceSuccessMap[address]; ok {
						continue
					}

					txBuilder := &tx.TransactionBuilder{
						EndPoint: sdk.GetUrl(),
						Client:   sdk.GetClient(),
					}
					walletHelper := wallet.NewWalletHelper(txBuilder, nil)
					_, gasBalance, err := walletHelper.GetBalance(address)

					if err == nil {
						balanceSuccessMap[address] = gasBalance
						delete(balanceFailedMap, address)
					} else {
						balanceFailedMap[address] = err.Error()
					}
				}
				logs.Info("balanceSuccessMap=%+v", balanceSuccessMap)
				logs.Info("balanceFailedMap=%+v", balanceFailedMap)
			}
		}
	}
}

func TestOntRelayerMonitor(t *testing.T) {
	config := conf.NewConfig("../../prod.json")
	relayerConfig := conf.NewRelayerConfig("../../relayer_prod.json")
	type args struct {
		config      *conf.Config
		relayConfig *conf.RelayerConfig
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{config: config, relayConfig: relayerConfig},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OntRelayerMonitor(tt.args.config, tt.args.relayConfig)
		})
	}
}

func OntRelayerMonitor(config *conf.Config, relayConfig *conf.RelayerConfig) {
	var precision float64 = 1000000000
	chainId := basedef.ONT_CROSSCHAIN_ID
	var relayAccount *conf.RelayAccountConfig
	for _, cfg := range relayConfig.RelayAccountConfig {
		if cfg.ChainId == chainId {
			relayAccount = cfg
		}
	}

	for _, chainNodeConfig := range config.ChainNodes {
		if chainNodeConfig.ChainId == chainId {
			for _, node := range chainNodeConfig.Nodes {
				balanceSuccessMap := make(map[string]float64, 0)
				balanceFailedMap := make(map[string]string, 0)
				sdk := ontologygosdk.NewOntologySdk()
				sdk.NewRpcClient().SetAddress(node.Url)
				for _, address := range relayAccount.Address {
					if _, ok := balanceSuccessMap[address]; ok {
						continue
					}

					account, err := common2.AddressFromBase58(address)
					logs.Info("1")
					if err != nil {
						logs.Info("2")
						account, err = common2.AddressFromHexString(address)

					}
					balance, err := sdk.Native.Ong.BalanceOf(account)
					if err == nil {
						balanceSuccessMap[address] = float64(balance) / precision
						delete(balanceFailedMap, address)
					} else {
						balanceFailedMap[address] = err.Error()
					}
				}
				logs.Info("balanceSuccessMap=%+v", balanceSuccessMap)
				logs.Info("balanceFailedMap=%+v", balanceFailedMap)
			}
		}
	}
}

func TestNeo3RelayerMonitor(t *testing.T) {
	config := conf.NewConfig("../../prod.json")
	relayerConfig := conf.NewRelayerConfig("../../relayer_prod.json")
	type args struct {
		config      *conf.Config
		relayConfig *conf.RelayerConfig
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{config: config, relayConfig: relayerConfig},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Neo3RelayerMonitor(tt.args.config, tt.args.relayConfig)
		})
	}
}

func Neo3RelayerMonitor(config *conf.Config, relayConfig *conf.RelayerConfig) {
	chainId := basedef.NEO3_CROSSCHAIN_ID
	var precision float64 = 100000000
	var relayAccount *conf.RelayAccountConfig
	for _, cfg := range relayConfig.RelayAccountConfig {
		if cfg.ChainId == chainId {
			relayAccount = cfg
		}
	}

	for _, chainNodeConfig := range config.ChainNodes {
		if chainNodeConfig.ChainId == chainId {
			for _, node := range chainNodeConfig.Nodes {
				balanceSuccessMap := make(map[string]float64, 0)
				balanceFailedMap := make(map[string]string, 0)
				sdk := chainsdk.NewNeo3Sdk(node.Url)
				if sdk.GetClient() == nil {
					logs.Info("node: %s,NewNeoSdk error: %s", node.Url)
					continue
				}
				for _, account := range relayAccount.Neo3Account {
					if _, ok := balanceSuccessMap[account.Address]; ok {
						continue
					}
					nep2Key := "6PYURwh7xSUBA6yEbBsnYCGESd7UNHi3iWHsDweCxwzEH8xPBSphb1z7ss"
					pwd := "poly"
					keypair, err := keys.NewKeyPairFromNEP2(nep2Key, pwd, helper.DefaultAddressVersion, keys.N, keys.R, keys.P)
					if err != nil {
						logs.Error("NewKeyPairFromNEP2 err:%s", err)
						continue
					}
					logs.Info("PrivateKey=%+v", keypair.PrivateKey)
					wh, err := wallet2.NewWalletHelperFromPrivateKey(sdk.GetClient(), keypair.PrivateKey)
					if err != nil {
						balanceFailedMap[account.Address] = err.Error()
					}
					accountAndBalances, err := wh.GetAccountAndBalance(tx2.GasToken)
					total := big.NewInt(0)
					if err == nil {
						for _, balance := range accountAndBalances {
							total = total.Add(total, balance.Value)
						}
						balanceSuccessMap[account.Address] = float64(total.Uint64()) / precision
						delete(balanceFailedMap, account.Address)
					} else {
						balanceFailedMap[account.Address] = err.Error()
					}
				}
				logs.Info("balanceSuccessMap=%+v", balanceSuccessMap)
				logs.Info("balanceFailedMap=%+v", balanceFailedMap)
			}
		}
	}
}
