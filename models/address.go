package models

//Location Model
type Location struct {
	Lat float32 `json:"lat" bson:"lat" query:"lat" validate:"required,numeric"`
	Lon float32 `json:"lon" bson:"lon" query:"lon" validate:"required,numeric"`
}
