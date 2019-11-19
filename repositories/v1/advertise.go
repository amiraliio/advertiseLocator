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
	//TODO move this query builder and use fluent api pattern
	//build query
	//TODO change this to raw query because is more clear
	var query bson.D
	for _, tag := range filter.Tags {
		if (tag.Min != "" && tag.Max != "" && tag.Min == tag.Max) || (tag.Value != "") {
			exactQueryTagKey := bson.E{
				Key:   "tags.key",
				Value: tag.Key,
			}
			var exactQueryTagValue bson.E
			if tag.Value == "" {
				exactQueryTagValue = bson.E{
					Key:   "tags.value",
					Value: tag.Min,
				}
			} else {
				exactQueryTagValue = bson.E{
					Key:   "tags.value",
					Value: tag.Value,
				}
			}
			// jjj := bson.D{exactQueryTagKey, exactQueryTagValue}
			query = append(query, exactQueryTagKey, exactQueryTagValue)
			continue
		}
		if tag.Min != "" {
			minQueryTagKey := bson.E{
				Key:   "tags.key",
				Value: tag.Key,
			}
			minQueryTagValue := bson.E{
				Key: "tags.value",
				Value: bson.E{
					Key:   "$gte",
					Value: tag.Value,
				},
			}
			query = append(query, minQueryTagKey, minQueryTagValue)
			continue
		}
		if tag.Max != "" {
			maxQueryTagKey := bson.E{
				Key:   "tags.key",
				Value: tag.Key,
			}
			maxQueryTagValue := bson.E{
				Key: "tags.value",
				Value: bson.E{
					Key:   "$lte",
					Value: tag.Value,
				},
			}
			query = append(query, maxQueryTagKey, maxQueryTagValue)
			continue
		}
	}
	//if request from auth user another query block will be added to final query builder
	if filter.UserID != primitive.NilObjectID {
		userQuery := bson.E{
			Key:   "person._id",
			Value: filter.UserID,
		}
		query = append(query, userQuery)
	}
	// skip := bson.E{
	// 	Key:   "$skip",
	// 	Value: filter.Page * filter.Limit,
	// }
	// limit := bson.E{
	// 	Key:   "$limit",
	// 	Value: filter.Limit,
	// }
	// queryBuilder = append(queryBuilder, skip, limit)

	//perform query
	cursor, err := helpers.Mongo().Find(models.AdvertiseCollection, query)
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
