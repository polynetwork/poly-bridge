package crosschainstats

import (
	"github.com/beego/beego/v2/core/logs"
	"math/big"
	"poly-bridge/http/tools"
)

func (this *Stats) censusAssetLines() (err error) {
	logs.Info("start censusAssetLines")
	sourceTokenStatistics, err := this.dao.GetSourceTokenStatistics()
	if err != nil {
		logs.Error("GetSourceTokenStatistics err", err)
		return err
	}
	for _, sourceTokenStatistic := range sourceTokenStatistics {
		if sourceTokenStatistic.InAmount.Cmp(big.NewInt(0)) >= 0 {
			tools.Record(new(big.Int).Div(&sourceTokenStatistic.InAmount.Int, big.NewInt(100)), "locked_asset_InAmount.%s", sourceTokenStatistic.TokenBasicName)
			tools.Record(new(big.Int).Div(&sourceTokenStatistic.InAmountUsd.Int, big.NewInt(10000)), "locked_asset_InAmountUsd.%s", sourceTokenStatistic.TokenBasicName)
		}
	}
	return nil
}
