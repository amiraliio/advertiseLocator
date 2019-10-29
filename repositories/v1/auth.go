package repositories

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"errors"
	"github.com/amiraliio/advertiselocator/helpers"
	"github.com/amiraliio/advertiselocator/models"
)

//AuthRepository interface
type AuthRepository interface {
	PersonRegister(person *models.Person, auth *models.Auth, client *models.Client) (*models.Client, error)
}

//AuthService repository
type AuthService struct{}

//PersonRegister method
func (service *AuthService) PersonRegister(person *models.Person, auth *models.Auth, client *models.Client) (*models.Client, error) {
	if !checkUserExistOrNot(auth) {
		return nil, errors.New("User with the requested email exist")
	}
	//insert auth
	_, err := helpers.Mongo().InsertOne(models.AuthCollection, auth)
	if err != nil {
		return nil, err
	}
	//insert person
	_, err = helpers.Mongo().InsertOne(models.PersonCollection, person)
	if err != nil {
		return nil, err
	}
	//insert client
	_, err = helpers.Mongo().InsertOne(models.ClientCollection, client)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func checkUserExistOrNot(auth *models.Auth) bool {
	query := bson.M{
			Key: "value",
			Value: auth.Value,
	}
	existUser, err := helpers.Mongo().FindOne(models.AuthCollection,query)

	if err !=nil
}
