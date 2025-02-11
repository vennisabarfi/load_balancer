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

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	s1 := mux.NewRouter()
	s1.HandleFunc("/", FirstServer).Methods("GET")

	srv1 := &http.Server{
		Handler: s1,
		Addr:    "127.0.0.1:3000",
		// server timeouts
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	s2 := mux.NewRouter()
	s2.HandleFunc("/", SecondServer).Methods("GET")

	srv2 := &http.Server{
		Handler: s2,
		Addr:    "127.0.0.1:3001",
		// server timeouts
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	s3 := mux.NewRouter()
	s3.HandleFunc("/", ThirdServer).Methods("GET")

	srv3 := &http.Server{
		Handler: s3,
		Addr:    "127.0.0.1:3002",
		// server timeouts
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// second server. set up a goroutine to use concurrency. update to wait group.

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			fmt.Println("Error with load balancer server", err)
		}
	}()
	go func() {
		err := srv1.ListenAndServe()
		if err != nil {
			fmt.Println("Error with first server", err)
		}
	}()
	go func() {
		err := srv2.ListenAndServe()
		if err != nil {
			fmt.Println("Error with second server", err)
		}
	}()

	// second server starting
	err := srv3.ListenAndServe()
	if err != nil {
		fmt.Println("Error with third server", err)
	}

}
