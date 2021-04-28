package meta

import (
	"fmt"
	"poly-bridge/conf"
	"poly-bridge/models"
	. "poly-bridge/nft_http/meta/common"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	configFile = "./../../conf/config_devnet.json"
	sf         *StoreFetcher

	asset   = "seascape"
	baseUri = "https://api.seascape.network/nft/metadata/"
)

func beforeTest() {
	cfg := conf.NewConfig(configFile)

	user := cfg.DBConfig.User
	password := cfg.DBConfig.Password
	scheme := cfg.DBConfig.Scheme
	url := cfg.DBConfig.URL
	lgr := logger.Default
	lgr = lgr.LogMode(logger.Info)

	db, err := gorm.Open(mysql.Open(user+":"+password+"@tcp("+url+")/"+scheme+"?charset=utf8"), &gorm.Config{Logger: lgr})
	if err != nil {
		panic(err)
	}

	sf := NewStoreFetcher(db)
	sf.Register(FetcherTypeStandard, asset, baseUri)
}

func TestStoreFetcher_FetchSeascape(t *testing.T) {
	beforeTest()

	tokenId := "141"
	profile, err := sf.Fetch(asset, &FetchRequestParams{
		TokenId: tokenId,
		Url:     "https://api.seascape.network/nft/metadata/141",
	})
	assert.NoError(t, err)
	t.Log(profile)

	persist := new(models.NFTProfile)
	res := sf.db.Where("token_basic_name = ? and nft_token_id = ?", asset, tokenId).Find(persist)
	assert.True(t, res.RowsAffected > 0)
	assert.Equal(t, persist.Name, profile.Name)
}

func TestStoreFetcher_BatchFetchSeascape(t *testing.T) {
	beforeTest()

	ids := []string{"137", "138", "140", "141"}
	reqs := make([]*FetchRequestParams, 0)
	for _, id := range ids {
		tid := fmt.Sprintf("%d", id)
		req := &FetchRequestParams{
			TokenId: tid,
			Url:     fmt.Sprintf("%s%d", baseUri, id)}
		reqs = append(reqs, req)
	}

	profile, err := sf.BatchFetch(asset, reqs)
	assert.NoError(t, err)
	t.Log(profile)

	persist := make([]*models.NFTProfile, 0)
	res := sf.db.Where("token_basic_name = ? and nft_token_id in (?)", asset, ids).Find(&persist)
	assert.True(t, res.RowsAffected == int64(len(ids)))
}
