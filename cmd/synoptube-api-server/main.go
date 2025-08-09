package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// The summarizeHandler function is responsible for handling requests to the /summarize endpoint.
// For now, it just logs the request and sends a simple response.
// In the future, this is where you would call the core logic from your pkg/synoptube package.
func summarizeHandler(w http.ResponseWriter, r *http.Request) {
	// Log the incoming request.
	log.Printf("Received request from %s for URL %s", r.RemoteAddr, r.URL.Path)

	// We'll only respond to POST requests.
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// For a real application, you would parse the request body here
	// to get the YouTube URL or playlist ID to summarize.
	// For this example, we'll just acknowledge the request.

	// Write a simple success message back to the client.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "Summarization request received. Processing...")
}

func main() {
	// Set the port for the server to listen on.
	// We'll use an environment variable if available, otherwise default to 8080.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	serverAddr := ":" + port

	// Register our handler function for the /summarize endpoint.
	http.HandleFunc("/summarize", summarizeHandler)

	// Start the HTTP server.
	log.Printf("Synoptube API server starting on http://localhost%s", serverAddr)
	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
