package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//BaseFilter filter
type BaseFilter struct {
	FromDate primitive.DateTime `json:"fromDate"`
	ToDate   primitive.DateTime `json:"toDate"`
}

//TagFilter filter
type TagFilter struct {
	BaseFilter
	Key   string  `json:"key"`
	Value string  `json:"value"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
}
