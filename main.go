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
	"strconv"
	"syscall"
	"time"

	dbinitializer "github.com/richardolarez/403Project/init"
	accountmanager "github.com/richardolarez/403Project/internal/account_manager"
	"github.com/richardolarez/403Project/internal/logger"
	"github.com/richardolarez/403Project/internal/models"
	"github.com/richardolarez/403Project/internal/service"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func main() {

	//

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
		CartItems     []*service.Cart `json:"cartItems"`
		CustomerID    string          `json:"customer_id"`
		ItemID        string          `json:"item_id"`
		PaymentMethod string          `json:"payment_method"`
	}

	// DeleteRequest represents a request to delete an employee/customer by ID and first name.
	type DeleteRequest struct {
		ID        int    `json:"id"`
		FirstName string `json:"firstName"`
	}

	type AddRequest struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Role      string `json:"role"`
	}

	type AddCustomerRequest struct {
		FirstName   string `json:"firstname"`
		LastName    string `json:"lastname"`
		DOB         string `json:"dob"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phonenumber"`
		Address     string `json:"address"`
		Insurance   string `json:"insurance"`
	}

	// Create a logger instance
	logDir := "/db"
	loggerInst := logger.NewLogger(logDir)

	// Define an endpoint to retrieve all inventory items
	http.HandleFunc("/inventory", func(w http.ResponseWriter, r *http.Request) {
		// Retrieve all inventory items from the database
		enableCors(&w)
		inventory, err := models.GetInventory()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error retrieving inventory to JSON", map[string]interface{}{"error": err.Error(), "inventory": inventory})
			return
		}

		loggerInst.Log(logger.Info, "Converting inventory to JSON", map[string]interface{}{"inventory": inventory})
		// Convert the list of inventory items to JSON
		inventoryJSON, err := json.Marshal(inventory)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error converting inventory to JSON", map[string]interface{}{"error": err.Error()})
			return
		}

		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response to the client
		w.Write(inventoryJSON)

		loggerInst.Log(logger.Info, "Inventory request completed", map[string]interface{}{"response_code": http.StatusOK})
	})

	// Define an endpoint to retrieve all prescriptions
	http.HandleFunc("/prescriptions", func(w http.ResponseWriter, r *http.Request) {
		// Retrieve all prescriptions from the database
		enableCors(&w)
		prescriptions, err := models.GetPrescriptions()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert the list of prescriptions to JSON
		prescriptionsJSON, err := json.Marshal(prescriptions)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response to the client
		w.Write(prescriptionsJSON)

		loggerInst.Log(logger.Info, "Prescriptions request completed", map[string]interface{}{"response_code": http.StatusOK})
	})

	// Define an endpoint to retrieve all employees
	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		// Retrieve all employees from the database
		enableCors(&w)
		employees, err := models.GetAllEmployees()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error retrieving employees to JSON", map[string]interface{}{"error": err.Error()})
			return
		}

		loggerInst.Log(logger.Info, "Converting employees to JSON", map[string]interface{}{"employees": employees})
		// Convert the list of employees to JSON
		employeesJSON, err := json.Marshal(employees)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error converting employees to JSON", map[string]interface{}{"error": err.Error()})
			return
		}

		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response to the client
		w.Write(employeesJSON)

		loggerInst.Log(logger.Info, "Employees request completed", map[string]interface{}{"response_code": http.StatusOK})
	})

	//Define an endpoint to retrieve all customers
	http.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		// Retrieve all customers from the database
		enableCors(&w)
		customers, err := models.GetAllCustomers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error retrieving customers to JSON", map[string]interface{}{"error": err.Error()})
			return
		}

		loggerInst.Log(logger.Info, "Converting customers to JSON", map[string]interface{}{"customers": customers})
		// Convert the list of customers to JSON
		customersJSON, err := json.Marshal(customers)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error converting customers to JSON", map[string]interface{}{"error": err.Error()})
			return
		}

		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response to the client
		w.Write(customersJSON)

		loggerInst.Log(logger.Info, "Customers request completed", map[string]interface{}{"response_code": http.StatusOK})
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
			loggerInst.Log(logger.Error, "Error parsing login request", map[string]interface{}{"error": err.Error(), "username": loginRequest.Username, "password": loginRequest.Password})
			return
		}

		// Authenticate the employee login
		loggerInst.Log(logger.Info, "Received login request", map[string]interface{}{"request_method": r.Method, "request_path": r.URL.Path, "username": loginRequest.Username, "password": loginRequest.Password})
		employee, err := accountmanager.AuthenticateEmployee(loginRequest.Username, loginRequest.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			loggerInst.Log(logger.Error, "Error authenticating employee login", map[string]interface{}{"error": err.Error()})
			return
		}

		// Convert the employee to JSON
		employeeJSON, err := json.Marshal(employee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error converting employee to JSON", map[string]interface{}{"error": err.Error(), "username": loginRequest.Username, "password": loginRequest.Password})
			return
		}

		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response to the client
		w.Write(employeeJSON)
		loggerInst.Log(logger.Info, "Login request completed", map[string]interface{}{"response_code": http.StatusOK, "username": loginRequest.Username, "password": loginRequest.Password})
	})

	http.HandleFunc("/checkout", func(w http.ResponseWriter, r *http.Request) {
		loggerInst.Log(logger.Info, "Received checkout request", map[string]interface{}{"request_method": r.Method, "request_path": r.URL.Path})
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
			loggerInst.Log(logger.Error, "Error parsing checkout request", map[string]interface{}{"error": err.Error()})
			return
		}

		// Call the Checkout function to process the order and get a sales transaction
		customerID, err := strconv.Atoi(checkoutRequest.CustomerID)
		receipt, err := service.Checkout(customerID, checkoutRequest.PaymentMethod, checkoutRequest.CartItems)
		print(receipt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error processing checkout request", map[string]interface{}{"error": err.Error()})
			return
		}

		// // Convert the transaction to JSON
		// transactionJSON, err := json.Marshal(transaction)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	loggerInst.Log(logger.Error, "Error converting transaction to JSON", map[string]interface{}{"error": err.Error()})
		// 	return
		// }

		// Convert the receipt to JSON
		receiptJSON, err := json.Marshal(&receipt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error converting receipt to JSON", map[string]interface{}{"error": err.Error()})
			return
		}

		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")

		// // Write the JSON response to the client
		// w.Write(transactionJSON)
		// loggerInst.Log(logger.Info, "Checkout request completed", map[string]interface{}{"response_code": http.StatusOK, "transaction": transaction})

		// Write the JSON response to the client
		w.Write(receiptJSON)
		loggerInst.Log(logger.Info, "Checkout request completed", map[string]interface{}{"response_code": http.StatusOK, "receipt": receipt})

	})

	// Define an endpoint to delete an employee by ID and first name
	http.HandleFunc("/deleteEmployee", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			loggerInst.Log(logger.Info, "Delete employee request completed", map[string]interface{}{"response_code": http.StatusOK})
			return
		}

		// Parse the request parameters
		var deleteRequest DeleteRequest
		err := json.NewDecoder(r.Body).Decode(&deleteRequest)
		if err != nil {
			http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
			loggerInst.Log(logger.Error, "Error parsing delete employee request", map[string]interface{}{"error": err.Error()})
			return
		}

		id := deleteRequest.ID
		firstName := deleteRequest.FirstName

		// Call the DeleteEmployee function to delete the employee
		err = models.DeleteEmployee(id, firstName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error deleting employee", map[string]interface{}{"error": err.Error()})
			return
		}

		// Respond with a success message
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Employee deleted successfully"))
		loggerInst.Log(logger.Info, "Delete employee request completed", map[string]interface{}{"response_code": http.StatusOK, "id": id, "firstName": firstName})
	})

	// Define an endpoint to add an employee by username, password, first name, last name,and role
	http.HandleFunc("/addEmployee", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			loggerInst.Log(logger.Info, "Add employee request completed", map[string]interface{}{"response_code": http.StatusOK})
			return
		}

		// Parse the request parameters
		var addRequest AddRequest
		err := json.NewDecoder(r.Body).Decode(&addRequest)
		if err != nil {
			http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
			loggerInst.Log(logger.Error, "Error parsing add employee request", map[string]interface{}{"error": err.Error()})
			return
		}

		/*username := addRequest.Username
		password := addRequest.Password
		firstName := addRequest.FirstName
		lastName := addRequest.LastName
		role := addRequest.Role*/

		// Call the AddEmployee function to delete the employee
		employee, err := models.AddEmployee(addRequest.Username, addRequest.Password, addRequest.FirstName, addRequest.LastName, addRequest.Role)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error adding employee", map[string]interface{}{"error": err.Error()})
			return
		}

		// Convert the employee to JSON
		employeeJSON, err := json.Marshal(employee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error converting employee to JSON", map[string]interface{}{"error": err.Error()})
			return
		}

		// Respond with a success message
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Employee added successfully"))
		w.Write(employeeJSON)
		loggerInst.Log(logger.Info, "Add employee request completed", map[string]interface{}{"response_code": http.StatusOK, "employee": employee})
	})

	// Define an endpoint to add an customer by first name, last name, email, phone number, and address
	http.HandleFunc("/addCustomer", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			loggerInst.Log(logger.Info, "Add customer request completed", map[string]interface{}{"response_code": http.StatusOK})
			return
		}

		// Parse the request parameters
		var addCustomerRequest AddCustomerRequest
		err := json.NewDecoder(r.Body).Decode(&addCustomerRequest)
		if err != nil {
			http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
			loggerInst.Log(logger.Error, "Error parsing add customer request", map[string]interface{}{"error": err.Error()})
			return
		}

		// Call the AddCustomer function to add the customer
		customer, err := models.AddCustomer(addCustomerRequest.FirstName, addCustomerRequest.LastName, addCustomerRequest.DOB, addCustomerRequest.Email, addCustomerRequest.PhoneNumber, addCustomerRequest.Address, addCustomerRequest.Insurance)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error adding customer", map[string]interface{}{"error": err.Error()})
			return
		}

		// Convert the customer to JSON
		customerJSON, err := json.Marshal(customer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error converting customer to JSON", map[string]interface{}{"error": err.Error()})
			return
		}

		// Respond with a success message
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Customer added successfully"))
		w.Write(customerJSON)
		loggerInst.Log(logger.Info, "Add customer request completed", map[string]interface{}{"response_code": http.StatusOK, "customer": customer})
	})

	// Define an endpoint to delete a customer by ID and first name
	http.HandleFunc("/deleteCustomer", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			loggerInst.Log(logger.Info, "Delete customer request completed", map[string]interface{}{"response_code": http.StatusOK})
			return
		}

		// Parse the request parameters
		var deleteRequest DeleteRequest
		err := json.NewDecoder(r.Body).Decode(&deleteRequest)
		if err != nil {
			http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
			loggerInst.Log(logger.Error, "Error parsing delete customer request", map[string]interface{}{"error": err.Error()})
			return
		}

		id := deleteRequest.ID
		firstName := deleteRequest.FirstName

		// Call the DeleteEmployee function to delete the employee
		err = models.DeleteCustomer(id, firstName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error deleting customer", map[string]interface{}{"error": err.Error()})
			return
		}

		// Respond with a success message
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Customer deleted successfully"))
		loggerInst.Log(logger.Info, "Delete customer request completed", map[string]interface{}{"response_code": http.StatusOK, "id": id, "firstName": firstName})
	})

	// Define an endpoint to add new inventory items
	http.HandleFunc("/addNewInventoryItem", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			loggerInst.Log(logger.Info, "Add inventory item request completed", map[string]interface{}{"response_code": http.StatusOK})
			return
		}

		// Parse the request parameters
		type AddInventoryRequest struct {
			Name           string  `json:"name"`
			Description    string  `json:"description"`
			Price          float64 `json:"price"`
			Quantity       int     `json:"quantity"`
			IsPrescription bool    `json:"isPrescription"`
		}

		var addInventoryRequest AddInventoryRequest

		err := json.NewDecoder(r.Body).Decode(&addInventoryRequest)
		if err != nil {
			http.Error(w, "Invalid request parameters", http.StatusBadRequest)
			loggerInst.Log(logger.Error, "Error parsing add inventory item request", map[string]interface{}{"error": err.Error()})
			return
		}

		// Call the NewInventoryItem function to add the inventory item
		err = models.NewInventoryItem(addInventoryRequest.Name, addInventoryRequest.Description, addInventoryRequest.Price, addInventoryRequest.Quantity, addInventoryRequest.IsPrescription)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error adding inventory item", map[string]interface{}{"error": err.Error()})
			return
		}

		// Respond with a success message
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Inventory item added successfully"))
		loggerInst.Log(logger.Info, "Add inventory item request completed", map[string]interface{}{"response_code": http.StatusOK, "name": addInventoryRequest.Name, "description": addInventoryRequest.Description, "price": addInventoryRequest.Price, "quantity": addInventoryRequest.Quantity, "isPrescription": addInventoryRequest.IsPrescription})
	})

	// Define an endpoint to update inventory items
	http.HandleFunc("/updateInventoryItem", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			loggerInst.Log(logger.Info, "Update inventory item request completed", map[string]interface{}{"response_code": http.StatusOK})
			return
		}

		// Parse the request parameters
		type UpdateInventoryRequest struct {
			ID             int     `json:"id"`
			Name           string  `json:"name"`
			Description    string  `json:"description"`
			Price          float64 `json:"price"`
			Quantity       int     `json:"quantity"`
			IsPrescription bool    `json:"isPrescription"`
		}

		var updateInventoryRequest UpdateInventoryRequest

		err := json.NewDecoder(r.Body).Decode(&updateInventoryRequest)
		if err != nil {
			http.Error(w, "Invalid request parameters", http.StatusBadRequest)
			loggerInst.Log(logger.Error, "Error parsing update inventory item request", map[string]interface{}{"error": err.Error()})
			return
		}

		// Call the GetInventoryItem function to get the inventory item
		item, err := models.GetInventoryItem(updateInventoryRequest.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error retrieving inventory item", map[string]interface{}{"error": err.Error()})
			return
		}

		// Update the inventory item
		item.Name = updateInventoryRequest.Name
		item.Description = updateInventoryRequest.Description
		item.Price = updateInventoryRequest.Price
		item.Quantity = updateInventoryRequest.Quantity

		item.Update(updateInventoryRequest.Name, updateInventoryRequest.Description, updateInventoryRequest.Price, updateInventoryRequest.Quantity, updateInventoryRequest.IsPrescription)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error updating inventory item", map[string]interface{}{"error": err.Error()})
			return
		}

		// Respond with a success message
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Inventory item updated successfully"))
		loggerInst.Log(logger.Info, "Update inventory item request completed", map[string]interface{}{"response_code": http.StatusOK, "id": updateInventoryRequest.ID, "name": updateInventoryRequest.Name, "description": updateInventoryRequest.Description, "price": updateInventoryRequest.Price, "quantity": updateInventoryRequest.Quantity, "isPrescription": updateInventoryRequest.IsPrescription})
	})

	// Define an endpoint to update an employee's password
	http.HandleFunc("/updatePassword", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			loggerInst.Log(logger.Info, "Update password request completed", map[string]interface{}{"response_code": http.StatusOK})
			return
		}

		// Parse the request parameters
		type UpdatePasswordRequest struct {
			Username    string `json:"username"`
			OldPassword string `json:"oldPassword"`
			NewPassword string `json:"newPassword"`
		}

		var updatePasswordRequest UpdatePasswordRequest

		err := json.NewDecoder(r.Body).Decode(&updatePasswordRequest)
		if err != nil {
			http.Error(w, "Invalid request parameters", http.StatusBadRequest)
			loggerInst.Log(logger.Error, "Error parsing update password request", map[string]interface{}{"error": err.Error()})
			return
		}

		//var employee *models.Employee
		// Call the UpdatePassword function to update the employee's password
		err = models.UpdatePassword(updatePasswordRequest.Username, updatePasswordRequest.OldPassword, updatePasswordRequest.NewPassword)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			loggerInst.Log(logger.Error, "Error updating password", map[string]interface{}{"error": err.Error()})
			return
		}

		// Respond with a success message
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Password updated successfully"))
		loggerInst.Log(logger.Info, "Update password request completed", map[string]interface{}{"response_code": http.StatusOK, "username": updatePasswordRequest.Username, "oldPassword": updatePasswordRequest.OldPassword, "newPassword": updatePasswordRequest.NewPassword})
	})

	// Define an endpoint to retrieve all logs
	http.HandleFunc("/logs", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			loggerInst.Log(logger.Info, "Logs request completed", map[string]interface{}{"response_code": http.StatusOK})
			return
		}

		// Retrieve all logs from the database
		logReader := logger.NewLogReader("./db/logs.json", logger.Info, "", time.Time{}, time.Time{})
		logs, err := logReader.ReadLogs()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error retrieving logs", map[string]interface{}{"error": err.Error()})
			return
		}

		// Convert the list of logs to JSON
		logsJSON, err := json.Marshal(logs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error converting logs to JSON", map[string]interface{}{"error": err.Error()})
			return
		}

		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response to the client
		w.Write(logsJSON)
		loggerInst.Log(logger.Info, "Logs request completed", map[string]interface{}{"response_code": http.StatusOK})
	})

	// Define an endpoint to retrieve logs of financial transactions during a specific time period
	http.HandleFunc("/financialLogs", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			loggerInst.Log(logger.Info, "Financial logs request completed", map[string]interface{}{"response_code": http.StatusOK})
			return
		}

		// Retrieve all logs from the database
		logReader := logger.NewLogReader("./db/logs.json", logger.Info, "Checkout", time.Time{}, time.Time{})
		logs, err := logReader.ReadLogs()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error retrieving financial logs", map[string]interface{}{"error": err.Error()})
			return
		}

		// Convert the list of logs to JSON
		logsJSON, err := json.Marshal(logs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error converting financial logs to JSON", map[string]interface{}{"error": err.Error()})
			return
		}

		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response to the client
		w.Write(logsJSON)
		loggerInst.Log(logger.Info, "Financial logs request completed", map[string]interface{}{"response_code": http.StatusOK})
	})

	// Define an endpoint to retrieve logs of inventory during a specific time period
	http.HandleFunc("/inventoryLogs", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			loggerInst.Log(logger.Info, "Inventory logs request completed", map[string]interface{}{"response_code": http.StatusOK})
			return
		}

		// Retrieve all logs from the database
		logReader := logger.NewLogReader("./db/logs.json", logger.Info, "Inventory", time.Time{}, time.Time{})
		logs, err := logReader.ReadLogs()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error retrieving inventory logs", map[string]interface{}{"error": err.Error()})
			return
		}

		// Convert the list of logs to JSON
		logsJSON, err := json.Marshal(logs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error converting inventory logs to JSON", map[string]interface{}{"error": err.Error()})
			return
		}

		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response to the client
		w.Write(logsJSON)
		loggerInst.Log(logger.Info, "Inventory logs request completed", map[string]interface{}{"response_code": http.StatusOK})
	})

	http.HandleFunc("/prescription", func(w http.ResponseWriter, r *http.Request) {
		loggerInst.Log(logger.Info, "Received add prescription request", map[string]interface{}{"request_method": r.Method, "request_path": r.URL.Path})
		enableCors(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			loggerInst.Log(logger.Info, "Add prescription request completed", map[string]interface{}{"response_code": http.StatusOK})
			return
		}

		// Parse the request parameters
		type AddPrescriptionRequest struct {
			ID           int     `json:"id"`
			Drug         string  `json:"drug"`
			Doses        int     `json:"doses"`
			Strength     string  `json:"strength"`
			Price        float64 `json:"price"`
			Doctor       string  `json:"doctor"`
			CustomerID   int     `json:"customerID"`
			IsFilled     string  `json:"isFilled"`
			PharmacistID int     `json:"pharmacistID"`
		}

		var addPrescriptionRequest AddPrescriptionRequest

		fmt.Println("Received JSON data:", r.Body)

		err := json.NewDecoder(r.Body).Decode(&addPrescriptionRequest)
		if err != nil {
			http.Error(w, "Invalid request parameters", http.StatusBadRequest)
			loggerInst.Log(logger.Error, "Error parsing add prescription request", map[string]interface{}{"error": err.Error()})
			return
		}

		// Call the NewPrescription function to add the prescription
		err = models.AddPrescription(addPrescriptionRequest.ID, addPrescriptionRequest.Drug, addPrescriptionRequest.Doses, addPrescriptionRequest.Strength, addPrescriptionRequest.Price, addPrescriptionRequest.Doctor, addPrescriptionRequest.CustomerID, addPrescriptionRequest.IsFilled, addPrescriptionRequest.PharmacistID)

		// TODO: Save the prescription to the database or perform any other necessary operations
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error adding prescription", map[string]interface{}{"error": err.Error()})
			return
		}

		// Respond with a success message
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Prescription added successfully"))
		loggerInst.Log(logger.Info, "Add prescription request completed", map[string]interface{}{"response_code": http.StatusOK})
	})

	http.HandleFunc("/medicines", func(w http.ResponseWriter, r *http.Request) {
		loggerInst.Log(logger.Info, "Received get medicines request", map[string]interface{}{"request_method": r.Method, "request_path": r.URL.Path})
		enableCors(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			loggerInst.Log(logger.Info, "Get medicines request completed", map[string]interface{}{"response_code": http.StatusOK})
			return
		}

		// Call the GetMedicine function to retrieve the medicines
		medicines, err := models.GetMedicine()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error retrieving medicines", map[string]interface{}{"error": err.Error()})
			return
		}

		// Convert the list of medicines to JSON
		medicinesJSON, err := json.Marshal(medicines)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			loggerInst.Log(logger.Error, "Error converting medicines to JSON", map[string]interface{}{"error": err.Error()})
			return
		}

		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response to the client
		w.Write(medicinesJSON)
		loggerInst.Log(logger.Info, "Get medicines request completed", map[string]interface{}{"response_code": http.StatusOK})
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
