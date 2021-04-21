package controllers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckNumberString(t *testing.T) {
	testErr := fmt.Errorf("invalid number string")

	var testdata = []struct {
		data   string
		expect error
	}{
		{
			data:   "0x124sjf",
			expect: testErr,
		},
		{
			data:   "124sjf",
			expect: testErr,
		},
		{
			data:   "124 f3",
			expect: testErr,
		},
		{
			data:   "jf35",
			expect: testErr,
		},
		{
			data:   "12345678987654321123456789876543s211234567898765432112345678987654321",
			expect: testErr,
		},
		{
			data:   "",
			expect: fmt.Errorf("number string is empty"),
		},
		{
			data:   "00000",
			expect: nil,
		},
		{
			data:   "0",
			expect: nil,
		},
		{
			data:   "1234",
			expect: nil,
		},
		{
			data:   "12345678987654321123456789876543211234567898765432112345678987654321",
			expect: nil,
		},
	}

	for _, v := range testdata {
		num, err := checkNumString(v.data)
		if err != nil {
			assert.Equal(t, v.expect.Error(), err.Error())
		} else {
			assert.True(t, v.expect == nil)
			t.Logf("origin %s, after convert %s", v.data, num.String())
		}
	}
}
