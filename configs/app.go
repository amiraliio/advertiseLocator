package configs

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

//TODO add all environment variable to unique context

const (
	DevelopEnvironment    string = "DEV"
	ProductionEnvironment string = "PRODUCTION"
)

var (
	//Server variable to use framework instance in the other packages
	Server *echo.Echo = framework()
)

//instantiate framework
func framework() (framework *echo.Echo) {
	framework = echo.New()
	//instance of custom validator
	framework.Validator = &validation{validator: instantiateValidator()}
	//active logger
	framework = logger(framework)
	// Debug mode
	framework = debugger(framework)
	//active gzip if in production
	framework = gzip(framework)
	//cors
	framework.Use(middleware.CORS())
	//body limit
	framework.Use(middleware.BodyLimit("10M"))
	//security
	framework.Use(middleware.Secure())
	return framework
}

func logger(framework *echo.Echo) *echo.Echo {
	logger := viper.GetBool("APP.LOGGER")
	if logger {
		framework.Use(middleware.Logger())
	}
	return framework
}

func debugger(framework *echo.Echo) *echo.Echo {
	debug := viper.GetBool("APP.DEBUG")
	framework.Debug = debug
	if debug {
		framework.Use(middleware.Recover())
	}
	return framework
}

func gzip(framework *echo.Echo) *echo.Echo {
	if viper.GetString("APP.ENV") == ProductionEnvironment {
		framework.Use(middleware.GzipWithConfig(middleware.GzipConfig{
			Level: 5,
		}))
	}
	return framework
}
