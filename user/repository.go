package user

import (
	"errors"
	"fmt"
	"github.com/stevezaluk/simple-idp-lib/server"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var ErrUserAlreadyExists = errors.New("user: User already exists")
var ErrUserDoesNotExist = errors.New("user: Does not exist")
var ErrFetchUserFailed = errors.New("user: Failed to fetch user")
var ErrCreateUserFailed = errors.New("user: Failed to create user")
var ErrDeleteUserFailed = errors.New("user: Failed to delete user")

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
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrUserDoesNotExist
		}
		return nil, fmt.Errorf("%w: (%s)", ErrFetchUserFailed, err)
	}

	return &ret, nil
}

/*
CheckUserExists - Check to see if a user already exists in the database
*/
func CheckUserExists(database *server.Database, email string) (bool, error) {
	ok, err := database.Exists(bson.M{"email": email})
	if err != nil {
		return false, err
	}

	if ok {
		return true, nil
	}

	return false, nil
}

/*
CreateUser - Insert a new user into the database, and return any errors that may occur
*/
func CreateUser(database *server.Database, user *User, password string, params *HashingParameters) error {
	ok, err := CheckUserExists(database, user.Email)
	if err != nil {
		return fmt.Errorf("%w: (%s)", ErrCreateUserFailed, err)
	}

	if ok {
		return ErrUserAlreadyExists
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
	if err != nil {
		return fmt.Errorf("%w: (%s)", ErrDeleteUserFailed, err)
	}

	if !ok {
		return ErrUserDoesNotExist
	}

	err = database.Delete(bson.M{"email": email})
	if err != nil {
		return err
	}

	return nil
}
