// id_generator.go
package idgen

import (
	"math/rand"
	"time"
)

const (
	// Define the character set for the generated IDs.
	chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	// Define the desired length of the generated IDs.
	idLength = 8
)

// Initialize the random number generator.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateID generates a short and readable ID.
func GenerateID() string {
	id := make([]byte, idLength)
	for i := 0; i < idLength; i++ {
		id[i] = chars[rand.Intn(len(chars))]
	}
	return string(id)
}
