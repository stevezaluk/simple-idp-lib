package user

import "github.com/spf13/viper"

/*
PasswordHashingParams - A structure representing password hashing values
that can be used to modify the way that simple-idp hashes user passwords
*/
type PasswordHashingParams struct {
	// saltLength - The length of the salt to generate when hashing a password
	saltLength int

	// keyLength - The length of the resulting password hash
	keyLength int

	// iterations - The number of times the result should be re-hashed
	iterations int

	// parallelism - The number of threads Argon2 should use when generating a password hash
	parallelism int

	// memory - The max amount of memory Argon2 can use when hashing a password
	memory int
}

/*
NewHashingParams - Constructor for PasswordHashingParams
*/
func NewHashingParams(saltLength int, keyLength int, iterations int, parallelism int, memory int) *PasswordHashingParams {
	return &PasswordHashingParams{
		saltLength:  saltLength,
		keyLength:   keyLength,
		iterations:  iterations,
		parallelism: parallelism,
		memory:      memory,
	}
}

/*
NewHashingParamsFromConfig - A wrapper around NewHashingParams that fills values in from viper
*/
func NewHashingParamsFromConfig() *PasswordHashingParams {
	return NewHashingParams(
		viper.GetInt("password_hashing.salt_length"),
		viper.GetInt("password_hashing.key_length"),
		viper.GetInt("password_hashing.iterations"),
		viper.GetInt("password_hashing.parallelism"),
		viper.GetInt("password_hashing.memory"),
	)
}
