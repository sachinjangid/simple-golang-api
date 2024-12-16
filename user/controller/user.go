package controller

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/simple-golang-api/user/model"
	"github.com/simple-golang-api/utils"
)

// GetUsers : get the list of all the user
func GetUsers(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	tokenInfo, err := utils.ValidateJWT(token)
	ctx.Writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusUnauthorized, map[string]string{"success": "false", "error": "Invalid authorization"})
		return
	}

	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.JSON(http.StatusUnauthorized, map[string]any{"success": "false", "data": tokenInfo})
}

// UserDetail :
type UserDetail struct {
	Username   string
	Authorized bool
}

// UserLogin :
func UserLogin(ctx *gin.Context) {
	logger := log.Default()
	var user model.User

	if err := utils.Bind(ctx, &user); err != nil {
		logger.Println(err)
		ctx.JSON(http.StatusBadRequest, map[string]string{"success": "false", "error": "Invaid request"})
		return
	}

	if user.Username != "test" || user.Password != "test" {
		ctx.JSON(http.StatusBadRequest, map[string]string{"success": "false", "error": "Invaid creds"})
		return
	}

	//  create sample user to create token
	userData := UserDetail{
		Username:   "test",
		Authorized: true,
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = userData.Authorized
	claims["username"] = userData.Username
	claims["exp"] = time.Now().Add(time.Hour * 1)

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_TOKEN")))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"success": "false", "error": "Could not generate token"})
		return
	}

	ctx.JSON(http.StatusBadRequest, map[string]string{"success": "true", "token": token})
}
