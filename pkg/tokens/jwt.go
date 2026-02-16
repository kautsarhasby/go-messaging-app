package tokens

import (
	"context"
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

func GenerateToken(ctx context.Context, username, fullname string, tokenType string) (string, error) {
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

func ValidateToken(ctx context.Context, token string) (*TokenClaims, error) {
	secretKey := []byte(env.GetEnv("APP_SECRET", ""))
	jwtToken, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("failed to validate method JWT %v", t.Header["alg"])

		}

		return secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to parse %v", err)
	}

	claimToken, ok := jwtToken.Claims.(*TokenClaims)
	if !ok && !jwtToken.Valid {
		return nil, fmt.Errorf("Token Invalid")
	}

	return claimToken, nil

}
