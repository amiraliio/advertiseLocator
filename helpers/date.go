package helpers

import (
	"errors"
	"time"
)

func ParsDateTime(date string) (dataTime time.Time, err error) {
	dataTime, err = time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		return time.Now(), errors.New(err.Error())
	}
	return dataTime, nil
}

