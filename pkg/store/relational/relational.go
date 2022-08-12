package relational

import (
	"github.com/shikharvashistha/krypto-alerts/pkg/store/relational/service"
	"gorm.io/gorm"
)

func NewRelational(db *gorm.DB) RL {
	return &relational{
		alert:     service.NewAlertSvc(db),     // Initialize the alert service
		websocket: service.NewWebsocketSvc(db), // Initialize the websocket service
		coins:     service.NewCoinsSvc(db),     // Initialize the coins service
		cache:     service.NewCacheSvc(db),     // Initialize the cache service
	}
}

type relational struct {
	alert     service.Alert
	websocket service.Websocket
	coins     service.Coins
	cache     service.Cache
}

func (r *relational) Alert() service.Alert {
	return r.alert
}

func (r *relational) Websocket() service.Websocket {
	return r.websocket
}

func (r *relational) Coins() service.Coins {
	return r.coins
}

func (r *relational) Cache() service.Cache {
	return r.cache
}
