package common

import (
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"strings"

	"github.com/beego/beego/v2/core/logs"
)

var (
	ethereumSdk   *chainsdk.EthereumSdkPro
	pltSdk        *chainsdk.EthereumSdkPro
	bscSdk        *chainsdk.EthereumSdkPro
	hecoSdk       *chainsdk.EthereumSdkPro
	okSdk         *chainsdk.EthereumSdkPro
	neoSdk        *chainsdk.NeoSdkPro
	ontologySdk   *chainsdk.OntologySdkPro
	maticSdk      *chainsdk.EthereumSdkPro
	swthSdk       *chainsdk.SwitcheoSdkPro
	arbitrumSdk   *chainsdk.EthereumSdkPro
	zilliqaSdk    *chainsdk.ZilliqaSdkPro
	xdaiSdk       *chainsdk.EthereumSdkPro
	fantomSdk     *chainsdk.EthereumSdkPro
	avaxSdk       *chainsdk.EthereumSdkPro
	optimisticSdk *chainsdk.EthereumSdkPro
	zionmainSdk   *chainsdk.EthereumSdkPro
	sidechainSdk  *chainsdk.EthereumSdkPro
	kovanSdk      *chainsdk.EthereumSdkPro
	rinkebySdk    *chainsdk.EthereumSdkPro
	goerliSdk     *chainsdk.EthereumSdkPro
	config        *conf.Config
)

func SetupChainsSDK(cfg *conf.Config) {
	if cfg == nil {
		panic("Missing config")
	}
	config = cfg
	newChainSdks(cfg)
}

