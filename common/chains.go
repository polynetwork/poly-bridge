package common

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/polynetwork/bridge-common/chains/eth"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	erc20 "poly-bridge/go_abi/mintable_erc20_abi"
	nftmapping "poly-bridge/go_abi/nft_mapping_abi"
	"strings"
	"time"
)

var (
	sdkMap map[uint64]interface{}
)

func GetSdk(chainId uint64) (interface{}, error) {
	sdk, ok := sdkMap[chainId]
	if !ok {
		return nil, fmt.Errorf("chain %d sdk not initialized", chainId)
	}
	return sdk, nil
}

func SetupChainsSDK(cfg *conf.Config) {
	if cfg == nil {
		panic("Missing config")
	}
	sdkMap = make(map[uint64]interface{}, 0)
	for _, chainListenConfig := range cfg.ChainListenConfig {
		switch chainListenConfig.ChainId {
		case basedef.ONT_CROSSCHAIN_ID:
		case basedef.NEO_CROSSCHAIN_ID:
		case basedef.NEO3_CROSSCHAIN_ID:
		case basedef.RIPPLE_CROSSCHAIN_ID:
		case basedef.STARCOIN_CROSSCHAIN_ID:
		case basedef.APTOS_CROSSCHAIN_ID:
		case basedef.ZKSYNC_CROSSCHAIN_ID:
		default:
			if basedef.IsETHChain(chainListenConfig.ChainId) {
				sdk, err := eth.WithOptions(chainListenConfig.ChainId, chainListenConfig.Nodes, time.Second*30, 1)
				if err != nil {
					panic(fmt.Sprintf("Create chain sdk failed. chain=%d, err=%s", chainListenConfig.ChainId, err))
				}
				sdkMap[chainListenConfig.ChainId] = sdk
			}
		}
	}
}

func GetBalance(chainId uint64, hash string) (balance *big.Int, err error) {
	sdk, err := GetSdk(chainId)
	if err != nil {
		return nil, err
	}
	maxBalance, balance := big.NewInt(0), big.NewInt(0)
	maxFun := func(balance *big.Int) {
		if balance.Cmp(maxBalance) > 0 {
			maxBalance = balance
		}
	}

	errMap := make(map[error]interface{}, 0)
	switch chainId {
	case basedef.ONT_CROSSCHAIN_ID:
	case basedef.NEO_CROSSCHAIN_ID:
	case basedef.NEO3_CROSSCHAIN_ID:
	case basedef.RIPPLE_CROSSCHAIN_ID:
	case basedef.STARCOIN_CROSSCHAIN_ID:
	case basedef.APTOS_CROSSCHAIN_ID:
	case basedef.ZKSYNC_CROSSCHAIN_ID:

	default:
		if basedef.IsETHChain(chainId) {
			ethConfig := conf.GlobalConfig.GetChainListenConfig(chainId)
			if ethConfig == nil {
				err = fmt.Errorf("chain %d is invalid", chainId)
				return
			}
			for _, v := range ethConfig.ProxyContract {
				if len(strings.TrimSpace(v)) == 0 {
					continue
				}
				balance, err := chainsdk.GetEthErc20Balance(hash, v, sdk.(*eth.SDK).Node())
				if err != nil {
					return nil, err
				}
				errMap[err] = nil
				maxFun(balance)
			}
			if maxBalance.Cmp(big.NewInt(0)) > 0 {
				return maxBalance, nil
			}
			for k := range errMap {
				if k == nil {
					return new(big.Int).SetUint64(0), nil
				} else {
					err = k
				}
			}
			return new(big.Int).SetUint64(0), err
		}

	}
	return new(big.Int).SetUint64(0), fmt.Errorf(fmt.Sprintf("chain id %d invalid", chainId))

}

