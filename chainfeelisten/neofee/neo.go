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

package neofee

import (
	"math/big"
	"poly-swap/chainsdk"
	"poly-swap/conf"
)

type NeoFee struct {
	neoCfg *conf.FeeListenConfig
	neoSdk *chainsdk.NeoSdkPro
}

func NewNeoFee(neoCfg *conf.FeeListenConfig, feeUpdateSlot int64) *NeoFee {
	neoFee := &NeoFee{}
	neoFee.neoCfg = neoCfg
	//
	urls := neoCfg.GetNodesUrl()
	sdk := chainsdk.NewNeoSdkPro(urls, uint64(feeUpdateSlot), neoCfg.ChainId)
	neoFee.neoSdk = sdk
	return neoFee
}

func (this *NeoFee) GetFee() (*big.Int, *big.Int, *big.Int, error) {
	return big.NewInt(1000000000), big.NewInt(1000000000), big.NewInt(1000000000), nil
}

func (this *NeoFee) GetChainId() uint64 {
	return this.neoCfg.ChainId
}

func (this *NeoFee) Name() string {
	return conf.ChainId2Name(this.neoCfg.ChainId)
}
