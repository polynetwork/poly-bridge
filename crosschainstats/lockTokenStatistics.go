package crosschainstats

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/shopspring/decimal"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/common"
	"poly-bridge/conf"
	"poly-bridge/models"
	"time"
)

type chainhashproxy struct {
	chainId   uint64
	hash      string
	ItemProxy string
}

var itemProxy2ItemName map[string]string

func assembleLockToken(chainid uint64, hash string, chainCfg []*conf.ChainListenConfig) []chainhashproxy {
	x := make([]chainhashproxy, 0)
	for _, chain := range chainCfg {
		if chain.ChainId == chainid {
			a := chainhashproxy{
				chainid,
				hash,
				chain.ProxyContract,
			}
			x = append(x, a)
			for _, other := range chain.OtherProxyContract {
				a := chainhashproxy{
					chainid,
					hash,
					other.ItemProxy,
				}
				x = append(x, a)
			}
		}
	}
	return x
}

func initItemProxyMap(chainCfg []*conf.ChainListenConfig) {
	mapItemProxy := make(map[string]string)
	for _, chain := range chainCfg {
		for _, other := range chain.OtherProxyContract {
			mapItemProxy[other.ItemProxy] = other.ItemName
		}
		mapItemProxy[chain.ProxyContract] = "poly"
	}
	itemProxy2ItemName = mapItemProxy
}

func (this *Stats) computeLockTokenStatistics() (err error) {
	logs.Info("start computeLockTokenStatistics")
	initItemProxyMap(this.chainCfg)
	lockTokenStatistics, err := this.dao.GetLockTokenStatistics()
	if err != nil {
		return fmt.Errorf("Failed to GetLockTokenStatistics %w", err)
	}
	tokenBasics, err := this.dao.GetTokenBasics()
	if err != nil {
		return fmt.Errorf("Failed to GetTokenBasics %w", err)
	}
	proxychainhashMap := make(map[chainhashproxy]bool)
	for _, lockTokenStatistic := range lockTokenStatistics {
		b := chainhashproxy{
			lockTokenStatistic.ChainId,
			lockTokenStatistic.Hash,
			lockTokenStatistic.ItemProxy,
		}
		proxychainhashMap[b] = true
	}
	for _, tokenBasic := range tokenBasics {
		if tokenBasic.Standard == uint8(0) && tokenBasic.ChainId != uint64(0) && tokenBasic.Tokens != nil {
			for _, token := range tokenBasic.Tokens {
				if token.ChainId == tokenBasic.ChainId {
					chainProxySlice := assembleLockToken(token.ChainId, token.Hash, this.chainCfg)
					for _, v := range chainProxySlice {
						if _, ok := proxychainhashMap[v]; !ok {
							proxychainhashMap[v] = true
							lockTokenStatistic := new(models.LockTokenStatistic)
							lockTokenStatistic.ChainId = v.chainId
							lockTokenStatistic.Hash = v.hash
							lockTokenStatistic.ItemProxy = v.ItemProxy
							lockTokenStatistic.ItemName = itemProxy2ItemName[v.ItemProxy]
							lockTokenStatistic.InAmount = models.NewBigIntFromInt(0)
							lockTokenStatistics = append(lockTokenStatistics, lockTokenStatistic)
						}
					}
				}
			}

		}
	}
	tokenBasicBTC, err := this.dao.GetBTCPrice()
	if err != nil {
		return fmt.Errorf("Failed to GetBTCPrice %w", err)
	}
	BTCPrice := decimal.NewFromInt(tokenBasicBTC.Price).Div(decimal.NewFromInt(basedef.PRICE_PRECISION))
	for _, lockTokenStatistic := range lockTokenStatistics {
		token, err := this.dao.GetTokenBasicByHash(lockTokenStatistic.ChainId, lockTokenStatistic.Hash)
		if err != nil {
			logs.Error("this_dao_GetTokenBasicByHash err", err)
			continue
		}
		price_new := decimal.New(token.TokenBasic.Price, 0).Div(decimal.NewFromInt(basedef.PRICE_PRECISION))
		precision_new := decimal.New(int64(1), int32(token.Precision))
		balance, err := getAndRetryProxyBalance(lockTokenStatistic.ChainId, lockTokenStatistic.Hash, lockTokenStatistic.ItemProxy)
		if err != nil {
			logs.Info("getAndRetryProxyBalance chainId: %v, Hash: %v, err:%v", lockTokenStatistic.ChainId, lockTokenStatistic.Hash, err)
		} else {
			amount_new := decimal.NewFromBigInt(balance, 0)
			lockTokenStatistic.InAmount = models.NewBigInt(amount_new.Div(precision_new).Mul(decimal.NewFromInt32(100)).BigInt())
		}
		amount_usd := decimal.NewFromBigInt(&lockTokenStatistic.InAmount.Int, 0).Mul(price_new)
		lockTokenStatistic.InAmountUsd = models.NewBigInt(amount_usd.Mul(decimal.NewFromInt32(100)).BigInt())
		amount_btc := amount_usd.Div(BTCPrice)
		lockTokenStatistic.InAmountBtc = models.NewBigInt(amount_btc.Mul(decimal.NewFromInt32(100)).BigInt())
		lockTokenStatistic.UpdateTime = uint64(time.Now().Unix())
	}
	err = this.dao.SaveLockTokenStatistics(lockTokenStatistics)
	if err != nil {
		logs.Error("SaveLockTokenStatistics err", err)
	}
	logs.Info("end computeLockTokenStatistics")
	return nil
}

func getAndRetryProxyBalance(chainId uint64, hash string, itemProxy string) (*big.Int, error) {
	balance, err := common.GetProxyBalance(chainId, hash, itemProxy)
	if err != nil {
		for i := 0; i < 4; i++ {
			time.Sleep(time.Second)
			balance, err = common.GetProxyBalance(chainId, hash, itemProxy)
			if err == nil {
				break
			}
		}
	}
	return balance, err
}
