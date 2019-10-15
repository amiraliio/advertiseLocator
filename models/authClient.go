package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AuthClient model
type AuthClient struct {
	BaseUser
	ClientID         primitive.ObjectID `json:"clientID"`
	Version          string             `json:"version"`
	LastLogin        primitive.DateTime `json:"lastLogin"`
	OSType           string             `json:"osType"`
	OSVersion        string             `json:"osVersion"`
	RefreshToken     string             `json:"OSVersion"`
	Token            string             `json:"token"`
	VerificationCode string             `json:"verificationCode"`
	ExpireDate       primitive.DateTime `json:"expireDate"`
	API              API                `json:"api"`
	Auth             Auth               `json:"auth"`
}

//API model
type API struct {
	Key        string             `json:"key"`
	Name       string             `json:"name"`
	ExpireDate primitive.DateTime `json:"expireDate"`
	Token      string             `json:"token"`
	Type       string             `json:"type"`
}
