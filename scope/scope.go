package scope

import (
	"github.com/stevezaluk/simple-idp-lib/metadata"
)

/*
Scope - A user defined permission that can restrict access to resources. Attached at the API level
*/
type Scope struct {
	// Metadata - General metadata for the structure
	Metadata *metadata.Metadata `json:"metadata" bson:"metadata"`

	// Name - The name of the permission scope. Should follow the following format: action:resource
	Name string `json:"name" bson:"name"`

	// Description - A description for what the scope does
	Description string `json:"description" bson:"description"`
}

/*
New - A constructor for the Scope object
*/
func New(name string, description string) (*Scope, error) {
	meta, err := metadata.New()
	if err != nil {
		return nil, err
	}

	return &Scope{
		Metadata:    meta,
		Name:        name,
		Description: description,
	}, nil
}
