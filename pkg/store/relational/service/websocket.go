package service

import (
	"github.com/shikharvashistha/krypto-alerts/pkg/store/relational/models.go"
	"gorm.io/gorm"
)

type Websocket interface {
	Create(model *models.Websocket) error                     // Create takes a model and creates it in the database for Websocket table
	Delete(model *models.Websocket) error                     // Delete takes a model and deletes it from the database for Websocket table
	Get(model *models.Websocket) error                        // Get takes a model and gets it from the database for Websocket table
	Update(model *models.Websocket) error                     // Update takes a model and updates it in the database for Websocket table
	List(model *models.Websocket) ([]models.Websocket, error) // List takes a model and returns a list of models from the database for Websocket table
}

type websocket struct {
	db *gorm.DB
}

func NewWebsocketSvc(db *gorm.DB) Websocket {
	return &websocket{db}
}

func (dep *websocket) Create(model *models.Websocket) error {
	return model.Create(dep.db)
}

func (dep *websocket) Delete(model *models.Websocket) error {
	return model.Delete(dep.db)
}

func (dep *websocket) Get(model *models.Websocket) error {
	return model.Get(dep.db)
}

func (dep *websocket) Update(model *models.Websocket) error {
	return model.Update(dep.db)
}

func (dep *websocket) List(model *models.Websocket) ([]models.Websocket, error) {
	return model.List(dep.db)
}
