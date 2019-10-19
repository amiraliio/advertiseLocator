package models


//AuthCollection collection
const AuthCollection string = "auth"

//Auth model
type Auth struct {
	BaseUser
	Value    string `json:"value"`
	Password string `json:"password"`
	Type     string `json:"type"`
}
