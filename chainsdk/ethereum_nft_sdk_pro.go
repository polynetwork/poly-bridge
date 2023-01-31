package chainsdk

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func (pro *EthereumSdkPro) GetAndCheckTokenUrl(
	inquirer, asset, owner string,
	tokenId string,
) (url string, err error) {

	info := pro.GetLatest()
	if info == nil {
		return "", fmt.Errorf("all node is not working")
	}
	for info != nil {
		if url, err = info.sdk.GetAndCheckNFTUri(common.HexToAddress(inquirer), common.HexToAddress(asset), common.HexToAddress(owner), tokenId); err != nil {
			info = pro.reset(info)
		} else {
			return
		}
	}
	return
}

func (pro *EthereumSdkPro) NFTBalance(asset, owner string) (balance *big.Int, err error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}

	for info != nil {
		if balance, err = info.sdk.GetNFTBalance(common.HexToAddress(asset), common.HexToAddress(owner)); err != nil {
			info = pro.reset(info)
		} else {
			return
		}
	}
	return
}

func (pro *EthereumSdkPro) GetNFTUrl(asset string, tokenId string) (url string, err error) {
	info := pro.GetLatest()
	if info == nil {
		return "", fmt.Errorf("all node is not working")
	}

	for info != nil {
		if url, err = info.sdk.GetNFTTokenUri(common.HexToAddress(asset), tokenId); err != nil {
			info = pro.reset(info)
		} else {
			return
		}
	}
	return
}

func (pro *EthereumSdkPro) GetUnCrossChainNFTsByIndex(
	inquirer, asset string,
	lockProxies []string,
	start, length int,
) (mp map[string]string, err error) {

	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}
	lockProxiesAddr := make([]common.Address, len(lockProxies))
	for i, lockProxy := range lockProxies {
		lockProxiesAddr[i] = common.HexToAddress(lockProxy)
	}
	for info != nil {
		if mp, err = info.sdk.GetUnCrossChainNFTsByIndex(common.HexToAddress(inquirer), common.HexToAddress(asset), lockProxiesAddr, start, length); err != nil {
			info = pro.reset(info)
		} else {
			return
		}
	}
	return
}

func (pro *EthereumSdkPro) GetTokensByIndex(
	inquirer, asset, owner string,
	start, length int,
) (res map[string]string, err error) {

	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}
	fmt.Println("begin look up", inquirer, asset, owner)
	for info != nil {
		if res, err = info.sdk.GetOwnerNFTsByIndex(common.HexToAddress(inquirer), common.HexToAddress(asset), common.HexToAddress(owner), start, length); err != nil {
			info = pro.reset(info)
		} else {
			return
		}
	}
	return
}

func (pro *EthereumSdkPro) GetNFTURLs(asset string, tokenIds []*big.Int) (res map[string]string, err error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}
	for info != nil {
		if res, err = info.sdk.GetOwnerNFTUrls(common.HexToAddress(asset), tokenIds); err != nil {
			info = pro.reset(info)
		} else {
			return
		}
	}
	return
}
