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

package neo3fee

import (
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
)

type Neo3Fee struct {
	neoCfg *conf.FeeListenConfig
	neoSdk *chainsdk.Neo3SdkPro
}

func NewNeo3Fee(neoCfg *conf.FeeListenConfig, feeUpdateSlot int64) *Neo3Fee {
	Neo3Fee := &Neo3Fee{}
	Neo3Fee.neoCfg = neoCfg
	//
	urls := neoCfg.GetNodesUrl()
	sdk := chainsdk.NewNeo3SdkPro(urls, uint64(feeUpdateSlot), neoCfg.ChainId)
	Neo3Fee.neoSdk = sdk
	return Neo3Fee
}

func (this *Neo3Fee) GetFee() (*big.Int, *big.Int, *big.Int, error) {
	gasPrice := new(big.Int).SetUint64(1)
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(basedef.FEE_PRECISION))
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(this.neoCfg.GasLimit))
	proxyFee := new(big.Int).Mul(gasPrice, new(big.Int).SetInt64(this.neoCfg.ProxyFee))
	proxyFee = new(big.Int).Div(proxyFee, new(big.Int).SetInt64(100))
	minFee := new(big.Int).Mul(gasPrice, new(big.Int).SetInt64(this.neoCfg.MinFee))
	minFee = new(big.Int).Div(minFee, new(big.Int).SetInt64(100))
	return minFee, gasPrice, proxyFee, nil
}

func (this *Neo3Fee) GetChainId() uint64 {
	return this.neoCfg.ChainId
}

func (this *Neo3Fee) Name() string {
	return this.neoCfg.ChainName
}
