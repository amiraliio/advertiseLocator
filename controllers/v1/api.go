package controllers

import (
	"github.com/labstack/echo/v4"
)

//GenerateXAPIKey controller
func GenerateXAPIKey(request echo.Context) (err error) {
	//TODO save API key per platform in mongo
	//TODO platform are web, ios, android
	//TODO that must be unique by platform and packageName
	//TODO just admin can add this type of api key
	return nil
}
