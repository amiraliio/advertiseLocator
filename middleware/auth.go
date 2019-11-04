package middleware

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/amiraliio/advertiselocator/helpers"
	lang "github.com/amiraliio/advertiselocator/lang/eng"
	"github.com/amiraliio/advertiselocator/models"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TODO detect user agent os to match apikey token and agent request
//TODO dynamic generate internal code

//CheckAPIKey middleware
func CheckAPIKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(request echo.Context) error {
		xAPIKey := request.Request().Header.Get(models.APIKeyHeaderKey)
		if xAPIKey == "" {
			return helpers.ResponseError(
				request,
				http.StatusUnauthorized,
				helpers.AccessTarget,
				http.StatusText(http.StatusUnauthorized),
				"M1000",
				helpers.ApiKeyTarget,
				lang.MustSetValidAPIKey,
			)
		}
		dataKey, err := helpers.DecodeToken(xAPIKey)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		expireTime, err := strconv.Atoi(os.Getenv("API_KEY_TOKEN_EXPIRE_DAY"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if dataKey.CreatedAt < primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, -expireTime)) {
			return echo.NewHTTPError(http.StatusUnauthorized, lang.TheAPIKeyExpired)
		}
		apiModel := new(models.API)
		apiModel.CreatedAt = dataKey.CreatedAt
		apiModel.Key = dataKey.Key
		apiModel.Type = dataKey.Type
		request.Set(models.APIKeyHeaderKey, apiModel)
		return next(request)
	}
}

func CheckIsPerson(next echo.HandlerFunc) echo.HandlerFunc {
	return func(request echo.Context) error {
		auth := request.Request().Header.Get(models.AuthorizationHeaderKey)
		if auth == "" {
			return helpers.ResponseError(
				request,
				http.StatusUnauthorized,
				helpers.AccessTarget,
				http.StatusText(http.StatusUnauthorized),
				"M1001",
				helpers.AuthTarget,
				"Must be authenticated",
			)
		}
		data, err := helpers.DecodeToken(auth)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		if data.Type != models.PersonUserType {
			return echo.NewHTTPError(http.StatusUnauthorized, "Access denied")
		}
		expireDate, err := strconv.Atoi(os.Getenv("CLIENT_TOKEN_EXPIRE_DAY"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if data.CreatedAt < primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, -expireDate)) {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token Expired")
		}
		client := new(models.Client)
		client.CreatedAt = data.CreatedAt
		objectID, err := primitive.ObjectIDFromHex(data.Key)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		client.ID = objectID
		request.Set(models.AuthorizationHeaderKey, client)
		return next(request)
	}
}
