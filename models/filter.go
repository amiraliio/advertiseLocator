package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//BaseFilter filter
type BaseFilter struct {
	FromDate primitive.DateTime `json:"fromDate" query:"fromDate"`
	ToDate   primitive.DateTime `json:"toDate" query:"toDate"`
}

//TagFilter filter
type TagFilter struct {
	BaseFilter
	Key   string  `json:"key" query:"key"`
	Value string  `json:"value" query:"value"`
	Min   float64 `json:"min" query:"min"`
	Max   float64 `json:"max" query:"max"`
}
