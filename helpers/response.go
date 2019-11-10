package helpers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/amiraliio/advertiselocator/configs"
	"github.com/labstack/echo/v4"
)

//the error and response message model implemented according to the OASIS Standard incorporating Approved standard
//referenced by link http://docs.oasis-open.org/odata/odata-json-format/v4.0/errata02/os/odata-json-format-v4.0-errata02-os-complete.html#_Toc403940655

//response targets
const (
	InsertTarget string = "Insert"
	UpdateTarget string = "Update"
	DeleteTarget string = "Delete"
	QueryTarget  string = "Query"
	AccessTarget string = "Access"
)

type ResponseModel struct {
	Success    bool             `json:"success"`
	Error      *ErrorModel      `json:"error"`
	Data       interface{}      `json:"data"`
	Pagination *PaginationModel `json:"pagination"`
}

type PaginationModel struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	LastIndex int `json:"lastIndex"`
	Total     int `json:"total"`
}

//ErrorMessage model
type ErrorModel struct {
	Code       int            `json:"code"`
	Message    string         `json:"message"`
	Target     string         `json:"target"`
	Details    []*ErrorDetail `json:"details"`
	InnerError *InnerError    `json:"innererror"`
}

//InnerError model
type InnerError struct {
	Trace   string `json:"trace"`
	Context string `json:"context"`
}

//ErrorDetail model
type ErrorDetail struct {
	Code    string `json:"code"`
	Target  string `json:"target"`
	Message string `json:"message"`
}

func httpTarget(httpCode int, requestMethod string) string {
	switch {
	case httpCode == http.StatusUnauthorized:
		return AccessTarget
	case requestMethod == http.MethodPost:
		return InsertTarget
	case requestMethod == http.MethodGet:
		return QueryTarget
	case requestMethod == http.MethodPut:
		return UpdateTarget
	case requestMethod == http.MethodDelete:
		return DeleteTarget
	default:
		return ""
	}
}

//ErrorResponse helper
func ResponseError(request echo.Context, err error, httpCode int, internalCode, detailTarget, detailMessage string) error {
	errorMessage := new(ErrorModel)
	errorMessage.Code = httpCode
	errorMessage.Message = http.StatusText(httpCode)
	errorMessage.Target = httpTarget(httpCode, request.Request().Method)
	body := new(ErrorDetail)
	body.Code = internalCode
	if httpCode == http.StatusUnprocessableEntity && len(detailMessage) > 0 {
		target := strings.Fields(detailMessage)
		body.Target = strings.ToLower(target[0])
	} else {
		body.Target = detailTarget
	}
	body.Message = strings.ToLower(detailMessage)
	errorMessage.Details = append(errorMessage.Details, body)
	if configs.Server.Debug && err != nil {
		innerError := new(InnerError)
		innerError.Trace = err.Error()
		errorMessage.InnerError = innerError
	}
	response := new(ResponseModel)
	response.Success = false
	response.Error = errorMessage
	response.Data = nil
	response.Pagination = nil
	return request.JSONPretty(httpCode, response, "	")
}

func ResponseOk(request echo.Context, httpCode int, data interface{}) error {
	response := new(ResponseModel)
	response.Success = true
	response.Error = nil
	response.Data = data
	pagination := new(PaginationModel)
	page, _ := strconv.Atoi(request.QueryParam("page"))
	pagination.Page = page
	limit, _ := strconv.Atoi(request.QueryParam("limit"))
	pagination.Limit = limit
	pagination.Total = pagination.Page * pagination.Limit
	return request.JSONPretty(httpCode, response, "	")
}