func newChainSdks(config *conf.Config) {
	for _, cfg := range config.ChainListenConfig {
		switch cfg.ChainId {
		case basedef.ETHEREUM_CROSSCHAIN_ID:
			conf := config.GetChainListenConfig(basedef.ETHEREUM_CROSSCHAIN_ID)
			if conf == nil {
				logs.Error("Missing ETHEREUM chain sdk config")
			}
			urls := conf.GetNodesUrl()
			ethereumSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
		case basedef.ZIONMAIN_CROSSCHAIN_ID:
			conf := config.GetChainListenConfig(basedef.ZIONMAIN_CROSSCHAIN_ID)
			if conf == nil {
				logs.Error("Missing ZIONMAIN chain sdk config")
			}
			urls := conf.GetNodesUrl()
			zionmainSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
		case basedef.SIDECHAIN_CROSSCHAIN_ID:
			conf := config.GetChainListenConfig(basedef.SIDECHAIN_CROSSCHAIN_ID)
			if conf == nil {
				logs.Error("Missing SIDECHAIN chain sdk config")
			}
			urls := conf.GetNodesUrl()
			sidechainSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
		case basedef.MATIC_CROSSCHAIN_ID:
			conf := config.GetChainListenConfig(basedef.MATIC_CROSSCHAIN_ID)
			if conf == nil {
				logs.Error("Missing MATIC chain sdk config")
			}
			urls := conf.GetNodesUrl()
			maticSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
		case basedef.BSC_CROSSCHAIN_ID:
			conf := config.GetChainListenConfig(basedef.BSC_CROSSCHAIN_ID)
			if conf == nil {
				logs.Error("Missing MATIC chain sdk config")
			}
			urls := conf.GetNodesUrl()
			bscSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
		case basedef.HECO_CROSSCHAIN_ID:
			conf := config.GetChainListenConfig(basedef.HECO_CROSSCHAIN_ID)
			if conf == nil {
				logs.Error("Missing HECO chain sdk config")
			}
			urls := conf.GetNodesUrl()
			hecoSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
		case basedef.OK_CROSSCHAIN_ID:
			conf := config.GetChainListenConfig(basedef.OK_CROSSCHAIN_ID)
			if conf == nil {
				logs.Error("Missing OK chain sdk config")
			}
			urls := conf.GetNodesUrl()
			okSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
		case basedef.PLT_CROSSCHAIN_ID:
			conf := config.GetChainListenConfig(basedef.PLT_CROSSCHAIN_ID)
			if conf == nil {
				logs.Error("Missing PLT chain sdk config")
			}
			urls := conf.GetNodesUrl()
			pltSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
		case basedef.KOVAN_CROSSCHAIN_ID:
			if basedef.ENV == basedef.TESTNET {
				conf := config.GetChainListenConfig(basedef.KOVAN_CROSSCHAIN_ID)
				if conf == nil {
					logs.Error("Missing KOVAN chain sdk config")
				}
				urls := conf.GetNodesUrl()
				kovanSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
			}
		case basedef.RINKEBY_CROSSCHAIN_ID:
			if basedef.ENV == basedef.TESTNET {
				conf := config.GetChainListenConfig(basedef.RINKEBY_CROSSCHAIN_ID)
				if conf == nil {
					logs.Error("Missing RINKEBY chain sdk config")
				}
				urls := conf.GetNodesUrl()
				rinkebySdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
			}
		case basedef.GOERLI_CROSSCHAIN_ID:
			if basedef.ENV == basedef.TESTNET {
				conf := config.GetChainListenConfig(basedef.GOERLI_CROSSCHAIN_ID)
				if conf == nil {
					logs.Error("Missing GOERLI chain sdk config")
				}
				urls := conf.GetNodesUrl()
				goerliSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
			}

		}
	}
	//{
	//	neoConfig := config.GetChainListenConfig(basedef.NEO_CROSSCHAIN_ID)
	//	if neoConfig == nil {
	//		panic("chain is invalid")
	//	}
	//	urls := neoConfig.GetNodesUrl()
	//	neoSdk = chainsdk.NewNeoSdkPro(urls, neoConfig.ListenSlot, neoConfig.ChainId)
	//}
	//{
	//	ontConfig := config.GetChainListenConfig(basedef.ONT_CROSSCHAIN_ID)
	//	if ontConfig == nil {
	//		panic("chain is invalid")
	//	}
	//	urls := ontConfig.GetNodesUrl()
	//	ontologySdk = chainsdk.NewOntologySdkPro(urls, ontConfig.ListenSlot, ontConfig.ChainId)
	//}
	//if basedef.ENV == basedef.MAINNET {
	//	swthConfig := config.GetChainListenConfig(basedef.SWITCHEO_CROSSCHAIN_ID)
	//	if swthConfig == nil {
	//		panic("swth chain is invalid")
	//	}
	//	urls := swthConfig.GetNodesUrl()
	//	swthSdk = chainsdk.NewSwitcheoSdkPro(urls, swthConfig.ListenSlot, swthConfig.ChainId)
	//}
	//{
	//	arbitrumConfig := config.GetChainListenConfig(basedef.ARBITRUM_CROSSCHAIN_ID)
	//	if arbitrumConfig == nil {
	//		panic("chain is invalid")
	//	}
	//	urls := arbitrumConfig.GetNodesUrl()
	//	arbitrumSdk = chainsdk.NewEthereumSdkPro(urls, arbitrumConfig.ListenSlot, arbitrumConfig.ChainId)
	//}
	//{
	//	xdaiConfig := config.GetChainListenConfig(basedef.XDAI_CROSSCHAIN_ID)
	//	if xdaiConfig == nil {
	//		panic("chain:XDAI is invalid")
	//	}
	//	urls := xdaiConfig.GetNodesUrl()
	//	xdaiSdk = chainsdk.NewEthereumSdkPro(urls, xdaiConfig.ListenSlot, xdaiConfig.ChainId)
	//}
	//{
	//	zilliqaCfg := config.GetChainListenConfig(basedef.ZILLIQA_CROSSCHAIN_ID)
	//	if zilliqaCfg == nil {
	//		panic("zilliqa GetChainListenConfig chain is invalid")
	//	}
	//	urls := zilliqaCfg.GetNodesUrl()
	//	zilliqaSdk = chainsdk.NewZilliqaSdkPro(urls, zilliqaCfg.ListenSlot, zilliqaCfg.ChainId)
	//}
	//{
	//	fantomConfig := config.GetChainListenConfig(basedef.FANTOM_CROSSCHAIN_ID)
	//	if fantomConfig == nil {
	//		panic("chain is invalid")
	//	}
	//	urls := fantomConfig.GetNodesUrl()
	//	fantomSdk = chainsdk.NewEthereumSdkPro(urls, fantomConfig.ListenSlot, fantomConfig.ChainId)
	//}
	//{
	//	avaxConfig := config.GetChainListenConfig(basedef.AVAX_CROSSCHAIN_ID)
	//	if avaxConfig == nil {
	//		panic("chain is invalid")
	//	}
	//	urls := avaxConfig.GetNodesUrl()
	//	avaxSdk = chainsdk.NewEthereumSdkPro(urls, avaxConfig.ListenSlot, avaxConfig.ChainId)
	//}
	//{
	//	optimisticConfig := config.GetChainListenConfig(basedef.OPTIMISTIC_CROSSCHAIN_ID)
	//	if optimisticConfig == nil {
	//		panic("chain is invalid")
	//	}
	//	urls := optimisticConfig.GetNodesUrl()
	//	optimisticSdk = chainsdk.NewEthereumSdkPro(urls, optimisticConfig.ListenSlot, optimisticConfig.ChainId)
	//}
}

