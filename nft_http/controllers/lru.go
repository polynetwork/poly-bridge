package controllers

import "fmt"

func SetHomePageCache(chainId uint64, rsp *HomeRsp) {
	key := formatHomePageCacheKey(chainId)
	lruDB.Add(key, rsp)
}

func GetHomePageCache(chainId uint64) (*HomeRsp, bool) {
	key := formatHomePageCacheKey(chainId)
	data, ok := lruDB.Get(key)
	if !ok {
		return nil, false
	}
	rsp, ok := data.(*HomeRsp)
	if !ok {
		return nil, false
	}
	return rsp, true
}

func SetItemCache(chainId uint64, asset string, tokenId string, item *Item) {
	key := formatItemKey(chainId, asset, tokenId)
	lruDB.Add(key, item)
}

func GetItemCache(chainId uint64, asset string, tokenId string) (*Item, bool) {
	key := formatItemKey(chainId, asset, tokenId)
	data, ok := lruDB.Get(key)
	if !ok {
		return nil, false
	}
	item, ok := data.(*Item)
	if !ok {
		return nil, false
	}
	return item, true
}

func formatHomePageCacheKey(chainId uint64) string {
	return fmt.Sprintf("homepage_%d", chainId)
}

func formatItemKey(chainId uint64, asset string, tokenId string) string {
	return fmt.Sprintf("item_%d_%s_%s", chainId, asset, tokenId)
}