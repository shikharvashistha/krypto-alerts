package utils

import (
	"encoding/json"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/shikharvashistha/krypto-alerts/pkg/store/relational/models.go"
)

func InitializeWebsocketService() {
	// Initialize the websocket logger service
	logger := NewLogger("websocket")

	// Initialize a new redis client or return the existing client
	rdb, ctx := RedisConnect()

	// Initialize the websocket service
	connection, _, err := websocket.DefaultDialer.Dial("wss://stream.binance.com/ws/btcusdt@miniTicker", nil)
	if err != nil {
		logger.WithError(WebSocket, err).Info("Failed to initialize the websocket client")
		return
	}
	go func() {
		for {
			// Read the message from the websocket
			_, message, err := connection.ReadMessage()
			if err != nil {
				logger.WithError(WebSocket, err).Info("Failed to read the message from the websocket")
				return
			}
			var response models.Websocket
			// Unmarshal the message into the response struct
			err = json.Unmarshal(message, &response)
			if err != nil {
				logger.WithError(WebSocket, err).Info("Failed to unmarshal the message from the websocket")
				return
			}

			// Prepare the message to be sent to the redis server
			var alerts []models.Alert
			price, err := strconv.ParseFloat(response.C, 64)
			if err != nil {
				logger.WithError(WebSocket, err).Info("Failed to parse the price from the websocket")
				return
			}
			//Find the alerts that are active and have the price more than the current price
			if err := GetDB().Model(&alerts).Where("status = ? AND alert_value < ?", "active", price).Find(&alerts).Error; err != nil {
				logger.WithError(WebSocket, err).Info("Failed to find the alerts from the database")
				return
			}
			//For all the alerts, update the status to Inactive and send an email if the trigger mail is true
			for _, alert := range alerts {
				logger.WithField("alert", alert).Info("Alert triggered")
				alert.Status = "Sent"
				if err := GetDB().Save(&alert).Error; err != nil {
					logger.WithError(WebSocket, err).Info("Failed to update the alert in the database")
					return
				}
				if alert.Status != "Sent" {
					if err := rdb.Publish(ctx, "mail", alert.Email).Err(); err != nil {
						logger.WithError(WebSocket, err).Info("Failed to publish the alert to the redis channel")
						return
					}
				}
			}

		}
	}()

}
