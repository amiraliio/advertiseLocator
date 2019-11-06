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

//Notice: this controller detail error code start with CU which is abbreviation for Controller Auth
//so each go file has their own unique code prefix, which implemented by responsibility + entity name

func authRepository() repositories.AuthInterface {
	return new(repositories.AuthRepository)
}

//TODO return internal status code from repo

//PersonRegister controller to register person
func PersonRegister(request echo.Context) (err error) {
	registerRequest, err := helpers.BindAndValidateRequest(request, new(requests.PersonRegister))
	if err != nil {
		return helpers.ResponseError(request, http.StatusUnprocessableEntity, "CU-1000", "Validate", err.Error())
	}
	requestModel := registerRequest.(*requests.PersonRegister)
	//person model
	person := new(models.Person)
	personID := primitive.NewObjectID()
	person.ID = personID
	person.Status = models.ActiveStatus
	person.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	person.CreatedBy = personID
	person.UserType = models.PersonUserType
	person.Email = requestModel.Email
	person.IP = request.RealIP()
	//auth model
	auth := new(models.Auth)
	authID := primitive.NewObjectID()
	auth.ID = authID
	auth.Status = models.ActiveStatus
	auth.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	auth.CreatedBy = personID
	auth.UserType = models.PersonUserType
	auth.Value = requestModel.Email
	auth.IP = request.RealIP()
	auth.UserID = personID
	hashedPassword, err := helpers.HashPassword(requestModel.Password)
	if err != nil {
		return helpers.ResponseError(request, http.StatusBadRequest, "CU-1001", "Encrypt Password", err.Error())
	}
	auth.Password = hashedPassword
	auth.Type = models.EmailAuthType
	client, err := clientMapper(request, auth, requestModel.Client, helpers.APIKeyData(request))
	if err != nil {
		return helpers.ResponseError(request, http.StatusBadRequest, "CU-1002", "Map Client", err.Error())
	}
	result, err := authRepository().PersonRegister(person, auth, client)
	if err != nil {
		return helpers.ResponseError(request, http.StatusBadRequest, "CU-1003", "Register Person", err.Error())
	}
	return helpers.ResponseOk(request, http.StatusCreated, result)
}

//PersonLogin controller
func PersonLogin(request echo.Context) (err error) {
	loginRequest, err := helpers.BindAndValidateRequest(request, new(requests.PersonLogin))
	if err != nil {
		return helpers.ResponseError(request, http.StatusUnprocessableEntity, "CU-1004", "Validate", err.Error())
	}
	requestModel := loginRequest.(*requests.PersonLogin)
	auth, err := authRepository().GetAuthData(requestModel.Email)
	if err != nil {
		return helpers.ResponseError(request, http.StatusNonAuthoritativeInfo, "CU-1005", "Fetch Auth Data", err.Error())
	}
	if !helpers.CheckPasswordHash(requestModel.Password, auth.Password) {
		return helpers.ResponseError(request, http.StatusNonAuthoritativeInfo, "CU-1006", "Fetch Auth Data", "auth value or password is wrong")
	}
	client, err := clientMapper(request, auth, requestModel.Client, helpers.APIKeyData(request))
	if err != nil {
		return helpers.ResponseError(request, http.StatusBadRequest, "CU-1007", "Map client", err.Error())
	}
	_, err = authRepository().InsertClient(client)
	if err != nil {
		return helpers.ResponseError(request, http.StatusBadRequest, "CU-1008", "Insert client", err.Error())
	}
	return helpers.ResponseOk(request, http.StatusOK, client)
}

func clientMapper(request echo.Context, auth *models.Auth, clientRequest *requests.Client, xAPIKeyData *models.API) (*models.Client, error) {
	//client model
	client := new(models.Client)
	clientID := primitive.NewObjectID()
	client.ID = clientID
	client.Status = models.ActiveStatus
	client.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	client.CreatedBy = auth.UserID
	client.UserID = auth.UserID
	client.UserType = models.PersonUserType
	client.IP = request.RealIP()
	client.ClientID = clientRequest.ID
	client.Version = clientRequest.Version
	client.OSType = clientRequest.OsType
	client.OSVersion = clientRequest.OsVersion
	client.LastLogin = primitive.NewDateTimeFromTime(time.Now())
	client.API = xAPIKeyData
	refreshToken, err := helpers.EncodeToken(uuid.New().String(), models.PersonUserType, os.Getenv("CLIENT_TOKEN_EXPIRE_DAY"))
	if err != nil {
		return nil, err
	}
	client.RefreshToken = refreshToken.Token
	clientToken, err := helpers.EncodeToken(auth.UserID.Hex(), models.PersonUserType, os.Getenv("CLIENT_TOKEN_EXPIRE_DAY"))
	if err != nil {
		return nil, err
	}
	client.Token = clientToken.Token
	rand.Seed(time.Now().UnixNano())
	client.VerificationCode = rand.Int()
	client.ExpireDate = clientToken.ExpireDate
	client.Auth = auth
	return client, nil
}
