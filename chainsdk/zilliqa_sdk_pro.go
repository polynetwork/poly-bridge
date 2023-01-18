package chainsdk

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"math/big"
	"runtime/debug"
	"sync"
	"time"
)

type ZilliqaInfo struct {
	sdk          *ZilliqaSdk
	latestHeight uint64
}

func NewZilliqaInfo(url string) *ZilliqaInfo {
	sdk := NewZilliqaSdk(url)
	return &ZilliqaInfo{
		sdk:          sdk,
		latestHeight: 0,
	}
}

func (info *ZilliqaInfo) GetLastHeight() (uint64, error) {
	return info.sdk.GetCurrentBlockHeight()
}

type ZilliqaSdkPro struct {
	infos         map[string]*ZilliqaInfo
	selectionSlot uint64
	id            uint64
	mutex         sync.Mutex
}

func NewZilliqaSdkPro(urls []string, slot uint64, id uint64) *ZilliqaSdkPro {
	infos := make(map[string]*ZilliqaInfo, len(urls))
	for _, url := range urls {
		infos[url] = NewZilliqaInfo(url)
	}
	pro := &ZilliqaSdkPro{infos: infos, selectionSlot: slot, id: id}
	pro.selection()
	go pro.NodeSelection()
	return pro
}

func (pro *ZilliqaSdkPro) selection() {
	for url, info := range pro.infos {
		height, err := info.sdk.GetCurrentBlockHeight()
		if err != nil {
			logs.Error("get current block height err,chain : %v, url: %s", pro.id, url)
		}
		pro.mutex.Lock()
		info.latestHeight = height
		pro.mutex.Unlock()
	}
}

func (pro *ZilliqaSdkPro) NodeSelection() {
	for {
		pro.nodeSelection()
	}
}

func (pro *ZilliqaSdkPro) nodeSelection() {
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

func (pro *ZilliqaSdkPro) GetLatestHeight() (uint64, error) {
	info := pro.GetLatest()
	if info == nil {
		return 0, fmt.Errorf("all node is not working")
	}
	return info.latestHeight, nil
}

func (pro *ZilliqaSdkPro) GetLatest() *ZilliqaInfo {
	pro.mutex.Lock()
	defer func() {
		pro.mutex.Unlock()
	}()
	height := uint64(0)
	var latestInfo *ZilliqaInfo = nil
	for _, info := range pro.infos {
		if info != nil && info.latestHeight > height {
			height = info.latestHeight
			latestInfo = info
		}
	}
	return latestInfo
}

func (pro *ZilliqaSdkPro) GetInfoByHeight(height uint64) (*ZilliqaInfo, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}
	for info != nil {
		index := height
		_, err := info.sdk.GetBlock(index)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return info, nil
		}
	}
	return nil, fmt.Errorf("all node is not working")
}

func (pro *ZilliqaSdkPro) GetBlockByHeight(height uint64) (*ZilBlock, error) {
	info, err := pro.GetInfoByHeight(height)
	if err != nil {
		logs.Error("GetInfoByHeight err: %v", err)
		return nil, fmt.Errorf("Zilliqa all node is not working")
	}
	block, _ := info.sdk.GetBlock(height)
	return block, nil
}

func (pro *ZilliqaSdkPro) Erc20Balance(tokenhash, addrhash string) (*big.Int, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}
	var err error
	for i := 0; i < 3; i++ {
		if info != nil {
			balance, err := info.sdk.GetTokenBalance(tokenhash, addrhash)
			if err != nil {
				info.latestHeight = 0
				info = pro.GetLatest()
			} else {
				return balance, nil
			}
		} else {
			info = pro.GetLatest()
		}
	}
	return new(big.Int).SetUint64(0), err
}

func (pro *ZilliqaSdkPro) GetMinimumGasPrice() (string, error) {
	info := pro.GetLatest()
	if info == nil {
		return "", fmt.Errorf("all node is not working")
	}
	return info.sdk.GetMinimumGasPrice()
}
