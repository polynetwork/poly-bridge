package ripplemonitor

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"math"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"time"
)

type RippleMonitor struct {
	monitorConfig *conf.HealthMonitorConfig
	sdks          map[string]*chainsdk.RippleSdk
	nodeHeight    map[string]uint64
}

func NewRippleHealthMonitor(monitorConfig *conf.HealthMonitorConfig) *RippleMonitor {
	rippleMonitor := &RippleMonitor{}
	rippleMonitor.monitorConfig = monitorConfig
	sdks := make(map[string]*chainsdk.RippleSdk, 0)
	for _, node := range monitorConfig.ChainNodes.Nodes {
		sdks[node.Url] = chainsdk.NewRippleSdk(node.Url)
	}
	rippleMonitor.sdks = sdks
	rippleMonitor.nodeHeight = make(map[string]uint64, len(sdks))
	return rippleMonitor
}

func (r *RippleMonitor) GetChainName() string {
	return r.monitorConfig.ChainName
}

func (r *RippleMonitor) GetChainId() uint64 {
	return r.monitorConfig.ChainId
}

func (r *RippleMonitor) RelayerBalanceMonitor() ([]*basedef.RelayerAccountStatus, error) {
	return nil, nil
}

func (r *RippleMonitor) NodeMonitor() ([]basedef.NodeStatus, error) {
	nodeStatuses := make([]basedef.NodeStatus, 0)
	for url, sdk := range r.sdks {
		status := basedef.NodeStatus{
			ChainId:   r.monitorConfig.ChainId,
			ChainName: r.monitorConfig.ChainName,
			Url:       url,
			Status:    make([]string, 0),
			Time:      time.Now().Unix(),
		}
		height, err := r.GetCurrentHeight(sdk)
		if err == nil {
			status.Height = height
			r.nodeHeight[url] = height
			err = r.CheckAbiCall(sdk)
		}
		if err != nil {
			status.Status = append(status.Status, err.Error())
		}
		nodeStatuses = append(nodeStatuses, status)
	}
	return nodeStatuses, nil
}

func (r *RippleMonitor) GetCurrentHeight(sdk *chainsdk.RippleSdk) (uint64, error) {
	height, err := sdk.GetCurrentBlockHeight()
	if err != nil || height == 0 || height == math.MaxUint64 {
		e := fmt.Errorf("get current block height err: %s", err)
		logs.Error(fmt.Sprintf("%s node: %s, %s ", r.GetChainName(), sdk.GetUrl(), e))
		return 0, e
	}
	logs.Info("%s node: %s, latest height: %d", r.GetChainName(), sdk.GetUrl(), height)
	return height, nil
}

func (r *RippleMonitor) CheckAbiCall(sdk *chainsdk.RippleSdk) error {
	_, err := sdk.GetLedger(r.nodeHeight[sdk.GetUrl()] - 1)
	if err != nil {
		e := fmt.Errorf("call GetLedger error: %s", err)
		logs.Error(fmt.Sprintf("%s node: %s, %s ", r.GetChainName(), sdk.GetUrl(), e))
		return e
	}
	return nil
}
