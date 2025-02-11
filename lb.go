// load balancer
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
)

// type Response struct {
// 	IPAddress netip.Addr //store ip addresses
// 	Method    string     //request method eg GET
// 	Host      string
// }

type Responses struct {
	StatusCode int
	// Header     Header
	Body io.ReadCloser
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

// list of backend servers
var servers = []string{
	"http://localhost:3000",
	"http://localhost:3001",
	"http://localhost:3002",
}

// implement round robin algorithm
//
//	set up index to keep track of the next server to use
var currentServerIndex int
var mu sync.Mutex //help keep concurreny

func roundRobin() string {
	mu.Lock()
	defer mu.Unlock()

	server := servers[currentServerIndex]
	currentServerIndex = (currentServerIndex + 1) % len(servers)

	return server
}

// main server for load balancer
func LoadBalancer(w http.ResponseWriter, r *http.Request) {
	// extract client ip address from parsed_ip
	clients := getIPAddress(r)
	host := strings.Split(r.Host, ":")[0]
	user_agent := r.UserAgent()

	response := fmt.Sprintf("Received request from %s. \n%s %s \n Host: %s \n User-Agent: %s", clients, r.Method, r.Proto, host, user_agent)

	// get next server in round robin algorithm
	serverURL := roundRobin()

	// change this to newrequest with context
	req, err := http.NewRequest("GET", serverURL, nil) //method, url and request body
	if err != nil {
		log.Println("Error reaching backend server", err)
	}
	// set user-agent as header
	req.Header.Set("User-Agent", user_agent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending http request to the backend server", err)
	}
	defer resp.Body.Close()
	// backend server response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading backend server body", err)
	}
	w.Write(body)
	io.WriteString(os.Stdout, string(body))
	// print to server and also write response
	io.WriteString(os.Stdout, response)
	// w.Write([]byte(response))
	// io.WriteString(w, response)

}
