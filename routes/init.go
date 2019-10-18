package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/amiraliio/advertiselocator/configs"
)

//Init routes in the package main
func Init() {
	api()
	printRoutesToConsole()
	configs.Server.Logger.Fatal(configs.Server.Start(":" + os.Getenv("SERVER_PORT")))
}

//print whole project routes in the startup console
func printRoutesToConsole() {
	routes, err := json.MarshalIndent(configs.Server.Routes(), "", "  ")
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(string(routes))
}
