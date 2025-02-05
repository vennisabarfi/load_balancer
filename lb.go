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

// type Responses struct {
// 	res string
// }

// main server for load balancer
func LoadBalancer(w http.ResponseWriter, r *http.Request) {
	// extract client ip address from parsed_ip
	clients := getIPAddress(r)
	host := strings.Split(r.Host, ":")[0]
	user_agent := r.UserAgent()

	response := fmt.Sprintf("Received request from %s. \n%s %s \n Host: %s \n User-Agent: %s", clients, r.Method, r.Proto, host, user_agent)
	// w.WriteHeader(http.StatusOK)

	// forward request made to load balancer to backend server

	// create a new request/connection to the backend server
	// serverPort := "3000"
	// requestURL := fmt.Sprintf("http://localhost:%d", serverPort)
	// fmt.Print(requestURL)

	// change this to newrequest with context
	req, err := http.NewRequest("GET", "http://localhost:3000", nil) //method, url and request body
	if err != nil {
		log.Println("Error reaching backend server", err)
	}
	// set user-agent as header
	req.Header.Set("User-Agent", user_agent)

	client := &http.Client{}
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
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
