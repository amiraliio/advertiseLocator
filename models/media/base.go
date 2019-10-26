package models


//Media model
type BaseMedia struct {
	OriginalURL string `json:"originalURL" bson:"originalURL"`
	URL         string `json:"url" bson:"url"`
	Type        string `json:"type" bson:"type"`
}