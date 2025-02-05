// backend server
// receives forwarded result from load balancer
package main

import (
	"fmt"
	"net/http"
)

func BackendServer(w http.ResponseWriter, r *http.Request) {
	// loadresponse := LoadBalancer(w, r)
	// io.WriteString(os.Stdout, response)
	backendResponse := "Hello From Backend Server"

	fmt.Println("Hello From Backend Server")
	w.Write([]byte(backendResponse))

}
