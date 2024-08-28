package models

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

func Sha265Random() string {
	hash := sha256.New()
	hash.Write([]byte(time.Now().String()))

	// Get the hash result
	hashBytes := hash.Sum(nil)

	// Convert the hash bytes to a hexadecimal string
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}
