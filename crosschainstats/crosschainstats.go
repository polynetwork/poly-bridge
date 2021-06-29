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

package crosschainstats

import (
	"context"
	"fmt"
	"github.com/shopspring/decimal"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/common"
	"poly-bridge/conf"
	"poly-bridge/crosschaindao/bridgedao"
	"poly-bridge/models"
	"sync"
	"time"

	"github.com/beego/beego/v2/core/logs"
)

type Stats struct {
	context.Context
	cancel context.CancelFunc
	cfg    *conf.StatsConfig
	dao    *bridgedao.BridgeDao
	wg     sync.WaitGroup
}

var ccs *Stats

// Start - Do stats aggregation/calculation
func StartCrossChainStats(server string, cfg *conf.StatsConfig, dbCfg *conf.DBConfig) {
	if server != basedef.SERVER_POLY_BRIDGE {
		panic("CrossChainStats Only runs on bridge server")
	}
	if cfg == nil || cfg.TokenBasicStatsInterval == 0 || cfg.TokenStatsInterval == 0 {
		panic("Invalid Stats config")
	}

	dao := bridgedao.NewBridgeDao(dbCfg, false)
	ctx, cancel := context.WithCancel(context.Background())
	ccs = &Stats{dao: dao, cfg: cfg, Context: ctx, cancel: cancel}
	ccs.Start()
}

// Stop
func StopCrossChainStats() {
	if ccs != nil {
		ccs.Stop()
	}
}

func (this *Stats) run(interval int64, f func() error) {
	this.wg.Add(1)
	ticker := time.NewTicker(time.Second * time.Duration(interval))
	for {
		select {
		case <-ticker.C:
			err := f()
			if err != nil {
				logs.Error("stats run error%s", err)
			}
		case <-this.Done():
			break
		}
	}
	this.wg.Done()
}

func (this *Stats) Start() {
	go this.run(this.cfg.TokenBasicStatsInterval, this.computeStats)
	go this.run(this.cfg.TokenStatsInterval, this.computeTokensStats)
	go this.run(this.cfg.TokenStatisticInterval, this.computeTokenStatistics)
	go this.run(this.cfg.ChainStatisticInterval, this.computeChainStatistics)
	go this.run(this.cfg.ChainStatisticAssetInterval, this.computeChainStatisticAssets)
}

func (this *Stats) Stop() {
	logs.Info("Stopping stats server")
	this.cancel()
	this.wg.Wait()
}

func (this *Stats) computeStats() (err error) {
	logs.Info("Computing cross chain token basic stats")
	tokens, err := this.dao.GetTokenBasics()
	if err != nil {
		return fmt.Errorf("Failed to fetch token basic list %w", err)
	}
	for _, basic := range tokens {
		err := this.computeTokenBasicStats(basic)
		if err != nil {
			return err
		}
	}
	return
}

func (this *Stats) computeTokenBasicStats(token *models.TokenBasic) (err error) {
	assets := make([][]interface{}, len(token.Tokens))
	for i, t := range token.Tokens {
		assets[i] = []interface{}{t.ChainId, t.Hash}
	}
	checkPoint := token.StatsUpdateTime
	last, err := this.dao.GetLastSrcTransferForToken(assets)
	if err != nil || last == nil || checkPoint >= last.Id {
		return err
	}
	totalAmount, totalCount, err := this.dao.AggregateTokenBasicSrcTransfers(assets, checkPoint, last.Id)
	if err != nil {
		return err
	}
	token.StatsUpdateTime = last.Id
	if checkPoint == 0 {
		token.TotalAmount = &models.BigInt{*totalAmount}
		token.TotalCount = totalCount
	} else {
		token.TotalAmount = &models.BigInt{*new(big.Int).Add(totalAmount, &token.TotalAmount.Int)}
		token.TotalCount += totalCount
	}
	err = this.dao.UpdateTokenBasicStatsWithCheckPoint(token, checkPoint)
	return
}

func (this *Stats) computeTokensStats() (err error) {
	logs.Info("Computing cross chain token stats")
	tokens, err := this.dao.GetTokens()
	if err != nil {
		return fmt.Errorf("Failed to fetch token basic list %w", err)
	}
	for _, t := range tokens {
		amount, err := common.GetBalance(t.ChainId, t.Hash)
		if err != nil || amount == nil {
			logs.Error("Failed to fetch token available amount for token %s %v %s", t.Hash, t.ChainId, err)
			continue
		}
		err = this.dao.UpdateTokenAvailableAmount(t.Hash, t.ChainId, amount)
		if err != nil {
			logs.Error("Failed to update token available amount for token %s %v %s", t.Hash, t.ChainId, err)
		}
	}
	return
}

