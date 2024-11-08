package routes

import (
	"net/http"

	"Backend_berkah/controller"
)

func URL(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Handle preflight OPTIONS requests
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// // Set environment variables if needed
	// config.SetEnv()

	// Route handling
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			controller.GetHome(w, r)
		default:
			controller.NotFound(w, r)
		}
	case http.MethodPost:
		switch r.URL.Path {
		case "/register":
			controller.Register(w, r)
		case "/login":
			controller.Login(w, r)
		default:
			controller.NotFound(w, r)
		}
	case http.MethodPut:
		// if r.URL.Path == "/data" {
		// 	controller.UpdateRoute(w, r)
		// } else {
		// 	controller.NotFound(w, r)
		// }
	case http.MethodDelete:
		// if r.URL.Path == "/data" {
		// 	controller.DeleteRoute(w, r)
		// } else {
		// 	controller.NotFound(w, r)
		// }
	default:
		controller.NotFound(w, r)
	}
}
