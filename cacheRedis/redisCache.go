package cacheRedis

import (
	"encoding/json"
	"errors"
	"fmt"
	goredis "github.com/go-redis/redis"
	"math/big"
	"poly-bridge/conf"
	"poly-bridge/models"
	"strconv"
	"time"
)

const (
	_CrossTxCounter        = "CrossTxCounter"
	_TransferStatisticResp = "TransferStatisticRes"
	//getfee TokenBalance time.Hour*72
	_TokenBalance = "TokenBalance"
)

type RedisCache struct {
	c      *goredis.Client
	config *conf.RedisConfig
}

var Redis *RedisCache

func Init() {
	redisConfig := conf.GlobalConfig.RedisConfig
	redisCache, err := GetRedisClient(redisConfig)
	if err != nil {
		panic("redis Init panic")
	}
	Redis = redisCache
}

func GetRedisClient(redisConfig *conf.RedisConfig) (*RedisCache, error) {
	if redisConfig.DialTimeout <= 0 || redisConfig.ReadTimeout <= 0 || redisConfig.WriteTimeout <= 0 {
		return &RedisCache{
			c:      &goredis.Client{},
			config: redisConfig,
		}, errors.New("DialTimeout ReadTimeout WriteTimeout must exist")
	}
	options := &goredis.Options{
		Network:      redisConfig.Proto,
		Addr:         redisConfig.Addr,
		Password:     redisConfig.Password,
		DialTimeout:  redisConfig.DialTimeout * time.Second,
		ReadTimeout:  redisConfig.ReadTimeout * time.Second,
		WriteTimeout: redisConfig.WriteTimeout * time.Second,
		PoolSize:     redisConfig.PoolSize,
		IdleTimeout:  redisConfig.IdleTimeout * time.Second,
	}
	return &RedisCache{
		c:      goredis.NewClient(options),
		config: redisConfig,
	}, nil
}

func (r *RedisCache) SetCrossTxCounter(counter int64) (err error) {
	key := _CrossTxCounter
	if _, err = r.c.Set(key, counter, r.config.Expiration*time.Second).Result(); err != nil {
		err = errors.New(err.Error() + "add SetCrossTxCounter")
	}
	return
}
func (r *RedisCache) GetCrossTxCounter() (counter int64, err error) {
	key := _CrossTxCounter
	resp, err := r.c.Get(key).Result()
	if err != nil {
		err = errors.New(err.Error() + "cache GetCrossTxCounter")
		return
	}
	count, err := strconv.Atoi(resp)
	counter = int64(count)
	if err != nil {
		err = errors.New(err.Error() + "cache GetCrossTxCounter Atoi")
	}
	return
}

func (r *RedisCache) SetAllTransferResp(resp *models.AllTransferStatisticResp) (err error) {
	key := _TransferStatisticResp
	jsons, err := json.Marshal(resp)
	if _, err = r.c.Set(key, string(jsons), time.Second*60).Result(); err != nil {
		err = errors.New(err.Error() + "add SetAllTransferResp")
	}
	return
}
func (r *RedisCache) GetAllTransferResp() (*models.AllTransferStatisticResp, error) {
	key := _TransferStatisticResp
	jsons, err := r.c.Get(key).Result()
	if err != nil {
		err = errors.New(err.Error() + "cache GetAllTransferResp")
		return nil, err
	}
	resp := new(models.AllTransferStatisticResp)
	err = json.Unmarshal([]byte(jsons), resp)
	if err != nil {
		err = errors.New(err.Error() + "cache GetAllTransferResp")
		return nil, err
	}
	return resp, nil
}

func (r *RedisCache) GetTokenBalance(dstChainId uint64, dstTokenHash string) (*big.Int, error) {
	key := formatTokenBalanceKey(dstChainId, dstTokenHash)
	resp, err := r.c.Get(key).Result()
	if err != nil {
		err = errors.New(err.Error() + "cache GetCrossTxCounter")
		return big.NewInt(0), err
	}
	balance, result := new(big.Int).SetString(resp, 10)
	if !result {
		return big.NewInt(0), errors.New("GetTokenBalance SetString err")
	}
	return balance, nil
}
func (r *RedisCache) SetTokenBalance(dstChainId uint64, dstTokenHash string, tokenBalance *big.Int) (err error) {
	key := formatTokenBalanceKey(dstChainId, dstTokenHash)
	value := tokenBalance.String()
	if _, err = r.c.Set(key, value, time.Hour*72).Result(); err != nil {
		err = errors.New(err.Error() + "add SetAllTransferResp")
	}
	return
}
func formatTokenBalanceKey(dstChainId uint64, dstTokenHash string) string {
	key := fmt.Sprintf("%s_%d_%s", _TokenBalance, dstChainId, dstTokenHash)
	return key
}
