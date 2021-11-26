package chainsdk

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"testing"
)

func TestEthereumSdk_GetBoundAssetHash(t *testing.T) {
	config := conf.NewConfig("../prod.json")
	srcChainId := basedef.BSC_CROSSCHAIN_ID
	dstChainId := basedef.METIS_CROSSCHAIN_ID
	//dstTokenHash := "2692be44a6e38b698731fddf417d060f0d20a0cb"
	dstTokenHash := "DeadDeAddeAddEAddeadDEaDDEAdDeaDDeAD0000"
	dstTokenAddress := common.HexToAddress(dstTokenHash)
	var lockProxies []string
	for _, listenConfig := range config.ChainListenConfig {
		if listenConfig.ChainId == dstChainId {
			lockProxies = listenConfig.ProxyContract
			break
		}
	}
	for _, chainNodeConfig := range config.ChainNodes {
		if chainNodeConfig.ChainId == dstChainId {
			for _, node := range chainNodeConfig.Nodes {
				sdk, err := NewEthereumSdk(node.Url)
				if err != nil || sdk == nil || sdk.GetClient() == nil {
					logs.Info("node: %s,NewEthereumSdk error: %s", node.Url, err)
					continue
				}
				for _, proxy := range lockProxies {
					proxyAddr := common.HexToAddress(proxy)
					addr, err := sdk.GetBoundAssetHash(dstTokenAddress, proxyAddr, srcChainId)
					if err != nil {
						logs.Error("GetBoundAssetHash err:%s", err)
						continue
					}
					if addr == nil {
						continue
					}
					addrHash := addr.Hex()
					logs.Info("GetBoundAssetHash addrHash=%s", addrHash)
					assert.Equal(t, "0xe552Fb52a4F19e44ef5A967632DBc320B0820639", addrHash, "GetBoundAssetHash error")
				}
			}
		}
	}
}
