package repositories

import (
	"github.com/amiraliio/advertiselocator/configs"
	"github.com/amiraliio/advertiselocator/models"
)


type System interface {

}

type systemService struct{}


func (systemService *systemService) CreateAPIKey(api *models.API) (*models.API, error){
    db := configs.DB().Collection(models.APIKeyCollection, opts ...*options.CollectionOptions)
}

