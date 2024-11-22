package main

import (
	"fmt"
	"gateway/service"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/gateway/consumerloan/", func(w http.ResponseWriter, r *http.Request) {
		const serviceAURL = "http://localhost:8081"
		// Extract the custom path by trimming the "/gateway/" prefix
		log.Printf("Extracted path: %s", r.URL.Path)
		customPath := strings.TrimPrefix(r.URL.Path, "/gateway/consumerloan")
		log.Printf("Extracted custom path: %s", customPath)

		// Build the full URL for the backend service
		fullURL := fmt.Sprintf("%s%s", serviceAURL, customPath)
		switch r.Method {
		case http.MethodGet:
			service.ServiceRequest(w, r, fullURL)
		}
	})

	http.HandleFunc("/gateway/consumerloan", func(w http.ResponseWriter, r *http.Request) {
		// Redirect to ensure trailing slash is handled consistently
		http.Redirect(w, r, "/gateway/consumerloan/", http.StatusMovedPermanently)
	})

	http.HandleFunc("/gateway/cashloan/", func(w http.ResponseWriter, r *http.Request) {
		const serviceAURL = "http://localhost:8082"
		// Extract the custom path by trimming the "/gateway/" prefix
		log.Printf("Extracted path: %s", r.URL.Path)
		customPath := strings.TrimPrefix(r.URL.Path, "/gateway/cashloan")
		log.Printf("Extracted custom path: %s", customPath)

		// Build the full URL for the backend service
		fullURL := fmt.Sprintf("%s%s", serviceAURL, customPath)
		switch r.Method {
		case http.MethodGet:
			service.ServiceRequest(w, r, fullURL)
		}
	})

	http.HandleFunc("/gateway/cashloan", func(w http.ResponseWriter, r *http.Request) {
		// Redirect to ensure trailing slash is handled consistently
		http.Redirect(w, r, "/gateway/cashloan/", http.StatusMovedPermanently)
	})

	http.HandleFunc("/gateway/risk/", func(w http.ResponseWriter, r *http.Request) {
		const serviceAURL = "http://localhost:8083"
		// Extract the custom path by trimming the "/gateway/" prefix
		log.Printf("Extracted path: %s", r.URL.Path)
		customPath := strings.TrimPrefix(r.URL.Path, "/gateway/risk")
		log.Printf("Extracted custom path: %s", customPath)

		// Build the full URL for the backend service
		fullURL := fmt.Sprintf("%s%s", serviceAURL, customPath)
		switch r.Method {
		case http.MethodGet:
			service.ServiceRequest(w, r, fullURL)
		}
	})

	http.HandleFunc("/gateway/risk", func(w http.ResponseWriter, r *http.Request) {
		// Redirect to ensure trailing slash is handled consistently
		http.Redirect(w, r, "/gateway/risk/", http.StatusMovedPermanently)
	})

	port := ":8080"
	log.Printf("Start Gateway on Port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to Start Gateway: %v", err)
	}
}
