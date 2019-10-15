package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Payment model
type Payment struct {
	Base
	RefID         string             `json:"refID"`
	Bank          string             `json:"bank"`
	Amount        string             `json:"amount"`
	Type          string             `json:"type"`
	Date          primitive.DateTime `json:"date"`
	OutputAccount string             `json:"outputAccount"`
	InputAccount  string             `json:"InputAccount"`
}
