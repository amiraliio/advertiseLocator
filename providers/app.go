package providers

import (
	"github.com/amiraliio/advertiselocator/configs"
)

//Start application and initialize application assets
func Start() {
	configs.Init()
	initRoutes()
	register()
}

func register() {
	//do what you want in startup in this method
}
