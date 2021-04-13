package chainsdk

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

type TestConfig struct {
	SDK         *EthereumSdk
	EthUrl      string
	WrapAddress common.Address
	Asset       common.Address
	Owner       common.Address
}

const (
	C_ETH_DEV  = "eth_dev"
	C_BSC_DEV  = "bsc_dev"
	C_ETH_TEST = "eth_test"
	C_BSC_TEST = "bsc_test"
)

var (
	testmode = C_BSC_DEV
	ctx      *TestConfig

	cs = map[string]*TestConfig{
		C_ETH_DEV: &TestConfig{
			EthUrl:      "http://127.0.0.1:8545",
			WrapAddress: common.HexToAddress("0xE7Db150e4095Cbb35914b5dC980906C77B5990d2"),
			Asset:       common.HexToAddress("0x03d84da9432F7Cb5364A8b99286f97c59f738001"),
			Owner:       common.HexToAddress("0x5fb03eb21303d39967a1a119b32dd744a0fa8986"),
		},
		C_BSC_DEV: &TestConfig{
			EthUrl:      "http://127.0.0.1:8546",
			WrapAddress: common.HexToAddress("0x8F967507Ae66ad78c12478E10cA07c9104eb24A7"),
			Asset:       common.HexToAddress("0x03d84da9432F7Cb5364A8b99286f97c59f738001"),
			Owner:       common.HexToAddress("0x5fb03eb21303d39967a1a119b32dd744a0fa8986"),
		},
		C_ETH_TEST: &TestConfig{
			EthUrl:      "https://ropsten.infura.io/v3/19e799349b424211b5758903de1c47ea",
			WrapAddress: common.HexToAddress("0xbaBaAF5CF7f63437755aAAFE7a4106463c5cD540"),
			Asset:       common.HexToAddress("0xa85c9FC8F2c9060d674E0CA97F703a0A30619305"),
			Owner:       common.HexToAddress("0x5fb03eb21303d39967a1a119b32dd744a0fa8986"),
		},
		C_BSC_TEST: &TestConfig{
			EthUrl:      "https://data-seed-prebsc-2-s2.binance.org:8545",
			WrapAddress: common.HexToAddress("0x2E830E0cf3dc8643B497F88C07c8A72EFE24B11f"),
			Asset:       common.HexToAddress("0x455B51D882571E244d03668f1a458ca74E70d196"),
			Owner:       common.HexToAddress("0x5fb03eb21303d39967a1a119b32dd744a0fa8986"),
		},
	}
)

func TestMain(m *testing.M) {
	ctx = cs[testmode]
	ethSdk, err := NewEthereumSdk(ctx.EthUrl)
	if err != nil {
		panic(err)
	}
	ctx.SDK = ethSdk
	m.Run()
}
