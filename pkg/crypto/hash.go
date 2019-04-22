package crypto

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var prefix = [...]string{"$2a$", "$2b$"}

// Hash a string using bcrypt.
func Hash(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

// IsEqual if the password matches the given hash.
func IsEqual(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// IsHash checks if the given string is a bcrypt hash.
func IsHash(s string) bool {
	if size := len(s); size < 59 || size > 60 {
		return false
	}
	for _, p := range prefix {
		if strings.HasPrefix(s, p) {
			return true
		}
	}
	return false
}
