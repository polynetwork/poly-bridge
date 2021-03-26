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
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/crosschaineffect/explorereffect"
	"poly-bridge/crosschaineffect/swapeffect"
	"runtime/debug"
	"time"
)

type Effect interface {
	Effect() error
	Name() string
	GetEffectSlot() int64
}

var crossChainEffect *CrossChainEffect

func StartCrossChainEffect(server string, effCfg *conf.EventEffectConfig, dbCfg *conf.DBConfig) {
	effect := NewEffect(server, effCfg, dbCfg)
	if effect == nil {
		panic("effect is not valid")
	}
	crossChainEffect = NewCrossChainEffect(effect)
	crossChainEffect.Start()
}

func StopCrossChainEffect() {
	if crossChainEffect != nil {
		crossChainEffect.Stop()
	}
}

func NewEffect(server string, effCfg *conf.EventEffectConfig, dbCfg *conf.DBConfig) Effect {
	if server == basedef.SERVER_POLY_SWAP {
		return swapeffect.NewSwapEffect(effCfg, dbCfg)
	} else if server == basedef.SERVER_EXPLORER {
		return explorereffect.NewExplorerEffect(effCfg, dbCfg)
	} else {
		return nil
	}
}

type CrossChainEffect struct {
	effect Effect
	exit   chan bool
}

func NewCrossChainEffect(monitor Effect) *CrossChainEffect {
	crossChainMonitor := &CrossChainEffect{
		effect: monitor,
		exit:   make(chan bool, 0),
	}
	return crossChainMonitor
}

func (eff *CrossChainEffect) Start() {
	logs.Info("start cross chain effect.")
	go eff.Check()
}

func (eff *CrossChainEffect) Stop() {
	eff.exit <- true
	logs.Info("stop cross chain effect.")
}

func (eff *CrossChainEffect) Check() {
	for {
		exit := eff.check()
		if exit {
			close(eff.exit)
			break
		}
		time.Sleep(time.Second * 5)
	}
}

func (eff *CrossChainEffect) check() (exit bool) {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("service start, recover info: %s", string(debug.Stack()))
			exit = false
		}
	}()
	logs.Debug("cross chain effect, server: %s......", eff.effect.Name())
	ticker := time.NewTicker(time.Second * time.Duration(eff.effect.GetEffectSlot()))
	for {
		select {
		case <-ticker.C:
			err := eff.effect.Effect()
			if err != nil {
				logs.Error("cross chain effect err: %v", err)
			}
		case <-eff.exit:
			logs.Info("cross chain effect exit, server: %s......", eff.effect.Name())
			return true
		}
	}
}
