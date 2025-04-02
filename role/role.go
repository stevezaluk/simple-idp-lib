package role

import "github.com/stevezaluk/simple-idp-lib/uuid"

/*
Role - A user defined role that be attached to a User. Provides a structure
to attach permissions to
*/
type Role struct {
	// Id - A UUID acting as a unique identifier for the Role
	Id string `json:"id"`

	// Name - The name of the role. Cannot include spaces
	Name string `json:"name"`

	// Description - A description for what the role provides
	Description string `json:"description"`

	// Permissions - A list of permission Id's that the role is assigned
	Permissions []string `json:"permissions"`
}

/*
NewRole -A constructor for the role structure
*/
func NewRole(name string) (*Role, error) {
	identifier, err := uuid.GenerateUUID()
	if err != nil {
		return nil, err
	}

	return &Role{
		Id:   identifier,
		Name: name,
	}, nil
}
