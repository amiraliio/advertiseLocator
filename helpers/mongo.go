package helpers

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"github.com/amiraliio/advertiselocator/configs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TODO improvement must be just interface not function

//Mongo build
func Mongo() MongoInterface {
	return new(mongoService)
}

//MongoInterface interface
type MongoInterface interface {
	InsertOne(collectionName string, object interface{}) (primitive.ObjectID, error)
	FindOne(collectionName string, query bson.M) (bson.M, error)
}

type mongoService struct{}

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

//FindOne helper
func (service *mongoService) FindOne(collectionName string, query bson.M) (bson.M, error) {
	db := configs.DB().Collection(collectionName)
	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor := db.FindOne(context, query)
	var result bson.M
	err := cursor.Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}


func (service *mongoService) List(collectionName string, query bson.D) (bson.D, error){
	db := configs.DB().Collection(collectionName)
	context, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := db.Find(context, query)
	if err !=nil{
		return nil, err
	}
	defer cursor.Close(context)
	for cursor.Next(context){
		var result bson.M
		if err := cursor.Decode(&result); err !=nil{
			return nil, err
		}

	}
	if cursor.Err() !=nil{
		return nil,cursor.Err()
	}
	return
}
