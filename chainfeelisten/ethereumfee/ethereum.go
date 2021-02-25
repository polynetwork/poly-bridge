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

package ethereumfee

import (
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
)

type EthereumFee struct {
	ethCfg *conf.FeeListenConfig
	ethSdk *chainsdk.EthereumSdkPro
}

func NewEthereumFee(ethCfg *conf.FeeListenConfig, feeUpdateSlot int64) *EthereumFee {
	ethereumFee := &EthereumFee{}
	ethereumFee.ethCfg = ethCfg
	//
	urls := ethCfg.GetNodesUrl()
	sdk := chainsdk.NewEthereumSdkPro(urls, uint64(feeUpdateSlot), ethCfg.ChainId)
	ethereumFee.ethSdk = sdk
	return ethereumFee
}

func (this *EthereumFee) GetFee() (*big.Int, *big.Int, *big.Int, error) {
	gasPrice, err := this.ethSdk.SuggestGasPrice()
	if err != nil {
		return nil, nil, nil, err
	}
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(basedef.FEE_PRECISION))
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(this.ethCfg.GasLimit))
	proxyFee := new(big.Int).Mul(gasPrice, new(big.Int).SetInt64(this.ethCfg.ProxyFee))
	proxyFee = new(big.Int).Div(proxyFee, new(big.Int).SetInt64(100))
	minFee := new(big.Int).Mul(gasPrice, new(big.Int).SetInt64(this.ethCfg.MinFee))
	minFee = new(big.Int).Div(minFee, new(big.Int).SetInt64(100))
	return minFee, gasPrice, proxyFee, nil
}

func (this *EthereumFee) GetChainId() uint64 {
	return this.ethCfg.ChainId
}

func (this *EthereumFee) Name() string {
	return this.ethCfg.ChainName
}
