package metadata

import (
	"github.com/google/uuid"
	"time"
)

/*
Metadata - Provides general metadata for all objects created in simple-idp
*/
type Metadata struct {
	// Id - A UUID acting as a unique identifier for the structure
	Id string `json:"id"`

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
func NewMetadata() (*Metadata, error) {
	identifier, err := uuid.NewV6()
	if err != nil {
		return nil, err
	}

	timestamp := time.Now().UTC().UnixNano()

	return &Metadata{
		Id:           identifier.String(),
		CreationDate: timestamp,
		ModifiedDate: timestamp,
		Tags:         map[string]string{},
	}, nil
}
