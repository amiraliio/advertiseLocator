//Package routes ..
package routes

import (
	"github.com/amiraliio/advertiselocator/configs"
	"github.com/amiraliio/advertiselocator/controllers/v1"
	"github.com/amiraliio/advertiselocator/middleware"
	"github.com/labstack/echo/v4"
)

var (
	apiV1       *echo.Group = configs.Server.Group("/api/v1")
	auth        *echo.Group = apiV1.Group("/auth", middleware.CheckAPIKey)
	advertise   *echo.Group = apiV1.Group("/advertises", middleware.CheckAPIKey)
	media       *echo.Group = apiV1.Group("/media", middleware.CheckAPIKey)
	publicMedia *echo.Group = configs.Server.Group("/media", middleware.CheckAPIKey)
)

//API routes
func API() {
	//generate api keys
	apiV1.POST("/x-api-key", controllers.GenerateAPIKey).Name = "api-v1-generate-x-api-key"

	//auth routes
	auth.POST("/person-register", controllers.PersonRegister).Name = "api-v1-auth-person-register"
	auth.POST("/person-login", controllers.PersonLogin).Name = "api-v1-auth-person-login"

	//advertise crud for person
	advertise.POST("", controllers.AddAdvertise, middleware.CheckIsPerson).Name = "api-v1-add-advertise"
	advertise.GET("", controllers.ListOfAdvertises, middleware.CheckIsPerson, middleware.PublicAccess).Name = "api-v1-list-advertise"
	advertise.GET("/:id", controllers.GetAdvertise, middleware.CheckIsPerson).Name = "api-v1-get-advertise"
	advertise.DELETE("/:id", controllers.DeleteAdvertise, middleware.CheckIsPerson).Name = "api-v1-delete-advertise"
	// advertise.PUT("/:id", controllers.UpdateAdvertise, middleware.CheckIsPerson).Name = "api-v1-update-advertise"

	//media crud
	media.POST("/:mediaType/upload", controllers.UploadMedia, middleware.CheckIsPerson)
	publicMedia.GET("/temp/:mediaType/:userID/:year/:month/:day/:uniqueID/:filename", controllers.GetMedia)

}
