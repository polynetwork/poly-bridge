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
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"poly-bridge/basedef"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumInfo struct {
	sdk          *EthereumSdk
	latestHeight uint64
}

func NewEthereumInfo(url string) *EthereumInfo {
	sdk, err := NewEthereumSdk(url)
	if err != nil || sdk == nil {
		panic(err)
	}
	return &EthereumInfo{
		sdk:          sdk,
		latestHeight: 0,
	}
}

type EthereumSdkPro struct {
	infos         map[string]*EthereumInfo
	selectionSlot uint64
	id            uint64
	mutex         sync.Mutex
}

func NewEthereumSdkPro(urls []string, slot uint64, id uint64) *EthereumSdkPro {
	infos := make(map[string]*EthereumInfo, len(urls))
	for _, url := range urls {
		infos[url] = NewEthereumInfo(url)
	}
	pro := &EthereumSdkPro{infos: infos, selectionSlot: slot, id: id}
	pro.selection()
	go pro.NodeSelection()
	return pro
}

func (pro *EthereumSdkPro) NodeSelection() {
	for {
		pro.nodeSelection()
	}
}

func (pro *EthereumSdkPro) nodeSelection() {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("node selection, recover info: %s", string(debug.Stack()))
		}
	}()
	logs.Debug("node selection of chain : %d......", pro.id)
	ticker := time.NewTicker(time.Second * time.Duration(pro.selectionSlot))
	for {
		select {
		case <-ticker.C:
			pro.selection()
		}
	}
}

func (pro *EthereumSdkPro) selection() {
	for url, info := range pro.infos {
		height, err := info.sdk.GetCurrentBlockHeight()
		if err != nil || height == math.MaxUint64 || height == 0 {
			logs.Error("nodeselection get current block height err, chain %v, url: %s", pro.id, url)
			height = 1
		}
		/*
			block, err := info.sdk.GetBlockByNumber(height)
			if err != nil || block == nil {
				logs.Error("get current block err: %v, url: %s", err, url)
				info.latestHeight = height - 1
				continue
			}
			transactions := block.Transactions()
			if len(transactions) > 0 {
				transaction := transactions[0]
				receipt, err := info.sdk.GetTransactionReceipt(transaction.Hash())
				if err != nil || receipt == nil {
					logs.Error("get transaction receipt err: %v, url: %s", err, url)
					info.latestHeight = height - 1
					continue
				}
			}
		*/
		pro.mutex.Lock()
		info.latestHeight = height - 1
		pro.mutex.Unlock()
	}
}

func (pro *EthereumSdkPro) GetLatest() *EthereumInfo {
	pro.mutex.Lock()
	defer func() {
		pro.mutex.Unlock()
	}()
	height := uint64(0)
	var latestInfo *EthereumInfo = nil
	for _, info := range pro.infos {
		if info != nil && info.latestHeight > height {
			height = info.latestHeight
			latestInfo = info
		}
	}
	return latestInfo
}

func (pro *EthereumSdkPro) GetClient() *ethclient.Client {
	info := pro.GetLatest()
	if info == nil {
		return nil
	}
	return info.sdk.GetClient()
}

func (pro *EthereumSdkPro) SetClientHeightZero(cli *ethclient.Client) {
	for node, info := range pro.infos {
		if info.sdk.GetClient() == cli {
			logs.Error("SetClientHeightZero node:%v is err", node)
			info.latestHeight = 0
			break
		}
	}
}

func (pro *EthereumSdkPro) GetLatestHeight() (uint64, error) {
	info := pro.GetLatest()
	if info == nil {
		return 0, fmt.Errorf("all node is not working")
	}
	return info.latestHeight, nil
}

func (pro *EthereumSdkPro) GetHeaderByNumber(number uint64) (*types.Header, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}
	flag := 0
	for info != nil {
		header, err := info.sdk.GetHeaderByNumber(number)
		if err != nil {
			flag++
			if flag > 3 {
				logs.Error("GetHeaderByNumber_chain:%v,node:%v,GetHeaderByNumber err", pro.id, info.sdk.url)
				flag = 0
				time.Sleep(time.Second)
			}
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			flag = 0
			return header, nil
		}
	}
	return nil, fmt.Errorf("all node is not working")
}

