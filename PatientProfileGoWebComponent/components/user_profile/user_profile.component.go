package user_profile

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

func GetPatientByID(id int) (*models.Patient, error) {
	patient, exists := patients[id]
	if !exists {
		return nil, fmt.Errorf("patient with ID %d not found", id)
	}
	return &patient, nil
}

// API endpoint to get patient data as JSON
func GetPatientHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	patientID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		return
	}

	patient, err := GetPatientByID(patientID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patient)
}

// Serve the patient profile as HTML (for HTMX)
func GetPatientProfileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	patientID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		return
	}

	patient, err := GetPatientByID(patientID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Render the patient profile component
	component := UserProfile(*patient)
	component.Render(r.Context(), w)
}
