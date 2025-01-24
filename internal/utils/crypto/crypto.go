package crypto

import (
	"crypto/rand"
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

// generate salt
func GenerateSalt(length int) (string, error) {
	salt := make([]byte, length)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	return hex.EncodeToString(salt), nil
}

// hash password with salt
func HashPasswordWithSalt(password, salt string) (string) {
    saltedPassword := password + salt
	hassPassword := sha256.Sum256([]byte(saltedPassword))
	return hex.EncodeToString(hassPassword[:])
}

// compare password with hashed password
func ComparePasswordWithHash(password, salt, hashedPasswordStore string) bool {
	hassPassword := HashPasswordWithSalt(password, salt)
	return hassPassword == hashedPasswordStore
}