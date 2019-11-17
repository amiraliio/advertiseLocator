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
//TODO add variable in return types and remove : from variables

//Mongo build
func Mongo() MongoInterface {
	return new(mongoService)
}

//MongoInterface interface
type MongoInterface interface {
	InsertOne(collectionName string, object interface{}) (primitive.ObjectID, error)
	FindOne(collectionName string, query bson.M) *mongo.SingleResult
	Find(collectionName string, query bson.D) (*mongo.Cursor, error)
	Aggregate(collectionName string, query bson.D) (*mongo.Cursor, error)
	FindOneAndUpdate(collectionName string, filter bson.D, update bson.D) *mongo.SingleResult
	DeleteOne(collectionName string, filter bson.M) (deleteResult *mongo.DeleteResult, err error)
}

type mongoService struct{}

//FindOne helper
func (service *mongoService) FindOne(collectionName string, query bson.M) *mongo.SingleResult {
	db := configs.DB().Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()
	return db.FindOne(ctx, query)
}

//Find helper
func (service *mongoService) Find(collectionName string, query bson.D) (*mongo.Cursor, error) {
	db := configs.DB().Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := db.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	return cursor, nil
}

func (service *mongoService) Aggregate(colllectionName string, query bson.D) (cursor *mongo.Cursor, err error) {
	db := configs.DB().Collection(colllectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err = db.Aggregate(ctx, query)
	if err != nil {
		return nil, err
	}
	return cursor, nil
}

//InsertOne document in mongo
func (service *mongoService) InsertOne(collectionName string, object interface{}) (primitive.ObjectID, error) {
	db := configs.DB().Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	entityModel, err := Flatten(object)
	if err != nil {
		return primitive.NilObjectID, err
	}
	result, err := db.InsertOne(ctx, entityModel)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

//FindOneAndUpdate helper
func (service *mongoService) FindOneAndUpdate(collectionName string, filter bson.D, update bson.D) *mongo.SingleResult {
	collection := configs.DB().Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()
	return collection.FindOneAndUpdate(ctx, filter, update)
}

func (service *mongoService) DeleteOne(collectionName string, filter bson.M) (deleteResult *mongo.DeleteResult, err error) {
	collection := configs.DB().Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	deleteResult, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	return deleteResult, nil
}
