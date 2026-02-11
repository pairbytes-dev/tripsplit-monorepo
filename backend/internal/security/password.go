package security

import (
	"golang.org/x/crypto/bcrypt"
)

const passwordCost = bcrypt.DefaultCost

func HashPassword(plain string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(plain), passwordCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
