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

package chainsdk

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"poly-bridge/chainsdk/usdt_abi"
)

type EthereumSdk struct {
	rpcClient *rpc.Client
	rawClient *ethclient.Client
	url       string
}

func NewEthereumSdk(url string) (*EthereumSdk, error) {
	rpcClient, err1 := rpc.Dial(url)
	rawClient, err2 := ethclient.Dial(url)
	if rpcClient == nil || err1 != nil || rawClient == nil || err2 != nil {
		return nil, fmt.Errorf("ethereum node is not working!, err1: %v, err2: %v", err1, err2)
	}
	return &EthereumSdk{
		rpcClient: rpcClient,
		rawClient: rawClient,
		url:       url,
	}, nil
}

func (ec *EthereumSdk) GetClient() *ethclient.Client {
	return ec.rawClient
}

func (ec *EthereumSdk) GetCurrentBlockHeight() (uint64, error) {
	var result hexutil.Big
	err := ec.rpcClient.CallContext(context.Background(), &result, "eth_blockNumber")
	for err != nil {
		return 0, err
	}
	return (*big.Int)(&result).Uint64(), err
}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	return hexutil.EncodeBig(number)
}

// GetHeaderByNumber returns the given header
func (ec *EthereumSdk) GetHeaderByNumber(number uint64) (*types.Header, error) {
	var header *types.Header
	var newNumber *big.Int
	if number < 0 {
		newNumber = nil
	} else {
		newNumber = big.NewInt(int64(number))
	}
	err := ec.rpcClient.CallContext(context.Background(), &header, "eth_getBlockByNumber", toBlockNumArg(newNumber), false)
	for err != nil {
		return nil, err
	}
	return header, err
}

func (ec *EthereumSdk) GetBlockByNumber(number uint64) (*types.Block, error) {
	return ec.rawClient.BlockByNumber(context.Background(), new(big.Int).SetUint64(number))
}

func (ec *EthereumSdk) GetTransactionByHash(hash common.Hash) (*types.Transaction, error) {
	tx, _, err := ec.rawClient.TransactionByHash(context.Background(), hash)
	for err != nil {
		return nil, err
	}
	return tx, err
}

func (ec *EthereumSdk) GetTransactionReceipt(hash common.Hash) (*types.Receipt, error) {
	receipt, err := ec.rawClient.TransactionReceipt(context.Background(), hash)
	for err != nil {
		return nil, err
	}
	return receipt, nil
}

func (ec *EthereumSdk) NonceAt(addr common.Address) (uint64, error) {
	nonce, err := ec.rawClient.PendingNonceAt(context.Background(), addr)
	for err != nil {
		return 0, err
	}
	return nonce, nil
}

func (ec *EthereumSdk) SendRawTransaction(tx *types.Transaction) error {
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return err
	}
	err = ec.rpcClient.CallContext(context.Background(), nil, "eth_sendRawTransaction", hexutil.Encode(data))
	for err != nil {
		return err
	}
	return nil
}

func (ec *EthereumSdk) TransactionByHash(hash common.Hash) (*types.Transaction, bool, error) {
	tx, isPending, err := ec.rawClient.TransactionByHash(context.Background(), hash)
	for err != nil {
		return nil, false, err
	}
	return tx, isPending, err
}

func (ec *EthereumSdk) SuggestGasPrice() (*big.Int, error) {
	gasPrice, err := ec.rawClient.SuggestGasPrice(context.Background())
	for err != nil {
		return nil, err
	}
	return gasPrice, err
}

func (ec *EthereumSdk) EstimateGas(msg ethereum.CallMsg) (uint64, error) {
	gasLimit, err := ec.rawClient.EstimateGas(context.Background(), msg)
	for err != nil {
		return 0, err
	}
	return gasLimit, err
}

func (ec *EthereumSdk) Erc20Info(hash string) (string, string, int64, string, error) {
	erc20Address := common.HexToAddress(hash)
	erc20Contract, err := usdt_abi.NewTetherToken(erc20Address, ec.rawClient)
	if err != nil {
		return "", "", 0, "", err
	}
	name, err := erc20Contract.Name(&bind.CallOpts{})
	if err != nil {
		return "", "", 0, "", err
	}
	/*
		totolSupply, err := erc20Contract.TotalSupply(&bind.CallOpts{})
		if err != nil {
			return "", "", 0, "", err
		}
	*/
	decimal, err := erc20Contract.Decimals(&bind.CallOpts{})
	if err != nil {
		return "", "", 0, "", err
	}
	symbol, err := erc20Contract.Symbol(&bind.CallOpts{})
	if err != nil {
		return "", "", 0, "", err
	}
	return hash, name, decimal.Int64(), symbol, nil
}
