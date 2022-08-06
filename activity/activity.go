package activity

import (
	"context"
	"github.com/beego/beego/v2/core/logs"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/crosschaindao/bridgedao"
	"sync"
	"time"
)

type ActivityStats struct {
	context.Context
	cancel context.CancelFunc
	cfg    *conf.ActivityConfig
	dao    *bridgedao.BridgeDao
	wg     sync.WaitGroup
}

var ccs *ActivityStats

func StartActivity(server string, cfg *conf.ActivityConfig, dbCfg *conf.DBConfig) {
	if server != basedef.SERVER_POLY_BRIDGE {
		panic("StartStartActivity Only runs on bridge server")
	}
	if cfg == nil || dbCfg == nil {
		panic("Invalid Activity config")
	}
	dao := bridgedao.NewBridgeDao(dbCfg, false)
	ctx, cancel := context.WithCancel(context.Background())
	ccs = &ActivityStats{dao: dao, cfg: cfg, Context: ctx, cancel: cancel}
	ccs.Start()
}

func StopActivity() {
	if ccs != nil {
		ccs.Stop()
	}
}

func (this *ActivityStats) run(interval int64, f func() error) {
	this.wg.Add(1)
	ticker := time.NewTicker(time.Second * time.Duration(interval))
	for {
		select {
		case <-ticker.C:
			err := f()
			if err != nil {
				logs.Error("stats run error%s", err)
			}
		case <-this.Done():
			break
		}
	}
	this.wg.Done()
}

func (this *ActivityStats) Start() {
	logs.Info("start ActivityStats")
	if this.cfg.AirDropStartTime <= 0 || this.cfg.AirDropEndTime <= 0 ||
		this.cfg.TokenPriceStartTime <= 0 || this.cfg.TokenPriceEndTime <= 0 ||
		this.cfg.TokenPriceAvgInterval <= 0 || this.cfg.AirDropInfoInterval <= 0 {
		logs.Error("ActivityConfig config err")
		return
	}
	go this.StartAirDrop()
	go this.StartTokenPrice()
	logs.Info("ActivityStats ...")

}

func (this *ActivityStats) Stop() {
	logs.Info("Stopping ActivityStats")
	this.cancel()
	this.wg.Wait()
}
