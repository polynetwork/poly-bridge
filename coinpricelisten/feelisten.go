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

package coinpricelisten

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/big"
	"poly-swap/chainsdk"
	"poly-swap/conf"
	"poly-swap/models"
	"runtime/debug"
	"time"
)

type FeeListen struct {
	ethCfg        *conf.EthereumFeeListenConfig
	neoCfg        *conf.NeoFeeListenConfig
	bscCfg        *conf.BscFeeListenConfig
	dbCfg         *conf.DBConfig
	feeUpdateSlot int64
	db            *gorm.DB
}

func StartFeeListen(cfg *conf.FeeListenConfig, dbCfg *conf.DBConfig) {
	feeListen := NewFeeListen(cfg.EthereumFeeListenConfig, cfg.NeoFeeListenConfig, cfg.BscFeeListenConfig, cfg.FeeUpdateSlot, dbCfg)
	feeListen.Start()
}

func NewFeeListen(ethCfg *conf.EthereumFeeListenConfig, neoCfg *conf.NeoFeeListenConfig, bscCfg *conf.BscFeeListenConfig, feeUpdateSlot int64, dbCfg *conf.DBConfig) *FeeListen {
	feeListen := &FeeListen{}
	feeListen.ethCfg = ethCfg
	feeListen.neoCfg = neoCfg
	feeListen.bscCfg = bscCfg
	feeListen.dbCfg = dbCfg
	feeListen.feeUpdateSlot = feeUpdateSlot
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	feeListen.db = db
	//
	chainFees := make([]*models.ChainFee, 0)
	res := db.Find(&chainFees)
	if res.RowsAffected == 0 {
		panic("there is no token basic!")
	}
	err = feeListen.getChainFee(chainFees)
	if err != nil {
		panic(err)
	}
	db.Save(chainFees)
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
				res := this.db.Find(&chainFees)
				if res.RowsAffected == 0 {
					continue
				}
				err := this.getChainFee(chainFees)
				if err != nil {
					continue
				}
				this.db.Save(chainFees)
				break
			}
		}
	}
}

func (this *FeeListen) getChainFee(chainFees []*models.ChainFee) error {
	chainName2Item := make(map[uint64]*models.ChainFee, 0)
	for _, item := range chainFees {
		chainName2Item[item.ChainId] = item
	}
	//
	maxFee, minFee, err1 := this.getEthFee()
	chainFee, ok := chainName2Item[conf.ETHEREUM_CROSSCHAIN_ID]
	if err1 == nil && ok {
		chainFee.MaxFee = &models.BigInt{Int: *maxFee}
		chainFee.MinFee = &models.BigInt{Int: *minFee}
		x := new(big.Int).Mul(minFee, big.NewInt(int64(this.ethCfg.ProxyFee)))
		y := new(big.Int).Div(x, big.NewInt(100))
		chainFee.ProxyFee = &models.BigInt{Int: *y}
	} else {
		logs.Error("get eth fee err: %v", err1)
	}
	//
	maxFee, minFee, err2 := this.getNeoFee()
	chainFee, ok = chainName2Item[conf.NEO_CROSSCHAIN_ID]
	if err2 == nil && ok {
		chainFee.MaxFee = &models.BigInt{Int: *maxFee}
		chainFee.MinFee = &models.BigInt{Int: *minFee}
		x := new(big.Int).Mul(minFee, big.NewInt(int64(this.neoCfg.ProxyFee)))
		y := new(big.Int).Div(x, big.NewInt(100))
		chainFee.ProxyFee = &models.BigInt{Int: *y}
	} else {
		logs.Error("get neo fee err: %v", err2)
	}
	//
	maxFee, minFee, err3 := this.getBscFee()
	chainFee, ok = chainName2Item[conf.BSC_CROSSCHAIN_ID]
	if err2 == nil && ok {
		chainFee.MaxFee = &models.BigInt{Int: *maxFee}
		chainFee.MinFee = &models.BigInt{Int: *minFee}
		x := new(big.Int).Mul(minFee, big.NewInt(int64(this.bscCfg.ProxyFee)))
		y := new(big.Int).Div(x, big.NewInt(100))
		chainFee.ProxyFee = &models.BigInt{Int: *y}
	} else {
		logs.Error("get bsc fee err: %v", err2)
	}
	if err1 != nil || err2 != nil || err3 != nil {
		return fmt.Errorf("can not get the fee information of all chains!")
	}
	return nil
}

func (this *FeeListen) getEthFee() (*big.Int, *big.Int, error) {
	sdk, err := chainsdk.NewEthereumSdk(this.ethCfg.RestURL)
	if err != nil {
		return nil, nil, err
	}
	gasPrice, err := sdk.SuggestGasPrice()
	if err != nil {
		return nil, nil, err
	}
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(conf.PRICE_PRECISION))
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(this.ethCfg.GasLimit))
	return gasPrice, gasPrice, nil
}

func (this *FeeListen) getNeoFee() (*big.Int, *big.Int, error) {
	return big.NewInt(1000000000), big.NewInt(1000000000), nil
}

func (this *FeeListen) getBscFee() (*big.Int, *big.Int, error) {
	sdk, err := chainsdk.NewEthereumSdk(this.bscCfg.RestURL)
	if err != nil {
		return nil, nil, err
	}
	gasPrice, err := sdk.SuggestGasPrice()
	if err != nil {
		return nil, nil, err
	}
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(conf.PRICE_PRECISION))
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(this.bscCfg.GasLimit))
	return gasPrice, gasPrice, nil
}
