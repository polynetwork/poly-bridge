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
	"poly-bridge/basedef"
	"poly-bridge/chainfeedao"
	"poly-bridge/chainfeelisten/ethereumfee"
	"poly-bridge/chainfeelisten/neofee"
	"poly-bridge/chainfeelisten/ontologyfee"
	"poly-bridge/conf"
	"poly-bridge/models"
	"runtime/debug"
	"strings"
	"time"
)

var feeListen *FeeListen

func StartFeeListen(server string, feeUpdateSlot int64, feeListenCfgs []*conf.FeeListenConfig, dbCfg *conf.DBConfig) {
	dao := chainfeedao.NewChainFeeDao(server, dbCfg)
	if dao == nil {
		panic("server is not valid")
	}
	chainFees := make([]ChainFee, 0)
	for _, cfg := range feeListenCfgs {
		chainFee := NewChainFee(cfg, feeUpdateSlot)
		if chainFee == nil {
			panic("chain fee is not valid")
		}
		chainFees = append(chainFees, chainFee)
	}
	feeListen = NewFeeListen(feeUpdateSlot, chainFees, dao)
	feeListen.Start()
}

func StopFeeListen() {
	if feeListen != nil {
		feeListen.Stop()
	}
}

type ChainFee interface {
	GetFee() (*big.Int, *big.Int, *big.Int, error)
	GetChainId() uint64
	Name() string
}

func NewChainFee(cfg *conf.FeeListenConfig, feeUpdateSlot int64) ChainFee {
	if cfg.ChainId == basedef.ETHEREUM_CROSSCHAIN_ID {
		return ethereumfee.NewEthereumFee(cfg, feeUpdateSlot)
	} else if cfg.ChainId == basedef.NEO_CROSSCHAIN_ID {
		return neofee.NewNeoFee(cfg, feeUpdateSlot)
	} else if cfg.ChainId == basedef.BSC_CROSSCHAIN_ID {
		return ethereumfee.NewEthereumFee(cfg, feeUpdateSlot)
	} else if cfg.ChainId == basedef.HECO_CROSSCHAIN_ID {
		return ethereumfee.NewEthereumFee(cfg, feeUpdateSlot)
	} else if cfg.ChainId == basedef.ONT_CROSSCHAIN_ID {
		return ontologyfee.NewOntologyFee(cfg, feeUpdateSlot)
	} else {
		return nil
	}
}

type FeeListen struct {
	feeUpdateSlot int64
	fees          map[uint64]ChainFee
	db            chainfeedao.ChainFeeDao
	exit          chan bool
}

func NewFeeListen(feeUpdateSlot int64, fees []ChainFee, db chainfeedao.ChainFeeDao) *FeeListen {
	feeListen := &FeeListen{}
	feeListen.feeUpdateSlot = feeUpdateSlot
	feeListen.db = db
	feeListen.exit = make(chan bool, 0)
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

func (fl *FeeListen) Start() {
	logs.Info("start chain fee listen.")
	go fl.ListenFee()
}

func (fl *FeeListen) Stop() {
	fl.exit <- true
	logs.Info("stop chain fee listen.")
}

func (fl *FeeListen) ListenFee() {
	for {
		exit := fl.listenFee()
		if exit {
			close(fl.exit)
			break
		}
		time.Sleep(time.Second * 5)
	}
}

func (fl *FeeListen) listenFee() (exit bool) {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("service start, recover info: %s", string(debug.Stack()))
			exit = false
		}
	}()

	logs.Debug("fee listen, chain: %s, dao: %s......", fl.GetChainFees(), fl.db.Name())
	ticker := time.NewTicker(time.Minute)
	for {
		select {
		case <-ticker.C:
			now := time.Now().Unix() / 60
			if now%fl.feeUpdateSlot != 0 {
				continue
			}
			counter := 0
			for counter < 5 {
				time.Sleep(time.Second * 5)
				counter++
				logs.Info("do fee update at time: %s", time.Now().Format("2006-01-02 15:04:05"))
				chainFees := make([]*models.ChainFee, 0)
				chainFees, err := fl.db.GetFees()
				if err != nil {
					logs.Error("get chain fees err: %v", err)
					continue
				}
				err = fl.updateChainFees(chainFees)
				if err != nil {
					logs.Error("updateChainFees err: %v", err)
					continue
				}
				err = fl.db.SaveFees(chainFees)
				if err != nil {
					logs.Error("save fees err: %v", err)
					continue
				}
				break
			}
		case <-fl.exit:
			logs.Info("fee listen exit, chain: %s, dao: %s......", fl.GetChainFees(), fl.db.Name())
			return true
		}
	}
}

func (fl *FeeListen) updateChainFees(chainFees []*models.ChainFee) error {
	chainFee := make(map[uint64]*models.ChainFee, 0)
	for _, fee := range chainFees {
		chainFee[fee.ChainId] = fee
		fee.Ind = 0
	}
	for chainId, query := range fl.fees {
		fee, ok := chainFee[chainId]
		if !ok {
			logs.Error("fl is no fee of chain: %d", chainId)
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
		fee.Time = time.Now().Unix()
		fee.Ind = 1
	}
	for _, fee := range chainFees {
		if fee.Ind == 0 {
			logs.Error("fee of chain %d is not update", fee.ChainId)
		}
	}
	return nil
}

func (fl *FeeListen) GetChainFees() string {
	fees := make([]string, 0)
	for _, fee := range fl.fees {
		fees = append(fees, fee.Name())
	}
	return strings.Join(fees, ",")
}
