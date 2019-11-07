package requests

import (
	"github.com/amiraliio/advertiselocator/models"
)

//TODO validation for enums in request

type Advertise struct {
	Location    *models.Location         `json:"location" validate:"required"`
	Tags        []*models.Tag            `json:"tags" validate:"required,unique,max=200"`
	Radius      uint16                   `json:"radius" validate:"omitempty,numeric"`
	Images      []*models.AdvertiseImage `json:"images" validate:"required,unique,max=30"`
	Title       string                   `json:"title" validate:"required,min=10,max=500"`
	Description string                   `json:"description" validate:"required,min=5,max=4000"`
	Visibility  string                   `json:"visibility" validate:"required"`
}
