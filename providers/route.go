package providers

import (
	"encoding/json"
	"fmt"

	"github.com/amiraliio/advertiselocator/configs"
	"github.com/amiraliio/advertiselocator/routes"
	"github.com/spf13/viper"
)

//Init routes in the package main
func initRoutes() {
	routes.API()
	if viper.GetBool("APP.SHOW_ROUTES") {
		printRoutesToConsole()
	}
	configs.Server.Logger.Fatal(configs.Server.Start(":" + viper.GetString("APP.PORT")))
}

//print whole project routes in the startup console
func printRoutesToConsole() {
	routesList, err := json.MarshalIndent(configs.Server.Routes(), "", "  ")
	if err != nil {
		configs.Server.Logger.Warn(err.Error())
	}
	fmt.Println(string(routesList))
}
