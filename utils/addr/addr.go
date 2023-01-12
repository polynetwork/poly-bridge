package addr

import (
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

func InSlice(a common.Address, b ...common.Address) bool {
	for _, v := range b {
		if strings.EqualFold(v.String(), a.String()) {
			return true
		}
	}
	return false
}
