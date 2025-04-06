package user

import (
	"errors"
	"fmt"
	"github.com/stevezaluk/simple-idp-lib/server"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var ErrCreateUserFailed = errors.New("failed to create user")

/*
GetUser - Fetch a users metadata using its email address
*/
func GetUser(database *server.Database, email string, excludeCreds bool) (*User, error) {
	var ret User

	exclusion := ""
	if excludeCreds {
		exclusion = "credentials"
	}

	err := database.Find(bson.M{"email": email}, &ret, exclusion)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

/*
CreateUser - Insert a new user into the database, and return any errors that may occur
*/
func CreateUser(database *server.Database, user *User, password string, params *HashingParameters) error {
	_, err := GetUser(database, user.Email, true)
	if !errors.Is(err, mongo.ErrNoDocuments) {
		return fmt.Errorf("%w: User already exists", ErrCreateUserFailed)
	}

	creds, err := NewCredentials(password, params)
	if err != nil {
		return err
	}

	user.Credentials = creds
	err = database.Insert(user)
	if err != nil {
		return err
	}

	return nil
}
