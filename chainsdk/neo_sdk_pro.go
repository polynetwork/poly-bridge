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
	"github.com/joeqian10/neo-gogogo/rpc/models"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

type NeoInfo struct {
	sdk          *NeoSdk
	latestHeight uint64
}

func NewNeoInfo(url string) *NeoInfo {
	sdk := NewNeoSdk(url)
	return &NeoInfo{
		sdk:          sdk,
		latestHeight: 0,
	}
}

type NeoSdkPro struct {
	infos         map[string]*NeoInfo
	selectionSlot uint64
	id            uint64
	mutex         sync.Mutex
}

func NewNeoSdkPro(urls []string, slot uint64, id uint64) *NeoSdkPro {
	infos := make(map[string]*NeoInfo, len(urls))
	for _, url := range urls {
		infos[url] = NewNeoInfo(url)
	}
	pro := &NeoSdkPro{infos: infos, selectionSlot: slot, id: id}
	pro.selection()
	go pro.NodeSelection()
	return pro
}

func (pro *NeoSdkPro) NodeSelection() {
	for {
		pro.nodeSelection()
	}
}

func (pro *NeoSdkPro) nodeSelection() {
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

func (pro *NeoSdkPro) selection() {
	for url, info := range pro.infos {
		height, err := info.sdk.GetBlockCount()
		if err != nil {
			logs.Error("get current block height err: %v, url: %s", err, url)
		}
		pro.mutex.Lock()
		info.latestHeight = height
		pro.mutex.Unlock()
	}
}

func (pro *NeoSdkPro) GetLatest() *NeoInfo {
	pro.mutex.Lock()
	defer func() {
		pro.mutex.Unlock()
	}()
	height := uint64(0)
	var latestInfo *NeoInfo = nil
	for _, info := range pro.infos {
		if info != nil && info.latestHeight > height {
			height = info.latestHeight
			latestInfo = info
		}
	}
	return latestInfo
}

func (pro *NeoSdkPro) GetBlockCount() (uint64, error) {
	info := pro.GetLatest()
	if info == nil {
		return 0, fmt.Errorf("all node is not working")
	}
	return info.latestHeight, nil
}

func (pro *NeoSdkPro) GetBlockByIndex(index uint64) (*models.RpcBlock, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}
	for info != nil {
		block, err := info.sdk.GetBlockByIndex(index)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return block, nil
		}
	}
	return nil, fmt.Errorf("all node is not working")
}

func (pro *NeoSdkPro) GetApplicationLog(txId string) (*models.RpcApplicationLog, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}
	for info != nil {
		log, err := info.sdk.GetApplicationLog(txId)
		if err != nil && !strings.Contains(err.Error(), "json: cannot") {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return log, nil
		}
	}
	return nil, fmt.Errorf("all node is not working")
}

func (pro *NeoSdkPro) Nep5Info(hash string) (string, string, int64, error) {
	info := pro.GetLatest()
	if info == nil {
		return "", "", 0, fmt.Errorf("all node is not working")
	}
	for info != nil {
		hash, name, decimal, err := info.sdk.Nep5Info(hash)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return hash, name, decimal, nil
		}
	}
	return "", "", 0, fmt.Errorf("all node is not working")
}

func (pro *NeoSdkPro) GetTransactionHeight(hash string) (uint64, error) {
	info := pro.GetLatest()
	if info == nil {
		return 0, fmt.Errorf("all node is not working")
	}
	for info != nil {
		height, err := info.sdk.GetTransactionHeight(hash)
		if err != nil || height == 0 {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return height, nil
		}
	}
	return 0, fmt.Errorf("all node is not working")
}

func (pro *NeoSdkPro) SendRawTransaction(txHex string) (bool, error) {
	info := pro.GetLatest()
	if info == nil {
		return false, fmt.Errorf("all node is not working")
	}
	for info != nil {
		result, err := info.sdk.SendRawTransaction(txHex)
		if err != nil || !result {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return result, nil
		}
	}
	return false, fmt.Errorf("all node is not working")
}

func (pro *NeoSdkPro) WaitTransactionConfirm(hash string) bool {
	num := 0
	for num < 150 {
		time.Sleep(time.Second * 2)
		height, err := pro.GetTransactionHeight(hash)
		if err != nil || height == 0 {
			num++
			continue
		}
		return true
	}
	return false
}
