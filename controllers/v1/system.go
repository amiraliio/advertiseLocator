package controllers

import (
	"net/http"
	"os"

	"github.com/amiraliio/advertiselocator/helpers"
	"github.com/amiraliio/advertiselocator/models"
	"github.com/amiraliio/advertiselocator/repositories/v1"
	"github.com/amiraliio/advertiselocator/requests"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Notice: this controller detail error code start with CS which is abbreviation for Controller System
//so each go file has their own unique code prefix, which implemented by responsibility + entity name

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

func getSystemRepo() repositories.SystemInterface {
	return new(repositories.SystemRepository)
}

//GenerateAPIKey controller
func GenerateAPIKey(request echo.Context) (err error) {
	requestAPIKey, err := helpers.BindAndValidateRequest(request, new(requests.APIKey))
	if err != nil {
		return helpers.ResponseError(request, err, http.StatusUnprocessableEntity, "CS-1000", "Validatation", err.Error())
	}
	requestModel := requestAPIKey.(*requests.APIKey)
	uuidAsString := uuid.New().String()
	token, err := helpers.EncodeToken(uuidAsString, requestModel.Type, os.Getenv("API_KEY_TOKEN_EXPIRE_DAY"))
	if err != nil {
		return helpers.ResponseError(request, err, http.StatusBadRequest, "CS-1001", "Encryption", err.Error())
	}
	api := new(models.API)
	api.Key = uuidAsString
	api.Name = requestModel.Name
	api.ExpireDate = token.ExpireDate
	api.Token = token.Token
	api.Type = requestModel.Type
	api.Description = requestModel.Description
	api.ID = primitive.NewObjectID()
	api.Status = models.ActiveStatus
	api.CreatedAt = token.CreatedAt
	api.CreatedBy = primitive.NilObjectID
	data, err := getSystemRepo().CreateAPIKey(api)
	if err != nil {
		return helpers.ResponseError(request, err, http.StatusBadRequest, "CS-1002", "Insert API Key", err.Error())
	}
	return helpers.ResponseOk(request, http.StatusCreated, data)
}
