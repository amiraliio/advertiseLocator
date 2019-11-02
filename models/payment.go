package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Payment model
type Payment struct {
	Base          ",inline"
	RefID         string             `json:"refID" bson:"refID"`
	Bank          string             `json:"bank" bson:"bank"`
	Amount        string             `json:"amount" bson:"amount"`
	Type          string             `json:"type" bson:"type"`
	OutputAccount string             `json:"outputAccount" bson:"outputAccount"`
	InputAccount  string             `json:"InputAccount" bson:"inputAccount"`
	Date          primitive.DateTime `json:"date" bson:"date"`
}