func (pro *EthereumSdkPro) GetBlockTimeByNumber(number uint64) (uint64, error) {
	info := pro.GetLatest()
	if info == nil {
		return 0, fmt.Errorf("all node is not working")
	}
	flag := 0
	for info != nil {
		timestamp, err := info.sdk.GetBlockTimeByNumber(pro.id, number)
		if err != nil {
			flag++
			if flag > 3 {
				logs.Error("GetBlockTimeByNumber_chain:%v,node:%v,eth_getBlockByNumber err %v", pro.id, info.sdk.url, err)
				flag = 0
				time.Sleep(time.Second)
			}
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			flag = 0
			return timestamp, nil
		}
	}
	return 0, fmt.Errorf("all node is not working")
}

func (pro *EthereumSdkPro) GetTransactionByHash(hash common.Hash) (*types.Transaction, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}
	flag := 0
	for info != nil {
		tx, err := info.sdk.GetTransactionByHash(hash)
		if err != nil {
			flag++
			if flag > 3 {
				logs.Error("chain:%v,node:%v,GetTransactionByHash err", pro.id, info.sdk.url)
				flag = 0
				time.Sleep(time.Second)
			}
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			flag = 0
			return tx, nil
		}
	}
	return nil, fmt.Errorf("all node is not working")
}

func (pro *EthereumSdkPro) GetTransactionReceipt(hash common.Hash) (*types.Receipt, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}
	flag := 0
	for info != nil {
		receipt, err := info.sdk.GetTransactionReceipt(hash)
		if err != nil {
			flag++
			if flag > 3 {
				logs.Error("chain:%v,node:%v,GetTransactionReceipt err", pro.id, info.sdk.url)
				flag = 0
				time.Sleep(time.Second)
			}
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			flag = 0
			return receipt, nil
		}
	}
	return nil, fmt.Errorf("all node is not working")
}

func (pro *EthereumSdkPro) NonceAt(addr common.Address) (uint64, error) {
	info := pro.GetLatest()
	if info == nil {
		return 0, fmt.Errorf("all node is not working")
	}

	for info != nil {
		nonce, err := info.sdk.NonceAt(addr)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return nonce, nil
		}
	}
	return 0, fmt.Errorf("all node is not working")
}

func (pro *EthereumSdkPro) SendRawTransaction(tx *types.Transaction) error {
	info := pro.GetLatest()
	if info == nil {
		return fmt.Errorf("all node is not working")
	}

	for info != nil {
		err := info.sdk.SendRawTransaction(tx)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return nil
		}
	}
	return fmt.Errorf("all node is not working")
}

func (pro *EthereumSdkPro) TransactionByHash(hash common.Hash) (*types.Transaction, bool, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, false, fmt.Errorf("all node is not working")
	}

	for info != nil {
		tx, isPending, err := info.sdk.TransactionByHash(hash)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return tx, isPending, nil
		}
	}
	return nil, false, fmt.Errorf("all node is not working")
}

func (pro *EthereumSdkPro) SuggestGasPrice() (*big.Int, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}

	for info != nil {
		gas, err := info.sdk.SuggestGasPrice()
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return gas, nil
		}
	}
	return nil, fmt.Errorf("all node is not working")
}

func (pro *EthereumSdkPro) EstimateGas(msg ethereum.CallMsg) (uint64, error) {
	info := pro.GetLatest()
	if info == nil {
		return 0, fmt.Errorf("all node is not working")
	}

	for info != nil {
		gas, err := info.sdk.EstimateGas(msg)
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return gas, nil
		}
	}
	return 0, fmt.Errorf("all node is not working")
}

