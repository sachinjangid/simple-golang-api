package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/simple-golang-api/src"

	"github.com/gorilla/mux"
)

func main() {

	// Initiate routes and setup
	r := mux.NewRouter()
	src.APIRoutes(r)

	// Bind to a port and pass our router in
	fmt.Println("Running server")
	log.Fatal(http.ListenAndServe(":8000", r))
}
