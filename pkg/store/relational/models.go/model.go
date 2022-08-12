package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrRequiredFieldNotPresent = errors.New("required field not present")
)

func RegisterSchema(db *gorm.DB) error {
	return db.AutoMigrate(
		&Websocket{}, // Register Websocket table
		&Alert{},     // Register Alert table
		&Coins{},     // Register Coins table
		&Cache{},     // Register Cache table
	)
}

type Model struct {
	UserID    string    `json:"user_id"`              // UserID is the user id of the model
	CreatedAt time.Time `json:"created_at,omitempty"` // CreatedAt is the created at of the model
	UpdatedAt time.Time `json:"updated_at,omitempty"` // UpdatedAt is the updated at of the model
}
