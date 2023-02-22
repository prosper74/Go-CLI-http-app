package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
)

// The client code makes an HTTP GET request to the server specified by the "-server" command-line argument. It then reads the response body and prints it to the console.
func main() {
	server := flag.String("server", "http://localhost:8080", "the server to connect to")
	flag.Parse()

	resp, err := http.Get(*server)
	if err != nil {
		fmt.Printf("Error connecting to server: %s\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return
	}

	fmt.Printf("Response from server: %s\n", string(body))
}
