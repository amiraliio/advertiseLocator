package controllers

import (
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

func AddAdvertise(request echo.Context) (err error) {
	advertiseRequest := new(requests.Advertise)
	if err = request.Bind(advertiseRequest); err != nil {
		return helpers.ResponseError(
			request,
			http.StatusBadRequest,
			helpers.InsertTarget,
			http.StatusText(http.StatusBadRequest),
			"CA1000",
			"Insert Advertise",
			err.Error(),
		)
	}
	if err = request.Validate(advertiseRequest); err != nil {
		return helpers.ResponseError( //TODO problem validation error response
			request,
			http.StatusUnprocessableEntity,
			helpers.InsertTarget,
			http.StatusText(http.StatusUnprocessableEntity),
			"CA1001",
			"Insert Advertise",
			err.Error(),
		)
	}
	advertise := new(models.Advertise)
	advertise.Status = models.ActiveStatus
	advertiseID := primitive.NewObjectID()
	advertise.ID = advertiseID
	advertise.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	advertise.CreatedBy = primitive.NilObjectID //TODO must change to user id
	advertise.Location = advertiseRequest.Location
	advertise.Tags = advertiseRequest.Tags
	person := new(models.Person)
	person.ID = primitive.NilObjectID //TODO must change to user id
	advertise.Advertiser = person
	advertise.Radius = advertiseRequest.Radius
	advertise.Images = advertiseRequest.Images
	advertise.Title = advertiseRequest.Title
	advertise.Description = advertiseRequest.Description
	advertise.Visibility = advertiseRequest.Visibility
	result, err := advertiseRepository().InsertAdvertise(advertise)
	if err != nil {
		return helpers.ResponseError(
			request,
			http.StatusBadRequest,
			helpers.InsertTarget,
			http.StatusText(http.StatusBadRequest),
			"CA1002",
			"Insert Advertise",
			err.Error(),
		)
	}
	return helpers.ResponseOk(request, http.StatusCreated, result)
}

func ListOfAdvertises(request echo.Context) (err error) {
	results, err := advertiseRepository().ListOfAdvertise()
	if err != nil {
		return helpers.ResponseError(
			request,
			http.StatusBadRequest,
			helpers.QueryTarget,
			http.StatusText(http.StatusBadRequest),
			"CA1003",
			"List Of Advertise",
			err.Error(),
		)
	}
	return helpers.ResponseOk(request, http.StatusCreated, results)
}

func GetAdvertise(request echo.Context) (err error) {
	return nil
}

func DeleteAdvertise(request echo.Context) (err error) {
	return nil
}

func UpdateAdvertise(request echo.Context) (err error) {
	return nil
}
