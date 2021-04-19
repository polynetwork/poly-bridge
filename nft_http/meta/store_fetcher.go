package meta

import (
	"errors"
	"math/big"
	"poly-bridge/models"
	. "poly-bridge/nft_http/meta/common"
	"poly-bridge/nft_http/meta/seascape"

	"gorm.io/gorm"
)

type MetaFetcher interface {

	// fetch single nft profile by NFT asset name and NFT token full url
	Fetch(req *FetchRequestParams) (*models.NFTProfile, error)

	// batch fetch nft profiles by params of asset and url list
	BatchFetch(list []*FetchRequestParams) ([]*models.NFTProfile, error)

	// full url format should be personality, e.g: fmt.Sprintf("%s%d", baseUri, tokenId)
	FullUrl(tokenId *big.Int) string
}

type FetcherType int

const (
	FetcherTypeUnknown = iota
	FetcherTypeSeascape
	FetcherTypeMockSeascape
)

var ErrFetcherNotExist = errors.New("fetcher not exist")

func NewFetcher(fetcherTyp FetcherType, assetName, baseUri string) (fetcher MetaFetcher) {
	switch fetcherTyp {
	case FetcherTypeMockSeascape:
		fetcher = seascape.NewMockFetcher(assetName, baseUri)
	case FetcherTypeSeascape:
		fetcher = seascape.NewFetcher(assetName, baseUri)
	default:
		fetcher = nil
	}
	return
}

type StoreFetcher struct {
	fetcher      map[FetcherType]MetaFetcher
	assetFetcher map[string]FetcherType // mapping asset to fetcher type
	db           *gorm.DB
}

func NewStoreFetcher(orm *gorm.DB) *StoreFetcher {
	sf := new(StoreFetcher)
	sf.db = orm
	sf.fetcher = make(map[FetcherType]MetaFetcher)
	sf.assetFetcher = make(map[string]FetcherType)
	return sf
}

func (s *StoreFetcher) Register(ft FetcherType, asset string, baseUri string) {
	fetcher := NewFetcher(ft, asset, baseUri)
	if fetcher == nil {
		return
	}
	s.fetcher[ft] = fetcher
	s.assetFetcher[asset] = ft
}

func (s *StoreFetcher) selectFetcher(asset string) MetaFetcher {
	typ, ok := s.assetFetcher[asset]
	if !ok {
		return nil
	}
	fetcher, ok := s.fetcher[typ]
	if !ok {
		return nil
	}
	return fetcher
}

func (s *StoreFetcher) Fetch(asset string, req *FetchRequestParams) (profile *models.NFTProfile, err error) {
	fetcher := s.selectFetcher(asset)
	if fetcher == nil {
		return nil, ErrFetcherNotExist
	}

	profile = new(models.NFTProfile)
	res := s.db.Model(&models.NFTProfile{}).
		Where("token_basic_name = ? and nft_token_id = ?", asset, req.TokenId).
		Find(profile)
	if res.RowsAffected > 0 {
		return profile, nil
	}

	if profile, err = fetcher.Fetch(req); err != nil {
		return nil, err
	}

	s.db.Save(profile)
	return
}

func (s *StoreFetcher) BatchFetch(asset string, reqs []*FetchRequestParams) ([]*models.NFTProfile, error) {
	fetcher := s.selectFetcher(asset)
	if fetcher == nil {
		return nil, ErrFetcherNotExist
	}

	finalList := make([]*models.NFTProfile, 0)
	unCacheList := make([]string, 0)
	needFetchMap := make(map[string]*FetchRequestParams, 0)

	for _, v := range reqs {
		unCacheList = append(unCacheList, v.TokenId)
		needFetchMap[v.TokenId] = v
	}

	persisted := make([]*models.NFTProfile, 0)
	s.db.Where("token_basic_name = ? and nft_token_id in (?)", asset, unCacheList).Find(&persisted)
	for _, v := range persisted {
		finalList = append(finalList, v)
		delete(needFetchMap, v.NftTokenId.String())
	}

	needFetchList := make([]*FetchRequestParams, 0)
	for _, v := range needFetchMap {
		needFetchList = append(needFetchList, v)
	}
	if len(needFetchList) == 0 {
		return finalList, nil
	}

	profiles, err := fetcher.BatchFetch(needFetchList)
	if err != nil {
		return nil, err
	}
	s.db.Save(profiles)

	for _, v := range profiles {
		finalList = append(finalList, v)
	}

	return finalList, nil
}
