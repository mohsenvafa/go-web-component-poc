// Check if the component is already defined
if (typeof PatientProfileComponent === "undefined") {
  class PatientProfileComponent extends HTMLElement {
    constructor() {
      super();
      this.attachShadow({ mode: "open" });
    }

    connectedCallback() {
      const patientId = this.getAttribute("patient-id");
      if (!patientId) {
        this.shadowRoot.innerHTML =
          "<p>Error: patient-id attribute is required</p>";
        return;
      }

      this.loadPatientProfile(patientId);
    }

    async loadPatientProfile(patientId) {
      try {
        const response = await fetch(
          "http://localhost:8091/patient/" + patientId
        );
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
