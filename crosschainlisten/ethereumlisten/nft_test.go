package ethereumlisten

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEthereumChainListen_IsContract(t *testing.T) {
	assert.False(t, isContract(""))
	assert.False(t, isContract("0000000000000000000000000000000000000000"))
	assert.False(t, isContract("111111111111111111111111111111111111111"))
	assert.True(t, isContract("1111111111111111111111111111111111111111"))
}
