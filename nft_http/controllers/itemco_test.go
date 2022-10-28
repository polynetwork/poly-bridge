package controllers

import (
	"fmt"
	mcm "poly-bridge/nft_http/meta/common"
	"sort"
	"testing"
	"time"
)

func Test_getItemsWithChainData(t *testing.T) {
	name := "CryptoSanguo"
	asset := "0x1819BFe00C0c0aEe24B88aCE7bff36d574d70180"
	chainId := uint64(701)
	tokenIdUrlMap := make(map[string]string)
	tokenIdUrlMap["3"] = "https://ikzttp.mypinata.cloud/ipfs/QmQFkLSQysj94s5GvTHPyzTxrawwtjgiiYS2TBLgrvw8CW/3"
	list := make([]*Item, 0)

	// get cache if exist
	profileReqs := make([]*mcm.FetchRequestParams, 0)
	for tokenId, url := range tokenIdUrlMap {
		cache, ok := GetItemCache(chainId, asset, tokenId)
		if ok {
			list = append(list, cache)
			delete(tokenIdUrlMap, tokenId)
			continue
		}

		req := &mcm.FetchRequestParams{
			TokenId: tokenId,
			Url:     url,
		}
		profileReqs = append(profileReqs, req)
	}

	// fetch metadata list and show rpc time
	tBeforeBatchFetch := time.Now().UnixNano()
	profiles, err := fetcher.BatchFetch(chainId, name, profileReqs)
	if err != nil {
		fmt.Printf("batch fetch NFT profiles err: %v", err)
	}
	tAfterBatchFetch := time.Now().UnixNano()
	debugBatchFetchTime := (tAfterBatchFetch - tBeforeBatchFetch) / int64(time.Microsecond)
	fmt.Printf("batchFetchNFTItems - batchFetchTime: %d microsecond， profiles %d", debugBatchFetchTime, len(profiles))

	// convert to items
	for _, v := range profiles {
		tokenId := v.NftTokenId
		item := new(Item).instance(name, tokenId, v)
		list = append(list, item)
		SetItemCache(chainId, asset, tokenId, item)
		delete(tokenIdUrlMap, tokenId)
	}
	for tokenId, _ := range tokenIdUrlMap {
		item := new(Item).instance(name, tokenId, nil)
		list = append(list, item)
	}

	// sort items with token id
	if len(list) < 2 {
		fmt.Println(list)
	}
	sort.Slice(list, func(i, j int) bool {
		itemi, _ := string2Big(list[i].TokenId)
		itemj, _ := string2Big(list[j].TokenId)
		return itemi.Cmp(itemj) < 0
	})

	fmt.Println(list)

}
