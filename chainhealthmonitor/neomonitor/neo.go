package neomonitor

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"math"
	"poly-bridge/basedef"
	"poly-bridge/cacheRedis"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"time"
)

type NeoMonitor struct {
	monitorConfig *conf.HealthMonitorConfig
	sdks          map[string]*chainsdk.NeoSdk
	nodeHeight    map[string]uint64
	nodeStatus    map[string]string
}

func NewNeoHealthMonitor(monitorConfig *conf.HealthMonitorConfig) *NeoMonitor {
	neoMonitor := &NeoMonitor{}
	neoMonitor.monitorConfig = monitorConfig
	sdks := make(map[string]*chainsdk.NeoSdk, 0)
	for _, node := range monitorConfig.ChainNodes.Nodes {
		sdk := chainsdk.NewNeoSdk(node.Url)
		if sdk.GetClient() == nil {
			if _, err := cacheRedis.Redis.Set(cacheRedis.NodeStatusPrefix+node.Url, fmt.Sprintf("initial sdk error:sdk.client is nil"), time.Hour*24); err != nil {
				logs.Error("set neo node[%s] status error: %s", node.Url, err)
			}
			logs.Error("neo node: %s, initial sdk error:sdk.client is nil", node.Url)
			continue
		}
		sdks[node.Url] = sdk
	}
	neoMonitor.sdks = sdks
	return neoMonitor
}

func (n *NeoMonitor) GetChainName() string {
	return n.monitorConfig.ChainName
}

func (n *NeoMonitor) NodeMonitor() ([]basedef.NodeStatus, error) {
	nodeStatuses := make([]basedef.NodeStatus, 0)
	for url, sdk := range n.sdks {
		status := basedef.NodeStatus{
			ChainId:   n.monitorConfig.ChainId,
			ChainName: n.monitorConfig.ChainName,
			Url:       url,
			Time:      time.Now().Format("2006-01-02 15:04:05"),
		}
		height, err := n.GetCurrentHeight(sdk)
		if err == nil {
			status.Height = height
			n.nodeHeight[url] = height
			err = n.CheckAbiCall(sdk)
		}
		err = n.CheckAbiCall(sdk)
		if err != nil {
			n.nodeStatus[url] = err.Error()
		} else {
			n.nodeStatus[url] = basedef.NodeStatusOk
		}
		status.Status = n.nodeStatus[url]
		nodeStatuses = append(nodeStatuses, status)
	}
	data, _ := json.Marshal(nodeStatuses)
	_, err := cacheRedis.Redis.Set(cacheRedis.NodeStatusPrefix+n.monitorConfig.ChainName, data, time.Hour*24)
	if err != nil {
		logs.Error("set neo node status error: %s", err)
	}
	return nodeStatuses, err
}

func (n *NeoMonitor) GetCurrentHeight(sdk *chainsdk.NeoSdk) (uint64, error) {
	height, err := sdk.GetBlockCount()
	if err != nil || height == 0 || height == math.MaxUint64 {
		err := fmt.Errorf("get current block height err: %s", err)
		logs.Error(fmt.Sprintf("neo node: %s, %s ", sdk.GetUrl(), err))
		return 0, err
	}
	logs.Info("neo node: %s, latest height: %d", sdk.GetUrl(), height)
	return height, nil
}

func (n *NeoMonitor) CheckAbiCall(sdk *chainsdk.NeoSdk) error {
	_, err := sdk.GetBlockByIndex(n.nodeHeight[sdk.GetUrl()] - 1)
	if err != nil {
		err := fmt.Errorf("call GetBlockByIndex error: %s", err)
		logs.Error(fmt.Sprintf("neo node: %s, %s ", sdk.GetUrl(), err))
		return err
	}
	return nil
}
