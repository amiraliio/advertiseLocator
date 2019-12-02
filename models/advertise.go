package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TODO models validation

//AdvertiseCollection collection name
const AdvertiseCollection string = "advertises"

//advertise visibility
const (
	PublicVisibility  string = "PUBLIC"
	PrivateVisibility string = "PRIVATE"
)

//Advertise model
type Advertise struct {
	Base        ",inline"
	Location    *Location          `json:"location" bson:"location"`
	Tags        []*Tag             `json:"tags" bson:"tags"`
	Advertiser  *Person            `json:"person" bson:"person"`
	Radius      uint16             `json:"radius" bson:"radius"`
	Images      []*AdvertiseImage  `json:"images" bson:"images"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	ExpireDate  primitive.DateTime `json:"expireDate" bson:"expireDate"`
	Charges     []*Charge          `json:"charges" bson:"charges"`
	Payments    []*Payment         `json:"payments" bson:"payments"`
	Visibility  string             `json:"visibility" bson:"visibility"`
}

type Tag struct {
	Key          string      `json:"key" bson:"key" validate:"required,min=1,max=100"`
	Value        string      `json:"value" bson:"value" validate:"required"`
	Min          string      `json:"min" bson:"min" validate:"omitempty"`
	Max          string      `json:"max" bson:"max" validate:"omitempty"`
	NumericValue interface{} `json:"numericValue" bson:"numericValue" validate:"omitempty,numeric"`
}
