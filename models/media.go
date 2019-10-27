package models

//BaseMedia model
type BaseMedia struct {
	OriginalURL string `json:"originalURL" bson:"originalURL"`
	URL         string `json:"url" bson:"url"`
	Type        string `json:"type" bson:"type"`
}

//File model
type File struct {
	BaseMedia ",inline"
}

//Image model
type Image struct {
	BaseMedia ",inline"
	Size      string `json:"size" bson:"size"`
}

//Video model
type Video struct {
	BaseMedia ",inline"
}

//AdvertiseImage model
type AdvertiseImage struct {
	Image    ",inline"
	IsMain   bool   `json:"isMain" bson:"isMain"`
	Caption  string `json:"caption" bson:"caption"`
	Show     bool   `json:"show" bson:"show"`
	Priority byte   `json:"priority" bson:"priority"`
}
