package controllers

import (
	"fmt"
	"time"
)

type CacheHomeRsp struct {
	Rsp  *HomeRsp
	Time time.Time
}

func SetHomePageItemsCache(chainId uint64, assetName string, list []*Item) {
	key := formatHomePageItemsCache(chainId, assetName)
	lruDB.Add(key, list)
}

func GetHomePageItemsCache(chainId uint64, assetName string) ([]*Item, bool) {
	key := formatHomePageItemsCache(chainId, assetName)
	data, ok := lruDB.Get(key)
	if !ok {
		return nil, false
	}
	list, ok := data.([]*Item)
	if !ok {
		return nil, false
	}
	return list, true
}

//func SetHomePageCache(chainId uint64, start, length int, rsp *CacheHomeRsp) {
//	key := formatHomePageCacheKey(chainId, start, length)
//	lruDB.Add(key, rsp)
//}
//
//func GetHomePageCache(chainId uint64, start, len int) (*CacheHomeRsp, bool) {
//	key := formatHomePageCacheKey(chainId, start, len)
//	data, ok := lruDB.Get(key)
//	if !ok {
//		return nil, false
//	}
//	rsp, ok := data.(*CacheHomeRsp)
//	if !ok {
//		return nil, false
//	}
//	return rsp, true
//}

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

func formatHomePageItemsCache(chainId uint64, assetName string) string {
	return fmt.Sprintf("homepage_%d_%s", chainId, assetName)
}

//func formatHomePageCacheKey(chainId uint64, start, len int) string {
//	return fmt.Sprintf("homepage_%d_%d_%d", chainId, start, len)
//}

func formatItemKey(chainId uint64, asset string, tokenId string) string {
	return fmt.Sprintf("item_%d_%s_%s", chainId, asset, tokenId)
}
