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
	"github.com/beego/beego/v2/core/logs"
	"github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/common"
	common1 "github.com/ontio/ontology/common"
	"github.com/ontio/ontology/core/types"
	"math/big"
	"runtime/debug"
	"sync"
	"time"
)

type OntologyInfo struct {
	sdk          *ontology_go_sdk.OntologySdk
	latestHeight uint64
}

func NewOntologyInfo(url string) *OntologyInfo {
	sdk := ontology_go_sdk.NewOntologySdk()
	sdk.NewRpcClient().SetAddress(url)
	return &OntologyInfo{
		sdk:          sdk,
		latestHeight: 0,
	}
}

type OntologySdkPro struct {
	infos         map[string]*OntologyInfo
	selectionSlot uint64
	id            uint64
	mutex         sync.Mutex
}

func NewOntologySdkPro(urls []string, slot uint64, id uint64) *OntologySdkPro {
	infos := make(map[string]*OntologyInfo, len(urls))
	for _, url := range urls {
		infos[url] = NewOntologyInfo(url)
	}
	pro := &OntologySdkPro{infos: infos, selectionSlot: slot, id: id}
	pro.selection()
	go pro.NodeSelection()
	return pro
}

func (pro *OntologySdkPro) NodeSelection() {
	for {
		pro.nodeSelection()
	}
}

func (pro *OntologySdkPro) nodeSelection() {
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

func (pro *OntologySdkPro) selection() {
	for url, info := range pro.infos {
		height, err := info.sdk.GetCurrentBlockHeight()
		if err != nil {
			logs.Error("get current block height err,chain : %v, url: %s", pro.id, url)
		}
		pro.mutex.Lock()
		info.latestHeight = uint64(height)
		pro.mutex.Unlock()
	}
}

func (pro *OntologySdkPro) GetLatest() *OntologyInfo {
	pro.mutex.Lock()
	defer func() {
		pro.mutex.Unlock()
	}()
	height := uint64(0)
	var latestInfo *OntologyInfo = nil
	for _, info := range pro.infos {
		if info != nil && info.latestHeight > height {
			height = info.latestHeight
			latestInfo = info
		}
	}
	return latestInfo
}

func (pro *OntologySdkPro) GetCurrentBlockHeight() (uint64, error) {
	info := pro.GetLatest()
	if info == nil {
		return 0, fmt.Errorf("all node is not working")
	}
	return info.latestHeight, nil
}

func (pro *OntologySdkPro) GetBlockByHeight(height uint32) (*types.Block, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}
	return info.sdk.GetBlockByHeight(height)
}

func (pro *OntologySdkPro) GetSmartContractEventByBlock(height uint32) ([]*common.SmartContactEvent, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}
	return info.sdk.GetSmartContractEventByBlock(height)
}

func (pro *OntologySdkPro) GetSdk() (*ontology_go_sdk.OntologySdk, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}
	return info.sdk, nil
}

func (pro *OntologySdkPro) IsOngAddress(hash string) bool {
	if hash == "0200000000000000000000000000000000000000" {
		return true
	} else {
		return false
	}
}

func (pro *OntologySdkPro) Oep4Balance(hash string, addr string) (*big.Int, error) {
	info := pro.GetLatest()
	if info == nil {
		return new(big.Int).SetUint64(0), fmt.Errorf("all node is not working")
	}
	if pro.IsOngAddress(hash) {
		address, err := common1.AddressFromHexString(addr)
		if err != nil {
			return new(big.Int).SetUint64(0), err
		}
		balance, err := info.sdk.Native.Ong.BalanceOf(address)
		if err != nil {
			return new(big.Int).SetUint64(0), err
		}
		return new(big.Int).SetUint64(balance), nil
	} else {
		contractAddr, err := common1.AddressFromHexString(hash)
		if err != nil {
			return new(big.Int).SetUint64(0), err
		}
		address, err := common1.AddressFromHexString(addr)
		if err != nil {
			return new(big.Int).SetUint64(0), err
		}
		preResult, err := info.sdk.NeoVM.PreExecInvokeNeoVMContract(contractAddr,
			[]interface{}{"balanceOf", []interface{}{address[:]}})
		if err != nil {
			return new(big.Int).SetUint64(0), err
		}
		balance, err := preResult.Result.ToInteger()
		if err != nil {
			return new(big.Int).SetUint64(0), err
		}
		return balance, nil
	}
}

func (pro *OntologySdkPro) Oep4TotalSupply(hash string, addr string) (*big.Int, error) {
	info := pro.GetLatest()
	if info == nil {
		return new(big.Int).SetUint64(0), fmt.Errorf("all node is not working")
	}
	if pro.IsOngAddress(hash) {
		return new(big.Int).SetUint64(0), nil
	} else {
		contractAddr, err := common1.AddressFromHexString(hash)
		if err != nil {
			return new(big.Int).SetUint64(0), err
		}
		address, err := common1.AddressFromHexString(addr)
		if err != nil {
			return new(big.Int).SetUint64(0), err
		}
		preResult, err := info.sdk.NeoVM.PreExecInvokeNeoVMContract(contractAddr,
			[]interface{}{"totalSupply", []interface{}{address[:]}})
		if err != nil {
			return new(big.Int).SetUint64(0), err
		}
		totalSupply, err := preResult.Result.ToInteger()
		if err != nil {
			return new(big.Int).SetUint64(0), err
		}
		return totalSupply, nil
	}
}

func (pro *OntologySdkPro) GetTransaction(txHash string) (*types.Transaction, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}
	return info.sdk.GetTransaction(txHash)
}
