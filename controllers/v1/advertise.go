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

func advertiseRepository() repositories.AdvertiseInterface {
	return new(repositories.AdvertiseRepository)
}

//AddAdvertise controller
func AddAdvertise(request echo.Context) error {
	advertiseRequest, err := helpers.BindAndValidateRequest(request, new(requests.Advertise))
	if err != nil {
		return helpers.ResponseError(request, http.StatusUnprocessableEntity, "CA1000", "Insert Advertise", err.Error())
	}
	requestModel := advertiseRequest.(*requests.Advertise)
	advertise := new(models.Advertise)
	advertise.Status = models.ActiveStatus
	advertise.ID = primitive.NewObjectID()
	advertise.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	advertise.CreatedBy = helpers.AuthData(request).UserID
	person := new(models.Person)
	person.ID = helpers.AuthData(request).UserID
	advertise.Advertiser = person
	advertise.Location = requestModel.Location
	advertise.Tags = requestModel.Tags
	advertise.Radius = requestModel.Radius
	advertise.Images = requestModel.Images
	advertise.Title = requestModel.Title
	advertise.Description = requestModel.Description
	advertise.Visibility = requestModel.Visibility
	result, err := advertiseRepository().InsertAdvertise(advertise)
	if err != nil {
		return helpers.ResponseError(request, http.StatusBadRequest, "CA1001", "Insert Advertise", err.Error())
	}
	return helpers.ResponseOk(request, http.StatusCreated, result)
}

//ListOfAdvertises controller
func ListOfAdvertises(request echo.Context) (err error) {
	filter := new(models.AdvertiseFilter)
	filter.UserID = helpers.AuthData(request).UserID
	results, err := advertiseRepository().ListOfAdvertise(filter)
	if err != nil {
		return helpers.ResponseError(request, http.StatusBadRequest, "CA1005", "List Of Advertise", err.Error())
	}
	return helpers.ResponseOk(request, http.StatusOK, results)
}

//GetAdvertise controller
func GetAdvertise(request echo.Context) (err error) {
	filter := new(models.AdvertiseFilter)
	filter.UserID = helpers.AuthData(request).UserID
	objectId, err := primitive.ObjectIDFromHex(request.Param("id"))
	if err != nil {
		return helpers.ResponseError(request, http.StatusBadRequest, "CA1007", "Get Advertise", err.Error())
	}
	filter.ID = objectId
	results, err := advertiseRepository().FindOne(filter)
	if err != nil {
		return helpers.ResponseError(request, http.StatusBadRequest, "CA1008", "Get Advertise", err.Error())
	}
	return helpers.ResponseOk(request, http.StatusOK, results)
}

func DeleteAdvertise(request echo.Context) (err error) {
	filter := new(models.AdvertiseFilter)
	filter.UserID = helpers.AuthData(request).UserID
	objectId, err := primitive.ObjectIDFromHex(request.Param("id"))
	if err != nil {
		return helpers.ResponseError(request, http.StatusBadRequest, "CA1009", "Delete Advertise", err.Error())
	}
	filter.ID = objectId
	results, err := advertiseRepository().DeleteOne(filter)
	if err != nil {
		return helpers.ResponseError(request, http.StatusBadRequest, "CA1010", "Delete Advertise", err.Error())
	}
	if results == 0 {
		return helpers.ResponseError(request, http.StatusBadRequest, "CA1013", "Delete Advertise", "Document for deleting doesn't exist")
	}
	return helpers.ResponseOk(request, http.StatusOK, "Deleted")
}
