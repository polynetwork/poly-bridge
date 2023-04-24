package number

import (
	"fmt"
	"math/big"
)

func BigIntDiv10X(a *big.Int, x int) string {
	nStr := a.String()
	if nStr == "0" || x <= 0 {
		return nStr
	}
	s := fmt.Sprintf("%0*s", x+1, nStr)
	resultStr := s[:len(s)-x] + "." + s[len(s)-x:]
	return resultStr
}
