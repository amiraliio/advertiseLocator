package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//BaseFilter filter
type BaseFilter struct {
	FromDate primitive.DateTime `json:"fromDate" query:"fromDate" bson:"fromDate"`
	ToDate   primitive.DateTime `json:"toDate" query:"toDate" bson:"toDate"`
}


