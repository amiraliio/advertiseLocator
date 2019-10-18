package configs

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"time"
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

//TimeOut context to use some cases like database connection or etc
func TimeOut(t byte) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(t)*time.Second)
}
