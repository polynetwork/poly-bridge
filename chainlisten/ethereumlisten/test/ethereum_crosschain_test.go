package test

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"poly-swap/chainlisten/ethereumlisten"
	"poly-swap/chainlisten/ethereumlisten/wrapper_abi"
	"poly-swap/conf"
	"strings"
	"testing"
	"time"
)

func NewPrivateKey(key string) *ecdsa.PrivateKey {
	priKey, err := crypto.HexToECDSA(key)
	if err != nil {
		panic(err)
	}
	return priKey
}

func waitTransactionConfirm(ethSdk *ethereumlisten.EthereumSdk, hash common.Hash) {
	errNum := 0
	for errNum < 100 {
		time.Sleep(time.Second * 1)
		_, ispending, err := ethSdk.TransactionByHash(hash)
		fmt.Printf("transaction %s is pending: %v\n",  hash.String(), ispending)
		if err != nil {
			errNum ++
			continue
		}
		if ispending == true{
			continue
		} else {
			break
		}
	}
}

func TestEthereumCross(t *testing.T) {
	config := conf.NewConfig("./../../../conf/config_testnet.json")
	if config == nil {
		panic("read config failed!")
	}
	ethSdk, err := ethereumlisten.NewEthereumSdk(config.ChainListenConfig.EthereumChainListenConfig.RestURL)
	if err != nil {
		panic(err)
	}
	contractabi, err := abi.JSON(strings.NewReader(polywrapper.IPolyWrapperABI))
	if err != nil {
		panic(err)
	}
	assetHash := common.HexToAddress("0000000000000000000000000000000000000000")
	toAddress := common.Hex2Bytes("6e43f9988f2771f1a2b140cb3faad424767d39fc")
	txData, err := contractabi.Pack("lock", assetHash, uint64(4), toAddress, big.NewInt(int64(100000000000000000)), big.NewInt(10000000000000000))
	if err != nil {
		panic(err)
	}
	fmt.Printf("TestInvokeContract - txdata:%s\n", hex.EncodeToString(txData))
	wrapperContractAddress := common.HexToAddress(config.ChainListenConfig.EthereumChainListenConfig.WrapperContract)
	privateKey := NewPrivateKey("56b446a2de5edfccee1581fbba79e8bb5c269e28ab4c0487860afb7e2c2d2b6e")
	fromAddr := crypto.PubkeyToAddress(privateKey.PublicKey)
	fmt.Printf("user address: %s\n", fromAddr.String())
	nonce, err := ethSdk.NonceAt(fromAddr)
	if err != nil {
		panic(err)
	}
	gasPrice, err := ethSdk.SuggestGasPrice()
	if err != nil {
		panic(err)
	}
	fmt.Printf("gas price: %s\n", gasPrice.String())
	callMsg := ethereum.CallMsg{
		From: fromAddr, To: &wrapperContractAddress, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(100000000000000000), Data: txData,
	}

	gasLimit, err := ethSdk.EstimateGas(callMsg)
	if err != nil || gasLimit == 0 {
		panic(err)
	}
	fmt.Printf("gas limit: %d\n", gasLimit)
	tx := types.NewTransaction(nonce, wrapperContractAddress, big.NewInt(100000000000000000), gasLimit, gasPrice, txData)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		panic(err)
	}
	err = ethSdk.SendRawTransaction(signedTx)
	if err != nil {
		panic(err)
	}
	waitTransactionConfirm(ethSdk, signedTx.Hash())
}
