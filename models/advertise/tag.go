package models

//Tag model
type Tag struct {
	Key   string `json:"key" bson:"key"`
	Value string `json:"value" bson:"value"`
	Min   string `json:"min" bson:"min"`
	Max   string `json:"max" bson:"max"`
}
