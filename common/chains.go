package common

import (
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"

	"github.com/beego/beego/v2/core/logs"
)

var (
	ethereumSdk   *chainsdk.EthereumSdkPro
	pltSdk        *chainsdk.EthereumSdkPro
	bscSdk        *chainsdk.EthereumSdkPro
	hecoSdk       *chainsdk.EthereumSdkPro
	okSdk         *chainsdk.EthereumSdkPro
	neoSdk        *chainsdk.NeoSdkPro
	neo3Sdk        *chainsdk.Neo3SdkPro
	ontologySdk   *chainsdk.OntologySdkPro
	maticSdk      *chainsdk.EthereumSdkPro
	swthSdk       *chainsdk.SwitcheoSdkPro
	arbitrumSdk   *chainsdk.EthereumSdkPro
	zilliqaSdk    *chainsdk.ZilliqaSdkPro
	xdaiSdk       *chainsdk.EthereumSdkPro
	fantomSdk     *chainsdk.EthereumSdkPro
	avaxSdk       *chainsdk.EthereumSdkPro
	optimisticSdk *chainsdk.EthereumSdkPro
	config        *conf.Config
)

var getfee0num uint64
var getfeename string

var bscgetfee0num uint64
var bscgetfeename string

func SetupChainsSDK(cfg *conf.Config, a uint64, b string) {
	if cfg == nil {
		panic("Missing config")
	}
	config = cfg
	getfee0num = a
	getfeename = b
	bscgetfee0num = a
	bscgetfeename = b
	newChainSdks(cfg)
}

