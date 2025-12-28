package validator

import (
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// IsValidEmail checks if the email format is valid
func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

// IsValidUsername checks if username meets requirements
func IsValidUsername(username string) bool {
	if len(username) < 3 || len(username) > 50 {
		return false
	}
	// Allow alphanumeric, underscore, and hyphen
	for _, char := range username {
		if !((char >= 'a' && char <= 'z') ||
			(char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') ||
			char == '_' || char == '-') {
			return false
		}
	}
	return true
}

// IsValidPassword checks if password meets minimum requirements
func IsValidPassword(password string) bool {
	return len(password) >= 8
}

// IsEmpty checks if a string is empty or only whitespace
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

// IsValidULID checks if a string is a valid ULID format
func IsValidULID(id string) bool {
	if len(id) != 26 {
		return false
	}
	// ULID is base32 encoded
	for _, char := range id {
		if !((char >= '0' && char <= '9') ||
			(char >= 'A' && char <= 'Z')) {
			return false
		}
	}
	return true
}
