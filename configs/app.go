package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	validator "gopkg.in/go-playground/validator.v9"
)

var (
	//Server variable to use framework instance in the other packages
	Server *echo.Echo = framework()
)

//Init configs in the package main
func Init() {
	environmentConfigs()
}

//instantiate framework
func framework() *echo.Echo {
	framework := echo.New()
	//instance of custom validator
	framework.Validator = &validation{validator: validator.New()}
	// Debug mode
	debug, err := strconv.ParseBool(os.Getenv("SERVER_DEBUG"))
	if err != nil {
		log.Println(err.Error())
	}
	framework.Debug = debug
	return framework
}

//load environment variables from .env files
func environmentConfigs() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
}
