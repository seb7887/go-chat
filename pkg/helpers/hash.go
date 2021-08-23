package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
}

func VerifyPassword(password string, hash string) bool {
	passErr := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return false
	}

	return true
}
