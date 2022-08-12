package models

import (
	"gorm.io/gorm"
)

type Websocket struct {
	C string `json:"c"` // C is the identifier used to unmarshal the websocket message
}

func (d *Websocket) Get(db *gorm.DB) error {

	return db.Where(d).First(d).Error
}

func (d *Websocket) Create(db *gorm.DB) error {
	return db.Session(&gorm.Session{FullSaveAssociations: true}).Create(d).Error
}

func (d *Websocket) Delete(db *gorm.DB) error {
	return db.Where(&Websocket{C: d.C}).Delete(d).Error
}

func (d *Websocket) Update(db *gorm.DB) error {
	return db.Session(&gorm.Session{FullSaveAssociations: true}).Where(&Websocket{C: d.C}).Updates(d).Error
}

func (d *Websocket) List(db *gorm.DB) ([]Websocket, error) {

	Websockets := []Websocket{}
	err := db.Where(d).Find(&Websockets).Error

	if err != nil {
		return nil, err
	}
	return Websockets, nil
}
