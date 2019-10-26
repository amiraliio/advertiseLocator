package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//APIKeyCollection collection
const APIKeyCollection string = "apiKey"

//APIKeyHeaderName const
const APIKeyHeaderName string = "x-api-key"

//InternalAPIKey const
const InternalAPIKey string = "internal"

//ExternalAPIKey const
const ExternalAPIKey string = "external"

//WebAPIKey const
const WebAPIKey string = "web"

//AndroidAPIKey const
const AndroidAPIKey string = "android"

//IosAPIKey const
const IosAPIKey string = "ios"

//API model
type API struct {
	Base       ",inline"
	Key        string             `json:"key" bson:"key"`
	Name       string             `json:"name" bson:"name"`
	ExpireDate primitive.DateTime `json:"expireDate" bson:"expireDate"`
	Token      string             `json:"token" bson:"token"`
	Type       string             `json:"type" bson:"type"`
}
