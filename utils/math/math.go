package math

import (
	"fmt"
	"math"
	"math/big"
	"poly-bridge/utils/decimal"

	emath "github.com/ethereum/go-ethereum/common/math"
)

var (
	_decimal int32 = 18
	N1       *big.Int
	_n1      decimal.Decimal

	EmptyDecimal = decimal.Zero
	EmptyBig     = big.NewInt(0)

	MinUint256 = big.NewInt(0)
	MaxUint256 = emath.MaxBig256

	MinDecimal256 = decimal.Zero
	MaxDecimal256 = decimal.NewFromBigInt(MaxUint256, 0)
)

func Init(precision int32) {
	_decimal = precision
	N1 = Pow10toBigInt(_decimal)
	_n1 = DecimalFromBigInt(N1)
}

func MultiT(n int) *big.Int {
	return SafeMul(N1, big.NewInt(int64(n)))
}

func MultiFloatT(f float64) decimal.Decimal {
	data := DecimalFromFloat(f)
	coin := DecimalFromBigInt(N1)
	value := DecimalSafeMul(data, coin)
	return value
}

func Mul1T(num decimal.Decimal) decimal.Decimal {
	return DecimalSafeMul(num, _n1)
}

func Div1T(num decimal.Decimal) decimal.Decimal {
	return DecimalUnSafeDiv(num, _n1)
}

func PrintUT(num *big.Int) uint64 {
	data := UnsafeDiv(num, N1)
	return data.Uint64()
}

func PrintFT(num decimal.Decimal) float64 {
	data := DecimalUnSafeDiv(num, _n1)
	value, _ := data.Float64()
	return value
}

func SafeUint32(bz []byte) uint32 {
	num := new(big.Int).SetBytes(bz)
	if num.Uint64() >= math.MaxUint32 {
		return math.MaxUint32
	}
	return uint32(num.Uint64())
}

func SafeUint8(bz []byte) uint8 {
	num := new(big.Int).SetBytes(bz)
	if num.Uint64() >= math.MaxUint8 {
		return math.MaxUint8
	}
	return uint8(num.Uint64())
}

// safe calculations for big int

func SafeAdd(a, b *big.Int) *big.Int {
	sum := new(big.Int).Add(a, b)
	if sum.Cmp(MaxUint256) >= 0 {
		return MaxUint256
	} else {
		return sum
	}
}

func SafeAddWithErr(a, b *big.Int) (*big.Int, error) {
	sum := new(big.Int).Add(a, b)
	if sum.Cmp(MaxUint256) > 0 {
		return nil, fmt.Errorf("out of range")
	}
	return sum, nil
}

func SafeMul(a, b *big.Int) *big.Int {
	sum := new(big.Int).Mul(a, b)
	if sum.Cmp(MaxUint256) >= 0 {
		return MaxUint256
	} else {
		return sum
	}
}

func UnSafeMod(a, b *big.Int) *big.Int {
	return new(big.Int).Mod(a, b)
}

func SafeSub(a, b *big.Int) *big.Int {
	if a.Cmp(b) <= 0 {
		return EmptyBig
	} else {
		return new(big.Int).Sub(a, b)
	}
}

func UnsafeSub(a, b *big.Int) (*big.Int, error) {
	if a.Cmp(b) < 0 {
		return nil, fmt.Errorf("sub amount invalid")
	}
	return new(big.Int).Sub(a, b), nil
}

func UnsafeDiv(a, b *big.Int) *big.Int {
	a1 := DecimalFromBigInt(a)
	b1 := DecimalFromBigInt(b)
	data := a1.Div(b1)
	return data.BigInt()
}

// safe calculations for decimal.Decimal

func DecimalZero() decimal.Decimal {
	return decimal.New(0, 1)
}

func DecimalFromInt64(i int64) decimal.Decimal {
	return decimal.NewFromInt(i)
}

func DecimalFromBigInt(i *big.Int) decimal.Decimal {
	return decimal.NewFromBigInt(i, 0)
}

func DecimalFromFloat(f float64) decimal.Decimal {
	return decimal.NewFromFloat(f)
}

func DecimalUnSafeDiv(a, b decimal.Decimal) decimal.Decimal {
	return a.Div(b)
}

func Decimal2BigInt(a decimal.Decimal) *big.Int {
	return a.BigInt()
}

func DecimalSafeAdd(a, b decimal.Decimal) decimal.Decimal {
	return a.Add(b)
}

func DecimalSafeSub(a, b decimal.Decimal) decimal.Decimal {
	if a.LessThanOrEqual(b) {
		return EmptyDecimal
	} else {
		return a.Sub(b)
	}
}

func DecimalSafeMul(a, b decimal.Decimal) decimal.Decimal {
	sum := a.Mul(b)
	if sum.Cmp(MaxDecimal256) > 0 {
		sum = MaxDecimal256
	}
	return sum
}

func Pow10toBigInt(n int32) *big.Int {
	data := decimal.NewFromBigInt(big.NewInt(1), n)
	return data.BigInt()
}

func String2BigInt(s string) *big.Int {
	data, _ := new(big.Int).SetString(s, 10)
	return data
}
