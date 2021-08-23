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

package coincheck

type JPY struct {
	JPY  float64 `json:"jpy,string"`
	USD  float64 `json:"usd,string"`
	BTC  float64 `json:"btc,string"`
	ETH  float64 `json:"eth,string"`
	ETC  float64 `json:"etc,string"`
	LSK  float64 `json:"lsk,string"`
	FCT  float64 `json:"fct,string"`
	XRP  float64 `json:"xrp,string"`
	XEM  float64 `json:"xem,string"`
	LTC  float64 `json:"ltc,string"`
	BCH  float64 `json:"bch,string"`
	MONA float64 `json:"mona,string"`
	XLM  float64 `json:"xlm,string"`
	QTUM float64 `json:"qtum,string"`
	BAT  float64 `json:"bat,string"`
	IOST float64 `json:"iost,string"`
	ENJ  float64 `json:"enj,string"`
	OMG  float64 `json:"omg,string"`
	PLT  float64 `json:"plt,string"`
}

type BTC struct {
	BTC  float64 `json:"btc,string"`
	ETH  float64 `json:"eth,string"`
	ETC  float64 `json:"etc,string"`
	LSK  float64 `json:"lsk,string"`
	FCT  float64 `json:"fct,string"`
	XRP  float64 `json:"xrp,string"`
	XEM  float64 `json:"xem,string"`
	LTC  float64 `json:"ltc,string"`
	BCH  float64 `json:"bch,string"`
	MONA float64 `json:"mona,string"`
	XLM  float64 `json:"xlm,string"`
	QTUM float64 `json:"qtum,string"`
	BAT  float64 `json:"bat,string"`
	IOST float64 `json:"iost,string"`
	ENJ  float64 `json:"enj,string"`
	OMG  float64 `json:"omg,string"`
	PLT  float64 `json:"plt,string"`
}

type Rate struct {
	Jpy JPY `json:"jpy"`
	Btc BTC `json:"btc"`
}
