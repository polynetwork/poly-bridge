package chainsdk

import (
	"fmt"
	"math/big"
)

func (pro *Neo3SdkPro) GetAndCheckTokenUrl(
	inquirer, asset, owner string,
	tokenId string,
) (url string, err error) {

	info := pro.GetLatest()
	if info == nil {
		return "", fmt.Errorf("all node is not working")
	}
	for info != nil {
		if url, err = info.sdk.GetAndCheckNFTUri(inquirer, asset, owner, tokenId); err != nil {
			info = pro.reset(info)
		} else {
			return
		}
	}
	return
}
func (pro *Neo3SdkPro) NFTBalance(asset, owner string) (balance *big.Int, err error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}

	for info != nil {
		if balance, err = info.sdk.GetNFTBalance(asset, owner); err != nil {
			info = pro.reset(info)
		} else {
			return
		}
	}
	return
}

func (pro *Neo3SdkPro) GetNFTUrl(asset string, tokenId string) (url string, err error) {
	info := pro.GetLatest()
	if info == nil {
		return "", fmt.Errorf("all node is not working")
	}

	for info != nil {
		if url, err = info.sdk.GetNFTTokenUri(asset, tokenId); err != nil {
			info = pro.reset(info)
		} else {
			return
		}
	}
	return
}

func (pro *Neo3SdkPro) GetUnCrossChainNFTsByIndex(
	inquirer, asset string,
	lockProxies []string,
	start, length int,
) (mp map[string]string, err error) {
	return
}

func (pro *Neo3SdkPro) GetTokensByIndex(
	inquirer, asset, owner string,
	start, length int,
) (res map[string]string, err error) {

	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}

	for info != nil {
		if res, err = info.sdk.GetOwnerNFTsByIndex(inquirer, asset, owner, start, length); err != nil {
			info = pro.reset(info)
		} else {
			return
		}
	}
	return
}

//
//func (pro *Neo3SdkPro) GetNFTURLs(asset string, tokenIds []*big.Int) (res map[string]string, err error) {
//	info := pro.GetLatest()
//	if info == nil {
//		return nil, fmt.Errorf("all node is not working")
//	}
//	for info != nil {
//		info.sdk.
//		if res, err = info.sdk.GetOwnerNFTUrls(asset, tokenIds); err != nil {
//			info = pro.reset(info)
//		} else {
//			return
//		}
//	}
//	return
//}
//
