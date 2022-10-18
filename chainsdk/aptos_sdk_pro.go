package chainsdk

import (
	"context"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/portto/aptos-go-sdk/client"
	"github.com/portto/aptos-go-sdk/models"
	"runtime/debug"
	"sync"
	"time"
)

type AptosInfo struct {
	Sdk          *AptosSdk
	latestHeight uint64
}

func NewAptosInfo(url string) *AptosInfo {
	sdk := NewAptosSdk(url)
	return &AptosInfo{
		Sdk:          sdk,
		latestHeight: 0,
	}
}

type AptosSdkPro struct {
	infos         map[string]*AptosInfo
	selectionSlot uint64
	id            uint64
	mutex         sync.Mutex
	ctx           context.Context
}

func NewAptosSdkPro(urls []string, slot uint64, id uint64) *AptosSdkPro {
	infos := make(map[string]*AptosInfo, len(urls))
	for _, url := range urls {
		infos[url] = NewAptosInfo(url)
	}
	pro := &AptosSdkPro{infos: infos, selectionSlot: slot, id: id, ctx: context.Background()}
	pro.selection()
	go pro.NodeSelection()
	return pro
}

func (pro *AptosSdkPro) NodeSelection() {
	for {
		pro.nodeSelection()
	}
}

func (pro *AptosSdkPro) nodeSelection() {
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

func (pro *AptosSdkPro) selection() {
	for url, info := range pro.infos {
		height, err := info.Sdk.GetCurrentBlockHeight()
		if err != nil {
			logs.Error("get current block height err: %v, url: %s", err, url)
		}
		pro.mutex.Lock()
		info.latestHeight = height
		pro.mutex.Unlock()
	}
}

func (pro *AptosSdkPro) GetLatest() *AptosInfo {
	pro.mutex.Lock()
	defer func() {
		pro.mutex.Unlock()
	}()
	height := uint64(0)
	var latestInfo *AptosInfo = nil
	for _, info := range pro.infos {
		if info != nil && info.latestHeight > height {
			height = info.latestHeight
			latestInfo = info
		}
	}
	return latestInfo
}

func (pro *AptosSdkPro) GetBlockCount() (uint64, error) {
	info := pro.GetLatest()
	if info == nil {
		return 0, fmt.Errorf("all nodes are not working")
	}
	return info.latestHeight, nil
}

func (pro *AptosSdkPro) GetBlockByIndex(index uint64) (*client.Block, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all nodes are not working")
	}
	for info != nil {
		block, err := info.Sdk.GetBlockByNumber(index)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return block, nil
		}
	}
	return nil, fmt.Errorf("all nodes are not working")
}

func (pro *AptosSdkPro) GetEvents(filter *AptosEventFilter) ([]models.Event, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all nodes are not working")
	}
	for info != nil {
		events, err := info.Sdk.GetEvents(pro.ctx, filter)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return events, nil
		}
	}
	return nil, fmt.Errorf("all nodes are not working")
}

func (pro *AptosSdkPro) GetTxByVersion(version uint64) (*client.TransactionResp, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all nodes are not working")
	}
	for info != nil {
		tx, err := info.Sdk.GetTxByVersion(pro.ctx, version)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return tx, nil
		}
	}
	return nil, fmt.Errorf("all nodes are not working")
}

func (pro *AptosSdkPro) GetBlockByVersion(version uint64) (*client.Block, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all nodes are not working")
	}
	for info != nil {
		block, err := info.Sdk.GetBlockByVersion(pro.ctx, version)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return block, nil
		}
	}
	return nil, fmt.Errorf("all nodes are not working")
}

func (pro *AptosSdkPro) GetGasPrice() (uint64, error) {
	info := pro.GetLatest()
	if info == nil {
		return 0, fmt.Errorf("all nodes are not working")
	}
	for info != nil {
		gasPrice, err := info.Sdk.GetGasPrice(pro.ctx)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return gasPrice, nil
		}
	}
	return 0, fmt.Errorf("all nodes are not working")
}

//func (pro *AptosSdkPro) GetTransactionInfoByHash(hash string) (*client.TransactionInfo, error) {
//	info := pro.GetLatest()
//	if info == nil {
//		return nil, fmt.Errorf("all nodes are not working")
//	}
//	for info != nil {
//		tx, err := info.sdk.GetTransactionInfoByHash(hash)
//		if err != nil {
//			info.latestHeight = 0
//			info = pro.GetLatest()
//		} else {
//			return tx, nil
//		}
//	}
//	return nil, fmt.Errorf("all nodes are not working")
//}
//
//func (pro *AptosSdkPro) GetGasPrice() (int, error) {
//	info := pro.GetLatest()
//	if info == nil {
//		return 1, fmt.Errorf("all nodes are not working")
//	}
//
//	for info != nil {
//		gasPrice, err := info.sdk.GetGasPrice()
//		if err != nil {
//			info.latestHeight = 0
//			info = pro.GetLatest()
//		} else {
//			return gasPrice, nil
//		}
//	}
//	return 1, fmt.Errorf("all nodes are not working")
//
//}
//
//func (pro *AptosSdkPro) GetBalance(tokenHash string, genesisAccountAddress string) (*big.Int, error) {
//	info := pro.GetLatest()
//	if info == nil {
//		return new(big.Int).SetUint64(0), fmt.Errorf("all node is not working")
//	}
//	for info != nil {
//		balance := new(big.Int).SetUint64(0)
//
//		balance, err := info.sdk.GetBalance(tokenHash, genesisAccountAddress)
//		if err != nil {
//			logs.Error("starcoin GetBalance [token hash=%s, genesisAccountAddress=%s] err=%s", tokenHash, genesisAccountAddress, err)
//			info.latestHeight = 0
//			info = pro.GetLatest()
//		} else {
//			return balance, nil
//		}
//	}
//	return new(big.Int).SetUint64(0), fmt.Errorf("all node is not working")
//}
