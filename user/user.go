package user

import (
	"github.com/stevezaluk/simple-idp-lib/metadata"
)

/*
User - A structure representing a user accessing an API
*/
type User struct {
	// Metadata - General metadata for the structure
	Metadata *metadata.Metadata `json:"metadata"`

	// Username - A distinct name for the user. Cannot conflict with other usernames
	Username string `json:"username"`

	// Email - An email that the user can be contacted at
	Email string `json:"email"`

	// EmailVerified - A boolean value describing if the user has validated there email
	EmailVerified bool `json:"email_verified"`

	// Roles - A list of role ID's that the user is assigned
	Roles []string `json:"roles"`

	// Permissions - A list of permission Id's that the user is assigned
	Permissions []string `json:"permissions"`

	// Applications - A list of application Id's that the user is authorized to access
	Applications []string `json:"applications"`
}

/*
NewUser - A constructor for the User structure
*/
func NewUser(username string, email string, password string) (*User, error) {
	meta, err := metadata.NewMetadata()
	if err != nil {
		return nil, err
	}

	return &User{
		Metadata:      meta,
		Username:      username,
		Email:         email,
		EmailVerified: false,
	}, nil
}
