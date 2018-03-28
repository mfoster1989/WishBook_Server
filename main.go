// main.go
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/mfoster1989/WishBook_Server/messages"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := messages.NewRouter() // create routes
	// these two lines are important in order to allow access from the front-end side to the methods
	allowedHeaders := handlers.AllowedHeaders([]string{"*"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	// launch server with CORS validations
	http.ListenAndServe(":"+port,
		handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router))
}
