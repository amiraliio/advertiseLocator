package requests

//TODO oneof validation for Type

//APIKey request model
type APIKey struct {
	Type        string `json:"type" validate:"required"`
	Name        string `json:"name" validate:"required,min=5,max=150"`
	Description string `json:"description" validate:"omitempty,min=10,max=400"`
}
