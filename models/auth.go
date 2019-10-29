package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

///////////// collection names //////////////////

//AuthCollection collection
const AuthCollection string = "auth"

//APIKeyCollection collection
const APIKeyCollection string = "apiKey"

//ClientCollection const
const ClientCollection string = "client"

///////////// header names //////////////////

//APIKeyHeaderKey const
const APIKeyHeaderKey string = "x-api-key"

//AuthorizationHeaderKey const
const AuthorizationHeaderKey string = "Authorization"

///////////// register and login ways //////////////////

//EmailAuthType auth type
const EmailAuthType string = "EMAIL"

//CellPhoneAuthType auth type
const CellPhoneAuthType string = "CELLPHONE"

//GoogleAuthType auth type
const GoogleAuthType string = "GOOGLE"

//FaceBookAuthType auth type
const FaceBookAuthType string = "FACEBOOK"

//TwitterAuthType auth type
const TwitterAuthType string = "TWITTER"

///////////// api key types //////////////////

//ExternalAPIKey const
const ExternalAPIKey string = "EXTERNAL"

//WebAPIKey const
const WebAPIKey string = "WEB"

//AndroidAPIKey const
const AndroidAPIKey string = "ANDROID"

//IosAPIKey const
const IosAPIKey string = "IOS"

///////////////////////////////////////////

//Auth model
type Auth struct {
	BaseUser ",inline"
	Value    string `json:"value" bson:"value"`
	Password string `json:"password" bson:"password"`
	Type     string `json:"type" bson:"type"`
}

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
	VerificationCode int                `json:"verificationCode" bson:"verificationCode"`
	ExpireDate       primitive.DateTime `json:"expireDate" bson:"expireDate"`
	API              API                `json:"api" bson:"api"`
	Auth             Auth               `json:"auth" bson:"auth"`
}

//API model
type API struct {
	Base        ",inline"
	Key         string             `json:"key" bson:"key"`
	Name        string             `json:"name" bson:"name"`
	ExpireDate  primitive.DateTime `json:"expireDate" bson:"expireDate"`
	Token       string             `json:"token" bson:"token"`
	Type        string             `json:"type" bson:"type"`
	Description string             `json:"description" bson:"description"`
}
