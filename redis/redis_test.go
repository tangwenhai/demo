package redis

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRedis(t *testing.T) {
	Init()
	_ = Set("k1", "v1", 2)
	time.Sleep(time.Duration(1) * time.Second)
	v, err := Get("k1")
	assert.Equal(t, v, "v1")
	assert.Equal(t, err, nil)

	_ = Set("k2", "v2", 2)
	time.Sleep(time.Duration(3) * time.Second)
	v, err = Get("k2")
	assert.Equal(t, v, "")
	assert.NotEqual(t, err, nil)

	vals, err := MGet("k3", "k4")
	fmt.Printf("vals: %v, err: %v\n", vals, err)

	for _, value := range vals {
		str, ok := value.(string)
		fmt.Printf("str:%v, ok:%v, len(str):%v\n", str, ok, len(str))
		if !ok || len(str) <= 0 {
			continue
		}

	}
}
