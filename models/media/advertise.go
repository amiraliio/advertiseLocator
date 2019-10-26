package models


//AdvertiseImage model
type AdvertiseImage struct {
	Image    ",inline"
	IsMain   bool   `json:"isMain" bson:"isMain"`
	Caption  string `json:"caption" bson:"caption"`
	Show     bool   `json:"show" bson:"show"`
	Priority byte   `json:"priority" bson:"priority"`
}
