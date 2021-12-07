/*
 * Copyright (C) 2020 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

// Notice: functions in this file only used for deploy_tool and test cases.

package chainsdk

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	erc20 "poly-bridge/go_abi/mintable_erc20_abi"
	nftmapping "poly-bridge/go_abi/nft_mapping_abi"
	nftquery "poly-bridge/go_abi/nft_query_abi"
	nftwrap "poly-bridge/go_abi/nft_wrap_abi"
	xecdsa "poly-bridge/utils/ecdsa"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	polycm "github.com/polynetwork/poly/common"
)

var NativeFeeToken = common.HexToAddress("0x0000000000000000000000000000000000000000")

func (s *EthereumSdk) TransferNative(
	key *ecdsa.PrivateKey,
	to common.Address,
	amount *big.Int,
) (common.Hash, error) {

	from := xecdsa.Key2address(key)
	nonce, err := s.NonceAt(from)
	if err != nil {
		return EmptyHash, err
	}

	gasPrice, err := s.SuggestGasPrice()
	if err != nil {
		return EmptyHash, err
	}

	gasLimit, err := s.EstimateGas(ethereum.CallMsg{
		From: from, To: &to, Gas: 0, GasPrice: gasPrice,
		Value: amount, Data: []byte{},
	})
	if err != nil {
		return EmptyHash, err
	}

	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, []byte{})
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, key)
	if err != nil {
		return EmptyHash, err
	}
	if err := s.SendRawTransaction(signedTx); err != nil {
		return EmptyHash, err
	}

	if err := s.waitTxConfirm(signedTx.Hash()); err != nil {
		return EmptyHash, err
	}
	return signedTx.Hash(), nil
}

func (s *EthereumSdk) GetNativeBalance(owner common.Address) (*big.Int, error) {
	return s.rawClient.BalanceAt(context.Background(), owner, nil)
}

func (s *EthereumSdk) MintERC20Token(
	key *ecdsa.PrivateKey,
	asset, to common.Address,
	amount *big.Int) (common.Hash, error) {

	contract, err := erc20.NewERC20Mintable(asset, s.backend())
	if err != nil {
		return EmptyHash, err
	}

	auth, err := s.makeAuth(key, DefaultGasLimit)
	if err != nil {
		return EmptyHash, err
	}

	tx, err := contract.Mint(auth, to, amount)
	if err != nil {
		return EmptyHash, err
	}

	if err := s.waitTxConfirm(tx.Hash()); err != nil {
		return EmptyHash, err
	}

	return tx.Hash(), nil
}

func (s *EthereumSdk) TransferERC20Token(
	key *ecdsa.PrivateKey,
	asset, to common.Address,
	amount *big.Int,
) (common.Hash, error) {

	contract, err := erc20.NewERC20Mintable(asset, s.backend())
	if err != nil {
		return EmptyHash, err
	}

	auth, err := s.makeAuth(key, DefaultGasLimit)
	if err != nil {
		return EmptyHash, err
	}
	tx, err := contract.Transfer(auth, to, amount)
	if err != nil {
		return EmptyHash, err
	}

	if err := s.waitTxConfirm(tx.Hash()); err != nil {
		return EmptyHash, err
	}

	return tx.Hash(), nil
}

func (s *EthereumSdk) GetERC20Balance(asset, owner common.Address) (*big.Int, error) {
	contract, err := erc20.NewERC20Mintable(asset, s.backend())
	if err != nil {
		return nil, err
	}
	return contract.BalanceOf(nil, owner)
}

func (s *EthereumSdk) GetERC20TotalSupply(asset common.Address) (*big.Int, error) {
	contract, err := erc20.NewERC20Mintable(asset, s.backend())
	if err != nil {
		return nil, err
	}
	return contract.TotalSupply(nil)
}

func (s *EthereumSdk) ApproveERC20Token(
	key *ecdsa.PrivateKey,
	asset, spender common.Address,
	amount *big.Int,
) (common.Hash, error) {

	contract, err := erc20.NewERC20Mintable(asset, s.backend())
	if err != nil {
		return EmptyHash, err
	}

	auth, err := s.makeAuth(key, DefaultGasLimit)
	if err != nil {
		return EmptyHash, err
	}
	tx, err := contract.Approve(auth, spender, amount)
	if err != nil {
		return EmptyHash, err
	}

	if err := s.waitTxConfirm(tx.Hash()); err != nil {
		return EmptyHash, err
	}

	return tx.Hash(), nil
}

func (s *EthereumSdk) GetERC20Allowance(asset, owner, spender common.Address) (*big.Int, error) {
	contract, err := erc20.NewERC20Mintable(asset, s.backend())
	if err != nil {
		return nil, err
	}
	return contract.Allowance(nil, owner, spender)
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

func (s *EthereumSdk) GetNFTTotalSupply(asset common.Address) (*big.Int, error) {
	cm, err := nftmapping.NewCrossChainNFTMapping(asset, s.backend())
	if err != nil {
		return nil, err
	}
	return cm.TotalSupply(nil)
}

func (s *EthereumSdk) GetNFTBalance(asset, owner common.Address) (*big.Int, error) {
	cm, err := nftmapping.NewCrossChainNFTMapping(asset, s.backend())
	if err != nil {
		return nil, err
	}
	return cm.BalanceOf(nil, owner)
}

func (s *EthereumSdk) GetNFTTokenUri(asset common.Address, tokenID *big.Int) (string, error) {
	cm, err := nftmapping.NewCrossChainNFTMapping(asset, s.backend())
	if err != nil {
		return "", err
	}
	return cm.TokenURI(nil, tokenID)
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

func (s *EthereumSdk) GetAndCheckNFTUrl(queryAddr, asset, owner common.Address, tokenId *big.Int) (string, error) {
	inquirer, err := nftquery.NewPolyNFTQuery(queryAddr, s.backend())
	if err != nil {
		return "", err
	}

	ok, url, err := inquirer.GetAndCheckTokenUrl(nil, asset, owner, tokenId)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", fmt.Errorf("owner token not exist")
	}
	return url, nil
}

func (s *EthereumSdk) GetOwnerNFTByIndex(asset, owner common.Address, index int) (*big.Int, error) {
	cm, err := nftmapping.NewCrossChainNFTMapping(asset, s.backend())
	if err != nil {
		return nil, err
	}

	return cm.TokenOfOwnerByIndex(nil, owner, big.NewInt(int64(index)))
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

func (s *EthereumSdk) WrapLockWithErc20FeeToken(
	key *ecdsa.PrivateKey,
	wrapAddr,
	fromAsset,
	toAddr common.Address,
	toChainId uint64,
	tokenID *big.Int,
	feeToken common.Address,
	feeAmount *big.Int,
	id *big.Int,
) (common.Hash, error) {

	wrapper, err := nftwrap.NewPolyNFTWrapper(wrapAddr, s.backend())
	if err != nil {
		return EmptyHash, err
	}

	auth, err := s.makeAuth(key, DefaultGasLimit)
	if err != nil {
		return EmptyHash, err
	}

	tx, err := wrapper.Lock(auth, fromAsset, toChainId, toAddr, tokenID, feeToken, feeAmount, id)
	if err != nil {
		return EmptyHash, err
	}
	if err := s.waitTxConfirm(tx.Hash()); err != nil {
		return EmptyHash, err
	}

	return tx.Hash(), nil
}

func (s *EthereumSdk) WrapLockWithNativeFeeToken(
	key *ecdsa.PrivateKey,
	wrapAddr,
	fromAsset,
	toAddr common.Address,
	toChainId uint64,
	tokenID *big.Int,
	feeAmount *big.Int,
	id *big.Int,
) (common.Hash, error) {

	auth, err := s.makeAuth(key, DefaultGasLimit)
	if err != nil {
		return EmptyHash, err
	}

	contractABI, err := abi.JSON(strings.NewReader(nftwrap.PolyNFTWrapperABI))
	if err != nil {
		return EmptyHash, err
	}

	raw, err := contractABI.Pack("lock", fromAsset, toChainId, toAddr, tokenID, NativeFeeToken, feeAmount, id)
	if err != nil {
		return EmptyHash, err
	}

	unsignedTx := types.NewTransaction(auth.Nonce.Uint64(), wrapAddr, feeAmount, auth.GasLimit, auth.GasPrice, raw)
	signedTx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, key)
	if err != nil {
		return EmptyHash, err
	}

	if err := s.rawClient.SendTransaction(context.Background(), signedTx); err != nil {
		return EmptyHash, err
	}

	hash := signedTx.Hash()
	if err := s.waitTxConfirm(hash); err != nil {
		return EmptyHash, err
	}

	return hash, nil
}

type WrapLockMethod struct {
	FromAsset common.Address
	ToChainId uint64
	ToAddress common.Address
	TokenId   *big.Int
	FeeToken  common.Address
	Fee       *big.Int
	Id        *big.Int
}

func assembleSafeTransferCallData(toAddress common.Address, chainID uint64) []byte {
	sink := polycm.NewZeroCopySink(nil)
	sink.WriteVarBytes(toAddress.Bytes())
	sink.WriteUint64(chainID)
	return sink.Bytes()
}

func filterTokenInfo(enc []byte) map[string]string {
	source := polycm.NewZeroCopySource(enc)
	var (
		num     polycm.Uint256
		url     string
		tokenId *big.Int
		eof     bool
		res     = make(map[string]string)
	)
	for {
		if num, eof = source.NextHash(); !eof {
			bz := polycm.ToArrayReverse(num[:])
			tokenId = new(big.Int).SetBytes(bz)
		} else {
			break
		}
		if url, eof = source.NextString(); !eof {
			res[tokenId.String()] = url
		} else {
			break
		}
	}

	return res
}
