package store

import (
	"github.com/shikharvashistha/krypto-alerts/pkg/store/keyvalue"
	"github.com/shikharvashistha/krypto-alerts/pkg/store/relational"
	"gorm.io/gorm"
)

type Store interface {
	RL() relational.RL // Store interface to get the relational store
	KV() keyvalue.KV   // Store interface to get the key value store
}

func NewStore(db *gorm.DB) Store {
	return &store{
		rl: relational.NewRelational(db), // Initialize the relational store
		kv: keyvalue.NewKVStore(),        // Initialize the key value store
	}
}
