package utils

import (
	"crypto/rand"
	"fmt"
)

// GenerateID creates a random 6-digit ID
func GenerateID() (string, error) {
	b := make([]byte, 3) // 3 bytes make 6 hex characters
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}
