package helpers

import (
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
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

//IsInstance helper
func IsInstance(src, dst interface{}) bool {
	return reflect.TypeOf(src) == reflect.TypeOf(dst)
}
