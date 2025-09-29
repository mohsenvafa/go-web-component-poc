import React, { useEffect } from 'react';
import './App.css';

function App() {
  useEffect(() => {
    // Check if the script is already loaded
    const existingScript = document.querySelector('script[src="http://localhost:8091/webcomponent.js"]');
    if (existingScript) {
      return; // Script already loaded
    }

    // Load the web component script from the Go service
    const script = document.createElement('script');
    script.src = 'http://localhost:8091/webcomponent.js';
    script.crossOrigin = 'anonymous';
    document.head.appendChild(script);

    return () => {
      // Cleanup: remove the script when component unmounts
      const scriptToRemove = document.querySelector('script[src="http://localhost:8091/webcomponent.js"]');
      if (scriptToRemove) {
        document.head.removeChild(scriptToRemove);
      }
    };
  }, []);

  return (
    <div className="App">
      <div className="container">
        <div className="header">
          <h1>POC Portal - React App</h1>
          <p>This React application demonstrates consuming the Patient Profile Web Component</p>
        </div>
        
        <div className="info">
          <h3>About this POC</h3>
          <p>This React application consumes the Patient Profile Web Component from the PatientProfileGoWebComponent service. 
          The web component is loaded dynamically and displays patient information.</p>
          <p><strong>Note:</strong> Make sure the PatientProfileGoWebComponent service is running on port 8080.</p>
        </div>
        
        <div className="patient-grid">
          <div className="patient-section">
            <h3>Patient 1 - John Doe</h3>
            <patient-profile patient-id="1"></patient-profile>
          </div>
          
          <div className="patient-section">
            <h3>Patient 2 - Jane Smith</h3>
            <patient-profile patient-id="2"></patient-profile>
          </div>
        </div>
        
        <div style={{ marginTop: '30px', textAlign: 'center', color: '#666' }}>
          <p>This demonstrates how a React application can consume web components from a Go service.</p>
        </div>
      </div>
    </div>
  );
}

export default App;
