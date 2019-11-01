package helpers

import (
	"context"
	"time"

	"github.com/amiraliio/advertiselocator/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//TODO improvement must be just interface not function
//TODO sort, filter, pagination for List
//TODO complete this helper

//Mongo build
func Mongo() MongoInterface {
	return new(mongoService)
}

//MongoInterface interface
type MongoInterface interface {
	InsertOne(collectionName string, object interface{}) (primitive.ObjectID, error)
	FindOne(collectionName string, query bson.M) *mongo.SingleResult
	List(collectionName string, query bson.D, modelToMap interface{}) ([]interface{}, error)
	FindOneAndUpdate(collectionName string, filter bson.D, update bson.D) *mongo.SingleResult
}

type mongoService struct{}

//FindOne helper
func (service *mongoService) FindOne(collectionName string, query bson.M) *mongo.SingleResult {
	db := configs.DB().Collection(collectionName)
	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return db.FindOne(context, query)
}

//List helper
func (service *mongoService) List(collectionName string, query bson.D, modelToMap interface{}) ([]interface{}, error) {
	db := configs.DB().Collection(collectionName)
	context, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := db.Find(context, query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context)
	var data []interface{}
	for cursor.Next(context) {
		if err := cursor.Decode(modelToMap); err != nil {
			return nil, err
		}
		data = append(data, modelToMap)
	}
	if cursor.Err() != nil {
		return nil, cursor.Err()
	}
	return data, nil
}

//InsertOne document in mongo
func (service *mongoService) InsertOne(collectionName string, object interface{}) (primitive.ObjectID, error) {
	db := configs.DB().Collection(collectionName)
	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	entityModel, err := Flatten(object)
	if err != nil {
		return primitive.NilObjectID, err
	}
	result, err := db.InsertOne(context, entityModel)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

//FindOneAndUpdate helper
func (service *mongoService) FindOneAndUpdate(collectionName string, filter bson.D, update bson.D) *mongo.SingleResult {
	collection := configs.DB().Collection(collectionName)
	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return collection.FindOneAndUpdate(context, filter, update)
}