func GetProxyBalance(chainId uint64, hash string, proxy string) (balance *big.Int, err error) {
	sdk, err := GetSdk(chainId)
	if err != nil {
		return nil, err
	}
	balance = big.NewInt(0)
	switch chainId {
	case basedef.ONT_CROSSCHAIN_ID:
	case basedef.NEO_CROSSCHAIN_ID:
	case basedef.NEO3_CROSSCHAIN_ID:
	case basedef.RIPPLE_CROSSCHAIN_ID:
	case basedef.STARCOIN_CROSSCHAIN_ID:
	case basedef.APTOS_CROSSCHAIN_ID:
	case basedef.ZKSYNC_CROSSCHAIN_ID:

	default:
		if basedef.IsETHChain(chainId) {
			balance, err = chainsdk.GetEthErc20Balance(hash, proxy, sdk.(*eth.SDK).Node())
			if err != nil {
				return big.NewInt(0), err
			}
			return
		}
	}
	return new(big.Int).SetUint64(0), fmt.Errorf(fmt.Sprintf("chain id %d invalid", chainId))

}

func GetTotalSupply(chainId uint64, hash string) (totalSupply *big.Int, err error) {
	sdk, err := GetSdk(chainId)
	if err != nil {
		return nil, err
	}
	totalSupply = big.NewInt(0)
	switch chainId {
	case basedef.ONT_CROSSCHAIN_ID:
	case basedef.NEO_CROSSCHAIN_ID:
	case basedef.NEO3_CROSSCHAIN_ID:
	case basedef.RIPPLE_CROSSCHAIN_ID:
	case basedef.STARCOIN_CROSSCHAIN_ID:
	case basedef.APTOS_CROSSCHAIN_ID:
	case basedef.ZKSYNC_CROSSCHAIN_ID:
	default:
		if basedef.IsETHChain(chainId) {
			if !basedef.IsNativeTokenAddress(hash) {
				erc20Address := common.HexToAddress(hash)
				contract, err := erc20.NewERC20Mintable(erc20Address, sdk.(*eth.SDK).Node().Client)
				if err != nil {
					totalSupply, err = contract.TotalSupply(nil)
				}
			}
			if err != nil {
				return big.NewInt(0), err
			}
		}
	}
	err = fmt.Errorf(fmt.Sprintf("chain id %d invalid", chainId))
	return
}

func GetNftOwner(chainId uint64, asset string, tokenId int) (owner common.Address, err error) {
	owner = common.Address{}
	sdk, err := GetSdk(chainId)
	if err != nil {
		return owner, err
	}
	assetAddr := common.HexToAddress(asset)
	switch chainId {
	case basedef.ONT_CROSSCHAIN_ID:
	case basedef.NEO_CROSSCHAIN_ID:
	case basedef.NEO3_CROSSCHAIN_ID:
	case basedef.RIPPLE_CROSSCHAIN_ID:
	case basedef.STARCOIN_CROSSCHAIN_ID:
	case basedef.APTOS_CROSSCHAIN_ID:
	case basedef.ZKSYNC_CROSSCHAIN_ID:

	default:
		if basedef.IsETHChain(chainId) {
			var contract *nftmapping.CrossChainNFTMapping
			contract, err = nftmapping.NewCrossChainNFTMapping(assetAddr, sdk.(*eth.SDK).Node())
			if err == nil {
				owner, err = contract.OwnerOf(nil, big.NewInt(int64(tokenId)))
			}
		}
		err = fmt.Errorf(fmt.Sprintf("chain id %d invalid", chainId))
	}
	return
}

func GetBoundLockProxy(lockProxies []string, srcTokenHash, DstTokenHash string, srcChainId, dstChainId uint64) (string, error) {
	sdk, err := GetSdk(dstChainId)
	if err != nil {
		return "", err
	}
	switch dstChainId {
	case basedef.ONT_CROSSCHAIN_ID:
	case basedef.NEO_CROSSCHAIN_ID:
	case basedef.NEO3_CROSSCHAIN_ID:
	case basedef.RIPPLE_CROSSCHAIN_ID:
	case basedef.STARCOIN_CROSSCHAIN_ID:
	case basedef.APTOS_CROSSCHAIN_ID:
	case basedef.ZKSYNC_CROSSCHAIN_ID:
	default:
		if basedef.IsETHChain(dstChainId) {
			return chainsdk.GetBoundLockProxy(lockProxies, srcTokenHash, DstTokenHash, srcChainId, sdk.(*eth.SDK))
		}
	}
	return "", fmt.Errorf("chain %d is not ethereum based", dstChainId)

}
