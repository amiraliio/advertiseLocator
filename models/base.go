package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//status of entities
const (
	ActiveStatus   string = "ACTIVE"
	InactiveStatus string = "INACTIVE"
	PendingStatus  string = "PENDING"
	DeleteStatus   string = "DELETE"
)

//Base model
type Base struct {
	Status    string             `json:"status" bson:"status"`
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedAt primitive.DateTime `json:"createdAt" bson:"createdAt"`
	CreatedBy primitive.ObjectID `json:"createdBy" bson:"createdBy"`
	UpdatedAt primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
	UpdatedBy primitive.ObjectID `json:"updatedBy" bson:"updatedBy"`
}
