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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"poly-bridge/basedef"

	//"github.com/polynetwork/eth-contracts/go_abi/erc20_abi"
	"math/big"
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

func (s *EthereumSdk) GetClient() *ethclient.Client {
	return s.rawClient
}

func (s *EthereumSdk) GetUrl() string {
	return s.url
}

func (s *EthereumSdk) GetCurrentBlockHeight() (uint64, error) {
	var result hexutil.Big
	err := s.rpcClient.CallContext(context.Background(), &result, "eth_blockNumber")
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
func (s *EthereumSdk) GetHeaderByNumber(number uint64) (*types.Header, error) {
	var header *types.Header
	var newNumber *big.Int
	if number < 0 {
		newNumber = nil
	} else {
		newNumber = big.NewInt(int64(number))
	}
	err := s.rpcClient.CallContext(context.Background(), &header, "eth_getBlockByNumber", toBlockNumArg(newNumber), false)
	for err != nil {
		return nil, err
	}
	return header, err
}

// GetBlockTimeByNumber returns the timestamp of given block number
func (s *EthereumSdk) GetBlockTimeByNumber(chainId, number uint64) (timestamp uint64, err error) {
	type Header struct {
		Time string `json:"timestamp"`
	}

	var header interface{}
	switch chainId {
	case basedef.ZKSYNC_CROSSCHAIN_ID, basedef.CELO_CROSSCHAIN_ID:
		header = &Header{}
	default:
		header = &types.Header{}
	}

	var newNumber *big.Int
	if number < 0 {
		newNumber = nil
	} else {
		newNumber = big.NewInt(int64(number))
	}

	err = s.rpcClient.CallContext(context.Background(), &header, "eth_getBlockByNumber", toBlockNumArg(newNumber), false)
	for err != nil {
		return 0, err
	}
	switch chainId {
	case basedef.ZKSYNC_CROSSCHAIN_ID, basedef.CELO_CROSSCHAIN_ID:
		if res, ok := header.(*Header); ok {
			timestamp, err = hexutil.DecodeUint64(res.Time)
		}
	default:
		if res, ok := header.(*types.Header); ok {
			timestamp = res.Time
		}
	}
	return
}

func (s *EthereumSdk) GetBlockByNumber(number uint64) (*types.Block, error) {
	return s.rawClient.BlockByNumber(context.Background(), new(big.Int).SetUint64(number))
}

func (s *EthereumSdk) GetTransactionByHash(hash common.Hash) (*types.Transaction, error) {
	tx, _, err := s.rawClient.TransactionByHash(context.Background(), hash)
	for err != nil {
		return nil, err
	}
	return tx, err
}

func (s *EthereumSdk) GetTransactionReceipt(hash common.Hash) (*types.Receipt, error) {
	receipt, err := s.rawClient.TransactionReceipt(context.Background(), hash)
	for err != nil {
		return nil, err
	}
	return receipt, nil
}

func (s *EthereumSdk) NonceAt(addr common.Address) (uint64, error) {
	nonce, err := s.rawClient.PendingNonceAt(context.Background(), addr)
	for err != nil {
		return 0, err
	}
	return nonce, nil
}

func (s *EthereumSdk) SendRawTransaction(tx *types.Transaction) error {
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return err
	}
	err = s.rpcClient.CallContext(context.Background(), nil, "eth_sendRawTransaction", hexutil.Encode(data))
	for err != nil {
		return err
	}
	return nil
}

func (s *EthereumSdk) TransactionByHash(hash common.Hash) (*types.Transaction, bool, error) {
	tx, isPending, err := s.rawClient.TransactionByHash(context.Background(), hash)
	for err != nil {
		return nil, false, err
	}
	return tx, isPending, err
}

func (s *EthereumSdk) SuggestGasPrice() (*big.Int, error) {
	gasPrice, err := s.rawClient.SuggestGasPrice(context.Background())
	for err != nil {
		return nil, err
	}
	return gasPrice, err
}

func (s *EthereumSdk) EstimateGas(msg ethereum.CallMsg) (uint64, error) {
	gasLimit, err := s.rawClient.EstimateGas(context.Background(), msg)
	for err != nil {
		return 0, err
	}
	return gasLimit, err
}

//func (ec *EthereumSdk) Erc20Info(hash string) (string, string, int64, string, error) {
//	erc20Address := common.HexToAddress(hash)
//	erc20Contract, err := usdt_abi.NewTetherToken(erc20Address, ec.rawClient)
//	if err != nil {
//		return "", "", 0, "", err
//	}
//	name, err := erc20Contract.Name(&bind.CallOpts{})
//	if err != nil {
//		return "", "", 0, "", err
//	}
//	/*
//		totolSupply, err := erc20Contract.TotalSupply(&bind.CallOpts{})
//		if err != nil {
//			return "", "", 0, "", err
//		}
//	*/
//	decimal, err := erc20Contract.Decimals(&bind.CallOpts{})
//	if err != nil {
//		return "", "", 0, "", err
//	}
//	symbol, err := erc20Contract.Symbol(&bind.CallOpts{})
//	if err != nil {
//		return "", "", 0, "", err
//	}
//	return hash, name, decimal.Int64(), symbol, nil
//}

//func (ec *EthereumSdk) Erc20Balance(erc20 string, addr string) (uint64, error) {
//	erc20Address := common.HexToAddress(erc20)
//	erc20Contract, err := erc20_abi.NewERC20(erc20Address, ec.rawClient)
//	if err != nil {
//		return 0, fmt.Errorf("erc20 address is not right")
//	}
//	userAddress := common.HexToAddress(addr)
//	balance, err := erc20Contract.BalanceOf(&bind.CallOpts{}, userAddress)
//	if err != nil {
//		return 0, err
//	}
//	return balance.Uint64(), nil
//}

func (s *EthereumSdk) EthBalance(addr string) (*big.Int, error) {
	var result hexutil.Big
	ctx := context.Background()
	err := s.rpcClient.CallContext(ctx, &result, "eth_getBalance", "0x"+addr, "latest")
	return (*big.Int)(&result), err
}

func (s *EthereumSdk) FilterLog(FromBlock *big.Int, ToBlock *big.Int, Addresses []common.Address) ([]types.Log, error) {
	ctx := context.Background()
	var filterQuery ethereum.FilterQuery
	filterQuery.FromBlock = FromBlock
	filterQuery.ToBlock = ToBlock
	filterQuery.Addresses = Addresses
	return s.rawClient.FilterLogs(ctx, filterQuery)
}
