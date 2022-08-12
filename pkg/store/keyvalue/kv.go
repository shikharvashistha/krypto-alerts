package keyvalue

import (
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/shikharvashistha/krypto-alerts/pkg/utils"
)

type kvs struct{}

func (k *kvs) Set(key, value string, time time.Duration) error {
	rdb, ctx := utils.RedisConnect()            // Initialize the redis client or return the existing client
	err := rdb.Set(ctx, key, value, time).Err() // Set the key value pair in the redis server
	return err
}

func (k *kvs) Get(key string) (string, error) {
	rdb, ctx := utils.RedisConnect()       // Initialize the redis client or return the existing client
	val, err := rdb.Get(ctx, key).Result() // Get the value for the key from the redis server
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

func (k *kvs) Remove(key string) error {
	rdb, ctx := utils.RedisConnect() // Initialize the redis client or return the existing client
	err := rdb.Del(ctx, key).Err()   // Remove the key from the redis server
	return err
}
