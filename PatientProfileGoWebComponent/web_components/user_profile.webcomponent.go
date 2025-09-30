package web_components

import (
	"net/http"
)

// Serve the web component definition
func GetWebComponentHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the static JavaScript file
	http.ServeFile(w, r, "./static/js/patient-profile-webcomponent.js")
}

// Serve the main page that demonstrates the web component
func GetIndexHandler(w http.ResponseWriter, r *http.Request) {
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
    <patient-profile patient-id="1" base-url="http://localhost:8091"></patient-profile>
    
    <h2>Patient 2:</h2>
    <patient-profile patient-id="2" base-url="http://localhost:8091"></patient-profile>
    
    <h2>Non-existent Patient (Error case):</h2>
    <patient-profile patient-id="999" base-url="http://localhost:8091"></patient-profile>
</body>
</html>
`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}
