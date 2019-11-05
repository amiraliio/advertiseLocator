package providers

import _ "github.com/joho/godotenv/autoload" //_ autoloaded config variables from .env file

//TODO validate for env parameter
//TODO for example .env app_key length must be 32 byte

//Start application and initialize application assets
func Start() {
	initRoutes()
	register()
}

func register() {
	//do what you want in startup in this method
}
