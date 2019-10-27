package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)



//BaseFilter filter
type BaseFilter struct {
	FromDate primitive.DateTime `json:"fromDate" query:"fromDate" bson:"fromDate"`
	ToDate   primitive.DateTime `json:"toDate" query:"toDate" bson:"toDate"`
}



//TagFilter filter
type TagFilter struct {
	BaseFilter ",inline"
	Key        string  `json:"key" query:"key" bson:"key"`
	Value      string  `json:"value" query:"value" bson:"value"`
	Min        float64 `json:"min" query:"min" bson:"min"`
	Max        float64 `json:"max" query:"max" bson:"max"`
}