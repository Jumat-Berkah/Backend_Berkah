package main

import (
	"Backend_berkah/helper"
	"Backend_berkah/model"
	"Backend_berkah/routes"
	"context"
	"log"
	"net/http"
	"os"
)

func main() {
	// Get MongoDB connection URI from environment variable
	mongoURI := os.Getenv("JUMAT_BERKAH")
	if mongoURI == "" {
		log.Fatal("Environment variable JUMAT_BERKAH is not set")
	}

	// Initialize MongoDB connection
	mongoInfo := model.DBInfo{
		DBString: mongoURI,
		DBName:   "jumat_berkah",
	}
	db, err := helper.MongoConnect(mongoInfo)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer db.Client().Disconnect(context.Background())

	// Get the server port from environment variable, default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize router
	router := routes.URL()

	log.Printf("Starting server on port %s...", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}