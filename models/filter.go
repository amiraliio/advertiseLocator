package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//BaseFilter filter
type BaseFilter struct {
	StartDate primitive.DateTime `json:"startDate" query:"startdate" bson:"startdate"`
	EndDate   primitive.DateTime `json:"endDate" query:"endDate" bson:"endDate"`
	UserID    primitive.ObjectID `json:"userID" bson:"userID"`
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Sort      string             `json:"sort" query:"sort" bson:"sort"`
	LastIndex string             `json:"lastIndex" query:"lastIndex" bson:"lastIndex"`
	Count     int                `json:"count" query:"count" bson:"count"`
}

//AdvertiseFilter
type AdvertiseFilter struct {
	BaseFilter ",inline"
	Tags       []*Tag `json:"tags" query:"tags" bson:"tags"`
}
