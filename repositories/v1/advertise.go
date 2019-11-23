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

	var mapper bson.A
	for _, tag := range filter.Tags {
		if (tag.Min != "" && tag.Max != "" && tag.Min == tag.Max) || (tag.Value != "") {
			if tag.Value == "" {
				exactValue := bson.D{
					bson.E{Key: "tags.key", Value: tag.Key},
					bson.E{Key: "tags.value", Value: tag.Min},
				}
				mapper = append(mapper, exactValue)
				continue
			} else {
				exactValue := bson.D{
					bson.E{Key: "tags.key", Value: tag.Key},
					bson.E{Key: "tags.value", Value: tag.Value},
				}
				mapper = append(mapper, exactValue)
				continue
			}
		}
		if tag.Min != "" && tag.Value == "" {
			minValue := bson.D{
				bson.E{
					Key:   "tags.key",
					Value: tag.Key,
				},
				bson.E{
					Key: "tags.value",
					Value: bson.D{
						bson.E{
							Key:   "$gte",
							Value: tag.Min,
						},
					},
				},
			}
			mapper = append(mapper, minValue)
		}
		if tag.Max != "" && tag.Value == "" {
			maxValue := bson.D{
				bson.E{
					Key:   "tags.key",
					Value: tag.Key,
				},
				bson.E{
					Key: "tags.value",
					Value: bson.D{
						bson.E{
							Key:   "$lte",
							Value: tag.Max,
						},
					},
				},
			}
			mapper = append(mapper, maxValue)
		}
	}
	//if request from auth user another query block will be added to final query builder
	if filter.UserID != primitive.NilObjectID {
		userValue := bson.D{
			bson.E{
				Key:   "person._id",
				Value: filter.UserID,
			},
		}
		mapper = append(mapper, userValue)
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

	builder := bson.D{bson.E{Key: "$and", Value: mapper}}
	cursor, err := helpers.Mongo().Find(models.AdvertiseCollection, builder)
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
