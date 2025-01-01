package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

// get hash
func GetHash(key string) (string) {
	hash := sha256.New()
	hash.Write([]byte(key))
	hashBytes := hash.Sum(nil)

	return hex.EncodeToString(hashBytes)
}