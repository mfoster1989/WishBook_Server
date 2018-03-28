// main.go
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/mfoster1989/WishBook_Server/messages"
)

// func Cleaner(w http.ResponseWriter, r *http.Request) {
// 	// Read body
// 	b, err := ioutil.ReadAll(r.Body)
// 	defer r.Body.Close()
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}

// 	// Unmarshal
// 	var msg Message
// 	err = json.Unmarshal(b, &msg)
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}

// 	output, err := json.Marshal(msg)
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}
// 	w.Header().Set("content-type", "application/json")
// 	w.Write(output)
// }

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := messages.NewRouter() // create routes
	// these two lines are important in order to allow access from the front-end side to the methods
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "OPTIONS"})
	// launch server with CORS validations
	log.Fatal(http.ListenAndServe(":"+port,
		handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)))
}
