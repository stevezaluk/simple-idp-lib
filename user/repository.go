package user

import (
	"github.com/stevezaluk/simple-idp-lib/server"
	"go.mongodb.org/mongo-driver/v2/bson"
)

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
