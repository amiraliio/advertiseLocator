package controllers

import (
	"github.com/amiraliio/advertiselocator/helpers"
	"github.com/amiraliio/advertiselocator/models"
	"github.com/amiraliio/advertiselocator/repositories/v1"
	"github.com/amiraliio/advertiselocator/requests"
	"github.com/labstack/echo/v4"
	"net/http"
)

func advertiseRepository() repositories.AdvertiseInterface {
	return new(repositories.AdvertiseRepository)
}

type AdvertiseService struct{
	 _ request.Get(models.APIKeyHeaderKey).(*models.API)


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
			"There Is a problem while binding the request",
		)
	}
	if err = request.Validate(advertiseRequest); err != nil {
		return helpers.ResponseError(request, http.StatusBadRequest, helpers.InsertTarget, http.StatusText(http.StatusBadRequest), "CA1000", "Insert Advertise", "There Is a problem while binding the request")
	}
}

func ListOfAdvertises(request echo.Context) (err error) {
	return nil
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
