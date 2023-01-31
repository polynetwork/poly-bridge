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
	"github.com/beego/beego/v2/core/logs"
	"math/big"
	"strings"

	erc20 "poly-bridge/go_abi/mintable_erc20_abi"
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
	logs.Info("filterTokenInfo: res=%+v", res)
	return res
}
