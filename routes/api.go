package routes

import (
	"github.com/amiraliio/advertiselocator/configs"
	"github.com/amiraliio/advertiselocator/controllers/v1"
	"github.com/labstack/echo/v4"
)

var (
	apiV1 *echo.Group = configs.Server.Group("/api/v1")
	auth  *echo.Group = apiV1.Group("/auth")
)

func api() {
	auth.POST("/person-register", controllers.PersonRegister).Name = "auth-person-register"
}
