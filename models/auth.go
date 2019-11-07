package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//authentication collection
const (
	AuthCollection   string = "auth"
	APIKeyCollection string = "apiKey"
	ClientCollection string = "client"
)

//authentication headers
const (
	APIKeyHeaderKey        string = "x-api-key"
	AuthorizationHeaderKey string = "Authorization"
)

//authentication types
const (
	EmailAuthType     string = "EMAIL"
	CellPhoneAuthType string = "CELLPHONE"
	GoogleAuthType    string = "GOOGLE"
	FaceBookAuthType  string = "FACEBOOK"
	TwitterAuthType   string = "TWITTER"
)

//api key types
const (
	ExternalAPIKey string = "EXTERNAL"
	WebAPIKey      string = "WEB"
	AndroidAPIKey  string = "ANDROID"
	IosAPIKey      string = "IOS"
)

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
	API              *API               `json:"api" bson:"api"`
	Auth             *Auth              `json:"auth" bson:"auth"`
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
