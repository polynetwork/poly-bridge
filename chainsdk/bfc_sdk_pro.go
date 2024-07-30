package chainsdk

import (
	"context"
	"github.com/beego/beego/v2/adapter/logs"
	suimodels "github.com/block-vision/sui-go-sdk/models"
	suisdk "github.com/block-vision/sui-go-sdk/sui"
	"runtime/debug"
	"strconv"
	"sync"
	"time"
)

type BfcInfo struct {
	Sdk          suisdk.ISuiAPI
	LatestHeight uint64
}

type BfcSdkPro struct {
	infos         map[string]*BfcInfo
	id            uint64
	selectionSlot uint64
	mutex         sync.Mutex
	ctx           context.Context
}

func NewBfcSdkPro(urls []string, slot uint64, id uint64) *BfcSdkPro {
	infos := make(map[string]*BfcInfo, 0)
	for _, url := range urls {
		sdk := suisdk.NewSuiClient(url)
		h, err := sdk.SuiGetTotalTransactionBlocks(context.Background())
		if err != nil {
			logs.Error("SuiGetTotalTransactionBlocks url: %v, err: %v", url, err)
		} else {
			infos[url] = &BfcInfo{
				Sdk:          sdk,
				LatestHeight: h,
			}
		}
	}
	pro := &BfcSdkPro{infos: infos, id: id, selectionSlot: slot, ctx: context.Background()}
	pro.selection()
	go pro.NodeSelection()
	return pro
}

func (pro *BfcSdkPro) NodeSelection() {
	for {
		pro.nodeSelection()
	}
}

func (pro *BfcSdkPro) nodeSelection() {
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

func (pro *BfcSdkPro) selection() {
	for url, info := range pro.infos {
		height, err := info.Sdk.SuiGetTotalTransactionBlocks(context.Background())
		if err != nil {
			logs.Error("get current block height chain: %v, err: %v, url: %s", pro.id, err, url)
		}
		pro.mutex.Lock()
		info.LatestHeight = height
		pro.mutex.Unlock()
	}
}

func (pro *BfcSdkPro) GetLatest() *BfcInfo {
	pro.mutex.Lock()
	defer func() {
		pro.mutex.Unlock()
	}()
	height := uint64(0)
	var latestInfo *BfcInfo
	for _, info := range pro.infos {
		if info != nil && info.LatestHeight > height {
			height = info.LatestHeight
			latestInfo = info
		}
	}
	return latestInfo
}

func (b *BfcSdkPro) QueryEvents(moveEventType, eventCursor string, batch uint64) (suimodels.PaginatedEventsResponse, error) {
	info := b.GetLatest()
	EventRequest := suimodels.SuiXQueryEventsRequest{
		SuiEventFilter: suimodels.EventFilterByMoveEventType{
			MoveEventType: moveEventType,
		},
		Limit: batch,
	}
	if eventCursor != "" {
		EventRequest.Cursor = suimodels.EventId{
			TxDigest: eventCursor,
			EventSeq: "0",
		}
	}
	return info.Sdk.SuiXQueryEvents(context.Background(), EventRequest)
}

func (b *BfcSdkPro) GetCheckpoint(txDigest string) (uint64, error) {
	info := b.GetLatest()
	tx, err := info.Sdk.SuiGetTransactionBlock(context.Background(), suimodels.SuiGetTransactionBlockRequest{
		Digest: txDigest,
	})
	if err != nil {
		return 0, err
	}
	checkpoint, _ := strconv.Atoi(tx.Checkpoint)
	return uint64(checkpoint), err
}

func (b *BfcSdkPro) GetTotalTransactionBlocks() (uint64, error) {
	info := b.GetLatest()
	return info.Sdk.SuiGetTotalTransactionBlocks(context.Background())
}
