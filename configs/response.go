package configs

import (
	"github.com/labstack/echo/v4"
)

func customErrorHandler(err error, c echo.Context) {
	// echo.JSONPretty(http.StatusBadRequest, ErrorResponse, " ")
}

//ErrorMessage model
type ErrorMessage struct {
	Code    int          `json:"code"`
	Title   string       `json:"title"`
	Details []*ErrorBody `json:"details"`
}

//ErrorBody model
type ErrorBody struct {
	Code    string `json:"code"`
	Target  string `json:"target"`
	Message string `json:"message"`
}

//ErrorResponse helper
func ErrorResponse(httpCode int, internal string, title, target, message string) *echo.HTTPError {
	body := new(ErrorBody)
	body.Code = internal
	body.Target = target
	body.Message = message
	errorMessage := new(ErrorMessage)
	errorMessage.Code = httpCode
	errorMessage.Title = title
	errorMessage.Details = append(errorMessage.Details, body)
	return echo.NewHTTPError(httpCode, errorMessage)
}
