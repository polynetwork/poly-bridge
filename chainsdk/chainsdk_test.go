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
			EthUrl:       "https://ropsten.infura.io/v3/19e799349b424211b5758903de1c47ea",
			QueryAddress: common.HexToAddress("0xD2B67aeeA3A5e85AEe9F77E98db094d1E30A00Ed"),
			Asset:        common.HexToAddress("0xa85c9FC8F2c9060d674E0CA97F703a0A30619305"),
			Owner:        common.HexToAddress("0xf1c7203ef81fb9663babd8516ebd30d33ee84ee8"),
		},
		C_BSC_TEST: &TestConfig{
			EthUrl:       "http://43.128.231.211:8575", //"https://data-seed-prebsc-2-s2.binance.org:8545",
			QueryAddress: common.HexToAddress("0x274DD6F0bA27C167821fD340cb3a3B2Ab3b5827D"),
			Asset:        common.HexToAddress("0x455B51D882571E244d03668f1a458ca74E70d196"),
			Owner:        common.HexToAddress("0x5fb03eb21303d39967a1a119b32dd744a0fa8986"),
		},
		C_ETH_SEASCAPE_TEST: &TestConfig{
			EthUrl:       "https://ropsten.infura.io/v3/19e799349b424211b5758903de1c47ea",
			QueryAddress: common.HexToAddress("0xD2B67aeeA3A5e85AEe9F77E98db094d1E30A00Ed"),
			Asset:        common.HexToAddress("0x3680fb34f55030326659cd9aaec522b6e355bdb6"),
			Owner:        common.HexToAddress("0xc12E333cdD2843c7719aFfca036cDe023579F192"),
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
