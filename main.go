package main

import (
	"Backend_berkah/routes"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", http.HandlerFunc(routes.URL)); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
