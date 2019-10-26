package requests

//APIKey request model
type APIKey struct {
	Type       string `json:"type" bson:"type" validate:"required"`
	Name       string `json:"name" bson:"name" validate:"required"`
	ExpireDate string `json:"expireDate" bson:"expireDate" validate:"required"`
}
