package main

import (
	"basic_crud/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", server.CreateUser).Methods("POST")
	router.HandleFunc("/users", server.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", server.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", server.DeleteUser).Methods("DELETE")

	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
