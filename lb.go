package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/netip"
	"os"
	"strings"
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
	host := strings.Split(r.Host, ":")[0]
	user_agent := r.UserAgent()
	response := fmt.Sprintf("Received request from %s. \n%s %s \n Host: %s \n User-Agent: %s", client, r.Method, r.Proto, host, user_agent)
	io.WriteString(os.Stdout, response)
}
