package services

import (
	"fmt"

	"patient-profile-webcomponent/models"
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

// GetPatientByID retrieves a patient by their ID
func GetPatientByID(id int) (*models.Patient, error) {
	patient, exists := patients[id]
	if !exists {
		return nil, fmt.Errorf("patient with ID %d not found", id)
	}
	return &patient, nil
}

// GetAllPatients retrieves all patients
func GetAllPatients() map[int]models.Patient {
	return patients
}

// CreatePatient creates a new patient (for future use)
func CreatePatient(patient models.Patient) error {
	patients[patient.ID] = patient
	return nil
}

// UpdatePatient updates an existing patient (for future use)
func UpdatePatient(id int, patient models.Patient) error {
	if _, exists := patients[id]; !exists {
		return fmt.Errorf("patient with ID %d not found", id)
	}
	patients[id] = patient
	return nil
}

// DeletePatient deletes a patient (for future use)
func DeletePatient(id int) error {
	if _, exists := patients[id]; !exists {
		return fmt.Errorf("patient with ID %d not found", id)
	}
	delete(patients, id)
	return nil
}
