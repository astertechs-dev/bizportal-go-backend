package token_util

import (
	"fmt"

	jwt "github.com/golang-jwt/jwt/v4"
)

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(
		requestToken,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		},
	)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractIDFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(
		requestToken,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		},
	)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return "", fmt.Errorf("invalid Token")
	}

	return claims["id"].(string), nil
}
