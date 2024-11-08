package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"Backend_berkah/config"
	"Backend_berkah/model"
)

// Register user
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("Register endpoint: Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("Register endpoint: Invalid input - %v\n", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Register endpoint: Failed to hash password - %v\n", err)
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)
	user.ID = primitive.NewObjectID()

	// Insert user into the "users" collection
	collection := config.DB.Collection("users")
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Printf("Register endpoint: Failed to register user - %v\n", err)
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	log.Println("Register endpoint: User registered successfully")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// Login user
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("Login endpoint: Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var loginRequest model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		log.Printf("Login endpoint: Invalid input - %v\n", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Retrieve user from the "users" collection
	collection := config.DB.Collection("users")
	var user model.User
	err := collection.FindOne(context.Background(), bson.M{"email": loginRequest.Email}).Decode(&user)
	if err != nil {
		log.Printf("Login endpoint: Invalid email or password - %v\n", err)
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		log.Printf("Login endpoint: Invalid email or password - %v\n", err)
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	log.Println("Login endpoint: User logged in successfully")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User logged in successfully"})
}
