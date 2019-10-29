package requests

//TODO oneof validation for Type

//APIKey request model
type APIKey struct {
	Type        string `json:"type" bson:"type" validate:"required"`
	Name        string `json:"name" bson:"name" validate:"required,min=5,max=150"`
	Description string `json:"description" bson:"description" validate:"min=10,max=400"`
}
