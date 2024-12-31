// Backend entry point
package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/email-verifier/handlers"
)

func main() {
	// Set up the HTTP server with a route for email verification
	http.HandleFunc("/verify", handlers.VerifyHandler)

	// Start the server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}