package standard

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	. "poly-bridge/nft_http/meta/common"
	"testing"
)

func TestFetcher_Fetch(t *testing.T) {
	tokenId := "174663285876834235977250332211344399785"
	url := "https://ipfs.io/ipfs/bafybeicvths73jv4pvczmebgv2beyurh74p6do2oyyu37p6v6vh372leqi/315.json"
	fetcher := NewFetcher("digicol", "https://digicol.io/ipfs/")
	profile, err := fetcher.Fetch(&FetchRequestParams{
		Url:     url,
		TokenId: tokenId,
	})
	assert.NoError(t, err)
	//
	//t.Logf("%v", profile)
	a, _ := json.MarshalIndent(profile, "", "	")
	fmt.Println(string(a))
}
