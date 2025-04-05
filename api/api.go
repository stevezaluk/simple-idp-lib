package api

import (
	"github.com/stevezaluk/simple-idp-lib/metadata"
	"github.com/stevezaluk/simple-idp-lib/scope"
)

type TokenType string

const (
	RS256 TokenType = "RS256"
	HS256 TokenType = "HS256"
)

/*
API - A user defined application that defines how you can issue tokens
*/
type API struct {
	// Metadata - General metadata for the structure
	Metadata *metadata.Metadata `json:"metadata" bson:"metadata"`

	// Name - The name of the API
	Name string `json:"name" bson:"name"`

	// Audience - The audience identifier that will be sent back in tokens
	Audience string `json:"audience" bson:"audience"`

	// TokenType - The type of JWT's that this API can issue. Can either be RS256 or HS256
	TokenType TokenType `json:"token_type" bson:"token_type"`

	// TokenLifetime - The number of seconds in which a token should expire
	TokenLifetime int `json:"token_lifetime" bson:"token_lifetime"`

	// Permissions - Any permissions that this API can utilize
	Permissions []scope.Scope `json:"permissions" bson:"permissions"`

	// AddPermissions - Determines if permissions should be added to tokens
	AddPermissions bool `json:"add_permissions" bson:"add_permissions"`
}

/*
New - A constructor for the API object
*/
func New(name string, audience string, tokenType TokenType) (*API, error) {
	meta, err := metadata.New()
	if err != nil {
		return nil, err
	}

	return &API{
		Metadata:      meta,
		Name:          name,
		Audience:      audience,
		TokenType:     tokenType,
		TokenLifetime: 86400,
	}, nil
}
