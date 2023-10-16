// main.go

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// DeleteRequest represents a request to delete an employee
	type DeleteRequest struct {
		ID        int    `json:"id"`
		FirstName string `json:"firstName"`
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

	// Define an endpoint to delete an employee by ID and first name
	http.HandleFunc("/deleteEmployee", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Parse the request parameters
		var deleteRequest DeleteRequest
		err := json.NewDecoder(r.Body).Decode(&deleteRequest)
		if err != nil {
			http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
			return
		}

		id := deleteRequest.ID
		firstName := deleteRequest.FirstName

		// Call the DeleteEmployee function to delete the employee
		err = models.DeleteEmployee(id, firstName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond with a success message
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Employee deleted successfully"))
	})

	// Start the server
	server := &http.Server{
		Addr: ":8080",
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for an interrupt signal (SIGINT) or a termination signal (SIGTERM)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	} else {
		log.Println("Server shutdown completed")
	}
}
