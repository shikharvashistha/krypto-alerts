package models

import (
	"gorm.io/gorm"
)

type Coins struct {
	Name  string  `json:"name"`  // Name is the name of the coin
	Price float64 `json:"price"` // Price is the price of the coin
}

func (d *Coins) Get(db *gorm.DB) error {

	return db.Where(d).First(d).Error
}

func (d *Coins) Create(db *gorm.DB) error {
	return db.Session(&gorm.Session{FullSaveAssociations: true}).Create(d).Error
}

func (d *Coins) Delete(db *gorm.DB) error {
	return db.Where(&Coins{Name: d.Name}).Delete(d).Error
}

func (d *Coins) Update(db *gorm.DB) error {
	return db.Session(&gorm.Session{FullSaveAssociations: true}).Where(&Coins{Name: d.Name}).Updates(d).Error
}

func (d *Coins) List(db *gorm.DB) ([]Coins, error) {

	Coinss := []Coins{}
	err := db.Where(d).Find(&Coinss).Error

	if err != nil {
		return nil, err
	}
	return Coinss, nil
}
