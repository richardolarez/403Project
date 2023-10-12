// main.go

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	dbinitializer "github.com/SFWE403/UArizonaPharmacy/init"
	"github.com/SFWE403/UArizonaPharmacy/internal/models"
)

func main() {
	// Initialize the database
	err := dbinitializer.InitializeDatabase()
	if err != nil {
		fmt.Printf("Error initializing database: %v\n", err)
		return
	}

	// Define an endpoint to retrieve all inventory items
	http.HandleFunc("/inventory", func(w http.ResponseWriter, r *http.Request) {
		// Retrieve all inventory items from the database
		inventory, err := models.GetInventory()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert the list of inventory items to JSON
		inventoryJSON, err := json.Marshal(inventory)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response to the client
		w.Write(inventoryJSON)
	})

	// Start the server
	server := &http.Server{
		Addr: ":8080",
	}
	server.ListenAndServe()
}
