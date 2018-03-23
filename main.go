package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

func AllMessagesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func FindMessagesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not working yet")
}

func CreateMessagesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not working yet")
}

func UpdateMessagesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not working yet")
}

func main() {
	fmt.Println("code running")
	r := mux.NewRouter()
	r.HandleFunc("/messages", AllMessagesEndPoint).Methods("GET")
	r.HandleFunc("/messages", CreateMessagesEndPoint).Methods("POST")
	r.HandleFunc("/messages", UpdateMessagesEndPoint).Methods("PUT")
	r.HandleFunc("/messages/{id}", FindMessagesEndPoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
