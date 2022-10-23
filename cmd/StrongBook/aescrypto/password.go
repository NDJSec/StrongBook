package aescrypto

import (
	"crypto/rand"
	"crypto/sha256"
)

func generateRandomSalt(saltSize int) []byte {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		panic(err)
	}

	return salt
}

func hashPassword(password string, salt []byte) []byte {
	var passwordBytes = []byte(password)

	var sha256Hasher = sha256.New()

	passwordBytes = append(passwordBytes, salt...)

	sha256Hasher.Write(passwordBytes)

	var hashedPasswordBytes = sha256Hasher.Sum(nil)

	return hashedPasswordBytes
}

func doPasswordsMatch(hashedPassword, currPassword string,
	salt []byte) bool {
	var currPasswordHash = hashPassword(currPassword, salt)

	return hashedPassword == string(currPasswordHash)
}
