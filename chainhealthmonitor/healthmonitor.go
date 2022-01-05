package chainhealthmonitor

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"poly-bridge/basedef"
	"poly-bridge/cacheRedis"
	"poly-bridge/chainhealthmonitor/ethereummonitor"
	"poly-bridge/chainhealthmonitor/neo3monitor"
	"poly-bridge/chainhealthmonitor/neomonitor"
	"poly-bridge/chainhealthmonitor/ontologymonitor"
	"poly-bridge/chainhealthmonitor/polymonitor"
	"poly-bridge/chainhealthmonitor/switcheomonitor"
	"poly-bridge/chainhealthmonitor/zilliqamonitor"
	polycommon "poly-bridge/common"
	"poly-bridge/conf"
	"runtime/debug"
	"time"
)

var healthMonitorConfigMap = make(map[uint64]*conf.HealthMonitorConfig, 0)

func StartHealthMonitor(config *conf.Config) {
	logs.Info("StartHealthMonitor")
	for _, cfg := range config.ChainNodes {
		monitorConfig := &conf.HealthMonitorConfig{ChainId: cfg.ChainId, ChainName: cfg.ChainName, ChainNodes: cfg}
		healthMonitorConfigMap[cfg.ChainId] = monitorConfig
	}
	for _, cfg := range config.ChainListenConfig {
		healthMonitorConfigMap[cfg.ChainId].CCMContract = cfg.CCMContract
	}
	//for i, i2 := range collection {
	//	TODO RelayerAddrs
	//}

	for _, monitorConfig := range healthMonitorConfigMap {
		healthMonitorHandle := NewHealthMonitorHandle(monitorConfig)
		if healthMonitorHandle == nil {
			logs.Error("chain %s handler is invalid", monitorConfig.ChainName)
			continue
		}
		monitor := &HealthMonitor{healthMonitorHandle}
		monitor.Start(config)
	}

}

type HealthMonitorHandle interface {
	NodeMonitor() ([]basedef.NodeStatus, error)
	GetChainName() string
}

type HealthMonitor struct {
	handle HealthMonitorHandle
}

func (h *HealthMonitor) Start(config *conf.Config) {
	go h.NodeMonitor(config)
}

func (h *HealthMonitor) NodeMonitor(config *conf.Config) {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("NodeMonitor restart, recover info: %s", string(debug.Stack()))
		}
	}()

	logs.Info("start %s NodeMonitor", h.handle.GetChainName())
	nodeMonitorTicker := time.NewTicker(time.Second * time.Duration(config.ChainNodeStatusCheckInterval))
	for {
		select {
		case <-nodeMonitorTicker.C:
			oldNodeStatusMap := make(map[string]*basedef.NodeStatus)
			if dataStr, err := cacheRedis.Redis.Get(cacheRedis.NodeStatusPrefix + h.handle.GetChainName()); err == nil {
				var oldNodeStatuses []basedef.NodeStatus
				if err := json.Unmarshal([]byte(dataStr), &oldNodeStatuses); err != nil {
					logs.Error("chain %s node status data Unmarshal error: ", h.handle.GetChainName(), err)
				} else {
					for _, oldNodeStatus := range oldNodeStatuses {
						oldNodeStatusMap[oldNodeStatus.Url] = &oldNodeStatus
					}
				}
			}
			if nodeStatuses, err := h.handle.NodeMonitor(); err == nil {
				data, _ := json.Marshal(nodeStatuses)
				_, err := cacheRedis.Redis.Set(cacheRedis.NodeStatusPrefix+h.handle.GetChainName(), data, time.Hour*24)
				if err != nil {
					logs.Error("set %s node status error: %s", h.handle.GetChainName(), err)
				}

				for _, nodeStatus := range nodeStatuses {
					oldNodeStatus := oldNodeStatusMap[nodeStatus.Url]
					var nodeHeightNoGrowthTime uint64
					if oldNodeStatus != nil && nodeStatus.Height == oldNodeStatus.Height {
						nodeHeightNoGrowthTime = nodeStatus.Height - oldNodeStatus.Height
						if nodeHeightNoGrowthTime > 180 {
							nodeStatus.Status = append(nodeStatus.Status, fmt.Sprintf("node height no growth more than %d s", nodeHeightNoGrowthTime))
						}
					}
					send, recover := needSendNodeStatusAlarm(&nodeStatus)
					if send {
						if err := sendNodeStatusDingAlarm(nodeStatus, recover); err != nil {
							logs.Error("%s node: %s sendNodeStatusDingAlarm err:", h.handle.GetChainName(), nodeStatus.Url, err)
							continue
						}
						if recover {
							if _, err := cacheRedis.Redis.Del(cacheRedis.NodeStatusAlarmPrefix + nodeStatus.Url); err != nil {
								logs.Error("clear %s node: %s alarm err: %s", h.handle.GetChainName(), nodeStatus.Url, err)
							}
						} else {
							if _, err := cacheRedis.Redis.Set(cacheRedis.NodeStatusAlarmPrefix+nodeStatus.Url, "alarm has been sent", time.Second*time.Duration(config.ChainNodeStatusAlarmInterval)); err != nil {
								logs.Error("mark %s node: %s alarm has been sent error: %s", h.handle.GetChainName(), nodeStatus.Url, err)
							}
						}
					}
				}
			}
		}
	}
}

