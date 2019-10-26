package controllers

import (
	"net/http"
	"time"

	"github.com/amiraliio/advertiselocator/helpers"
	"github.com/amiraliio/advertiselocator/models"
	"github.com/amiraliio/advertiselocator/repositories/v1"
	"github.com/amiraliio/advertiselocator/requests"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	dataTime, err := time.Parse("2020-03-12 15:04:05", requestAPIKey.ExpireDate)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	token, err := helpers.EncodeToken(uuid, requestAPIKey.Type, primitive.NewDateTimeFromTime(dataTime))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	api := new(models.API)
	api.Key = uuid
	api.Name = requestAPIKey.Name
	api.ExpireDate = requestAPIKey.ExpireDate
	api.Token = token
	api.Type = requestAPIKey.Type
	api.ID = primitive.NewObjectID()
	api.Status = models.ActiveStatus
	api.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	api.CreatedBy = primitive.NilObjectID //TODO must be admin id and read from token
	api.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	api.UpdatedBy = primitive.NilObjectID //TODO must be admin id and read from token

	data, err := getSystemRepo().CreateAPIKey(api)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//TODO must have admin access and without checkAPI middleware
	//TODO validation for type and request
	//TODO save API key per platform in mongo
	//TODO platform are web, ios, android
	//TODO that must be unique by platform and packageName
	//TODO just admin can add this type of api key
	return request.JSON(http.StatusCreated, data)
}
