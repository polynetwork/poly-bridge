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

package crosschaineffect

import (
	"github.com/astaxie/beego/logs"
	"poly-bridge/conf"
	"poly-bridge/crosschaineffect/explorereffect"
	"poly-bridge/crosschaineffect/swapeffect"
	"runtime/debug"
	"time"
)

type Effect interface {
	Effect() error
	Name() string
}

func StartCrossChainEffect(effCfg *conf.EventEffectConfig, dbCfg *conf.DBConfig) {
	effect := NewEffect(effCfg, dbCfg)
	if effect == nil {
		panic("effect is not valid")
	}
	crossChainEffect := NewCrossChainEffect(effect)
	crossChainEffect.Start()
}

func NewEffect(effCfg *conf.EventEffectConfig, dbCfg *conf.DBConfig) Effect {
	if effCfg.Server == conf.SERVER_POLY_SWAP {
		return swapeffect.NewSwapEffect(effCfg, dbCfg)
	} else if effCfg.Server == conf.SERVER_EXPLORER {
		return explorereffect.NewExplorerEffect(effCfg, dbCfg)
	} else {
		return nil
	}
}

type CrossChainEffect struct {
	monitor Effect
}

func NewCrossChainEffect(monitor Effect) *CrossChainEffect {
	crossChainMonitor := &CrossChainEffect{
		monitor: monitor,
	}
	return crossChainMonitor
}

func (eff *CrossChainEffect) Start() {
	go eff.Check()
}

func (eff *CrossChainEffect) Check() {
	for {
		eff.check()
		time.Sleep(time.Second * 5)
	}
}

func (eff *CrossChainEffect) check() {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("service start, recover info: %s", string(debug.Stack()))
		}
	}()
	logs.Debug("cross chain effect, server: %s......", eff.monitor.Name())
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			err := eff.monitor.Effect()
			if err != nil {
				logs.Error("cross chain monitor err: %v", err)
			}
		}
	}
}
