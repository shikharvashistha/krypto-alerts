package models

import (
	"gorm.io/gorm"
)

type Cache struct {
	Model
	PageNo   int    `json:"page_no"`   // PageNo is the page number of the cache
	PageSize int    `json:"page_size"` // PageSize is the page size of the cache
	Status   string `json:"status"`    // Status is the status of the alert
}

func (d *Cache) Get(db *gorm.DB) error {

	return db.Where(d).First(d).Error
}

func (d *Cache) Create(db *gorm.DB) error {
	return db.Session(&gorm.Session{FullSaveAssociations: true}).Create(d).Error
}

func (d *Cache) Delete(db *gorm.DB) error {
	return db.Where(&Cache{Model: Model{UserID: d.Model.UserID}}).Delete(d).Error
}

func (d *Cache) Update(db *gorm.DB) error {
	return db.Session(&gorm.Session{FullSaveAssociations: true}).Where(&Cache{Model: Model{UserID: d.Model.UserID}}).Updates(d).Error
}

func (d *Cache) List(db *gorm.DB) ([]Cache, error) {

	Caches := []Cache{}
	err := db.Where(d).Find(&Caches).Error

	if err != nil {
		return nil, err
	}
	return Caches, nil
}