func GetBalance(chainId uint64, hash string) (*big.Int, error) {
	maxBalance := big.NewInt(0)
	maxFun := func(balance *big.Int) {
		if balance.Cmp(maxBalance) > 0 {
			maxBalance = balance
		}
	}
	errMap := make(map[error]bool, 0)
	switch chainId {
	case basedef.ETHEREUM_CROSSCHAIN_ID:
		config := config.GetChainListenConfig(basedef.ETHEREUM_CROSSCHAIN_ID)
		if config == nil {
			panic("chain is invalid")
		}
		for _, v := range config.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := ethereumSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.ZIONMAIN_CROSSCHAIN_ID:
		config := config.GetChainListenConfig(basedef.ZIONMAIN_CROSSCHAIN_ID)
		if config == nil {
			panic("chain is invalid")
		}
		for _, v := range config.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := zionmainSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.SIDECHAIN_CROSSCHAIN_ID:
		config := config.GetChainListenConfig(basedef.SIDECHAIN_CROSSCHAIN_ID)
		if config == nil {
			panic("chain is invalid")
		}
		for _, v := range config.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := sidechainSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.MATIC_CROSSCHAIN_ID:
		config := config.GetChainListenConfig(basedef.MATIC_CROSSCHAIN_ID)
		if config == nil {
			panic("chain is invalid")
		}
		for _, v := range config.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := maticSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.BSC_CROSSCHAIN_ID:
		config := config.GetChainListenConfig(basedef.BSC_CROSSCHAIN_ID)
		if config == nil {
			panic("chain is invalid")
		}
		for _, v := range config.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := bscSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.HECO_CROSSCHAIN_ID:
		config := config.GetChainListenConfig(basedef.HECO_CROSSCHAIN_ID)
		if config == nil {
			panic("chain is invalid")
		}
		for _, v := range config.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := hecoSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.OK_CROSSCHAIN_ID:
		config := config.GetChainListenConfig(basedef.OK_CROSSCHAIN_ID)
		if config == nil {
			panic("chain is invalid")
		}
		for _, v := range config.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := okSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.KOVAN_CROSSCHAIN_ID:
		config := config.GetChainListenConfig(basedef.KOVAN_CROSSCHAIN_ID)
		if config == nil {
			panic("Missing kovan chain sdk config")
		}
		for _, v := range config.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := kovanSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.RINKEBY_CROSSCHAIN_ID:
		config := config.GetChainListenConfig(basedef.RINKEBY_CROSSCHAIN_ID)
		if config == nil {
			panic("Missing rinkeby chain sdk config")
		}
		for _, v := range config.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := rinkebySdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	case basedef.GOERLI_CROSSCHAIN_ID:
		config := config.GetChainListenConfig(basedef.GOERLI_CROSSCHAIN_ID)
		if config == nil {
			panic("Missing goerli chain sdk config")
		}
		for _, v := range config.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := goerliSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	default:
		return new(big.Int).SetUint64(0), nil
	}

	if maxBalance.Cmp(big.NewInt(0)) > 0 {
		return maxBalance, nil
	}
	var err error
	for k, _ := range errMap {
		if k == nil {
			return new(big.Int).SetUint64(0), nil
		} else {
			err = k
		}
	}
	return new(big.Int).SetUint64(0), err

}

