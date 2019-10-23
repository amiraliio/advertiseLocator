package providers

import 	_ "github.com/joho/godotenv/autoload" //_ autoload the config variables from .env file

//Start application and initialize application assets
func Start() {
	initRoutes()
	register()
}

func register() {
	//do what you want in startup in this method
}
