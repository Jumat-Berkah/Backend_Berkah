package main

import (
	"Backend_berkah/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("JUMAT_BERKAH")
	if port == "" {
		port = "8080"
	}

	router := routes.URL()

	log.Printf("Starting server on port %s...", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}