//if chainId == basedef.NEO_CROSSCHAIN_ID {
//	neoConfig := config.GetChainListenConfig(basedef.NEO_CROSSCHAIN_ID)
//	if neoConfig == nil {
//		panic("chain is invalid")
//	}
//	return neoSdk.Nep5Balance(hash, neoConfig.ProxyContract)
//}
//if chainId == basedef.ONT_CROSSCHAIN_ID {
//	ontConfig := config.GetChainListenConfig(basedef.ONT_CROSSCHAIN_ID)
//	if ontConfig == nil {
//		panic("chain is invalid")
//	}
//	return ontologySdk.Oep4Balance(hash, ontConfig.ProxyContract)
//}
//if chainId == basedef.ARBITRUM_CROSSCHAIN_ID {
//	arbitrumConfig := config.GetChainListenConfig(basedef.ARBITRUM_CROSSCHAIN_ID)
//	if arbitrumConfig == nil {
//		panic("chain is invalid")
//	}
//	return arbitrumSdk.Erc20Balance(hash, arbitrumConfig.ProxyContract)
//}
//if chainId == basedef.XDAI_CROSSCHAIN_ID {
//	xdaiConfig := config.GetChainListenConfig(basedef.XDAI_CROSSCHAIN_ID)
//	if xdaiConfig == nil {
//		panic("chain is invalid")
//	}
//	return xdaiSdk.Erc20Balance(hash, xdaiConfig.ProxyContract)
//}
//if chainId == basedef.ZILLIQA_CROSSCHAIN_ID {
//	zilliqaCfg := config.GetChainListenConfig(basedef.ZILLIQA_CROSSCHAIN_ID)
//	if zilliqaCfg == nil {
//		panic("zilliqa GetChainListenConfig chain is invalid")
//	}
//	return zilliqaSdk.Erc20Balance(hash, zilliqaCfg.ProxyContract)
//}
//if chainId == basedef.FANTOM_CROSSCHAIN_ID {
//	fantomConfig := config.GetChainListenConfig(basedef.FANTOM_CROSSCHAIN_ID)
//	if fantomConfig == nil {
//		panic("chain is invalid")
//	}
//	return fantomSdk.Erc20Balance(hash, fantomConfig.ProxyContract)
//}
//if chainId == basedef.AVAX_CROSSCHAIN_ID {
//	avaxConfig := config.GetChainListenConfig(basedef.AVAX_CROSSCHAIN_ID)
//	if avaxConfig == nil {
//		panic("chain is invalid")
//	}
//	return avaxSdk.Erc20Balance(hash, avaxConfig.ProxyContract)
//}
//if chainId == basedef.OPTIMISTIC_CROSSCHAIN_ID {
//	optimisticConfig := config.GetChainListenConfig(basedef.OPTIMISTIC_CROSSCHAIN_ID)
//	if optimisticConfig == nil {
//		panic("chain is invalid")
//	}
//	return optimisticSdk.Erc20Balance(hash, optimisticConfig.ProxyContract)
//}
/*if chainId == basedef.PLT_CROSSCHAIN_ID {
  	conf := config.GetChainListenConfig(basedef.PLT_CROSSCHAIN_ID)
  	if conf == nil {
  		panic("chain is invalid")
  	}
  	return pltSdk.Erc20Balance(hash,conf.ProxyContract)
  }
*/

