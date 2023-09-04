package idgen

import (
	"math/rand"
	"time"
)

const (
	chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	idLength = 8
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateID() string {
	id := make([]byte, idLength)
	for i := 0; i < idLength; i++ {
		id[i] = chars[rand.Intn(len(chars))]
	}
	return string(id)
}
