package chainsdk

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

type TestConfig struct {
	SDK          *EthereumSdk
	EthUrl       string
	QueryAddress common.Address
	Asset        common.Address
	Owner        common.Address
}

const (
	C_ETH_DEV           = "eth_dev"
	C_BSC_DEV           = "bsc_dev"
	C_ETH_TEST          = "eth_test"
	C_BSC_TEST          = "bsc_test"
	C_ETH_SEASCAPE_TEST = "eth_seascape_test"
	C_BSC_SEASCAPE_TEST = "bsc_seascape_test"
)

var (
	testmode = C_ETH_SEASCAPE_TEST
	ctx      *TestConfig

	cs = map[string]*TestConfig{
		C_ETH_DEV: &TestConfig{
			EthUrl:       "http://127.0.0.1:8545",
			QueryAddress: common.HexToAddress("0xBb0e8D6CFd87C6A07e312f1F31fd1F1cC9949F2C"),
			Asset:        common.HexToAddress("0x03d84da9432F7Cb5364A8b99286f97c59f738001"),
			Owner:        common.HexToAddress("0x5fb03eb21303d39967a1a119b32dd744a0fa8986"),
		},
		C_BSC_DEV: &TestConfig{
			EthUrl:       "http://127.0.0.1:8546",
			QueryAddress: common.HexToAddress("0x8F967507Ae66ad78c12478E10cA07c9104eb24A7"),
			Asset:        common.HexToAddress("0x03d84da9432F7Cb5364A8b99286f97c59f738001"),
			Owner:        common.HexToAddress("0x5fb03eb21303d39967a1a119b32dd744a0fa8986"),
		},
		C_ETH_TEST: &TestConfig{
			EthUrl:       "https://goerli.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161",
			QueryAddress: common.HexToAddress("0x0eaAe6F5D8Be3177D9286825c3878Dd4AAAFD2e6"),
			Asset:        common.HexToAddress("0x52d3d38ceca694c3dd163a85ef8a08e0edc7b07b"),
			Owner:        common.HexToAddress("0x1eeDc8e7A4c708acF64106205F79CA4CDe11Ce3A"),
		},
		C_BSC_TEST: &TestConfig{
			EthUrl:       "http://43.128.231.211:8575", //"https://data-seed-prebsc-2-s2.binance.org:8545",
			QueryAddress: common.HexToAddress("0x274DD6F0bA27C167821fD340cb3a3B2Ab3b5827D"),
			Asset:        common.HexToAddress("0x455B51D882571E244d03668f1a458ca74E70d196"),
			Owner:        common.HexToAddress("0x5fb03eb21303d39967a1a119b32dd744a0fa8986"),
		},
		C_ETH_SEASCAPE_TEST: &TestConfig{
			EthUrl:       "https://goerli.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161",
			QueryAddress: common.HexToAddress("0x0eaAe6F5D8Be3177D9286825c3878Dd4AAAFD2e6"),
			Asset:        common.HexToAddress("0x52d3d38ceca694c3dd163a85ef8a08e0edc7b07b"),
			Owner:        common.HexToAddress("0x1eeDc8e7A4c708acF64106205F79CA4CDe11Ce3A"),
		},
		C_BSC_SEASCAPE_TEST: &TestConfig{
			EthUrl:       "https://data-seed-prebsc-2-s2.binance.org:8545",
			QueryAddress: common.HexToAddress("0x274DD6F0bA27C167821fD340cb3a3B2Ab3b5827D"),
			Asset:        common.HexToAddress("0x66638f4970c2ae63773946906922c07a583b6069"),
			Owner:        common.HexToAddress("0xc12E333cdD2843c7719aFfca036cDe023579F192"),
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
