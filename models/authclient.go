package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ClientCollection const
const ClientCollection string = "client"

//AuthClient model
type AuthClient struct {
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

//API model
type API struct {
	Key        string             `json:"key" bson:"key"`
	Name       string             `json:"name" bson:"name"`
	ExpireDate primitive.DateTime `json:"expireDate" bson:"expireDate"`
	Token      string             `json:"token" bson:"token"`
	Type       string             `json:"type" bson:"type"`
}
