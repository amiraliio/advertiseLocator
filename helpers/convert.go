package helpers

import (
	"errors"

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
