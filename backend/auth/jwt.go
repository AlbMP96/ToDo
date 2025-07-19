package auth

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateJwtToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
