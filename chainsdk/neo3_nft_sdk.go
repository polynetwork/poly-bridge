package chainsdk

import (
	"fmt"
	"math/big"
	"sort"
)

func (sdk *Neo3Sdk) GetAndCheckNFTUri(queryAddr, asset, owner, tokenId string) (string, error) {
	tokenOwner, err := sdk.Nep11OwnerOf(asset, tokenId)
	if tokenOwner != owner {
		fmt.Println(tokenOwner, owner)
		return "", fmt.Errorf("owner token not exist")
	}
	tokenUrl, err := sdk.Nep11TokenUri(asset, tokenId)
	if err != nil {
		return "", err
	}
	return tokenUrl, nil
}

func (sdk *Neo3Sdk) GetNFTTokenUri(asset, tokenId string) (string, error) {
	tokenUri, err := sdk.Nep11TokenUri(asset, tokenId)
	if err != nil {
		return "", err
	}
	return tokenUri, nil
}

func (sdk *Neo3Sdk) GetNFTBalance(asset, owner string) (*big.Int, error) {
	balanceStr, err := sdk.Nep11BalanceOf(asset, owner)
	if err != nil {
		return nil, err
	}
	balance, _ := big.NewInt(0).SetString(balanceStr, 10)
	return balance, nil
}

func (sdk *Neo3Sdk) GetOwnerNFTsByIndex(queryAddr, asset, owner string, start, length int) (map[string]string, error) {
	tokenIdsAll, err := sdk.Nep11TokensOf(asset, owner)
	var tokenIds []string
	if err != nil {
		return nil, err
	}
	sort.Strings(tokenIdsAll)
	end := start + length
	if end < len(tokenIdsAll) {
		tokenIds = tokenIdsAll[start:end]
	} else {
		tokenIds = tokenIdsAll[start:]
	}
	return sdk.Nep11UriByBatchInvoke(asset, tokenIds)
}
