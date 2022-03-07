package neo3monitor

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/keys"
	"github.com/joeqian10/neo3-gogogo/tx"
	"github.com/joeqian10/neo3-gogogo/wallet"
	"math"
	"math/big"
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

func (n *Neo3Monitor) RelayerBalanceMonitor() ([]*basedef.RelayerAccountStatus, error) {
	var precision float64 = 100000000
	var sdk *chainsdk.Neo3Sdk
	var maxHeight uint64
	isMaxHeight := func(height uint64) bool {
		if height >= maxHeight {
			maxHeight = height
			return true
		}
		return false
	}
	for _, s := range n.sdks {
		height, _ := s.GetBlockCount()
		if isMaxHeight(height) {
			sdk = s
		}
	}

	balanceSuccessMap := make(map[string]*big.Int, 0)
	balanceFailedMap := make(map[string]string, 0)
	for _, account := range n.monitorConfig.RelayerAccount.Neo3Account {
		if _, ok := balanceSuccessMap[account.Address]; ok {
			continue
		}
		keypair, err := keys.NewKeyPairFromNEP2(account.Key, account.Pwd, helper.DefaultAddressVersion, keys.N, keys.R, keys.P)
		if err != nil {
			balanceFailedMap[account.Address] = err.Error()
			continue
		}
		wh, err := wallet.NewWalletHelperFromPrivateKey(sdk.GetClient(), keypair.PrivateKey)
		if err != nil {
			balanceFailedMap[account.Address] = err.Error()
		}
		accountAndBalances, err := wh.GetAccountAndBalance(tx.GasToken)
		total := big.NewInt(0)

		if err != nil {
			balanceFailedMap[account.Address] = err.Error()
		} else {
			for _, balance := range accountAndBalances {
				total = total.Add(total, balance.Value)
			}
			if total.Uint64() != 0 {
				balanceSuccessMap[account.Address] = total
				delete(balanceFailedMap, account.Address)
			} else {
				balanceFailedMap[account.Address] = "balance is 0 or all nodes are unavailable"
			}
		}
	}
	relayerStatus := make([]*basedef.RelayerAccountStatus, 0)
	for address, balance := range balanceSuccessMap {
		status := basedef.RelayerAccountStatus{
			ChainId:   n.monitorConfig.ChainId,
			ChainName: n.monitorConfig.ChainName,
			Address:   address,
			Balance:   float64(balance.Uint64()) / precision,
			Threshold: n.monitorConfig.RelayerAccount.Threshold / precision,
			Time:      time.Now().Unix(),
		}
		relayerStatus = append(relayerStatus, &status)
	}
	for address, err := range balanceFailedMap {
		logs.Error("get %s relayer[%s] balance failed. err: %s", n.monitorConfig.ChainName, address, err)
		status := basedef.RelayerAccountStatus{
			ChainId:   n.monitorConfig.ChainId,
			ChainName: n.monitorConfig.ChainName,
			Address:   address,
			Balance:   0,
			Threshold: n.monitorConfig.RelayerAccount.Threshold / precision,
			Status:    err,
			Time:      time.Now().Unix(),
		}
		relayerStatus = append(relayerStatus, &status)
	}
	return relayerStatus, nil
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
			n.nodeStatus[url] = basedef.StatusOk
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
