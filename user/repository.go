package user

import (
	"github.com/stevezaluk/simple-idp-lib/server"
	"go.mongodb.org/mongo-driver/v2/bson"
)

/*
GetUser - Fetch a users metadata using its email address
*/
func GetUser(database *server.Database, email string) (*User, error) {
	var ret User

	err := database.Find(bson.M{"email": email}, &ret)

	if err != nil {
		return nil, err
	}

	return &ret, nil
}
