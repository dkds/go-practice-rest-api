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

func parseToken(token string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(
		token,
		func(t *jwt.Token) (any, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("Unexpected signing method")
			}
			return []byte(signingKey), nil
		})
	if err != nil {
		return nil, errors.New("Could not parse token, " + err.Error())
	}
	return parsedToken, nil
}

func ValidateToken(token string) error {
	parsedToken, err := parseToken(token)
	if err != nil {
		return err
	}
	if !parsedToken.Valid {
		return errors.New("Invalid token")
	}

	return nil
}

func ExtractUserIdFromToken(token string) (int64, error) {
	parsedToken, err := parseToken(token)
	if err != nil {
		return 0, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid token claims")
	}

	userId := int64(claims["userId"].(float64))
	// email := claims["email"].(string)

	return userId, nil
}
