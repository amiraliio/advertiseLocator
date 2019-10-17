package routes

import (
	"github.com/amiraliio/advertiselocator/controllers/v1"
    "github.com/labstack/echo/v4"
	"fmt"
)

func Init(){
    server := echo.New()
    apiGroup := server.Group("/api/v1/auth", m ...echo.MiddlewareFunc)
    apiGroup.POST("", controllers.PersonRegister).Name = "auth-person-register"

}