package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Base model
type Base struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Status    string             `json:"status" bson:"status"`
	CreatedAt primitive.DateTime `json:"createdAt" bson:"createdAt"`
	CreatedBy primitive.ObjectID `json:"createdBy" bson:"createdBy"`
	UpdatedAt primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
	UpdatedBy primitive.ObjectID `json:"updatedBy" bson:"updatedBy"`
}
