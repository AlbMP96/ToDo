package auth

import (
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
var passwordRegex = regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*]{8,}$`)

func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func IsValidPassword(password string) bool {
	return passwordRegex.MatchString(password)
}
