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
	"poly-bridge/basedef"
	"poly-bridge/cacheRedis"
	"poly-bridge/common"
	"poly-bridge/conf"
	"poly-bridge/explorer"
	"poly-bridge/http"
	"poly-bridge/nft_http"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"github.com/urfave/cli"
	//http2 "net/http"
	_ "net/http/pprof"
)

func main() {
	//go func() {
	//	http2.ListenAndServe("0.0.0.0:3344", nil)
	//}()
	if err := setupApp().Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func setupApp() *cli.App {
	app := cli.NewApp()
	app.Name = "poly bridge server"
	app.Usage = "poly-bridge http server"
	app.Action = run
	app.Version = "1.0.0"
	app.Copyright = "Copyright in 2019 The PolyNetwork Authors"
	app.Flags = []cli.Flag{
		conf.ConfigPathFlag,
	}
	return app
}

func run(ctx *cli.Context) {
	// Initialize
	configFile := ctx.GlobalString("config")
	config := conf.NewConfig(configFile)
	if config == nil || config.HttpConfig == nil {
		//fmt.Println(config)
		//fmt.Println(config.HttpConfig)
		panic("Invalid server config")
	}
	logs.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename":"%s"}`, config.LogFile))

	basedef.ConfirmEnv(config.Env)
	common.SetupChainsSDK(config)

	web.InsertFilter("*", web.BeforeRouter, cors.Allow(
		&cors.Options{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
			AllowCredentials: true,
		},
	))

	// bridge http
	http.Init()
	// explorer http
	explorer.Init()
	// redis
	cacheRedis.Init()

	// register http routers
	web.AddNamespace(
		web.NewNamespace("/v1",
			nft_http.Init(config),
			http.GetRouter(config),
			explorer.GetRouter(),
		),
	)

	// Insert web config
	web.BConfig.Listen.HTTPAddr = config.HttpConfig.Address
	web.BConfig.Listen.HTTPPort = config.HttpConfig.Port
	web.BConfig.RunMode = config.RunMode
	web.BConfig.AppName = "bridgehttp"
	web.BConfig.CopyRequestBody = true
	web.BConfig.EnableErrorsRender = false

	if config.RunMode == "dev" {
		var FilterLog = func(ctx *context.Context) {
			url, _ := json.Marshal(ctx.Input.Data()["RouterPattern"])
			params := string(ctx.Input.RequestBody)
			outputBytes, _ := json.Marshal(ctx.Input.Data()["json"])
			divider := " - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -"
			topDivider := "┌" + divider
			middleDivider := "├" + divider
			bottomDivider := "└" + divider
			outputStr := "\n" + topDivider + "\n│ url:" + string(url) + "\n" + middleDivider + "\n│ request:" + string(params) + "\n│ response:" + string(outputBytes) + "\n" + bottomDivider
			logs.Info(outputStr)
		}
		web.InsertFilter("/*", web.FinishRouter, FilterLog)
	}

	web.Run()
}
