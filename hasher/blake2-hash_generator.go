package hasher

import (
    "fmt"
    "golang.org/x/crypto/blake2b"
    "io"
    "os"
)

func CalculateBLAKE2Hash(filePath string) (string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return "", err
    }
    defer file.Close()

    hash, err := blake2b.New256(nil)
    if err != nil {
        return "", err
    }

    _, err = io.Copy(hash, file)
    if err != nil {
        return "", err
    }

    hashBytes := hash.Sum(nil)
    hashString := fmt.Sprintf("%x", hashBytes)
    return hashString, nil
}
