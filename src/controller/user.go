package controller

import (
	"net/http"
)

// GetUsers : get the list of all the user
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This endpoint will provide the list of the users"))
}

