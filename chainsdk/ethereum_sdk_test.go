package chainsdk

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var (
	// testnet
	//asset = common.HexToAddress("0xa85c9FC8F2c9060d674E0CA97F703a0A30619305")
	//owner = common.HexToAddress("0x5fb03eb21303d39967a1a119b32dd744a0fa8986")

	// devnet
	asset = common.HexToAddress("0x03d84da9432F7Cb5364A8b99286f97c59f738001")
	owner = common.HexToAddress("0x5fb03eb21303d39967a1a119b32dd744a0fa8986")
)

func TestNewEthereumSdk_GetTokensByIndex(t *testing.T) {
	start := 0
	length := 10
	data, err := ethSdk.GetTokensByIndex(wrapContract, asset, owner, start, length)
	assert.NoError(t, err)
	for tokenId, url := range data {
		t.Logf("token %d url %s", tokenId.Uint64(), url)
	}
}

func TestNewEthereumSdk_GetTokensById(t *testing.T) {
	tokenIds := []*big.Int{
		big.NewInt(1),
		big.NewInt(2),
	}

	data, err := ethSdk.GetTokensById(wrapContract, asset, tokenIds)
	assert.NoError(t, err)
	for tokenId, url := range data {
		t.Logf("token %d url %s", tokenId.Uint64(), url)
	}
}

func TestNewEthereumSdk_GetAndCheckTokenUrl(t *testing.T) {
	tokenId := big.NewInt(1)
	url, err := ethSdk.GetAndCheckTokenUrl(wrapContract, asset, owner, tokenId)
	assert.NoError(t, err)
	t.Logf("token %d url is %s", tokenId.Uint64(), url)
}