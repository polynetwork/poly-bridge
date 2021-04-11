package ethereumlisten

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEthereumChainListen_IsContract(t *testing.T) {
	assert.False(t, isContract(""))
	assert.False(t, isContract("0000000000000000000000000000000000000000"))
	assert.False(t, isContract("111111111111111111111111111111111111111"))
	assert.True(t, isContract("1111111111111111111111111111111111111111"))
	assert.True(t, isContract("88aD4fD94a05602E595101a3e3171f91289C8f6b"))
	assert.True(t, isContract("9bEF1AE7304D3d2F344ea00e796ADa18cE1beb03"))

	x := []int{}

	var x1, x2 []int
	x1 = nil
	x2 = nil
	x = append(x, x1...)
	x = append(x, x2...)
	t.Log(x)
}

func TestNewEthereumChainListen_AllErr(t *testing.T) {
	y1 := errors.New("hi")
	var y2 error = nil
	assert.NoError(t, allErr(y1, y2))

	y2 = errors.New("mi")
	assert.NotNil(t, allErr(y1, y2))
}
