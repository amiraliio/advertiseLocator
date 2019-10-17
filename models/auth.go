package models

//Auth model
type Auth struct {
	BaseUser
	Value    string `json:"value"`
	Password string `json:"password"`
	Type     string `json:"type"`
}
