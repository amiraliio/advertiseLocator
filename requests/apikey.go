package requests

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//APIKey request model
type APIKey struct {
	Type       string             `json:"type" bson:"type" validate:"required"`
	Name       string             `json:"name" bson:"name" validate:"required"`
	ExpireName primitive.DateTime `json:"expireTime" bson:"expireTime" validate:"required"`
}
