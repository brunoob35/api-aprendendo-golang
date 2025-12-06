package security

import "golang.org/x/crypto/bcrypt"

//go get golang.org/x/crypto/bcrypt

// Hash receives a string and hashes it
// its necessary to decode the password to a slice of bytes and to set a cost for the operation.
// Bcrypt has a default value, this is a balanced value it won't consume too much computing resources.
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// ValidatePassword compares a string password from the request to the hash saved in DB, validating the operation
// and returning if equals.
func ValidatePassword(hashedPassword, stringPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(stringPassword))
}
