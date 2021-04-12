package chainsdk

import (
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

var (
	ethSdk *EthereumSdk

	// testnet
	wrapContract = common.HexToAddress("0x2E830E0cf3dc8643B497F88C07c8A72EFE24B11f")// bsc
	//wrapContract = common.HexToAddress("0x46Fc99509c4Aab0c958B8b2175edAa9C4963Ac09") //eth
	ethUrl = "https://data-seed-prebsc-2-s2.binance.org:8545" //bsc
	// ethUrl = "https://ropsten.infura.io/v3/19e799349b424211b5758903de1c47ea" // eth
	asset = common.HexToAddress("0x455B51D882571E244d03668f1a458ca74E70d196") //bsc
	//asset = common.HexToAddress("0xa85c9FC8F2c9060d674E0CA97F703a0A30619305") // eth
	owner = common.HexToAddress("0x5fb03eb21303d39967a1a119b32dd744a0fa8986")

	// devnet
	//wrapContract = common.HexToAddress("0x6DcD3dEfb145e59634d5281b19d555Cf933E4601")
	//ethUrl = "http://127.0.0.1:8545"
	//asset = common.HexToAddress("0x03d84da9432F7Cb5364A8b99286f97c59f738001")
	//owner = common.HexToAddress("0x5fb03eb21303d39967a1a119b32dd744a0fa8986")
)

func TestMain(m *testing.M) {
	var err error

	if ethSdk, err = NewEthereumSdk(ethUrl); err != nil {
		panic(err)
	}

	m.Run()
}
