# Go Web Component POC

This repository contains a Proof of Concept demonstrating how to create and consume web components using Go, HTMX, and Templ.

## Projects

### 1. PatientProfileGoWebComponent
- **Location**: `./PatientProfileGoWebComponent/`
- **Purpose**: Go service that serves patient profile data as a web component
- **Port**: 8091
- **Features**:
  - Go backend with HTMX and Templ
  - Patient profile API endpoints
  - Web component definition
  - Demo page showing the web component

### 2. PocPortalGoSurfaceApp
- **Location**: `./PocPortalGoSurfaceApp/`
- **Purpose**: Go application that consumes the patient profile web component
- **Port**: 8092
- **Features**:
  - Go frontend with HTMX and Templ
  - Consumes web component from PatientProfileGoWebComponent
  - Demonstrates Go-to-Go web component consumption

### 3. PocPortalReactApp
- **Location**: `./PocPortalReactApp/`
- **Purpose**: React application that consumes the patient profile web component
- **Port**: 8093
- **Features**:
  - React frontend
  - Consumes web component from PatientProfileGoWebComponent
  - Demonstrates React-to-Go web component consumption

## Quick Start

### Prerequisites
- Go 1.21+
- Node.js 16+
- npm

### Running the POC

1. **Start the Patient Profile Web Component Service**:
   ```bash
   cd PatientProfileGoWebComponent
   make deps
   make run
   ```
   This starts the service on http://localhost:8091

2. **Start the Go Portal App** (in a new terminal):
   ```bash
   cd PocPortalGoSurfaceApp
   make deps
   make run
   ```
   This starts the Go portal on http://localhost:8092

3. **Start the React Portal App** (in a new terminal):
   ```bash
   cd PocPortalReactApp
   make install
   make start
   ```
   This starts the React app on http://localhost:8093

### Testing the Web Component

1. Visit http://localhost:8091 to see the web component demo
2. Visit http://localhost:8092 to see the Go app consuming the web component
3. Visit http://localhost:8093 to see the React app consuming the web component

## Architecture

```
┌─────────────────────────────────────┐
│  PatientProfileGoWebComponent       │
│  (Port 8091)                        │
│  - Go + HTMX + Templ                │
│  - Patient API                      │
│  - Web Component Definition         │
└─────────────────┬───────────────────┘
                  │
                  │ Web Component
                  │ (patient-profile)
                  │
        ┌─────────┴─────────┐
        │                   │
┌───────▼────────┐ ┌────────▼────────┐
│ PocPortalGo    │ │ PocPortalReact  │
│ SurfaceApp     │ │ App             │
│ (Port 8092)    │ │ (Port 8093)     │
│ - Go + HTMX    │ │ - React         │
│ - Templ        │ │ - Web Component │
└────────────────┘ └─────────────────┘
```

## Web Component Usage

The web component can be used in any HTML page:

```html
<script src="http://localhost:8091/webcomponent.js"></script>
<patient-profile patient-id="1"></patient-profile>
```

## Key Features Demonstrated

1. **Web Component Creation**: How to create a web component in Go
2. **Cross-Framework Consumption**: How to consume the same web component in both Go and React
3. **API Integration**: How the web component fetches data from Go APIs
4. **Shadow DOM**: How the web component uses Shadow DOM for encapsulation
5. **HTMX Integration**: How HTMX works with Go and Templ

## Development

Each project has its own Makefile with common targets:
- `make deps` - Install dependencies
- `make build` - Build the application
- `make run` - Run the application
- `make clean` - Clean build artifacts
- `make help` - Show available targets

## Notes

- This is a POC and uses mock data
- In production, you'd want proper error handling, authentication, and database integration
- The web component uses Shadow DOM for style encapsulation
- CORS is not configured - in production, you'd need to handle cross-origin requests properly
