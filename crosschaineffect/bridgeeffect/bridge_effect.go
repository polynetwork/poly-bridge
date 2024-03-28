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

package bridgeeffect

import (
	"encoding/json"
	"fmt"
	"poly-bridge/basedef"
	"poly-bridge/cacheRedis"
	"poly-bridge/conf"
	"poly-bridge/models"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var checkTime int = 0
var counterTime int = 0

type BridgeEffect struct {
	dbCfg    *conf.DBConfig
	cfg      *conf.EventEffectConfig
	db       *gorm.DB
	redis    *cacheRedis.RedisCache
	redisCfg *conf.RedisConfig
	chains   []*models.Chain
	time     int64
}

func NewBridgeEffect(cfg *conf.EventEffectConfig, dbCfg *conf.DBConfig, redisCfg *conf.RedisConfig) *BridgeEffect {
	swapEffect := &BridgeEffect{
		dbCfg:    dbCfg,
		cfg:      cfg,
		redisCfg: redisCfg,
		chains:   nil,
		time:     0,
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
	redis, err := cacheRedis.GetRedisClient(redisCfg)
	if err != nil {
		panic(err)
	}
	swapEffect.redis = redis
	chains := make([]*models.Chain, 0)
	res := db.Model(&models.Chain{}).Find(&chains)
	if res.Error != nil || res.RowsAffected == 0 {
		panic(err)
	}
	swapEffect.chains = chains
	swapEffect.time = time.Now().Unix()
	return swapEffect
}

func (eff *BridgeEffect) Effect() error {
	err := eff.updateHash()
	if err != nil {
		logs.Error("update hash- err: %s", err)
	}
	err = eff.updateBfcHash()
	if err != nil {
		logs.Error("updateBfcHash hash- err: %s", err)
	}
	err = eff.updateDstHash()
	if err != nil {
		logs.Error("update dsthash- err: %s", err)
	}
	/*
		err = eff.checkStatus()
		if err != nil {
			logs.Error("check status- err: %s", err)
		}
	*/
	err = eff.updateStatus()
	if err != nil {
		logs.Error("update status- err: %s", err)
	}
	err = eff.doStatistic()
	if err != nil {
		logs.Error("update status- err: %s", err)
	}
	err = eff.checkChainListening()
	if err != nil {
		logs.Error("check chain listening- err: %s", err)
	}
	counterTime++
	if counterTime > 180 {
		counterTime = 0
		err = eff.StartUpdateCrossCount()
		if err != nil {
			logs.Error("UpdateCrossCount err: %s", err)
		}
	}
	return nil
}

func (eff *BridgeEffect) StartUpdateCrossCount() error {
	logs.Info("StartUpdateCrossCount start")
	var counter int64
	res := eff.db.Model(&models.PolyTransaction{}).
		Where("src_transactions.standard = ?", 0).
		Joins("left join src_transactions on src_transactions.hash = poly_transactions.src_hash").
		Count(&counter)
	if res.RowsAffected == 0 {
		return fmt.Errorf("StartUpdateCrossCount counter err %w", res.Error)
	}
	err := eff.redis.SetCrossTxCounter(counter)
	if err != nil {
		return fmt.Errorf("StartUpdateCrossCount SetCrossTxCounter err %w", err)
	}
	return nil
}
func (eff *BridgeEffect) Name() string {
	return basedef.SERVER_POLY_SWAP
}

func (eff *BridgeEffect) GetEffectSlot() int64 {
	return eff.cfg.EffectSlot
}

func (eff *BridgeEffect) updateHash() error {
	batch := 500
	index := 0

	for {
		polySrcRelations := make([]*models.PolySrcRelation, 0)
		eff.db.Table("poly_transactions").Where("poly_transactions.src_hash like ? and poly_transactions.time > ?", "00000000%", 1622476800).Select("poly_transactions.hash as poly_hash, src_transactions.hash as src_hash").Joins("inner join src_transactions on poly_transactions.src_hash = src_transactions.key and poly_transactions.src_chain_id = src_transactions.chain_id").Preload("SrcTransaction").Preload("PolyTransaction").Limit(batch).Offset(batch * index).Order("poly_transactions.time desc").Find(&polySrcRelations)
		updatePolyTransactions := make([]*models.PolyTransaction, 0)
		for _, polySrcRelation := range polySrcRelations {
			if polySrcRelation.SrcTransaction != nil {
				polySrcRelation.PolyTransaction.Key = polySrcRelation.PolyTransaction.SrcHash
				polySrcRelation.PolyTransaction.SrcHash = polySrcRelation.SrcHash
				updatePolyTransactions = append(updatePolyTransactions, polySrcRelation.PolyTransaction)
			}
		}
		if len(updatePolyTransactions) > 0 {
			logs.Info("updateHash now min PolyTransaction.id", updatePolyTransactions[0].Id)
			eff.db.Save(updatePolyTransactions)
			index++
		} else {
			break
		}
	}
	logs.Info("Update hash finished with at most %d * 500 checked", index+1)
	return nil
}

func (eff *BridgeEffect) updateBfcHash() error {
	batch := 500
	index := 0

	for {
		polySrcRelations := make([]*models.PolySrcRelation, 0)
		eff.db.Table("poly_transactions").Where("poly_transactions.src_chain_id = ? and poly_transactions.time > ? and poly_transactions.key = ?", basedef.BFC_CROSSCHAIN_ID, 1622476800, "").Select("poly_transactions.hash as poly_hash, src_transactions.hash as src_hash").Joins("inner join src_transactions on poly_transactions.src_hash = src_transactions.key and poly_transactions.src_chain_id = src_transactions.chain_id").Preload("SrcTransaction").Preload("PolyTransaction").Limit(batch).Offset(batch * index).Order("poly_transactions.time desc").Find(&polySrcRelations)
		updatePolyTransactions := make([]*models.PolyTransaction, 0)
		for _, polySrcRelation := range polySrcRelations {
			if polySrcRelation.SrcTransaction != nil {
				polySrcRelation.PolyTransaction.Key = polySrcRelation.PolyTransaction.SrcHash
				polySrcRelation.PolyTransaction.SrcHash = polySrcRelation.SrcHash
				updatePolyTransactions = append(updatePolyTransactions, polySrcRelation.PolyTransaction)
			}
		}
		if len(updatePolyTransactions) > 0 {
			logs.Info("updateBfcHash now min PolyTransaction.id", updatePolyTransactions[0].Id)
			eff.db.Save(updatePolyTransactions)
			index++
		} else {
			break
		}
	}
	logs.Info("Update updateBfcHash finished with at most %d * 500 checked", index+1)
	return nil
}

func (eff *BridgeEffect) updateDstHash() error {
	batch := 500
	index := 0

	for {
		dstPolyRelations := make([]*models.DstPolyRelation, 0)
		eff.db.Table("dst_transactions").
			Where("dst_transactions.poly_hash = '' and dst_transactions.chain_id = ? and dst_transactions.sequence > 0", basedef.RIPPLE_CROSSCHAIN_ID).
			Select("dst_transactions.hash as dst_hash, poly_transactions.hash as poly_hash").
			Joins("inner join poly_transactions on dst_transactions.sequence = poly_transactions.dst_sequence and poly_transactions.dst_chain_id = dst_transactions.chain_id").
			Preload("PolyTransaction").
			Preload("DstTransaction").
			Limit(batch).
			Offset(batch * index).
			Order("poly_transactions.id").
			Find(&dstPolyRelations)
		updateDstTransactions := make([]*models.DstTransaction, 0)
		for _, dstPolyRelation := range dstPolyRelations {
			if dstPolyRelation.PolyTransaction != nil && dstPolyRelation.DstTransaction != nil {
				dstPolyRelation.DstTransaction.PolyHash = dstPolyRelation.PolyTransaction.Hash
				updateDstTransactions = append(updateDstTransactions, dstPolyRelation.DstTransaction)
			}
		}
		if len(updateDstTransactions) > 0 {
			logs.Info("updateHash now min DstTransaction.id", updateDstTransactions[0].Id)
			eff.db.Save(updateDstTransactions)
			index++
		} else {
			break
		}
	}
	logs.Info("Update DstHash finished with at most %d * 500 checked", index+1)
	return nil
}

func (eff *BridgeEffect) checkStatus() error {
	{
		wrapperTransactions := make([]*models.WrapperTransaction, 0)
		now := time.Now().Unix() - eff.cfg.HowOld2
		eff.db.Model(models.WrapperTransaction{}).Where("(status NOT IN ? and time < ?) and ((src_chain_id = ? and dst_chain_id = ?) or (src_chain_id = ? and dst_chain_id = ?))",
			[]int{basedef.STATE_FINISHED, basedef.STATE_WAIT, basedef.STATE_SKIP},
			now, basedef.BSC_CROSSCHAIN_ID, basedef.HECO_CROSSCHAIN_ID, basedef.HECO_CROSSCHAIN_ID, basedef.BSC_CROSSCHAIN_ID).Find(&wrapperTransactions)
		if len(wrapperTransactions) > 0 {
			wrapperTransactionsJson, _ := json.Marshal(wrapperTransactions)
			logs.Error("There is unfinished transactions(%d) %s", now, string(wrapperTransactionsJson))
		}
	}
	{
		wrapperTransactions := make([]*models.WrapperTransaction, 0)
		now := time.Now().Unix() - eff.cfg.HowOld
		eff.db.Model(models.WrapperTransaction{}).Where("status != ? and time < ?", basedef.STATE_FINISHED, now).Find(&wrapperTransactions)
		if len(wrapperTransactions) > 0 {
			wrapperTransactionsJson, _ := json.Marshal(wrapperTransactions)
			logs.Error("There is unfinished transactions(%d) %s", now, string(wrapperTransactionsJson))
		}
	}
	return nil
}

func (eff *BridgeEffect) updateStatus() error {
	chains := make([]*models.Chain, 0)
	id2Chains := make(map[uint64]*models.Chain)
	eff.db.Model(&models.Chain{}).Find(&chains)
	for _, chain := range chains {
		id2Chains[chain.ChainId] = chain
	}

	batch := 500
	index := 0
	for {
		wrapperPolyDstRelations := make([]*models.SrcPolyDstRelation, 0)
		wrapperTransactions := make([]*models.WrapperTransaction, 0)
		eff.db.Table("wrapper_transactions").Where("wrapper_transactions.status != ? and wrapper_transactions.time > 1622476800", basedef.STATE_FINISHED).Select("wrapper_transactions.hash as src_hash, poly_transactions.hash as poly_hash, dst_transactions.hash as dst_hash").Joins("left join poly_transactions on wrapper_transactions.hash = poly_transactions.src_hash").Joins("left join dst_transactions on poly_transactions.hash = dst_transactions.poly_hash").Preload("WrapperTransaction").Preload("DstTransaction").Limit(batch).Offset(batch * index).Order("wrapper_transactions.time desc").Find(&wrapperPolyDstRelations)
		for _, wrapperPolyDstRelation := range wrapperPolyDstRelations {
			wrapperTransaction := wrapperPolyDstRelation.WrapperTransaction
			pending := wrapperTransaction.Status == basedef.STATE_SKIP || wrapperTransaction.Status == basedef.STATE_WAIT
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
			if !pending || wrapperTransaction.Status == basedef.STATE_FINISHED {
				wrapperTransactions = append(wrapperTransactions, wrapperTransaction)
			}
		}
		if len(wrapperTransactions) > 0 {
			eff.db.Save(wrapperTransactions)
		}
		if len(wrapperPolyDstRelations) == 0 {
			break
		}
		index++
	}
	logs.Info("Update wrapper tx status finished with at most %d * 500 checked", index+1)
	return nil
}

func (eff *BridgeEffect) checkChainListening() error {
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
		id2Chains[chain.ChainId] = chain
	}
	chains := make([]*models.Chain, 0)
	eff.db.Model(&models.Chain{}).Find(&chains)
	for _, chain := range chains {
		old, ok := id2Chains[chain.ChainId]
		if !ok {
			continue
		}
		if chain.Height == old.Height && chain.HeightSwap == old.HeightSwap {
			logs.Error("Chain %d is not listening!", chain.ChainId)
		}
	}
	eff.chains = chains
	eff.time = now
	return nil
}

type TimeStatistic struct {
	SrcChainId uint64
	DstChainId uint64
	Time       float64
}

func (eff *BridgeEffect) doStatistic() error {
	timeStatistics := make([]*TimeStatistic, 0)
	start := time.Now().Unix() - eff.cfg.TimeStatisticSlot
	res := eff.db.Raw("select avg(c.time - a.time) * 100000000 as time, a.chain_id as src_chain_id, c.chain_id as dst_chain_id from src_transactions a inner join poly_transactions b on a.hash = b.src_hash inner join dst_transactions c on b.hash = c.poly_hash inner join wrapper_transactions d on a.hash = d.hash where c.time > ? group by a.chain_id,c.chain_id;", start).Scan(&timeStatistics)
	if res.Error != nil {
		logs.Error("do avg time statistic err: %v", res.Error.Error())
	}
	newTimeStatistics := make([]*models.TimeStatistic, 0)
	for _, item := range timeStatistics {
		if item.Time < 0 {
			continue
		}
		newTimeStatistics = append(newTimeStatistics, &models.TimeStatistic{
			SrcChainId: item.SrcChainId,
			DstChainId: item.DstChainId,
			Time:       uint64(item.Time),
		})
	}
	if len(newTimeStatistics) > 0 {
		res = eff.db.Save(newTimeStatistics)
		if res.Error != nil {
			logs.Error("save avg time statistic err: %v", res.Error.Error())
		}
	}
	return nil
}
