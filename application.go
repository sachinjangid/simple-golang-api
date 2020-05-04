package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/simple-golang-api/vendor"
)

func main() {
	// load env variables
	godotenv.Load(".env")

	// Initiate routes and setup
	r := mux.NewRouter()
	vendor.APIRoutes(r)

	// Bind to a port and pass our router in
	fmt.Println("Running server")
	log.Fatal(http.ListenAndServe(":8000", r))
}
