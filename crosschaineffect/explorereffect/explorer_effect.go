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

package explorereffect

import (
	"github.com/astaxie/beego/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/crosschaindao/explorerdao"
	"time"
)

type ExplorerEffect struct {
	dbCfg  *conf.DBConfig
	effCfg *conf.EventEffectConfig
	db     *gorm.DB
	chains []*explorerdao.Chain
	time   int64
}

func NewExplorerEffect(effCfg *conf.EventEffectConfig, dbCfg *conf.DBConfig) *ExplorerEffect {
	explorerEff := &ExplorerEffect{
		dbCfg:  dbCfg,
		effCfg: effCfg,
		chains: nil,
		time:   0,
	}
	Logger := logger.Default
	if dbCfg.Debug == true {
		Logger = Logger.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{Logger: Logger})
	if err != nil {
		panic(err)
	}
	explorerEff.db = db
	chains := make([]*explorerdao.Chain, 0)
	res := db.Model(&explorerdao.Chain{}).Find(&chains)
	if res.Error != nil || res.RowsAffected == 0 {
		panic(err)
	}
	explorerEff.chains = chains
	explorerEff.time = time.Now().Unix()
	return explorerEff
}

func (eff *ExplorerEffect) Effect() error {
	err := eff.updateHash()
	if err != nil {
		logs.Error("update hash- err: %s", err)
	}
	err = eff.checkChainListening()
	if err != nil {
		logs.Error("check chain listening- err: %s", err)
	}
	return nil
}

func (eff *ExplorerEffect) Name() string {
	return basedef.SERVER_EXPLORER
}

func (eff *ExplorerEffect) GetEffectSlot() int64 {
	return eff.effCfg.EffectSlot
}

func (eff *ExplorerEffect) updateHash() error {
	polySrcRelations := make([]*explorerdao.PolySrcRelation, 0)
	eff.db.Table("mchain_tx").Where("left(mchain_tx.ftxhash, 8) = ? and mchain_tx.fchain != ?", "00000000", basedef.ETHEREUM_CROSSCHAIN_ID).Select("mchain_tx.txhash as poly_hash, fchain_tx.txhash as src_hash").Joins("inner join fchain_tx on mchain_tx.ftxhash = fchain_tx.xkey and mchain_tx.fchain = fchain_tx.chain_id").Preload("SrcTransaction").Preload("PolyTransaction").Find(&polySrcRelations)
	updatePolyTransactions := make([]*explorerdao.PolyTransaction, 0)
	for _, polySrcRelation := range polySrcRelations {
		if polySrcRelation.SrcTransaction != nil {
			polySrcRelation.PolyTransaction.SrcHash = polySrcRelation.SrcHash
			updatePolyTransactions = append(updatePolyTransactions, polySrcRelation.PolyTransaction)
		}
	}
	if len(updatePolyTransactions) > 0 {
		eff.db.Save(updatePolyTransactions)
	}
	return nil
}

func (eff *ExplorerEffect) checkChainListening() error {
	slot := eff.effCfg.ChainListening
	if slot == 0 {
		slot = 300
	}
	old := eff.time / slot
	now := time.Now().Unix()
	new := now / slot
	if new == old {
		return nil
	}
	id2Chains := make(map[uint64]*explorerdao.Chain)
	for _, chain := range eff.chains {
		id2Chains[chain.ChainId] = chain
	}
	chains := make([]*explorerdao.Chain, 0)
	eff.db.Model(&explorerdao.Chain{}).Find(&chains)
	for _, chain := range chains {
		old, ok := id2Chains[chain.ChainId]
		if !ok {
			continue
		}
		if chain.Height == old.Height {
			logs.Error("Chain %d is not listening!", chain.ChainId)
		}
	}
	eff.chains = chains
	eff.time = now
	return nil
}
