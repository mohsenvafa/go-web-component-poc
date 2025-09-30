// Shared stylesheets cache (one per base URL)
// Use window to avoid redeclaration errors in React dev mode
if (typeof window.patientProfileStyleSheets === "undefined") {
  window.patientProfileStyleSheets = {};
}

// Load CSS once per base URL and share across instances
if (typeof window.loadPatientProfileStyleSheet === "undefined") {
  window.loadPatientProfileStyleSheet = async function (baseUrl) {
    if (!window.patientProfileStyleSheets[baseUrl]) {
      try {
        const response = await fetch(
          `${baseUrl}/static/css/patient-profile.css`
        );
        const css = await response.text();
        const styleSheet = new CSSStyleSheet();
        await styleSheet.replace(css);
        window.patientProfileStyleSheets[baseUrl] = styleSheet;
      } catch (error) {
        console.error("Failed to load stylesheet:", error);
      }
    }
    return window.patientProfileStyleSheets[baseUrl];
  };
}

// Check if the component is already defined
if (typeof PatientProfileComponent === "undefined") {
  class PatientProfileComponent extends HTMLElement {
    constructor() {
      super();
      this.attachShadow({ mode: "open" });
    }

    async connectedCallback() {
      const patientId = this.getAttribute("patient-id");
      const baseUrl = this.getAttribute("base-url") || "http://localhost:8091";

      if (!patientId) {
        this.shadowRoot.innerHTML =
          "<p>Error: patient-id attribute is required</p>";
        return;
      }

      // Load shared stylesheet once and adopt it
      const styleSheet = await window.loadPatientProfileStyleSheet(baseUrl);
      if (styleSheet) {
        this.shadowRoot.adoptedStyleSheets = [styleSheet];
      }

      this.loadPatientProfile(patientId, baseUrl);
    }

    async loadPatientProfile(patientId, baseUrl) {
      try {
        const response = await fetch(`${baseUrl}/patient/${patientId}`);
        if (!response.ok) {
          throw new Error("Failed to load patient profile");
        }
        const html = await response.text();
        this.shadowRoot.innerHTML = html;
      } catch (error) {
        this.shadowRoot.innerHTML =
          "<p>Error loading patient profile: " + error.message + "</p>";
      }
    }
  }

  // Only define the custom element if it hasn't been defined already
  if (!customElements.get("patient-profile")) {
    customElements.define("patient-profile", PatientProfileComponent);
  }
}
