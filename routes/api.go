package routes

import (
	"github.com/amiraliio/advertiselocator/configs"
	"github.com/amiraliio/advertiselocator/controllers/v1"
	"github.com/amiraliio/advertiselocator/middleware"
	"github.com/labstack/echo/v4"
)

var (
	apiV1 *echo.Group = configs.Server.Group("/api/v1")
	auth  *echo.Group = apiV1.Group("/auth")
)

//API routes
func API() {
	apiV1.POST("/x-api-key", controllers.GenerateAPIKey).Name = "api-v1-generate-x-api-key"

	auth.Use(middleware.CheckAPIKey)
	auth.POST("/person-register", controllers.PersonRegister).Name = "api-v1-auth-person-register"
}
