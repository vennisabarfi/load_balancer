// backend server
// receives forwarded result from load balancer
package main

import (
	"fmt"
	"net/http"
)

func FirstServer(w http.ResponseWriter, r *http.Request) {
	// loadresponse := LoadBalancer(w, r)
	// io.WriteString(os.Stdout, response)
	backendResponse := "Hello From First Backend Server"

	fmt.Println("Hello From First Backend Server")
	w.Write([]byte(backendResponse))

}

func SecondServer(w http.ResponseWriter, r *http.Request) {
	// loadresponse := LoadBalancer(w, r)
	// io.WriteString(os.Stdout, response)
	backendResponse := "Hello From Second Backend Server"

	fmt.Println("Hello From Second Backend Server")
	w.Write([]byte(backendResponse))

}

func ThirdServer(w http.ResponseWriter, r *http.Request) {
	// loadresponse := LoadBalancer(w, r)
	// io.WriteString(os.Stdout, response)
	backendResponse := "Hello From Third Backend Server"

	fmt.Println("Hello From Third Backend Server")
	w.Write([]byte(backendResponse))

}

func HealthServer(w http.ResponseWriter, r *http.Request) {
	// loadresponse := LoadBalancer(w, r)
	// io.WriteString(os.Stdout, response)
	backendResponse := "Healthy Endpoint Yayy"

	fmt.Println("Healthy Endpoint Yayy")
	w.Write([]byte(backendResponse))

}
