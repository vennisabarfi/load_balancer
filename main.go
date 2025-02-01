package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", MainServer).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
