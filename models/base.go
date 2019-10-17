package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Base model
type Base struct {
	ID        primitive.ObjectID `json:"_id"`
	Status    string             `json:"status"`
	CreatedAt primitive.DateTime `json:"createdAt"`
	CreatedBy primitive.ObjectID `json:"createdBy"`
	UpdatedAt primitive.DateTime `json:"updatedAt"`
	UpdatedBy primitive.ObjectID `json:"updatedBy"`
}
