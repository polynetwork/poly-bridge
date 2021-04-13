package chainsdk

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEthereumSdk_GetTokens(t *testing.T) {
	t.Logf("current context: %s", ctx.EthUrl)

	start := 0
	length := 10
	balance, err := ctx.SDK.GetNFTBalance(ctx.Asset, ctx.Owner)
	assert.NoError(t, err)
	t.Logf("user NFT balance %d", balance.Uint64())

	if balance.Uint64() == 0 {
		return
	}

	tokensByIndex, err := ctx.SDK.GetTokensByIndex(ctx.WrapAddress, ctx.Asset, ctx.Owner, start, length)
	assert.NoError(t, err)

	tokenIds := make([]*big.Int, 0)
	for tokenId, url := range tokensByIndex {
		t.Logf("getTokensByIndex: token %d url %s", tokenId.Uint64(), url)
		tokenIds = append(tokenIds, tokenId)
	}

	tokensWithId, err := ctx.SDK.GetTokensById(ctx.WrapAddress, ctx.Asset, tokenIds)
	assert.NoError(t, err)
	for tokenId, url := range tokensWithId {
		t.Logf("getTokensById: token %d url %s", tokenId.Uint64(), url)
	}

	for _, tokenId := range tokenIds {
		url, err := ctx.SDK.GetAndCheckTokenUrl(ctx.WrapAddress, ctx.Asset, ctx.Owner, tokenId)
		assert.NoError(t, err)
		t.Logf("getAndCheckTokenUrl: token %d url is %s", tokenId.Uint64(), url)
	}
}
