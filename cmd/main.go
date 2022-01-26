/*
 * Copyright (C) 2020 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"poly-bridge/cacheRedis"
	"runtime"
	"syscall"

	"poly-bridge/basedef"
	"poly-bridge/chainfeelisten"
	"poly-bridge/coinpricelisten"
	"poly-bridge/common"
	"poly-bridge/conf"
	"poly-bridge/crosschaineffect"
	"poly-bridge/crosschainlisten"
	"poly-bridge/crosschainstats"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/urfave/cli"
)

func setupApp() *cli.App {
	app := cli.NewApp()
	app.Usage = "poly-bridge Service"
	app.Action = StartServer
	app.Version = "1.0.0"
	app.Copyright = "Copyright in 2019 The Ontology Authors"
	app.Flags = []cli.Flag{
		conf.ConfigPathFlag,
	}
	app.Commands = []cli.Command{}
	app.Before = func(context *cli.Context) error {
		runtime.GOMAXPROCS(runtime.NumCPU())
		return nil
	}
	return app
}

func StartServer(ctx *cli.Context) {
	for true {
		startServer(ctx)
		sig := waitSignal()
		stopServer()
		if sig != syscall.SIGHUP {
			break
		} else {
			continue
		}
	}
}

func startServer(ctx *cli.Context) {
	configFile := ctx.GlobalString("config")
	config := conf.NewConfig(configFile)
	if config == nil {
		logs.Error("startServer - read config failed!")
		return
	}
	logs.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename":"%s"}`, config.LogFile))

	{
		conf, _ := json.Marshal(config)
		logs.Info("%s\n", string(conf))
	}
	//initialize redis
	cacheRedis.Init()

	//metrics.Init("bridge")
	basedef.ConfirmEnv(config.Env)
	common.SetupChainsSDK(config)
	if config.Backup {
		crosschainlisten.StartCrossChainListen(config)
		crosschainlisten.StartCrossChainListenPatch(config)
		return
	}
	crosschainlisten.StartCrossChainListen(config)
	//coinpricelisten.StartCoinPriceListen(config.Server, config.CoinPriceUpdateSlot, config.CoinPriceListenConfig, config.DBConfig)
	//chainfeelisten.StartFeeListen(config.Server, config.FeeUpdateSlot, config.FeeListenConfig, config.DBConfig)
	crosschaineffect.StartCrossChainEffect(config.Server, config.EventEffectConfig, config.DBConfig, config.RedisConfig)
	//crosschainstats.StartCrossChainStats(config.Server, config.StatsConfig, config.DBConfig, config.IPPortConfig, config.ChainListenConfig)

	metricConfig := config.MetricConfig
	if metricConfig == nil {
		metricConfig = &conf.HttpConfig{
			Address: "0.0.0.0",
			Port:    6222,
		}
	}

	// Insert web config
	web.BConfig.Listen.HTTPAddr = metricConfig.Address
	web.BConfig.Listen.HTTPPort = metricConfig.Port
	web.BConfig.RunMode = config.RunMode
	web.BConfig.AppName = "bridge-server"
	web.BConfig.CopyRequestBody = true
	web.BConfig.EnableErrorsRender = false
	go web.Run()
}

func waitSignal() os.Signal {
	exit := make(chan os.Signal, 0)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer signal.Stop(sc)
	go func() {
		for sig := range sc {
			logs.Info("cross chain listen received signal:(%s).", sig.String())
			exit <- sig
			close(exit)
			break
		}
	}()
	sig := <-exit
	return sig
}

func stopServer() {
	crosschainlisten.StopCrossChainListen()
	coinpricelisten.StopCoinPriceListen()
	chainfeelisten.StopFeeListen()
	crosschaineffect.StopCrossChainEffect()
	crosschainstats.StopCrossChainStats()
}

func main() {
	if err := setupApp().Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
