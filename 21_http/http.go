package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Tudo ok!!"))
}

func main() {
	// definindo handlers das rotas
	http.HandleFunc("/home", home)

	// abrindo um servidor conexão com servidor para um servidor local
	log.Fatal(http.ListenAndServe(":8080", nil))
}
