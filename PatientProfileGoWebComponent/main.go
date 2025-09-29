package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"patient-profile-webcomponent/components"
	"patient-profile-webcomponent/models"

	"github.com/gorilla/mux"
)

// Mock data - in a real app this would come from a database
var patients = map[int]models.Patient{
	1: {
		ID:          1,
		Name:        "John Doe",
		Email:       "john.doe@example.com",
		Phone:       "+1-555-0123",
		DateOfBirth: "1985-03-15",
		Address:     "123 Main St, Anytown, USA",
		MedicalID:   "MED-001",
	},
	2: {
		ID:          2,
		Name:        "Jane Smith",
		Email:       "jane.smith@example.com",
		Phone:       "+1-555-0456",
		DateOfBirth: "1990-07-22",
		Address:     "456 Oak Ave, Somewhere, USA",
		MedicalID:   "MED-002",
	},
}

func getPatientByID(id int) (*models.Patient, error) {
	patient, exists := patients[id]
	if !exists {
		return nil, fmt.Errorf("patient with ID %d not found", id)
	}
	return &patient, nil
}

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

// API endpoint to get patient data as JSON
func getPatientHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	patientID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		return
	}

	patient, err := getPatientByID(patientID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patient)
}

// Serve the patient profile as HTML (for HTMX)
func getPatientProfileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	patientID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		return
	}

	patient, err := getPatientByID(patientID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Render the patient profile component
	component := components.UserProfile(*patient)
	component.Render(r.Context(), w)
}

// Serve the web component definition
func getWebComponentHandler(w http.ResponseWriter, r *http.Request) {
	webComponentJS := `
// Check if the component is already defined
if (typeof PatientProfileComponent === 'undefined') {
class PatientProfileComponent extends HTMLElement {
	constructor() {
		super();
		this.attachShadow({ mode: 'open' });
	}

	connectedCallback() {
		const patientId = this.getAttribute('patient-id');
		if (!patientId) {
			this.shadowRoot.innerHTML = '<p>Error: patient-id attribute is required</p>';
			return;
		}

		this.loadPatientProfile(patientId);
	}

	async loadPatientProfile(patientId) {
		try {
			const response = await fetch('http://localhost:8091/api/patient/' + patientId);
			if (!response.ok) {
				throw new Error('Failed to load patient data');
			}
			const patient = await response.json();
			
			// Create the patient profile HTML
			const html = '<style>' +
				'.patient-profile { font-family: Arial, sans-serif; border: 1px solid #ddd; border-radius: 8px; padding: 20px; background: #f9f9f9; max-width: 400px; margin: 10px; }' +
				'.patient-profile h3 { margin-top: 0; color: #333; border-bottom: 2px solid #007bff; padding-bottom: 10px; }' +
				'.patient-profile .field { margin: 10px 0; }' +
				'.patient-profile .label { font-weight: bold; color: #555; }' +
				'.patient-profile .value { color: #333; }' +
				'</style>' +
				'<div class="patient-profile">' +
				'<h3>Patient Profile</h3>' +
				'<div class="field"><span class="label">Name:</span><span class="value">' + patient.name + '</span></div>' +
				'<div class="field"><span class="label">Email:</span><span class="value">' + patient.email + '</span></div>' +
				'<div class="field"><span class="label">Phone:</span><span class="value">' + patient.phone + '</span></div>' +
				'<div class="field"><span class="label">Date of Birth:</span><span class="value">' + patient.dateOfBirth + '</span></div>' +
				'<div class="field"><span class="label">Address:</span><span class="value">' + patient.address + '</span></div>' +
				'<div class="field"><span class="label">Medical ID:</span><span class="value">' + patient.medicalId + '</span></div>' +
				'</div>';
			this.shadowRoot.innerHTML = html;
		} catch (error) {
			this.shadowRoot.innerHTML = '<p>Error loading patient profile: ' + error.message + '</p>';
		}
	}
}

// Only define the custom element if it hasn't been defined already
if (!customElements.get('patient-profile')) {
	customElements.define('patient-profile', PatientProfileComponent);
}
}`

	w.Header().Set("Content-Type", "application/javascript")
	w.Write([]byte(webComponentJS))
}

// Serve the main page that demonstrates the web component
func getIndexHandler(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Patient Profile Web Component</title>
    <script src="https://unpkg.com/htmx.org@1.9.10"></script>
    <script src="/webcomponent.js"></script>
</head>
<body>
    <h1>Patient Profile Web Component Demo</h1>
    <p>This page demonstrates the patient profile web component:</p>
    
    <h2>Patient 1:</h2>
    <patient-profile patient-id="1"></patient-profile>
    
    <h2>Patient 2:</h2>
    <patient-profile patient-id="2"></patient-profile>
    
    <h2>Non-existent Patient (Error case):</h2>
    <patient-profile patient-id="999"></patient-profile>
</body>
</html>
`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func main() {
	r := mux.NewRouter()

	// API routes
	r.HandleFunc("/api/patient/{id}", getPatientHandler).Methods("GET")
	r.HandleFunc("/patient/{id}", getPatientProfileHandler).Methods("GET")

	// Web component and demo page
	r.HandleFunc("/webcomponent.js", getWebComponentHandler).Methods("GET")
	r.HandleFunc("/", getIndexHandler).Methods("GET")

	// Serve static files (if any)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Apply CORS middleware to all routes
	handler := corsMiddleware(r)

	fmt.Println("Patient Profile Web Component server starting on :8091")
	fmt.Println("Visit http://localhost:8091 to see the demo")
	log.Fatal(http.ListenAndServe(":8091", handler))
}
