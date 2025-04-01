package uuid

import "github.com/google/uuid"

/*
GenerateUUID - Generate a random V6 UUID and return it as a string
*/
func GenerateUUID() (string, error) {
	identifier, err := uuid.NewV6()
	if err != nil {
		return "", err
	}

	return identifier.String(), nil
}
