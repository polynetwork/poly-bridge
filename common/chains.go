package common

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/polynetwork/bridge-common/base"
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

	ethSdkMap map[uint64]*eth.SDK
)

func GetSdk(chainId uint64) interface{} {
	if base.SameAsETH(chainId) {
		return ethSdkMap[chainId]
	}
	return sdkMap[chainId]
}

func SetupChainsSDK(cfg *conf.Config) {
	if cfg == nil {
		panic("Missing config")
	}
	ethSdkMap = make(map[uint64]*eth.SDK, 0)
	for _, chainListenConfig := range cfg.ChainListenConfig {
		switch chainListenConfig.ChainId {
		case base.HECO, base.BSC, base.MATIC:
			sdk, err := eth.WithOptions(chainListenConfig.ChainId, chainListenConfig.Nodes, time.Second*30, 1)
			if err != nil {
				panic(fmt.Sprintf("Create chain sdk failed. chain=%d, err=%s", chainListenConfig.ChainId, err))
			}
			ethSdkMap[chainListenConfig.ChainId] = sdk
		}
	}
}

func getEthErc20Balance(token, owner string, client *eth.Client) (balance *big.Int, err error) {
	tokenAddress := common.HexToAddress(token)
	ownerAddr := common.HexToAddress(owner)
	if basedef.IsNativeTokenAddress(token) {
		var result hexutil.Big
		ctx := context.Background()
		err = client.Rpc.CallContext(ctx, &result, "eth_getBalance", "0x"+owner, "latest")
		balance = (*big.Int)(&result)
	} else {
		var contract *erc20.ERC20Extended
		contract, err = erc20.NewERC20Mintable(tokenAddress, client)
		if err == nil {
			balance, err = contract.BalanceOf(nil, ownerAddr)
		}
	}
	return
}

func GetBalance(chainId uint64, hash string) (balance *big.Int, err error) {
	maxBalance, balance := big.NewInt(0), big.NewInt(0)
	maxFun := func(balance *big.Int) {
		if balance.Cmp(maxBalance) > 0 {
			maxBalance = balance
		}
	}

	errMap := make(map[error]interface{}, 0)
	switch chainId {
	case base.HECO, base.BSC, base.MATIC:
		ethConfig := conf.GlobalConfig.GetChainListenConfig(chainId)
		if ethConfig == nil {
			err = fmt.Errorf("chain %d is invalid", chainId)
			return
		}
		sdk := ethSdkMap[chainId]
		if sdk == nil {
			err = fmt.Errorf("chain %d sdk not initialized", chainId)
			return
		}
		for _, v := range ethConfig.ProxyContract {
			if len(strings.TrimSpace(v)) == 0 {
				continue
			}
			balance, err := getEthErc20Balance(hash, v, sdk.Node())
			if err != nil {
				return nil, err
			}
			errMap[err] = nil
			maxFun(balance)
		}
		if maxBalance.Cmp(big.NewInt(0)) > 0 {
			return maxBalance, nil
		}
		for k, _ := range errMap {
			if k == nil {
				return new(big.Int).SetUint64(0), nil
			} else {
				err = k
			}
		}
		return new(big.Int).SetUint64(0), err
	default:
		return new(big.Int).SetUint64(0), fmt.Errorf(fmt.Sprintf("chain id %d invalid", chainId))
	}
}

func GetProxyBalance(chainId uint64, hash string, proxy string) (balance *big.Int, err error) {
	balance = big.NewInt(0)
	switch chainId {
	case base.HECO, base.BSC, base.MATIC:
		sdk := ethSdkMap[chainId]
		if sdk == nil {
			err = fmt.Errorf("chain %d sdk not initialized", chainId)
			return
		}
		balance, err = getEthErc20Balance(hash, proxy, sdk.Node())
		if err != nil {
			return big.NewInt(0), err
		}
		return
	default:
		return new(big.Int).SetUint64(0), fmt.Errorf(fmt.Sprintf("chain id %d invalid", chainId))
	}
}

func GetTotalSupply(chainId uint64, hash string) (totalSupply *big.Int, err error) {
	totalSupply = big.NewInt(0)
	switch chainId {
	case base.HECO, base.BSC, base.MATIC:
		sdk := ethSdkMap[chainId]
		if sdk == nil {
			err = fmt.Errorf("chain %s sdk not initialized", chainId)
			return
		}
		if !basedef.IsNativeTokenAddress(hash) {
			erc20Address := common.HexToAddress(hash)
			contract, err := erc20.NewERC20Mintable(erc20Address, sdk.Node().Client)
			if err != nil {
				totalSupply, err = contract.TotalSupply(nil)
			}
		}
		if err != nil {
			return big.NewInt(0), err
		}
	default:
		err = fmt.Errorf(fmt.Sprintf("chain id %d invalid", chainId))
		return
	}
	return
}

func GetNftOwner(chainId uint64, asset string, tokenId int) (owner common.Address, err error) {
	owner = common.Address{}
	assetAddr := common.HexToAddress(asset)
	switch chainId {
	case base.ETH, base.MATIC, base.BSC, base.HECO, base.OK, base.ARBITRUM, base.XDAI:
		sdk := ethSdkMap[chainId]
		if sdk == nil {
			err = fmt.Errorf("chain %s sdk not initialized", chainId)
			return
		}
		var contract *nftmapping.CrossChainNFTMapping
		contract, err = nftmapping.NewCrossChainNFTMapping(assetAddr, sdk.Node())
		if err == nil {
			owner, err = contract.OwnerOf(nil, big.NewInt(int64(tokenId)))
		}
	default:
		err = fmt.Errorf(fmt.Sprintf("chain id %d invalid", chainId))
	}
	return
}

func GetBoundLockProxy(lockProxies []string, srcTokenHash, DstTokenHash string, srcChainId, dstChainId uint64) (string, error) {
	if sdk, exist := sdkMap[dstChainId]; exist {
		if value, ok := sdk.(*chainsdk.EthereumSdkPro); ok {
			return value.GetBoundLockProxy(lockProxies, srcTokenHash, DstTokenHash, srcChainId)
		}
	}
	return "", fmt.Errorf("chain %d is not ethereum based", dstChainId)
}
