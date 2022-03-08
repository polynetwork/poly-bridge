package healthmonitor

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"poly-bridge/basedef"
	"poly-bridge/cacheRedis"
	"poly-bridge/common"
	"poly-bridge/conf"
	"poly-bridge/monitor/healthmonitor/ethereummonitor"
	"poly-bridge/monitor/healthmonitor/neo3monitor"
	"poly-bridge/monitor/healthmonitor/neomonitor"
	"poly-bridge/monitor/healthmonitor/ontologymonitor"
	"poly-bridge/monitor/healthmonitor/polymonitor"
	"poly-bridge/monitor/healthmonitor/switcheomonitor"
	"poly-bridge/monitor/healthmonitor/zilliqamonitor"
	"runtime/debug"
	"time"
)

var healthMonitorConfigMap = make(map[uint64]*conf.HealthMonitorConfig, 0)

func StartHealthMonitor(config *conf.Config, relayerConfig *conf.RelayerConfig) {
	for _, cfg := range config.ChainNodes {
		monitorConfig := &conf.HealthMonitorConfig{ChainId: cfg.ChainId, ChainName: cfg.ChainName, ChainNodes: cfg}
		healthMonitorConfigMap[cfg.ChainId] = monitorConfig
	}
	for _, cfg := range config.ChainListenConfig {
		healthMonitorConfigMap[cfg.ChainId].CCMContract = cfg.CCMContract
	}
	for _, cfg := range relayerConfig.RelayAccountConfig {
		healthMonitorConfigMap[cfg.ChainId].RelayerAccount = cfg
	}
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

type MonitorHandle interface {
	NodeMonitor() ([]basedef.NodeStatus, error)
	RelayerBalanceMonitor() ([]*basedef.RelayerAccountStatus, error)
	GetChainName() string
}

type HealthMonitor struct {
	handle MonitorHandle
}

func (h *HealthMonitor) Start(config *conf.Config) {
	go h.NodeMonitor(config)
	go h.RelayerAccountMonitor(config)
}

func (h *HealthMonitor) NodeMonitor(config *conf.Config) {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("NodeMonitor restart, recover info: %s", string(debug.Stack()))
		}
	}()

	logs.Info("start %s NodeMonitor", h.handle.GetChainName())
	nodeMonitorTicker := time.NewTicker(time.Second * time.Duration(config.BotConfig.ChainNodeStatusCheckInterval))
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
					sendAlarm, recoverAlarm := needSendNodeStatusAlarm(&nodeStatus)
					if sendAlarm {
						if err := sendNodeStatusDingAlarm(nodeStatus, recoverAlarm); err != nil {
							logs.Error("%s node: %s sendNodeStatusDingAlarm err:", h.handle.GetChainName(), nodeStatus.Url, err)
							continue
						}
						if recoverAlarm {
							if _, err := cacheRedis.Redis.Del(cacheRedis.NodeStatusAlarmPrefix + nodeStatus.Url); err != nil {
								logs.Error("clear %s node: %s alarm err: %s", h.handle.GetChainName(), nodeStatus.Url, err)
							}
						} else {
							if _, err := cacheRedis.Redis.Set(cacheRedis.NodeStatusAlarmPrefix+nodeStatus.Url, "alarm has been sent", time.Second*time.Duration(config.BotConfig.ChainNodeStatusAlarmInterval)); err != nil {
								logs.Error("mark %s node: %s alarm has been sent error: %s", h.handle.GetChainName(), nodeStatus.Url, err)
							}
						}
					}
				}
			}
		}
	}
}

