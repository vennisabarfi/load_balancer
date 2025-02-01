package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/netip"
)

type Response struct {
	IPAddress netip.Addr //store ip addresses
	Method    string     //request method eg GET
	Host      string
}

func getIPAddress(r *http.Request) string {
	// retrive ip address from request
	ip := r.RemoteAddr
	parsed_ip, _, err := net.SplitHostPort(ip)
	if err != nil {
		log.Print("Error parsing ip address of client")
	}
	return parsed_ip
}

// main server for load balancer
func MainServer(w http.ResponseWriter, r *http.Request) {
	// extract client ip address from parsed_ip
	client := getIPAddress(r)
	fmt.Print(client)
	fmt.Println("Hello World")

}
