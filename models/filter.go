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
	Page      int                `json:"page" query:"page"`
	Limit     int                `json:"limit" query:"limit"`
}

//TagFilter filter
type TagFilter struct {
	BaseFilter ",inline"
	Key        string  `json:"key" query:"key" bson:"key"`
	Value      string  `json:"value" query:"value" bson:"value"`
	Min        float64 `json:"min" query:"min" bson:"min"`
	Max        float64 `json:"max" query:"max" bson:"max"`
}

//AdvertiseFilter
type AdvertiseFilter struct {
	BaseFilter ",inline"
	Tags       []*Tag `json:"tags" query:"tags" bson:"tags"`
}
