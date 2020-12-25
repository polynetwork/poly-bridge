package ethereumlisten

import (
	"context"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

type EthereumSdk struct {
	rpcClient *rpc.Client
	rawClient *ethclient.Client
	urls      []string
	node      int
}

func NewEthereumSdk(url []string) (*EthereumSdk, error) {
	node := -1
	var rpcClient *rpc.Client
	var rawClient *ethclient.Client
	var err1 error
	var err2 error
	for (node + 1) < len(url) {
		node++
		rpcClient, err1 = rpc.Dial(url[node])
		rawClient, err2 = ethclient.Dial(url[node])
		if err1 == nil && err2 == nil {
			break
		}
	}
	if rpcClient == nil || err1 != nil || rawClient == nil || err2 != nil {
		return nil, fmt.Errorf("all ethereum node is not working!")
	}
	return &EthereumSdk{
		rpcClient: rpcClient,
		rawClient: rawClient,
		urls:      url,
		node:      node,
	}, nil
}

func (ec *EthereumSdk) NextClient() (int, error) {
	ec.node++
	ec.node = ec.node % len(ec.urls)
	var err1 error
	ec.rpcClient, err1 = rpc.Dial(ec.urls[ec.node])
	var err2 error
	ec.rawClient, err2 = ethclient.Dial(ec.urls[ec.node])
	err := err1
	if err == nil {
		err = err2
	}
	return ec.node, err
}

func (ec *EthereumSdk) GetLatestHeight() (uint64, error) {
	var result hexutil.Big
	latestHeight := uint64(0)
	for i, url := range ec.urls {
		rpcClient, err := rpc.Dial(url)
		if err != nil {
			continue
		}
		err = rpcClient.CallContext(context.Background(), &result, "eth_blockNumber")
		if err != nil {
			continue
		}
		height := (*big.Int)(&result).Uint64()
		if height > latestHeight {
			ec.node = i
			latestHeight = height
		}
	}
	var err1 error
	ec.rpcClient, err1 = rpc.Dial(ec.urls[ec.node])
	var err2 error
	ec.rawClient, err2 = ethclient.Dial(ec.urls[ec.node])
	err := err1
	if err == nil {
		err = err2
	}
	return latestHeight, err
}

func (ec *EthereumSdk) GetCurrentBlockHeight() (uint64, error) {
	var result hexutil.Big
	cur := ec.node
	err := ec.rpcClient.CallContext(context.Background(), &result, "eth_blockNumber")
	for err != nil {
		logs.Error("EthereumClient.GetCurrentBlockHeight err:%s, url: %s", err.Error(), ec.urls[ec.node])
		next, err := ec.NextClient()
		if next == cur {
			return 0, fmt.Errorf("all node is not working!")
		}
		if err != nil {
			continue
		}
		err = ec.rpcClient.CallContext(context.Background(), &result, "eth_blockNumber")
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
	cur := ec.node
	var newNumber *big.Int
	if number < 0 {
		newNumber = nil
	} else {
		newNumber = big.NewInt(int64(number))
	}
	err := ec.rpcClient.CallContext(context.Background(), &header, "eth_getBlockByNumber", toBlockNumArg(newNumber), false)
	for err != nil {
		logs.Error("EthereumClient.GetHeaderByNumber err:%s, url: %s", err.Error(), ec.urls[ec.node])
		next, err := ec.NextClient()
		if next == cur {
			return nil, fmt.Errorf("all node is not working!")
		}
		if err != nil {
			continue
		}
		err = ec.rpcClient.CallContext(context.Background(), &header, "eth_getBlockByNumber", toBlockNumArg(newNumber), false)
	}
	return header, err
}

func (ec *EthereumSdk) GetTransactionByHash(hash common.Hash) (*types.Transaction, error) {
	cur := ec.node
	tx, _, err := ec.rawClient.TransactionByHash(context.Background(), hash)
	for err != nil {
		logs.Error("EthereumClient.GetTransactionByHash err:%s, url: %s", err.Error(), ec.urls[ec.node])
		next, err := ec.NextClient()
		if next == cur {
			return nil, fmt.Errorf("all node is not working!")
		}
		if err != nil {
			continue
		}
		tx, _, err = ec.rawClient.TransactionByHash(context.Background(), hash)
	}
	return tx, err
}

func (ec *EthereumSdk) GetTransactionReceipt(hash common.Hash) (*types.Receipt, error) {
	cur := ec.node
	receipt, err := ec.rawClient.TransactionReceipt(context.Background(), hash)
	for err != nil {
		logs.Error("EthereumClient.GetTransactionReceipt err:%s, url: %s", err.Error(), ec.urls[ec.node])
		next, err := ec.NextClient()
		if next == cur {
			return nil, fmt.Errorf("all node is not working!")
		}
		if err != nil {
			continue
		}
		receipt, err = ec.rawClient.TransactionReceipt(context.Background(), hash)
	}
	return receipt, nil
}

func (ec *EthereumSdk) NonceAt(addr common.Address) (uint64, error) {
	cur := ec.node
	nonce, err := ec.rawClient.PendingNonceAt(context.Background(), addr)
	for err != nil {
		logs.Error("EthereumClient.NonceAt err:%s, url: %s", err.Error(), ec.urls[ec.node])
		next, err := ec.NextClient()
		if next == cur {
			return 0, fmt.Errorf("all node is not working!")
		}
		if err != nil {
			continue
		}
		nonce, err = ec.rawClient.PendingNonceAt(context.Background(), addr)
	}
	return nonce, nil
}


func (ec *EthereumSdk) SendRawTransaction(tx *types.Transaction) error {
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return err
	}
	cur := ec.node
	err = ec.rpcClient.CallContext(context.Background(), nil, "eth_sendRawTransaction", hexutil.Encode(data))
	for err != nil {
		logs.Error("EthereumClient.SendRawTransaction err:%s, url: %s", err.Error(), ec.urls[ec.node])
		next, err := ec.NextClient()
		if next == cur {
			return fmt.Errorf("all node is not working!")
		}
		if err != nil {
			continue
		}
		err = ec.rpcClient.CallContext(context.Background(), nil, "eth_sendRawTransaction", hexutil.Encode(data))
	}
	return nil
}

func (ec *EthereumSdk) TransactionByHash(hash common.Hash) (*types.Transaction, bool, error) {
	cur := ec.node
	tx, isPending, err := ec.rawClient.TransactionByHash(context.Background(), hash)
	for err != nil {
		logs.Error("EthereumClient.TransactionByHash err:%s, url: %s", err.Error(), ec.urls[ec.node])
		next, err := ec.NextClient()
		if next == cur {
			return nil, false, fmt.Errorf("all node is not working!")
		}
		if err != nil {
			continue
		}
		tx, isPending, err = ec.rawClient.TransactionByHash(context.Background(), hash)
	}
	return tx, isPending, err
}

func (ec *EthereumSdk) SuggestGasPrice() (*big.Int, error) {
	cur := ec.node
	gasPrice, err := ec.rawClient.SuggestGasPrice(context.Background())
	for err != nil {
		logs.Error("EthereumClient.SuggestGasPrice err:%s, url: %s", err.Error(), ec.urls[ec.node])
		next, err := ec.NextClient()
		if next == cur {
			return nil, fmt.Errorf("all node is not working!")
		}
		if err != nil {
			continue
		}
		gasPrice, err = ec.rawClient.SuggestGasPrice(context.Background())
	}
	return gasPrice, err
}

func (ec *EthereumSdk) EstimateGas(msg ethereum.CallMsg) (uint64, error) {
	cur := ec.node
	gasLimit, err := ec.rawClient.EstimateGas(context.Background(), msg)
	for err != nil {
		logs.Error("EthereumClient.EstimateGas err:%s, url: %s", err.Error(), ec.urls[ec.node])
		next, err := ec.NextClient()
		if next == cur {
			return 0, fmt.Errorf("all node is not working!")
		}
		if err != nil {
			continue
		}
		gasLimit, err = ec.rawClient.EstimateGas(context.Background(), msg)
	}
	return gasLimit, err
}