//func (pro *EthereumSdkPro) Erc20Info(hash string) (string, string, int64, string, error) {
//	info := pro.GetLatest()
//	if info == nil {
//		return "", "", 0, "", fmt.Errorf("all node is not working")
//	}
//	for info != nil {
//		hash, name, decimal, symbol, err := info.sdk.Erc20Info(hash)
//		if err != nil {
//			info.latestHeight = 0
//			info = pro.GetLatest()
//		} else {
//			return hash, name, decimal, symbol, nil
//		}
//	}
//	return "", "", 0, "", fmt.Errorf("all node is not working")
//}

func (pro *EthereumSdkPro) IsEthAddress(addr string) bool {
	if addr == "0000000000000000000000000000000000000000" {
		return true
	} else {
		return false
	}
}

func (pro *EthereumSdkPro) Erc20Balance(erc20 string, addr string) (*big.Int, error) {
	info := pro.GetLatest()
	if info == nil {
		return new(big.Int).SetUint64(0), fmt.Errorf("all node is not working")
	}
	erc20Address := common.HexToAddress(erc20)
	userAddress := common.HexToAddress(addr)
	for info != nil {
		balance := new(big.Int).SetUint64(0)
		var err error
		if pro.IsEthAddress(erc20) {
			balance, err = info.sdk.EthBalance(addr)
			//logs.Error("eth, addr: %s, balance: %s", addr, balance.String())
		} else {
			balance, err = info.sdk.GetERC20Balance(erc20Address, userAddress)
			//logs.Error("erc20, addr: %s, balance: %s", addr, balance.String())
		}
		if err != nil {
			info.latestHeight = 0
			info = pro.GetLatest()
		} else {
			return balance, nil
		}
	}
	return new(big.Int).SetUint64(0), fmt.Errorf("all node is not working")
}
func (pro *EthereumSdkPro) Erc20TotalSupply(erc20 string) (*big.Int, error) {
	info := pro.GetLatest()
	if info == nil {
		return new(big.Int).SetUint64(0), fmt.Errorf("all node is not working")
	}
	erc20Address := common.HexToAddress(erc20)
	for info != nil {
		totalSupply := new(big.Int).SetUint64(0)
		if pro.IsEthAddress(erc20) {
			return totalSupply, nil
		} else {
			totalSupply, err := info.sdk.GetERC20TotalSupply(erc20Address)
			if err != nil {
				info.latestHeight = 0
				info = pro.GetLatest()
			} else {
				return totalSupply, nil
			}
		}
	}
	return new(big.Int).SetUint64(0), fmt.Errorf("all node is not working")
}

func (pro *EthereumSdkPro) WaitTransactionConfirm(hash common.Hash) bool {
	num := 0
	for num < 300 {
		time.Sleep(time.Second * 2)
		_, ispending, err := pro.TransactionByHash(hash)
		if err != nil {
			num++
			continue
		}
		if ispending {
			num++
			continue
		} else {
			return true
		}
	}
	return false
}

func (pro *EthereumSdkPro) NFTBalance(asset, owner common.Address) (balance *big.Int, err error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}

	for info != nil {
		if balance, err = info.sdk.GetNFTBalance(asset, owner); err != nil {
			info = pro.reset(info)
		} else {
			return
		}
	}
	return
}

func (pro *EthereumSdkPro) GetNFTOwner(asset string, tokenId *big.Int) (owner common.Address, err error) {
	assetAddr := common.HexToAddress(asset)
	info := pro.GetLatest()
	if info == nil {
		return EmptyAddress, fmt.Errorf("all node is not working")
	}

	for info != nil {
		if owner, err = info.sdk.GetNFTOwner(assetAddr, tokenId); err != nil {
			info = pro.reset(info)
		} else {
			return
		}
	}
	return
}

func (pro *EthereumSdkPro) GetTokensByIndex(
	inquirer, asset, owner common.Address,
	start, length int,
) (res map[string]string, err error) {

	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}

	for info != nil {
		if res, err = info.sdk.GetOwnerNFTsByIndex(inquirer, asset, owner, start, length); err != nil {
			info = pro.reset(info)
		} else {
			return
		}
	}
	return
}

