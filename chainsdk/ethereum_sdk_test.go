package chainsdk

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
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
		big.NewInt(10),
		big.NewInt(12),
		big.NewInt(13),
		big.NewInt(14),
		big.NewInt(16),
		big.NewInt(17),
	}

	data, err := ethSdk.GetTokensById(wrapContract, asset, tokenIds)
	assert.NoError(t, err)
	for tokenId, url := range data {
		t.Logf("token %d url %s", tokenId.Uint64(), url)
	}
}

func TestNewEthereumSdk_GetAndCheckTokenUrl(t *testing.T) {
	tokenId := big.NewInt(12)
	url, err := ethSdk.GetAndCheckTokenUrl(wrapContract, asset, owner, tokenId)
	assert.NoError(t, err)
	t.Logf("token %d url is %s", tokenId.Uint64(), url)
}
