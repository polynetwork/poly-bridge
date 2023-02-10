package chainsdk

import (
	"fmt"
	ctypes "github.com/tendermint/tendermint/rpc/coretypes"
	"runtime/debug"
	"sync"
	"time"

	"github.com/beego/beego/v2/core/logs"
	//ctypes "github.com/tendermint/tendermint/rpc/corety"
)

type SwitcheoInfo struct {
	sdk          *SwitcheoSDK
	latestHeight uint64
}

func NewSwitcheoInfo(url string) *SwitcheoInfo {
	sdk := NewSwitcheoSDK(url)
	return &SwitcheoInfo{
		sdk:          sdk,
		latestHeight: 0,
	}
}

type SwitcheoSdkPro struct {
	infos         map[string]*SwitcheoInfo
	selectionSlot uint64
	id            uint64
	mutex         sync.Mutex
}

func NewSwitcheoSdkPro(urls []string, slot uint64, id uint64) *SwitcheoSdkPro {
	infos := make(map[string]*SwitcheoInfo, len(urls))
	for _, url := range urls {
		infos[url] = NewSwitcheoInfo(url)
	}
	pro := &SwitcheoSdkPro{infos: infos, selectionSlot: slot, id: id}
	pro.selection()
	go pro.NodeSelection()
	return pro
}

func (pro *SwitcheoSdkPro) GetInfoByHeight(height uint64) (*SwitcheoInfo, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}
	for info != nil {
		index := int64(height)
		_, err := info.sdk.Block(&index)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return info, nil
		}
	}
	return nil, fmt.Errorf("all node is not working")
}

func (pro *SwitcheoSdkPro) GetBlockByHeight(height uint64) (*ctypes.ResultBlock, error) {
	info, err := pro.GetInfoByHeight(height)
	if err != nil {
		logs.Error("GetInfoByHeight err: %v", err)
	} else {
		index := int64(height)
		block, _ := info.sdk.Block(&index)
		return block, nil
	}
	return nil, fmt.Errorf("all node is not working")
}

func (pro *SwitcheoSdkPro) selection() {
	for url, info := range pro.infos {
		height, err := info.sdk.GetCurrentBlockHeight()
		if err != nil {
			logs.Error("get current block height err, chain: %v, url: %s, err: %v", pro.id, url, err)
		}
		pro.mutex.Lock()
		info.latestHeight = height
		pro.mutex.Unlock()
	}
}

func (pro *SwitcheoSdkPro) NodeSelection() {
	for {
		pro.nodeSelection()
	}
}

func (pro *SwitcheoSdkPro) nodeSelection() {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("node selection, recover info: %s,  err: %s", string(debug.Stack()), r)
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

func (pro *SwitcheoSdkPro) GetLatestHeight() (uint64, error) {
	info := pro.GetLatest()
	if info == nil {
		return 0, fmt.Errorf("all node is not working")
	}
	return info.latestHeight, nil
}

func (pro *SwitcheoSdkPro) GetLatest() *SwitcheoInfo {
	pro.mutex.Lock()
	defer func() {
		pro.mutex.Unlock()
	}()
	height := uint64(0)
	var latestInfo *SwitcheoInfo = nil
	for _, info := range pro.infos {
		if info != nil && info.latestHeight > height {
			height = info.latestHeight
			latestInfo = info
		}
	}
	return latestInfo
}

func (pro *SwitcheoSdkPro) TxSearch(height uint64, query string, prove bool, page, perPage int, orderBy string) (*ctypes.ResultTxSearch, error) {
	info, _ := pro.GetInfoByHeight(height)
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}
	return info.sdk.TxSearch(query, prove, page, perPage, orderBy)

}
