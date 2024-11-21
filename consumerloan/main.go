package main

import (
	"gateway/consumerloan/service"
	"log"
	"net/http"
)

func main() {
	port := ":8081"
	http.HandleFunc("/", service.HandleRequest)
	http.HandleFunc("/search", service.HandleSearch)
	log.Printf("Starting server on %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Failed to start service a: %v", err)
	}
}
