package providers

//this is improtant to first load environment autoload config
import _ "github.com/joho/godotenv/autoload" //_ autoloaded config variables from .env file

import (
	"os"

	"github.com/amiraliio/advertiselocator/configs"
)

//Start application and initialize application assets
func Start() {
	register()
	initRoutes()
}

//do what you want in startup in this method
func register() {
	if len(os.Getenv("APP_KEY")) != 32 {
		configs.Server.Logger.Fatal("Length of APP_KEY must be 32 byte")
	}
	if os.Getenv("APP_ENV") != configs.DevelopEnvironment && os.Getenv("APP_ENV") != configs.ProductionEnvironment {
		configs.Server.Logger.Fatal("APP_ENV must be one of the [ DEV, PRODUCTION ] ")
	}
}
