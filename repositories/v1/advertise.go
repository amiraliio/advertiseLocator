package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/amiraliio/advertiselocator/helpers"
	"github.com/amiraliio/advertiselocator/models"
	"go.mongodb.org/mongo-driver/bson"
)

type AdvertiseInterface interface {
	InsertAdvertise(advertise *models.Advertise) (*models.Advertise, error)
	ListOfAdvertise(filter *models.AdvertiseFilter) ([]*models.Advertise, error)
	FindOne(filter *models.AdvertiseFilter) (advertise *models.Advertise, err error)
	DeleteOne(filter *models.AdvertiseFilter) (int64, error)
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

func (service *AdvertiseRepository) FindOne(filter *models.AdvertiseFilter) (advertise *models.Advertise, err error) {
	query := bson.M{"_id": filter.ID, "person._id": filter.UserID}
	if err = helpers.Mongo().FindOne(models.AdvertiseCollection, query).Decode(&advertise); err != nil {
		return nil, err
	}
	return advertise, nil
}

func (service *AdvertiseRepository) DeleteOne(filter *models.AdvertiseFilter) (int64, error) {
	query := bson.M{"_id": filter.ID, "person._id": filter.UserID}
	result, err := helpers.Mongo().DeleteOne(models.AdvertiseCollection, query)
	if err != nil {
		return 0, err
	}
	if result.DeletedCount == 0 {
		return 0, errors.New("Document for deleting doesn't exist")
	}
	return result.DeletedCount, nil
}
