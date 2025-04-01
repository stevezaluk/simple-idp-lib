package api

import "github.com/stevezaluk/simple-idp-lib/uuid"

/*
Scope - A user defined permission that can restrict access to resources. Attached at the API level
*/
type Scope struct {
	// Id - A UUID acting as a unique identifier for the Scope
	Id string `json:"id"`

	// Name - The name of the permission scope. Should follow the following format: action:resource
	Name string `json:"identifier"`

	// Description - A description for what the scope does
	Description string `json:"description"`
}

/*
NewScope - A constructor for the Scope object
*/
func NewScope(name string, description string) (*Scope, error) {
	identifier, err := uuid.GenerateUUID()
	if err != nil {
		return nil, err
	}

	return &Scope{
		Id:          identifier,
		Name:        name,
		Description: description,
	}
}
