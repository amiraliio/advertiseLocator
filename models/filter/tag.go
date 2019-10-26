package models


//TagFilter filter
type TagFilter struct {
	BaseFilter ",inline"
	Key        string  `json:"key" query:"key" bson:"key"`
	Value      string  `json:"value" query:"value" bson:"value"`
	Min        float64 `json:"min" query:"min" bson:"min"`
	Max        float64 `json:"max" query:"max" bson:"max"`
}