package gateio

import (
	"context"
	"github.com/antihax/optional"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gateio/gateapi-go/v6"
	"poly-bridge/basedef"
	"poly-bridge/models"
)

type GateioSdk struct {
	client *gateapi.APIClient
}

func NewGateioSdk() *GateioSdk {
	return &GateioSdk{
		gateapi.NewAPIClient(gateapi.NewConfiguration()),
	}
}

func (g *GateioSdk) GetMarketName() string {
	return basedef.MARKET_GATEIO
}

func (g *GateioSdk) GetCoinPriceAndRank(coins []models.NameAndmarketId) (map[string]float64, map[string]int, error) {
	coinPrice := make(map[string]float64, 0)
	coinRank := make(map[string]int, 0)
	for _, coin := range coins {
		tickers, _, err := g.client.SpotApi.ListTickers(context.Background(), &gateapi.ListTickersOpts{optional.NewString(coin.PriceMarketName)})
		if err != nil {
			if e, ok := err.(gateapi.GateAPIError); ok {
				logs.Error("[token: %s] gate api error: %s", coin.PriceMarketName, e.Error())
				continue
			} else {
				logs.Error("[token: %s] generic error: %s", coin.PriceMarketName, err.Error())
				continue
			}
		}
		if len(tickers) != 1 {
			logs.Error("[token: %s] tickers length is not 1", coin.PriceMarketName, err.Error())
			continue
		}
		lastPriceString := tickers[0].Last
		if lastPriceString == "" {
			logs.Error("[token: %s] lastPriceString is empty", coin.PriceMarketName, err.Error())
		}
		coinPrice[coin.PriceMarketName] = basedef.String2Float64(lastPriceString)
		coinRank[coin.PriceMarketName] = 0
	}
	logs.Info("gateio coin price: %+v", coinPrice)
	return coinPrice, coinRank, nil
}
