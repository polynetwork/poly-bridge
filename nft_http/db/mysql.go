package nftdb

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/models"
)

var (
	db         *gorm.DB
	nativeHash = []string{"0000000000000000000000000000000000000000", "0000000000000000000000000000000000000103"}
)

func NewDB(cfg *conf.DBConfig) *gorm.DB {
	user := cfg.User
	password := cfg.Password
	url := cfg.URL
	scheme := cfg.Scheme
	Logger := logger.Default
	if cfg.Debug {
		Logger = Logger.LogMode(logger.Info)
	}
	format := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", user, password, url, scheme)
	var err error
	db, err = gorm.Open(mysql.Open(format), &gorm.Config{Logger: Logger})
	if err != nil {
		panic(err)
	}
	return db
}

func GetDB() *gorm.DB {
	return db
}

func InitNftAssets() []*models.Token {
	assets := make([]*models.Token, 0)
	db.Where("standard = ?", models.TokenTypeErc721).
		Preload("TokenBasic").
		Find(&assets)
	for _, v := range assets {
		logs.Info("load asset %s, chainid %d, hash %s", v.TokenBasicName, v.ChainId, v.Hash)
	}
	return assets
}
func InitFeeTokens() map[uint64]*models.Token {
	feeTokenList := make([]*models.Token, 0)
	feeTokens := make(map[uint64]*models.Token)
	db.Where("hash in ?", nativeHash).
		Preload("TokenBasic").
		Find(&feeTokenList)
	for _, v := range feeTokenList {
		feeTokens[v.ChainId] = v
		logs.Info("load chainid %d feeToken %s", v.ChainId, v.TokenBasicName)
	}
	return feeTokens
}

func FindFeeToken(cid uint64, hash string) *models.Token {
	feeTokens := make([]*models.Token, 0)
	db.Model(&models.Token{}).
		Where("hash in ?", nativeHash).
		Preload("TokenBasic").
		Find(&feeTokens)

	for _, v := range feeTokens {
		if cid == v.ChainId && hash == v.Hash {
			return v
		}
	}
	feeToken := new(models.Token)
	err := db.Model(&models.Token{}).
		Where("chain_id = ? and hash = ?", cid, hash).
		Preload("TokenBasic").
		First(feeToken).Error
	if err == nil {
		return feeToken
	}
	return nil
}

func InitNeo3MapAssetHash() map[string]string {
	tokenMap := make([]models.TokenMap, 0)
	db.Where("src_chain_id = ? AND standard = 1", basedef.NEO3_CROSSCHAIN_ID).Find(&tokenMap)
	res := make(map[string]string)
	for _, v := range tokenMap {
		res[v.DstTokenHash] = v.SrcTokenHash
	}
	return res
}
