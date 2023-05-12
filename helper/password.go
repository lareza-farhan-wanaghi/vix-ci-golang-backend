package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func UnsafeHashPassword(pw string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pw), 12)
	return string(hash)
}

func HashPassword(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), 12)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ValidatePassword(hash, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
}
