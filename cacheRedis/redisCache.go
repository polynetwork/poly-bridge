package cacheRedis

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	goredis "github.com/go-redis/redis"
	"math/big"
	"poly-bridge/conf"
	"poly-bridge/models"
	"strconv"
	"sync"
	"time"
)

const (
	_CrossTxCounter        = "CrossTxCounter"
	_TransferStatisticResp = "TransferStatisticRes"
	_ShortTokenBalance     = "ShortTokenBalance"
	//getfee TokenBalance time.Hour*72
	_LongTokenBalance               = "LongTokenBalance"
	TxCheckBot                      = "TxCheckBot"
	LargeTxAlarmPrefix              = "LargeTxAlarm_"
	LargeTxList                     = "LargeTxList"
	MarkTxAsPaidPrefix              = "MarkTxAsPaid_"
	MarkTxAsSkipPrefix              = "MarkTxAsSkip_"
	StuckTxAlarmHasSendPrefix       = "StuckTxAlarmHasSendPrefix_"
	NodeStatusPrefix                = "NodeStatusPrefix_"
	NodeStatusAlarmPrefix           = "NodeStatusAlarmPrefix_"
	IgnoreNodeStatusAlarmPrefix     = "IgnoreNodeStatusAlarmPrefix_"
	ChainStatusPrefix               = "ChainStatusPrefix_"
	AssetBoundDstLockProxyPrefix    = "AssetBoundDstLockProxyPrefix_"
	_GetManualTxData                = "GetManualTxData_"
	RelayerAccountStatusPrefix      = "RelayerAccountStatusPrefix_"
	RelayerAccountStatusAlarmPrefix = "RelayerAccountStatusAlarmPrefix_"
	_ChainTVLAmount                 = "ChainTVLAmount_"
	MarkTokenAsDying                = "MarkTokenAsDying_"
)

type RedisCache struct {
	c      *goredis.Client
	config *conf.RedisConfig
}

var Redis *RedisCache
var mutex sync.Mutex

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
func (r *RedisCache) SetTokenBalance(srcChainId, dstChainId uint64, dstTokenHash string, tokenBalance *big.Int) (err error) {
	key := formatTokenBalanceKey(_ShortTokenBalance, srcChainId, dstChainId, dstTokenHash)
	value := tokenBalance.String()
	if _, err = r.c.Set(key, value, time.Second*2).Result(); err != nil {
		err = errors.New(err.Error() + "add SetTokenBalance")
	}
	return
}
func (r *RedisCache) GetTokenBalance(srcChainId, dstChainId uint64, dstTokenHash string) (*big.Int, error) {
	key := formatTokenBalanceKey(_ShortTokenBalance, srcChainId, dstChainId, dstTokenHash)
	resp, err := r.c.Get(key).Result()
	if err != nil {
		err = errors.New(err.Error() + "cache GetTokenBalance")
		return big.NewInt(0), err
	}
	balance, result := new(big.Int).SetString(resp, 10)
	if !result {
		return big.NewInt(0), errors.New("GetTokenBalance SetString err")
	}
	return balance, nil
}
func (r *RedisCache) SetLongTokenBalance(srcChainId, dstChainId uint64, dstTokenHash string, tokenBalance *big.Int) (err error) {
	key := formatTokenBalanceKey(_LongTokenBalance, srcChainId, dstChainId, dstTokenHash)
	value := tokenBalance.String()
	if _, err = r.c.Set(key, value, time.Hour*72).Result(); err != nil {
		err = errors.New(err.Error() + "add SetLongTokenBalance")
	}
	return
}
func (r *RedisCache) GetLongTokenBalance(srcChainId, dstChainId uint64, dstTokenHash string) (*big.Int, error) {
	key := formatTokenBalanceKey(_LongTokenBalance, srcChainId, dstChainId, dstTokenHash)
	resp, err := r.c.Get(key).Result()
	if err != nil {
		err = errors.New(err.Error() + "cache GetLongTokenBalance")
		return big.NewInt(0), err
	}
	balance, result := new(big.Int).SetString(resp, 10)
	if !result {
		return big.NewInt(0), errors.New("GetLongTokenBalance SetString err")
	}
	return balance, nil
}
func formatTokenBalanceKey(_key string, srcChainId, dstChainId uint64, dstTokenHash string) string {
	key := fmt.Sprintf("%s_%d_%d_%s", _key, srcChainId, dstChainId, dstTokenHash)
	return key
}

