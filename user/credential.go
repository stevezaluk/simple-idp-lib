package user

import (
	"crypto/subtle"
	"encoding/base64"
	"github.com/spf13/viper"
	"github.com/stevezaluk/simple-idp-lib/rand"
	"golang.org/x/crypto/argon2"
)

/*
HashingParameters - A structure representing the parameters that should be used
when generating argon2 user password hashes. These can be modified in the simple-idp's
configuration file to improve runtime complexity on resource limited machines
*/
type HashingParameters struct {
	// KeyLength - The length of the resulting password hash
	KeyLength uint32 `json:"key_len"`

	// SaltLength - The length of the salt used for hashing the password
	SaltLength uint32 `json:"salt_len"`

	// Time - The number of iterations used to hash the password
	Time uint32 `json:"time"`

	// Memory - The bytes of memory that Argon2 is allowed to use when hashing passwords
	Memory uint32 `json:"memory"`

	// Threads - The number of go-routines to use when hashing a password
	Threads uint8 `json:"threads"`
}

/*
NewHashingParameters - A constructor for HashingParameters
*/
func NewHashingParameters(keyLength uint32, saltLength uint32, time uint32, memory uint32, threads uint8) *HashingParameters {
	return &HashingParameters{
		KeyLength:  keyLength,
		SaltLength: saltLength,
		Time:       time,
		Memory:     memory,
		Threads:    threads,
	}
}

/*
NewHashingParametersFromConfig - A wrapper around NewHashingParameters that fills in parameters
from Viper
*/
func NewHashingParametersFromConfig() *HashingParameters {
	return NewHashingParameters(
		uint32(viper.GetInt("argon.key_len")),
		uint32(viper.GetInt("argon.salt_len")),
		uint32(viper.GetInt("argon.time")),
		uint32(viper.GetInt("argon.memory")),
		uint8(viper.GetInt("argon.threads")),
	)
}

type Credentials struct {
	// Params - The HashingParameters that were used to generate the hash
	Params *HashingParameters

	// Salt - The base64 encoded versionsalt that is used when hashing the password
	Salt string

	// Key - The base64 encoded version of the password hash
	Key string
}

/*
NewCredentials -
*/
func NewCredentials(password string, params *HashingParameters) (*Credentials, error) {
	salt, err := rand.Seed(int(params.SaltLength))
	if err != nil {
		return nil, err
	}

	key := argon2.IDKey([]byte(password), salt, params.Time, params.Memory, params.Threads, params.KeyLength)

	return &Credentials{
		Params: params,
		Salt:   base64.URLEncoding.EncodeToString(salt),
		Key:    base64.URLEncoding.EncodeToString(key),
	}, nil
}

/*
ValidateCredential - Validates if two secrets are the same. If they are then user
has entered the correct password
*/
func (credential *Credentials) ValidateCredential(password string) (bool, error) {
	/*
		Credentials are always stored as URL Safe Base64 encoded strings,
		so the salt needs to be decoded here before we can compare
	*/
	decodedSalt, err := base64.URLEncoding.DecodeString(credential.Salt)
	if err != nil {
		return false, err
	}

	/*
		Hashing the password parameter here so that we can compare it with what is stored
		in the credential structure. The same salt needs to be used here, as using a salt
		generates a different hash even if the passwords are the same
	*/
	key := argon2.IDKey(
		[]byte(password),
		decodedSalt,
		credential.Params.Time,
		credential.Params.Memory,
		credential.Params.Threads,
		credential.Params.KeyLength,
	)

	/*
		Just like with the salt, we are going to decode the credential here. Encoding
		credentials to base64 strings makes them slightly easier to represent, however
		you now need to be responsible for encoding/decoding them.
	*/
	decodedKey, err := base64.URLEncoding.DecodeString(credential.Key)
	if err != nil {
		return false, err
	}

	/*
		We are using subtle.CompareTimeCompare here as this can prevent timing based attacks
		when comparing two strings. Timing based attacks are a problem when comparing secrets
		of differing lengths, as this can expose information about the compared secrets.

		The return value of subtle.CompareTimeCompare is an integer representing its value. If 1,
		they match, otherwise they don't.
	*/
	ok := subtle.ConstantTimeCompare(key, decodedKey)
	if ok == 1 {
		return true, nil
	}

	return false, nil
}
