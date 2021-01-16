package lru

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLRU(t *testing.T) {
	lru := New(2)
	err := lru.Put("k1", "v1")
	assert.Equal(t, err, nil)

	err = lru.Put("k2", "v2")
	assert.Equal(t, err, nil)

	v, _ := lru.Get("k1")
	assert.Equal(t, v.(string), "v1")

	err = lru.Put("k3", "v3")
	assert.Equal(t, err, nil)

	v, _ = lru.Get("k2")
	assert.Equal(t, v, nil)
}