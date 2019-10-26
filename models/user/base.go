package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


//BaseUser model
type BaseUser struct {
	Base     ",inline"
	UserID   primitive.ObjectID `json:"userID" bson:"userID"`
	UserType string             `json:"userType" bson:"userType"`
	IP       string             `json:"ip" bson:"ip"`
}