func (h *HealthMonitor) RelayerAccountMonitor(config *conf.Config) {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("%s RelayerAccountMonitor recover info: %s", h.handle.GetChainName(), string(debug.Stack()))
		}
	}()
	logs.Info("start %s RelayerBalanceMonitor", h.handle.GetChainName())

	monitorTicker := time.NewTicker(time.Minute * 5)
	for {
		select {
		case <-monitorTicker.C:
			if relayerAccountStatuses, err := h.handle.RelayerBalanceMonitor(); relayerAccountStatuses != nil && err == nil {
				for _, accountStatus := range relayerAccountStatuses {
					if len(accountStatus.Status) == 0 {
						if accountStatus.Balance < accountStatus.Threshold {
							accountStatus.Status = "insufficient"
							logs.Error("%s relayer %s", h.handle.GetChainName(), accountStatus.Status)
							continue
						} else {
							accountStatus.Status = basedef.StatusOk
						}
					}
				}
				data, _ := json.Marshal(relayerAccountStatuses)
				_, err := cacheRedis.Redis.Set(cacheRedis.RelayerAccountStatusPrefix+h.handle.GetChainName(), data, time.Hour*24)
				if err != nil {
					logs.Error("set %s node status error: %s", h.handle.GetChainName(), err)
				}
				for _, accountStatus := range relayerAccountStatuses {
					sendAlarm, recoverAlarm := needSendRelayerAccountStatusAlarm(accountStatus)
					if sendAlarm {
						if err := sendRelayerAccountStatusDingAlarm(accountStatus, recoverAlarm); err != nil {
							logs.Error("%s relayer address: %s sendRelayerAccountStatusDingAlarm err:", h.handle.GetChainName(), accountStatus.Address, err)
							continue
						}
						alarmKey := fmt.Sprintf("%s%s-%s", cacheRedis.RelayerAccountStatusAlarmPrefix, accountStatus.ChainName, accountStatus.Address)
						if recoverAlarm {
							if _, err := cacheRedis.Redis.Del(alarmKey); err != nil {
								logs.Error("clear %s relayer address: %s alarm err: %s", h.handle.GetChainName(), accountStatus.Address, err)
							}
							logs.Info("clear %s relayer address: %s alarm", h.handle.GetChainName(), accountStatus.Address)
						} else {
							if _, err := cacheRedis.Redis.Set(alarmKey, "alarm has been sent", time.Hour*12); err != nil {
								logs.Error("mark %s relayer address: %s alarm has been sent error: %s", h.handle.GetChainName(), accountStatus.Address, err)
							}
							logs.Info("mark %s relayer address: %s alarm has been sent", h.handle.GetChainName(), accountStatus.Address)
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
			if len(nodeStatus.Status) == 1 && nodeStatus.Status[0] == basedef.StatusOk {
				send = true
				recover = true
			}
		} else {
			if len(nodeStatus.Status) >= 1 && nodeStatus.Status[0] != basedef.StatusOk {
				send = true
				recover = false
			}
		}
	}

	if send {
		ignore, _ := cacheRedis.Redis.Exists(cacheRedis.IgnoreNodeStatusAlarmPrefix + nodeStatus.Url)
		if ignore {
			send = false
			logs.Info("ignore %s node: %s alarm", nodeStatus.ChainName, nodeStatus.Url)
		}
	}
	return
}

func needSendRelayerAccountStatusAlarm(relayerStatus *basedef.RelayerAccountStatus) (send, recover bool) {
	alarmKey := fmt.Sprintf("%s%s-%s", cacheRedis.RelayerAccountStatusAlarmPrefix, relayerStatus.ChainName, relayerStatus.Address)
	exist, err := cacheRedis.Redis.Exists(alarmKey)
	if err == nil {
		if exist {
			if relayerStatus.Status == basedef.StatusOk {
				send = true
				recover = true
			}
		} else {
			if relayerStatus.Status != basedef.StatusOk {
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
		status = "OK"
	} else {
		title = fmt.Sprintf("%s Node Alarm", nodeStatus.ChainName)
		for _, info := range nodeStatus.Status {
			status = fmt.Sprintf("%s\n%s", status, info)
		}
	}
	body := fmt.Sprintf("## %s\n- Node: %s\n- Height: %d\n- Status: %s\n- Time: %s\n",
		title,
		nodeStatus.Url,
		nodeStatus.Height,
		status,
		time.Unix(nodeStatus.Time, 0).Format("2006-01-02 15:04:05"),
	)

	buttons := make([]map[string]string, 0)
	if !isRecover {
		buttons = append(buttons, []map[string]string{
			{
				"title":     "Ignore For 1 Day",
				"actionURL": fmt.Sprintf("%stoken=%s&node=%s&day=%d", conf.GlobalConfig.BotConfig.BaseUrl+conf.GlobalConfig.BotConfig.IgnoreNodeStatusAlarmUrl, conf.GlobalConfig.BotConfig.ApiToken, nodeStatus.Url, 1),
			},
			{
				"title":     "Ignore For 10 Day",
				"actionURL": fmt.Sprintf("%stoken=%s&node=%s&day=%d", conf.GlobalConfig.BotConfig.BaseUrl+conf.GlobalConfig.BotConfig.IgnoreNodeStatusAlarmUrl, conf.GlobalConfig.BotConfig.ApiToken, nodeStatus.Url, 10),
			},
			{
				"title":     "Cancel Ignore",
				"actionURL": fmt.Sprintf("%stoken=%s&node=%s&day=%d", conf.GlobalConfig.BotConfig.BaseUrl+conf.GlobalConfig.BotConfig.IgnoreNodeStatusAlarmUrl, conf.GlobalConfig.BotConfig.ApiToken, nodeStatus.Url, 0),
			},
		}...)
	}
	buttons = append(buttons, map[string]string{
		"title":     "List All",
		"actionURL": fmt.Sprintf("%stoken=%s", conf.GlobalConfig.BotConfig.BaseUrl+conf.GlobalConfig.BotConfig.ListNodeStatusUrl, conf.GlobalConfig.BotConfig.ApiToken),
	})

	logs.Info(body)
	logs.Info(buttons)
	return common.PostDingCard(title, body, buttons, conf.GlobalConfig.BotConfig.NodeStatusDingUrl)
}

func sendRelayerAccountStatusDingAlarm(relayerStatus *basedef.RelayerAccountStatus, isRecover bool) error {
	title := ""
	if isRecover {
		title = fmt.Sprintf("%s relayer refilled", relayerStatus.ChainName)
	} else {
		title = fmt.Sprintf("%s relayer insufficient", relayerStatus.ChainName)
	}

	body := fmt.Sprintf("## %s\n- Address: %s\n- Balance: %f\n-  Threshold:%f\n- Time: %s\n",
		title,
		relayerStatus.Address,
		relayerStatus.Balance,
		relayerStatus.Threshold,
		time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
	)

	buttons := make([]map[string]string, 0)
	buttons = append(buttons, map[string]string{
		"title":     "List All",
		"actionURL": fmt.Sprintf("%stoken=%s", conf.GlobalConfig.BotConfig.BaseUrl+conf.GlobalConfig.BotConfig.ListRelayerAccountStatusUrl, conf.GlobalConfig.BotConfig.ApiToken),
	})

	logs.Info(body)
	logs.Info(buttons)
	return common.PostDingCard(title, body, buttons, conf.GlobalConfig.BotConfig.RelayerAccountStatusDingUrl)
}

func NewHealthMonitorHandle(monitorConfig *conf.HealthMonitorConfig) MonitorHandle {
	switch monitorConfig.ChainId {
	case basedef.POLY_CROSSCHAIN_ID:
		return polymonitor.NewPolyHealthMonitor(monitorConfig)
	case basedef.ETHEREUM_CROSSCHAIN_ID, basedef.O3_CROSSCHAIN_ID, basedef.BSC_CROSSCHAIN_ID, basedef.PLT_CROSSCHAIN_ID,
		basedef.OK_CROSSCHAIN_ID, basedef.HECO_CROSSCHAIN_ID, basedef.MATIC_CROSSCHAIN_ID, basedef.ARBITRUM_CROSSCHAIN_ID,
		basedef.XDAI_CROSSCHAIN_ID, basedef.FANTOM_CROSSCHAIN_ID, basedef.AVAX_CROSSCHAIN_ID, basedef.OPTIMISTIC_CROSSCHAIN_ID,
		basedef.METIS_CROSSCHAIN_ID, basedef.PIXIE_CROSSCHAIN_ID, basedef.RINKEBY_CROSSCHAIN_ID, basedef.BOBA_CROSSCHAIN_ID,
		basedef.OASIS_CROSSCHAIN_ID, basedef.OASIS1_CROSSCHAIN_ID, basedef.HARMONY_CROSSCHAIN_ID:
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
