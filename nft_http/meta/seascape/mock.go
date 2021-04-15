package seascape

import (
	"fmt"
	"math/big"
	"poly-bridge/models"
	. "poly-bridge/nft_http/meta/common"
	"poly-bridge/nft_http/meta/utils"
)

type MockSeascapeFetcher struct {
	Asset string // asset name, e.g: seascape
	BaseUri string // oss base uri of agency storage, e.g: https://api.seascape.network/nft/metadata/
}

func NewMockFetcher(assetName,baseUri string) *MockSeascapeFetcher {
	return &MockSeascapeFetcher{
		Asset:assetName,
		BaseUri: baseUri,
	}
}

func (f *MockSeascapeFetcher) Fetch(req *FetchRequestParams) (*models.NFTProfile, error) {
	raw, err := utils.Request(req.Url)
	if err != nil {
		return nil, err
	}

	origin := new(Profile)
	if err = origin.Unmarshal(raw); err != nil {
		return nil, err
	}

	return origin.Convert(f.Asset, req.TokenId)
}

func (f *MockSeascapeFetcher) BatchFetch(reqs []*FetchRequestParams) ([]*models.NFTProfile, error) {
	list := make([]*models.NFTProfile, 0)
	for _, v := range reqs {
		if data, err := f.Fetch(v); err == nil {
			list = append(list, data)
		}
	}
	return list, nil
}

func (f *MockSeascapeFetcher) FullUrl(tokenId *big.Int) string {
	return fmt.Sprintf("%s%d", f.BaseUri, tokenId.Uint64())
}