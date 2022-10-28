package chainsdk

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	polycm "github.com/polynetwork/poly/common"
	"math/big"
	nftmapping "poly-bridge/go_abi/nft_mapping_abi"
	nftquery "poly-bridge/go_abi/nft_query_abi"
	"regexp"
	"strings"
)

func (s *EthereumSdk) GetNFTTokenUri(asset common.Address, tokenID string) (string, error) {
	cm, err := nftmapping.NewCrossChainNFTMapping(asset, s.backend())
	if err != nil {
		return "", err
	}
	tokenIdBigInt, err := checkNumString(tokenID)
	if err != nil {
		return "", err
	}
	return cm.TokenURI(nil, tokenIdBigInt)
}

func (s *EthereumSdk) GetOwnerNFTUrls(asset common.Address, tokenIds []*big.Int) (map[string]string, error) {
	cm, err := nftmapping.NewCrossChainNFTMapping(asset, s.backend())
	if err != nil {
		return nil, err
	}

	res := make(map[string]string)
	for _, tokenId := range tokenIds {
		url, err := cm.TokenURI(nil, tokenId)
		if err == nil {
			res[tokenId.String()] = url
		}
	}
	return res, nil
}

func (s *EthereumSdk) GetNFTApproved(asset common.Address, tokenID *big.Int) (common.Address, error) {
	cm, err := nftmapping.NewCrossChainNFTMapping(asset, s.backend())
	if err != nil {
		return EmptyAddress, err
	}
	return cm.GetApproved(nil, tokenID)
}

func (s *EthereumSdk) GetNFTOwner(asset common.Address, tokenID *big.Int) (common.Address, error) {
	cm, err := nftmapping.NewCrossChainNFTMapping(asset, s.backend())
	if err != nil {
		return EmptyAddress, err
	}
	return cm.OwnerOf(nil, tokenID)
}

func checkNumString(numStr string) (*big.Int, error) {
	numStr = strings.Trim(numStr, " ")
	if numStr == "" {
		return nil, fmt.Errorf("number string is empty")
	}

	ok, err := regexp.Match(`^\d+$`, []byte(numStr))
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("invalid number string")
	}
	data, ok := new(big.Int).SetString(numStr, 10)
	if !ok {
		return nil, fmt.Errorf("convert string to big int err")
	}
	return data, nil
}

func (s *EthereumSdk) GetAndCheckNFTUri(queryAddr, asset, owner common.Address, tokenId string) (string, error) {
	inquirer, err := nftquery.NewPolyNFTQuery(queryAddr, s.backend())
	if err != nil {
		return "", err
	}
	tokenIdBigInt, err := checkNumString(tokenId)
	if err != nil {
		return "", err
	}
	ok, url, err := inquirer.GetAndCheckTokenUrl(nil, asset, owner, tokenIdBigInt)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", fmt.Errorf("owner token not exist")
	}
	return url, nil
}

func (s *EthereumSdk) GetNFTBalance(asset, owner common.Address) (*big.Int, error) {
	cm, err := nftmapping.NewCrossChainNFTMapping(asset, s.backend())
	if err != nil {
		return nil, err
	}
	return cm.BalanceOf(nil, owner)
}

func (s *EthereumSdk) GetUnCrossChainNFTsByIndex(
	queryAddr,
	asset common.Address,
	lockProxies []common.Address,
	start, length int,
) (map[string]string, error) {

	inquirer, err := nftquery.NewPolyNFTQuery(queryAddr, s.backend())
	if err != nil {
		return nil, err
	}

	st, ln := big.NewInt(int64(start)), big.NewInt(int64(length))
	var encs []byte
	for _, lockProxy := range lockProxies {
		ok, enc, err := inquirer.GetFilterTokensByIndex(nil, asset, lockProxy, st, ln)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, nil
		}
		encs = append(encs, enc...)
	}
	res := filterTokenInfo(encs)
	return res, nil
}

func (s *EthereumSdk) GetNFTsById(queryAddr, asset common.Address, tokenIdList []*big.Int) (map[string]string, error) {
	if len(tokenIdList) == 0 {
		return nil, fmt.Errorf("empty id list")
	}

	inquirer, err := nftquery.NewPolyNFTQuery(queryAddr, s.backend())
	if err != nil {
		return nil, err
	}

	sink := polycm.NewZeroCopySink(nil)
	list := []*big.Int{big.NewInt(int64(len(tokenIdList)))}
	list = append(list, tokenIdList...)
	for _, v := range list {
		hash := common.BytesToHash(v.Bytes())
		reversed := polycm.ToArrayReverse(hash[:])
		data, err := polycm.Uint256ParseFromBytes(reversed[:])
		if err != nil {
			return nil, err
		}
		sink.WriteHash(data)
	}

	ok, enc, err := inquirer.GetTokensByIds(nil, asset, sink.Bytes())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}
	res := filterTokenInfo(enc)
	return res, nil
}

