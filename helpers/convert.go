package helpers

import (
	"errors"
	"reflect"
	"strconv"

	"code.cloudfoundry.org/bytefmt"
)

func ConvertByte(byteNumber int64, convertTo string) (uint64, error) {
	switch convertTo {
	case "MB":
		return bytefmt.ToMegabytes(bytefmt.ByteSize(uint64(byteNumber)))
	default:
		return uint64(0), errors.New("argument for converting is invalid")
	}
}

func CheckAndReturnNumeric(value string) (response interface{}, dataType reflect.Kind, err error) {
	intValue, err := strconv.Atoi(value)
	if err == nil {
		return intValue, reflect.Int, nil
	} else {
		floatValue, err := strconv.ParseFloat(value, 64)
		if err == nil {
			return floatValue, reflect.Float64, nil
		}
	}
	return nil, 0, errors.New("the string isn't numeric")
}
