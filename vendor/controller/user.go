package controller

import (
	"encoding/json"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"time"
	"utils"
)

// GetUsers : get the list of all the user
func GetUsers(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	tokenInfo, err := utils.ValidateJWT(token)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid authorization"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tokenInfo)
}

// UserDetail :
type UserDetail struct {
	Username   string
	Authorized bool
}

// UserLogin :
func UserLogin(w http.ResponseWriter, r *http.Request) {
	claims := jwt.MapClaims{}

	//  create sample user to create token
	userData := UserDetail{
		Username:   "Test",
		Authorized: true,
	}

	claims["authorized"] = userData.Authorized
	claims["username"] = userData.Username
	claims["exp"] = time.Now().Add(time.Hour * 1)

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_TOKEN")))
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"success": "false", "error": "Could not generate token"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"success": "true", "token": token})
}
