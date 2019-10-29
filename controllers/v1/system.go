package controllers

import (
	"os"
	"net/http"

	"github.com/amiraliio/advertiselocator/helpers"
	"github.com/amiraliio/advertiselocator/models"
	"github.com/amiraliio/advertiselocator/repositories/v1"
	"github.com/amiraliio/advertiselocator/requests"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TODO must have admin access and without checkAPI middleware
//TODO validation for type and request
//TODO save API key per platform in mongo
//TODO platform are web, ios, android
//TODO that must be unique by platform and packageName
//TODO just admin can add this type of api key
//TODO when a same token generated with name and same type the old token must inactive
//TODO move uuid to helpers
//TODO move date parser to helpers
//TODO change answer of errors
//TODO must be admin id and read from token for created_by

func getSystemRepo() repositories.SystemRepository {
	return new(repositories.SystemService)
}

//GenerateAPIKey controller
func GenerateAPIKey(request echo.Context) (err error) {
	requestAPIKey := new(requests.APIKey)
	if err = request.Bind(requestAPIKey); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = request.Validate(requestAPIKey); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	uuid := uuid.New().String()
	token, err := helpers.EncodeToken(uuid, requestAPIKey.Type, os.Getenv("API_KEY_TOKEN_EXPIRE_DAY"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	api := new(models.API)
	api.Key = uuid
	api.Name = requestAPIKey.Name
	api.ExpireDate = token.ExpireDate
	api.Token = token.Token
	api.Type = requestAPIKey.Type
	api.Description = requestAPIKey.Description
	api.ID = primitive.NewObjectID()
	api.Status = models.ActiveStatus
	api.CreatedAt = token.CreatedAt
	api.CreatedBy = primitive.NilObjectID
	data, err := getSystemRepo().CreateAPIKey(api)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return request.JSON(http.StatusCreated, data)
}
