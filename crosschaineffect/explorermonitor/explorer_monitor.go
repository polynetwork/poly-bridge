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

package explorermonitor

import (
	"github.com/astaxie/beego/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"poly-swap/conf"
	"poly-swap/crosschaindao/explorerdao"
)

type ExplorerMonitor struct {
	dbCfg  *conf.DBConfig
	monCfg *conf.CrossChainMonitorConfig
	db     *gorm.DB
}

func NewExplorerMonitor(monCfg *conf.CrossChainMonitorConfig, dbCfg *conf.DBConfig) *ExplorerMonitor {
	explorerMonitor := &ExplorerMonitor{
		dbCfg:  dbCfg,
		monCfg: monCfg,
	}
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	explorerMonitor.db = db
	return explorerMonitor
}

func (monitor *ExplorerMonitor) Monitor() error {
	err := monitor.updateHash()
	if err != nil {
		logs.Error("update hash- err: %s", err)
	}
	return nil
}

func (monitor *ExplorerMonitor) updateHash() error {
	polySrcRelations := make([]*explorerdao.PolySrcRelation, 0)
	monitor.db.Debug().Table("mchain_tx").Where("left(mchain_tx.ftxhash, 8) = ? and mchain_tx.fchain != ?", "00000000", conf.ETHEREUM_CROSSCHAIN_ID).Select("mchain_tx.txhash as poly_hash, fchain_tx.txhash as src_hash").Joins("inner join fchain_tx on mchain_tx.ftxhash = fchain_tx.xkey").Preload("SrcTransaction").Preload("PolyTransaction").Find(&polySrcRelations)
	updatePolyTransactions := make([]*explorerdao.PolyTransaction, 0)
	for _, polySrcRelation := range polySrcRelations {
		if polySrcRelation.SrcTransaction != nil {
			polySrcRelation.PolyTransaction.SrcHash = polySrcRelation.SrcHash
			updatePolyTransactions = append(updatePolyTransactions, polySrcRelation.PolyTransaction)
		}
	}
	if len(updatePolyTransactions) > 0 {
		monitor.db.Save(updatePolyTransactions)
	}
	return nil
}
