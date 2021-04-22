package chainsdk

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/stretchr/testify/assert"
	"math/big"
	nftwrap "poly-bridge/go_abi/nft_wrap_abi"
	pabi "poly-bridge/utils/abi"
	"strings"
	"testing"
)

func TestNewEthereumSdk_GetAndCheckOwnerNFT(t *testing.T) {
	t.Logf("current context: %s", ctx.EthUrl)

	tokenId := big.NewInt(201)
	url, err := ctx.SDK.GetAndCheckNFTUrl(ctx.WrapAddress, ctx.Asset, ctx.Owner, tokenId)
	assert.NoError(t, err)
	t.Logf("token %d url %s", tokenId.Uint64(), url)
}

func TestNewEthereumSdk_GetUserTokenIds(t *testing.T) {
	t.Logf("current context: %s", ctx.EthUrl)

	index := 1
	tokenId, err := ctx.SDK.GetOwnerNFTByIndex(ctx.Asset, ctx.Owner, index)
	assert.NoError(t, err)
	t.Logf("%s %dth NFT is %d", ctx.Owner.Hex(), index, tokenId.Uint64())
}

func TestNewEthereumSdk_GetTokens(t *testing.T) {
	t.Logf("current context: %s", ctx.EthUrl)

	start := 0
	length := 10

	totalSupply, err := ctx.SDK.GetNFTTotalSupply(ctx.Asset)
	assert.NoError(t, err)
	t.Logf("total supply %d", totalSupply.Uint64())

	ignore := EmptyAddress
	data, err := ctx.SDK.GetUnCrossChainNFTsByIndex(ctx.WrapAddress, ctx.Asset, ignore, start, length)
	assert.NoError(t, err)
	for tokenId, url := range data {
		t.Logf("getUnCrossChainNFTsByIndex: token %s url is %s", tokenId, url)
	}

	balance, err := ctx.SDK.GetNFTBalance(ctx.Asset, ctx.Owner)
	assert.NoError(t, err)
	t.Logf("user NFT balance %d", balance.Uint64())

	if balance.Uint64() == 0 || totalSupply.Uint64() == 0 {
		return
	}

	tokensByIndex, err := ctx.SDK.GetOwnerNFTsByIndex(ctx.WrapAddress, ctx.Asset, ctx.Owner, start, length)
	assert.NoError(t, err)

	tokenIds := make([]*big.Int, 0)
	for tokenId, url := range tokensByIndex {
		t.Logf("getTokensByIndex: token %s url %s", tokenId, url)
		tid, _ := new(big.Int).SetString(tokenId, 10)
		tokenIds = append(tokenIds, tid)
	}

	tokensWithId, err := ctx.SDK.GetNFTsById(ctx.WrapAddress, ctx.Asset, tokenIds)
	assert.NoError(t, err)
	for tokenId, url := range tokensWithId {
		t.Logf("getTokensById: token %s url %s", tokenId, url)
	}

	for _, tokenId := range tokenIds {
		url, err := ctx.SDK.GetAndCheckNFTUrl(ctx.WrapAddress, ctx.Asset, ctx.Owner, tokenId)
		assert.NoError(t, err)
		t.Logf("getAndCheckTokenUrl: token %d url is %s", tokenId.Uint64(), url)
	}
}

func TestABIUnpackWrapperLockParameters(t *testing.T) {
	code := "0985b87f0000000000000000000000000c3c33da088abeee376418d3e384528c5aadba11000000000000000000000000000000000000000000000000000000000000004f000000000000000000000000a107c23029c31da1b5ab19eab8228a2a44024c7d00000000000000000000000000000000000000000000000000000000000000c90000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000009898b76ae74c000000000000000000000000000000000000000000000000000000000000000000"
	abiStr := strings.NewReader(nftwrap.PolyNFTWrapperABI)
	wrapperABI, err := abi.JSON(abiStr)
	assert.NoError(t, err)

	enc, err := hex.DecodeString(code)
	assert.NoError(t, err)

	data := new(WrapLockMethod)
	err = pabi.UnpackMethod(wrapperABI, "lock", data, enc[:])
	assert.NoError(t, err)

	t.Logf("data: {\r\n toChainId %d\r\n tokenId %d\r\n fromAsset %s\r\n toAddress %s\r\n feeToken %s\r\n fee %s\r\n dataId %d\r\n}",
		data.ToChainId, data.TokenId.Uint64(), data.FromAsset.Hex(), data.ToAddress.Hex(), data.FeeToken.Hex(), data.Fee.String(), data.Id)
}
