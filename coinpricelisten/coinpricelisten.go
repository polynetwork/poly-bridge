package coinpricelisten

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/prometheus/common/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/big"
	"poly-swap/coinpricelisten/binance"
	"poly-swap/coinpricelisten/coinmarketcap"
	"poly-swap/conf"
	"poly-swap/models"
	"runtime/debug"
	"strings"
	"time"
)

type CoinPriceListen struct {
	cmcCfg   *conf.CoinMarketCapPriceListenConfig
	binCfg   *conf.BinancePriceListenConfig
	dbCfg    *conf.DBConfig
	priceUpdateSlot int64
	db       *gorm.DB
}

func StartCoinPriceListen(cfg *conf.CoinPriceListenConfig, dbCfg *conf.DBConfig) {
	cpListen := NewCoinPriceListen(cfg.CoinMarketCapPriceListenConfig, cfg.BinancePriceListenConfig, cfg.PriceUpdateSlot, dbCfg)
	cpListen.Start()
}

func NewCoinPriceListen(cmcCfg *conf.CoinMarketCapPriceListenConfig, binCfg *conf.BinancePriceListenConfig, priceUpdateSlot int64, dbCfg *conf.DBConfig) *CoinPriceListen {
	cpListen := &CoinPriceListen{}
	cpListen.cmcCfg = cmcCfg
	cpListen.binCfg = binCfg
	cpListen.dbCfg = dbCfg
	cpListen.priceUpdateSlot = priceUpdateSlot
	db, err := gorm.Open(mysql.Open(dbCfg.User + ":" + dbCfg.Password + "@tcp(" + dbCfg.URL + ")/" +
		dbCfg.Scheme + "?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	cpListen.db = db
	//
	tokenBasics := make([]*models.TokenBasic, 0)
	res := db.Find(&tokenBasics)
	if res.RowsAffected == 0 {
		panic("there is no token basic!")
	}
	err = cpListen.getCoinPrice(tokenBasics)
	if err != nil {
		panic(err)
	}
	db.Debug().Save(tokenBasics)
	return cpListen
}


func (this *CoinPriceListen) Start() {
	go this.ListenPrice()
}

func (this *CoinPriceListen) ListenPrice() {
	for {
		this.listenPrice()
	}
}

func (this *CoinPriceListen) listenPrice() {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("service start, recover info: %s", string(debug.Stack()))
		}
	}()

	logs.Debug("listen price......")
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			now := time.Now().Unix() / 60
			if now % this.priceUpdateSlot != 0 {
				continue
			}
			log.Infof("do price update at time: %s", time.Now().Format("2006-01-02 15:04:05"))
			tokenBasics := make([]*models.TokenBasic, 0)
			res := this.db.Find(&tokenBasics)
			if res.RowsAffected == 0 {
				log.Errorf("there is no token basic!")
				continue
			}
			err := this.getCoinPrice(tokenBasics)
			if err != nil {
				log.Errorf("updateCoinPrice err: %v", err)
				continue
			}
			this.db.Save(tokenBasics)
		}
	}
}

