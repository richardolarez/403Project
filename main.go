// main.go

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	dbinitializer "github.com/richardolarez/403Project/init"
	accountmanager "github.com/richardolarez/403Project/internal/account_manager"
	"github.com/richardolarez/403Project/internal/models"
	"github.com/richardolarez/403Project/internal/service"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func main() {
	// Check if the database file exists
	if _, err := os.Stat("./db/database.json"); os.IsNotExist(err) {
		// Initialize the database
		err := dbinitializer.InitializeDatabase()
		if err != nil {
			fmt.Printf("Error initializing database: %v\n", err)
			return
		}
	} else {
		fmt.Println("Database file already exists.")
	}

	// LoginRequest represents a request to authenticate an employee login.
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// CheckoutRequest represents a request to process a sales transaction.
	type CheckoutRequest struct {
		CustomerID    int                     `json:"customer_id"`
		Items         []*models.InventoryItem `json:"items"`
		PaymentMethod string                  `json:"payment_method"`
	}

	// Define an endpoint to retrieve all inventory items
	http.HandleFunc("/inventory", func(w http.ResponseWriter, r *http.Request) {
		// Retrieve all inventory items from the database
		enableCors(&w)
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

	// Define an endpoint to retrieve all employees
	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		// Retrieve all employees from the database
		enableCors(&w)
		employees, err := models.GetAllEmployees()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert the list of employees to JSON
		employeesJSON, err := json.Marshal(employees)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response to the client
		w.Write(employeesJSON)
	})

	// Define an endpoint to authenticate an employee login
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
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

	http.HandleFunc("/checkout", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Parse the checkout request from the request body
		var checkoutRequest CheckoutRequest
		err := json.NewDecoder(r.Body).Decode(&checkoutRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var receipt *string

		// Call the Checkout function to process the order and get a sales transaction
		receipt, transaction, err := service.Checkout(checkoutRequest.CustomerID, checkoutRequest.Items, checkoutRequest.PaymentMethod)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert the transaction to JSON
		transactionJSON, err := json.Marshal(transaction)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert the receipt to JSON
		receiptJSON, err := json.Marshal(receipt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response to the client
		w.Write(transactionJSON)

		// Write the JSON response to the client
		w.Write(receiptJSON)

	})

	// Start the server

	server := &http.Server{
		Addr: ":8080",
	}
	server.ListenAndServe()
}
