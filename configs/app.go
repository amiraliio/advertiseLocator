package configs

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
)

var (
	//Server variable to start echo framework instance
	Server *echo.Echo = echo.New()
)

//Init configs
func Init() {
	environmentConfigs()
}

//load environment variables from .env files
func environmentConfigs() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
}
