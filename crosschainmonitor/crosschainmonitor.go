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

package crosschainmonitor

import (
	"github.com/astaxie/beego/logs"
	"poly-swap/conf"
	"poly-swap/crosschainmonitor/explorermonitor"
	"poly-swap/crosschainmonitor/swapmonitor"
	"runtime/debug"
	"time"
)

type Monitor interface {
	Monitor() error
}

func StartCrossChainMonitor(monitorCfg *conf.CrossChainMonitorConfig, dbCfg *conf.DBConfig) {
	monitor := NewMonitor(monitorCfg, dbCfg)
	if monitor == nil {
		panic("monitor is not valid")
	}
	crossChainMonitor := NewCrossChainMonitor(monitor)
	crossChainMonitor.Start()
}

func NewMonitor(monCfg *conf.CrossChainMonitorConfig, dbCfg *conf.DBConfig) Monitor {
	if monCfg.Server == conf.SERVER_POLY_SWAP {
		return swapmonitor.NewSwapMonitor(monCfg, dbCfg)
	} else if monCfg.Server == conf.SERVER_EXPLORER {
		return explorermonitor.NewExplorerMonitor(monCfg, dbCfg)
	} else {
		return nil
	}
}

type CrossChainMonitor struct {
	monitor Monitor
}

func NewCrossChainMonitor(monitor Monitor) *CrossChainMonitor {
	crossChainMonitor := &CrossChainMonitor{
		monitor: monitor,
	}
	return crossChainMonitor
}

func (mon *CrossChainMonitor) Start() {
	go mon.Check()
}

func (mon *CrossChainMonitor) Check() {
	for {
		mon.check()
	}
}

func (mon *CrossChainMonitor) check() {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("service start, recover info: %s", string(debug.Stack()))
		}
	}()
	logs.Debug("check events %s......")
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			err := mon.monitor.Monitor()
			if err != nil {
				logs.Error("cross chain monitor err: %v", err)
			}
		}
	}
}

