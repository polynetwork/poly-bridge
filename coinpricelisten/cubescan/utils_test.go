package cubescan

import (
	"fmt"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/models"
	"testing"
)

func Test_GetcubePrice(t *testing.T) {
	sdk := NewCubescanSdk()
	coinPrices, coinRanks, err := sdk.GetCoinPriceAndRank([]models.NameAndmarketId{{"cube", 0}})
	if err != nil {
		panic(err)
	}
	fmt.Println(coinPrices, coinRanks)
	price, _ := new(big.Float).Mul(big.NewFloat(coinPrices["cube"]), big.NewFloat(float64(basedef.PRICE_PRECISION))).Int64()
	fmt.Println(price)
}
