package rand

import (
	"crypto/rand"
)

/*
Seed - Create a byte slice of fixed size and insert randomness to it
*/
func Seed(length int) ([]byte, error) {
	seed := make([]byte, length)

	_, err := rand.Read(seed)
	if err != nil {
		return make([]byte, 0), err
	}

	return seed, nil
}
