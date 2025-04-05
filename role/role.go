package role

import (
	"github.com/stevezaluk/simple-idp-lib/metadata"
)

/*
Role - A user defined role that be attached to a User. Provides a structure
to attach permissions to
*/
type Role struct {
	// Metadata - General metadata for the structure
	Metadata *metadata.Metadata `json:"metadata"`

	// Name - The name of the role. Cannot include spaces
	Name string `json:"name"`

	// Description - A description for what the role provides
	Description string `json:"description"`

	// Permissions - A list of permission Id's that the role is assigned
	Permissions []string `json:"permissions"`
}

/*
New -A constructor for the role structure
*/
func New(name string) (*Role, error) {
	meta, err := metadata.New()
	if err != nil {
		return nil, err
	}

	return &Role{
		Metadata: meta,
		Name:     name,
	}, nil
}
