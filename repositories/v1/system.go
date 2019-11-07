package repositories

import (
	"github.com/amiraliio/advertiselocator/helpers"
	"github.com/amiraliio/advertiselocator/models"
)

//SystemInterface service
type SystemInterface interface {
	CreateAPIKey(api *models.API) (*models.API, error)
}

//SystemRepository service
type SystemRepository struct{}

//CreateAPIKey service
func (systemService *SystemRepository) CreateAPIKey(api *models.API) (*models.API, error) {
	_, err := helpers.Mongo().InsertOne(models.APIKeyCollection, api)
	if err != nil {
		return nil, err
	}
	return api, nil
}
