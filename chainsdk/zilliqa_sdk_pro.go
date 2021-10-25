package chainsdk

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/beego/beego/v2/core/logs"
	"math/big"
	"sync"
)


type ZilliqaInfo struct {
	sdk *ZilliqaSdk
	latestHeight uint64
}

func NewZilliqaInfo(url string) (*ZilliqaInfo) {
	sdk := NewZilliqaSdk(url)
	return &ZilliqaInfo{
		sdk:          sdk,
		latestHeight: 0,
	}
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
			logs.Error("get current block height err: %v, url: %s", err, url)
		}
		pro.mutex.Lock()
		info.latestHeight = height
		pro.mutex.Unlock()
	}
}

func (pro *ZilliqaSdkPro) GetLatestHeight() (uint64, error) {
	info := pro.GetLatest()
	if info == nil {
		return 0, fmt.Errorf("all node is not working")
	}
	return info.latestHeight, nil
}


func (client *ZilliqaSdkPro) GetCurrentBlockHeight() (uint64, error) {
	client.client.
	var result hexutil.Big
	err := ec.rpcClient.CallContext(context.Background(), &result, "eth_blockNumber")
	for err != nil {
		return 0, err
	}
	return (*big.Int)(&result).Uint64(), err
}
