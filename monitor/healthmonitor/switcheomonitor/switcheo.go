package switcheomonitor

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"math"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"time"
)

type SwitcheoMonitor struct {
	monitorConfig *conf.HealthMonitorConfig
	sdks          map[string]*chainsdk.SwitcheoSDK
	nodeHeight    map[string]uint64
}

func NewSwitcheoHealthMonitor(monitorConfig *conf.HealthMonitorConfig) *SwitcheoMonitor {
	switcheoMonitor := &SwitcheoMonitor{}
	switcheoMonitor.monitorConfig = monitorConfig
	sdks := make(map[string]*chainsdk.SwitcheoSDK, 0)
	for _, node := range monitorConfig.ChainNodes.Nodes {
		sdks[node] = chainsdk.NewSwitcheoSDK(node)
	}
	switcheoMonitor.sdks = sdks
	switcheoMonitor.nodeHeight = make(map[string]uint64, len(sdks))
	return switcheoMonitor
}

func (s *SwitcheoMonitor) GetChainName() string {
	return s.monitorConfig.ChainName
}

func (s *SwitcheoMonitor) GetChainId() uint64 {
	return s.monitorConfig.ChainId
}

func (s *SwitcheoMonitor) RelayerBalanceMonitor() ([]*basedef.RelayerAccountStatus, error) {
	return nil, nil
}

func (s *SwitcheoMonitor) NodeMonitor() ([]basedef.NodeStatus, error) {
	nodeStatuses := make([]basedef.NodeStatus, 0)
	for url, sdk := range s.sdks {
		status := basedef.NodeStatus{
			ChainId:   s.monitorConfig.ChainId,
			ChainName: s.monitorConfig.ChainName,
			Url:       url,
			Status:    make([]string, 0),
			Time:      time.Now().Unix(),
		}
		height, err := s.GetCurrentHeight(sdk, url)
		if err == nil {
			status.Height = height
			s.nodeHeight[url] = height
			err = s.CheckAbiCall(sdk, url)
		}
		if err != nil {
			status.Status = append(status.Status, err.Error())
		}
		nodeStatuses = append(nodeStatuses, status)
	}
	//data, _ := json.Marshal(nodeStatuses)
	//_, err := cacheRedis.Redis.Set(cacheRedis.NodeStatusPrefix+s.monitorConfig.ChainName, data, time.Hour*24)
	//if err != nil {
	//	logs.Error("set %s node status error: %s", s.GetChainName(), err)
	//}
	return nodeStatuses, nil
}

func (s *SwitcheoMonitor) GetCurrentHeight(sdk *chainsdk.SwitcheoSDK, url string) (uint64, error) {
	height, err := sdk.GetCurrentBlockHeight()
	if err != nil || height == 0 || height == math.MaxUint64 {
		e := fmt.Errorf("get current block height err: %s", err)
		logs.Error(fmt.Sprintf("%s node: %s, %s ", s.GetChainName(), url, e))
		return 0, e
	}
	logs.Info("%s node: %s, latest height: %d", s.GetChainName(), url, height)
	return height, nil
}

func (s *SwitcheoMonitor) CheckAbiCall(sdk *chainsdk.SwitcheoSDK, url string) error {
	height := s.nodeHeight[url] - 1
	index := int64(height)
	block, err := sdk.Block(&index)
	if err != nil {
		e := fmt.Errorf("call Block err: %s", err)
		logs.Error(fmt.Sprintf("%s node: %s, %s ", s.GetChainName(), url, e))
		return e
	}
	if block == nil {
		e := fmt.Errorf("there is no %s block", s.GetChainName())
		logs.Error(fmt.Sprintf("%s node: %s, %s ", s.GetChainName(), url, e))
		return e
	}

	lockQuery := fmt.Sprintf("tx.height=%d AND make_from_cosmos_proof.status='1'", height)
	_, err = sdk.TxSearch(lockQuery, false, 1, 100, "asc")
	if err != nil {
		e := fmt.Errorf("call TxSearch get lock events err: %s", err)
		logs.Error(fmt.Sprintf("%s node: %s, %s ", s.GetChainName(), url, e))
		return e
	}

	unlockQuery := fmt.Sprintf("tx.height=%d", height)
	_, err = sdk.TxSearch(unlockQuery, false, 1, 100, "asc")
	if err != nil {
		e := fmt.Errorf("call TxSearch get unlock events err: %s", err)
		logs.Error(fmt.Sprintf("%s node: %s, %s ", s.GetChainName(), url, e))
		return e
	}
	return nil
}
