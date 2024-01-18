package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

// Server interface definition
type Server interface {
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request)
}

// simpleServer struct implementation
type simpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

// LoadBalancer struct definition
type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
}

// newSimpleServer creates a new simpleServer instance
func newSimpleServer(addr string) *simpleServer {
	serverURL, err := url.Parse(addr)
	handleError(err)
	return &simpleServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverURL),
	}
}

// NewLoadBalancer creates a new LoadBalancer instance
func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}
}

// Address returns the address of a simpleServer
func (ss *simpleServer) Address() string {
	return ss.addr
}

// IsAlive returns true for a simpleServer (dummy implementation)
func (s *simpleServer) IsAlive() bool {
	return true
}

// Serve proxies the request to the underlying server
func (s *simpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)
}

// getNextAvailableServer returns the next available server in round-robin fashion
func (lb *LoadBalancer) getNextAvailableServer() Server {
	server := lb.servers[lb.roundRobinCount]
	for !server.IsAlive() {
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}
	lb.roundRobinCount++
	return server
}

// serveProxy proxies the request to the next available server
func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, r *http.Request) Server {
	targetServer := lb.getNextAvailableServer()
	fmt.Printf("forwarding request to address %q\n", targetServer.Address())
	targetServer.Serve(rw, r)
	return targetServer
}

func main() {
	servers := []Server{
		newSimpleServer("https://www.facebook.com"),
		newSimpleServer("https://www.bing.com"),
		newSimpleServer("https://www.duckduckgo.com"),
	}
	lb := NewLoadBalancer("8000", servers)

	handleRedirect := func(rw http.ResponseWriter, r *http.Request) {
		lb.serveProxy(rw, r)
	}

	http.HandleFunc("/", handleRedirect)

	fmt.Printf("serving requests at 'localhost:%s'\n", lb.port)
	http.ListenAndServe(":"+lb.port, nil)
}
