package meta

import (
	"fmt"
	"math/big"
	"poly-bridge/models"
	"poly-bridge/nft_http/meta/cache"
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

func NewFetcher(fetcherTyp FetcherType, assetName, baseUri string) (MetaFetcher, error) {
	var fetcher MetaFetcher

	switch fetcherTyp {
	case FetcherTypeMockSeascape:
		fetcher = seascape.NewMockFetcher(assetName, baseUri)
	case FetcherTypeSeascape:
		fetcher = seascape.NewFetcher(assetName, baseUri)
	default:
		fetcher = nil
	}
	if fetcher == nil {
		return nil, fmt.Errorf("invalid fetcher type")
	}
	return fetcher, nil
}

type StoreFetcher struct {
	fetcher      map[FetcherType]MetaFetcher
	assetFetcher map[string]FetcherType // mapping asset to fetcher type
	db           *gorm.DB
	cache        *cache.Cache
}

func NewStoreFetcher(orm *gorm.DB, cacheSize int) (*StoreFetcher, error) {
	sf := new(StoreFetcher)
	sf.db = orm
	c, err := cache.NewLRU(cacheSize)
	if err != nil {
		return nil, err
	}
	sf.cache = c
	sf.fetcher = make(map[FetcherType]MetaFetcher)
	sf.assetFetcher = make(map[string]FetcherType)
	return sf, nil
}

func (s *StoreFetcher) Register(ft FetcherType, asset string, baseUri string) error {
	fetcher, err := NewFetcher(ft, asset, baseUri)
	if err != nil {
		return err
	}
	s.fetcher[ft] = fetcher
	s.assetFetcher[asset] = ft
	return nil
}

func (s *StoreFetcher) selectFetcher(asset string) MetaFetcher {
	typ := s.assetFetcher[asset]
	fetcher := s.fetcher[typ]
	return fetcher
}

func (s *StoreFetcher) Fetch(asset string, req *FetchRequestParams) (profile *models.NFTProfile, err error) {
	var ok bool
	if profile, ok = s.cache.Get(asset, req.TokenId); ok {
		return
	}

	profile = new(models.NFTProfile)
	res := s.db.Model(&models.NFTProfile{}).
		Where("token_basic_name = ? and nft_token_id = ?", asset, req.TokenId).
		Find(profile)
	if res.RowsAffected > 0 && profile.Name != "" {
		return profile, nil
	}

	fetcher := s.selectFetcher(asset)
	if profile, err = fetcher.Fetch(req); err != nil {
		return nil, err
	}

	s.db.Save(profile)
	s.cache.Set(asset, req.TokenId, profile)
	return
}

func (s *StoreFetcher) BatchFetch(asset string, reqs []*FetchRequestParams) ([]*models.NFTProfile, error) {
	finalList := make([]*models.NFTProfile, 0)
	uncacheList := make([]*models.BigInt, 0)
	needFetchMap := make(map[string]*FetchRequestParams, 0)

	for _, v := range reqs {
		tid := v.TokenId.String()

		if cached, ok := s.cache.Get(asset, v.TokenId); ok {
			finalList = append(finalList, cached)
			continue
		}

		uncacheList = append(uncacheList, v.TokenId)
		needFetchMap[tid] = v
	}

	persisted := make([]*models.NFTProfile, 0)
	s.db.Where("token_basic_name = ? and nft_token_id in (?)", asset, uncacheList).Find(&persisted)
	for _,  v := range persisted {
		tid := v.NftTokenId.String()
		finalList = append(finalList, v)
		delete(needFetchMap, tid)
		s.cache.Set(asset, v.NftTokenId, v)
	}

	needFetchList := make([]*FetchRequestParams, 0)
	for _, v := range needFetchMap {
		needFetchList = append(needFetchList, v)
	}
	if len(needFetchList) == 0 {
		return finalList, nil
	}

	fetcher := s.selectFetcher(asset)
	profiles, err := fetcher.BatchFetch(needFetchList)
	if err != nil {
		return nil, err
	}

	for _, v := range profiles {
		s.cache.Set(asset, v.NftTokenId, v)
	}
	s.db.Save(profiles)

	for _, v := range profiles {
		finalList = append(finalList, v)
	}

	return finalList, nil
}