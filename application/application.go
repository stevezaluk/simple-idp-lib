package application

import "github.com/stevezaluk/simple-idp-lib/uuid"

type GrantType string

const (
	ClientCredentials     GrantType = "client_credentials"
	AuthorizationCodePKCE GrantType = "authorization_code"
)

/*
Application - A user defined application. Users will define these and authorize there applications
to use API's that are defined
*/
type Application struct {
	// Id - A UUID acting as a unique identifier for the Application
	Id string `json:"id"`

	// Name - The name of the Application
	Name string `json:"name"`

	// GrantType - The OAuth grant type that you want to assign to this application
	GrantType GrantType `json:"grantType"`

	// ClientID - A base64 random string representing the clientId
	ClientID string `json:"clientId"`

	// ClientSecret - A SHA256 hash of a randomly generated string
	ClientSecret string `json:"clientSecret"`
}

/*
NewApplication - A constructor for the Application struct
*/
func NewApplication(name string, grantType GrantType) (*Application, error) {
	identifier, err := uuid.GenerateUUID()
	if err != nil {
		return nil, err
	}

	return &Application{
		Id:        identifier,
		Name:      name,
		GrantType: grantType,
	}, nil
}
