package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//auth client model

//ClientCollection const
const ClientCollection string = "client"

//Client model
type Client struct {
	BaseUser         ",inline"
	ClientID         string             `json:"clientID" bson:"clientID"`
	Version          string             `json:"Version" bson:"version"`
	LastLogin        primitive.DateTime `json:"lastLogin" bson:"lastLogin"`
	OSType           string             `json:"osType" bson:"osType"`
	OSVersion        string             `json:"osVersion" bson:"osVersion"`
	RefreshToken     string             `json:"refreshToken" bson:"refreshToken"`
	Token            string             `json:"token" bson:"token"`
	VerificationCode string             `json:"verificationCode" bson:"verificationCode"`
	ExpireDate       primitive.DateTime `json:"expireDate" bson:"expireDate"`
	API              API                `json:"api" bson:"api"`
	Auth             Auth               `json:"auth" bson:"auth"`
}
