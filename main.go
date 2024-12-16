package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/simple-golang-api/user"
)

func main() {
	// load env variables
	godotenv.Load(".env")

	route := gin.Default()
	// Initiate routes and setup

	user.APIRoutes(route)

	// Bind to a port and pass our router in
	fmt.Println("Running server")
	log.Fatal(http.ListenAndServe(":8080", route))
}
