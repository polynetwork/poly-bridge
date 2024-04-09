package pltelfrate

import (
	"fmt"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/models"
	"testing"
)

func Test_GetELFPrice(t *testing.T) {
	cfg := new(conf.CoinPriceListenConfig)
	cfg.MarketName = "elfrate"
	cfg.Nodes = make([]*conf.Restful, 0)
	cfg.Nodes = append(cfg.Nodes, &conf.Restful{Url: "https://api.pltplace.io/api/v1/elf_rate"})
	sdk := NewElfRateSdk(cfg)
	coinPrices, coinRanks, err := sdk.GetCoinPriceAndRank([]models.NameAndmarketId{{"ELF", 0}})
	if err != nil {
		panic(err)
	}
	fmt.Println(coinPrices, coinRanks)
	price, _ := new(big.Float).Mul(big.NewFloat(coinPrices["ELF"]), big.NewFloat(float64(basedef.PRICE_PRECISION))).Int64()
	fmt.Println(price)
}
