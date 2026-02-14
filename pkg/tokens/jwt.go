package tokens

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kautsarhasby/go-messaging-app/pkg/env"
)

type TokenClaims struct {
	Username string
	Fullname string
	jwt.RegisteredClaims
}

var TokenType = map[string]time.Duration{
	"token":        5 * time.Minute,
	"refreshToken": (24 * time.Hour),
}

func GenerateToken(username, fullname string, tokenType string) (string, error) {
	secretKey := []byte(env.GetEnv("APP_SECRET", ""))

	tokenClaims := TokenClaims{
		Username: username,
		Fullname: fullname,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    env.GetEnv("APP_NAME", ""),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenType[tokenType])),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	resultToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("error %v", err)
	}

	return resultToken, nil
}
