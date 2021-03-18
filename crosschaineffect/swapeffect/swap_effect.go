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

package swapeffect

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/models"
	"time"
)

type SwapEffect struct {
	dbCfg  *conf.DBConfig
	cfg    *conf.EventEffectConfig
	db     *gorm.DB
	chains []*models.Chain
	time   int64
}

func NewSwapEffect(cfg *conf.EventEffectConfig, dbCfg *conf.DBConfig) *SwapEffect {
	swapEffect := &SwapEffect{
		dbCfg:  dbCfg,
		cfg:    cfg,
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
	swapEffect.db = db
	chains := make([]*models.Chain, 0)
	res := db.Model(&models.Chain{}).Find(&chains)
	if res.Error != nil || res.RowsAffected == 0 {
		panic(err)
	}
	swapEffect.chains = chains
	swapEffect.time = time.Now().Unix()
	return swapEffect
}

func (eff *SwapEffect) Effect() error {
	err := eff.updateHash()
	if err != nil {
		logs.Error("update hash- err: %s", err)
	}
	err = eff.checkStatus()
	if err != nil {
		logs.Error("check status- err: %s", err)
	}
	err = eff.updateStatus()
	if err != nil {
		logs.Error("update status- err: %s", err)
	}
	err = eff.checkChainListening()
	if err != nil {
		logs.Error("check chain listening- err: %s", err)
	}
	return nil
}
func (eff *SwapEffect) Name() string {
	return basedef.SERVER_POLY_SWAP
}

func (eff *SwapEffect) GetEffectSlot() int64 {
	return eff.cfg.EffectSlot
}

func (eff *SwapEffect) updateHash() error {
	polySrcRelations := make([]*models.PolySrcRelation, 0)
	eff.db.Table("poly_transactions").Where("left(poly_transactions.src_hash, 8) = ?", "00000000").Select("poly_transactions.hash as poly_hash, src_transactions.hash as src_hash").Joins("inner join src_transactions on poly_transactions.src_hash = src_transactions.key and poly_transactions.src_chain_id = src_transactions.chain_id").Preload("SrcTransaction").Preload("PolyTransaction").Find(&polySrcRelations)
	updatePolyTransactions := make([]*models.PolyTransaction, 0)
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

func (eff *SwapEffect) checkStatus() error {
	wrapperTransactions := make([]*models.WrapperTransaction, 0)
	now := time.Now().Unix() - eff.cfg.HowOld
	eff.db.Model(models.WrapperTransaction{}).Where("status != ? and time < ?", basedef.STATE_FINISHED, now).Find(&wrapperTransactions)
	if len(wrapperTransactions) > 0 {
		wrapperTransactionsJson, _ := json.Marshal(wrapperTransactions)
		logs.Error("There is unfinished transactions %s", string(wrapperTransactionsJson))
	}
	return nil
}

func (eff *SwapEffect) updateStatus() error {
	chains := make([]*models.Chain, 0)
	id2Chains := make(map[uint64]*models.Chain)
	eff.db.Model(&models.Chain{}).Find(&chains)
	for _, chain := range chains {
		id2Chains[*chain.ChainId] = chain
	}
	wrapperPolyDstRelations := make([]*models.SrcPolyDstRelation, 0)
	wrapperTransactions := make([]*models.WrapperTransaction, 0)
	eff.db.Table("wrapper_transactions").Where("status != ?", basedef.STATE_FINISHED).Select("wrapper_transactions.hash as src_hash, poly_transactions.hash as poly_hash, dst_transactions.hash as dst_hash").Joins("left join poly_transactions on wrapper_transactions.hash = poly_transactions.src_hash").Joins("left join dst_transactions on poly_transactions.hash = dst_transactions.poly_hash").Preload("WrapperTransaction").Preload("DstTransaction").Find(&wrapperPolyDstRelations)
	for _, wrapperPolyDstRelation := range wrapperPolyDstRelations {
		wrapperTransaction := wrapperPolyDstRelation.WrapperTransaction
		if wrapperPolyDstRelation.PolyHash == "" {
			chain, ok := id2Chains[wrapperPolyDstRelation.WrapperTransaction.SrcChainId]
			if ok {
				if chain.Height-wrapperPolyDstRelation.WrapperTransaction.BlockHeight >= chain.BackwardBlockNumber {
					wrapperTransaction.Status = basedef.STATE_SOURCE_CONFIRMED
				} else {
					wrapperTransaction.Status = basedef.STATE_SOURCE_DONE
				}
			} else {
				wrapperTransaction.Status = basedef.STATE_SOURCE_DONE
			}
		} else if wrapperPolyDstRelation.DstHash == "" {
			wrapperTransaction.Status = basedef.STATE_POLY_CONFIRMED
		} else {
			chain, ok := id2Chains[wrapperPolyDstRelation.DstTransaction.ChainId]
			if ok {
				if chain.Height-wrapperPolyDstRelation.DstTransaction.Height >= 1 {
					wrapperTransaction.Status = basedef.STATE_FINISHED
				} else {
					wrapperTransaction.Status = basedef.STATE_DESTINATION_DONE
				}
			} else {
				wrapperTransaction.Status = basedef.STATE_FINISHED
			}
		}
		wrapperTransactions = append(wrapperTransactions, wrapperTransaction)
	}
	if len(wrapperTransactions) > 0 {
		eff.db.Save(wrapperTransactions)
	}
	return nil
}

func (eff *SwapEffect) checkChainListening() error {
	slot := eff.cfg.ChainListening
	if slot == 0 {
		slot = 300
	}
	old := eff.time / slot
	now := time.Now().Unix()
	new := now / slot
	if new == old {
		return nil
	}
	id2Chains := make(map[uint64]*models.Chain)
	for _, chain := range eff.chains {
		id2Chains[*chain.ChainId] = chain
	}
	chains := make([]*models.Chain, 0)
	eff.db.Model(&models.Chain{}).Find(&chains)
	for _, chain := range chains {
		old, ok := id2Chains[*chain.ChainId]
		if !ok {
			continue
		}
		if chain.Height == old.Height {
			logs.Error("Chain %d is not listening!", *chain.ChainId)
		}
	}
	eff.chains = chains
	eff.time = now
	return nil
}
