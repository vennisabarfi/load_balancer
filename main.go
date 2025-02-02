package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// main server
	r := mux.NewRouter()
	r.HandleFunc("/", MainServer).Methods("GET")

	s := mux.NewRouter()
	s.HandleFunc("/", SecondServer).Methods("GET")

	// second server. set up a goroutine to use concurrency

	go func() {
		err := http.ListenAndServe(":8080", r)
		if err != nil {
			fmt.Println("Error with first server", err)
		}
	}()

	// second server starting
	// log.Fatal(http.ListenAndServe(":3000", s))
	err := http.ListenAndServe(":3000", s)
	if err != nil {
		fmt.Println("Error with second server", err)
	}

}
