package controllers

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/amiraliio/advertiselocator/helpers"
	"github.com/amiraliio/advertiselocator/models"
	"github.com/amiraliio/advertiselocator/repositories/v1"
	"github.com/amiraliio/advertiselocator/requests"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func authRepository() repositories.AuthRepository {
	return new(repositories.AuthService)
}

//TODO BaseResponse
//return internal status code from repo

//PersonRegister controller to register person
func PersonRegister(request echo.Context) (err error) {
	//added from apikey middleware to context
	xAPIKeyData := request.Get(models.APIKeyHeaderKey)
	if !helpers.IsInstance(xAPIKeyData, (*models.API)(nil)) {
		return echo.NewHTTPError(http.StatusForbidden, "API key must be instance of API model")
	}
	registerRequest := new(requests.PersonRegister)
	if err = request.Bind(registerRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = request.Validate(registerRequest); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	//person model
	person := new(models.Person)
	personID := primitive.NewObjectID()
	person.ID = personID
	person.Status = models.ActiveStatus
	person.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	person.CreatedBy = personID
	person.UserType = models.PersonUserType
	person.Email = registerRequest.Email
	person.IP = request.RealIP()
	//auth model
	auth := new(models.Auth)
	authID := primitive.NewObjectID()
	auth.ID = authID
	auth.Status = models.ActiveStatus
	auth.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	auth.CreatedBy = personID
	auth.UserType = models.PersonUserType
	auth.Value = registerRequest.Email
	auth.IP = request.RealIP()
	auth.UserID = personID
	hashedPassword, err := helpers.HashPassword(registerRequest.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	auth.Password = hashedPassword
	auth.Type = models.EmailAuthType
	client, err := clientMapper(request, personID, registerRequest, xAPIKeyData.(*models.API), authID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	result, err := authRepository().PersonRegister(person, auth, client)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotModified, err.Error())
	}
	return request.JSON(http.StatusCreated, result)
}

//PersonLogin controller
func PersonLogin() (err error) {
	return nil
}

func clientMapper(request echo.Context, personID primitive.ObjectID, registerRequest *requests.PersonRegister, xAPIKeyData *models.API, authID primitive.ObjectID) (*models.Client, error) {
	//client model
	client := new(models.Client)
	clientID := primitive.NewObjectID()
	client.ID = clientID
	client.Status = models.ActiveStatus
	client.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	client.CreatedBy = personID
	client.UserID = personID
	client.UserType = models.PersonUserType
	client.IP = request.RealIP()
	client.ClientID = registerRequest.Client.ID
	client.Version = registerRequest.Client.Version
	client.LastLogin = primitive.NewDateTimeFromTime(time.Now())
	client.OSType = registerRequest.Client.OsType
	client.OSVersion = registerRequest.Client.OsVersion
	client.API.Key = xAPIKeyData.Key
	client.API.ExpireDate = xAPIKeyData.ExpireDate
	client.API.Type = xAPIKeyData.Type
	client.API.CreatedAt = xAPIKeyData.CreatedAt
	refreshToken, err := helpers.EncodeToken(uuid.New().String(), models.PersonUserType, os.Getenv("CLIENT_TOKEN_EXPIRE_DAY"))
	if err != nil {
		return nil, err
	}
	client.RefreshToken = refreshToken.Token
	clientToken, err := helpers.EncodeToken(personID.String(), models.PersonUserType, os.Getenv("CLIENT_TOKEN_EXPIRE_DAY"))
	if err != nil {
		return nil, err
	}
	client.Token = clientToken.Token
	rand.Seed(time.Now().UnixNano())
	client.VerificationCode = rand.Int()
	client.ExpireDate = clientToken.ExpireDate
	client.Auth.ID = authID
	return client, nil
}
