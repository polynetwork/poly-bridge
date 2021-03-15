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
	"poly-bridge/bridge_tools/conf"
	"poly-bridge/crosschaindao"
)

func startTransactions(cfg *conf.TransactionsConfig) {
	dao := crosschaindao.NewCrossChainDao(cfg.Server, cfg.Backup, cfg.DBConfig)
	if dao == nil {
		panic("server is invalid")
	}
	//
	err := dao.UpdateEvents(nil, cfg.WrapperTransactions, cfg.SrcTransactions, cfg.PolyTransactions, cfg.DstTransactions)
	if err != nil {
		panic(err)
	}
	err = dao.RemoveEvents(cfg.RemoveTransactions, cfg.RemoveTransactions, cfg.RemoveTransactions)
	if err != nil {
		panic(err)
	}
}
