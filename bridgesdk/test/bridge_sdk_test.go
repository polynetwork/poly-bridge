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

package test

import (
	"encoding/json"
	"fmt"
	"math/big"
	"poly-bridge/bridgesdk"
	"testing"
	"time"
)

func TestBridageSdk(t *testing.T) {
	sdk := bridgesdk.NewBridgeSdkPro([]string{"https://bridge.poly.network/testnet/v1/"}, 1)
	rsp, err := sdk.GetFee(79, 2, "0000000000000000000000000000000000000000", "155040625D7ae3e9caDA9a73E3E44f76D3Ed1409")
	if err != nil {
		panic(err)
	}
	rspJson, _ := json.Marshal(rsp)
	fmt.Printf("rsp: %s\n", string(rspJson))
}

func Diff(a *big.Float, b *big.Float) float32 {
	diff := new(big.Float).Sub(a, b)
	aaa := new(big.Float).Quo(diff, b)
	bbb, _ := aaa.Float32()
	return bbb
}

func TestBridageSdkStable(t *testing.T) {
	sdk := bridgesdk.NewBridgeSdkPro([]string{"https://bridge.poly.network/testnet/v1/"}, 10)
	avgAmount := new(big.Float)
	{
		rsp, err := sdk.GetFee(79, 2, "0000000000000000000000000000000000000000", "155040625D7ae3e9caDA9a73E3E44f76D3Ed1409")
		if err != nil {
			panic(err)
		}
		amount, result := new(big.Float).SetString(rsp.TokenAmount)
		if !result {
			panic("float error")
		}
		avgAmount = amount
	}

	for true {
		time.Sleep(time.Second * 5)
		rsp, err := sdk.GetFee(79, 2, "0000000000000000000000000000000000000000", "155040625D7ae3e9caDA9a73E3E44f76D3Ed1409")
		if err != nil {
			panic(err)
		}
		rspJson, _ := json.Marshal(rsp)
		fmt.Printf("rsp: %s\n", string(rspJson))
		amount, result := new(big.Float).SetString(rsp.TokenAmount)
		if !result {
			panic("float error")
		}
		diff := Diff(amount, avgAmount)
		if diff < -0.2 {
			panic(err)
		}
	}
}
