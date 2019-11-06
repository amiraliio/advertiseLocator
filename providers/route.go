package providers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/amiraliio/advertiselocator/configs"
	"github.com/amiraliio/advertiselocator/routes"
)

//Init routes in the package main
func initRoutes() {
	routes.API()
	printRoutesToConsole()
	configs.Server.Logger.Fatal(configs.Server.Start(":" + os.Getenv("SERVER_PORT")))
}

//print whole project routes in the startup console
func printRoutesToConsole() {
	routes, err := json.MarshalIndent(configs.Server.Routes(), "", "  ")
	if err != nil {
		configs.Server.Logger.Warn(err.Error())
	}
	fmt.Println(string(routes))
}
