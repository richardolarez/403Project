// main.go

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	dbinitializer "github.com/SFWE403/UArizonaPharmacy/init"
	accountmanager "github.com/SFWE403/UArizonaPharmacy/internal/account_manager"
	"github.com/SFWE403/UArizonaPharmacy/internal/models"
)

func main() {
	// Initialize the database
	err := dbinitializer.InitializeDatabase()
	if err != nil {
		fmt.Printf("Error initializing database: %v\n", err)
		return
	}

	// LoginRequest represents a request to authenticate an employee login.
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
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

	// Define an endpoint to authenticate an employee login
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		// Parse the username and password from the request body
		var loginRequest LoginRequest
		err := json.NewDecoder(r.Body).Decode(&loginRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Authenticate the employee login
		employee, err := accountmanager.AuthenticateEmployee(loginRequest.Username, loginRequest.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Convert the employee to JSON
		employeeJSON, err := json.Marshal(employee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response to the client
		w.Write(employeeJSON)
	})

	// Start the server
	server := &http.Server{
		Addr: ":8080",
	}
	server.ListenAndServe()
}
