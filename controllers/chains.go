package controllers

import (
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"

	"github.com/astaxie/beego"
)

var (
	ethereumSdk *chainsdk.EthereumSdkPro
	bscSdk      *chainsdk.EthereumSdkPro
	hecoSdk     *chainsdk.EthereumSdkPro
	okSdk       *chainsdk.EthereumSdkPro
	neoSdk      *chainsdk.NeoSdkPro
	ontologySdk *chainsdk.OntologySdkPro
	config      *conf.Config
)

func init() {
	newChainSdks()
}

func newChainSdks() {
	configFile := beego.AppConfig.String("chain_config")
	config = conf.NewConfig(configFile)
	if config == nil {
		panic("startServer - read config failed!")
	}
	{
		ethereumConfig := config.GetChainListenConfig(basedef.ETHEREUM_CROSSCHAIN_ID)
		if ethereumConfig == nil {
			panic("chain is invalid")
		}
		urls := ethereumConfig.GetNodesUrl()
		ethereumSdk = chainsdk.NewEthereumSdkPro(urls, ethereumConfig.ListenSlot, ethereumConfig.ChainId)
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
		ontConfig := config.GetChainListenConfig(basedef.ONT_CROSSCHAIN_ID)
		if ontConfig == nil {
			panic("chain is invalid")
		}
		urls := ontConfig.GetNodesUrl()
		ontologySdk = chainsdk.NewOntologySdkPro(urls, ontConfig.ListenSlot, ontConfig.ChainId)
	}
}

func GetBalance(chainId uint64, hash string) (*big.Int, error) {
	if chainId == basedef.ETHEREUM_CROSSCHAIN_ID {
		ethereumConfig := config.GetChainListenConfig(basedef.ETHEREUM_CROSSCHAIN_ID)
		if ethereumConfig == nil {
			panic("chain is invalid")
		}
		return ethereumSdk.Erc20Balance(hash, ethereumConfig.ProxyContract)
	}
	if chainId == basedef.BSC_CROSSCHAIN_ID {
		bscConfig := config.GetChainListenConfig(basedef.BSC_CROSSCHAIN_ID)
		if bscConfig == nil {
			panic("chain is invalid")
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
	if chainId == basedef.ONT_CROSSCHAIN_ID {
		ontConfig := config.GetChainListenConfig(basedef.ONT_CROSSCHAIN_ID)
		if ontConfig == nil {
			panic("chain is invalid")
		}
		return ontologySdk.Oep4Balance(hash, ontConfig.ProxyContract)
	}
	return new(big.Int).SetUint64(0), nil
}
