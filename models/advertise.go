package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Advertise model
type Advertise struct {
	Base
	Location    Location          `json:"location"`
	Tags        []Tag             `json:"tags"`
	Advertiser  Person            `json:"person"`
	Radius      uint16             `json:"radius"`
	Images      []AdvertiseImage  `json:"images"`
	Description string             `json:"description"`
	ExpireDate  primitive.DateTime `json:"expireDate"`
	Charges     []Charge          `json:"charges"`
	Payments    []Payment         `json:"payments"`
}
