package utils

import (
	"fmt"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(obj, b)
}
