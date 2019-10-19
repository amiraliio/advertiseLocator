package repositories

import (
	"github.com/amiraliio/advertiselocator/helpers"
	"github.com/amiraliio/advertiselocator/models"
)

//AuthRepository interface
type AuthRepository interface {
	PersonRegister(person *models.Person, auth *models.Auth, client *models.AuthClient) (*models.Person, error)
}

//PersonRegister method
func PersonRegister(person *models.Person, auth *models.Auth, client *models.AuthClient) (*models.Person, error) {
	//insert person
	personID, err := helpers.Mongo().InsertOne(models.PersonCollection, person)
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
	person.ID = personID
	return person, nil
}
