package zilliqamonitor

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

type ZilliqaMonitor struct {
	monitorConfig *conf.HealthMonitorConfig
	sdks          map[string]*chainsdk.ZilliqaSdk
	nodeHeight    map[string]uint64
	nodeStatus    map[string]string
}

func NewZilliqaHealthMonitor(monitorConfig *conf.HealthMonitorConfig) *ZilliqaMonitor {
	zilliqaMonitor := &ZilliqaMonitor{}
	zilliqaMonitor.monitorConfig = monitorConfig
	sdks := make(map[string]*chainsdk.ZilliqaSdk, 0)
	for _, node := range monitorConfig.ChainNodes.Nodes {
		sdks[node.Url] = chainsdk.NewZilliqaSdk(node.Url)
	}
	zilliqaMonitor.sdks = sdks
	return zilliqaMonitor
}

func (z *ZilliqaMonitor) GetChainName() string {
	return z.monitorConfig.ChainName
}

func (z *ZilliqaMonitor) NodeMonitor() error {
	nodeStatuses := make([]basedef.NodeStatus, 0)
	for url, sdk := range z.sdks {
		status := basedef.NodeStatus{
			ChainId:   z.monitorConfig.ChainId,
			ChainName: z.monitorConfig.ChainName,
			Url:       url,
			Time:      time.Now().Format("2006-01-02 15:04:05"),
		}
		height, err := z.GetCurrentHeight(sdk)
		if err == nil {
			status.Height = height
			z.nodeHeight[url] = height
			err = z.CheckAbiCall(sdk)
		}
		if err != nil {
			z.nodeStatus[url] = err.Error()
		} else {
			z.nodeStatus[url] = "OK"
		}
		status.Status = z.nodeStatus[url]
		nodeStatuses = append(nodeStatuses, status)
	}
	data, _ := json.Marshal(nodeStatuses)
	_, err := cacheRedis.Redis.Set(cacheRedis.NodeStatusPrefix+z.monitorConfig.ChainName, data, time.Hour*24)
	if err != nil {
		logs.Error("set neo3 node status error: %s", err)
	}
	return err
}

func (z *ZilliqaMonitor) GetCurrentHeight(sdk *chainsdk.ZilliqaSdk) (uint64, error) {
	height, err := sdk.GetCurrentBlockHeight()
	if err != nil || height == 0 || height == math.MaxUint64 {
		err := fmt.Errorf("get current block height err: %s", err)
		logs.Error(fmt.Sprintf("zilliqa node: %s, %s ", sdk.GetUrl(), err))
		return 0, err
	}
	logs.Info("zilliqa node: %s, latest height: %d", sdk.GetUrl(), height)
	return height, nil
}

func (z *ZilliqaMonitor) CheckAbiCall(sdk *chainsdk.ZilliqaSdk) error {
	_, err := sdk.GetBlock(z.nodeHeight[sdk.GetUrl()] - 1)
	if err != nil {
		err := fmt.Errorf("call GetBlock error: %s", err)
		logs.Error(fmt.Sprintf("zilliqa node: %s, %s ", sdk.GetUrl(), err))
		return err
	}
	return nil
}
