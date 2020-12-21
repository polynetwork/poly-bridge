package ethereumlisten

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/prometheus/common/log"
	"math/big"
)

type EthereumSdk struct {
	rpcClient   *rpc.Client
	rawClient   *ethclient.Client
	urls        []string
	node        int
}

func NewEthereumSdk(url []string) (*EthereumSdk,error) {
	node := -1
	var rpcClient *rpc.Client
	var rawClient *ethclient.Client
	var err1 error
	var err2 error
	for (node + 1) < len(url) {
		node ++
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
		urls: url,
		node: node,
	}, nil
}

func (ec *EthereumSdk) NextClient() (int, error) {
	ec.node ++
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

func (ec *EthereumSdk) GetCurrentBlockHeight() (uint64, error) {
	var result hexutil.Big
	cur := ec.node
	err := ec.rpcClient.CallContext(context.Background(), &result, "eth_blockNumber")
	for err != nil {
		log.Errorf("EthereumClient.GetCurrentBlockHeight err:%s, url: %s", err.Error(), ec.urls[ec.node])
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
		log.Errorf("EthereumClient.GetHeaderByNumber err:%s, url: %s", err.Error(), ec.urls[ec.node])
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

