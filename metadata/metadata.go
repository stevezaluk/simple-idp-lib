package metadata

import (
	"github.com/stevezaluk/simple-idp-lib/uuid"
	"time"
)

/*
Metadata - Provides general metadata for all objects created in simple-idp
*/
type Metadata struct {
	// Id - A UUID acting as a unique identifier for the structure
	Id string `json:"id"`

	// Name - The name of the structure
	Name string `json:"name"`

	// CreationDate - The date that this structure was stored in MongoDB
	CreationDate int64 `json:"creation_date"`

	// ModifiedDate - The date that this structure was last modified
	ModifiedDate int64 `json:"modified_date"`

	// Tags - Arbitrary user defined tags
	Tags map[string]string `json:"tags"`
}

/*
NewMetadata - A constructor for the Metadata structure
*/
func NewMetadata(name string) (*Metadata, error) {
	identifier, err := uuid.GenerateUUID()
	if err != nil {
		return nil, err
	}

	timestamp := time.Now().UTC().UnixNano()

	return &Metadata{
		Id:           identifier,
		Name:         name,
		CreationDate: timestamp,
		ModifiedDate: timestamp,
		Tags:         map[string]string{},
	}, nil
}