func needSendNodeStatusAlarm(nodeStatus *basedef.NodeStatus) (send, recover bool) {
	exist, err := cacheRedis.Redis.Exists(cacheRedis.NodeStatusAlarmPrefix + nodeStatus.Url)
	if err == nil {
		if exist {
			if len(nodeStatus.Status) == 1 && nodeStatus.Status[0] == basedef.NodeStatusOk {
				send = true
				recover = true
			}
		} else {
			if len(nodeStatus.Status) >= 1 && nodeStatus.Status[0] != basedef.NodeStatusOk {
				send = true
				recover = false
			}
		}
	}
	return
}

func sendNodeStatusDingAlarm(nodeStatus basedef.NodeStatus, isRecover bool) error {
	title := ""
	status := ""
	if isRecover {
		title = fmt.Sprintf("%s Node Recover", nodeStatus.ChainName)
		status = "<font color=green>OK</font>"
	} else {
		title = fmt.Sprintf("%s Node ALarm", nodeStatus.ChainName)
		for i, info := range nodeStatus.Status {
			status = fmt.Sprintf("%s\n%d. <font color=red>%s</font>", status, i+1, info)
		}
	}
	body := fmt.Sprintf("## %s\n- Node: %s\n- Height: %d\n- Status: %s\n- Time: %d\n",
		title,
		nodeStatus.Url,
		nodeStatus.Height,
		status,
		nodeStatus.Time,
	)
	buttons := []map[string]string{
		{
			"title":     "List All",
			"actionURL": fmt.Sprintf("%stoken=%s", conf.GlobalConfig.BotConfig.BaseUrl+conf.GlobalConfig.BotConfig.ListNodeStatusUrl, conf.GlobalConfig.BotConfig.ApiToken),
		},
	}
	logs.Info(body)
	return polycommon.PostDingCard(title, body, buttons, conf.GlobalConfig.BotConfig.NodeStatusDingUrl)
}

func NewHealthMonitorHandle(monitorConfig *conf.HealthMonitorConfig) HealthMonitorHandle {
	switch monitorConfig.ChainId {
	case basedef.POLY_CROSSCHAIN_ID:
		return polymonitor.NewPolyHealthMonitor(monitorConfig)
	case basedef.ETHEREUM_CROSSCHAIN_ID, basedef.O3_CROSSCHAIN_ID, basedef.BSC_CROSSCHAIN_ID, basedef.PLT_CROSSCHAIN_ID,
		basedef.OK_CROSSCHAIN_ID, basedef.HECO_CROSSCHAIN_ID, basedef.MATIC_CROSSCHAIN_ID, basedef.ARBITRUM_CROSSCHAIN_ID,
		basedef.XDAI_CROSSCHAIN_ID, basedef.FANTOM_CROSSCHAIN_ID, basedef.AVAX_CROSSCHAIN_ID, basedef.OPTIMISTIC_CROSSCHAIN_ID,
		basedef.METIS_CROSSCHAIN_ID:
		return ethereummonitor.NewEthereumHealthMonitor(monitorConfig)
	case basedef.NEO_CROSSCHAIN_ID:
		return neomonitor.NewNeoHealthMonitor(monitorConfig)
	case basedef.ONT_CROSSCHAIN_ID:
		return ontologymonitor.NewOntologyHealthMonitor(monitorConfig)
	case basedef.SWITCHEO_CROSSCHAIN_ID:
		return switcheomonitor.NewSwitcheoHealthMonitor(monitorConfig)
	case basedef.NEO3_CROSSCHAIN_ID:
		return neo3monitor.NewNeo3HealthMonitor(monitorConfig)
	case basedef.ZILLIQA_CROSSCHAIN_ID:
		return zilliqamonitor.NewZilliqaHealthMonitor(monitorConfig)
	default:
		return nil
	}
}
