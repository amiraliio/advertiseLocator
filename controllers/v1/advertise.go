package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/amiraliio/advertiselocator/helpers"
	"github.com/amiraliio/advertiselocator/models"
	"github.com/amiraliio/advertiselocator/repositories/v1"
	"github.com/amiraliio/advertiselocator/requests"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TODO check access just person can add advertise
//TODO use controller service

func advertiseRepository() repositories.AdvertiseInterface {
	return new(repositories.AdvertiseRepository)
}

//TODO move to a base place
func authData(request echo.Context) (*models.Client, error) {
	auth := request.Get(models.AuthorizationHeaderKey)
	if !helpers.IsInstance(auth, (*models.Client)(nil)) {
		return nil, errors.New("auth need")
	}
	return auth.(*models.Client), nil
}

//AddAdvertise controller
func AddAdvertise(request echo.Context) (err error) {
	authData, err := authData(request)
	if err != nil {
		return helpers.ResponseError(request, http.StatusBadRequest, helpers.InsertTarget, http.StatusText(http.StatusBadRequest), "CA1001", "Find Advertise", "Auth data must be instance of client model")
	}
	advertiseRequest := new(requests.Advertise)
	if err = request.Bind(advertiseRequest); err != nil {
		return helpers.ResponseError(request, http.StatusBadRequest, helpers.InsertTarget, http.StatusText(http.StatusBadRequest), "CA1002", "Insert Advertise", err.Error())
	}
	if err = request.Validate(advertiseRequest); err != nil {
		//TODO problem validation error response
		return helpers.ResponseError(request, http.StatusUnprocessableEntity, helpers.InsertTarget, http.StatusText(http.StatusUnprocessableEntity), "CA1003", "Insert Advertise", err.Error())
	}
	advertise := new(models.Advertise)
	advertise.Status = models.ActiveStatus
	advertiseID := primitive.NewObjectID()
	advertise.ID = advertiseID
	advertise.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	advertise.CreatedBy = authData.UserID
	advertise.Location = advertiseRequest.Location
	advertise.Tags = advertiseRequest.Tags
	person := new(models.Person)
	person.ID = authData.UserID
	advertise.Advertiser = person
	advertise.Radius = advertiseRequest.Radius
	advertise.Images = advertiseRequest.Images
	advertise.Title = advertiseRequest.Title
	advertise.Description = advertiseRequest.Description
	advertise.Visibility = advertiseRequest.Visibility
	result, err := advertiseRepository().InsertAdvertise(advertise)
	if err != nil {
		return helpers.ResponseError(request, http.StatusBadRequest, helpers.InsertTarget, http.StatusText(http.StatusBadRequest), "CA1004", "Insert Advertise", err.Error())
	}
	return helpers.ResponseOk(request, http.StatusCreated, result)
}

//ListOfAdvertises controller
func ListOfAdvertises(request echo.Context) (err error) {
	authData, err := authData(request)
	if err != nil {
		return helpers.ResponseError(request, http.StatusBadRequest, helpers.InsertTarget, http.StatusText(http.StatusBadRequest), "CA1001", "Find Advertise", "Auth data must be instance of client model")
	}
	// queries := request.QueryParams()
	filter := new(models.AdvertiseFilter)
	filter.UserID = authData.UserID
	results, err := advertiseRepository().ListOfAdvertise(filter)
	if err != nil {
		return helpers.ResponseError(request, http.StatusBadRequest, helpers.QueryTarget, http.StatusText(http.StatusBadRequest), "CA1005", "List Of Advertise", err.Error())
	}
	return helpers.ResponseOk(request, http.StatusOK, results)
}

//GetAdvertise controller
func GetAdvertise(request echo.Context) (err error) {
	authData, err := authData(request)
	if err != nil {
		return helpers.ResponseError(request, http.StatusBadRequest, helpers.InsertTarget, http.StatusText(http.StatusBadRequest), "CA1001", "Find Advertise", "Auth data must be instance of client model")
	}
	filter := new(models.AdvertiseFilter)
	filter.UserID = authData.UserID
	objectId, err := primitive.ObjectIDFromHex(request.Param("id"))
	if err != nil {
		return helpers.ResponseError(request, http.StatusBadRequest, helpers.QueryTarget, http.StatusText(http.StatusBadRequest), "CA1007", "Get Advertise", err.Error())
	}
	filter.ID = objectId
	results, err := advertiseRepository().FindOne(filter)
	if err != nil {
		return helpers.ResponseError(request, http.StatusBadRequest, helpers.QueryTarget, http.StatusText(http.StatusBadRequest), "CA1008", "Get Advertise", err.Error())
	}
	return helpers.ResponseOk(request, http.StatusOK, results)
}

func DeleteAdvertise(request echo.Context) (err error) {
	return nil
}

func UpdateAdvertise(request echo.Context) (err error) {
	return nil
}
