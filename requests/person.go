package requests

//PersonRegister request models
type PersonRegister struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=36"`
	Client   Client `json:"client" validate:"required"`
}



//PersonLogin struct
type PersonLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=36"`
	Client   Client `json:"client" validate:"required"`
}