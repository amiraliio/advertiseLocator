package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Charge model
type Charge struct {
	Base       ",inline"
	Title      string             `json:"title" bson:"title"`
	Type       string             `json:"type" bson:"type"`
	ExpireDate primitive.DateTime `json:"expireDate" bson:"expireDate"`
}
