package main

import (
	"github.com/amiraliio/advertiselocator/configs"
	"fmt"
	app "github.com/amiraliio/advertiselocator/providers"
	"github.com/spf13/viper"
	"os"
)

func main() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("config.yaml")
	currentPath, err := os.Getwd()
	if err != nil {
		configs.Server.Logger.Fatal("cannot get current directory")
	}
	viper.AddConfigPath(currentPath)
	viper.AddConfigPath(".")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	if err := viper.ReadInConfig(); err != nil {
		configs.Server.Logger.Fatal("Cannot read config file")
	}
	fmt.Println(viper.ConfigFileUsed())
	app.Start()
}
