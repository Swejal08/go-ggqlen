package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenMetadata struct {
	ID    string
	Email string
}

func SignJwtToken(payload TokenMetadata) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	return signToken(map[string]interface{}{
		"id":    payload.ID,
		"email": payload.Email,
	}, secret, 365*24*time.Hour)
}

func signToken(metaData map[string]interface{}, secret string, ExpiredAt time.Duration) (string, error) {
	claims := jwt.MapClaims{}

	for key, value := range metaData {
		claims[key] = value
	}
	claims["exp"] = time.Now().Add(ExpiredAt).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func VerifyToken(accessToken, secret string) (*TokenMetadata, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}
	return &TokenMetadata{ID: fmt.Sprint(claims["id"]), Email: fmt.Sprint(claims["email"])}, nil
}
