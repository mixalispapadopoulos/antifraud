package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 1. Define the handler (like a route in Flask)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from OpenShift! AUTO v5! You requested: %s\n", r.URL.Path)
	})

	// 2. Start the server on port 8080
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
