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
	"math/big"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/joeqian10/neo3-gogogo/rpc/models"
)

type Neo3Info struct {
	sdk          *Neo3Sdk
	latestHeight uint64
}

func NewNeo3Info(url string) *Neo3Info {
	sdk := NewNeo3Sdk(url)
	return &Neo3Info{
		sdk:          sdk,
		latestHeight: 0,
	}
}

type Neo3SdkPro struct {
	infos         map[string]*Neo3Info
	selectionSlot uint64
	id            uint64
	mutex         sync.Mutex
}

func NewNeo3SdkPro(urls []string, slot uint64, id uint64) *Neo3SdkPro {
	infos := make(map[string]*Neo3Info, len(urls))
	for _, url := range urls {
		infos[url] = NewNeo3Info(url)
	}
	pro := &Neo3SdkPro{infos: infos, selectionSlot: slot, id: id}
	pro.selection()
	go pro.NodeSelection()
	return pro
}

func (pro *Neo3SdkPro) NodeSelection() {
	for {
		pro.nodeSelection()
	}
}

func (pro *Neo3SdkPro) nodeSelection() {
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

func (pro *Neo3SdkPro) selection() {
	for url, info := range pro.infos {
		height, err := info.sdk.GetBlockCount()
		if err != nil {
			logs.Error("get current block height err, chain: %v, url: %s", pro.id, url)
		}
		pro.mutex.Lock()
		info.latestHeight = height
		pro.mutex.Unlock()
	}
}

func (pro *Neo3SdkPro) GetLatest() *Neo3Info {
	pro.mutex.Lock()
	defer func() {
		pro.mutex.Unlock()
	}()
	height := uint64(0)
	var latestInfo *Neo3Info = nil
	for _, info := range pro.infos {
		if info != nil && info.latestHeight > height {
			height = info.latestHeight
			latestInfo = info
		}
	}
	return latestInfo
}
func (pro *Neo3SdkPro) reset(info *Neo3Info) *Neo3Info {
	info.latestHeight = 0
	return pro.GetLatest()
}

func (pro *Neo3SdkPro) GetBlockCount() (uint64, error) {
	info := pro.GetLatest()
	if info == nil {
		return 0, fmt.Errorf("all node is not working")
	}
	return info.latestHeight, nil
}

func (pro *Neo3SdkPro) GetBlockByIndex(index uint64) (*models.RpcBlock, error) {
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

func (pro *Neo3SdkPro) GetApplicationLog(txId string) (*models.RpcApplicationLog, error) {
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

func (pro *Neo3SdkPro) Nep17Info(hash string) (string, string, int64, error) {
	info := pro.GetLatest()
	if info == nil {
		return "", "", 0, fmt.Errorf("all node is not working")
	}
	for info != nil {
		hash, name, decimal, err := info.sdk.Nep17Info(hash)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return hash, name, decimal, nil
		}
	}
	return "", "", 0, fmt.Errorf("all node is not working")
}

func (pro *Neo3SdkPro) Nep17Balance(hash string, addr string) (*big.Int, error) {
	info := pro.GetLatest()
	if info == nil {
		logs.Info("Nep17Balance: info is nil")
		return new(big.Int).SetUint64(0), fmt.Errorf("all node is not working")
	}
	for info != nil {
		balance, err := info.sdk.Nep17Balance(hash, addr)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return balance, nil
		}
	}
	return new(big.Int).SetUint64(0), fmt.Errorf("all node is not working")
}

func (pro *Neo3SdkPro) Nep17TotalSupply(hash string) (*big.Int, error) {
	info := pro.GetLatest()
	if info == nil {
		return new(big.Int).SetUint64(0), fmt.Errorf("all node is not working")
	}
	for info != nil {
		totalSupply, err := info.sdk.Nep17TotalSupply(hash)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return totalSupply, nil
		}
	}
	return new(big.Int).SetUint64(0), fmt.Errorf("all node is not working")
}

func (pro *Neo3SdkPro) GetTransactionHeight(hash string) (uint64, error) {
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

func (pro *Neo3SdkPro) SendRawTransaction(txHex string) (bool, error) {
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

func (pro *Neo3SdkPro) WaitTransactionConfirm(hash string) bool {
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
