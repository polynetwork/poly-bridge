package standard

import (
	log "github.com/beego/beego/v2/core/logs"
	"poly-bridge/models"
	. "poly-bridge/nft_http/meta/common"
	"poly-bridge/nft_http/meta/utils"
)

type Fetcher struct {
	Asset   string // asset name, e.g: seascape
	BaseUri string // oss base uri of agency storage, e.g: https://api.seascape.network/nft/metadata/
}

func NewFetcher(assetName, baseUri string) *Fetcher {
	return &Fetcher{
		Asset:   assetName,
		BaseUri: baseUri,
	}
}

func (f *Fetcher) Fetch(req *FetchRequestParams) (*models.NFTProfile, error) {
	raw, err := utils.Request(req.Url)
	if err != nil {
		return nil, err
	}

	origin := new(Profile)
	if err = origin.Unmarshal(raw); err != nil {
		return nil, err
	}

	return origin.Convert(f.Asset, req.TokenId, req.Url)
}

func (f *Fetcher) BatchFetch(reqs []*FetchRequestParams) ([]*models.NFTProfile, error) {
	list := make([]*models.NFTProfile, 0)
	ch := make(chan *models.NFTProfile)
	for _, v := range reqs {
		go f.concurrentFetch(v, ch)
	}
	for range reqs {
		if data := <-ch; data != nil {
			list = append(list, data)
		}
	}
	return list, nil
}

func (f *Fetcher) concurrentFetch(req *FetchRequestParams, ch chan<- *models.NFTProfile) {
	data, err := f.Fetch(req)
	if err != nil {
		log.Error("fail to request %v, err: %v", req.Url, err)
	}
	ch <- data
}
