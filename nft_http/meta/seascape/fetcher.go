package seascape

import (
	"math/big"
	"poly-bridge/models"
	. "poly-bridge/nft_http/meta/common"
)

type SeascapeFetcher struct {
	Asset   string
	BaseUri string
}

func NewFetcher(asset, baseUri string) *SeascapeFetcher {
	return &SeascapeFetcher{}
}

func (f *SeascapeFetcher) Fetch(req *FetchRequestParams) (*models.NFTProfile, error) {
	return nil, nil
}

// batch fetch nft profiles by params of asset and url list
func (f *SeascapeFetcher) BatchFetch(reqs []*FetchRequestParams) ([]*models.NFTProfile, error) {
	return nil, nil
}

func (f *SeascapeFetcher) FullUrl(tokenId *big.Int) string {
	return ""
}
