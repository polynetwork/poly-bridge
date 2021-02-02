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
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
	_ "poly-bridge/routers"
)

func main() {
	logs.SetLogger(logs.AdapterFile, `{"filename":"logs/bridge_http.log"}`)
	mode := beego.AppConfig.String("runmode")
	if mode == "dev" {
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
		AllowCredentials: true}))
	beego.Run()
}
