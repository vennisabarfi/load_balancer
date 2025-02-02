// backend server
// receives forwarded result from load balancer
package main

import (
	"fmt"
	"net/http"
)

func BackendServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
}
