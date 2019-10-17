package models

//Tag model
type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Min   string `json:"min"`
	Max   string `json:"max"`
}
