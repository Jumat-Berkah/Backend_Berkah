package routes

import (
	"Backend_berkah/controller"
	"log"
	"net/http"
)

func URL(w http.ResponseWriter, r *http.Request) {
    log.Printf("Received %s request for: %s", r.Method, r.URL.Path)

    // Set CORS headers
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

    // Handle preflight OPTIONS requests
    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    // Route handling based on HTTP method and path
    switch r.Method {
    case http.MethodGet:
        switch r.URL.Path {
        case "/":
            controller.GetHome(w, r)
        default:
            log.Printf("Not Found: %s %s", r.Method, r.URL.Path)
            controller.NotFound(w, r)
        }
    case http.MethodPost:
        switch r.URL.Path {
        case "/register":
            log.Println("Handling /register route")
            controller.Register(w, r)
        case "/login":
            log.Println("Handling /login route")
            controller.Login(w, r)
        default:
            log.Printf("Not Found: %s %s", r.Method, r.URL.Path)
            controller.NotFound(w, r)
        }
    case http.MethodPut:
        // if r.URL.Path == "/data" {
        //     controller.UpdateRoute(w, r)
        // } else {
        //     log.Printf("Not Found: %s %s", r.Method, r.URL.Path)
        //     controller.NotFound(w, r)
        // }
    case http.MethodDelete:
        // if r.URL.Path == "/data" {
        //     controller.DeleteRoute(w, r)
        // } else {
        //     log.Printf("Not Found: %s %s", r.Method, r.URL.Path)
        //     controller.NotFound(w, r)
        // }
    default:
        log.Printf("Method Not Allowed: %s %s", r.Method, r.URL.Path)
        controller.NotFound(w, r)
    }
}
