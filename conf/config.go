/*
 * Copyright (C) 2020 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package conf

import (
	"encoding/json"
	"poly-bridge/basedef"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/urfave/cli"
)

var GlobalConfig *Config
var PolyProxy map[string]bool

var (
	ConfigPathFlag = cli.StringFlag{
		Name:  "config",
		Usage: "Server config file `<path>`",
		Value: "./conf/config_testnet.json",
	}
)

//getFlagName deal with short flag, and return the flag name whether flag name have short name
func GetFlagName(flag cli.Flag) string {
	return strings.TrimSpace(strings.Split(flag.GetName(), ",")[0])
}

type DBConfig struct {
	URL      string
	User     string
	Password string
	Scheme   string
	Debug    bool
}

type RedisConfig struct {
	Proto        string        `json:"proto"`
	Addr         string        `json:"addr"`
	Password     string        `json:"password"`
	PoolSize     int           `json:"pool_size"`
	MinIdleConns int           `json:"min_idle_conns"`
	DialTimeout  time.Duration `json:"dial_timeout"`
	ReadTimeout  time.Duration `json:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout"`
	IdleTimeout  time.Duration `json:"idle_timeout"`
	Expiration   time.Duration `json:"expiration"`
}

type ExpConfig struct {
	URL      string
	User     string
	Password string
	Scheme   string
	Debug    bool
}

type Restful struct {
	Url string
	Key string
}

type OtherItemProxy struct {
	ItemName  string
	ItemProxy string
}

type ChainNodes struct {
	ChainName   string
	ChainId     uint64
	Nodes       []*Restful
	ExtendNodes []*Restful
}

type ChainListenConfig struct {
	ChainName          string
	ChainId            uint64
	ListenSlot         uint64
	Defer              uint64
	BatchSize          uint64
	Nodes              []*Restful
	ExtendNodes        []*Restful
	WrapperContract    []string
	CCMContract        string
	ProxyContract      []string
	OtherProxyContract []*OtherItemProxy
	NFTWrapperContract []string
	NFTProxyContract   []string
	NFTQueryContract   string
	SwapContract       string
}

type HealthMonitorConfig struct {
	ChainId        uint64
	ChainName      string
	ChainNodes     *ChainNodes
	CCMContract    string
	RelayerAccount *RelayAccountConfig
}

type RelayAccountConfig struct {
	ChainName   string
	ChainId     uint64
	Address     []string
	Neo3Account []Neo3Account
	Threshold   float64
}

type Neo3Account struct {
	Address string
	Key     string
	Pwd     string
}

type RelayerConfig struct {
	RelayAccountConfig []*RelayAccountConfig
}

func (cfg *ChainListenConfig) GetNodesUrl() []string {
	urls := make([]string, 0)
	for _, node := range cfg.Nodes {
		urls = append(urls, node.Url)
	}
	return urls
}

func (cfg *ChainListenConfig) GetNodesKey() []string {
	keys := make([]string, 0)
	for _, node := range cfg.Nodes {
		keys = append(keys, node.Key)
	}
	return keys
}

func (cfg *ChainListenConfig) GetExtendNodesUrl() []string {
	urls := make([]string, 0)
	for _, node := range cfg.ExtendNodes {
		urls = append(urls, node.Url)
	}
	return urls
}

func (cfg *ChainListenConfig) GetExtendNodesKey() []string {
	keys := make([]string, 0)
	for _, node := range cfg.ExtendNodes {
		keys = append(keys, node.Key)
	}
	return keys
}

type CoinPriceListenConfig struct {
	MarketName string
	Nodes      []*Restful
}

func (cfg *CoinPriceListenConfig) GetNodesUrl() []string {
	urls := make([]string, 0)
	for _, node := range cfg.Nodes {
		urls = append(urls, node.Url)
	}
	return urls
}

func (cfg *CoinPriceListenConfig) GetNodesKey() []string {
	keys := make([]string, 0)
	for _, node := range cfg.Nodes {
		keys = append(keys, node.Key)
	}
	return keys
}

type FeeListenConfig struct {
	ChainId       uint64
	ChainName     string
	Nodes         []*Restful
	ProxyFee      int64
	MinFee        int64
	GasLimit      int64
	EthL1GasLimit int64
}

func (cfg *FeeListenConfig) GetNodesUrl() []string {
	urls := make([]string, 0)
	for _, node := range cfg.Nodes {
		urls = append(urls, node.Url)
	}
	return urls
}

func (cfg *FeeListenConfig) GetNodesKey() []string {
	keys := make([]string, 0)
	for _, node := range cfg.Nodes {
		keys = append(keys, node.Key)
	}
	return keys
}

type StatsConfig struct {
	TokenBasicStatsInterval    int64 // Chain token basic stats aggregation interval in seconds
	TokenAmountCheckInterval   int64 // Chain token stats aggregation interval in seconds
	TokenStatisticInterval     int64 // TokenStatistic aggregation interval in seconds
	LockTokenStatisticInterval int64 // All lockproxy aggregation interval in seconds
	ChainStatisticInterval     int64 // ChainStatisticInterval except asset aggregation interval in seconds
	ChainAddressCheckInterval  int64 // ChainStatistic's asset Interval aggregation interval in seconds
	AssetStatisticInterval     int64 // AssetStatistic aggregation interval in seconds
	AssetAdressInterval        int64 // AssetAdress aggregation interval in seconds
	CensusTimeLinesInterval    int64 // CensusTimeLines interval in seconds
	CensusAssetLinesInterval   int64 // CensusAssetLinesInterval interval in seconds
}

type EventEffectConfig struct {
	HowOld            int64
	HowOld2           int64
	ChainListening    int64
	EffectSlot        int64
	TimeStatisticSlot int64
}

type BotConfig struct {
	DingUrl                                   string
	LargeTxDingUrl                            string
	NodeStatusDingUrl                         string
	RelayerAccountStatusDingUrl               string
	CheckFrom                                 int64
	Interval                                  int64
	BaseUrl                                   string
	DetailUrl                                 string
	FinishUrl                                 string
	MarkAsPaidUrl                             string
	TxUrl                                     string
	ListLargeTxUrl                            string
	ListNodeStatusUrl                         string
	ListRelayerAccountStatusUrl               string
	IgnoreNodeStatusAlarmUrl                  string
	ApiToken                                  string
	ChainNodeStatusCheckInterval              uint64
	ChainNodeStatusAlarmInterval              uint64
	TxStuckCountMarkChainUnhealthy            int
	AllNodesUnavailableTimeMarkChainUnhealthy int64
	AllNodesNoGrowthTimeMarkChainUnhealthy    int64
}

type HttpConfig struct {
	Address string
	Port    int
}

type IPPortConfig struct {
	WBTCIP string
	USDTIP string
	DingIP string
}

type Config struct {
	Server                string
	Env                   string
	RunMode               string
	Backup                bool
	LargeTxAmount         int64
	LogFile               string
	HttpConfig            *HttpConfig
	MetricConfig          *HttpConfig
	ChainNodes            []*ChainNodes
	ChainListenConfig     []*ChainListenConfig
	CoinPriceUpdateSlot   int64
	CoinPriceListenConfig []*CoinPriceListenConfig
	FeeUpdateSlot         int64
	FeeListenConfig       []*FeeListenConfig
	EventEffectConfig     *EventEffectConfig
	StatsConfig           *StatsConfig
	DBConfig              *DBConfig
	BotConfig             *BotConfig
	RedisConfig           *RedisConfig
	IPPortConfig          *IPPortConfig
	NftConfig             *NftConfig
	RelayUrl              string
}

func (cfg *Config) GetChainListenConfig(chainId uint64) *ChainListenConfig {
	for _, chainListenConfig := range cfg.ChainListenConfig {
		if chainListenConfig.ChainId == chainId {
			return chainListenConfig
		}
	}
	return nil
}

func (cfg *Config) GetCoinPriceListenConfig(market string) *CoinPriceListenConfig {
	for _, coinPriceListenConfig := range cfg.CoinPriceListenConfig {
		if coinPriceListenConfig.MarketName == market {
			return coinPriceListenConfig
		}
	}
	return nil
}

func (cfg *Config) GetFeeListenConfig(chainId uint64) *FeeListenConfig {
	for _, feeListenConfig := range cfg.FeeListenConfig {
		if feeListenConfig.ChainId == chainId {
			return feeListenConfig
		}
	}
	return nil
}

func NewConfig(filePath string) *Config {
	fileContent, err := basedef.ReadFile(filePath)
	if err != nil {
		logs.Error("NewServiceConfig: failed, err: %s", err)
		return nil
	}
	config := &Config{}
	err = json.Unmarshal(fileContent, config)
	if err != nil {
		logs.Error("NewServiceConfig: failed, err: %s", err)
		return nil
	}

	chainNodeMap := make(map[uint64]*ChainNodes, 0)
	for _, node := range config.ChainNodes {
		chainNodeMap[node.ChainId] = node
	}

	for _, listenConfig := range config.ChainListenConfig {
		if chainNode, ok := chainNodeMap[listenConfig.ChainId]; ok {
			listenConfig.Nodes = chainNode.Nodes
			listenConfig.ExtendNodes = chainNode.ExtendNodes
		}
	}
	for _, listenConfig := range config.FeeListenConfig {
		if chainNode, ok := chainNodeMap[listenConfig.ChainId]; ok {
			listenConfig.Nodes = chainNode.Nodes
		}
	}

	GlobalConfig = config
	initPolyProxy()
	return config
}

func initPolyProxy() {
	if len(PolyProxy) > 0 {
		return
	}
	PolyProxy = make(map[string]bool, 0)
	proxyConfigs := GlobalConfig.ChainListenConfig
	for _, v := range proxyConfigs {
		//some chain only listen,don't need our relayer cross
		if v.ChainId == basedef.SWITCHEO_CROSSCHAIN_ID || v.ChainId == basedef.ZILLIQA_CROSSCHAIN_ID {
			continue
		}
		for _, proxy := range v.ProxyContract {
			PolyProxy[strings.ToUpper(proxy)] = true
			PolyProxy[strings.ToUpper(basedef.HexStringReverse(proxy))] = true
		}
		PolyProxy[strings.ToUpper(v.SwapContract)] = true
		PolyProxy[strings.ToUpper(basedef.HexStringReverse(v.SwapContract))] = true
		for _, contract := range v.NFTProxyContract {
			PolyProxy[strings.ToUpper(contract)] = true
		}
		for _, contract := range v.NFTProxyContract {
			PolyProxy[strings.ToUpper(basedef.HexStringReverse(contract))] = true
		}
	}
	if len(PolyProxy) == 0 {
		panic("init PolyProxy err,polyProxy is nil")
	}
	PolyProxy[""] = true
	logs.Info("init polyProxy:", PolyProxy)
}

type NftConfig struct {
	Description string
	ExternalUrl string
	ColImage    string
	DfImage     string
	ColName     string
	DfName      string
	IpfsUrl     string
	Pwd         string
}

func NewRelayerConfig(filePath string) *RelayerConfig {
	fileContent, err := basedef.ReadFile(filePath)
	if err != nil {
		logs.Error("NewRelayerConfig: failed, err: %s", err)
		return nil
	}
	config := &RelayerConfig{}
	err = json.Unmarshal(fileContent, config)
	if err != nil {
		logs.Error("NewRelayerConfig: failed, err: %s", err)
		return nil
	}
	return config
}
