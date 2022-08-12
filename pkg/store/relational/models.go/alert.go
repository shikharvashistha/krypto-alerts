package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Alert struct {
	Model
	AlertID     string  `json:"alert_id"`     // AlertID is the unique identifier for an alert
	Email       string  `json:"email"`        // Email is the email address of the user who created the alert
	TriggerMail bool    `json:"trigger_mail"` // TriggerMail is a boolean value that determines if the user wants to receive an email notification when the alert is triggered
	AlertValue  float64 `json:"alert_value"`  // AlertValue is the value that the user wants to be alerted when the price of the coin reaches
	Status      string  `json:"status"`       // Status is the status of the alert. It can be either active or inactive
}

func (c *Alert) Get(db *gorm.DB) error {

	return db.Preload(clause.Associations).Where(c).First(c).Error
}

func (c *Alert) Create(db *gorm.DB) error {
	return db.Session(&gorm.Session{FullSaveAssociations: true}).Create(c).Error
}

func (c *Alert) Delete(db *gorm.DB) error {
	return db.Where(&Alert{AlertID: c.AlertID}).Delete(c).Error
}

func (c *Alert) Update(db *gorm.DB) error {
	return db.Session(&gorm.Session{FullSaveAssociations: true}).Where(&Alert{AlertID: c.AlertID}).Updates(c).Error
}

func (c *Alert) List(db *gorm.DB) ([]Alert, error) {

	Alert := []Alert{}
	err := db.Where(c).Find(&Alert).Error

	if err != nil {
		return nil, err
	}
	return Alert, nil
}
func (c *Alert) GetByOffset(db *gorm.DB, offset int, limit int) ([]Alert, error) {

	Alert := []Alert{}
	err := db.Where(c).Offset(offset).Limit(limit).Find(&Alert).Error

	if err != nil {
		return nil, err
	}
	return Alert, nil
}
