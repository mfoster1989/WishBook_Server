package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

// Message object
type Message struct {
	ID      int64  `json:"id"`
	Token   int64  `bson:"token" json:"token"`
	Message string `bson:"message" json:"message"`
}

// Create Books slice
var messages []Message

// AllMessagesEndPoint routing
func AllMessagesEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
	// fmt.Fprintln(w, "not implemented yet !")
}

// CreateMessagesEndPoint routing
func CreateMessagesEndPoint(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "not working yet")
	w.Header().Set("Content-Type", "application/json")
	var message Message
	_ = json.NewDecoder(r.body).Decode(&Message)
	book.ID = 

}

// UpdateMessagesEndPoint routing
func UpdateMessagesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not working yet")
}

// FindMessagesEndPoint routing
func FindMessagesEndPoint(w http.ResponseWriter, r *http.Request) {
	// // fmt.Fprintln(w, "not working yet")
	// w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(r) //get params
	// for _, item := range messages {
	// 	if item.ID == params[id] {
	// 		json.NewEncoder(w).Encode(item)
	// 	}
	// }
	// json.NewEncoder(w).Encode(&Message{})
}

func main() {
	// making sure code is running in development
	fmt.Println("code running")
	// creates new router
	r := mux.NewRouter()

	// I don't want to delete this code

	// allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	// allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	//inserting test data
	messages = append(messages, Message{ID: 1234, Token: 123456789123456789, Message: "This is a test to make sure that my API is working"})

	// CRUD functionality
	r.HandleFunc("/messages", AllMessagesEndPoint).Methods("GET")
	r.HandleFunc("/messages", CreateMessagesEndPoint).Methods("POST")
	r.HandleFunc("/messages/{id}", UpdateMessagesEndPoint).Methods("PUT")
	r.HandleFunc("/messages/{id}", FindMessagesEndPoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
