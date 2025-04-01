package api

type TokenType string

const (
	RS256 TokenType = "RS256"
	HS256 TokenType = "HS256"
)

/*
API - A user defined application that defines how you can issue tokens
*/
type API struct {
	// Id - A UUID acting as a unique identifier for the API
	Id string `json:"id"`

	// Name - The name of the API
	Name string `json:"name"`

	// Audience - The audience identifier that will be sent back in tokens
	Audience string `json:"audience"`

	// TokenType - The type of JWT's that this API can issue. Can either be RS256 or HS256
	TokenType TokenType `json:"tokenType"`

	// TokenLifetime - The number of seconds in which a token should expire
	TokenLifetime int `json:"tokenLifetime"`

	// Permissions - Any permissions that this API can utilize
	Permissions []Scope `json:"permissions"`

	// AddPermissions - Determines if permissions should be added to tokens
	AddPermissions bool `json:"addPermissions"`

	// AddRefreshToken - Determines if refresh tokens can be issued for this API
	AddRefreshToken bool `json:"addRefreshToken"`
}

/*
NewAPI - A constructor for the API object
*/
func NewAPI(name string, audience string, tokenType TokenType) *API {
	// identifier still needs to be added

	return &API{
		Name:      name,
		Audience:  audience,
		TokenType: tokenType,
	}
}
