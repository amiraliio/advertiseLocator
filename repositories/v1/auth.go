package repositories

import (
	"errors"
	"github.com/amiraliio/advertiselocator/helpers"
	"github.com/amiraliio/advertiselocator/models"
	"go.mongodb.org/mongo-driver/bson"
)

//AuthRepository interface
type AuthRepository interface {
	PersonRegister(person *models.Person, auth *models.Auth, client *models.Client) (*models.Client, error)
}

//AuthService repository
type AuthService struct{}

//PersonRegister method
func (service *AuthService) PersonRegister(person *models.Person, auth *models.Auth, client *models.Client) (*models.Client, error) {
	userExist, err := checkUserExistOrNot(auth)
	if err != nil {
		return nil, err
	}
	if userExist {
		return nil, errors.New("User with the requested email exist")
	}
	//insert auth
	_, err = helpers.Mongo().InsertOne(models.AuthCollection, auth)
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

func checkUserExistOrNot(auth *models.Auth) (bool, error) {
	query := bson.M{"value": auth.Value, "status": models.ActiveStatus, "userType": auth.UserType, "type": auth.Type}
	existUser, err := helpers.Mongo().FindOne(models.AuthCollection, query, new(models.Auth))
	if err != nil {
		return false, err
	}
	if existUser != nil && helpers.IsInstance(existUser, (*models.Auth)(nil)) {
		if existUser.(models.Auth).Value == auth.Value {
			return true, nil
		}
	}
	return false, nil
}
