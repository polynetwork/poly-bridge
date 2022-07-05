package common

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
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
	bobaSdk       *chainsdk.EthereumSdkPro
	rinkebySdk    *chainsdk.EthereumSdkPro
	bytomSdk      *chainsdk.EthereumSdkPro
	oasisSdk      *chainsdk.EthereumSdkPro
	harmonySdk    *chainsdk.EthereumSdkPro
	kccSdk        *chainsdk.EthereumSdkPro
	hscSdk        *chainsdk.EthereumSdkPro
	starcoinSdk   *chainsdk.StarcoinSdkPro
	kavaSdk       *chainsdk.EthereumSdkPro
	cubeSdk       *chainsdk.EthereumSdkPro
	zkSyncSdk     *chainsdk.EthereumSdkPro
	celoSdk       *chainsdk.EthereumSdkPro
	cloverSdk     *chainsdk.EthereumSdkPro
	sdkMap        map[uint64]interface{}
	config        *conf.Config
)

func GetSdk(chainId uint64) interface{} {
	return sdkMap[chainId]
}

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
		bobaConfig := config.GetChainListenConfig(basedef.BOBA_CROSSCHAIN_ID)
		if bobaConfig == nil {
			panic("boba chain is invalid")
		}
		urls := bobaConfig.GetNodesUrl()
		bobaSdk = chainsdk.NewEthereumSdkPro(urls, bobaConfig.ListenSlot, bobaConfig.ChainId)
		sdkMap[basedef.BOBA_CROSSCHAIN_ID] = bobaSdk
	}
	{
		starcoinConfig := config.GetChainListenConfig(basedef.STARCOIN_CROSSCHAIN_ID)
		if starcoinConfig == nil {
			panic("starcoin chain is invalid")
		}
		urls := starcoinConfig.GetNodesUrl()
		starcoinSdk = chainsdk.NewStarcoinSdkPro(urls, starcoinConfig.ListenSlot, starcoinConfig.ChainId)
		sdkMap[basedef.STARCOIN_CROSSCHAIN_ID] = starcoinSdk
	}
	if basedef.ENV == basedef.TESTNET {
		{
			rinkebyConfig := config.GetChainListenConfig(basedef.RINKEBY_CROSSCHAIN_ID)
			if rinkebyConfig == nil {
				panic("rinkeby chain is invalid")
			}
			urls := rinkebyConfig.GetNodesUrl()
			rinkebySdk = chainsdk.NewEthereumSdkPro(urls, rinkebyConfig.ListenSlot, rinkebyConfig.ChainId)
			sdkMap[basedef.RINKEBY_CROSSCHAIN_ID] = rinkebySdk
		}
	}
	{
		chainConfig := config.GetChainListenConfig(basedef.OASIS_CROSSCHAIN_ID)
		if chainConfig == nil {
			panic("oasis chain is invalid")
		}
		urls := chainConfig.GetNodesUrl()
		oasisSdk = chainsdk.NewEthereumSdkPro(urls, chainConfig.ListenSlot, chainConfig.ChainId)
		sdkMap[basedef.OASIS_CROSSCHAIN_ID] = oasisSdk
	}
	{
		chainConfig := config.GetChainListenConfig(basedef.HARMONY_CROSSCHAIN_ID)
		if chainConfig == nil {
			panic("harmony chain is invalid")
		}
		urls := chainConfig.GetNodesUrl()
		harmonySdk = chainsdk.NewEthereumSdkPro(urls, chainConfig.ListenSlot, chainConfig.ChainId)
		sdkMap[basedef.HARMONY_CROSSCHAIN_ID] = harmonySdk
	}
	{
		chainConfig := config.GetChainListenConfig(basedef.KCC_CROSSCHAIN_ID)
		if chainConfig == nil {
			panic("kcc chain is invalid")
		}
		urls := chainConfig.GetNodesUrl()
		kccSdk = chainsdk.NewEthereumSdkPro(urls, chainConfig.ListenSlot, chainConfig.ChainId)
		sdkMap[basedef.KCC_CROSSCHAIN_ID] = kccSdk
	}
	{
		bytomConfig := config.GetChainListenConfig(basedef.BYTOM_CROSSCHAIN_ID)
		if bytomConfig == nil {
			panic("bytom chain is invalid")
		}
		urls := bytomConfig.GetNodesUrl()
		bytomSdk = chainsdk.NewEthereumSdkPro(urls, bytomConfig.ListenSlot, bytomConfig.ChainId)
		sdkMap[basedef.BYTOM_CROSSCHAIN_ID] = bytomSdk
	}
	{
		chainConfig := config.GetChainListenConfig(basedef.HSC_CROSSCHAIN_ID)
		if chainConfig == nil {
			panic("chain HSC is invalid")
		}
		urls := chainConfig.GetNodesUrl()
		hscSdk = chainsdk.NewEthereumSdkPro(urls, chainConfig.ListenSlot, chainConfig.ChainId)
		sdkMap[basedef.HSC_CROSSCHAIN_ID] = hscSdk
	}
	{
		kavaConfig := config.GetChainListenConfig(basedef.KAVA_CROSSCHAIN_ID)
		if kavaConfig == nil {
			panic("kava chain is invalid")
		}
		urls := kavaConfig.GetNodesUrl()
		kavaSdk = chainsdk.NewEthereumSdkPro(urls, kavaConfig.ListenSlot, kavaConfig.ChainId)
		sdkMap[basedef.KAVA_CROSSCHAIN_ID] = kavaSdk
	}
	{
		cubeConfig := config.GetChainListenConfig(basedef.CUBE_CROSSCHAIN_ID)
		if cubeConfig == nil {
			panic("cube chain is invalid")
		}
		urls := cubeConfig.GetNodesUrl()
		cubeSdk = chainsdk.NewEthereumSdkPro(urls, cubeConfig.ListenSlot, cubeConfig.ChainId)
		sdkMap[basedef.CUBE_CROSSCHAIN_ID] = cubeSdk
	}
	{
		chainConfig := config.GetChainListenConfig(basedef.ZKSYNC_CROSSCHAIN_ID)
		if chainConfig == nil {
			panic("zkSync chain is invalid")
		}
		urls := chainConfig.GetNodesUrl()
		zkSyncSdk = chainsdk.NewEthereumSdkPro(urls, chainConfig.ListenSlot, chainConfig.ChainId)
		sdkMap[basedef.ZKSYNC_CROSSCHAIN_ID] = zkSyncSdk
	}
	{
		chainConfig := config.GetChainListenConfig(basedef.CELO_CROSSCHAIN_ID)
		if chainConfig == nil {
			panic("celo chain is invalid")
		}
		urls := chainConfig.GetNodesUrl()
		celoSdk = chainsdk.NewEthereumSdkPro(urls, chainConfig.ListenSlot, chainConfig.ChainId)
		sdkMap[basedef.CELO_CROSSCHAIN_ID] = celoSdk
	}
	{
		chainConfig := config.GetChainListenConfig(basedef.CLOVER_CROSSCHAIN_ID)
		if chainConfig == nil {
			panic("clover chain is invalid")
		}
		urls := chainConfig.GetNodesUrl()
		cloverSdk = chainsdk.NewEthereumSdkPro(urls, chainConfig.ListenSlot, chainConfig.ChainId)
		sdkMap[basedef.CLOVER_CROSSCHAIN_ID] = cloverSdk
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
	if chainId == basedef.BOBA_CROSSCHAIN_ID {
		bobaConfig := config.GetChainListenConfig(basedef.BOBA_CROSSCHAIN_ID)
		if bobaConfig == nil {
			panic("boba chain is invalid")
		}
		for _, v := range bobaConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := bobaSdk.Erc20Balance(hash, v)
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
	if chainId == basedef.OASIS_CROSSCHAIN_ID {
		chainConfig := config.GetChainListenConfig(basedef.OASIS_CROSSCHAIN_ID)
		if chainConfig == nil {
			panic("oasis chain is invalid")
		}
		for _, v := range chainConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := oasisSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.HARMONY_CROSSCHAIN_ID {
		chainConfig := config.GetChainListenConfig(basedef.HARMONY_CROSSCHAIN_ID)
		if chainConfig == nil {
			panic("harmony chain is invalid")
		}
		for _, v := range chainConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := harmonySdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.KCC_CROSSCHAIN_ID {
		chainConfig := config.GetChainListenConfig(basedef.KCC_CROSSCHAIN_ID)
		if chainConfig == nil {
			panic("kcc chain is invalid")
		}
		for _, v := range chainConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := kccSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.BYTOM_CROSSCHAIN_ID {
		bytomConfig := config.GetChainListenConfig(basedef.BYTOM_CROSSCHAIN_ID)
		if bytomConfig == nil {
			panic("bytom chain is invalid")
		}
		for _, v := range bytomConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := bytomSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.HSC_CROSSCHAIN_ID {
		chainConfig := config.GetChainListenConfig(basedef.HSC_CROSSCHAIN_ID)
		if chainConfig == nil {
			panic("hsc chain is invalid")
		}
		for _, v := range chainConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := hscSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.STARCOIN_CROSSCHAIN_ID {
		starcoinConfig := config.GetChainListenConfig(basedef.STARCOIN_CROSSCHAIN_ID)
		if starcoinConfig == nil {
			panic("starcoin chain is invalid")
		}
		for _, v := range starcoinConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := starcoinSdk.GetBalance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.KAVA_CROSSCHAIN_ID {
		kavaConfig := config.GetChainListenConfig(basedef.KAVA_CROSSCHAIN_ID)
		if kavaConfig == nil {
			panic("kava chain is invalid")
		}
		for _, v := range kavaConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := kavaSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.CUBE_CROSSCHAIN_ID {
		cubeConfig := config.GetChainListenConfig(basedef.CUBE_CROSSCHAIN_ID)
		if cubeConfig == nil {
			panic("cube chain is invalid")
		}
		for _, v := range cubeConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := cubeSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.ZKSYNC_CROSSCHAIN_ID {
		chainConfig := config.GetChainListenConfig(basedef.ZKSYNC_CROSSCHAIN_ID)
		if chainConfig == nil {
			panic("zkSync chain is invalid")
		}
		for _, v := range chainConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := zkSyncSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.CELO_CROSSCHAIN_ID {
		chainConfig := config.GetChainListenConfig(basedef.CELO_CROSSCHAIN_ID)
		if chainConfig == nil {
			panic("celo chain is invalid")
		}
		for _, v := range chainConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := celoSdk.Erc20Balance(hash, v)
			maxFun(balance)
			errMap[err] = true
		}
	}
	if chainId == basedef.CLOVER_CROSSCHAIN_ID {
		chainConfig := config.GetChainListenConfig(basedef.CLOVER_CROSSCHAIN_ID)
		if chainConfig == nil {
			panic("clover chain is invalid")
		}
		for _, v := range chainConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := cloverSdk.Erc20Balance(hash, v)
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
	if chainId == basedef.BOBA_CROSSCHAIN_ID {
		bobaConfig := config.GetChainListenConfig(basedef.BOBA_CROSSCHAIN_ID)
		if bobaConfig == nil {
			panic("boba chain GetTotalSupply invalid")
		}
		return bobaSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.RINKEBY_CROSSCHAIN_ID {
		rinkebyConfig := config.GetChainListenConfig(basedef.RINKEBY_CROSSCHAIN_ID)
		if rinkebyConfig == nil {
			panic("rinkeby chain GetTotalSupply invalid")
		}
		return rinkebySdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.OASIS_CROSSCHAIN_ID {
		chainConfig := config.GetChainListenConfig(basedef.OASIS_CROSSCHAIN_ID)
		if chainConfig == nil {
			panic("oasis chain GetTotalSupply invalid")
		}
		return oasisSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.KCC_CROSSCHAIN_ID {
		chainConfig := config.GetChainListenConfig(basedef.KCC_CROSSCHAIN_ID)
		if chainConfig == nil {
			panic("kcc chain GetTotalSupply invalid")
		}
		return kccSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.BYTOM_CROSSCHAIN_ID {
		bytomConfig := config.GetChainListenConfig(basedef.BYTOM_CROSSCHAIN_ID)
		if bytomConfig == nil {
			panic("bytom chain GetTotalSupply invalid")
		}
		return bytomSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.HSC_CROSSCHAIN_ID {
		chainConfig := config.GetChainListenConfig(basedef.HSC_CROSSCHAIN_ID)
		if chainConfig == nil {
			panic("hsc chain GetTotalSupply invalid")
		}
		return hscSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.KAVA_CROSSCHAIN_ID {
		kavaConfig := config.GetChainListenConfig(basedef.KAVA_CROSSCHAIN_ID)
		if kavaConfig == nil {
			panic("kava chain GetTotalSupply invalid")
		}
		return kavaSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.CUBE_CROSSCHAIN_ID {
		cubeConfig := config.GetChainListenConfig(basedef.CUBE_CROSSCHAIN_ID)
		if cubeConfig == nil {
			panic("cube chain GetTotalSupply invalid")
		}
		return cubeSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.ZKSYNC_CROSSCHAIN_ID {
		chainConfig := config.GetChainListenConfig(basedef.ZKSYNC_CROSSCHAIN_ID)
		if chainConfig == nil {
			panic("zkSync chain GetTotalSupply invalid")
		}
		return zkSyncSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.CELO_CROSSCHAIN_ID {
		chainConfig := config.GetChainListenConfig(basedef.CELO_CROSSCHAIN_ID)
		if chainConfig == nil {
			panic("celo chain GetTotalSupply invalid")
		}
		return celoSdk.Erc20TotalSupply(hash)
	}
	if chainId == basedef.CLOVER_CROSSCHAIN_ID {
		chainConfig := config.GetChainListenConfig(basedef.CLOVER_CROSSCHAIN_ID)
		if chainConfig == nil {
			panic("clover chain GetTotalSupply invalid")
		}
		return cloverSdk.Erc20TotalSupply(hash)
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
	case basedef.BOBA_CROSSCHAIN_ID:
		return bobaSdk.Erc20Balance(hash, proxy)
	case basedef.RINKEBY_CROSSCHAIN_ID:
		return rinkebySdk.Erc20Balance(hash, proxy)
	case basedef.OASIS_CROSSCHAIN_ID:
		return oasisSdk.Erc20Balance(hash, proxy)
	case basedef.HARMONY_CROSSCHAIN_ID:
		return harmonySdk.Erc20Balance(hash, proxy)
	case basedef.KCC_CROSSCHAIN_ID:
		return kccSdk.Erc20Balance(hash, proxy)
	case basedef.BYTOM_CROSSCHAIN_ID:
		return bytomSdk.Erc20Balance(hash, proxy)
	case basedef.HSC_CROSSCHAIN_ID:
		return hscSdk.Erc20Balance(hash, proxy)
	case basedef.STARCOIN_CROSSCHAIN_ID:
		return starcoinSdk.GetBalance(hash, proxy)
	case basedef.KAVA_CROSSCHAIN_ID:
		return kavaSdk.Erc20Balance(hash, proxy)
	case basedef.CUBE_CROSSCHAIN_ID:
		return cubeSdk.Erc20Balance(hash, proxy)
	case basedef.ZKSYNC_CROSSCHAIN_ID:
		return zkSyncSdk.Erc20Balance(hash, proxy)
	case basedef.CELO_CROSSCHAIN_ID:
		return celoSdk.Erc20Balance(hash, proxy)
	case basedef.CLOVER_CROSSCHAIN_ID:
		return cloverSdk.Erc20Balance(hash, proxy)
	default:
		return new(big.Int).SetUint64(0), nil
	}
}

func GetNftOwner(chainId uint64, asset string, tokenId int) (owner common.Address, err error) {
	switch chainId {
	case basedef.ETHEREUM_CROSSCHAIN_ID:
		return ethereumSdk.GetNFTOwner(asset, big.NewInt(int64(tokenId)))
	case basedef.MATIC_CROSSCHAIN_ID:
		return maticSdk.GetNFTOwner(asset, big.NewInt(int64(tokenId)))
	case basedef.BSC_CROSSCHAIN_ID:
		return bscSdk.GetNFTOwner(asset, big.NewInt(int64(tokenId)))
	case basedef.HECO_CROSSCHAIN_ID:
		return hecoSdk.GetNFTOwner(asset, big.NewInt(int64(tokenId)))
	case basedef.OK_CROSSCHAIN_ID:
		return okSdk.GetNFTOwner(asset, big.NewInt(int64(tokenId)))
	case basedef.ARBITRUM_CROSSCHAIN_ID:
		return arbitrumSdk.GetNFTOwner(asset, big.NewInt(int64(tokenId)))
	case basedef.XDAI_CROSSCHAIN_ID:
		return xdaiSdk.GetNFTOwner(asset, big.NewInt(int64(tokenId)))
	default:
		return common.Address{}, fmt.Errorf("has nat func with chain:%v", chainId)
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
