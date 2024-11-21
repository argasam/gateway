package main

import (
	"gateway/cashloan/service"
	"log"
	"net/http"
)

func main() {
	port := ":8082"
	http.HandleFunc("/", service.HandleRequest)
	log.Printf("Starting server on %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Failed to start service a: %v", err)
	}
}
