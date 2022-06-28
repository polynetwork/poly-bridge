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

package http

import (
	"github.com/beego/beego/v2/server/web"
	"poly-bridge/conf"
)

func GetRouter(config *conf.Config) web.LinkNamespace {
	SetCoinRankFilterInfo(config.RiskyCoinHandleConfig)
	SetProxyFeeRatioMap(config)
	ns := web.NSNamespace("/bridge",
		web.NSRouter("/", &InfoController{}, "*:Get"),
		web.NSRouter("/token/", &TokenController{}, "post:Token"),
		web.NSRouter("/tokens/", &TokenController{}, "post:Tokens"),
		web.NSRouter("/tokenbasics/", &TokenController{}, "post:TokenBasics"),
		web.NSRouter("/tokenbasicsinfo/", &TokenController{}, "post:TokenBasicsInfo"),
		web.NSRouter("/tokenmap/", &TokenMapController{}, "post:TokenMap"),
		web.NSRouter("/tokenmapreverse/", &TokenMapController{}, "post:TokenMapReverse"),
		web.NSRouter("/getfee/", &FeeController{}, "post:GetFee"),
		web.NSRouter("/oldgetfee/", &FeeController{}, "post:OldGetFee"),
		web.NSRouter("/checkfee/", &FeeController{}, "post:CheckFee"),
		web.NSRouter("/newcheckfee/", &FeeController{}, "post:NewCheckFee"),
		web.NSRouter("/checkswapfee/", &FeeController{}, "post:CheckSwapFee"),
		web.NSRouter("/transactions/", &TransactionController{}, "post:Transactions"),
		web.NSRouter("/transactionswithfilter/", &TransactionController{}, "post:TransactionsOfAddressWithFilter"),
		web.NSRouter("/transactionsofaddress/", &TransactionController{}, "post:TransactionsOfAddress"),
		web.NSRouter("/transactionofhash/", &TransactionController{}, "post:TransactionOfHash"),
		web.NSRouter("/transactionofcurve/", &TransactionController{}, "post:TransactionOfCurve"),
		web.NSRouter("/transactionsofstate/", &TransactionController{}, "post:TransactionsOfState"),
		web.NSRouter("/transactionsofunfinished/", &TransactionController{}, "post:TransactionsOfUnfinished"),
		web.NSRouter("/transactionsofasset/", &TransactionController{}, "post:TransactionsOfAsset"),
		web.NSRouter("/expecttime/", &StatisticController{}, "post:ExpectTime"),
		web.NSRouter("/gettokenasset/", &TokenAssetController{}, "post:Gettokenasset"),
		web.NSRouter("/getmanualtxdata/", &TransactionController{}, "post:GetManualTxData"),
		web.NSRouter("/chainhealth/", &ChainHealthController{}, "post:Health"),
		web.NSRouter("/wrappercheck/", &WrapperController{}, "post:WrapperCheck"),
	)
	return ns
}