func newChainSdks(config *conf.Config) {
	{
		ethereumConfig := config.GetChainListenConfig(basedef.ETHEREUM_CROSSCHAIN_ID)
		if ethereumConfig == nil {
			panic("chain is invalid")
		}
		urls := ethereumConfig.GetNodesUrl()
		ethereumSdk = chainsdk.NewEthereumSdkPro(urls, ethereumConfig.ListenSlot, ethereumConfig.ChainId)
	}
	{
		maticConfig := config.GetChainListenConfig(basedef.MATIC_CROSSCHAIN_ID)
		if maticConfig == nil {
			panic("chain is invalid")
		}
		urls := maticConfig.GetNodesUrl()
		maticSdk = chainsdk.NewEthereumSdkPro(urls, maticConfig.ListenSlot, maticConfig.ChainId)
	}
	{
		bscConfig := config.GetChainListenConfig(basedef.BSC_CROSSCHAIN_ID)
		if bscConfig == nil {
			panic("chain is invalid")
		}
		urls := bscConfig.GetNodesUrl()
		bscSdk = chainsdk.NewEthereumSdkPro(urls, bscConfig.ListenSlot, bscConfig.ChainId)
	}
	{
		hecoConfig := config.GetChainListenConfig(basedef.HECO_CROSSCHAIN_ID)
		if hecoConfig == nil {
			panic("chain is invalid")
		}
		urls := hecoConfig.GetNodesUrl()
		hecoSdk = chainsdk.NewEthereumSdkPro(urls, hecoConfig.ListenSlot, hecoConfig.ChainId)
	}
	{
		okConfig := config.GetChainListenConfig(basedef.OK_CROSSCHAIN_ID)
		if okConfig == nil {
			panic("chain is invalid")
		}
		urls := okConfig.GetNodesUrl()
		okSdk = chainsdk.NewEthereumSdkPro(urls, okConfig.ListenSlot, okConfig.ChainId)
	}
	{
		neoConfig := config.GetChainListenConfig(basedef.NEO_CROSSCHAIN_ID)
		if neoConfig == nil {
			panic("chain is invalid")
		}
		urls := neoConfig.GetNodesUrl()
		neoSdk = chainsdk.NewNeoSdkPro(urls, neoConfig.ListenSlot, neoConfig.ChainId)
	}
	{
		neo3Config := config.GetChainListenConfig(basedef.NEO3_CROSSCHAIN_ID)
		if neo3Config == nil {
			panic("chain is invalid")
		}
		urls := neo3Config.GetNodesUrl()
		neo3Sdk = chainsdk.NewNeo3SdkPro(urls, neo3Config.ListenSlot, neo3Config.ChainId)
	}
	{
		ontConfig := config.GetChainListenConfig(basedef.ONT_CROSSCHAIN_ID)
		if ontConfig == nil {
			panic("chain is invalid")
		}
		urls := ontConfig.GetNodesUrl()
		ontologySdk = chainsdk.NewOntologySdkPro(urls, ontConfig.ListenSlot, ontConfig.ChainId)
	}
	if basedef.ENV == basedef.MAINNET {
		swthConfig := config.GetChainListenConfig(basedef.SWITCHEO_CROSSCHAIN_ID)
		if swthConfig == nil {
			panic("swth chain is invalid")
		}
		urls := swthConfig.GetNodesUrl()
		swthSdk = chainsdk.NewSwitcheoSdkPro(urls, swthConfig.ListenSlot, swthConfig.ChainId)
	}
	{
		conf := config.GetChainListenConfig(basedef.PLT_CROSSCHAIN_ID)
		if conf != nil {
			urls := conf.GetNodesUrl()
			pltSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
		} else {
			logs.Error("Missing plt chain sdk config")
		}
	}
	{
		arbitrumConfig := config.GetChainListenConfig(basedef.ARBITRUM_CROSSCHAIN_ID)
		if arbitrumConfig == nil {
			panic("chain is invalid")
		}
		urls := arbitrumConfig.GetNodesUrl()
		arbitrumSdk = chainsdk.NewEthereumSdkPro(urls, arbitrumConfig.ListenSlot, arbitrumConfig.ChainId)
	}
	{
		xdaiConfig := config.GetChainListenConfig(basedef.XDAI_CROSSCHAIN_ID)
		if xdaiConfig == nil {
			panic("chain:XDAI is invalid")
		}
		urls := xdaiConfig.GetNodesUrl()
		xdaiSdk = chainsdk.NewEthereumSdkPro(urls, xdaiConfig.ListenSlot, xdaiConfig.ChainId)
	}
	{
		zilliqaCfg := config.GetChainListenConfig(basedef.ZILLIQA_CROSSCHAIN_ID)
		if zilliqaCfg == nil {
			panic("zilliqa GetChainListenConfig chain is invalid")
		}
		urls := zilliqaCfg.GetNodesUrl()
		zilliqaSdk = chainsdk.NewZilliqaSdkPro(urls, zilliqaCfg.ListenSlot, zilliqaCfg.ChainId)
	}
	{
		fantomConfig := config.GetChainListenConfig(basedef.FANTOM_CROSSCHAIN_ID)
		if fantomConfig == nil {
			panic("chain is invalid")
		}
		urls := fantomConfig.GetNodesUrl()
		fantomSdk = chainsdk.NewEthereumSdkPro(urls, fantomConfig.ListenSlot, fantomConfig.ChainId)
	}
	{
		avaxConfig := config.GetChainListenConfig(basedef.AVAX_CROSSCHAIN_ID)
		if avaxConfig == nil {
			panic("chain is invalid")
		}
		urls := avaxConfig.GetNodesUrl()
		avaxSdk = chainsdk.NewEthereumSdkPro(urls, avaxConfig.ListenSlot, avaxConfig.ChainId)
	}
	{
		optimisticConfig := config.GetChainListenConfig(basedef.OPTIMISTIC_CROSSCHAIN_ID)
		if optimisticConfig == nil {
			panic("chain is invalid")
		}
		urls := optimisticConfig.GetNodesUrl()
		optimisticSdk = chainsdk.NewEthereumSdkPro(urls, optimisticConfig.ListenSlot, optimisticConfig.ChainId)
	}
}

