package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"patient-profile-webcomponent/services"

	"github.com/gorilla/mux"
)

// GetPatientHandler - API endpoint to get patient data as JSON
func GetPatientHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	patientID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		return
	}

	patient, err := services.GetPatientByID(patientID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patient)
}
