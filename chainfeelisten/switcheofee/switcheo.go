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

package switcheofee

import (
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
)

type SwitcheoFee struct {
	swthCfg *conf.FeeListenConfig
	swthSdk *chainsdk.SwitcheoSdkPro
}

func NewSwitcheoFee(swthCfg *conf.FeeListenConfig, feeUpdateSlot int64) *SwitcheoFee {
	switcheoFee := &SwitcheoFee{}
	switcheoFee.swthCfg = swthCfg
	sdk := chainsdk.NewSwitcheoSdkPro(swthCfg.Nodes, uint64(feeUpdateSlot), swthCfg.ChainId)
	switcheoFee.swthSdk = sdk
	return switcheoFee
}

func (this *SwitcheoFee) GetFee() (*big.Int, *big.Int, *big.Int, error) {
	gasPrice := new(big.Int).SetUint64(0)
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(basedef.FEE_PRECISION))
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(this.swthCfg.GasLimit))
	proxyFee := new(big.Int).Mul(gasPrice, new(big.Int).SetInt64(this.swthCfg.ProxyFee))
	proxyFee = new(big.Int).Div(proxyFee, new(big.Int).SetInt64(100))
	minFee := new(big.Int).Mul(gasPrice, new(big.Int).SetInt64(this.swthCfg.MinFee))
	minFee = new(big.Int).Div(minFee, new(big.Int).SetInt64(100))
	return minFee, gasPrice, proxyFee, nil
}

func (this *SwitcheoFee) GetChainId() uint64 {
	return this.swthCfg.ChainId
}

func (this *SwitcheoFee) Name() string {
	return this.swthCfg.ChainName
}
