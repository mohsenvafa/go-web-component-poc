package models

// Patient represents a patient profile
type Patient struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	DateOfBirth string `json:"dateOfBirth"`
	Address     string `json:"address"`
	MedicalID   string `json:"medicalId"`
}
