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

package bridgesdk

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"runtime/debug"
	"sync"
	"time"
)

type BridgeInfo struct {
	sdk    *BridgeSdk
	online bool
}

func NewBridgeInfo(url string) *BridgeInfo {
	sdk := NewBridgeSdk(url)
	return &BridgeInfo{
		sdk:    sdk,
		online: true,
	}
}

type BridgeSdkPro struct {
	infos         map[string]*BridgeInfo
	selectionSlot uint64
	mutex         sync.Mutex
}

func NewBridgeSdkPro(urls []string, slot uint64) *BridgeSdkPro {
	infos := make(map[string]*BridgeInfo, len(urls))
	for _, url := range urls {
		infos[url] = NewBridgeInfo(url)
	}
	pro := &BridgeSdkPro{infos: infos, selectionSlot: slot}
	pro.selection()
	go pro.NodeSelection()
	return pro
}

func (pro *BridgeSdkPro) NodeSelection() {
	for {
		pro.nodeSelection()
	}
}

func (pro *BridgeSdkPro) nodeSelection() {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("node selection, recover info: %s", string(debug.Stack()))
		}
	}()
	logs.Debug("node selection of bridge sdk......")
	ticker := time.NewTicker(time.Second * time.Duration(pro.selectionSlot))
	for {
		select {
		case <-ticker.C:
			pro.selection()
		}
	}
}

func (pro *BridgeSdkPro) selection() {
	pro.mutex.Lock()
	defer func() {
		pro.mutex.Unlock()
	}()
	logs.Debug("select node of bridge sdk......")
	for url, info := range pro.infos {
		if info == nil {
			info = NewBridgeInfo(url)
			pro.infos[url] = info
		}
		if info == nil {
			continue
		}
		online, err := info.sdk.Info()
		if err != nil {
			logs.Error("get server info err: %v, url: %s", err, url)
		}
		info.online = online
	}
}

func (pro *BridgeSdkPro) GetLatest() *BridgeInfo {
	pro.mutex.Lock()
	defer func() {
		pro.mutex.Unlock()
	}()
	for _, info := range pro.infos {
		if info != nil && info.online {
			return info
		}
	}
	return nil
}

func (pro *BridgeSdkPro) CheckFee(checks []*CheckFeeReq) ([]*CheckFeeRsp, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}
	for info != nil {
		rsp, err := info.sdk.CheckFee(checks)
		if err != nil {
			logs.Error("check fee err: %v, url: %s", err, info.sdk.url)
			info.online = false
			info = pro.GetLatest()
		} else {
			return rsp, nil
		}
	}
	return nil, fmt.Errorf("all node is not working")
}
