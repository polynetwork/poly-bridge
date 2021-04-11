package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString2BigInt(t *testing.T) {
	expect := "123456789876543212345678987654321"

	data := String2BigInt(expect).String()
	assert.Equal(t, expect, data)
}
