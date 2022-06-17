package huobi

import (
	"fmt"
	"poly-bridge/models"
	"testing"
)

func TestHuobiSdk_GetCoinPrice(t *testing.T) {

	sdk := NewHuobiSdk()
	cube := models.NameAndmarketId{
		"CUBE", 0,
	}
	coins := []models.NameAndmarketId{
		cube,
	}
	fmt.Println(sdk.GetCoinPriceAndRank(coins))

}
