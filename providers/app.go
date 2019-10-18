package providers

import (
	"github.com/amiraliio/advertiselocator/configs"
	"github.com/amiraliio/advertiselocator/routes"
)

//Start application and initialize application assets
func Start() {
	configs.Init()
	routes.Init()
	register()
}

func register() {
	//do what you want in startup in this method
}
