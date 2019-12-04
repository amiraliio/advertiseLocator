//Package middleware ...
package middleware

import (
	"net/http"
	"time"

	"github.com/amiraliio/advertiselocator/helpers"
	lang "github.com/amiraliio/advertiselocator/lang/eng"
	"github.com/amiraliio/advertiselocator/models"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TODO detect user agent os to match apikey token and agent request
//TODO dynamic generate internal code

//CheckAPIKey middleware - check the requested
func CheckAPIKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(request echo.Context) error {
		xAPIKey := request.Request().Header.Get(models.APIKeyHeaderKey)
		if xAPIKey == "" {
			return helpers.ResponseError(request, nil, http.StatusUnauthorized, "MA-1000", helpers.AccessTarget, lang.MustSetValidAPIKey)
		}
		dataKey, err := helpers.DecodeToken(xAPIKey)
		if err != nil {
			return helpers.ResponseError(request, err, http.StatusUnauthorized, "MA-1001", "Decode API Key", lang.MustSetValidAPIKey)
		}
		if dataKey.CreatedAt < primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, -viper.GetInt("AUTH.API_KEY_TOKEN_EXPIRE_DAY"))) {
			return helpers.ResponseError(request, nil, http.StatusUnauthorized, "MA-1003", helpers.AccessTarget, lang.TheAPIKeyExpired)
		}
		apiModel := new(models.API)
		apiModel.CreatedAt = dataKey.CreatedAt
		apiModel.Key = dataKey.Key
		apiModel.Type = dataKey.Type
		request.Set(models.APIKeyHeaderKey, apiModel)
		return next(request)
	}
}

//CheckIsPerson middleware - checking the requested user is person or not and has access
func CheckIsPerson(next echo.HandlerFunc) echo.HandlerFunc {
	return userAccess(next, models.PersonUserType)
}

//CheckIsAdmin middleware - checking the requested user is admin or not and access
func CheckIsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return userAccess(next, models.AdminUserType)
}

func userAccess(next echo.HandlerFunc, userType string) echo.HandlerFunc {
	return func(request echo.Context) error {
		auth := request.Request().Header.Get(models.AuthorizationHeaderKey)
		if auth == "" {
			return helpers.ResponseError(request, nil, http.StatusUnauthorized, "MA-1004", helpers.AccessTarget, "Access denied")
		}
		data, err := helpers.DecodeToken(auth)
		if err != nil {
			return helpers.ResponseError(request, err, http.StatusUnauthorized, "MA-1005", helpers.AccessTarget, err.Error())
		}
		if data.Type != userType {
			return helpers.ResponseError(request, nil, http.StatusUnauthorized, "MA-1006", helpers.AccessTarget, "Access denied")
		}
		if data.CreatedAt < primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, -viper.GetInt("AUTH.CLIENT_TOKEN_EXPIRE_DAY"))) {
			return helpers.ResponseError(request, nil, http.StatusUnauthorized, "MA-1007", helpers.AccessTarget, "Token expired")
		}
		client := new(models.Client)
		client.CreatedAt = data.CreatedAt
		objectID, err := primitive.ObjectIDFromHex(data.Key)
		if err != nil {
			return helpers.ResponseError(request, err, http.StatusBadRequest, "MA-1008", "Create ObjectID", err.Error())
		}
		client.UserID = objectID
		request.Set(models.AuthorizationHeaderKey, client)
		return next(request)
	}
}

func PublicAccess(next echo.HandlerFunc) echo.HandlerFunc {
	return func(request echo.Context) error {
		return next(request)
	}
}
