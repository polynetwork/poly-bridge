package chainhealthmonitor

import (
	"context"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math"
	"poly-bridge/basedef"
	"poly-bridge/chainhealthmonitor/ethereummonitor"
	"poly-bridge/chainhealthmonitor/neo3monitor"
	"poly-bridge/chainhealthmonitor/neomonitor"
	"poly-bridge/chainhealthmonitor/ontologymonitor"
	"poly-bridge/chainhealthmonitor/polymonitor"
	"poly-bridge/chainhealthmonitor/switcheomonitor"
	"poly-bridge/chainhealthmonitor/zilliqamonitor"
	"poly-bridge/chainsdk"
	polycommon "poly-bridge/common"
	"poly-bridge/conf"
	"poly-bridge/go_abi/eccm_abi"
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
		monitor.Start()
	}

}

type HealthMonitorHandle interface {
	NodeMonitor() ([]basedef.NodeStatus, error)
	GetChainName() string
}

type HealthMonitor struct {
	handle HealthMonitorHandle
}

func (h *HealthMonitor) Start() {
	go h.NodeMonitor()
}

func (h *HealthMonitor) NodeMonitor() {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("NodeMonitor restart, recover info: %s", string(debug.Stack()))
		}
	}()

	logs.Info("start %s NodeMonitor", h.handle.GetChainName())
	nodeMonitorTicker := time.NewTicker(time.Second * time.Duration(20))
	for {
		select {
		case <-nodeMonitorTicker.C:
			if nodeStatuses, err := h.handle.NodeMonitor(); err == nil {
				logs.Info("%s nodeStatuses:%+v", h.handle.GetChainName(), nodeStatuses)
				for _, nodeStatus := range nodeStatuses {
					if nodeStatus.Status != basedef.NodeStatusOk {
						if err := sendNodeStatusDingAlarm(nodeStatus); err != nil {
							logs.Error("sendNodeStatusDingAlarm err:", err)
						}
					}
				}
			}
		}
	}
}

func sendNodeStatusDingAlarm(nodeStatus basedef.NodeStatus) error {
	title := fmt.Sprintf("%s Node ALarm:", nodeStatus.ChainName)
	body := fmt.Sprintf("## %s\n- Node: %s\n- Height: %d\n- Status: %s\n",
		title,
		nodeStatus.Url,
		nodeStatus.Height,
		nodeStatus.Status,
	)
	buttons := make([]map[string]string, 0)
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

func EthNodeMonitor(config *conf.Config) {
	logs.Info("EthNodeMonitor")
	var ccmContractAddr string
	for _, listenConfig := range config.ChainListenConfig {
		if listenConfig.ChainId == basedef.HECO_CROSSCHAIN_ID {
			ccmContractAddr = listenConfig.CCMContract
			break
		}
	}

	for _, chainNodeConfig := range config.ChainNodes {
		if chainNodeConfig.ChainId == basedef.HECO_CROSSCHAIN_ID {
			for _, node := range chainNodeConfig.Nodes {
				sdk, err := chainsdk.NewEthereumSdk(node.Url)
				if err != nil || sdk == nil || sdk.GetClient() == nil {
					logs.Info("node: %s,NewEthereumSdk error: %s", node.Url, err)
					continue
				}
				height, err := sdk.GetCurrentBlockHeight()
				if err != nil || height == 0 || height == math.MaxUint64 {
					logs.Error("node: %s, get current block height err: %s, ", sdk.GetUrl(), err)
					continue
				}
				height -= 1
				//height = 13881338

				logs.Info("node: %s, height: %d", node.Url, height)

				eccmContractAddress := common.HexToAddress(ccmContractAddr)
				client := sdk.GetClient()
				eccmContract, err := eccm_abi.NewEthCrossChainManager(eccmContractAddress, client)
				if err != nil {
					logs.Error("node: %s, NewEthCrossChainManager error: %s", sdk.GetUrl(), err)
					continue
				}
				opt := &bind.FilterOpts{
					Start:   height,
					End:     &height,
					Context: context.Background(),
				}
				// get ethereum lock events from given block
				_, err = eccmContract.FilterCrossChainEvent(opt, nil)
				if err != nil {
					logs.Error("node: %s, FilterCrossChainEvent error: %s", sdk.GetUrl(), err)
					continue
				}
				// ethereum unlock events from given block
				_, err = eccmContract.FilterVerifyHeaderAndExecuteTxEvent(opt)
				if err != nil {
					logs.Error("node: %s, FilterVerifyHeaderAndExecuteTxEvent error: %s", sdk.GetUrl(), err)
					continue
				}
			}
		}
	}
}
