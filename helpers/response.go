package helpers

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

//sample of structure response
// {
//     "success": false,
//     "error": {
//         "code": 422,
//         "target": "insert",
//         "message": "unprocessable entity",
//         "details": [
//             {
//                 "code": "m-1000",
//                 "target": "password",
//                 "message": "must be at least 8 character"
//             }
//         ]
//     },
//     "data": []
// }

//the error and response message model implemented according to the OASIS Standard incorporating Approved standard
//referenced by link http://docs.oasis-open.org/odata/odata-json-format/v4.0/errata02/os/odata-json-format-v4.0-errata02-os-complete.html#_Toc403940655

const (
	INSERT_TARGET     string = "Insert"
	UPDATE_TARGET     string = "Update"
	DELETE_TARGET     string = "Delete"
	LIST              string = "Query"
	AUTH_TARGET       string = "Authentication"
	AUTHORIZED_TARGET string = "Authorization"
	REGISTER_TARGET   string = "Register"
	APIKEY_TARGET     string = "APIKey"
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
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Target  string         `json:"target"`
	Details []*ErrorDetail `json:"details"`
}

//ErrorDetail model
type ErrorDetail struct {
	Code    string `json:"code"`
	Target  string `json:"target"`
	Message string `json:"message"`
}

//ErrorResponse helper
func ResponseError(request echo.Context, httpCode int, httpTarget, httpMessage, internalCode, detailTarget, detailMessage string) error {
	errorMessage := new(ErrorModel)
	errorMessage.Code = httpCode
	errorMessage.Message = httpMessage
	errorMessage.Target = httpTarget
	body := new(ErrorDetail)
	body.Code = internalCode
	body.Target = detailTarget
	body.Message = detailMessage
	errorMessage.Details = append(errorMessage.Details, body)
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
