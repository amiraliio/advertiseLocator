package controllers

import (
	"errors"

	"github.com/amiraliio/advertiselocator/requests"
	"github.com/labstack/echo/v4"
)



func repository(){
	
}

//PersonRegister controller to register person
func PersonRegister(request echo.Context) (err error) {
	person := new(requests.PersonRegister)
	if err = request.Bind(person); err != nil {
		return errors.New(err.Error())
	}
	if err = request.Validate(person); err !=nil{
		return errors.New(err.Error())
	}
	//TODO sd
	return nil
}

// func PersonLogin(context echo.Context) error{

// }
