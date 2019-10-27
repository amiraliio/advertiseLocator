package middleware

import (
	"net/http"
	"time"

	"github.com/amiraliio/advertiselocator/helpers"
	"github.com/amiraliio/advertiselocator/models"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TODO detect user agent os to match apikey token and agent

//CheckAPIKey middleware
func CheckAPIKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(request echo.Context) error {
		xAPIKey := request.Request().Header.Get(models.APIKeyHeaderKey)
		if xAPIKey == "" {
			return echo.NewHTTPError(http.StatusForbidden, "A valid API key must be set")
		}
		dataKey, err := helpers.DecodeToken(xAPIKey)
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		if dataKey.ExpireDate < primitive.NewDateTimeFromTime(time.Now()) {
			return echo.NewHTTPError(http.StatusForbidden, "The API key has been expired")
		}
		request.Set(models.APIKeyHeaderKey, dataKey)
		return next(request)
	}
}
