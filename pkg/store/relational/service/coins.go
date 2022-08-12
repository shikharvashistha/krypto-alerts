package service

import (
	"github.com/shikharvashistha/krypto-alerts/pkg/store/relational/models.go"
	"gorm.io/gorm"
)

type Coins interface {
	Create(model *models.Coins) error                 // Create takes a model and creates it in the database for Coin table
	Delete(model *models.Coins) error                 // Delete takes a model and deletes it from the database for Coin table
	Get(model *models.Coins) error                    // Get takes a model and gets it from the database for Coin table
	Update(model *models.Coins) error                 // Update takes a model and updates it in the database for Coin table
	List(model *models.Coins) ([]models.Coins, error) // List takes a model and returns a list of models from the database for Coin table
}
type coins struct {
	db *gorm.DB
}

func NewCoinsSvc(db *gorm.DB) Coins {
	return &coins{db}
}

func (conf *coins) Create(model *models.Coins) error {
	return model.Create(conf.db)
}

func (c *coins) Delete(model *models.Coins) error {
	return model.Delete(c.db)
}

func (c *coins) Get(model *models.Coins) error {
	return model.Get(c.db)
}

func (c *coins) Update(model *models.Coins) error {
	return model.Update(c.db)
}

func (c *coins) List(model *models.Coins) ([]models.Coins, error) {
	return model.List(c.db)
}
