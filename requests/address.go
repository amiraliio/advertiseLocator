package requests

//Location request Model
type Location struct {
	Lat float32 `json:"lat" validate:"required,latitude"`
	Lon float32 `json:"lon" validate:"required,longitude"`
}
