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

package chainsdk

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/polynetwork/poly-go-sdk/common"
	"github.com/polynetwork/poly/core/types"
	"runtime/debug"
	"sync"
	"time"
)

type PolyInfo struct {
	sdk          *PolySDK
	latestHeight uint64
}

func NewPolyInfo(url string) *PolyInfo {
	sdk := NewPolySDK(url)
	return &PolyInfo{
		sdk:          sdk,
		latestHeight: 0,
	}
}

type PolySDKPro struct {
	infos         map[string]*PolyInfo
	selectionSlot uint64
	id            uint64
	mutex         sync.Mutex
}

func NewPolySDKPro(urls []string, slot uint64, id uint64) *PolySDKPro {
	infos := make(map[string]*PolyInfo, len(urls))
	for _, url := range urls {
		infos[url] = NewPolyInfo(url)
	}
	pro := &PolySDKPro{infos: infos, selectionSlot: slot, id: id}
	pro.selection()
	go pro.NodeSelection()
	return pro
}

func (pro *PolySDKPro) NodeSelection() {
	for {
		pro.nodeSelection()
	}
}

func (pro *PolySDKPro) nodeSelection() {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("node selection, recover info: %s", string(debug.Stack()))
		}
	}()
	logs.Debug("node selection of chain : %d......", pro.id)
	ticker := time.NewTicker(time.Second * time.Duration(pro.selectionSlot))
	for {
		select {
		case <-ticker.C:
			pro.selection()
		}
	}
}

func (pro *PolySDKPro) selection() {
	for url, info := range pro.infos {
		height, err := info.sdk.GetCurrentBlockHeight()
		if err != nil {
			logs.Error("get current block height err: %v, url: %s", err, url)
		}
		pro.mutex.Lock()
		info.latestHeight = height
		pro.mutex.Unlock()
	}
}

func (pro *PolySDKPro) GetLatest() *PolyInfo {
	pro.mutex.Lock()
	defer func() {
		pro.mutex.Unlock()
	}()
	height := uint64(0)
	var latestInfo *PolyInfo = nil
	for _, info := range pro.infos {
		if info != nil && info.latestHeight > height {
			height = info.latestHeight
			latestInfo = info
		}
	}
	return latestInfo
}

func (pro *PolySDKPro) GetCurrentBlockHeight() (uint64, error) {
	info := pro.GetLatest()
	if info == nil {
		return 0, fmt.Errorf("all node is not working")
	}
	return info.latestHeight, nil
}

func (pro *PolySDKPro) GetBlockByHeight(height uint64) (*types.Block, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}
	for info != nil {
		block, err := info.sdk.GetBlockByHeight(height)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return block, nil
		}
	}
	return nil, fmt.Errorf("all node is not working")
}

func (pro *PolySDKPro) GetSmartContractEventByBlock(height uint64) ([]*common.SmartContactEvent, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}
	for info != nil {
		event, err := info.sdk.GetSmartContractEventByBlock(height)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return event, nil
		}
	}
	return nil, fmt.Errorf("all node is not working")
}
