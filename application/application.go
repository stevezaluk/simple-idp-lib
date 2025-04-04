package application

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/stevezaluk/simple-idp-lib/rand"
	"github.com/stevezaluk/simple-idp-lib/uuid"
)

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
	GrantType GrantType `json:"grant_type"`

	// ClientID - A base64 random string representing the clientId
	ClientID string `json:"client_id"`

	// ClientSecret - A 256-bit random string
	ClientSecret string `json:"client_secret"`
}

/*
NewApplication - A constructor for the Application struct
*/
func NewApplication(name string, grantType GrantType) (*Application, error) {
	identifier, err := uuid.GenerateUUID()
	if err != nil {
		return nil, err
	}

	app := &Application{
		Id:        identifier,
		Name:      name,
		GrantType: grantType,
	}

	err = app.generateClientCredentials()
	if err != nil {
		return nil, err
	}

	return app, nil
}

/*
generateClientCredentials - Generates a new ClientID and ClientSecret for
the application to use
*/
func (application *Application) generateClientCredentials() error {

	clientIdSeed, err := rand.Seed(64)
	if err != nil {
		return err
	}

	clientSecretSeed, err := rand.Seed(64)
	if err != nil {
		return err
	}

	clientSecret := sha256.New()

	application.ClientID = base64.URLEncoding.EncodeToString(clientIdSeed)
	application.ClientSecret = hex.EncodeToString(clientSecret.Sum(clientSecretSeed))

	return nil
}
