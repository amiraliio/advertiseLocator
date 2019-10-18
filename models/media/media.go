package models

//Media model
type Media struct {
	OriginalURL string `json:"originalURL"`
	URL         string `json:"url"`
	Type        string `json:"type"`
}

//Image model
type Image struct {
	Media
	Size string `json:"size"`
}

//Video model
type Video struct {
	Media
}

//File model
type File struct {
	Media
}

//AdvertiseImage model
type AdvertiseImage struct {
	Image
	IsMain   bool   `json:"isMain"`
	Caption  string `json:"caption"`
	Show     bool   `json:"show"`
	Priority byte   `json:"priority"`
}
