package crosschainstats

import (
	"github.com/beego/beego/v2/core/logs"
	"poly-bridge/http/tools"
	"time"
)

func (this *Stats) censusTimeLines() (err error) {
	logs.Info("start censusTimeLines")
	allChain, err := this.dao.GetChains()
	if err != nil {
		return err
	}
	chains := make(map[uint64]string, 0)
	for _, chain := range allChain {
		chains[chain.ChainId] = chain.Name
	}
	timeNow := time.Now().Unix()
	timeLast := timeNow - 60
	chainAvgTimes, err := this.dao.GetAvgTimeSrc2Poly(timeLast, timeNow)
	if chainAvgTimes != nil {
		for _, srcAvgTime := range chainAvgTimes {
			if name, ok := chains[srcAvgTime.ChainId]; ok {
				tools.Record(srcAvgTime.AvgTime, "avg_src_to_poly.%s", name)
			}
		}
	}
	chainAvgTimes, err = this.dao.GetAvgTimePoly2Dst(timeLast, timeNow)
	if chainAvgTimes != nil {
		for _, polyAvgTime := range chainAvgTimes {
			if name, ok := chains[polyAvgTime.ChainId]; ok {
				tools.Record(polyAvgTime.AvgTime, "avg_poly_to_dst.%s", name)
			}
		}
	}
	return nil
}
