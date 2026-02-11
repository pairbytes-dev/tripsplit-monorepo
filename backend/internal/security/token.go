package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("maximusdecimusmeridiusg")

func GenerateToken(userID int64, email string) (string, error) {
	claims := jwt.MapClaims{
		"sub":   userID,
		"email": email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
