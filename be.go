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
	fmt.Println("Hello World")

}