func (this *Stats) computeTokenStatistics() (err error) {
	nowInId := this.dao.GetNewDstTransfer().Id
	nowOutId := this.dao.GetNewSrcTransfer().Id
	nowTokenStatistic := this.dao.GetNewTokenSta()

	inTokenStatistics := make([]*models.TokenStatistic, 0)
	if nowInId > nowTokenStatistic.LastInCheckId {
		err = this.dao.CalculateInTokenStatistics(nowTokenStatistic.LastInCheckId, nowInId, inTokenStatistics)
		if err != nil {
			return fmt.Errorf("Failed to CalculateInTokenStatistics %w", err)
		}
		for _, tokenStatistic := range inTokenStatistics {
			precision_new := decimal.New(int64(tokenStatistic.Token.Precision), 0)
			inAmount_new := decimal.New(tokenStatistic.InAmount.Int64(), 0)
			price_new := decimal.New(tokenStatistic.Token.TokenBasic.Price, 0)
			tokenStatistic.InAmountUsdt = models.NewBigInt(inAmount_new.Div(precision_new).Mul(price_new).BigInt())
		}
	}
	outTokenStatistics := make([]*models.TokenStatistic, 0)
	if nowOutId > nowTokenStatistic.LastOutCheckId {
		err = this.dao.CalculateInTokenStatistics(nowTokenStatistic.LastOutCheckId, nowOutId, outTokenStatistics)
		if err != nil {
			return fmt.Errorf("Failed to CalculateInTokenStatistics %w", err)
		}
		for _, tokenStatistic := range outTokenStatistics {
			precision_new := decimal.New(int64(tokenStatistic.Token.Precision), 0)
			outAmount_new := decimal.New(tokenStatistic.OutAmount.Int64(), 0)
			price_new := decimal.New(tokenStatistic.Token.TokenBasic.Price, 0)
			tokenStatistic.OutAmountUsdt = models.NewBigInt(outAmount_new.Div(precision_new).Mul(price_new).BigInt())
		}
	}
	if nowInId > nowTokenStatistic.LastInCheckId || nowOutId > nowTokenStatistic.LastOutCheckId {
		tokenStatistics := make([]*models.TokenStatistic, 0)
		err = this.dao.GetTokenStatistics(tokenStatistics)
		if err != nil {
			return fmt.Errorf("Failed to GetTokenStatistics %w", err)
		}
		for _, statistic := range tokenStatistics {
			for _, in := range inTokenStatistics {
				if statistic.ChainId == in.ChainId && statistic.Hash == in.Hash {
					statistic.InAmount = addDecimalBigInt(statistic.InAmount, in.InAmount)
					statistic.InCounter = addDecimalInt64(statistic.InCounter, in.InCounter)
					statistic.InAmountUsdt = addDecimalBigInt(statistic.InAmountUsdt, in.InAmountUsdt)
					statistic.LastInCheckId = nowInId
				}
			}
			for _, out := range outTokenStatistics {
				if statistic.ChainId == out.ChainId && statistic.Hash == out.Hash {
					statistic.OutAmount = addDecimalBigInt(statistic.OutAmount, out.OutAmount)
					statistic.OutCounter = addDecimalInt64(statistic.OutCounter, out.OutCounter)
					statistic.OutAmountUsdt = addDecimalBigInt(statistic.OutAmountUsdt, out.OutAmountUsdt)
					statistic.LastOutCheckId = nowOutId
				}
			}
		}
		err = this.dao.SaveTokenStatistics(tokenStatistics)
		if err != nil {
			return fmt.Errorf("Failed to SaveTokenStatistics %w", err)
		}
	}
	return
}

func (this *Stats) computeChainStatistics() (err error) {
	nowChainStatistic := this.dao.GetNewChainSta()
	nowInId := this.dao.GetNewDstTransfer().Id
	nowOutId := this.dao.GetNewSrcTransfer().Id

	inChainStatistics := make([]*models.ChainStatistic, 0)
	if nowInId > nowChainStatistic.LastInCheckId {
		err = this.dao.CalculateInChainStatistics(nowChainStatistic.LastInCheckId, nowInId, inChainStatistics)
		if err != nil {
			logs.Error("Failed to CalculateInTokenStatistics %w", err)
		}
	}
	outChainStatistics := make([]*models.ChainStatistic, 0)
	if nowOutId > nowChainStatistic.LastOutCheckId {
		err = this.dao.CalculateOutChainStatistics(nowChainStatistic.LastOutCheckId, nowOutId, outChainStatistics)
		if err != nil {
			logs.Error("Failed to CalculateInTokenStatistics %w", err)
		}
	}
	if nowInId > nowChainStatistic.LastInCheckId || nowOutId > nowChainStatistic.LastOutCheckId {
		chainStatistics := make([]*models.ChainStatistic, 0)
		err = this.dao.GetChainStatistic(chainStatistics)
		if err != nil {
			return fmt.Errorf("Failed to CalculateInTokenStatistics %w", err)
		}
		for _, chainStatistic := range chainStatistics {
			for _, in := range inChainStatistics {
				if chainStatistic.ChainId == in.ChainId {
					chainStatistic.In = addDecimalInt64(chainStatistic.In, in.In)
					chainStatistic.LastInCheckId = nowInId
				}
			}
			for _, out := range outChainStatistics {
				if chainStatistic.ChainId == out.ChainId {
					chainStatistic.Out = addDecimalInt64(chainStatistic.Out, out.Out)
					chainStatistic.LastOutCheckId = nowOutId
				}
			}
		}
	}
	return
}
func (this *Stats) computeChainStatisticAssets() (err error) {
	computeChainStatistics := make([]*models.ChainStatistic, 0)
	err = this.dao.CalculateChainStatisticAssets(computeChainStatistics)
	if err != nil {
		return fmt.Errorf("Failed to CalculateChainStatisticAssets %w", err)
	}
	return
	chainStatistics := make([]*models.ChainStatistic, 0)
	this.dao.GetChainStatistic(chainStatistics)
	for _, chainStatistic := range chainStatistics {
		for _, new := range computeChainStatistics {
			if chainStatistic.ChainId == new.ChainId {
				chainStatistic.Addresses = new.Addresses
			}
		}
	}
	return
}

func addDecimalBigInt(a, b *models.BigInt) *models.BigInt {
	a_new := decimal.New(a.Int64(), 0)
	b_new := decimal.New(b.Int64(), 0)
	c := a_new.Add(b_new)
	return models.NewBigInt(c.BigInt())
}
func addDecimalInt64(a, b int64) int64 {
	a_new := decimal.New(a, 0)
	b_new := decimal.New(b, 0)
	c := a_new.Add(b_new)
	return c.IntPart()
}
