package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	Description string             `json:"description" bson:"description"`
	ExpireDate  primitive.DateTime `json:"expireDate" bson:"expireDate"`
	Charges     []*Charge          `json:"charges" bson:"charges"`
	Payments    []*Payment         `json:"payments" bson:"payments"`
	Visibility  string             `json:"visibility" bson:"visibility"`
}

//Tag model
type Tag struct {
	Key   string `json:"key" bson:"key"`
	Value string `json:"value" bson:"value"`
	Min   string `json:"min" bson:"min"`
	Max   string `json:"max" bson:"max"`
}
