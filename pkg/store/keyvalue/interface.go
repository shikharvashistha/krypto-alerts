package keyvalue

import "time"

type KV interface {
	Get(key string) (string, error)                  // Get implements KV to get a value from the store
	Set(key, value string, time time.Duration) error // Set implements KV to set a value in the store
	Remove(key string) error                         // Remove implements KV to remove a value from the store
}

func NewKVStore() KV {
	return &kvs{}
}
