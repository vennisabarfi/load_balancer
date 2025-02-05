package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// main server
	r := mux.NewRouter()

	r.HandleFunc("/", LoadBalancer).Methods("GET")

	srv1 := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	s := mux.NewRouter()
	s.HandleFunc("/", BackendServer).Methods("GET")

	srv2 := &http.Server{
		Handler: s,
		Addr:    "127.0.0.1:3000",
		// server timeouts
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// second server. set up a goroutine to use concurrency

	go func() {
		err := srv1.ListenAndServe()
		if err != nil {
			fmt.Println("Error with first server", err)
		}
	}()

	// second server starting
	err := srv2.ListenAndServe()
	if err != nil {
		fmt.Println("Error with second server", err)
	}

}
