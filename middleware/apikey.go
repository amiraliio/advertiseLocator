package middleware

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/amiraliio/advertiselocator/configs"
	"github.com/amiraliio/advertiselocator/helpers"
	lang "github.com/amiraliio/advertiselocator/lang/eng"
	"github.com/amiraliio/advertiselocator/models"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TODO detect user agent os to match apikey token and agent request

//CheckAPIKey middleware
func CheckAPIKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(request echo.Context) error {
		xAPIKey := request.Request().Header.Get(models.APIKeyHeaderKey)
		if xAPIKey == "" {
			return configs.ErrorResponse(http.StatusForbidden, "m1000", "Access Forbidden", "apiKey", lang.MustSetValidAPIKey)
		}
		dataKey, err := helpers.DecodeToken(xAPIKey)
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		expireTime, err := strconv.Atoi(os.Getenv("API_KEY_TOKEN_EXPIRE_DAY"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if dataKey.CreatedAt < primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, -expireTime)) {
			return echo.NewHTTPError(http.StatusForbidden, lang.TheAPIKeyExpired)
		}
		request.Set(models.APIKeyHeaderKey, dataKey)
		return next(request)
	}
}
