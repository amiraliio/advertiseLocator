package routes

import (
	"github.com/amiraliio/advertiselocator/configs"
	"os"
)

//Init routes
func Init() {
	api()
	configs.Server.Logger.Fatal(configs.Server.Start(":" + os.Getenv("SERVER_PORT")))
}
