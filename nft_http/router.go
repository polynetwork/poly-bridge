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

package nft_http

import (
	"poly-bridge/conf"
	"poly-bridge/nft_http/controllers"

	"github.com/beego/beego/v2/server/web"
)

func Init(config *conf.Config) web.LinkNamespace {
	controllers.SetBaseInfo(config.HttpConfig.Address, config.HttpConfig.Port)
	controllers.Initialize(config)

	ns := web.NSNamespace("/nft",
		web.NSRouter("/", &controllers.InfoController{}, "*:Get"),

		web.NSRouter("/assetshow/", &controllers.InfoController{}, "post:Home"),
		web.NSRouter("/asset/", &controllers.AssetController{}, "post:Asset"),
		web.NSRouter("/assets/", &controllers.AssetController{}, "post:Assets"),

		//web.NSRouter("/assetbasics/", &controllers.AssetController{}, "post:AssetBasics"),
		//web.NSRouter("/assetmap/", &controllers.AssetMapController{}, "post:AssetMap"),
		//web.NSRouter("/assetmapreverse/", &controllers.AssetMapController{}, "post:AssetMapReverse"),

		web.NSRouter("/items/", &controllers.ItemController{}, "post:Items"),

		web.NSRouter("/exp_transactions/", &controllers.ExplorerController{}, "post:Transactions"),
		web.NSRouter("/exp_transactionsofaddress/", &controllers.ExplorerController{}, "post:TransactionsOfAddress"),
		web.NSRouter("/exp_transactionofhash/", &controllers.ExplorerController{}, "post:TransactionDetail"),

		web.NSRouter("/transactionsofaddress/", &controllers.TransactionController{}, "post:TransactionsOfAddress"),
		web.NSRouter("/transactionofhash/", &controllers.TransactionController{}, "post:TransactionOfHash"),

		//web.NSRouter("/transactionsofstate/", &controllers.TransactionController{}, "post:TransactionsOfState"),
	)
	return ns
}
