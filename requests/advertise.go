package requests

type Advertise struct {
	Location    *Location         `json:"location" validate:"required"`
	Tags        []Tag             `json:"tags" validate:"required,unique,max=200"`
	Radius      uint16            `json:"radius" validate:"omitempty,numeric,max=5"`
	Images      []*AdvertiseImage `json:"images" validate:"required,unique,max=30"`
	Description string            `json:"description" validate:"required,min=5,max=4000"`
	Visibility  string            `json:"visibility" validate:"required"`
}

//Tag model
type Tag struct {
	Key   string `json:"key" validate:"required,min=1,max=100"`
	Value string `json:"value" validate:"required,min=1,max=1000"`
	Min   string `json:"min" validate:"omitempty,min=1,max=100"`
	Max   string `json:"max" validate:"omitempty,min=1,max=100"`
}

//AdvertiseImage model
type AdvertiseImage struct {
	URL      string `json:"url" validate:"required,min=10,max=1000"`
	IsMain   bool   `json:"isMain" validate:"omitempty"`
	Caption  string `json:"caption" validate:"omitempty,min=10,max=255"`
	Show     bool   `json:"show" validate:"required"`
	Priority byte   `json:"priority" validate:"required,numeric,max=3"`
}
