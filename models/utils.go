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

func (bigInt *BigInt) MarshalJSON() ([]byte, error) {
	if bigInt == nil {
		return []byte("null"), nil
	}
	return []byte(bigInt.String()), nil
}

func (bigInt *BigInt) UnmarshalJSON(p []byte) error {
	if string(p) == "null" {
		return nil
	}
	data, ok := new(big.Int).SetString(string(p), 10)
	if !ok {
		return fmt.Errorf("not a valid big integer: %s", p)
	}
	bigInt.Int = *data
	return nil
}

