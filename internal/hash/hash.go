package hash

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword securely hashes a password using bcrypt
func HashPassword(p string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}