func (this *CoinPriceListen) getCoinPrice(tokenBasics []*models.TokenBasic) error {
	tokenCmcNames := make([]string, 0)
	tokenBinNames := make([]string, 0)
	for _, item := range tokenBasics {
		tokenCmcNames = append(tokenCmcNames, item.CmcName)
		tokenBinNames = append(tokenBinNames, item.BinName)
	}
	cmcPrices, cmcerr := this.getCmcCoinPrice(tokenCmcNames)
	binPrices, binerr := this.getBinancePrice(tokenBinNames)
	if cmcerr != nil && binerr != nil {
		return fmt.Errorf("cmcerr: %s, binerr: %s", cmcerr.Error(), binerr.Error())
	}
	for _, token := range tokenBasics {
		avgPrices := make([]uint64, 0)
		cmcPrice, ok := cmcPrices[token.CmcName]
		if ok {
			price, _ := new(big.Float).Mul(big.NewFloat(cmcPrice), big.NewFloat(100000000)).Uint64()
			token.CmcPrice = price
			token.CmcInd = 1
			token.Time = uint64(time.Now().Unix())
			avgPrices = append(avgPrices, price)
		} else {
			token.CmcInd = 0
			log.Errorf("can not get the coinmarketcap price of coin %s", token.CmcName)
		}
		binPrice, ok := binPrices[token.BinName]
		if ok {
			price, _ := new(big.Float).Mul(big.NewFloat(binPrice), big.NewFloat(100000000)).Uint64()
			token.BinPrice = price
			token.BinInd = 1
			token.Time = uint64(time.Now().Unix())
			avgPrices = append(avgPrices, price)
		} else {
			token.BinInd = 0
			log.Errorf("can not get the binance price of coin %s", token.BinName)
		}
		if len(avgPrices) > 0 {
			token.AvgPrice = avg(avgPrices)
			token.AvgInd = 1
		} else {
			token.AvgInd = 0
		}
	}
	return nil
}

func avg(data []uint64) uint64 {
	sum := uint64(0)
	for _, item := range data {
		sum += item
	}
	return sum / uint64(len(data))
}


func (this *CoinPriceListen) getCmcCoinPrice(coins []string) (map[string]float64, error) {
	var cmcSdk *coinmarketcap.CoinMarketCapSdk
	if this.cmcCfg == nil || len(this.cmcCfg.RestURL) == 0 {
		cmcSdk = coinmarketcap.DefaultCoinMarketCapSdk()
	} else {
		cmcSdk = coinmarketcap.NewCoinMarketCapSdk(this.cmcCfg.RestURL, this.cmcCfg.Key)
	}
	listings, err := cmcSdk.ListingsLatest()
	if err != nil {
		return nil, err
	}
	//
	coinName2Id := make(map[string]string, 0)
	for _, listing := range listings {
		coinName2Id[listing.Name] = fmt.Sprintf("%d", listing.ID)
	}
	//
	coinIds := make([]string, 0)
	for _, coin := range coins {
		coinId, ok := coinName2Id[coin]
		if !ok {
			log.Warnf("There is no coin %s in CoinMarketCap!", coin)
			continue
		}
		coinIds = append(coinIds, coinId)
	}
	//
	requestCoinIds := strings.Join(coinIds, ",")
	quotes, err := cmcSdk.QuotesLatest(requestCoinIds)
	if err != nil {
		return nil, err
	}
	//
	coinName2Price := make(map[string]float64)
	for _, v := range quotes {
		name := v.Name
		if v.Quote == nil || v.Quote["USD"] == nil {
			log.Warnf(" There is no price for coin %s in CoinMarketCap!", name)
			continue
		}
		coinName2Price[name] = v.Quote["USD"].Price
	}
	return coinName2Price, nil
}

func (this *CoinPriceListen) getBinancePrice(coins []string) (map[string]float64, error) {
	var binSdk *binance.BinanceSdk
	if this.binCfg == nil || len(this.binCfg.RestURL) == 0 {
		binSdk = binance.DefaultBinanceSdk()
	} else {
		binSdk = binance.NewBinanceSdk(this.binCfg.RestURL)
	}
	quotes, err := binSdk.QuotesLatest()
	if err != nil {
		return nil, err
	}
	coinSymbol2Price := make(map[string]float64, 0)
	for _, v := range quotes {
		coinSymbol2Price[v.Symbol] = v.Price
	}
	coinPrice := make(map[string]float64, 0)
	for _, coin := range coins {
		price, ok := coinSymbol2Price[coin]
		if !ok {
			log.Warnf("There is no coin price %s in Binance!", coin)
			continue
		}
		coinPrice[coin] = price
	}
	return coinPrice, nil
}