func (r *RedisCache) Get(key string) (string, error) {
	res, err := r.c.Get(key).Result()
	if err != nil {
		logs.Error("Get key %s err: %s", key, err)
		return "", err
	}
	return res, nil
}

func (r *RedisCache) Set(key string, value interface{}, expiration time.Duration) (bool, error) {
	err := r.c.Set(key, value, expiration).Err()
	if err != nil {
		logs.Error("Set key %s err: %s", key, err)
		return false, err
	}
	return true, nil
}

func (r *RedisCache) Unlink(key string) (int64, error) {
	cnt, err := r.c.Unlink(key).Result()
	if err != nil {
		logs.Error("Unlink key: %s err: %s", key, err)
		return 0, err
	}
	return cnt, nil
}

func (r *RedisCache) Del(key string) (int64, error) {
	cnt, err := r.c.Del(key).Result()
	if err != nil {
		logs.Error("Del key: %s err: %s", key, err)
		return 0, err
	}
	return cnt, nil
}

func (r *RedisCache) Exists(key string) (bool, error) {
	existed, err := r.c.Exists(key).Result()
	if err != nil {
		logs.Error("check key: %s exists err: %s", key, err)
		return false, err
	}
	if existed == 0 {
		return false, nil
	}
	return true, nil
}

func (r *RedisCache) Expire(key string, expiration time.Duration) (bool, error) {
	result, err := r.c.Expire(key, expiration).Result()
	if err != nil {
		return false, err
	}
	return result, nil
}

func (r *RedisCache) Lock(key string, value interface{}, expiration time.Duration) (bool, error) {
	mutex.Lock()
	defer mutex.Unlock()
	isSet, err := r.c.SetNX(key, value, expiration).Result()
	if err != nil {
		logs.Error("Lock err:%s", err)
		return false, err
	}
	return isSet, nil
}

func (r *RedisCache) UnLock(key string) (int64, error) {
	mutex.Lock()
	defer mutex.Unlock()
	cnt, err := r.c.Del(key).Result()
	if err != nil {
		logs.Error("UnLock err:%s", err)
		return 0, err
	}
	return cnt, nil
}
func (r *RedisCache) GetManualTx(polyhash string) (string, error) {
	key := _GetManualTxData + polyhash
	resp, err := r.c.Get(key).Result()
	if err != nil {
		err = errors.New(err.Error() + "cache GetManualTx")
		return "", err
	}
	return resp, nil
}
func (r *RedisCache) SetManualTx(polyhash string, manualTx string) (err error) {
	key := _GetManualTxData + polyhash
	value := manualTx
	if _, err = r.c.Set(key, value, time.Second*1).Result(); err != nil {
		err = errors.New(err.Error() + "cache SetManualTx")
	}
	return
}

func (r *RedisCache) RPush(key string, value ...interface{}) error {
	if err := r.c.RPush(key, value).Err(); err != nil {
		logs.Error("Redis Push[%s: %v] err: %s", key, value, err)
		return err
	}
	return nil
}

func (r *RedisCache) LRange(key string, start, stop int64) ([]string, error) {
	if vals, err := r.c.LRange(key, start, stop).Result(); err != nil {
		logs.Error("Redis LRange[key:%s, start:%d, stop:%d] err: %s", key, start, stop, err)
		return nil, err
	} else {
		return vals, nil
	}
}

func (r *RedisCache) SetChainTvl(chain uint64, amount string) (err error) {
	key := _ChainTVLAmount + fmt.Sprintf("%v", chain)
	value := amount
	if _, err = r.c.Set(key, value, time.Second*10).Result(); err != nil {
		err = errors.New(err.Error() + "cache SetChainTvl")
	}
	return
}

func (r *RedisCache) GetChainTvl(chain uint64) (amount string, err error) {
	key := _ChainTVLAmount + fmt.Sprintf("%v", chain)
	resp, err := r.c.Get(key).Result()
	if err != nil {
		err = errors.New(err.Error() + "cache GetChainTvl")
		return "0", err
	}
	return resp, nil
}
