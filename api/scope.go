package api

/*
Scope - A user defined permission that can restrict access to resources. Attached at the API level
*/
type Scope struct {
	// Identifier - The identifier of the permission scope. Should follow the following format: action:resource
	Identifier string `json:"identifier"`

	// Description - A description for what the scope does
	Description string `json:"description"`
}

/*
NewScope - A constructor for the Scope object
*/
func NewScope(identifier string, description string) *Scope {
	return &Scope{
		Identifier:  identifier,
		Description: description,
	}
}
