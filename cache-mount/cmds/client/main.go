package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// URL of the server endpoint
	url := "http://localhost:8080/status"

	// Make the HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	// Print the response body
	fmt.Printf("Response from server: %s\n", body)
}
