package controllers

import (
	"github.com/amiraliio/advertiselocator/repositories/v1"
)

func advertiseRepository() repositories.AdvertiseInterface {
	return new(repositories.AdvertiseRepository)
}
