package controllers

import (
	"net/http"
	"time"

	"github.com/amiraliio/advertiselocator/models"
	"github.com/amiraliio/advertiselocator/repositories/v1"
	"github.com/amiraliio/advertiselocator/requests"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func authRepository() repositories.AuthRepository {
	return new(repositories.AuthService)
}

//PersonRegister controller to register person
func PersonRegister(request echo.Context) (err error) {
	registerRequest := new(requests.PersonRegister)
	if err = request.Bind(registerRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = request.Validate(registerRequest); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	person := new(models.Person)
	personID := primitive.NewObjectID()
	person.ID = personID
	person.Status = models.ActiveStatus
	person.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	person.CreatedBy = personID
	person.UpdatedBy = personID
	person.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	person.UserType = models.PersonUserType
	person.Email = registerRequest.Email
	person.IP = request.RealIP()

	auth := new(models.Auth)
	authID := primitive.NewObjectID()
	auth.ID = authID
	auth.Status = models.ActiveStatus
	auth.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	auth.CreatedBy = personID
	auth.UpdatedBy = personID
	auth.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	auth.UserType = models.PersonUserType
	auth.Value = registerRequest.Email
	auth.IP = request.RealIP()
	auth.UserID = personID
	auth.Password = registerRequest.Password //TODO hash this shit
	auth.Type = models.EmailAuthType

	client := new(models.AuthClient)
	clientID := primitive.NewObjectID()
	client.ID = clientID
	client.Status = models.ActiveStatus
	client.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	client.CreatedBy = personID
	client.UpdatedBy = personID
	client.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	client.UserID = personID
	client.UserType = models.PersonUserType
	client.IP = request.RealIP()
	client.ClientID = registerRequest.Client.ID
	client.Version = registerRequest.Client.Version
	client.LastLogin = primitive.NewDateTimeFromTime(time.Now())
	client.OSType = registerRequest.Client.OsType
	client.OSVersion = registerRequest.Client.OsVersion
	client.API.Key = ""
	client.API.Name = ""
	client.API.ExpireDate = ""
	client.API.Type = ""
	client.API.Token = ""
	client.RefreshToken = ""     //TODO refresh token
	client.Token = ""            //TODO token
	client.VerificationCode = "" //TODO
	//TODO other client fields

	result, err := authRepository().PersonRegister(person, auth, client)
	//TODO BaseResponse
	//return status code from repo
	if err != nil {
		return echo.NewHTTPError(http.StatusNotModified, err.Error())
	}
	return request.JSON(http.StatusCreated, result)
}
