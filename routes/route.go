package route

import (
	"net/http"

	"Backend_berkah/config"
	"Backend_berkah/controller"
)

func URL(w http.ResponseWriter, r *http.Request) {
	// Set Access Control Headers
	if config.SetAccessControlHeaders(w, r) {
		return
	}

	// Load environment variables (if necessary)
	// config.SetEnv()

	// Route handling based on HTTP method and URL path
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			controller.GetHome(w, r)  // Home handler
		default:
			controller.NotFound(w, r)
		}

	case http.MethodPost:
		switch r.URL.Path {
		case "/register":
			controller.Register(w, r)  // Register a new user
		case "/login":
			controller.Login(w, r)  // Login for an existing user
		default:
			controller.NotFound(w, r)
		}

	case http.MethodPut:
		// if r.URL.Path == "/data" {
		// 	controller.UpdateRoute(w, r)  // Update existing route data
		// } else {
		// 	controller.NotFound(w, r)
		// }

	case http.MethodDelete:
		// if r.URL.Path == "/data" {
		// 	controller.DeleteRoute(w, r)  // Delete route data
		// } else {
		// 	controller.NotFound(w, r)
		// }

	default:
		controller.NotFound(w, r)
	}
}
