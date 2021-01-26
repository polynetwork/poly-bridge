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

package crosschainlisten

import (
	"github.com/astaxie/beego/logs"
	"poly-bridge/conf"
	"poly-bridge/crosschaindao"
	"poly-bridge/crosschainlisten/ethereumlisten"
	"poly-bridge/crosschainlisten/neolisten"
	"poly-bridge/crosschainlisten/polylisten"
	"poly-bridge/models"
	"runtime/debug"
	"time"
)

func StartCrossChainListen(server string, listenCfg []*conf.ChainListenConfig, dbCfg *conf.DBConfig) {
	dao := crosschaindao.NewCrossChainDao(server, dbCfg)
	if dao == nil {
		panic("server is not valid")
	}
	for _, cfg := range listenCfg {
		chainHandle := NewChainHandle(cfg)
		chainListen := NewCrossChainListen(chainHandle, dao)
		chainListen.Start()
	}
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

func NewChainHandle(chainListenConfig *conf.ChainListenConfig) ChainHandle {
	if chainListenConfig.ChainId == conf.ETHEREUM_CROSSCHAIN_ID {
		return ethereumlisten.NewEthereumChainListen(chainListenConfig)
	} else if chainListenConfig.ChainId == conf.POLY_CROSSCHAIN_ID {
		return polylisten.NewPolyChainListen(chainListenConfig)
	} else if chainListenConfig.ChainId == conf.NEO_CROSSCHAIN_ID {
		return neolisten.NewNeoChainListen(chainListenConfig)
	} else if chainListenConfig.ChainId == conf.BSC_CROSSCHAIN_ID {
		return ethereumlisten.NewEthereumChainListen(chainListenConfig)
	} else if chainListenConfig.ChainId == conf.HECO_CROSSCHAIN_ID {
		return ethereumlisten.NewEthereumChainListen(chainListenConfig)
	} else {
		return nil
	}
}

type CrossChainListen struct {
	handle ChainHandle
	db     crosschaindao.CrossChainDao
}

func NewCrossChainListen(handle ChainHandle, db crosschaindao.CrossChainDao) *CrossChainListen {
	crossChainListen := &CrossChainListen{
		handle: handle,
		db:     db,
	}
	return crossChainListen
}

func (this *CrossChainListen) Start() {
	go this.ListenChain()
}

func (this *CrossChainListen) ListenChain() {
	for {
		this.listenChain()
		time.Sleep(time.Second * 5)
	}
}

func (this *CrossChainListen) listenChain() {
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
	if chain.BackwardBlockNumber == 0 {
		chain.BackwardBlockNumber = this.handle.GetBackwardBlockNumber()
	}
	this.db.UpdateChain(chain)
	logs.Debug("cross chain listen, chain: %s, dao: %s......", this.handle.GetChainName(), this.db.Name())
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
			} else if extendHeight >= height+this.handle.GetBackwardBlockNumber() {
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
