//Package repositories ...
package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/amiraliio/advertiselocator/helpers"
	"github.com/amiraliio/advertiselocator/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (service *AdvertiseRepository) ListOfAdvertise(filter *models.AdvertiseFilter) (advertises []*models.Advertise, err error) {
	//TODO move this query builder and use fluent structure
	//build query
	var queryBuilder bson.D
	for _, tag := range filter.Tags {
		var tagFilter bson.E
		if tag.Min != "" {
			tagFilter = bson.E{
				Key: "$match",
				Value: bson.E{
					Key:   tag.Key,
					Value: tag.Value,
				},
			}
		}
		queryBuilder = append(queryBuilder, tagFilter)
	}
	//if request from auth user another query block will be added to final query builder
	if filter.UserID != primitive.NilObjectID {
		userQuery := bson.E{
			Key: "$match",
			Value: bson.E{
				Key:   "person._id",
				Value: filter.UserID,
			},
		}
		queryBuilder = append(queryBuilder, userQuery)
	}
	skip := bson.E{
		Key:   "$skip",
		Value: filter.Page * filter.Limit,
	}
	limit := bson.E{
		Key:   "$limit",
		Value: filter.Limit,
	}
	queryBuilder = append(queryBuilder, skip, limit)
	//perform query
	cursor, err := helpers.Mongo().Find(models.AdvertiseCollection, queryBuilder)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var advertise *models.Advertise
		if err = cursor.Decode(&advertise); err != nil {
			return nil, err
		}
		advertises = append(advertises, advertise)
	}
	if cursor.Err() != nil {
		return nil, cursor.Err()
	}
	return advertises, nil
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
		return 0, errors.New("document for deleting doesn't exist")
	}
	return result.DeletedCount, nil
}
