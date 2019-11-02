package controllers

import (
	"github.com/amiraliio/advertiselocator/repositories/v1"
)

func advertiseRepository() repositories.AdvertiseRepository {
	return new(repositories.AdvertiseService)
}


func AddAdvertise(){

}