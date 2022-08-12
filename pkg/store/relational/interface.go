package relational

import "github.com/shikharvashistha/krypto-alerts/pkg/store/relational/service"

type RL interface {
	Alert() service.Alert         // Alert implementation for the relational store
	Cache() service.Cache         // Cache implementation for the relational store
	Coins() service.Coins         // Coins implementation for the relational store
	Websocket() service.Websocket // Websocket implementation for the relational store
}
