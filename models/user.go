package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//BaseUser model
type BaseUser struct {
	Base
	UserID   primitive.ObjectID `json:"userID"`
	UserType string             `json:"userType"`
	IP       string             `json:"ip"`
}

//Person model
type Person struct {
	BaseUser
	Location  Location `json:"location"`
	Avatar    Image    `json:"avatar"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"LastName"`
	CellPhone string   `json:"cellPhone"`
	Email     string   `json:"email"`
	Radius    uint16   `json:"radius"`
}

//Admin model
type Admin struct {
	BaseUser
}
