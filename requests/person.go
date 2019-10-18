package requests

//PersonRegister request models
type PersonRegister struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Client   Client `json:"client" validate:"required"`
}
