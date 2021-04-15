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
	"github.com/astaxie/beego"
	"poly-bridge/nft_http/controllers"
)

func init() {
	ns := beego.NewNamespace("/nft/v1",
		beego.NSRouter("/", &controllers.InfoController{}, "*:Get"),

		beego.NSRouter("/assetshow/", &controllers.InfoController{}, "post:Home"),
		beego.NSRouter("/asset/", &controllers.AssetController{}, "post:Asset"),
		beego.NSRouter("/assets/", &controllers.AssetController{}, "post:Assets"),

		//beego.NSRouter("/assetbasics/", &controllers.AssetController{}, "post:AssetBasics"),
		//beego.NSRouter("/assetmap/", &controllers.AssetMapController{}, "post:AssetMap"),
		//beego.NSRouter("/assetmapreverse/", &controllers.AssetMapController{}, "post:AssetMapReverse"),

		beego.NSRouter("/items/", &controllers.ItemController{}, "post:Items"),
		beego.NSRouter("/getfee/", &controllers.FeeController{}, "post:GetFee"),

		beego.NSRouter("/exp_rtransactions/", &controllers.ExplorerController{}, "post:Transactions"),
		beego.NSRouter("/exp_transactionsofaddress/", &controllers.ExplorerController{}, "post:TransactionsOfAddress"),
		beego.NSRouter("/exp_transactionofhash/", &controllers.ExplorerController{}, "post:TransactionDetail"),

		beego.NSRouter("/transactionsofaddress/", &controllers.TransactionController{}, "post:TransactionsOfAddress"),
		beego.NSRouter("/transactionofhash/", &controllers.TransactionController{}, "post:TransactionOfHash"),

		//beego.NSRouter("/transactionsofstate/", &controllers.TransactionController{}, "post:TransactionsOfState"),
	)
	beego.AddNamespace(ns)
	beego.Router("/", &controllers.InfoController{}, "*:Get")
}
