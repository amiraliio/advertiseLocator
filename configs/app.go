package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	validator "gopkg.in/go-playground/validator.v9"
)

//Init configs in the package main
func Init() {
	//set your initial config
}

var (
	//Server variable to use framework instance in the other packages
	Server *echo.Echo = framework()
)

//instantiate framework
func framework() *echo.Echo {
	framework := echo.New()
	//instance of custom validator
	framework.Validator = &validation{validator: validator.New()}
	//custom error response handler
	// framework.HTTPErrorHandler = customErrorHandler
	// Debug mode
	debug, err := strconv.ParseBool(os.Getenv("SERVER_DEBUG"))
	if err != nil {
		log.Println(err.Error())
	}
	framework.Debug = debug
	return framework
}
