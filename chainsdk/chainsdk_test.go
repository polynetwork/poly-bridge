package chainsdk

import (
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

var (
	ethSdk *EthereumSdk

	// testnet
	//wrapContract = common.HexToAddress("0x46Fc99509c4Aab0c958B8b2175edAa9C4963Ac09")
	//ethUrl = "https://ropsten.infura.io/v3/19e799349b424211b5758903de1c47ea"

	// devnet
	wrapContract = common.HexToAddress("0x2E830E0cf3dc8643B497F88C07c8A72EFE24B11f")
	ethUrl = "http://127.0.0.1:8545"
)

func TestMain(m *testing.M) {
	var err error

	if ethSdk, err = NewEthereumSdk(ethUrl); err != nil {
		panic(err)
	}

	m.Run()
}
