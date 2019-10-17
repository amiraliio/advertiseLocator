package routes

import (
	"github.com/amiraliio/advertiselocator/controllers/v1"
	"github.com/labstack/echo/v4"
)

//Init function to initialize routes
func Init() {
	server := echo.New()
	apiGroup := server.Group("/api/v1/auth")
	apiGroup.POST("/person/register", controllers.PersonRegister).Name = "auth-person-register"

}
