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
)

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
