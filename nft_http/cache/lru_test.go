package cache

import (
	lru "github.com/hashicorp/golang-lru"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleRepeatSet(t *testing.T) {
	db, err := lru.NewARC(10)
	assert.NoError(t, err)

	type Simple struct {
		Name string
		Age int
	}

	data1 := &Simple{
		Name: "hi",
		Age:  10,
	}
	data2 := &Simple{
		Name: "hi",
		Age:  10,
	}
	db.Add(data1, data1)
	_, exist := db.Get(data2)
	assert.False(t, exist)
}
