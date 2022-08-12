package service

import (
	"github.com/shikharvashistha/krypto-alerts/pkg/store/relational/models.go"
	"gorm.io/gorm"
)

type Cache interface {
	Create(model *models.Cache) error                 // Create takes a model and creates it in the database for Cache table
	Delete(model *models.Cache) error                 // Delete takes a model and deletes it from the database for Cache table
	Get(model *models.Cache) error                    // Get takes a model and gets it from the database for Cache table
	Update(model *models.Cache) error                 // Update takes a model and updates it in the database for Cache table
	List(model *models.Cache) ([]models.Cache, error) // List takes a model and returns a list of models from the database for Cache table
}
type cache struct {
	db *gorm.DB
}

func NewCacheSvc(db *gorm.DB) Cache {
	return &cache{db}
}

func (conf *cache) Create(model *models.Cache) error {
	return model.Create(conf.db)
}

func (c *cache) Delete(model *models.Cache) error {
	return model.Delete(c.db)
}

func (c *cache) Get(model *models.Cache) error {
	return model.Get(c.db)
}

func (c *cache) Update(model *models.Cache) error {
	return model.Update(c.db)
}

func (c *cache) List(model *models.Cache) ([]models.Cache, error) {
	return model.List(c.db)
}
