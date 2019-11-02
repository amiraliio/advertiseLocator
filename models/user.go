package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TODO order struct field from high to low field type size

//users collection names
const (
	PersonCollection string = "persons"
	AdminCollection  string = "admins"
)

//user types
const (
	PersonUserType string = "PERSON"
	AdminUserType  string = "ADMIN"
)

//BaseUser model
type BaseUser struct {
	Base     ",inline"
	UserType string             `json:"userType" bson:"userType"`
	IP       string             `json:"ip" bson:"ip"`
	UserID   primitive.ObjectID `json:"userID" bson:"userID"`
}

//Person model
type Person struct {
	BaseUser  ",inline"
	FirstName string    `json:"firstName" bson:"firstName"`
	LastName  string    `json:"lastName" bson:"lastName"`
	CellPhone string    `json:"cellPhone" bson:"cellPhone"`
	Email     string    `json:"email" bson:"email"`
	Radius    uint16    `json:"radius" bson:"radius"`
	Location  *Location `json:"location" bson:"location"`
	Avatar    *Image    `json:"avatar" bson:"avatar"`
}