func GetBalance(chainId uint64, hash string) (*big.Int, error) {
	if chainId == basedef.ETHEREUM_CROSSCHAIN_ID {
		ethereumConfig := config.GetChainListenConfig(basedef.ETHEREUM_CROSSCHAIN_ID)
		if ethereumConfig == nil {
			panic("chain is invalid")
		}
		if hash == "0000000000000000000000000000000000000000" {
			getfee0num++
			logs.Error(getfeename+"getfeename is:", getfee0num)
		}
		return ethereumSdk.Erc20Balance(hash, ethereumConfig.ProxyContract)
	}
	if chainId == basedef.MATIC_CROSSCHAIN_ID {
		maticConfig := config.GetChainListenConfig(basedef.MATIC_CROSSCHAIN_ID)
		if maticConfig == nil {
			panic("chain is invalid")
		}
		return maticSdk.Erc20Balance(hash, maticConfig.ProxyContract)
	}
	if chainId == basedef.BSC_CROSSCHAIN_ID {
		bscConfig := config.GetChainListenConfig(basedef.BSC_CROSSCHAIN_ID)
		if bscConfig == nil {
			panic("chain is invalid")
		}
		if hash == "0000000000000000000000000000000000000000" {
			bscgetfee0num++
			logs.Error(getfeename+"bscgetfeename is:", bscgetfee0num)
		}
		return bscSdk.Erc20Balance(hash, bscConfig.ProxyContract)
	}
	if chainId == basedef.HECO_CROSSCHAIN_ID {
		hecoConfig := config.GetChainListenConfig(basedef.HECO_CROSSCHAIN_ID)
		if hecoConfig == nil {
			panic("chain is invalid")
		}
		return hecoSdk.Erc20Balance(hash, hecoConfig.ProxyContract)
	}
	if chainId == basedef.OK_CROSSCHAIN_ID {
		okConfig := config.GetChainListenConfig(basedef.OK_CROSSCHAIN_ID)
		if okConfig == nil {
			panic("chain is invalid")
		}
		return okSdk.Erc20Balance(hash, okConfig.ProxyContract)
	}
	if chainId == basedef.NEO_CROSSCHAIN_ID {
		neoConfig := config.GetChainListenConfig(basedef.NEO_CROSSCHAIN_ID)
		if neoConfig == nil {
			panic("chain is invalid")
		}
		return neoSdk.Nep5Balance(hash, neoConfig.ProxyContract)
	}
	if chainId == basedef.NEO3_CROSSCHAIN_ID {
		neo3Config := config.GetChainListenConfig(basedef.NEO3_CROSSCHAIN_ID)
		if neo3Config == nil {
			panic("chain is invalid")
		}
		logs.Info("get neo3 balance. hash=%s, ProxyContract=%s", hash, neo3Config.ProxyContract)
		return neo3Sdk.Nep17Balance(hash, neo3Config.ProxyContract)
	}
	if chainId == basedef.ONT_CROSSCHAIN_ID {
		ontConfig := config.GetChainListenConfig(basedef.ONT_CROSSCHAIN_ID)
		if ontConfig == nil {
			panic("chain is invalid")
		}
		return ontologySdk.Oep4Balance(hash, ontConfig.ProxyContract)
	}
	if chainId == basedef.MATIC_CROSSCHAIN_ID {
		maticConfig := config.GetChainListenConfig(basedef.MATIC_CROSSCHAIN_ID)
		if maticConfig == nil {
			panic("chain is invalid")
		}
		return maticSdk.Erc20Balance(hash, maticConfig.ProxyContract)
	}
	if chainId == basedef.ARBITRUM_CROSSCHAIN_ID {
		arbitrumConfig := config.GetChainListenConfig(basedef.ARBITRUM_CROSSCHAIN_ID)
		if arbitrumConfig == nil {
			panic("chain is invalid")
		}
		return arbitrumSdk.Erc20Balance(hash, arbitrumConfig.ProxyContract)
	}
	if chainId == basedef.XDAI_CROSSCHAIN_ID {
		xdaiConfig := config.GetChainListenConfig(basedef.XDAI_CROSSCHAIN_ID)
		if xdaiConfig == nil {
			panic("chain is invalid")
		}
		return xdaiSdk.Erc20Balance(hash, xdaiConfig.ProxyContract)
	}
	if chainId == basedef.ZILLIQA_CROSSCHAIN_ID {
		zilliqaCfg := config.GetChainListenConfig(basedef.ZILLIQA_CROSSCHAIN_ID)
		if zilliqaCfg == nil {
			panic("zilliqa GetChainListenConfig chain is invalid")
		}
		return zilliqaSdk.Erc20Balance(hash, zilliqaCfg.ProxyContract)
	}
	if chainId == basedef.FANTOM_CROSSCHAIN_ID {
		fantomConfig := config.GetChainListenConfig(basedef.FANTOM_CROSSCHAIN_ID)
		if fantomConfig == nil {
			panic("chain is invalid")
		}
		return fantomSdk.Erc20Balance(hash, fantomConfig.ProxyContract)
	}
	if chainId == basedef.AVAX_CROSSCHAIN_ID {
		avaxConfig := config.GetChainListenConfig(basedef.AVAX_CROSSCHAIN_ID)
		if avaxConfig == nil {
			panic("chain is invalid")
		}
		return avaxSdk.Erc20Balance(hash, avaxConfig.ProxyContract)
	}
	if chainId == basedef.OPTIMISTIC_CROSSCHAIN_ID {
		optimisticConfig := config.GetChainListenConfig(basedef.OPTIMISTIC_CROSSCHAIN_ID)
		if optimisticConfig == nil {
			panic("chain is invalid")
		}
		return optimisticSdk.Erc20Balance(hash, optimisticConfig.ProxyContract)
	}
	/*if chainId == basedef.PLT_CROSSCHAIN_ID {
		conf := config.GetChainListenConfig(basedef.PLT_CROSSCHAIN_ID)
		if conf == nil {
			panic("chain is invalid")
		}
		return pltSdk.Erc20Balance(hash,conf.ProxyContract)
	}
	*/
	return new(big.Int).SetUint64(0), nil
}

