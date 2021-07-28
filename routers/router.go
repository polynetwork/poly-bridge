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

package routers

import (
	"poly-bridge/conf"
	"poly-bridge/controllers"

	"github.com/astaxie/beego"
)

func init() {

	configFile := beego.AppConfig.String("chain_config")
	config := conf.NewConfig(configFile)
	if config == nil {
		panic("startServer - read config failed!")
	}

	// run checks
	bot := &controllers.BotController{Conf: config}
	go bot.RunChecks()

	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/", &controllers.InfoController{}, "*:Get"),
		beego.NSRouter("/token/", &controllers.TokenController{}, "post:Token"),
		beego.NSRouter("/tokens/", &controllers.TokenController{}, "post:Tokens"),
		beego.NSRouter("/tokenbasics/", &controllers.TokenController{}, "post:TokenBasics"),
		beego.NSRouter("/tokenbasicsinfo/", &controllers.TokenController{}, "post:TokenBasicsInfo"),
		beego.NSRouter("/tokenmap/", &controllers.TokenMapController{}, "post:TokenMap"),
		beego.NSRouter("/tokenmapreverse/", &controllers.TokenMapController{}, "post:TokenMapReverse"),
		beego.NSRouter("/getfee/", &controllers.FeeController{}, "post:GetFee"),
		beego.NSRouter("/checkfee/", &controllers.FeeController{}, "post:CheckFee"),
		beego.NSRouter("/checkswapfee/", &controllers.FeeController{}, "post:CheckSwapFee"),
		beego.NSRouter("/transactions/", &controllers.TransactionController{}, "post:Transactions"),
		beego.NSRouter("/transactionswithfilter/", &controllers.TransactionController{}, "post:TransactionsOfAddressWithFilter"),
		beego.NSRouter("/transactionsofaddress/", &controllers.TransactionController{}, "post:TransactionsOfAddress"),
		beego.NSRouter("/transactionofhash/", &controllers.TransactionController{}, "post:TransactionOfHash"),
		beego.NSRouter("/transactionofcurve/", &controllers.TransactionController{}, "post:TransactionOfCurve"),
		beego.NSRouter("/transactionsofstate/", &controllers.TransactionController{}, "post:TransactionsOfState"),
		beego.NSRouter("/transactionsofunfinished/", &controllers.TransactionController{}, "post:TransactionsOfUnfinished"),
		beego.NSRouter("/transactionsofasset/", &controllers.TransactionController{}, "post:TransactionsOfAsset"),
		beego.NSRouter("/bot/", &controllers.BotController{Conf: config}, "get:BotPage"),
		beego.NSRouter("/bottxs/", &controllers.BotController{Conf: config}, "get:GetTxs"),
		beego.NSRouter("/botcheck/", &controllers.BotController{Conf: config}, "get:CheckTxs"),
		beego.NSRouter("/botcheckfee/", &controllers.BotController{Conf: config}, "post:CheckFees"),
		beego.NSRouter("/botfinishtx/", &controllers.BotController{Conf: config}, "get:FinishTx"),
		beego.NSRouter("/expecttime/", &controllers.StatisticController{}, "post:ExpectTime"),
	)
	beego.AddNamespace(ns)
	beego.Router("/", &controllers.InfoController{}, "*:Get")
}
