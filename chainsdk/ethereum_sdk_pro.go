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
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math"
	"math/big"
	"runtime/debug"
	"sync"
	"time"
)

type EthereumInfo struct {
	sdk          *EthereumSdk
	latestHeight uint64
}

func NewEthereumInfo(url string) *EthereumInfo {
	sdk, err := NewEthereumSdk(url)
	if err != nil || sdk == nil {
		panic(err)
	}
	return &EthereumInfo{
		sdk:          sdk,
		latestHeight: 0,
	}
}

type EthereumSdkPro struct {
	infos         map[string]*EthereumInfo
	selectionSlot uint64
	id            uint64
	mutex         sync.Mutex
}

func NewEthereumSdkPro(urls []string, slot uint64, id uint64) *EthereumSdkPro {
	infos := make(map[string]*EthereumInfo, len(urls))
	for _, url := range urls {
		infos[url] = NewEthereumInfo(url)
	}
	pro := &EthereumSdkPro{infos: infos, selectionSlot: slot, id: id}
	pro.selection()
	go pro.NodeSelection()
	return pro
}

func (pro *EthereumSdkPro) NodeSelection() {
	for {
		pro.nodeSelection()
	}
}

func (pro *EthereumSdkPro) nodeSelection() {
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

func (pro *EthereumSdkPro) selection() {
	for url, info := range pro.infos {
		height, err := info.sdk.GetCurrentBlockHeight()
		if err != nil || height == math.MaxUint64 {
			logs.Error("get current block height err: %v, url: %s", err, url)
			height = 1
		}
		/*
			block, err := info.sdk.GetBlockByNumber(height)
			if err != nil || block == nil {
				logs.Error("get current block err: %v, url: %s", err, url)
				info.latestHeight = height - 1
				continue
			}
			transactions := block.Transactions()
			if len(transactions) > 0 {
				transaction := transactions[0]
				receipt, err := info.sdk.GetTransactionReceipt(transaction.Hash())
				if err != nil || receipt == nil {
					logs.Error("get transaction receipt err: %v, url: %s", err, url)
					info.latestHeight = height - 1
					continue
				}
			}
		*/
		pro.mutex.Lock()
		info.latestHeight = height - 1
		pro.mutex.Unlock()
	}
}

func (pro *EthereumSdkPro) GetLatest() *EthereumInfo {
	pro.mutex.Lock()
	defer func() {
		pro.mutex.Unlock()
	}()
	height := uint64(0)
	var latestInfo *EthereumInfo = nil
	for _, info := range pro.infos {
		if info != nil && info.latestHeight > height {
			height = info.latestHeight
			latestInfo = info
		}
	}
	return latestInfo
}

func (pro *EthereumSdkPro) GetClient() *ethclient.Client {
	info := pro.GetLatest()
	if info == nil {
		return nil
	}
	return info.sdk.GetClient()
}

func (pro *EthereumSdkPro) GetLatestHeight() (uint64, error) {
	info := pro.GetLatest()
	if info == nil {
		return 0, fmt.Errorf("all node is not working")
	}
	return info.latestHeight, nil
}

func (pro *EthereumSdkPro) GetHeaderByNumber(number uint64) (*types.Header, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}

	for info != nil {
		header, err := info.sdk.GetHeaderByNumber(number)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return header, nil
		}
	}
	return nil, fmt.Errorf("all node is not working")
}

func (pro *EthereumSdkPro) GetTransactionByHash(hash common.Hash) (*types.Transaction, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}

	for info != nil {
		tx, err := info.sdk.GetTransactionByHash(hash)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return tx, nil
		}
	}
	return nil, fmt.Errorf("all node is not working")
}

func (pro *EthereumSdkPro) GetTransactionReceipt(hash common.Hash) (*types.Receipt, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}

	for info != nil {
		receipt, err := info.sdk.GetTransactionReceipt(hash)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return receipt, nil
		}
	}
	return nil, fmt.Errorf("all node is not working")
}

func (pro *EthereumSdkPro) NonceAt(addr common.Address) (uint64, error) {
	info := pro.GetLatest()
	if info == nil {
		return 0, fmt.Errorf("all node is not working")
	}

	for info != nil {
		nonce, err := info.sdk.NonceAt(addr)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return nonce, nil
		}
	}
	return 0, fmt.Errorf("all node is not working")
}

func (pro *EthereumSdkPro) SendRawTransaction(tx *types.Transaction) error {
	info := pro.GetLatest()
	if info == nil {
		return fmt.Errorf("all node is not working")
	}

	for info != nil {
		err := info.sdk.SendRawTransaction(tx)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return nil
		}
	}
	return fmt.Errorf("all node is not working")
}

func (pro *EthereumSdkPro) TransactionByHash(hash common.Hash) (*types.Transaction, bool, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, false, fmt.Errorf("all node is not working")
	}

	for info != nil {
		tx, isPending, err := info.sdk.TransactionByHash(hash)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return tx, isPending, nil
		}
	}
	return nil, false, fmt.Errorf("all node is not working")
}

func (pro *EthereumSdkPro) SuggestGasPrice() (*big.Int, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}

	for info != nil {
		gas, err := info.sdk.SuggestGasPrice()
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return gas, nil
		}
	}
	return nil, fmt.Errorf("all node is not working")
}

func (pro *EthereumSdkPro) EstimateGas(msg ethereum.CallMsg) (uint64, error) {
	info := pro.GetLatest()
	if info == nil {
		return 0, fmt.Errorf("all node is not working")
	}

	for info != nil {
		gas, err := info.sdk.EstimateGas(msg)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return gas, nil
		}
	}
	return 0, fmt.Errorf("all node is not working")
}

func (pro *EthereumSdkPro) Erc20Info(hash string) (string, string, int64, string, error) {
	info := pro.GetLatest()
	if info == nil {
		return "", "", 0, "", fmt.Errorf("all node is not working")
	}
	for info != nil {
		hash, name, decimal, symbol, err := info.sdk.Erc20Info(hash)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return hash, name, decimal, symbol, nil
		}
	}
	return "", "", 0, "", fmt.Errorf("all node is not working")
}

func (pro *EthereumSdkPro) WaitTransactionConfirm(hash common.Hash) bool {
	num := 0
	for num < 300 {
		time.Sleep(time.Second * 2)
		_, ispending, err := pro.TransactionByHash(hash)
		if err != nil {
			num++
			continue
		}
		if ispending {
			num++
			continue
		} else {
			return true
		}
	}
	return false
}
