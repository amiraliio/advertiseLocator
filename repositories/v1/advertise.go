package repositories

import (
	"github.com/amiraliio/advertiselocator/helpers"
	"github.com/amiraliio/advertiselocator/models"
)

type AdvertiseInterface interface {
	InsertAdvertise(advertise *models.Advertise) (*models.Advertise, error)
}

type AdvertiseRepository struct{}

func (service *AdvertiseRepository) InsertAdvertise(advertise *models.Advertise) (*models.Advertise, error) {
	_, err := helpers.Mongo().InsertOne(models.AdvertiseCollection, advertise)
	if err != nil {
		return nil, err
	}
	return advertise, nil
}
