package models



//Image model
type Image struct {
	BaseMedia ",inline"
	Size  string `json:"size" bson:"size"`
}