func GetTotalSupply(chainId uint64, hash string) (*big.Int, error) {
	switch chainId {
	case basedef.ETHEREUM_CROSSCHAIN_ID:
		ethereumConfig := config.GetChainListenConfig(basedef.ETHEREUM_CROSSCHAIN_ID)
		if ethereumConfig == nil {
			panic("chain is invalid")
		}
		return ethereumSdk.Erc20TotalSupply(hash)
	case basedef.ZIONMAIN_CROSSCHAIN_ID:
		zionmainConfig := config.GetChainListenConfig(basedef.ZIONMAIN_CROSSCHAIN_ID)
		if zionmainConfig == nil {
			panic("chain is invalid")
		}
		return zionmainSdk.Erc20TotalSupply(hash)
	case basedef.SIDECHAIN_CROSSCHAIN_ID:
		sidechainConfig := config.GetChainListenConfig(basedef.SIDECHAIN_CROSSCHAIN_ID)
		if sidechainConfig == nil {
			panic("chain is invalid")
		}
		return sidechainSdk.Erc20TotalSupply(hash)
	case basedef.MATIC_CROSSCHAIN_ID:
		maticConfig := config.GetChainListenConfig(basedef.MATIC_CROSSCHAIN_ID)
		if maticConfig == nil {
			panic("chain is invalid")
		}
		return maticSdk.Erc20TotalSupply(hash)
	case basedef.BSC_CROSSCHAIN_ID:
		bscConfig := config.GetChainListenConfig(basedef.BSC_CROSSCHAIN_ID)
		if bscConfig == nil {
			panic("chain is invalid")
		}
		return bscSdk.Erc20TotalSupply(hash)
	case basedef.HECO_CROSSCHAIN_ID:
		hecoConfig := config.GetChainListenConfig(basedef.HECO_CROSSCHAIN_ID)
		if hecoConfig == nil {
			panic("chain is invalid")
		}
		return hecoSdk.Erc20TotalSupply(hash)
	case basedef.OK_CROSSCHAIN_ID:
		okConfig := config.GetChainListenConfig(basedef.OK_CROSSCHAIN_ID)
		if okConfig == nil {
			panic("chain is invalid")
		}
		return okSdk.Erc20TotalSupply(hash)
	case basedef.KOVAN_CROSSCHAIN_ID:
		config := config.GetChainListenConfig(basedef.KOVAN_CROSSCHAIN_ID)
		if config == nil {
			panic("Missing kovan chain sdk config")
		}
		return kovanSdk.Erc20TotalSupply(hash)
	case basedef.RINKEBY_CROSSCHAIN_ID:
		config := config.GetChainListenConfig(basedef.RINKEBY_CROSSCHAIN_ID)
		if config == nil {
			panic("Missing rinkeby chain sdk config")
		}
		return rinkebySdk.Erc20TotalSupply(hash)

	case basedef.GOERLI_CROSSCHAIN_ID:
		config := config.GetChainListenConfig(basedef.GOERLI_CROSSCHAIN_ID)
		if config == nil {
			panic("Missing goerli chain sdk config")
		}
		return goerliSdk.Erc20TotalSupply(hash)
	default:
		return new(big.Int).SetUint64(0), nil
	}
	//if chainId == basedef.NEO_CROSSCHAIN_ID {
	//	neoConfig := config.GetChainListenConfig(basedef.NEO_CROSSCHAIN_ID)
	//	if neoConfig == nil {
	//		panic("chain is invalid")
	//	}
	//	return neoSdk.Nep5TotalSupply(hash)
	//}
	//if chainId == basedef.ONT_CROSSCHAIN_ID {
	//	ontConfig := config.GetChainListenConfig(basedef.ONT_CROSSCHAIN_ID)
	//	if ontConfig == nil {
	//		panic("chain is invalid")
	//	}
	//	return ontologySdk.Oep4TotalSupply(hash, ontConfig.ProxyContract)
	//}
	//if chainId == basedef.ARBITRUM_CROSSCHAIN_ID {
	//	arbitrumConfig := config.GetChainListenConfig(basedef.ARBITRUM_CROSSCHAIN_ID)
	//	if arbitrumConfig == nil {
	//		panic("chain is invalid")
	//	}
	//	return arbitrumSdk.Erc20TotalSupply(hash)
	//}
	//if chainId == basedef.XDAI_CROSSCHAIN_ID {
	//	xdaiConfig := config.GetChainListenConfig(basedef.XDAI_CROSSCHAIN_ID)
	//	if xdaiConfig == nil {
	//		panic("chain is invalid")
	//	}
	//	return xdaiSdk.Erc20TotalSupply(hash)
	//}
	//if chainId == basedef.FANTOM_CROSSCHAIN_ID {
	//	fantomConfig := config.GetChainListenConfig(basedef.FANTOM_CROSSCHAIN_ID)
	//	if fantomConfig == nil {
	//		panic("chain is invalid")
	//	}
	//	return fantomSdk.Erc20TotalSupply(hash)
	//}
	//if chainId == basedef.AVAX_CROSSCHAIN_ID {
	//	avaxConfig := config.GetChainListenConfig(basedef.AVAX_CROSSCHAIN_ID)
	//	if avaxConfig == nil {
	//		panic("chain is invalid")
	//	}
	//	return avaxSdk.Erc20TotalSupply(hash)
	//}
	//if chainId == basedef.OPTIMISTIC_CROSSCHAIN_ID {
	//	optimisticConfig := config.GetChainListenConfig(basedef.OPTIMISTIC_CROSSCHAIN_ID)
	//	if optimisticConfig == nil {
	//		panic("chain is invalid")
	//	}
	//	return optimisticSdk.Erc20TotalSupply(hash)
	//}

}

