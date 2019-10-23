package repositories

import (
	"go.mongodb.org/mongo-driver/bson"
	"github.com/amiraliio/advertiselocator/helpers"
	"github.com/amiraliio/advertiselocator/models"
)

//AuthRepository interface
type AuthRepository interface {
	PersonRegister(person *models.Person, auth *models.Auth, client *models.AuthClient) (*models.Person, error)
}

//AuthService repository
type AuthService struct{}

//PersonRegister method
func (service *AuthService) PersonRegister(person *models.Person, auth *models.Auth, client *models.AuthClient) (*models.Person, error) {
	//TODO move this mapping to helpers
	personData, err := bson.Marshal(person)
	if err != nil {
		return nil, err
	}
	var x map[string]interface{}
	err = bson.Unmarshal(personData,&x)
	//insert person
	_, err = helpers.Mongo().InsertOne(models.PersonCollection, x)
	if err != nil {
		return nil, err
	}
	//insert auth
	_, err = helpers.Mongo().InsertOne(models.AuthCollection, auth)
	if err != nil {
		return nil, err
	}
	//insert client
	_, err = helpers.Mongo().InsertOne(models.ClientCollection, client)
	if err != nil {
		return nil, err
	}
	return person, nil
}
