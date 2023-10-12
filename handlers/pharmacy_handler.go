package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SFWE403/UArizonaPharmacy/internal/models"
)

// Handler function for retrieving a list of pharmacies
func GetPharmaciesHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the list of pharmacies from the models package
	pharmacies := models.GetPharmacies()

	// Convert the list of pharmacies to JSON
	pharmaciesJSON, err := json.Marshal(pharmacies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the client
	w.Write(pharmaciesJSON)
}
