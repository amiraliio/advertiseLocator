package routes

import (
	"github.com/amiraliio/advertiselocator/configs"
	"github.com/amiraliio/advertiselocator/controllers/v1"
	"github.com/labstack/echo/v4"
)

var (
	apiGroup     *echo.Group = configs.Server.Group("/api/v1")
	authAPIGroup *echo.Group = apiGroup.Group("/auth")
)

func api() {
	authAPIGroup.POST("auth/person/register", controllers.PersonRegister).Name = "auth-person-register"
}
