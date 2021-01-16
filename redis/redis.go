package redis

import (
	"time"

	"github.com/go-redis/redis"
)

var Redis *redis.Client

func Init() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

}

func Set(key string, value interface{}, expireSecond int) error {
	_, err := Redis.Set(key, value, time.Duration(expireSecond)*time.Second).Result()
	return err
}

func Get(key string) (string, error) {
	return Redis.Get(key).Result()
}

func MGet(keys ...string) ([]interface{}, error) {
	return Redis.MGet(keys...).Result()
}
