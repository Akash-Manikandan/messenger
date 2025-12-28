package crypto

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

const (
	// Argon2id parameters
	memory      = 64 * 1024 // 64 MB
	iterations  = 3
	parallelism = 2
	saltLength  = 32
	keyLength   = 32
)

// GenerateSalt generates a random salt for password hashing
func GenerateSalt() (string, error) {
	salt := make([]byte, saltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("failed to generate salt: %w", err)
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

// HashPassword hashes a password using Argon2id with the provided salt
func HashPassword(password, saltStr string) (string, error) {
	salt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		return "", fmt.Errorf("failed to decode salt: %w", err)
	}

	hash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, keyLength)
	return base64.StdEncoding.EncodeToString(hash), nil
}

// VerifyPassword verifies a password against a hash using Argon2id
func VerifyPassword(password, hashedPassword, saltStr string) (bool, error) {
	computedHash, err := HashPassword(password, saltStr)
	if err != nil {
		return false, err
	}
	return computedHash == hashedPassword, nil
}

// ValidatePasswordStrength checks if password meets minimum requirements
func ValidatePasswordStrength(password string) error {
	if len(password) <= 4 {
		return fmt.Errorf("password must be longer than 4 characters")
	}
	return nil
}
