package ontologymonitor

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/ontio/ontology-go-sdk"
	"math"
	"poly-bridge/basedef"
	"poly-bridge/cacheRedis"
	"poly-bridge/conf"
	"time"
)

type OntologyMonitor struct {
	monitorConfig *conf.HealthMonitorConfig
	sdks          map[string]*ontology_go_sdk.OntologySdk
	nodeHeight    map[string]uint64
	nodeStatus    map[string]string
}

func NewOntologyHealthMonitor(monitorConfig *conf.HealthMonitorConfig) *OntologyMonitor {
	ontologyMonitor := &OntologyMonitor{}
	ontologyMonitor.monitorConfig = monitorConfig
	sdks := make(map[string]*ontology_go_sdk.OntologySdk, 0)
	for _, node := range monitorConfig.ChainNodes.Nodes {
		sdk := ontology_go_sdk.NewOntologySdk()
		sdk.NewRpcClient().SetAddress(node.Url)
		sdks[node.Url] = sdk
	}
	ontologyMonitor.sdks = sdks
	return ontologyMonitor
}

func (o *OntologyMonitor) GetChainName() string {
	return o.monitorConfig.ChainName
}

func (o *OntologyMonitor) NodeMonitor() ([]basedef.NodeStatus, error) {
	nodeStatuses := make([]basedef.NodeStatus, 0)
	for url, sdk := range o.sdks {
		status := basedef.NodeStatus{
			ChainId:   o.monitorConfig.ChainId,
			ChainName: o.monitorConfig.ChainName,
			Url:       url,
			Status:    make([]string, 0),
			Time:      time.Now().Unix(),
		}
		height, err := o.GetCurrentHeight(sdk, url)
		if err == nil {
			status.Height = height
			o.nodeHeight[url] = height
			err = o.CheckAbiCall(sdk, url)
		}
		if err != nil {
			o.nodeStatus[url] = err.Error()
		} else {
			o.nodeStatus[url] = basedef.NodeStatusOk
		}
		status.Status = append(status.Status, o.nodeStatus[url])
		nodeStatuses = append(nodeStatuses, status)
	}
	data, _ := json.Marshal(nodeStatuses)
	_, err := cacheRedis.Redis.Set(cacheRedis.NodeStatusPrefix+o.monitorConfig.ChainName, data, time.Hour*24)
	if err != nil {
		logs.Error("set %s node status error: %s", o.GetChainName(), err)
	}
	return nodeStatuses, err
}

func (o *OntologyMonitor) GetCurrentHeight(sdk *ontology_go_sdk.OntologySdk, url string) (uint64, error) {
	height, err := sdk.GetCurrentBlockHeight()
	if err != nil || height == 0 || height == math.MaxUint32 {
		err := fmt.Errorf("%s node: %s, get current block height err: %s", o.GetChainName(), url, err)
		logs.Error(fmt.Sprintf("%s node: %s, %s ", o.GetChainName(), url, err))
		return 0, err
	}
	logs.Info("%s node: %s, latest height: %d", o.GetChainName(), url, height)
	return uint64(height), nil
}

func (o *OntologyMonitor) CheckAbiCall(sdk *ontology_go_sdk.OntologySdk, url string) error {
	height := uint32(o.nodeHeight[url]) - 1
	_, err := sdk.GetBlockByHeight(height)
	if err != nil {
		err := fmt.Errorf("GetBlockByHeight err: %s", err)
		logs.Error(fmt.Sprintf("%s node: %s, %s ", o.GetChainName(), url, err))
		return err
	}
	_, err = sdk.GetSmartContractEventByBlock(height)
	if err != nil {
		err := fmt.Errorf("call GetSmartContractEventByBlock err: %s", err)
		logs.Error(fmt.Sprintf("%s node: %s, %s ", o.GetChainName(), url, err))
		return err
	}
	return nil
}
