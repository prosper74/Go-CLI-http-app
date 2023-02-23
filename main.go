package main

import (
	"flag"
	"fmt"
	"net/http"
)

// The code sets up a simple HTTP server that listens for incoming requests on the specified port. When a request is received, it responds with the message "Hello World!".
func main() {
	port := flag.Int("port", 8080, "the port to listen on")
	flag.Parse()

	http.HandleFunc("/", handler)

	fmt.Printf("Server listening on port %d\n", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
