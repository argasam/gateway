package main

import (
	"fmt"
	"gateway/service"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/gateway/service_a/", func(w http.ResponseWriter, r *http.Request) {
		const serviceAURL = "http://localhost:8081"
		// Extract the custom path by trimming the "/gateway/" prefix
		log.Printf("Extracted path: %s", r.URL.Path)
		customPath := strings.TrimPrefix(r.URL.Path, "/gateway/service_a/")
		log.Printf("Extracted custom path: %s", customPath)

		// Build the full URL for the backend service
		fullURL := fmt.Sprintf("%s/%s", serviceAURL, customPath)
		switch r.Method {
		case http.MethodGet:
			service.ServiceRequest(w, r, fullURL)
		}
	})

	port := ":8080"
	log.Printf("Start Gateway on Port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to Start Gateway: %v", err)
	}
}
