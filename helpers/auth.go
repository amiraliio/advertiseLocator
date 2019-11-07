package helpers

import (
	"github.com/amiraliio/advertiselocator/models"
	"github.com/labstack/echo/v4"
)

func AuthData(request echo.Context) *models.Client {
	return request.Get(models.AuthorizationHeaderKey).(*models.Client)
}

func APIKeyData(request echo.Context) *models.API {
	return request.Get(models.APIKeyHeaderKey).(*models.API)
}
