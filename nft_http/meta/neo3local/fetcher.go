package neo3local

import (
	log "github.com/beego/beego/v2/core/logs"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/models"
	nftdb "poly-bridge/nft_http/db"
	. "poly-bridge/nft_http/meta/common"
	"poly-bridge/nft_http/meta/utils"
	"strings"
)

type Fetcher struct {
	Asset   string // asset name, e.g: seascape
	BaseUri string // oss base uri of agency storage, e.g: https://api.seascape.network/nft/metadata/
}

var Neo3JsonUrlLocal string
var Neo3ImageUrlLocal string
var Neo3AssetMap map[string]string

func NewFetcher(assetName, baseUri string) *Fetcher {
	Neo3JsonUrlLocal = conf.GlobalConfig.Neo3LocalNftFetcherConfig.Neo3JsonUrlLocal
	Neo3ImageUrlLocal = conf.GlobalConfig.Neo3LocalNftFetcherConfig.Neo3ImageUrlLocal
	if Neo3JsonUrlLocal == "" || Neo3ImageUrlLocal == "" {
		panic("no Neo3JsonUrlLocal or Neo3ImageUrlLocal provided for neo3 nft fetcher")
	}
	Neo3AssetMap = nftdb.InitNeo3MapAssetHash()
	return &Fetcher{
		Asset:   assetName,
		BaseUri: baseUri,
	}
}

func (f *Fetcher) Fetch(req *FetchRequestParams) (*models.NFTProfile, error) {
	var assetHash string
	var err error
	if req.ChainId == basedef.NEO3_CROSSCHAIN_ID {
		assetHash = req.AssetHash
	} else {
		if hash, ok := Neo3AssetMap[req.AssetHash]; !ok {
			return nil, err
		} else {
			assetHash = hash
		}
	}
	url := AddNeoLocalPrefix(req.Url, assetHash, Neo3JsonUrlLocal)
	raw, err := utils.Request(url)
	if err != nil {
		log.Error("fail to request %v, err: %v", url, err)
		return nil, err
	}

	origin := new(Profile)
	if err = origin.Unmarshal(raw); err != nil {
		log.Error("fail to read resp from %v, err: %v", url, err)
		return nil, err
	}

	return origin.Convert(f.Asset, req.TokenId, req.Url, assetHash, Neo3ImageUrlLocal)
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
	data, _ := f.Fetch(req)
	ch <- data
}

func AddNeoLocalPrefix(url, assetHash, prefix string) string {
	if strings.HasPrefix(url, "ipfs.io") {
		url = "https:" + url
	} else if strings.HasPrefix(url, "ipfs://") {
		url = strings.TrimPrefix(url, "ipfs://")
		url = "https:ipfs.ioipfs" + url
	}
	if !strings.HasPrefix(assetHash, "0x") {
		assetHash = "0x" + assetHash
	}
	return prefix + assetHash + "/" + strings.ReplaceAll(url, "/", "")
}
