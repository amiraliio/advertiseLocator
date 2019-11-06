package helpers

import (
	"reflect"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

//Flatten can flat data models of embedded struct
func Flatten(object interface{}) (interface{}, error) {
	data, err := bson.Marshal(object)
	if err != nil {
		return nil, err
	}
	var entityModel map[string]interface{}
	if err = bson.Unmarshal(data, &entityModel); err != nil {
		return nil, err
	}
	return entityModel, nil
}

//IsInstance helper- checking the source model is instance of distance model
func IsInstance(src, dst interface{}) bool {
	return reflect.TypeOf(src) == reflect.TypeOf(dst)
}

func BindAndValidateRequest(request echo.Context, requestModel interface{}) (interface{}, error) {
	if err := request.Bind(requestModel); err != nil {
		return nil, err
	}
	if err := request.Validate(requestModel); err != nil {
		return nil, err
	}
	return requestModel, nil
}
