package redis

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	Init()
	Set("k1", "v1", 2)
	time.Sleep(time.Duration(1)*time.Second)
	v, err := Get("k1")
	assert.Equal(t, v, "v1")
	assert.Equal(t, err, nil)

	Set("k2", "v2", 2)
	time.Sleep(time.Duration(3)*time.Second)
	v, err = Get("k2")
	assert.Equal(t, v, "")
	assert.NotEqual(t, err, nil)
}