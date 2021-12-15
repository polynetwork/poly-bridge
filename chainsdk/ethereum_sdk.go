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
	"errors"
	"fmt"
	"time"

	log "github.com/beego/beego/v2/core/logs"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/polynetwork/bridge-common/wallet"

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

func (ec *EthereumSdk) EthBalance(addr string) (*big.Int, error) {
	var result hexutil.Big
	ctx := context.Background()
	err := ec.rpcClient.CallContext(ctx, &result, "eth_getBalance", "0x"+addr, "latest")
	return (*big.Int)(&result), err
}

type SwapNativeTokenMeta struct {
	SrcChain               uint64
	SrcHeight              uint64
	DstChain               uint64
	SrcAddress, DstAddress string
	SrcHash, DstHash       string
	SrcAmount, DstAmount   *big.Int
}

func (m *SwapNativeTokenMeta) DecodeSrcData(data []byte) (err error) {
	if len(data) < 24 {
		err = fmt.Errorf("Invalid src data: %x", data)
		return
	}
	m.SrcChain = uint64(btoi32(data[0:4]))
	m.DstAddress = common.BytesToAddress(data[4:24]).String()
	return
}

func (m *SwapNativeTokenMeta) EncodeDstData() (data []byte, err error) {
	src := common.HexToHash(m.SrcHash)
	data = append(i32tob(uint32(m.SrcChain)), src.Bytes()...)
	return
}

func IsEmpty(v []byte) bool {
	for _, b := range v {
		if b > 0 {
			return false
		}
	}
	return true
}

func i32tob(val uint32) []byte {
	r := make([]byte, 4)
	for i := uint32(0); i < 4; i++ {
		r[i] = byte((val >> (8 * i)) & 0xff)
	}
	return r
}

func btoi32(val []byte) uint32 {
	r := uint32(0)
	for i := uint32(0); i < 4; i++ {
		r |= uint32(val[i]) << (8 * i)
	}
	return r
}

// Swap native token to dst receiver
func (ec *EthereumSdk) SwapETH(wallet wallet.IWallet, tx *SwapNativeTokenMeta) (err error) {
	address := common.HexToAddress(tx.DstAddress)
	src := common.HexToHash(tx.SrcHash)
	if tx.DstAddress == "" || IsEmpty(address.Bytes()) || IsEmpty(src.Bytes()) || tx.DstAmount == nil {
		err = fmt.Errorf("Invalid swap tx %+v", *tx)
		return
	}
	data, err := tx.EncodeDstData()
	if err != nil {
		return
	}
	tx.DstHash, err = wallet.Send(address, tx.DstAmount, 0, nil, nil, data)
	return
}

// Fetch src transaction detail
func (ec *EthereumSdk) FetchSwapTx(hash string) (tx *SwapNativeTokenMeta, err error) {
	srcHash := common.HexToHash(hash)
	t, err := ec.TransactionWithExtraByHash(context.Background(), srcHash)
	if err != nil || t == nil || t.BlockNumber == nil || t.Tx() == nil {
		return
	}
	v := big.NewInt(0)
	v.SetString((*t.BlockNumber)[2:], 16)
	height := v.Uint64()

	tx = &SwapNativeTokenMeta{
		SrcHeight:  height,
		SrcHash:    srcHash.String(),
		SrcAddress: t.From.String(),
		SrcAmount:  t.Tx().Value(),
	}

	err = tx.DecodeSrcData(t.Tx().Data())
	return
}

// TransactionByHash returns the transaction with the given hash.
func (ec *EthereumSdk) TransactionWithExtraByHash(ctx context.Context, hash common.Hash) (json *rpcTransaction, err error) {
	err = ec.rpcClient.CallContext(ctx, &json, "eth_getTransactionByHash", hash)
	return
	/*
		if err != nil {
			return nil, err
		} else if json == nil || json.tx == nil {
			return nil, nil
		} else if _, r, _ := json.tx.RawSignatureValues(); r == nil {
			return nil, fmt.Errorf("server returned transaction without signature")
		}
		if json.From != nil && json.BlockHash != nil {
			setSenderFromServer(json.tx, *json.From, *json.BlockHash)
		}
		return json, nil
	*/
}

func (ec *EthereumSdk) GetTxHeight(ctx context.Context, hash common.Hash) (height uint64, pending bool, err error) {
	tx, err := ec.TransactionWithExtraByHash(context.Background(), hash)
	if err != nil || tx == nil {
		return
	}
	pending = tx.BlockNumber == nil
	if !pending {
		v := big.NewInt(0)
		v.SetString((*tx.BlockNumber)[2:], 16)
		height = v.Uint64()
	}
	return
}

func (ec *EthereumSdk) Confirm(hash common.Hash, blocks uint64, count int) (height, confirms uint64, pending bool, err error) {
	var current uint64
	for count > 0 {
		count--
		confirms = 0
		height, pending, err = ec.GetTxHeight(context.Background(), hash)
		if height > 0 {
			if blocks == 0 {
				return
			}
			current, err = ec.rawClient.BlockNumber(context.Background())
			if current >= height {
				confirms = current - height
				if confirms >= blocks {
					return
				}
			}
		}
		if err != nil {
			log.Info("Wait poly tx confirmation error", "count", count, "hash", hash, "err", err)
		}
		time.Sleep(time.Second)
	}
	return
}

type rpcTransaction struct {
	tx *types.Transaction
	txExtraInfo
}

func (t *rpcTransaction) Tx() *types.Transaction {
	return t.tx
}

type txExtraInfo struct {
	BlockNumber *string         `json:"blockNumber,omitempty"`
	BlockHash   *common.Hash    `json:"blockHash,omitempty"`
	From        *common.Address `json:"from,omitempty"`
}

// senderFromServer is a types.Signer that remembers the sender address returned by the RPC
// server. It is stored in the transaction's sender address cache to avoid an additional
// request in TransactionSender.
type senderFromServer struct {
	addr      common.Address
	blockhash common.Hash
}

func setSenderFromServer(tx *types.Transaction, addr common.Address, block common.Hash) {
	// Use types.Sender for side-effect to store our signer into the cache.
	types.Sender(&senderFromServer{addr, block}, tx)
}

func (s *senderFromServer) Equal(other types.Signer) bool {
	os, ok := other.(*senderFromServer)
	return ok && os.blockhash == s.blockhash
}

func (s *senderFromServer) Sender(tx *types.Transaction) (common.Address, error) {
	if s.blockhash == (common.Hash{}) {
		return common.Address{}, errors.New("sender not cached")
	}
	return s.addr, nil
}

func (s *senderFromServer) ChainID() *big.Int {
	panic("can't sign with senderFromServer")
}
func (s *senderFromServer) Hash(tx *types.Transaction) common.Hash {
	panic("can't sign with senderFromServer")
}
func (s *senderFromServer) SignatureValues(tx *types.Transaction, sig []byte) (R, S, V *big.Int, err error) {
	panic("can't sign with senderFromServer")
}
