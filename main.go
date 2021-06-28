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

	"poly-bridge/common"
	"poly-bridge/conf"
	"poly-bridge/controllers"
	"poly-bridge/nft_http"
	_ "poly-bridge/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/urfave/cli"
)

func main() {
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
		panic("Invalid server config")
	}
	logs.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename":"%s"}`, config.LogFile))

	controllers.Init()
	common.SetupChainsSDK(config)
	// NFT http
	nft_http.Init(config)

	// Insert beego config
	beego.BConfig.Listen.HTTPAddr = config.HttpConfig.Address
	beego.BConfig.Listen.HTTPPort = config.HttpConfig.Port
	beego.BConfig.RunMode = config.RunMode
	beego.BConfig.AppName = "bridgehttp"
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.EnableErrorsRender = false

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
		beego.InsertFilter("/*", beego.FinishRouter, FilterLog, false)
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true}),
	)

	beego.Run()
}
