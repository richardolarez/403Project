// api/authenticate.go

package api

import (
	"encoding/json"
	"net/http"

	accountmanager "github.com/richardolarez/403Project/internal/account_manager"
)

// AuthenticateEmployeeHandler handles requests to authenticate an employee login.
func AuthenticateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the username and password
	var requestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Authenticate the employee login
	employee, err := accountmanager.AuthenticateEmployee(requestBody.Username, requestBody.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Return the authenticated employee as JSON
	json.NewEncoder(w).Encode(employee)
}
