package utils

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"os"
	"strings"
)

// ValidateJWT :
func ValidateJWT(token string) (interface{}, error) {
	tokenDetail := strings.Split(token, " ")

	if len(tokenDetail) != 2 {
		return "", fmt.Errorf("Invalid token")
	}

	tokenString := tokenDetail[1]

	tokenData, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_TOKEN")), nil
	})

	return tokenData.Claims, err
}
