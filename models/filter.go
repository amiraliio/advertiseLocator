package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//BaseFilter filter
type BaseFilter struct {
	CreatedAt primitive.DateTime `json:"createdAt" query:"createdAt" bson:"createdAt"`
	UserID    primitive.ObjectID `json:"userID" bson:"userID"`
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Sort      string             `json:"sort" query:"sort" bson:"sort"`
	Page      uint16             `json:"page" query:"page"`
	Limit     uint16             `json:"limit" query:"limit"`
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
	Tags       []*Tag
}
