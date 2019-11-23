//package providers ...
package providers

import (
	"github.com/amiraliio/advertiselocator/configs"
	"github.com/spf13/viper"
)

//Start application and initialize application assets
func Start() {
	register()
	initRoutes()
}

//do what you want in startup in this method
func register() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		configs.Server.Logger.Fatal("Cannot read config file")
	}
	if len(viper.GetString("APP.KEY")) != 32 {
		configs.Server.Logger.Fatal("Length of APP_KEY must be 32 byte")
	}
	if viper.GetString("APP.ENV") != configs.DevelopEnvironment && viper.GetString("APP.ENV") != configs.ProductionEnvironment {
		configs.Server.Logger.Fatal("APP_ENV must be one of the [ DEV, PRODUCTION ] ")
	}
}
