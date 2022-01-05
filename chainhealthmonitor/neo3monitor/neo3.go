package neo3monitor

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

type Neo3Monitor struct {
	monitorConfig *conf.HealthMonitorConfig
	sdks          map[string]*chainsdk.Neo3Sdk
	nodeHeight    map[string]uint64
	nodeStatus    map[string]string
}

func NewNeo3HealthMonitor(monitorConfig *conf.HealthMonitorConfig) *Neo3Monitor {
	neo3Monitor := &Neo3Monitor{}
	neo3Monitor.monitorConfig = monitorConfig
	sdks := make(map[string]*chainsdk.Neo3Sdk, 0)
	for _, node := range monitorConfig.ChainNodes.Nodes {
		sdk := chainsdk.NewNeo3Sdk(node.Url)
		if sdk.GetClient() == nil {
			if _, err := cacheRedis.Redis.Set(cacheRedis.NodeStatusPrefix+node.Url, fmt.Sprintf("initial sdk error:sdk.client is nil"), time.Hour*24); err != nil {
				logs.Error("set %s node[%s] status error: %s", monitorConfig.ChainName, node.Url, err)
			}
			logs.Error("%s node: %s, initial sdk error: sdk.client is nil", monitorConfig.ChainName, node.Url)
			continue
		}
		sdks[node.Url] = sdk
	}
	neo3Monitor.sdks = sdks
	neo3Monitor.nodeHeight = make(map[string]uint64, len(sdks))
	neo3Monitor.nodeStatus = make(map[string]string, len(sdks))
	return neo3Monitor
}

func (n *Neo3Monitor) GetChainName() string {
	return n.monitorConfig.ChainName
}

func (n *Neo3Monitor) NodeMonitor() ([]basedef.NodeStatus, error) {
	nodeStatuses := make([]basedef.NodeStatus, 0)
	for url, sdk := range n.sdks {
		status := basedef.NodeStatus{
			ChainId:   n.monitorConfig.ChainId,
			ChainName: n.monitorConfig.ChainName,
			Url:       url,
			Status:    make([]string, 0),
			Time:      time.Now().Unix(),
		}
		height, err := n.GetCurrentHeight(sdk)
		if err == nil {
			status.Height = height
			n.nodeHeight[url] = height
			err = n.CheckAbiCall(sdk)
		}
		if err != nil {
			n.nodeStatus[url] = err.Error()
		} else {
			n.nodeStatus[url] = basedef.NodeStatusOk
		}
		status.Status = append(status.Status, n.nodeStatus[url])
		nodeStatuses = append(nodeStatuses, status)
	}

	data, _ := json.Marshal(nodeStatuses)
	_, err := cacheRedis.Redis.Set(cacheRedis.NodeStatusPrefix+n.monitorConfig.ChainName, data, time.Hour*24)
	if err != nil {
		logs.Error("set %s node status error: %s", n.GetChainName(), err)
	}
	return nodeStatuses, err
}

func (n *Neo3Monitor) GetCurrentHeight(sdk *chainsdk.Neo3Sdk) (uint64, error) {
	height, err := sdk.GetBlockCount()
	if err != nil || height == 0 || height == math.MaxUint64 {
		err := fmt.Errorf("get current block height err: %s", err)
		logs.Error(fmt.Sprintf("%s node: %s, %s ", n.GetChainName(), sdk.GetUrl(), err))
		return 0, err
	}
	logs.Info("%s node: %s, latest height: %d", n.GetChainName(), sdk.GetUrl(), height)
	return height, nil
}

func (n *Neo3Monitor) CheckAbiCall(sdk *chainsdk.Neo3Sdk) error {
	_, err := sdk.GetBlockByIndex(n.nodeHeight[sdk.GetUrl()] - 1)
	if err != nil {
		err := fmt.Errorf("call GetBlockByIndex error: %s", err)
		logs.Error(fmt.Sprintf("%s node: %s, %s ", n.GetChainName(), sdk.GetUrl(), err))
		return err
	}
	return nil
}
