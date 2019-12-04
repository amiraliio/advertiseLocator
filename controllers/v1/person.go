package controllers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/amiraliio/advertiselocator/helpers"
	"github.com/amiraliio/advertiselocator/models"
	"github.com/labstack/echo/v4"
)

//PersonListOfAdvertises controller
func PersonListOfAdvertises(request echo.Context) (err error) {
	//filter mapper
	filter := new(models.AdvertiseFilter)
	if request.Get(models.AuthorizationHeaderKey) != nil {
		filter.UserID = helpers.AuthData(request).UserID
	}
	if request.QueryParam("page") == "" {
		filter.Page = 1
	} else {
		page, err := strconv.ParseInt(request.QueryParam("page"), 10, 64)
		if err != nil {
			return helpers.ResponseError(request, err, http.StatusBadRequest, "CP-1001", "Str Page To Int", err.Error())
		}
		filter.Page = page
	}
	if request.QueryParam("limit") == "" {
		filter.Limit = 50
	} else {
		limit, err := strconv.ParseInt(request.QueryParam("limit"), 10, 64)
		if err != nil {
			return helpers.ResponseError(request, err, http.StatusBadRequest, "CP-1002", "Str Limit To Int", err.Error())
		}
		filter.Limit = limit
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
			return helpers.ResponseError(request, err, http.StatusBadRequest, "CP-1003", "Decode Base64", err.Error())
		}
		if err = json.Unmarshal(tagsAsString, filter); err != nil {
			return helpers.ResponseError(request, nil, http.StatusBadRequest, "CP-1004", "Map Tags", "query data must be tags[] model")
		}
	}
	results, err := advertiseRepository().ListOfAdvertise(filter)
	if err != nil {
		return helpers.ResponseError(request, err, http.StatusBadRequest, "CP-1005", "List Of Advertise", err.Error())
	}
	pagination := new(helpers.PaginationModel)
	if request.QueryParam("page") != "" {
		page, _ := strconv.Atoi(request.QueryParam("page"))
		pagination.Page = page
	} else {
		pagination.Page = 1
	}
	if request.QueryParam("limit") != "" {
		limit, _ := strconv.Atoi(request.QueryParam("limit"))
		pagination.Limit = limit
	} else {
		pagination.Limit = 50
	}
	request.Set("pagination", pagination)
	return helpers.ResponseOk(request, http.StatusOK, results)
}
