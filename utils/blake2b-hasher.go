package utils

import (
	"fmt"
	"io"

	"golang.org/x/crypto/blake2b"
)

func GenerateHash(file io.Reader) (string, error) {
	//hasher is an instance hash writer provided by the blake2b, nil - no key provided
	hasher, err := blake2b.New256(nil)
	if err != nil {
		return "", err
	}
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}
