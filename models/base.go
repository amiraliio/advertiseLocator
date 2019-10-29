package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


//ActiveStatus status
const ActiveStatus string = "ACTIVE"

//InactiveStatus status
const InactiveStatus string = "INACTIVE"

//PendingStatus status
const PendingStatus string = "PENDING"

//DeleteStatus status
const DeleteStatus string = "DELETE"



//Base model
type Base struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Status    string             `json:"status" bson:"status"`
	CreatedAt primitive.DateTime `json:"createdAt" bson:"createdAt"`
	CreatedBy primitive.ObjectID `json:"createdBy" bson:"createdBy"`
	UpdatedAt primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
	UpdatedBy primitive.ObjectID `json:"updatedBy" bson:"updatedBy"`
}
