package helper

import (
	"golang.org/x/crypto/bcrypt"
)

// UnsafeHashPassword returns the bcrypt hash of a string
func UnsafeHashPassword(pw string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pw), 12)
	return string(hash)
}

// HashPassword returns the bcrypt hash of a string and possible error while hashing
func HashPassword(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), 12)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ValidatePassword returns an error if the hash of the plain text is not equal to the provided hash
func ValidatePassword(hash, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
}
