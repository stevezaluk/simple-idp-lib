package user

import (
	"errors"
	"fmt"
	"github.com/stevezaluk/simple-idp-lib/server"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var ErrCreateUserFailed = errors.New("failed to create user")
var ErrDeleteUserFailed = errors.New("failed to delete user")

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
CheckUserExists - Check to see if a user already exists in the database
*/
func CheckUserExists(database *server.Database, email string) (bool, error) {
	_, err := GetUser(database, email, true)
	if !errors.Is(err, mongo.ErrNoDocuments) {
		return true, nil
	}

	return false, err
}

/*
CreateUser - Insert a new user into the database, and return any errors that may occur
*/
func CreateUser(database *server.Database, user *User, password string, params *HashingParameters) error {
	ok, err := CheckUserExists(database, user.Email)
	if ok {
		return fmt.Errorf("%w: User already exists", ErrCreateUserFailed)
	} else if err != nil {
		return err
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

/*
DeleteUser - Remove a single user from the database, and return any errors that may occur
*/
func DeleteUser(database *server.Database, email string) error {
	ok, err := CheckUserExists(database, email)
	if !ok {
		return fmt.Errorf("%w: User does not exist", ErrDeleteUserFailed)
	} else if err != nil {
		return err
	}

	err = database.Delete(bson.M{"email": email})
	if err != nil {
		return err
	}

	return nil
}
