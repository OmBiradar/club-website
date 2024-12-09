package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

// Data structure for the response
type Secretary struct {
	Photo    string `json:"photo"`
	Name     string `json:"name"`
	OneLiner string `json:"oneLiner"`
}

type ResponseData struct {
	WelcomeMessage   string      `json:"welcomeMessage"`
	Secretary        Secretary   `json:"secretary"`
	JointSecretaries []Secretary `json:"jointSecretaries"`
}

// Mock data
var yearData = map[string]ResponseData{
	"2024-25": {
		WelcomeMessage: "Welcome to CSI | IIT BHU for 2024-25!",
		Secretary: Secretary{
			Photo:    "/images/Mayank_Mehta.jpg",
			Name:     "Mayank Mehta",
			OneLiner: "Leading with Excellence",
		},
		JointSecretaries: []Secretary{
			{
				Photo:    "/images/Arnav_Raj.jpg",
				Name:     "Arnav Raj",
				OneLiner: "Striving for Success",
			},
			{
				Photo:    "/images/Gagan.jpg",
				Name:     "Gagan",
				OneLiner: "Innovation at its Best",
			},
		},
	},
}

func handleYearRequest(w http.ResponseWriter, r *http.Request) {
	// Extract the year from the URL path
	year := r.URL.Path[len("/api/year/"):]

	// Check if the year exists in mock data
	data, exists := yearData[year]
	if !exists {
		http.Error(w, "Data not found", http.StatusNotFound)
		return
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	mux := http.NewServeMux()

	// Serve static files (images folder)
	mux.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))

	// Handle API requests
	mux.HandleFunc("/api/year/", handleYearRequest)

	// Enable CORS
	handler := cors.Default().Handler(mux)

	// Start the server
	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
