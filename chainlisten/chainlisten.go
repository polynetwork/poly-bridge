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

package chainlisten

import (
	"github.com/astaxie/beego/logs"
	"poly-swap/chainlisten/ethereumlisten"
	"poly-swap/chainlisten/neolisten"
	"poly-swap/chainlisten/polylisten"
	"poly-swap/conf"
	"poly-swap/dao"
	"poly-swap/dao/swap_dao"
	"poly-swap/models"
	"runtime/debug"
	"time"
)

func StartChainListen(listenCfg *conf.ChainListenConfig, dbCfg *conf.DBConfig) {
	db := swap_dao.NewSwapDao(dbCfg)
	db.Start()
	ethereumListen := ethereumlisten.NewEthereumChainListen(listenCfg.EthereumChainListenConfig)
	chainListen := NewChainListen(ethereumListen, db)
	chainListen.Start()
	neoListen := neolisten.NewNeoChainListen(listenCfg.NeoChainListenConfig)
	chainListen = NewChainListen(neoListen, db)
	chainListen.Start()
	polyListen := polylisten.NewPolyChainListen(listenCfg.PolyChainListenConfig)
	chainListen = NewChainListen(polyListen, db)
	chainListen.Start()
}

type ChainHandle interface {
	GetExtendLatestHeight() (uint64, error)
	GetLatestHeight() (uint64, error)
	HandleNewBlock(height uint64) ([]*models.WrapperTransaction, []*models.SrcTransaction, []*models.PolyTransaction, []*models.DstTransaction, error)
	GetBackwardBlockNumber() uint64
	GetChainListenSlot() uint64
	GetChainId() uint64
	GetChainName() string
}

type ChainListen struct {
	handle ChainHandle
	db     dao.CrossChainEventDao
}

func NewChainListen(handle ChainHandle, db dao.CrossChainEventDao) *ChainListen {
	chainListen := &ChainListen{
		handle: handle,
		db:     db,
	}
	return chainListen
}

func (this *ChainListen) Start() {
	go this.ListenChain()
}

func (this *ChainListen) ListenChain() {
	for {
		this.listenChain()
	}
}

func (this *ChainListen) listenChain() {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("service start, recover info: %s", string(debug.Stack()))
		}
	}()
	chain, err := this.db.GetChain(this.handle.GetChainId())
	if err != nil {
		panic(err)
	}
	height, err := this.handle.GetLatestHeight()
	if err != nil || height == 0 {
		panic(err)
	}
	if chain.Height == 0 {
		chain.Height = height
	}
	this.db.UpdateChain(chain)
	logs.Debug("listen chain %s......", this.handle.GetChainName())
	ticker := time.NewTicker(time.Second * time.Duration(this.handle.GetChainListenSlot()))
	for {
		select {
		case <-ticker.C:
			var height, err = this.handle.GetLatestHeight()
			if err != nil || height == 0 {
				logs.Error("listenChain - cannot get chain %s height, err: %s", this.handle.GetChainName(), err)
				continue
			}
			extendHeight, err := this.handle.GetExtendLatestHeight()
			if err != nil || extendHeight == 0 {
				logs.Error("ListenChain - cannot get chain %s extend height, err: %s", this.handle.GetChainName(), err)
			} else if extendHeight-height >= this.handle.GetBackwardBlockNumber() {
				logs.Error("ListenChain - chain %s node is too slow, node height: %d, really height: %d", this.handle.GetChainName(), height, extendHeight)
			}
			if chain.Height >= height {
				continue
			}
			logs.Info("ListenChain - chain %s latest height is %d, listen height: %d", this.handle.GetChainName(), height, chain.Height)
			for chain.Height < height {
				wrapperTransactions, srcTransactions, polyTransactions, dstTransactions, err := this.handle.HandleNewBlock(chain.Height + 1)
				if err != nil {
					logs.Error("HandleNewBlock err: %v", err)
					break
				}
				chain.Height += 1
				err = this.db.UpdateEvents(chain, wrapperTransactions, srcTransactions, polyTransactions, dstTransactions)
				if err != nil {
					logs.Error("UpdateEvents err: %v", err)
					chain.Height -= 1
					break
				}
			}
		}
	}
}
