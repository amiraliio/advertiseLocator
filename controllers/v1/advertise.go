package controllers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/amiraliio/advertiselocator/helpers"
	"github.com/amiraliio/advertiselocator/models"
	"github.com/amiraliio/advertiselocator/repositories/v1"
	"github.com/amiraliio/advertiselocator/requests"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TODO add image url to originalURL
//TODO there must be one main image
//TODO update images type and size

//Notice: this controller detail error code start with CA which is abbreviation for Controller Advertise
//so each go file has their own unique code prefix, which implemented by responsibility + entity name

func advertiseRepository() repositories.AdvertiseInterface {
	return new(repositories.AdvertiseRepository)
}

//AddAdvertise controller
func AddAdvertise(request echo.Context) error {
	advertiseRequest, err := helpers.BindAndValidateRequest(request, new(requests.Advertise))
	if err != nil {
		return helpers.ResponseError(request, err, http.StatusUnprocessableEntity, "CA-1000", "Validatation", err.Error())
	}
	requestModel := advertiseRequest.(*requests.Advertise)
	if requestModel.Visibility != models.PrivateVisibility && requestModel.Visibility != models.PublicVisibility {
		return helpers.ResponseError(request, nil, http.StatusUnprocessableEntity, "CA-1001", "Validate Visibility Type", "type of visibility is invalid and must be one of the "+models.PrivateVisibility+", "+models.PublicVisibility)
	}
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
		return helpers.ResponseError(request, err, http.StatusNotModified, "CA-1002", "Insert Advertise", err.Error())
	}
	return helpers.ResponseOk(request, http.StatusCreated, result)
}

//ListOfAdvertises controller
func ListOfAdvertises(request echo.Context) (err error) {
	//filter mapper
	filter := new(models.AdvertiseFilter)
	if request.Get(models.AuthorizationHeaderKey) != nil {
		filter.UserID = helpers.AuthData(request).UserID
	}
	if request.QueryParam("startDate") != "" {
		startDate, err := helpers.ParsDateTime(request.QueryParam("startDate"))
		if err != nil {
			return helpers.ResponseError(request, err, http.StatusBadRequest, "CA-1003", "Parse StartDate", err.Error())
		}
		filter.StartDate = primitive.NewDateTimeFromTime(startDate)
	}
	if request.QueryParam("endDate") != "" {
		endDate, err := helpers.ParsDateTime(request.QueryParam("endDate"))
		if err != nil {
			return helpers.ResponseError(request, err, http.StatusBadRequest, "CA-1004", "Parse EndDate", err.Error())
		}
		filter.EndDate = primitive.NewDateTimeFromTime(endDate)
	}
	filter.LastIndex = request.QueryParam("lastIndex")
	if request.QueryParam("count") == "" {
		filter.Count = 20
	} else {
		count, err := strconv.Atoi(request.QueryParam("count"))
		if err != nil {
			return helpers.ResponseError(request, err, http.StatusBadRequest, "CA-1006", "Str count To Int", err.Error())
		}
		filter.Count = count
	}
	filter.Sort = request.QueryParam("sort")
	queryParam := request.QueryParam("query")
	//TODO implement below mapper as functional
	if queryParam != "" {
		query := &queryParam
		if strings.HasSuffix(*query, "==") {
			*query = strings.TrimSuffix(*query, "==")
		}
		if strings.HasSuffix(*query, "=") {
			*query = strings.TrimSuffix(*query, "=")
		}
		tagsAsString, err := base64.RawStdEncoding.DecodeString(*query)
		if err != nil {
			return helpers.ResponseError(request, err, http.StatusBadRequest, "CA-1007", "Decode Base64", err.Error())
		}
		if err = json.Unmarshal(tagsAsString, filter); err != nil {
			return helpers.ResponseError(request, nil, http.StatusBadRequest, "CA-1008", "Map Tags", "query data must be tags[] model")
		}
	}
	results, err := advertiseRepository().ListOfAdvertise(filter)
	if err != nil {
		return helpers.ResponseError(request, err, http.StatusBadRequest, "CA-1009", "List Of Advertise", err.Error())
	}
	return helpers.ResponseOk(request, http.StatusOK, results)
}

//GetAdvertise controller
func GetAdvertise(request echo.Context) (err error) {
	filter := new(models.AdvertiseFilter)
	if request.Get(models.AuthorizationHeaderKey) != nil {
		filter.UserID = helpers.AuthData(request).UserID
	}
	objectID, err := primitive.ObjectIDFromHex(request.Param("id"))
	if err != nil {
		return helpers.ResponseError(request, err, http.StatusBadRequest, "CA-1010", "Create ObjectID", err.Error())
	}
	filter.ID = objectID
	results, err := advertiseRepository().FindOne(filter)
	if err != nil {
		return helpers.ResponseError(request, err, http.StatusBadRequest, "CA-10011", "Get Advertise", err.Error())
	}
	return helpers.ResponseOk(request, http.StatusOK, results)
}

func DeleteAdvertise(request echo.Context) (err error) {
	filter := new(models.AdvertiseFilter)
	filter.UserID = helpers.AuthData(request).UserID
	objectID, err := primitive.ObjectIDFromHex(request.Param("id"))
	if err != nil {
		return helpers.ResponseError(request, err, http.StatusBadRequest, "CA-1012", "Create ObjectID", err.Error())
	}
	filter.ID = objectID
	_, err = advertiseRepository().DeleteOne(filter)
	if err != nil {
		return helpers.ResponseError(request, err, http.StatusBadRequest, "CA-1013", "Delete Advertise", err.Error())
	}
	return helpers.ResponseOk(request, http.StatusOK, "Deleted")
}
