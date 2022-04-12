package ethereummonitor

import (
	"context"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math"
	"math/big"
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
}

func NewEthereumHealthMonitor(monitorConfig *conf.HealthMonitorConfig) *EthereumHealthMonitor {
	ethMonitor := &EthereumHealthMonitor{}
	ethMonitor.monitorConfig = monitorConfig
	sdks := make(map[string]*chainsdk.EthereumSdk, 0)
	for _, node := range monitorConfig.ChainNodes.Nodes {
		sdk, err := chainsdk.NewEthereumSdk(node.Url)
		if err != nil || sdk == nil || sdk.GetClient() == nil {
			if _, e := cacheRedis.Redis.Set(cacheRedis.NodeStatusPrefix+node.Url, fmt.Sprintf("initial sdk error:%s", err), time.Hour*24); e != nil {
				logs.Error("set %s node[%s] status error: %s", monitorConfig.ChainName, node.Url, e)
			}
			logs.Error("%s node: %s, NewEthereumSdk error: %s", monitorConfig.ChainName, node.Url, err)
			continue
		}
		sdks[node.Url] = sdk
	}
	ethMonitor.sdks = sdks
	ethMonitor.nodeHeight = make(map[string]uint64, len(sdks))
	return ethMonitor
}

func (e *EthereumHealthMonitor) GetChainName() string {
	return e.monitorConfig.ChainName
}

func (e *EthereumHealthMonitor) GetChainId() uint64 {
	return e.monitorConfig.ChainId
}

func (e *EthereumHealthMonitor) RelayerBalanceMonitor() ([]*basedef.RelayerAccountStatus, error) {
	switch e.monitorConfig.ChainId {
	case basedef.PLT_CROSSCHAIN_ID, basedef.BCSPALETTE_CROSSCHAIN_ID, basedef.O3_CROSSCHAIN_ID:
		return nil, nil
	}
	balanceSuccessMap := make(map[string]*big.Int, 0)
	balanceFailedMap := make(map[string]string, 0)
	var precision float64 = 1000000000000000000
	for _, sdk := range e.sdks {
		for _, address := range e.monitorConfig.RelayerAccount.Address {
			if _, ok := balanceSuccessMap[address]; ok {
				continue
			}
			balance, err := sdk.GetNativeBalance(common.HexToAddress(address))
			if err == nil {
				balanceSuccessMap[address] = balance
				delete(balanceFailedMap, address)
			} else {
				balanceFailedMap[address] = err.Error()
			}
		}
	}
	relayerStatus := make([]*basedef.RelayerAccountStatus, 0)
	for address, balance := range balanceSuccessMap {
		status := basedef.RelayerAccountStatus{
			ChainId:   e.monitorConfig.ChainId,
			ChainName: e.monitorConfig.ChainName,
			Address:   address,
			Balance:   float64(balance.Uint64()) / precision,
			Threshold: e.monitorConfig.RelayerAccount.Threshold / precision,
			Time:      time.Now().Unix(),
		}
		relayerStatus = append(relayerStatus, &status)
	}
	for address, err := range balanceFailedMap {
		status := basedef.RelayerAccountStatus{
			ChainId:   e.monitorConfig.ChainId,
			ChainName: e.monitorConfig.ChainName,
			Address:   address,
			Balance:   0,
			Threshold: e.monitorConfig.RelayerAccount.Threshold / precision,
			Status:    err,
			Time:      time.Now().Unix(),
		}
		relayerStatus = append(relayerStatus, &status)
	}
	return relayerStatus, nil
}

func (e *EthereumHealthMonitor) NodeMonitor() ([]basedef.NodeStatus, error) {
	nodeStatuses := make([]basedef.NodeStatus, 0)
	for url, sdk := range e.sdks {
		status := basedef.NodeStatus{
			ChainId:   e.monitorConfig.ChainId,
			ChainName: e.monitorConfig.ChainName,
			Url:       url,
			Status:    make([]string, 0),
			Time:      time.Now().Unix(),
		}
		height, err := e.GetCurrentHeight(sdk, e.GetChainName())
		if err == nil {
			status.Height = height
			e.nodeHeight[url] = height
			err = e.CheckAbiCall(sdk)
		}
		if err != nil {
			status.Status = append(status.Status, err.Error())
		}
		nodeStatuses = append(nodeStatuses, status)
	}
	//data, _ := json.Marshal(nodeStatuses)
	//_, err := cacheRedis.Redis.Set(cacheRedis.NodeStatusPrefix+e.GetChainName(), data, time.Hour*24)
	//if err != nil {
	//	logs.Error("set %s node status error: %s", e.GetChainName(), err)
	//}
	return nodeStatuses, nil
}

func (e *EthereumHealthMonitor) GetCurrentHeight(sdk *chainsdk.EthereumSdk, chainName string) (uint64, error) {
	height, err := sdk.GetCurrentBlockHeight()
	if err != nil || height == 0 || height == math.MaxUint64 {
		logs.Info("%s height=%d", chainName, height)
		err2 := fmt.Errorf("get current block height err: %s", err)
		logs.Error(fmt.Sprintf("%s node: %s, %s ", chainName, sdk.GetUrl(), err2))
		return 0, err2
	}
	logs.Info("%s node: %s, latest height: %d", chainName, sdk.GetUrl(), height)
	return height, nil
}

func (e *EthereumHealthMonitor) CheckAbiCall(sdk *chainsdk.EthereumSdk) error {
	eccmContractAddress := common.HexToAddress(e.monitorConfig.CCMContract)
	client := sdk.GetClient()
	ethCrossChainManager, err := eccm_abi.NewEthCrossChainManager(eccmContractAddress, client)
	if err != nil {
		err2 := fmt.Errorf("call NewEthCrossChainManager error: %s", err)
		logs.Error(fmt.Sprintf("%s node: %s, %s ", e.GetChainName(), sdk.GetUrl(), err2))
		return err2
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
		err2 := fmt.Errorf("call FilterCrossChainEvent get lock events err: %s", err)
		logs.Error(fmt.Sprintf("%s node: %s, %s ", e.GetChainName(), sdk.GetUrl(), err2))
		return err2
	}
	// get unlock events from given block
	_, err = ethCrossChainManager.FilterVerifyHeaderAndExecuteTxEvent(opt)
	if err != nil {
		err2 := fmt.Errorf("call FilterVerifyHeaderAndExecuteTxEvent get unlock events err: %s", err)
		logs.Error(fmt.Sprintf("%s node: %s, %s ", e.GetChainName(), sdk.GetUrl(), err2))
		return err2
	}
	return nil
}
