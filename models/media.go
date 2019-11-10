package models

//image size
const (
	OriginalSize  string = "ORIGINAL"
	MediumSize    string = "MEDIUM"
	ThumbnailSize string = "THUMBNAIL"
)

//media groups
const (
	ImageMediaType  string = "images"
	VideosMediaType string = "videos"
	AudiosMediaType string = "audios"
	FilesMediaType  string = "files"
)

//media mimetypes
var (
	ImageMimeTypes []string = []string{"image/jpg", "image/jpeg", "image/pjpeg", "image/gif", "image/png"}
	VideMimeTypes  []string = []string{"video/mp4", "video/avi", "video/webm"}
	AudioMimeTypes []string = []string{"audio/mpeg", "audio/wave"}
	FileMimeTypes  []string = []string{"application/pdf"}
)

//BaseMedia model
type BaseMedia struct {
	OriginalURL string `json:"originalURL" bson:"originalURL"`
	URL         string `json:"url" bson:"url" validate:"required,min=10,max=1000"`
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
	IsMain   bool   `json:"isMain" bson:"isMain" validate:"omitempty"`
	Caption  string `json:"caption" bson:"caption" validate:"omitempty,min=10,max=255"`
	Show     bool   `json:"show" bson:"show" validate:"required"`
	Priority byte   `json:"priority" bson:"priority" validate:"required,numeric,max=3"`
}
