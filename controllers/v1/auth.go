package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

//PersonRegister
func PersonRegister(context echo.Context) error {
	fmt.Println(context.RealIP())
	return nil
}

// func PersonLogin(context echo.Context) error{

// }
