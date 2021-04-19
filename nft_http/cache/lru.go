package cache
//
//import (
//	"fmt"
//	"poly-bridge/models"
//
//	lru "github.com/hashicorp/golang-lru"
//)
//
//type Cache struct {
//	lru lru.ARCCache
//}
//
//func NewLRU(size int) (*Cache, error) {
//	ins, err := lru.NewARC(size)
//	if err != nil {
//		return nil, err
//	}
//
//	return &Cache{lru: *ins}, nil
//}
//
//func (c *Cache) Set(asset string, tokenId *models.BigInt, profile *models.NFTProfile) {
//	key := c.format(asset, tokenId)
//	c.lru.Add(key, profile)
//}
//
//func (c *Cache) Get(asset string, tokenId *models.BigInt) (*models.NFTProfile, bool) {
//	key := c.format(asset, tokenId)
//	ptr, ok := c.lru.Get(key)
//	if !ok {
//		return nil, false
//	}
//	data, ok := ptr.(*models.NFTProfile)
//	if !ok {
//		return nil, false
//	}
//	return data, true
//}
//
//func (c *Cache) format(asset string, tokenId *models.BigInt) string {
//	return fmt.Sprintf("nft_%s_%s", asset, tokenId.String())
//}
