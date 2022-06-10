package meta

import (
	"errors"
	"gorm.io/gorm"
	"poly-bridge/models"
	. "poly-bridge/nft_http/meta/common"
	"poly-bridge/nft_http/meta/standard"
)

type MetaFetcher interface {

	// fetch single nft profile by NFT asset name and NFT token full url
	Fetch(req *FetchRequestParams) (*models.NFTProfile, error)

	// batch fetch nft profiles by params of asset and url list
	BatchFetch(list []*FetchRequestParams) ([]*models.NFTProfile, error)
}

type FetcherType int

const (
	FetcherTypeUnknown  = 0
	FetcherTypeOpensea  = 1
	FetcherTypeStandard = 2
)

var ErrFetcherNotExist = errors.New("fetcher not exist")

func NewFetcher(fetcherTyp FetcherType, assetName, baseUri string) (fetcher MetaFetcher) {
	switch fetcherTyp {
	case FetcherTypeStandard, FetcherTypeOpensea:
		fetcher = standard.NewFetcher(assetName, baseUri)
	default:
		fetcher = nil
	}
	return
}

type StoreFetcher struct {
	fetcher      map[FetcherType]MetaFetcher
	assetFetcher map[uint64]map[string]FetcherType // mapping asset to fetcher type
	db           *gorm.DB
}

func NewStoreFetcher(orm *gorm.DB) *StoreFetcher {
	sf := new(StoreFetcher)
	sf.db = orm
	sf.fetcher = make(map[FetcherType]MetaFetcher)
	sf.assetFetcher = make(map[uint64]map[string]FetcherType)
	return sf
}

func (s *StoreFetcher) Register(ft FetcherType, chainId uint64, asset string, baseUri string) {
	fetcher := NewFetcher(ft, asset, baseUri)
	if fetcher == nil {
		return
	}
	s.fetcher[ft] = fetcher
	if _, ok := s.assetFetcher[chainId]; !ok {
		s.assetFetcher[chainId] = make(map[string]FetcherType)
	}
	s.assetFetcher[chainId][asset] = ft
}

func (s *StoreFetcher) selectFetcher(chainId uint64, asset string) MetaFetcher {
	if _, ok := s.assetFetcher[chainId]; !ok {
		return nil
	}

	typ, ok := s.assetFetcher[chainId][asset]
	if !ok {
		return nil
	}

	fetcher, ok := s.fetcher[typ]
	if !ok {
		return nil
	}
	return fetcher
}

func (s *StoreFetcher) Fetch(chainId uint64, asset string, req *FetchRequestParams) (profile *models.NFTProfile, err error) {
	fetcher := s.selectFetcher(chainId, asset)
	if fetcher == nil {
		return nil, ErrFetcherNotExist
	}

	profile = new(models.NFTProfile)
	res := s.db.Model(&models.NFTProfile{}).
		Where("token_basic_name = ? and nft_token_id = ? and name <> ''", asset, req.TokenId).
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

func (s *StoreFetcher) BatchFetch(chainId uint64, asset string, reqs []*FetchRequestParams) ([]*models.NFTProfile, error) {
	fetcher := s.selectFetcher(chainId, asset)
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
		delete(needFetchMap, v.NftTokenId)
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
