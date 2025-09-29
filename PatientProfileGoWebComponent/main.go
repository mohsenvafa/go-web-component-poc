package main

import (
	"fmt"
	"log"
	"net/http"

	"patient-profile-webcomponent/api"
	"patient-profile-webcomponent/components/user_profile"
	"patient-profile-webcomponent/web_components"

	"github.com/gorilla/mux"
)

// CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	// API routes
	r.HandleFunc("/api/patient/{id}", api.GetPatientHandler).Methods("GET")
	r.HandleFunc("/patient/{id}", user_profile.GetPatientProfileHandler).Methods("GET")

	// Web component and demo page
	r.HandleFunc("/webcomponent.js", web_components.GetWebComponentHandler).Methods("GET")
	r.HandleFunc("/", web_components.GetIndexHandler).Methods("GET")

	// Serve static files (if any)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Apply CORS middleware to all routes
	handler := corsMiddleware(r)

	fmt.Println("Patient Profile Web Component server starting on :8091")
	fmt.Println("Visit http://localhost:8091 to see the demo")
	log.Fatal(http.ListenAndServe(":8091", handler))
}
