package models

//Location Model
type Location struct {
	Lat float32 `json:"lat" bson:"lat"`
	Lon float32 `json:"lon" bson:"lon"`
}
