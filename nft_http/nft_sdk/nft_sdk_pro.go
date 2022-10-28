package nft_sdk

import (
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
)

type INftSdkPro interface {
	GetAndCheckTokenUrl(inquirer string, asset string, owner string, tokenId string) (url string, err error)
	NFTBalance(asset string, owner string) (balance *big.Int, err error)
	GetTokensByIndex(inquirer string, asset string, owner string, start int, length int) (res map[string]string, err error)
	GetNFTUrl(asset string, tokenId string) (url string, err error)
	GetUnCrossChainNFTsByIndex(inquirer string, asset string, lockProxies []string, start int, length int) (mp map[string]string, err error)
}

func SelectNftSdkPro(chainId uint64, urls []string, slot uint64) INftSdkPro {
	switch chainId {
	case basedef.NEO3_CROSSCHAIN_ID:
		return chainsdk.NewNeo3SdkPro(urls, slot, chainId)
	default:
		return chainsdk.NewEthereumSdkPro(urls, slot, chainId)

	}
}