func GetTotalSupply(chainId uint64, hash string) (*big.Int, error) {
	if chainId == basedef.ETHEREUM_CROSSCHAIN_ID {
		ethereumConfig := config.GetChainListenConfig(basedef.ETHEREUM_CROSSCHAIN_ID)
		if ethereumConfig == nil {
			panic("chain is invalid")
		}
		return ethereumSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.BSC_CROSSCHAIN_ID {
		bscConfig := config.GetChainListenConfig(basedef.BSC_CROSSCHAIN_ID)
		if bscConfig == nil {
			panic("chain is invalid")
		}
		return bscSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.HECO_CROSSCHAIN_ID {
		hecoConfig := config.GetChainListenConfig(basedef.HECO_CROSSCHAIN_ID)
		if hecoConfig == nil {
			panic("chain is invalid")
		}
		return hecoSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.OK_CROSSCHAIN_ID {
		okConfig := config.GetChainListenConfig(basedef.OK_CROSSCHAIN_ID)
		if okConfig == nil {
			panic("chain is invalid")
		}
		return okSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.NEO_CROSSCHAIN_ID {
		neoConfig := config.GetChainListenConfig(basedef.NEO_CROSSCHAIN_ID)
		if neoConfig == nil {
			panic("chain is invalid")
		}
		return neoSdk.Nep5TotalSupply(hash)
	}
	if chainId == basedef.NEO3_CROSSCHAIN_ID {
		neo3Config := config.GetChainListenConfig(basedef.NEO3_CROSSCHAIN_ID)
		if neo3Config == nil {
			panic("chain is invalid")
		}
		return neo3Sdk.Nep17TotalSupply(hash)
	}
	if chainId == basedef.ONT_CROSSCHAIN_ID {
		ontConfig := config.GetChainListenConfig(basedef.ONT_CROSSCHAIN_ID)
		if ontConfig == nil {
			panic("chain is invalid")
		}
		return ontologySdk.Oep4TotalSupply(hash, ontConfig.ProxyContract)
	}
	if chainId == basedef.MATIC_CROSSCHAIN_ID {
		maticConfig := config.GetChainListenConfig(basedef.MATIC_CROSSCHAIN_ID)
		if maticConfig == nil {
			panic("chain is invalid")
		}
		return maticSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.ARBITRUM_CROSSCHAIN_ID {
		arbitrumConfig := config.GetChainListenConfig(basedef.ARBITRUM_CROSSCHAIN_ID)
		if arbitrumConfig == nil {
			panic("chain is invalid")
		}
		return arbitrumSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.XDAI_CROSSCHAIN_ID {
		xdaiConfig := config.GetChainListenConfig(basedef.XDAI_CROSSCHAIN_ID)
		if xdaiConfig == nil {
			panic("chain is invalid")
		}
		return xdaiSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.FANTOM_CROSSCHAIN_ID {
		fantomConfig := config.GetChainListenConfig(basedef.FANTOM_CROSSCHAIN_ID)
		if fantomConfig == nil {
			panic("chain is invalid")
		}
		return fantomSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.AVAX_CROSSCHAIN_ID {
		avaxConfig := config.GetChainListenConfig(basedef.AVAX_CROSSCHAIN_ID)
		if avaxConfig == nil {
			panic("chain is invalid")
		}
		return avaxSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.OPTIMISTIC_CROSSCHAIN_ID {
		optimisticConfig := config.GetChainListenConfig(basedef.OPTIMISTIC_CROSSCHAIN_ID)
		if optimisticConfig == nil {
			panic("chain is invalid")
		}
		return optimisticSdk.Erc20TotalSupply(hash)
	}
	return new(big.Int).SetUint64(0), nil
}

type ProxyBalance struct {
	Amount    *big.Int
	ItemName  string
	ItemProxy string
}

func GetProxyBalance(chainId uint64, hash string, proxy string) (*big.Int, error) {
	switch chainId {
	case basedef.ETHEREUM_CROSSCHAIN_ID:
		if hash == "0000000000000000000000000000000000000000" {
			getfee0num++
			logs.Error(getfeename+"getfeename is:", getfee0num)
		}
		return ethereumSdk.Erc20Balance(hash, proxy)
	case basedef.MATIC_CROSSCHAIN_ID:
		return maticSdk.Erc20Balance(hash, proxy)
	case basedef.BSC_CROSSCHAIN_ID:
		if hash == "0000000000000000000000000000000000000000" {
			bscgetfee0num++
			logs.Error(getfeename+"bscgetfeename is:", bscgetfee0num)
		}
		return bscSdk.Erc20Balance(hash, proxy)
	case basedef.HECO_CROSSCHAIN_ID:
		return hecoSdk.Erc20Balance(hash, proxy)
	case basedef.OK_CROSSCHAIN_ID:
		return okSdk.Erc20Balance(hash, proxy)
	case basedef.NEO_CROSSCHAIN_ID:
		return neoSdk.Nep5Balance(hash, proxy)
	case basedef.NEO3_CROSSCHAIN_ID:
		return neo3Sdk.Nep17Balance(hash, proxy)
	case basedef.ONT_CROSSCHAIN_ID:
		return ontologySdk.Oep4Balance(hash, proxy)
	case basedef.ARBITRUM_CROSSCHAIN_ID:
		return arbitrumSdk.Erc20Balance(hash, proxy)
	case basedef.XDAI_CROSSCHAIN_ID:
		return xdaiSdk.Erc20Balance(hash, proxy)
	case basedef.ZILLIQA_CROSSCHAIN_ID:
		return zilliqaSdk.Erc20Balance(hash, proxy)
	case basedef.FANTOM_CROSSCHAIN_ID:
		return fantomSdk.Erc20Balance(hash, proxy)
	case basedef.AVAX_CROSSCHAIN_ID:
		return avaxSdk.Erc20Balance(hash, proxy)
	case basedef.OPTIMISTIC_CROSSCHAIN_ID:
		return optimisticSdk.Erc20Balance(hash, proxy)
	default:
		return new(big.Int).SetUint64(0), nil
	}
}
