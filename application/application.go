package application

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/stevezaluk/simple-idp-lib/metadata"
	"github.com/stevezaluk/simple-idp-lib/rand"
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
	// Metadata - General metadata for the structure
	Metadata *metadata.Metadata `json:"metadata"`

	// Name - The name of the Application
	Name string `json:"name"`

	// GrantType - The OAuth grant type that you want to assign to this application
	GrantType []GrantType `json:"grant_type"`

	// ClientID - A base64 random string representing the clientId
	ClientID string `json:"client_id"`

	// ClientSecret - A 256-bit random string
	ClientSecret string `json:"client_secret"`
}

/*
New - A constructor for the Application struct
*/
func New(name string, grantType []GrantType) (*Application, error) {
	meta, err := metadata.New()
	if err != nil {
		return nil, err
	}

	app := &Application{
		Metadata:  meta,
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
