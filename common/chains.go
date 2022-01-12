package common

import (
	"fmt"
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
	neo3Sdk       *chainsdk.Neo3SdkPro
	ontologySdk   *chainsdk.OntologySdkPro
	maticSdk      *chainsdk.EthereumSdkPro
	swthSdk       *chainsdk.SwitcheoSdkPro
	arbitrumSdk   *chainsdk.EthereumSdkPro
	zilliqaSdk    *chainsdk.ZilliqaSdkPro
	xdaiSdk       *chainsdk.EthereumSdkPro
	fantomSdk     *chainsdk.EthereumSdkPro
	avaxSdk       *chainsdk.EthereumSdkPro
	optimisticSdk *chainsdk.EthereumSdkPro
	metisSdk      *chainsdk.EthereumSdkPro
	pixieSdk      *chainsdk.EthereumSdkPro
	rinkebySdk    *chainsdk.EthereumSdkPro
	sdkMap        map[uint64]interface{}
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
	sdkMap = make(map[uint64]interface{}, 0)
	{
		ethereumConfig := config.GetChainListenConfig(basedef.ETHEREUM_CROSSCHAIN_ID)
		if ethereumConfig == nil {
			panic("chain is invalid")
		}
		urls := ethereumConfig.GetNodesUrl()
		ethereumSdk = chainsdk.NewEthereumSdkPro(urls, ethereumConfig.ListenSlot, ethereumConfig.ChainId)
		sdkMap[basedef.ETHEREUM_CROSSCHAIN_ID] = ethereumSdk
	}
	{
		maticConfig := config.GetChainListenConfig(basedef.MATIC_CROSSCHAIN_ID)
		if maticConfig == nil {
			panic("chain is invalid")
		}
		urls := maticConfig.GetNodesUrl()
		maticSdk = chainsdk.NewEthereumSdkPro(urls, maticConfig.ListenSlot, maticConfig.ChainId)
		sdkMap[basedef.MATIC_CROSSCHAIN_ID] = maticSdk
	}
	{
		bscConfig := config.GetChainListenConfig(basedef.BSC_CROSSCHAIN_ID)
		if bscConfig == nil {
			panic("chain is invalid")
		}
		urls := bscConfig.GetNodesUrl()
		bscSdk = chainsdk.NewEthereumSdkPro(urls, bscConfig.ListenSlot, bscConfig.ChainId)
		sdkMap[basedef.BSC_CROSSCHAIN_ID] = bscSdk
	}
	{
		hecoConfig := config.GetChainListenConfig(basedef.HECO_CROSSCHAIN_ID)
		if hecoConfig == nil {
			panic("chain is invalid")
		}
		urls := hecoConfig.GetNodesUrl()
		hecoSdk = chainsdk.NewEthereumSdkPro(urls, hecoConfig.ListenSlot, hecoConfig.ChainId)
		sdkMap[basedef.HECO_CROSSCHAIN_ID] = hecoSdk
	}
	{
		okConfig := config.GetChainListenConfig(basedef.OK_CROSSCHAIN_ID)
		if okConfig == nil {
			panic("chain is invalid")
		}
		urls := okConfig.GetNodesUrl()
		okSdk = chainsdk.NewEthereumSdkPro(urls, okConfig.ListenSlot, okConfig.ChainId)
		sdkMap[basedef.OK_CROSSCHAIN_ID] = okSdk
	}
	{
		neoConfig := config.GetChainListenConfig(basedef.NEO_CROSSCHAIN_ID)
		if neoConfig == nil {
			panic("chain is invalid")
		}
		urls := neoConfig.GetNodesUrl()
		neoSdk = chainsdk.NewNeoSdkPro(urls, neoConfig.ListenSlot, neoConfig.ChainId)
		sdkMap[basedef.NEO_CROSSCHAIN_ID] = neoSdk
	}
	{
		neo3Config := config.GetChainListenConfig(basedef.NEO3_CROSSCHAIN_ID)
		if neo3Config == nil {
			panic("chain is invalid")
		}
		urls := neo3Config.GetNodesUrl()
		neo3Sdk = chainsdk.NewNeo3SdkPro(urls, neo3Config.ListenSlot, neo3Config.ChainId)
		sdkMap[basedef.NEO3_CROSSCHAIN_ID] = neo3Sdk
	}
	{
		ontConfig := config.GetChainListenConfig(basedef.ONT_CROSSCHAIN_ID)
		if ontConfig == nil {
			panic("chain is invalid")
		}
		urls := ontConfig.GetNodesUrl()
		ontologySdk = chainsdk.NewOntologySdkPro(urls, ontConfig.ListenSlot, ontConfig.ChainId)
		sdkMap[basedef.ONT_CROSSCHAIN_ID] = ontologySdk
	}
	if basedef.ENV == basedef.MAINNET {
		swthConfig := config.GetChainListenConfig(basedef.SWITCHEO_CROSSCHAIN_ID)
		if swthConfig == nil {
			panic("swth chain is invalid")
		}
		urls := swthConfig.GetNodesUrl()
		swthSdk = chainsdk.NewSwitcheoSdkPro(urls, swthConfig.ListenSlot, swthConfig.ChainId)
		sdkMap[basedef.SWITCHEO_CROSSCHAIN_ID] = swthSdk
	}
	{
		conf := config.GetChainListenConfig(basedef.PLT_CROSSCHAIN_ID)
		if conf != nil {
			urls := conf.GetNodesUrl()
			pltSdk = chainsdk.NewEthereumSdkPro(urls, conf.ListenSlot, conf.ChainId)
			sdkMap[basedef.PLT_CROSSCHAIN_ID] = pltSdk
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
		sdkMap[basedef.ARBITRUM_CROSSCHAIN_ID] = arbitrumSdk
	}
	{
		xdaiConfig := config.GetChainListenConfig(basedef.XDAI_CROSSCHAIN_ID)
		if xdaiConfig == nil {
			panic("chain:XDAI is invalid")
		}
		urls := xdaiConfig.GetNodesUrl()
		xdaiSdk = chainsdk.NewEthereumSdkPro(urls, xdaiConfig.ListenSlot, xdaiConfig.ChainId)
		sdkMap[basedef.XDAI_CROSSCHAIN_ID] = xdaiSdk
	}
	{
		zilliqaCfg := config.GetChainListenConfig(basedef.ZILLIQA_CROSSCHAIN_ID)
		if zilliqaCfg == nil {
			panic("zilliqa GetChainListenConfig chain is invalid")
		}
		urls := zilliqaCfg.GetNodesUrl()
		zilliqaSdk = chainsdk.NewZilliqaSdkPro(urls, zilliqaCfg.ListenSlot, zilliqaCfg.ChainId)
		sdkMap[basedef.ZILLIQA_CROSSCHAIN_ID] = zilliqaSdk
	}
	{
		fantomConfig := config.GetChainListenConfig(basedef.FANTOM_CROSSCHAIN_ID)
		if fantomConfig == nil {
			panic("chain is invalid")
		}
		urls := fantomConfig.GetNodesUrl()
		fantomSdk = chainsdk.NewEthereumSdkPro(urls, fantomConfig.ListenSlot, fantomConfig.ChainId)
		sdkMap[basedef.FANTOM_CROSSCHAIN_ID] = fantomSdk
	}
	{
		avaxConfig := config.GetChainListenConfig(basedef.AVAX_CROSSCHAIN_ID)
		if avaxConfig == nil {
			panic("chain is invalid")
		}
		urls := avaxConfig.GetNodesUrl()
		avaxSdk = chainsdk.NewEthereumSdkPro(urls, avaxConfig.ListenSlot, avaxConfig.ChainId)
		sdkMap[basedef.AVAX_CROSSCHAIN_ID] = avaxSdk
	}
	{
		optimisticConfig := config.GetChainListenConfig(basedef.OPTIMISTIC_CROSSCHAIN_ID)
		if optimisticConfig == nil {
			panic("chain is invalid")
		}
		urls := optimisticConfig.GetNodesUrl()
		optimisticSdk = chainsdk.NewEthereumSdkPro(urls, optimisticConfig.ListenSlot, optimisticConfig.ChainId)
		sdkMap[basedef.OPTIMISTIC_CROSSCHAIN_ID] = optimisticSdk
	}
	{
		metisConfig := config.GetChainListenConfig(basedef.METIS_CROSSCHAIN_ID)
		if metisConfig == nil {
			panic("metis chain is invalid")
		}
		urls := metisConfig.GetNodesUrl()
		metisSdk = chainsdk.NewEthereumSdkPro(urls, metisConfig.ListenSlot, metisConfig.ChainId)
		sdkMap[basedef.METIS_CROSSCHAIN_ID] = metisSdk
	}
	{
		pixieConfig := config.GetChainListenConfig(basedef.PIXIE_CROSSCHAIN_ID)
		if pixieConfig == nil {
			panic("pixie chain is invalid")
		}
		urls := pixieConfig.GetNodesUrl()
		pixieSdk = chainsdk.NewEthereumSdkPro(urls, pixieConfig.ListenSlot, pixieConfig.ChainId)
		sdkMap[basedef.PIXIE_CROSSCHAIN_ID] = pixieSdk
	}
	{
		rinkebyConfig := config.GetChainListenConfig(basedef.RINKEBY_CROSSCHAIN_ID)
		if rinkebyConfig == nil {
			panic("metis chain is invalid")
		}
		urls := rinkebyConfig.GetNodesUrl()
		rinkebySdk = chainsdk.NewEthereumSdkPro(urls, rinkebyConfig.ListenSlot, rinkebyConfig.ChainId)
		sdkMap[basedef.RINKEBY_CROSSCHAIN_ID] = rinkebySdk
	}
}

func GetBalance(chainId uint64, hash string) (*big.Int, error) {
	maxBalance := big.NewInt(0)
	maxFun := func(balance *big.Int) {
		if balance.Cmp(maxBalance) > 0 {
			maxBalance = balance
		}
	}
	errMap := make(map[error]bool, 0)
	if chainId == basedef.ETHEREUM_CROSSCHAIN_ID {
		ethereumConfig := config.GetChainListenConfig(basedef.ETHEREUM_CROSSCHAIN_ID)
		if ethereumConfig == nil {
			panic("chain is invalid")
		}
		for _, v := range ethereumConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := ethereumSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.MATIC_CROSSCHAIN_ID {
		maticConfig := config.GetChainListenConfig(basedef.MATIC_CROSSCHAIN_ID)
		if maticConfig == nil {
			panic("chain is invalid")
		}
		for _, v := range maticConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := maticSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.BSC_CROSSCHAIN_ID {
		bscConfig := config.GetChainListenConfig(basedef.BSC_CROSSCHAIN_ID)
		if bscConfig == nil {
			panic("chain is invalid")
		}
		for _, v := range bscConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := bscSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.HECO_CROSSCHAIN_ID {
		hecoConfig := config.GetChainListenConfig(basedef.HECO_CROSSCHAIN_ID)
		if hecoConfig == nil {
			panic("chain is invalid")
		}
		for _, v := range hecoConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := hecoSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.OK_CROSSCHAIN_ID {
		okConfig := config.GetChainListenConfig(basedef.OK_CROSSCHAIN_ID)
		if okConfig == nil {
			panic("chain is invalid")
		}
		for _, v := range okConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := okSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.NEO_CROSSCHAIN_ID {
		neoConfig := config.GetChainListenConfig(basedef.NEO_CROSSCHAIN_ID)
		if neoConfig == nil {
			panic("chain is invalid")
		}
		for _, v := range neoConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := neoSdk.Nep5Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.NEO3_CROSSCHAIN_ID {
		neo3Config := config.GetChainListenConfig(basedef.NEO3_CROSSCHAIN_ID)
		if neo3Config == nil {
			panic("chain is invalid")
		}
		for _, v := range neo3Config.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := neo3Sdk.Nep17Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.ONT_CROSSCHAIN_ID {
		ontConfig := config.GetChainListenConfig(basedef.ONT_CROSSCHAIN_ID)
		if ontConfig == nil {
			panic("chain is invalid")
		}
		for _, v := range ontConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := ontologySdk.Oep4Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.ARBITRUM_CROSSCHAIN_ID {
		arbitrumConfig := config.GetChainListenConfig(basedef.ARBITRUM_CROSSCHAIN_ID)
		if arbitrumConfig == nil {
			panic("chain is invalid")
		}
		for _, v := range arbitrumConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := arbitrumSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.XDAI_CROSSCHAIN_ID {
		xdaiConfig := config.GetChainListenConfig(basedef.XDAI_CROSSCHAIN_ID)
		if xdaiConfig == nil {
			panic("chain is invalid")
		}
		for _, v := range xdaiConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := xdaiSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.ZILLIQA_CROSSCHAIN_ID {
		zilliqaCfg := config.GetChainListenConfig(basedef.ZILLIQA_CROSSCHAIN_ID)
		if zilliqaCfg == nil {
			panic("zilliqa GetChainListenConfig chain is invalid")
		}
		for _, v := range zilliqaCfg.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := zilliqaSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.FANTOM_CROSSCHAIN_ID {
		fantomConfig := config.GetChainListenConfig(basedef.FANTOM_CROSSCHAIN_ID)
		if fantomConfig == nil {
			panic("chain is invalid")
		}
		for _, v := range fantomConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := fantomSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.AVAX_CROSSCHAIN_ID {
		avaxConfig := config.GetChainListenConfig(basedef.AVAX_CROSSCHAIN_ID)
		if avaxConfig == nil {
			panic("chain is invalid")
		}
		for _, v := range avaxConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := avaxSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.OPTIMISTIC_CROSSCHAIN_ID {
		optimisticConfig := config.GetChainListenConfig(basedef.OPTIMISTIC_CROSSCHAIN_ID)
		if optimisticConfig == nil {
			panic("chain is invalid")
		}
		for _, v := range optimisticConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := optimisticSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.METIS_CROSSCHAIN_ID {
		metisConfig := config.GetChainListenConfig(basedef.METIS_CROSSCHAIN_ID)
		if metisConfig == nil {
			panic("metis chain is invalid")
		}
		for _, v := range metisConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := metisSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.PIXIE_CROSSCHAIN_ID {
		pixieConfig := config.GetChainListenConfig(basedef.PIXIE_CROSSCHAIN_ID)
		if pixieConfig == nil {
			panic("pixie chain is invalid")
		}
		for _, v := range pixieConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := pixieSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.RINKEBY_CROSSCHAIN_ID {
		rinkebyConfig := config.GetChainListenConfig(basedef.RINKEBY_CROSSCHAIN_ID)
		if rinkebyConfig == nil {
			panic("rinkeby chain is invalid")
		}
		for _, v := range rinkebyConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := rinkebySdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
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
		for _, v := range ontConfig.ProxyContract {
			if len(strings.TrimSpace(v)) != 0 {
				return ontologySdk.Oep4TotalSupply(hash, v)
			}
		}

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
	if chainId == basedef.METIS_CROSSCHAIN_ID {
		metisConfig := config.GetChainListenConfig(basedef.METIS_CROSSCHAIN_ID)
		if metisConfig == nil {
			panic("metis chain GetTotalSupply invalid")
		}
		return metisSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.PIXIE_CROSSCHAIN_ID {
		pixieConfig := config.GetChainListenConfig(basedef.PIXIE_CROSSCHAIN_ID)
		if pixieConfig == nil {
			panic("pixie chain GetTotalSupply invalid")
		}
		return pixieSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.RINKEBY_CROSSCHAIN_ID {
		rinkebyConfig := config.GetChainListenConfig(basedef.RINKEBY_CROSSCHAIN_ID)
		if rinkebyConfig == nil {
			panic("rinkeby chain GetTotalSupply invalid")
		}
		return rinkebySdk.Erc20TotalSupply(hash)
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
		return ethereumSdk.Erc20Balance(hash, proxy)
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
	case basedef.METIS_CROSSCHAIN_ID:
		return metisSdk.Erc20Balance(hash, proxy)
	case basedef.PIXIE_CROSSCHAIN_ID:
		return pixieSdk.Erc20Balance(hash, proxy)
	case basedef.RINKEBY_CROSSCHAIN_ID:
		return rinkebySdk.Erc20Balance(hash, proxy)
	default:
		return new(big.Int).SetUint64(0), nil
	}
}

func GetBoundLockProxy(lockProxies []string, srcTokenHash, DstTokenHash string, srcChainId, dstChainId uint64) (string, error) {
	if sdk, exist := sdkMap[dstChainId]; exist {
		if value, ok := sdk.(*chainsdk.EthereumSdkPro); ok {
			return value.GetBoundLockProxy(lockProxies, srcTokenHash, DstTokenHash, srcChainId)
		}
	}
	return "", fmt.Errorf("chain %d is not ethereum based", dstChainId)
}
