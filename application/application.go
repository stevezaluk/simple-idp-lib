package application

import "github.com/stevezaluk/simple-idp-lib/uuid"

/*
Application - A user defined application. Users will define these and authorize there applications
to use API's that are defined
*/
type Application struct {
	// Id - A UUID acting as a unique identifier for the Application
	Id string `json:"id"`

	// Name - The name of the Application
	Name string `json:"name"`
}

/*
NewApplication - A constructor for the Application struct
*/
func NewApplication(name string) (*Application, error) {
	identifier, err := uuid.GenerateUUID()
	if err != nil {
		return nil, err
	}
	return &Application{
		Id:   identifier,
		Name: name,
	}, nil
}
