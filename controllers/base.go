package controllers

import (
	"net/http"

	"github.com/amiraliio/advertiselocator/helpers"
	"github.com/amiraliio/advertiselocator/models"
	"github.com/labstack/echo/v4"
)

//check api key exist and pass to controller
func checkAPIKey(request echo.Context) (*models.API, error) {
    xAPIKeyData := request.Get(models.APIKeyHeaderKey).(*models.API)

	if !helpers.IsInstance(xAPIKeyData, (*models.API)(nil)) {
		return nil, helpers.ResponseError(
			request,
			http.StatusForbidden,
			helpers.ApiKeyTarget,
			http.StatusText(http.StatusForbidden),
			"C1000",
			"Check API Key",
			"API key must be instance of API model",
		)
	}
	return xAPIKeyData, nil
}