func (pro *EthereumSdkPro) GetNFTUrl(asset common.Address, tokenId *big.Int) (url string, err error) {
	info := pro.GetLatest()
	if info == nil {
		return "", fmt.Errorf("all node is not working")
	}

	for info != nil {
		if url, err = info.sdk.GetNFTTokenUri(asset, tokenId); err != nil {
			info = pro.reset(info)
		} else {
			return
		}
	}
	return
}

func (pro *EthereumSdkPro) GetTokensById(
	inquirer, asset common.Address,
	tokenIdList []*big.Int,
) (res map[string]string, err error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}

	for info != nil {
		if res, err = info.sdk.GetNFTsById(inquirer, asset, tokenIdList); err != nil {
			info = pro.reset(info)
		} else {
			return
		}
	}
	return
}

func (pro *EthereumSdkPro) GetAndCheckTokenUrl(
	inquirer, asset, owner common.Address,
	tokenId *big.Int,
) (url string, err error) {

	info := pro.GetLatest()
	if info == nil {
		return "", fmt.Errorf("all node is not working")
	}

	for info != nil {
		if url, err = info.sdk.GetAndCheckNFTUrl(inquirer, asset, owner, tokenId); err != nil {
			info = pro.reset(info)
		} else {
			return
		}
	}
	return
}

func (pro *EthereumSdkPro) GetUnCrossChainNFTsByIndex(
	inquirer, asset common.Address,
	lockProxies []common.Address,
	start, length int,
) (mp map[string]string, err error) {

	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}
	for info != nil {
		if mp, err = info.sdk.GetUnCrossChainNFTsByIndex(inquirer, asset, lockProxies, start, length); err != nil {
			info = pro.reset(info)
		} else {
			return
		}
	}
	return
}

func (pro *EthereumSdkPro) GetNFTURLs(asset common.Address, tokenIds []*big.Int) (res map[string]string, err error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("all node is not working")
	}
	for info != nil {
		if res, err = info.sdk.GetOwnerNFTUrls(asset, tokenIds); err != nil {
			info = pro.reset(info)
		} else {
			return
		}
	}
	return
}

func (pro *EthereumSdkPro) reset(info *EthereumInfo) *EthereumInfo {
	info.latestHeight = 0
	return pro.GetLatest()
}

func (pro *EthereumSdkPro) GetBoundLockProxy(lockProxies []string, srcTokenHash, dstTokenHash string, chainId uint64) (string, error) {
	info := pro.GetLatest()
	dstTokenAddress := common.HexToAddress(dstTokenHash)

	if info != nil {
		for _, proxy := range lockProxies {
			proxyAddr := common.HexToAddress(proxy)
			boundAsset, err := info.sdk.GetBoundAssetHash(dstTokenAddress, proxyAddr, chainId)
			if err != nil || boundAsset == nil {
				logs.Info("GetBoundAssetHash err:%s", err)
				continue
			}
			if boundAsset == nil {
				continue
			}
			addrHash := (boundAsset.Hex())[2:]
			logs.Info("GetBoundAssetHash addrHash=%s", addrHash)

			if chainId == basedef.STARCOIN_CROSSCHAIN_ID {
				srcTokenHashByteString := strings.ToLower(hex.EncodeToString([]byte(srcTokenHash)))
				logs.Info("srcTokenHashByteString =%s", srcTokenHashByteString)
				if strings.Contains(srcTokenHashByteString, strings.ToLower(addrHash)) {
					return proxy, nil
				}
			} else if strings.EqualFold(addrHash, srcTokenHash) || strings.EqualFold(basedef.HexStringReverse(addrHash), srcTokenHash) {
				return proxy, nil
			}
		}
	}
	return "", fmt.Errorf("catnot get bounded asset hash of %s", dstTokenHash)
}

func (pro *EthereumSdkPro) FilterLog(FromBlock *big.Int, ToBlock *big.Int, Addresses []common.Address) ([]types.Log, error) {
	info := pro.GetLatest()
	if info == nil {
		return nil, fmt.Errorf("chain:%v FilterLog all node is not working", pro.id)
	}
	return info.sdk.FilterLog(FromBlock, ToBlock, Addresses)
}
