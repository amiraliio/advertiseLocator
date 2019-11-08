package configs

import (
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

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
	logger, err := strconv.ParseBool(os.Getenv("APP_LOGGER"))
	if err != nil {
		framework.Logger.Fatal(err.Error())
	}
	if logger {
		framework.Use(middleware.Logger())
	}
	return framework
}

func debugger(framework *echo.Echo) *echo.Echo {
	debug, err := strconv.ParseBool(os.Getenv("APP_DEBUG"))
	if err != nil {
		framework.Logger.Fatal(err.Error())
	}
	framework.Debug = debug
	if debug {
		framework.Use(middleware.Recover())
	}
	return framework
}

func gzip(framework *echo.Echo) *echo.Echo {
	if os.Getenv("APP_ENV") == ProductionEnvironment {
		framework.Use(middleware.GzipWithConfig(middleware.GzipConfig{
			Level: 5,
		}))
	}
	return framework
}
