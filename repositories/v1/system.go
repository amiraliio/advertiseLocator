package repositories

import (
	"github.com/amiraliio/advertiselocator/helpers"
	"github.com/amiraliio/advertiselocator/models"
)

//SystemRepository service
type SystemRepository interface {
	CreateAPIKey(api *models.API) (*models.API, error)
}

//SystemService service
type SystemService struct{}

//CreateAPIKey service
func (systemService *SystemService) CreateAPIKey(api *models.API) (*models.API, error) {
	_, err := helpers.Mongo().InsertOne(models.APIKeyCollection, api)
	if err != nil {
		return nil, err
	}
	return api, nil
}
