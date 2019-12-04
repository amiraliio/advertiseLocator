//Package repositories ...
package repositories

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"strings"
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
	var builder bson.D
	//tags mapper
	if filter.Tags != nil {
		var mapper bson.A
		for _, tag := range filter.Tags {
			if (tag.Min != "" && tag.Max != "" && tag.Min == tag.Max) || (tag.Value != "") {
				if tag.Value == "" {
					exactValue := bson.D{
						bson.E{
							Key: "$elemMatch",
							Value: bson.D{
								bson.E{Key: "key", Value: tag.Key},
								bson.E{Key: "value", Value: tag.Min},
							},
						},
					}
					mapper = append(mapper, exactValue)
				} else {
					exactValue := bson.D{
						bson.E{
							Key: "$elemMatch",
							Value: bson.D{
								bson.E{Key: "key", Value: tag.Key},
								bson.E{Key: "value", Value: tag.Value},
							},
						},
					}
					mapper = append(mapper, exactValue)
				}
			}
			if tag.Min != "" && tag.Value == "" {
				minValue, minDataType, err := helpers.CheckAndReturnNumeric(tag.Min)
				if err == nil {
					var minDataValue interface{}
					switch minDataType {
					case reflect.Int:
						minDataValue = minValue.(int)
					case reflect.Float64:
						minDataValue = minValue.(float64)
					}
					minValue := bson.D{
						bson.E{
							Key: "$elemMatch",
							Value: bson.D{
								bson.E{
									Key:   "key",
									Value: tag.Key,
								},
								bson.E{
									Key: "numericValue",
									Value: bson.D{
										bson.E{
											Key:   "$gte",
											Value: minDataValue,
										},
										bson.E{
											Key:   "$type",
											Value: "number",
										},
										bson.E{
											Key:   "$ne",
											Value: bson.TypeNull,
										},
									},
								},
							},
						},
					}
					mapper = append(mapper, minValue)
				}
			}
			if tag.Max != "" && tag.Value == "" {
				maxValue, maxDataType, err := helpers.CheckAndReturnNumeric(tag.Max)
				if err == nil {
					var maxDataValue interface{}
					switch maxDataType {
					case reflect.Int:
						maxDataValue = maxValue.(int)
					case reflect.Float64:
						maxDataValue = maxValue.(float64)
					}
					maxValue := bson.D{
						bson.E{
							Key: "$elemMatch",
							Value: bson.D{
								bson.E{
									Key:   "key",
									Value: tag.Key,
								},
								bson.E{
									Key: "numericValue",
									Value: bson.D{
										bson.E{
											Key:   "$lte",
											Value: maxDataValue,
										},
										bson.E{
											Key:   "$type",
											Value: "number",
										},
										bson.E{
											Key:   "$ne",
											Value: bson.TypeNull,
										},
									},
								},
							},
						},
					}
					mapper = append(mapper, maxValue)
				}
			}
		}
		//if request from auth user another query block will be added to final query builder
		if filter.UserID != primitive.NilObjectID {
			userValue := bson.E{
				Key:   "person._id",
				Value: filter.UserID,
			}
			builder = bson.D{userValue, bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$all", Value: mapper}}}}
		} else {
			builder = bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$all", Value: mapper}}}}
		}
	} else if filter.UserID != primitive.NilObjectID {
		//if request from auth user another query block will be added to final query builder
		builder = bson.D{
			bson.E{
				Key:   "person._id",
				Value: filter.UserID,
			},
		}
	} else {
		builder = bson.D{}
	}
	//perform query
	var option *options.FindOptions
	var skip int64
	if filter.Page > 0 {
		skip = (filter.Page - 1) * (filter.Limit)
	} else {
		skip = 0
	}
	if filter.Sort != "" {
		var sortKey string
		var sortValue int
		if strings.HasPrefix(filter.Sort, "-") {
			sortKey = strings.TrimPrefix(filter.Sort, "-")
			sortValue = -1
		} else {
			sortValue = 1
			sortKey = filter.Sort
		}
		option = options.Find().SetSkip(skip).SetLimit(filter.Limit).SetSort(bson.M{sortKey: sortValue})
	} else {
		option = options.Find().SetSkip(skip).SetLimit(filter.Limit)
	}
	cursor, err := helpers.Mongo().Find(models.AdvertiseCollection, builder, option)
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
	var query bson.M
	if filter.UserID != primitive.NilObjectID {
		query = bson.M{"_id": filter.ID, "person._id": filter.UserID}
	} else {
		query = bson.M{"_id": filter.ID}
	}
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
