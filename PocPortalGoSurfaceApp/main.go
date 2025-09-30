package main

import (
	"fmt"
	"log"
	"net/http"

	"poc-portal-go-surface/components"

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

	// Serve the main portal page
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		component := components.PortalPage()
		component.Render(r.Context(), w)
	}).Methods("GET")

	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Serve the patient profile web component from the other service
	r.HandleFunc("/webcomponent.js", func(w http.ResponseWriter, r *http.Request) {
		// In a real application, this would proxy to the PatientProfileGoWebComponent service
		// For this POC, we'll serve it directly from the other service
		resp, err := http.Get("http://localhost:8091/webcomponent.js")
		if err != nil {
			http.Error(w, "Failed to load web component", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		w.Header().Set("Content-Type", "application/javascript")
		w.WriteHeader(resp.StatusCode)
		resp.Write(w)
	}).Methods("GET")

	// Apply CORS middleware to all routes
	handler := corsMiddleware(r)

	fmt.Println("POC Portal Go Surface App starting on :8092")
	fmt.Println("Visit http://localhost:8092 to see the portal")
	log.Fatal(http.ListenAndServe(":8092", handler))
}
