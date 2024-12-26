package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	port = 8080
)

func delayHandler(w http.ResponseWriter, r *http.Request) {
	// Extract delay value from path
	parts := strings.Split(r.URL.Path, "/")
	// Find the delay value after "delay" in the path
	var delayStr string
	for i := 0; i < len(parts)-1; i++ {
		if parts[i] == "delay" && i+1 < len(parts) {
			delayStr = parts[i+1]
			break
		}
	}

	if delayStr == "" {
		http.Error(w, "Invalid path. Use /delay/{seconds}", http.StatusBadRequest)
		return
	}

	seconds, err := strconv.Atoi(delayStr)
	if err != nil {
		http.Error(w, "Invalid delay value. Must be a number.", http.StatusBadRequest)
		return
	}

	if seconds < 0 {
		http.Error(w, "Delay value must be non-negative.", http.StatusBadRequest)
		return
	}

	// Log the incoming request
	log.Printf("Received request for %d second delay", seconds)

	// Sleep for the specified duration
	time.Sleep(time.Duration(seconds) * time.Second)

	// Send response
	w.Header().Set("Content-Type", "application/json")
	response := fmt.Sprintf(`{"message": "Delayed response after %d seconds"}`, seconds)
	fmt.Fprint(w, response)
}

func healthHandler(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status": "healthy"}`)
}

func main() {
	// Create a new mux for more flexible routing
	mux := http.NewServeMux()
	
	// Register handlers - using "/" to catch all paths
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// If it's a health check request
		if r.URL.Path == "/health" {
			healthHandler(w)
			return
		}
		
		// For all other paths, try to handle as delay request
		if strings.Contains(r.URL.Path, "/delay/") {
			delayHandler(w, r)
			return
		}

		// If no handlers match, return 404
		http.NotFound(w, r)
	})

	// Start server with our custom mux
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Starting delay server on port %d", port)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}