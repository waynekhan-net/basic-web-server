package main

import (
	"fmt"
	"net/http"
)

// A simple handler function
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler) // Register the handler for the root path
	fmt.Println("Server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Start the server
}
