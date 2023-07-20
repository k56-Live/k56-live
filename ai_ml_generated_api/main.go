package main

import (
	"fmt"
	"net/http"

	"github.com/project_root/ai_ml_generated_api/handlers"
)

func main() {
	// Set up the API routes and handlers
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/users", handlers.UsersHandler)
	http.HandleFunc("/data", handlers.DataHandler)

	// Start the API server
	fmt.Println("AI/ML API server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