func (s *EthereumSdk) GetNFTTotalSupply(asset common.Address) (*big.Int, error) {
	cm, err := nftmapping.NewCrossChainNFTMapping(asset, s.backend())
	if err != nil {
		return nil, err
	}
	return cm.TotalSupply(nil)
}

func (s *EthereumSdk) GetOwnerNFTByIndex(asset, owner common.Address, index int) (*big.Int, error) {
	cm, err := nftmapping.NewCrossChainNFTMapping(asset, s.backend())
	if err != nil {
		return nil, err
	}

	return cm.TokenOfOwnerByIndex(nil, owner, big.NewInt(int64(index)))
}

func (s *EthereumSdk) MintNFT(
	ownerKey *ecdsa.PrivateKey,
	asset,
	to common.Address,
	tokenID *big.Int,
	uri string,
) (common.Hash, error) {

	contract, err := nftmapping.NewCrossChainNFTMapping(asset, s.rawClient)
	if err != nil {
		return EmptyHash, err
	}

	auth, err := s.makeAuth(ownerKey, DefaultGasLimit)
	if err != nil {
		return EmptyHash, err
	}

	tx, err := contract.MintWithURI(auth, to, tokenID, uri)
	if err != nil {
		return EmptyHash, err
	}

	if err := s.waitTxConfirm(tx.Hash()); err != nil {
		return EmptyHash, err
	}
	return tx.Hash(), nil
}

func (s *EthereumSdk) NFTSafeTransferTo(
	nftOwnerKey *ecdsa.PrivateKey,
	asset,
	from,
	to common.Address,
	tokenID *big.Int,
) (common.Hash, error) {

	cm, err := nftmapping.NewCrossChainNFTMapping(asset, s.backend())
	if err != nil {
		return EmptyHash, err
	}

	auth, err := s.makeAuth(nftOwnerKey, DefaultGasLimit)
	if err != nil {
		return EmptyHash, err
	}

	tx, err := cm.SafeTransferFrom(auth, from, to, tokenID)
	if err != nil {
		return EmptyHash, err
	}

	if err := s.waitTxConfirm(tx.Hash()); err != nil {
		return EmptyHash, err
	}
	return tx.Hash(), nil
}

func (s *EthereumSdk) NFTSafeTransferFrom(
	nftOwnerKey *ecdsa.PrivateKey,
	asset,
	from,
	proxy common.Address,
	tokenID *big.Int,
	to common.Address,
	toChainID uint64,
) (common.Hash, error) {

	cm, err := nftmapping.NewCrossChainNFTMapping(asset, s.backend())
	if err != nil {
		return EmptyHash, err
	}

	auth, err := s.makeAuth(nftOwnerKey, DefaultGasLimit)
	if err != nil {
		return EmptyHash, err
	}
	data := assembleSafeTransferCallData(to, toChainID)
	tx, err := cm.SafeTransferFrom0(auth, from, proxy, tokenID, data)
	if err != nil {
		return EmptyHash, err
	}

	if err := s.waitTxConfirm(tx.Hash()); err != nil {
		return EmptyHash, err
	}
	return tx.Hash(), nil
}

func (s *EthereumSdk) NFTApprove(key *ecdsa.PrivateKey, asset, to common.Address, token *big.Int) (common.Hash, error) {
	cm, err := nftmapping.NewCrossChainNFTMapping(asset, s.backend())
	if err != nil {
		return EmptyHash, err
	}
	auth, err := s.makeAuth(key, DefaultGasLimit)
	if err != nil {
		return EmptyHash, err
	}
	tx, err := cm.Approve(auth, to, token)
	if err != nil {
		return EmptyHash, err
	}
	if err := s.waitTxConfirm(tx.Hash()); err != nil {
		return EmptyHash, err
	}
	return tx.Hash(), nil
}

func (s *EthereumSdk) GetOwnerNFTsByIndex(queryAddr, asset common.Address, owner common.Address, start, length int) (map[string]string, error) {
	inquirer, err := nftquery.NewPolyNFTQuery(queryAddr, s.backend())
	if err != nil {
		return nil, err
	}

	st, ln := big.NewInt(int64(start)), big.NewInt(int64(length))
	ok, enc, err := inquirer.GetOwnerTokensByIndex(nil, asset, owner, st, ln)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}
	res := filterTokenInfo(enc)
	return res, nil
}
