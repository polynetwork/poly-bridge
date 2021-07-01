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

package models

import (
	"database/sql/driver"
	"fmt"
	"math/big"

	"poly-bridge/basedef"
	"poly-bridge/utils/decimal"
)

var chainsNamesCache = map[uint64]string{}

func Init(chains []*Chain) {
	for _, chain := range chains {
		chainsNamesCache[chain.ChainId] = chain.Name
	}
}

func ChainId2Name(id uint64) string {
	name, ok := chainsNamesCache[id]
	if ok {
		return name
	}
	return fmt.Sprintf("%v", id)
}

type BigInt struct {
	big.Int
}

func NewBigIntFromInt(value int64) *BigInt {
	x := new(big.Int).SetInt64(value)
	return NewBigInt(x)
}

func NewBigInt(value *big.Int) *BigInt {
	return &BigInt{Int: *value}
}

func (bigInt *BigInt) Value() (driver.Value, error) {
	if bigInt == nil {
		return "null", nil
	}
	return bigInt.String(), nil
}

func (bigInt *BigInt) Scan(v interface{}) error {
	value, ok := v.([]byte)
	if !ok {
		return fmt.Errorf("type error, %v", v)
	}
	if string(value) == "null" {
		return nil
	}
	data, ok := new(big.Int).SetString(string(value), 10)
	if !ok {
		return fmt.Errorf("not a valid big integer: %s", value)
	}
	bigInt.Int = *data
	return nil
}

func FormatAmount(precision uint64, amount *BigInt) string {
	precision_new := decimal.New(int64(precision), 0)
	amount_new := decimal.New(amount.Int64(), 0)
	return amount_new.Div(precision_new).String()
}

func FormatFee(chain uint64, fee *BigInt) string {
	if chain == basedef.BTC_CROSSCHAIN_ID {
		precision_new := decimal.New(int64(100000000), 0)
		fee_new := decimal.New(fee.Int64(), 0)
		return fee_new.Div(precision_new).String()
	} else if chain == basedef.ONT_CROSSCHAIN_ID {
		precision_new := decimal.New(int64(1000000000), 0)
		fee_new := decimal.New(fee.Int64(), 0)
		return fee_new.Div(precision_new).String()
	} else if chain == basedef.ETHEREUM_CROSSCHAIN_ID {
		precision_new := decimal.New(int64(1000000000000000000), 0)
		fee_new := decimal.New(fee.Int64(), 0)
		return fee_new.Div(precision_new).String()
	} else {
		precision_new := decimal.New(int64(1), 0)
		fee_new := decimal.New(fee.Int64(), 0)
		return fee_new.Div(precision_new).String()
	}
}

func TxType2Name(ty uint32) string {
	return "cross chain transfer"
}
func Precent(a uint64, b uint64) string {
	c := float64(a) / float64(b)
	return fmt.Sprintf("%.2f%%", c * 100)
}
