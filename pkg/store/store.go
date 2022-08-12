package store

import (
	"github.com/shikharvashistha/krypto-alerts/pkg/store/keyvalue"
	"github.com/shikharvashistha/krypto-alerts/pkg/store/relational"
)

type store struct {
	rl relational.RL // relational store for Postgresql
	kv keyvalue.KV   // key value store for Redis
}

func (s *store) RL() relational.RL {
	return s.rl
}
func (s *store) KV() keyvalue.KV {
	return s.kv
}
