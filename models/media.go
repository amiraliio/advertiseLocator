package models

//Media model
type Media struct {
	OriginalURL string `json:"originalURL" bson:"originalURL"`
	URL         string `json:"url" bson:"url"`
	Type        string `json:"type" bson:"type"`
}

//Image model
type Image struct {
	Media ",inline"
	Size  string `json:"size" bson:"size"`
}

//Video model
type Video struct {
	Media ",inline"
}

//File model
type File struct {
	Media ",inline"
}

//AdvertiseImage model
type AdvertiseImage struct {
	Image    ",inline"
	IsMain   bool   `json:"isMain" bson:"isMain"`
	Caption  string `json:"caption" bson:"caption"`
	Show     bool   `json:"show" bson:"show"`
	Priority byte   `json:"priority" bson:"priority"`
}
