package service

import (
	"github.com/shikharvashistha/krypto-alerts/pkg/store/relational/models.go"
	"gorm.io/gorm"
)

type Alert interface {
	Create(model *models.Alert) error                                               // Create takes a model and creates it in the database for Alert table
	Delete(model *models.Alert) error                                               // Delete takes a model and deletes it from the database for Alert table
	Get(model *models.Alert) error                                                  // Get takes a model and gets it from the database for Alert table
	Update(model *models.Alert) error                                               // Update takes a model and updates it in the database for Alert table
	List(model *models.Alert) ([]models.Alert, error)                               // List takes a model and returns a list of models from the database for Alert table
	GetByOffset(model *models.Alert, offset int, limit int) ([]models.Alert, error) // GetByOffset takes a model and returns a list of models from the databases for Alert table
}
type alert struct {
	db *gorm.DB
}

func NewAlertSvc(db *gorm.DB) Alert {
	return &alert{db}
}

func (conf *alert) Create(model *models.Alert) error {
	return model.Create(conf.db)
}

func (c *alert) Delete(model *models.Alert) error {
	return model.Delete(c.db)
}

func (c *alert) Get(model *models.Alert) error {
	return model.Get(c.db)
}

func (c *alert) Update(model *models.Alert) error {
	return model.Update(c.db)
}

func (c *alert) List(model *models.Alert) ([]models.Alert, error) {
	return model.List(c.db)
}

func (c *alert) GetByOffset(model *models.Alert, offset int, limit int) ([]models.Alert, error) {
	return model.GetByOffset(c.db, offset, limit)
}
