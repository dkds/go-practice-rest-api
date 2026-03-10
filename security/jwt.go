package security

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const signingKey = "s3cr3t"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	signedToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", errors.New("Could not generate token, " + err.Error())
	}

	return signedToken, nil
}