type ProxyBalance struct {
	Amount    *big.Int
	ItemName  string
	ItemProxy string
}

func GetProxyBalance(chainId uint64, hash string, proxy string) (*big.Int, error) {
	switch chainId {
	case basedef.ETHEREUM_CROSSCHAIN_ID:
		return ethereumSdk.Erc20Balance(hash, proxy)
	case basedef.ZIONMAIN_CROSSCHAIN_ID:
		return zionmainSdk.Erc20Balance(hash, proxy)
	case basedef.SIDECHAIN_CROSSCHAIN_ID:
		return sidechainSdk.Erc20Balance(hash, proxy)
	case basedef.MATIC_CROSSCHAIN_ID:
		return maticSdk.Erc20Balance(hash, proxy)
	case basedef.BSC_CROSSCHAIN_ID:
		return bscSdk.Erc20Balance(hash, proxy)
	case basedef.HECO_CROSSCHAIN_ID:
		return hecoSdk.Erc20Balance(hash, proxy)
	case basedef.OK_CROSSCHAIN_ID:
		return okSdk.Erc20Balance(hash, proxy)
	case basedef.NEO_CROSSCHAIN_ID:
		return neoSdk.Nep5Balance(hash, proxy)
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
	case basedef.KOVAN_CROSSCHAIN_ID:
		return kovanSdk.Erc20Balance(hash, proxy)
	case basedef.RINKEBY_CROSSCHAIN_ID:
		return rinkebySdk.Erc20Balance(hash, proxy)
	case basedef.GOERLI_CROSSCHAIN_ID:
		return goerliSdk.Erc20Balance(hash, proxy)
	default:
		return new(big.Int).SetUint64(0), nil
	}
}
