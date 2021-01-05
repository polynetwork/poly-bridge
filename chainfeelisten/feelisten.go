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

package chainfeelisten

import (
	"github.com/astaxie/beego/logs"
	"math/big"
	"poly-swap/chainfeedao"
	"poly-swap/chainfeelisten/ethereumfee"
	"poly-swap/chainfeelisten/neofee"
	"poly-swap/conf"
	"poly-swap/models"
	"runtime/debug"
	"time"
)

func StartFeeListen(server string, feeUpdateSlot int64, feeListenCfgs []*conf.FeeListenConfig, dbCfg *conf.DBConfig) {
	dao := chainfeedao.NewChainFeeDao(server, dbCfg)
	if dao == nil {
		panic("server is not valid")
	}
	chainFees := make([]ChainFee, 0)
	for _, cfg := range feeListenCfgs {
		chainFee := NewChainFee(cfg, feeUpdateSlot)
		chainFees = append(chainFees, chainFee)
	}
	feeListen := NewFeeListen(feeUpdateSlot, chainFees, dao)
	feeListen.Start()
}

type ChainFee interface {
	GetFee() (*big.Int, *big.Int, *big.Int, error)
	GetChainId() uint64
}

func NewChainFee(cfg *conf.FeeListenConfig, feeUpdateSlot int64) ChainFee {
	if cfg.ChainId == conf.ETHEREUM_CROSSCHAIN_ID {
		return ethereumfee.NewEthereumFee(cfg, feeUpdateSlot)
	} else if cfg.ChainId == conf.NEO_CROSSCHAIN_ID {
		return neofee.NewNeoFee(cfg, feeUpdateSlot)
	} else if cfg.ChainId == conf.BSC_CROSSCHAIN_ID {
		return ethereumfee.NewEthereumFee(cfg, feeUpdateSlot)
	} else {
		return nil
	}
}

type FeeListen struct {
	feeUpdateSlot int64
	fees          map[uint64]ChainFee
	db            chainfeedao.ChainFeeDao
}

func NewFeeListen(feeUpdateSlot int64, fees []ChainFee, db chainfeedao.ChainFeeDao) *FeeListen {
	feeListen := &FeeListen{}
	feeListen.feeUpdateSlot = feeUpdateSlot
	feeListen.db = db
	feeListen.fees = make(map[uint64]ChainFee)
	for _, fee := range fees {
		feeListen.fees[fee.GetChainId()] = fee
	}
	//
	chainFees, err := db.GetFees()
	if err != nil {
		panic(err)
	}
	err = feeListen.updateChainFees(chainFees)
	if err != nil {
		panic(err)
	}
	err = db.SaveFees(chainFees)
	if err != nil {
		panic(err)
	}
	return feeListen
}

func (this *FeeListen) Start() {
	go this.ListenFee()
}

func (this *FeeListen) ListenFee() {
	for {
		this.listenFee()
	}
}

func (this *FeeListen) listenFee() {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("service start, recover info: %s", string(debug.Stack()))
		}
	}()

	logs.Debug("listen fee......")
	ticker := time.NewTicker(time.Minute)
	for {
		select {
		case <-ticker.C:
			now := time.Now().Unix() / 60
			if now%this.feeUpdateSlot != 0 {
				continue
			}
			counter := 0
			for counter < 5 {
				time.Sleep(time.Second * 5)
				counter++
				logs.Info("do fee update at time: %s", time.Now().Format("2006-01-02 15:04:05"))
				chainFees := make([]*models.ChainFee, 0)
				chainFees, err := this.db.GetFees()
				if err != nil {
					logs.Error("get chain fees err: %v", err)
					continue
				}
				err = this.updateChainFees(chainFees)
				if err != nil {
					logs.Error("updateChainFees err: %v", err)
					continue
				}
				err = this.db.SaveFees(chainFees)
				if err != nil {
					logs.Error("save fees err: %v", err)
					continue
				}
				break
			}
		}
	}
}

func (this *FeeListen) updateChainFees(chainFees []*models.ChainFee) error {
	chainFee := make(map[uint64]*models.ChainFee, 0)
	for _, fee := range chainFees {
		chainFee[fee.ChainId] = fee
		fee.Ind = 0
	}
	for chainId, query := range this.fees {
		fee, ok := chainFee[chainId]
		if !ok {
			logs.Error("this is no fee of chain: %d", chainId)
			continue
		}
		minFee, maxFee, proxyFee, err := query.GetFee()
		if err != nil {
			logs.Error("get fee of chain: %d err: %v", chainId, err)
			continue
		}
		logs.Info("get fee of chain: %d successful", chainId)
		fee.MinFee = models.NewBigInt(minFee)
		fee.MaxFee = models.NewBigInt(maxFee)
		fee.ProxyFee = models.NewBigInt(proxyFee)
		fee.Ind = 1
	}
	for _, fee := range chainFees {
		if fee.Ind == 0 {
			logs.Error("fee of chain %d is not update", fee.ChainId)
		}
	}
	return nil
}
