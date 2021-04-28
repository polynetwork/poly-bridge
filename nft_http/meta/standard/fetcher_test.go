package standard

import (
	"github.com/stretchr/testify/assert"
	. "poly-bridge/nft_http/meta/common"
	"testing"
)

func TestFetcher_Fetch(t *testing.T) {
	tokenId := "174663285876834235977250332211344399785"
	url := "https://digicol.io/ipfs/QmVG1FYVDKJk9X5cvahoswANis1TDBKUMagvEV9tX6ssB6"
	fetcher := NewFetcher("digicol", "https://digicol.io/ipfs/")
	profile, err := fetcher.Fetch(&FetchRequestParams{
		Url: url,
		TokenId: tokenId,
	})
	assert.NoError(t, err)

	t.Logf("%v", profile)
}
