package chainsdk

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/rubblelabs/ripple/data"
	"math/big"
	"runtime/debug"
	"sync"
	"time"
)

type RippleInfo struct {
	sdk          *RippleSdk
	latestHeight uint64
}

func NewRippleInfo(url string) *RippleInfo {
	sdk := NewRippleSdk(url)
	return &RippleInfo{
		sdk:          sdk,
		latestHeight: 0,
	}
}

func (info *RippleInfo) GetLastHeight() (uint64, error) {
	return info.sdk.GetCurrentBlockHeight()
}

type RippleSdkPro struct {
	infos         map[string]*RippleInfo
	selectionSlot uint64
	id            uint64
	mutex         sync.Mutex
}

func NewRippleSdkPro(urls []string, slot uint64, id uint64) *RippleSdkPro {
	infos := make(map[string]*RippleInfo, len(urls))
	for _, url := range urls {
		infos[url] = NewRippleInfo(url)
	}
	pro := &RippleSdkPro{infos: infos, selectionSlot: slot, id: id}
	pro.selection()
	go pro.NodeSelection()
	return pro
}

func (pro *RippleSdkPro) selection() {
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

func (pro *RippleSdkPro) NodeSelection() {
	for {
		pro.nodeSelection()
	}
}

func (pro *RippleSdkPro) nodeSelection() {
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

func (pro *RippleSdkPro) GetLatestHeight() (uint64, error) {
	info := pro.GetLatest()
	if info == nil {
		return 0, fmt.Errorf("chain %v all node is not working", pro.id)
	}
	return info.latestHeight, nil
}

func (pro *RippleSdkPro) GetLatest() *RippleInfo {
	pro.mutex.Lock()
	defer func() {
		pro.mutex.Unlock()
	}()
	height := uint64(0)
	var latestInfo *RippleInfo = nil
	for _, info := range pro.infos {
		if info != nil && info.latestHeight > height {
			height = info.latestHeight
			latestInfo = info
		}
	}
	return latestInfo
}

func (pro *RippleSdkPro) GetTransationsByHeight(height uint64) (data.TransactionSlice, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("chain %v all node is not working", pro.id)
	}
	for info != nil {
		ledger, err := info.sdk.GetLedger(height)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
			time.Sleep(time.Second * time.Duration(pro.selectionSlot))
		} else {
			return ledger.Ledger.Transactions, nil
		}
	}
	return nil, fmt.Errorf("chain %v all node is not working", pro.id)
}

func (pro *RippleSdkPro) XRPBalance(tokenhash, addrhash string) (*big.Int, error) {
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

func (pro *RippleSdkPro) GetMinimumGasPrice() (string, error) {
	info := pro.GetLatest()
	if info == nil {
		return "", fmt.Errorf("all node is not working")
	}
	return "", nil
}
