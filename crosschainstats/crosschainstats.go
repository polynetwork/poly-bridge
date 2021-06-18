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
	"math/big"
	"time"

	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/crosschaindao/bridgedao"
	"poly-bridge/models"

	"github.com/astaxie/beego/logs"
)

type Stats struct {
	context.Context
	cancel context.CancelFunc
	cfg    *conf.StatsConfig
	dao    *bridgedao.BridgeDao
	exit   chan struct{}
}

var ccs *Stats

// Start - Do stats aggregation/calculation
func StartCrossChainStats(server string, cfg *conf.StatsConfig, dbCfg *conf.DBConfig) {
	if server != basedef.SERVER_POLY_BRIDGE {
		panic("CrossChainStats Only runs on bridge server")
	}
	if cfg == nil || cfg.Interval == 0 {
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

func (this *Stats) Start() {
	ticker := time.NewTicker(time.Second * time.Duration(this.cfg.Interval))
	for {
		select {
		case <-ticker.C:
			err := this.computeStats()
			if err != nil {
				logs.Error("computeStats failed for %s", err)
			}
		case <-this.Done():
			close(this.exit)
			break
		}
	}
}

func (this *Stats) Stop() {
	logs.Info("Stopping stats server")
	this.cancel()
	<-this.exit
}

func (this *Stats) computeStats() (err error) {
	logs.Info("Computing cross chain stats")
	tokens, err := this.dao.GetTokens()
	if err != nil {
		return fmt.Errorf("Failed to fetch token basic list %w", err)
	}
	for _, basic := range tokens {
		err := this.computeTokenStats(basic)
		if err != nil {
			return err
		}
	}
	return
}

func (this *Stats) computeTokenStats(token *models.TokenBasic) (err error) {
	assets := make([]string, len(token.Tokens))
	for i, t := range token.Tokens {
		assets[i] = t.Hash
	}
	last, err := this.dao.GetLastSrcTransferForToken(assets)
	if err != nil || last == nil {
		return err
	}
	checkPoint := token.StatsUpdateTime
	totalAmount, totalCount, err := this.dao.AggregateTokenBasicSrcTransfers(assets, checkPoint, last.Time)
	if err != nil {
		return err
	}
	token.StatsUpdateTime = last.Time
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
