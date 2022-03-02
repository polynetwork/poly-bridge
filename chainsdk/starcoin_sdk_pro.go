package chainsdk

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/starcoinorg/starcoin-go/client"
	"runtime/debug"
	"sync"
	"time"
)

type StarcoinInfo struct {
	sdk          *StarCoinSdk
	latestHeight uint64
}

func NewStarcoinInfo(url string) *StarcoinInfo {
	sdk := NewStarCoinSdk(url)
	return &StarcoinInfo{
		sdk:          sdk,
		latestHeight: 0,
	}
}

type StarcoinSdkPro struct {
	infos         map[string]*StarcoinInfo
	selectionSlot uint64
	id            uint64
	mutex         sync.Mutex
}

func NewStarcoinSdkPro(urls []string, slot uint64, id uint64) *StarcoinSdkPro {
	infos := make(map[string]*StarcoinInfo, len(urls))
	for _, url := range urls {
		infos[url] = NewStarcoinInfo(url)
	}
	pro := &StarcoinSdkPro{infos: infos, selectionSlot: slot, id: id}
	pro.selection()
	go pro.NodeSelection()
	return pro
}

func (pro *StarcoinSdkPro) NodeSelection() {
	for {
		pro.nodeSelection()
	}
}

func (pro *StarcoinSdkPro) nodeSelection() {
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

func (pro *StarcoinSdkPro) selection() {
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

func (pro *StarcoinSdkPro) GetLatest() *StarcoinInfo {
	pro.mutex.Lock()
	defer func() {
		pro.mutex.Unlock()
	}()
	height := uint64(0)
	var latestInfo *StarcoinInfo = nil
	for _, info := range pro.infos {
		if info != nil && info.latestHeight > height {
			height = info.latestHeight
			latestInfo = info
		}
	}
	return latestInfo
}

func (pro *StarcoinSdkPro) GetBlockCount() (uint64, error) {
	info := pro.GetLatest()
	if info == nil {
		return 0, fmt.Errorf("all nodes are not working")
	}
	return info.latestHeight, nil
}

func (pro *StarcoinSdkPro) GetBlockByIndex(index uint64) (*client.Block, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all nodes are not working")
	}
	for info != nil {
		block, err := info.sdk.GetBlockByNumber(index)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return block, nil
		}
	}
	return nil, fmt.Errorf("all nodes are not working")
}

func (pro *StarcoinSdkPro) GetEvents(filter *client.EventFilter) ([]client.Event, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all nodes are not working")
	}
	for info != nil {
		events, err := info.sdk.GetEvents(filter)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return events, nil
		}
	}
	return nil, fmt.Errorf("all nodes are not working")
}

func (pro *StarcoinSdkPro) GetTransactionInfoByHash(hash string) (*client.TransactionInfo, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all nodes are not working")
	}
	for info != nil {
		tx, err := info.sdk.GetTransactionInfoByHash(hash)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return tx, nil
		}
	}
	return nil, fmt.Errorf("all nodes are not working")
}

func (pro *StarcoinSdkPro) GetGasPrice() (int, error) {
	info := pro.GetLatest()
	if info == nil {
		return 1, fmt.Errorf("all nodes are not working")
	}

	for info != nil {
		gasPrice, err := info.sdk.GetGasPrice()
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return gasPrice, nil
		}
	}
	return 1, fmt.Errorf("all nodes are not working")

}
