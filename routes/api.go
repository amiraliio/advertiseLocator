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

//API routes
func API() {
	apiV1.GET("x-api-key", controllers.GenerateXAPIKey).Name = "generate-x-api-key"

	auth.POST("/person-register", controllers.PersonRegister).Name = "auth-person-register"
}
