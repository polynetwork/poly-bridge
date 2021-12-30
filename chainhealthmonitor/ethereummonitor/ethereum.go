package ethereummonitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math"
	"poly-bridge/basedef"
	"poly-bridge/cacheRedis"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"poly-bridge/go_abi/eccm_abi"
	"time"
)

type EthereumHealthMonitor struct {
	monitorConfig *conf.HealthMonitorConfig
	sdks          map[string]*chainsdk.EthereumSdk
	nodeHeight    map[string]uint64
	nodeStatus    map[string]string
}

func NewEthereumHealthMonitor(monitorConfig *conf.HealthMonitorConfig) *EthereumHealthMonitor {
	ethMonitor := &EthereumHealthMonitor{}
	ethMonitor.monitorConfig = monitorConfig
	sdks := make(map[string]*chainsdk.EthereumSdk, 0)
	for _, node := range monitorConfig.ChainNodes.Nodes {
		sdk, err := chainsdk.NewEthereumSdk(node.Url)
		if err != nil || sdk == nil || sdk.GetClient() == nil {
			if _, err := cacheRedis.Redis.Set(cacheRedis.NodeStatusPrefix+node.Url, fmt.Sprintf("initial sdk error:%s", err), time.Hour*24); err != nil {
				logs.Error("set eth node[%s] status error: %s", node.Url, err)
			}
			logs.Error("eth node: %s, NewEthereumSdk error: %s", node.Url, err)
			continue
		}
		sdks[node.Url] = sdk
	}
	ethMonitor.sdks = sdks
	ethMonitor.nodeHeight = make(map[string]uint64, len(sdks))
	ethMonitor.nodeStatus = make(map[string]string, len(sdks))
	return ethMonitor
}

func (e *EthereumHealthMonitor) GetChainName() string {
	return e.monitorConfig.ChainName
}

func (e *EthereumHealthMonitor) NodeMonitor() error {
	nodeStatuses := make([]basedef.NodeStatus, 0)
	for url, sdk := range e.sdks {
		status := basedef.NodeStatus{
			ChainId:   e.monitorConfig.ChainId,
			ChainName: e.monitorConfig.ChainName,
			Url:       url,
			Time:      time.Now().Format("2006-01-02 15:04:05"),
		}
		height, err := e.GetCurrentHeight(sdk, e.monitorConfig.ChainName)
		if err == nil {
			status.Height = height
			e.nodeHeight[url] = height
			err = e.CheckAbiCall(sdk, e.monitorConfig.ChainName)
		}
		if err != nil {
			e.nodeStatus[url] = err.Error()
		} else {
			e.nodeStatus[url] = "OK"
		}
		status.Status = e.nodeStatus[url]
		nodeStatuses = append(nodeStatuses, status)
	}
	data, _ := json.Marshal(nodeStatuses)
	_, err := cacheRedis.Redis.Set(cacheRedis.NodeStatusPrefix+e.monitorConfig.ChainName, data, time.Hour*24)
	if err != nil {
		logs.Error("set neo3 node status error: %s", err)
	}
	return err
}

func (e *EthereumHealthMonitor) GetCurrentHeight(sdk *chainsdk.EthereumSdk, chainName string) (uint64, error) {
	height, err := sdk.GetCurrentBlockHeight()
	if err != nil || height == 0 || height == math.MaxUint64 {
		err := fmt.Errorf("get current block height err: %s, ", err)
		logs.Error(fmt.Sprintf("%s node: %s, %s ", chainName, sdk.GetUrl(), err))
		return 0, err
	}
	logs.Info("%s node: %s, latest height: %d", chainName, sdk.GetUrl(), height)
	return height, nil
}

func (e *EthereumHealthMonitor) CheckAbiCall(sdk *chainsdk.EthereumSdk, chainName string) error {
	eccmContractAddress := common.HexToAddress(e.monitorConfig.CCMContract)
	client := sdk.GetClient()
	ethCrossChainManager, err := eccm_abi.NewEthCrossChainManager(eccmContractAddress, client)
	if err != nil {
		err := fmt.Errorf("call NewEthCrossChainManager error: %s", err)
		logs.Error(fmt.Sprintf("%s node: %s, %s ", chainName, sdk.GetUrl(), err))
		e.nodeStatus[sdk.GetUrl()] = err.Error()
		return err
	}
	height := e.nodeHeight[sdk.GetUrl()] - 1
	opt := &bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: context.Background(),
	}
	// get lock events from given block
	_, err = ethCrossChainManager.FilterCrossChainEvent(opt, nil)
	if err != nil {
		err := fmt.Errorf("call FilterCrossChainEvent get lock events err: %s", err)
		logs.Error(fmt.Sprintf("%s node: %s, %s ", chainName, sdk.GetUrl(), err))
		e.nodeStatus[sdk.GetUrl()] = err.Error()
		return err
	}
	// get unlock events from given block
	_, err = ethCrossChainManager.FilterVerifyHeaderAndExecuteTxEvent(opt)
	if err != nil {
		err := fmt.Errorf("call FilterVerifyHeaderAndExecuteTxEvent get unlock events err: %s", err)
		logs.Error(fmt.Sprintf("%s node: %s, %s ", chainName, sdk.GetUrl(), err))
		e.nodeStatus[sdk.GetUrl()] = err.Error()
		return err
	}
	return nil
}
