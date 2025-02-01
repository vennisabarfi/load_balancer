package main

import (
	"fmt"
	"net/http"
)

// main server for load balancer
func MainServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
}
