package repositories

import (
	"context"
	"github.com/amiraliio/advertiselocator/helpers"
	"github.com/amiraliio/advertiselocator/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type AdvertiseInterface interface {
	InsertAdvertise(advertise *models.Advertise) (*models.Advertise, error)
	ListOfAdvertise(filter *models.AdvertiseFilter) ([]*models.Advertise, error)
}

type AdvertiseRepository struct{}

func (service *AdvertiseRepository) InsertAdvertise(advertise *models.Advertise) (*models.Advertise, error) {
	_, err := helpers.Mongo().InsertOne(models.AdvertiseCollection, advertise)
	if err != nil {
		return nil, err
	}
	return advertise, nil
}

func (service *AdvertiseRepository) ListOfAdvertise(filter *models.AdvertiseFilter) ([]*models.Advertise, error) {
	//TODO query with person id from token
	query := bson.D{
		bson.E{
			Key:   "person._id",
			Value: filter.UserID,
		},
	}
	cursor, err := helpers.Mongo().List(models.AdvertiseCollection, query)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer cursor.Close(ctx)
	var data []*models.Advertise
	for cursor.Next(ctx) {
		var advertise *models.Advertise
		if err := cursor.Decode(&advertise); err != nil {
			return nil, err
		}
		data = append(data, advertise)
	}
	if cursor.Err() != nil {
		return nil, cursor.Err()
	}
	return data, nil
}
