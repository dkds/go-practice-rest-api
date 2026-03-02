package security

import (
	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
)

func GenerateSalt() string {
	return rand.Text()
}

func HashPassword(password, salt string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
